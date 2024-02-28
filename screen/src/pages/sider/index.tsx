import {useEffect, useState} from "react";
import {Tree, FloatButton, Popconfirm} from 'antd';
import axios from "axios";
import env from "@/env"
import NewFilePage from "@/pages/sider/newFilePage";
import {PlusCircleOutlined,CloseCircleOutlined} from '@ant-design/icons';

function XS_Sider({messageApi,fileKey,setFileKey,treeData,setTreeData,updateSider,setContent,editorRef}){
    // 新增笔记的抽屉
    const [isOpen, setOpen] = useState(false);
    const openDrawer = () => {
        setOpen(true);
    };
    const closeDrawer = () => {
        setOpen(false);
    };
    //初始化顶级菜单
    useEffect(() => {
        axios.get(env.apiUrl + '/api/menu/init').then((resp) => {
            if (resp.data.code === 0){
                setTreeData([{title:resp.data.data.Label,key:"/",children:resp.data.data.MenuInitData}])
            }else {
                messageApi.open({
                    type: 'error',
                    content: resp.data.message,
                    duration: 2,
                });
            }
        })
    }, []);

    //动态加载选中的菜单
    const onLoadData = ({ key, children }: any) =>
        new Promise<void>((resolve) => {
            if (children) {resolve();return;}
            axios.get(env.apiUrl + '/api/menu/list?key=' + key).then((resp) => {
                if (resp.data.code === 0){
                    updateSider(key, resp.data.data.MenuListData)
                }else {
                    messageApi.open({
                        type: 'error',
                        content: resp.data.message,
                        duration: 2,
                    });
                }
            })
            resolve()
        });

    const onSelect = (selectedKeys,info) =>{
        setFileKey(info.node.key);
        //获取文件数据
        if (info.node.isLeaf){
            new Promise<void>((resolve) => {
                axios.get(env.apiUrl + '/api/content/get?key=' + info.node.key).then((resp) => {
                    if (resp.data.code === 0){
                        editorRef.current.setContent(resp.data.data.Content);
                        setContent(resp.data.data.Content);
                    }else {
                        messageApi.open({
                            type: 'error',
                            content: resp.data.message,
                            duration: 2,
                        });
                    }
                })
                resolve()
            });
        }
    };

    //更新树
    const updateTree = (key) =>
        new Promise<void>((resolve) => {
            //获取当前路径
            key = key.substring(0,key.lastIndexOf("/"))
            //如果是根路径，则初始化
            if (key === ""){
                axios.get(env.apiUrl + '/api/menu/init').then((resp) => {
                    if (resp.data.code === 0){
                        setTreeData([{title:resp.data.data.Label,key:"/",children:resp.data.data.MenuInitData}])
                    }else {
                        messageApi.open({
                            type: 'error',
                            content: resp.data.message,
                            duration: 2,
                        });
                    }
                })
                resolve();
            }
            //更新
            axios.get(env.apiUrl + '/api/menu/list?key=' + key).then((resp) => {
                if (resp.data.code === 0){
                    updateSider(key, resp.data.data.MenuListData)
                }else {
                    messageApi.open({
                        type: 'error',
                        content: resp.data.message,
                        duration: 2,
                    });
                }
            })
            resolve()
        });

    // 删除笔记
    const deleteContent  = ()  => {
        new Promise<void>((resolve) => {
            axios.get(env.apiUrl + '/api/content/delete?key='+fileKey).then((resp) => {
                if (resp.data.code === 0){
                    messageApi.open({
                        type: 'success',
                        content: '删除 '+fileKey+' 成功',
                        duration: 1,
                    });
                    //更新树
                    updateTree(fileKey);
                }else {
                    messageApi.open({
                        type: 'error',
                        content: resp.data.message,
                        duration: 2,
                    });
                }
            })
            resolve()
        });
    }

    return (
        <>
                <FloatButton
                    icon={<PlusCircleOutlined/>}
                    description="新增"
                    onClick={openDrawer}
                    shape="square"
                    style={{ left: 10,bottom:500,backgroundColor: '#ffcccc',}}
                />
                <NewFilePage messageApi={messageApi}
                             fileKey={fileKey}
                             setFileKey={setFileKey}
                             setContent={setContent}
                             updateTree={updateTree}
                             isOpen={isOpen}
                             closeDrawer={closeDrawer}/>
                <Popconfirm title={"确定删除:"+fileKey+" ?"} okText="Yes" cancelText="No" onConfirm={deleteContent}>
                    <FloatButton
                        icon={<CloseCircleOutlined/>}
                        description="删除"
                        shape="square"
                        danger
                        style={{ left: 10,bottom:440,backgroundColor: '#ffcccc',}}
                    />
                </Popconfirm>
            <Tree style={{ left: 20}}
                  showLine={true}
                  onSelect={onSelect}
                  loadData={onLoadData}
                  treeData={treeData}
                  defaultExpandAll={true}
            />
        </>
    );
}

export default XS_Sider;