import XS_Editor from "@/pages/content/tinymce";
import {FloatButton} from "antd";
import axios from "axios";
import env from "@/env";
import AlertLinkPage from "@/pages/content/alertLinkPage";
import {useRef,useState} from "react";
import {CheckCircleOutlined} from '@ant-design/icons';

function XS_Content({messageApi,fileKey,setFileKey,content,setContent,editorRef,saveContent}){
    const linksRef = useRef([]);//图片外链
    // 图片外链的抽屉
    const [isOpen, setOpen] = useState(false);
    const openDrawer = () => {
        setOpen(true);
    };
    const closeDrawer = () => {
        setOpen(false);
    };
    // 获得图片外链
    function getImageLinks(html) {
        let tmpLinks = []
        const regex = /<img[^>]+src="(http[^">]+)"/g;
        let match;
        while ((match = regex.exec(html)) !== null) {
            if (!match[1].includes(env.apiUrl)){
                tmpLinks.push(match[1]);
            }
        }
        linksRef.current = tmpLinks;
    }

    //点击保存：分析图片外链
    const doClick = ()=> {
        getImageLinks(editorRef.current.getContent());
        //没有外链，直接保存；有外链，打开抽屉，询问是否下载
        if (linksRef.current.length === 0) {
            saveContent();
        }else {
            openDrawer();
        }
    };

    return (
        <>
            <XS_Editor editorRef={editorRef}
                       content={content}
                       fileKey={fileKey}
                       setFileKey={setFileKey}
                       doClick={doClick}/>
            <FloatButton
                icon={<CheckCircleOutlined />}
                onClick={doClick}
                description="保存"
                shape="square"
                style={{bottom:440,backgroundColor: '#ffcccc',}}
            />
            <AlertLinkPage messageApi={messageApi}
                           links={linksRef.current}
                           editorRef={editorRef}
                           setContent={setContent}
                           saveContent={saveContent}
                           isOpen={isOpen}
                           closeDrawer={closeDrawer}/>
        </>
    )
};

export  default XS_Content;