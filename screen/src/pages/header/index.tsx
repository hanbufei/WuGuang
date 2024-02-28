import {useState} from "react";
import {Space, Button,Popover} from 'antd';
import SetRootPathPage from "@/pages/header/setRootPathPage";
import iconPng from "@/assets/Icon.png"
import Md2html from "@/pages/header/md2html";

function XS_Header({messageApi,setFileKey,setContent,setTreeData,saveContent,changeToMd}){
    // 更换笔记本的抽屉
    const [isOpen, setOpen] = useState(false);
    const openDrawer = () => {
        setOpen(true);
    };
    const closeDrawer = () => {
        setOpen(false);
    };

    // md2html的抽屉
    const [isMdOpen, setMdOpen] = useState(false);
    const openMdDrawer = () => {
        setMdOpen(true);
    };
    const closeMdDrawer = () => {
        setMdOpen(false);
    };

    return (
        <Space>
            <Popover content={"那些你记录下的，像刻在雾里的光，每一缕都相伴相随^_^"} title="雾光笔记">
                <img src={iconPng} width={100}  align='right'/>
            </Popover>
            <Button shape="round" onClick={openDrawer} type="primary"> 切换笔记本</Button>
            <Button shape="round" onClick={changeToMd} type="primary">复制为MarkDown</Button>
            <Button shape="round" onClick={openMdDrawer} type="primary">MarkDown转换为html</Button>
            <Button shape="round" onClick={saveContent} type="primary" danger>保存笔记</Button>
            <SetRootPathPage messageApi = {messageApi}
                             setFileKey = {setFileKey}
                             setContent = {setContent}
                             setTreeData = {setTreeData}
                             isOpen = {isOpen}
                             closeDrawer = {closeDrawer}/>
            <Md2html messageApi = {messageApi}
                     isMdOpen = {isMdOpen}
                     closeMdDrawer = {closeMdDrawer}/>
        </Space>
    )
}

export default XS_Header;