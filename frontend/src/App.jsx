import "./App.css"
import React, {useEffect, useState} from "react"
import {Application} from "./rpc.js"

function blobToBase64(blob) {
    return new Promise((resolve, _) => {
        const reader = new FileReader()
        reader.onloadend = () => resolve(reader.result)
        reader.readAsDataURL(blob)
    })
}

export default function () {
    let [imgUrl, setImgUrl] = useState("")

    useEffect(function () {
        let imgBase64 = localStorage.getItem("paste-img-base64")
        if (imgBase64 === null) {
            Application.DefaultSizeMainWindow()
        } else {
            setBase64Img(imgBase64)
        }

        window.runtime.EventsOn("set_base64_img", data => setBase64Img(data))
        window.runtime.EventsOn("clean_img", () => {
            setImgUrl("")
            localStorage.removeItem("paste-img-base64")
        })

        document.addEventListener("paste", onPaste, false)
        return () => document.removeEventListener("paste", onPaste)
    }, [])

    const setBlobImg = blob => {
        setImgUrl(URL.createObjectURL(blob))
        cacheBlob(blob)
    }

    const setBase64Img = imgBase64 => {
        fetch(imgBase64).then(res => res.blob()).then(blob => setBlobImg(blob))
    }

    const cacheBlob = blob => {
        blobToBase64(blob).then(res => {
            localStorage.setItem("paste-img-base64", res)
        })
    }

    const onPaste = e => {
        const cbd = e.clipboardData
        for (let i = 0; i < cbd.items.length; i++) {
            let item = cbd.items[i]
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

    const onImgLoad = e => Application.ResizeMainWindow(e.target.width, e.target.height)

    return (
        <div>
            <div className="draggable" style={{"--wails-draggable": "drag"}}></div>
            <img style={{display: "none"}} src={imgUrl} onLoad={onImgLoad} alt=""/>
            {
                imgUrl === "" ?
                    <div className="notice">
                        请在应用内执行粘贴(Command + V)
                        <br/>
                        或使用菜单中的 File -> Load Image
                        <br/>
                    </div> :
                    <img className="paste-img" src={imgUrl} alt="粘贴图片"/>
            }
        </div>
    )
}
