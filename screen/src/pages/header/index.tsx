import {useState} from "react";
import {Space, Button,Popover} from 'antd';
import SetRootPathPage from "@/pages/header/setRootPathPage";
import iconPng from "@/assets/Icon.png"

function XS_Header({messageApi,setFileKey,setContent,setTreeData,saveContent}){
    // 更换笔记本的抽屉
    const [isOpen, setOpen] = useState(false);
    const openDrawer = () => {
        setOpen(true);
    };
    const closeDrawer = () => {
        setOpen(false);
    };

    return (
        <Space>
            <Popover content={"那些你记录下的，像刻在雾里的光，每一缕都相伴相随^_^"} title="雾光笔记">
                <img src={iconPng} width={100}  align='right'/>
            </Popover>
            <Button shape="round" onClick={openDrawer} type="primary"> 切换笔记本</Button>
            <Button shape="round" onClick={saveContent} type="primary" danger>保存笔记</Button>
            <SetRootPathPage messageApi = {messageApi}
                             setFileKey = {setFileKey}
                             setContent = {setContent}
                             setTreeData = {setTreeData}
                             isOpen = {isOpen}
                             closeDrawer = {closeDrawer}/>
        </Space>
    )
}

export default XS_Header;