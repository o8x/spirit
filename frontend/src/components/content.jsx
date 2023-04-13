import React, {useEffect, useState} from "react"
import {Application} from "../rpc.js"

export default props => {
    if (props.url !== "") {
        return <img className="paste-img" src={props.url} alt="粘贴图片" style={{
            width: props.width ? props.width : 350, height: props.height ? props.height : 350,
        }}/>
    }

    let [style, setStyle] = useState({
        fontSize: "18px",
    })


    useEffect(() => {
        if (props.text !== "") {
            if (props.text.length <= 30) {
                setStyle({fontSize: "35px"})
                Application.ResizeMainWindow(500, 150)
            }

            if (props.text.length <= 15) {
                Application.ResizeMainWindow(500, 100)
            }

            if (props.text.length > 30) {
                setStyle({fontSize: "18px"})
                Application.DefaultSizeMainWindow()
            }
        }
    }, [props.text])

    if (props.text !== "") {
        return <div className="paste-text" style={style}>
            {props.text}
        </div>
    }

    return <div className="notice">
        请在应用内执行粘贴 (Ctrl + V)
        <br/>
        图像同时支持手动操作
        <br/>
        输入图像：Content -> Load Image
        <br/>
        清除当前内容：Content -> Clean
    </div>
}
