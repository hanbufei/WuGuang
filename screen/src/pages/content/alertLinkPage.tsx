import {Button, Drawer, Form, Space,List} from "antd";
import Base64 from "base-64";
import env from "@/env";
import axios from "axios";

function AlertLinkPage({messageApi,links,editorRef,setContent,saveContent,isOpen,closeDrawer}){
    // 替换为本地链接: 原链接的Base64编码，拼接到本地图片文件夹download
    const onFinish = () => {
        let text = editorRef.current.getContent();
        for (let i = 0; i < links.length; i++) {
            downloadImg(links[i]);
            text = text.replaceAll(links[i],"download/"+Base64.encode(links[i]));
        }
        editorRef.current.setContent(text);
        setContent(text);
        doSimpleSave();
    }
    const doSimpleSave = () => {
        saveContent();
        closeDrawer();
    }

    //下载图片
    const downloadImg  = (imgUrl)  => {
        new Promise<void>((resolve) => {
            axios.post(env.apiUrl + '/api/content/downloadImg',{
                ImgUrl: imgUrl,
            },{
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded'
                }}).then((resp) => {
                if (resp.data.code === 0){
                    messageApi.open({
                        type: 'success',
                        content: '图片下载成功'+imgUrl,
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
        <Drawer title="笔记中存在图片外链，是否下载到本地？"
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
                    <List
                        bordered
                        dataSource={links}
                        renderItem={(item) => <List.Item>{item}</List.Item>}
                    />
                </Form.Item>
                <Space>
                    <Form.Item>
                        <Button type="primary" htmlType="submit">
                            下载图片，并替换为本地链接
                        </Button>
                    </Form.Item>
                    <Form.Item>
                        <Button onClick={doSimpleSave}>
                            不下载，保留原文中外链
                        </Button>
                    </Form.Item>
                </Space>
            </Form>
        </Drawer>
    )
}

export default AlertLinkPage;