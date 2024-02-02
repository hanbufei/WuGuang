import React, {ReactNode, useRef, useState} from 'react';
import { Layout,message } from 'antd';
import XS_Header from "@/pages/header";
import XS_Sider from "@/pages/sider";
import XS_Content from "@/pages/content";
import XS_Tag from "@/pages/footer/tag";
import axios from "axios";
import env from "@/env";

const { Header, Footer, Sider, Content } = Layout;

interface DataNode {
    title: string;
    key: string;
    icon?:ReactNode;
    isLeaf?: boolean;
    children?: DataNode[];
}

const updateTreeData = (list: DataNode[], key: React.Key, children: DataNode[]): DataNode[] =>
    list.map((node) => {
        if (node.key === key) {
            return {
                ...node,
                children,
            };
        }
        if (node.children) {
            return {
                ...node,
                children: updateTreeData(node.children, key, children),
            };
        }
        return node;
    });


function App() {
    const [messageApi, contextHolder] = message.useMessage();//全局消息
    const [fileKey, setFileKey] = useState("/"); // 当前选择的节点文件路径
    const [content, setContent] = useState("");//当前显示的内容
    const [treeData, setTreeData] = useState<DataNode[] | []>([{title:"/",key:"/",isLeaf:true}]);//菜单数据
    const editorRef = useRef(null);//编辑器

    //更新指定节点的tree数据
    const updateSider =  (key,data)=>{
        setTreeData((origin) =>
            updateTreeData(origin, key, data),
        );
    }

    //保存笔记
    const saveContent  = ()  => {
        new Promise<void>((resolve) => {
            axios.post(env.apiUrl + '/api/content/save',{
                Key: fileKey,
                Content: editorRef.current.getContent(),
            },{
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded'
                }}).then((resp) => {
                if (resp.data.code === 0){
                    messageApi.open({
                        type: 'success',
                        content: '保存文件成功',
                        duration: 1,
                    });
                }else {
                    messageApi.open({
                        type: 'error',
                        content: resp.data.message,
                        duration: 2,
                    });
                }
            })
            resolve();
        });
    }
        
    return (
            <Layout>
                {contextHolder}
                <Sider width="25%" style={siderStyle} theme={"light"}>
                    <XS_Sider messageApi = {messageApi}
                              fileKey = {fileKey}
                              setFileKey = {setFileKey}
                              treeData = {treeData}
                              setTreeData = {setTreeData}
                              updateSider = {updateSider}
                              setContent = {setContent}
                              editorRef={editorRef}/>
                </Sider>
                    <Layout>
                        <Header style={headerStyle} theme={'light'}>
                            <XS_Header messageApi = {messageApi}
                                       setFileKey = {setFileKey}
                                       setContent = {setContent}
                                       setTreeData = {setTreeData}
                                       saveContent = {saveContent}/>
                        </Header>
                        <XS_Tag fileKey = {fileKey}/>
                        <Content style={contentStyle}>
                            <XS_Content messageApi = {messageApi}
                                        fileKey={fileKey}
                                        setFileKey = {setFileKey}
                                        content={content}
                                        setContent={setContent}
                                        editorRef={editorRef}
                                        saveContent={saveContent}
                            />
                        </Content>
                        {/*<Footer style={footerStyle}>*/}
                        {/*    <XS_Footer/>*/}
                        {/*</Footer>*/}
                        {/*<XS_Tag fileKey = {fileKey}/>*/}
                    </Layout>
            </Layout>
    );
}

const headerStyle: React.CSSProperties = {
    overflow: 'auto',
    backgroundColor: '#fff',
    height: '8vh',
};

const contentStyle: React.CSSProperties = {
    backgroundColor: '#FFFFFF',
};
const siderStyle: React.CSSProperties = {
    overflow: 'auto',
    height: '100vh',
    left: 0,
    top: 0,
};

const footerStyle: React.CSSProperties = {
    overflow: 'auto',
    color: '#fff',
    height: '8vh',
};


export default App;
