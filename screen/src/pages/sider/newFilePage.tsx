import {useRef, useState} from "react";
import {BankOutlined} from "@ant-design/icons";
import env from "@/env";
import {Button, Drawer, Space, Form, Input, Tag,Radio} from 'antd';
import axios from "axios";

function NewFilePage({messageApi,fileKey,setFileKey,setContent,updateTree,isOpen,closeDrawer}){
    const typeRef = useRef("目录或笔记");
    const [newName, setNewName] = useState("");
    const [extName, setExtName] = useState("");
    const onFinish = (values: any) => {
        let type = "f";
        if (extName === ""){type = "d";}
        new Promise<void>((resolve) => {
            axios.get(env.apiUrl + '/api/content/add?type='+type+'&key='+fileKey+newName+extName).then((resp) => {
                if (resp.data.code === 0){
                    messageApi.open({
                        type: 'success',
                        content: "新增成功："+fileKey+newName+extName,
                        duration: 1,
                    });
                    setFileKey(fileKey+newName+extName);
                    setContent("");
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
        closeDrawer();
        updateTree(fileKey+newName+extName);
    }

    //监听输入文件名变化
    const onNameChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        if (fileKey.endsWith("/")){
            setNewName(e.target.value);
        }else {
            setNewName("/"+e.target.value)
        }
    };
    //监听后缀名
    const onChange = (e: RadioChangeEvent) =>{
        setExtName(e.target.value);
        if (extName === ""){
            typeRef.current = "笔记";
        }else {
            typeRef.current = "目录";
        }
    }

    return(
        <Drawer title="新增目录或笔记"
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
                <Form.Item>
                    <Tag icon={<BankOutlined />}>{" 新增["+ typeRef.current + "] " + fileKey + newName + extName}</Tag>
                </Form.Item>
                <Form.Item>
                    <Radio.Group onChange={onChange} value={extName}>
                            <Space>
                                <Radio value={""}>目录</Radio>
                                <Radio value={".wg"}>笔记(富文本)</Radio>
                            </Space>
                    </Radio.Group>
                </Form.Item>
                <Form.Item<FieldType>
                    label={typeRef.current+"名"}
                    name="newName"
                    rules={[{ required: true}]}
                >
                    <Input onChange={onNameChange}/>
                </Form.Item>
                <Form.Item>
                    <Button type="primary" htmlType="submit">
                        新增{typeRef.current}
                    </Button>
                </Form.Item>
            </Form>
        </Drawer>)
};

export default NewFilePage;