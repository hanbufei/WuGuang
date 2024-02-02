import {Tag} from 'antd';
import {DeliveredProcedureOutlined} from "@ant-design/icons";

function XS_Tag({fileKey}){
    return (
        // <Tag icon={<DeliveredProcedureOutlined />} style={{position:"fixed",bottom:8}}>
        <Tag icon={<DeliveredProcedureOutlined />}>
            {fileKey}
        </Tag>
    )
};

export default XS_Tag;