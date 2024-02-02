import ReactQuill from "react-quill";
import "react-quill/dist/quill.snow.css";

// snow 或 bubble
function XS_Editor({contentRef,content,setContent}){
    const doChange = (value: string)=>{
        contentRef.current = value;
        setContent(value);
    };
    return (
        <ReactQuill placeholder="请输入..." theme="snow" value={content} onChange={doChange} style={{
            backgroundColor:'#f1e5d2',
            borderColor:'black',
            minHeight:'96%',
        }}/>
    )
}

export default XS_Editor;