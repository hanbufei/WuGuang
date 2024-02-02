import axios from "axios";
import {useEffect, useState} from "react";
import env from "@/env"
import {Space, Button, Drawer, Form, Input} from 'antd';

function SetRootPathPage({messageApi,setFileKey,setContent,setTreeData,isOpen,closeDrawer}){
    const [rootPath, setRootPath] = useState("");//笔记本绝对路径
    //获取笔记本绝对路径
    useEffect(() => {
        axios.get(env.apiUrl + '/api/storage/getRootPath').then((resp) => {
            if (resp.data.code === 0){
                setRootPath(resp.data.data.RootPath);
            }else {
                messageApi.open({
                    type: 'error',
                    content: resp.data.message,
                    duration: 2,
                });
                // closeDrawer();
            }
        })
    }, []);

    const onFinish = (values: any) => {
        // 重置数据
        setTreeData([{title:"/",key:"/",isLeaf:true}]);
        setFileKey("/");
        setContent("");
        new Promise<void>((resolve) => {
            axios.get(env.apiUrl + '/api/storage/setRootPath?rootPath=' + values.rootPath).then((resp) => {
                if (resp.data.code === 0){
                    messageApi.open({
                        type: 'success',
                        content: "设置笔记本路径："+ values.rootPath,
                        duration: 1,
                    });
                    //重置顶级菜单
                    axios.get(env.apiUrl + '/api/menu/init').then((resp) => {
                        if (resp.data.code === 0){
                            setTreeData([{title:resp.data.data.Label,key:"/",children:resp.data.data.MenuInitData}])
                        }else {
                            return Promise.reject({
                                status: resp.status,
                            })
                        }
                    })
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
        closeDrawer();
    };

    return (
        <Drawer title="选择笔记本"
                width={720}
                onClose={closeDrawer}
                open={isOpen}
                styles={{
                    body: {
                        paddingBottom: 80,
                    },
                }}
                extra={
                    <Space>
                        <Button onClick={closeDrawer}>Cancel</Button>
                    </Space>
                }>
            <Form layout="vertical" hideRequiredMark onFinish={onFinish}>
                <Form.Item<FieldType>
                    label="笔记本路径"
                    name="rootPath"
                    rules={[{ required: true}]}
                >
                    <Input placeholder={"当前路径为:"+ rootPath}/>
                </Form.Item>
                <Form.Item>
                    <Button type="primary" htmlType="submit">
                        更换笔记本
                    </Button>
                </Form.Item>
            </Form>
        </Drawer>
    )
}

export default SetRootPathPage;