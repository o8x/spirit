import "./App.css"
import React, {useEffect, useState} from "react"
import {Application} from "./rpc.js"
import Content from "./components/content.jsx"

function blobToBase64(blob) {
    return new Promise((resolve, _) => {
        const reader = new FileReader()
        reader.onloadend = () => resolve(reader.result)
        reader.readAsDataURL(blob)
    })
}

export default function () {
    let [text, setText] = useState("")
    let [imgUrl, setImgUrl] = useState("")
    let [imgSize, setImgSize] = useState([350, 350])
    let [windowSize, setWindowSize] = useState([350, 350])

    useEffect(function () {
        let imgBase64 = localStorage.getItem("paste-img-base64")
        if (imgBase64 === null) {
            Application.DefaultSizeMainWindow()
        } else {
            setBase64Img(imgBase64)
            // 有图像时才有必要恢复尺寸
            let w = localStorage.getItem("window-width")
            let h = localStorage.getItem("window-height")
            if (w !== null && h !== null) {
                Application.ResizeMainWindow(Number(w), Number(h))
            }
        }

        let textCache = localStorage.getItem("paste-text-cache")
        if (textCache !== null) {
            setText(textCache)
        }


        window.runtime.EventsOn("set_base64_img", data => {
            cleanBlob()
            setBase64Img(data)
        })
        window.runtime.EventsOn("clean_img", cleanBlob)

        onResize()
        document.addEventListener("paste", onPaste, false)
        window.onresize = onResize

        return () => {
            window.onresize = () => null
            document.removeEventListener("paste", onPaste)
        }
    }, [])

    const onResize = () => setImgSize(v => {
        let w = v[0]
        let h = v[1]
        let nw = window.innerWidth
        // 按照宽度缩放的比例缩放高度，实现等比缩放
        let nh = h * (nw / w)
        setWindowSize([nw, nh])
        Application.ResizeMainWindow(nw, nh)
        localStorage.setItem("window-width", `${nw}`)
        localStorage.setItem("window-height", `${nh}`)

        return v
    })

    const setBlobImg = blob => {
        setImgUrl(URL.createObjectURL(blob))
        blobToBase64(blob).then(res => {
            localStorage.setItem("paste-img-base64", res)
        })
    }

    const setBase64Img = imgBase64 => {
        fetch(imgBase64).then(res => res.blob()).then(blob => {
            setImgUrl(URL.createObjectURL(blob))
            localStorage.setItem("paste-img-base64", imgBase64)
        })
    }

    const cleanBlob = () => {
        setImgUrl("")
        setText("")
        setImgSize([350, 350])
        setWindowSize([350, 350])

        localStorage.removeItem("paste-img-base64")
        localStorage.removeItem("paste-text-cache")
        localStorage.removeItem("window-width")
        localStorage.removeItem("window-height")
    }

    const onPaste = e => {
        const cbd = e.clipboardData
        for (let i = 0; i < cbd.items.length; i++) {
            let item = cbd.items[i]
            if (item.kind === "string" && item.type === "text/plain") {
                item.getAsString(data => {
                    cleanBlob()
                    setText(data)
                    localStorage.setItem("paste-text-cache", data)
                })
            }

            if (item.kind === "file") {
                let blob = item.getAsFile()
                if (blob.size === 0) {
                    return
                }

                setBlobImg(blob)
                return
            }
        }
    }

    const onImgLoad = e => {
        let w = e.target.width / 2
        let h = e.target.height / 2

        setImgSize([w, h])


        if (localStorage.getItem("window-width") === null && localStorage.getItem("window-height") === null) {
            Application.ResizeMainWindow(w, h)
        }
    }

    return (<div className="main">
        <div className="draggable" style={{"--wails-draggable": "drag"}}></div>
        <img style={{display: "none"}} src={imgUrl} onLoad={onImgLoad} alt=""/>
        <Content url={imgUrl} text={text} width={windowSize[0]} height={windowSize[1]}/>
    </div>)
}
