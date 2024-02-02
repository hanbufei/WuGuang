import React, { useRef } from 'react';
import { Editor } from '@tinymce/tinymce-react';

export  function XS_Editor({editorRef,content,fileKey,setFileKey,doClick}) {
    // const editorRef = useRef(null);
    // const keyRef = useRef(null);
    // const doSave  =  () => {
    //     setFileKey(keyRef.current);
    //     contentRef.current = editorRef.current.getContent();
    //     doClick();
    // }
    return (
        <>
            <Editor
                tinymceScriptSrc={'/tinymce/tinymce.min.js'}
                onInit={(evt, editor) => {
                    editorRef.current = editor;
                }}
                initialValue={content}
                init={{
                    selector: '#textarea',
                    skin: 'oxide-dark',
                    language:'zh-Hans',
                    height: '96%',
                    resize:false,
                    menubar: false,
                    plugins: [
                        'lists','advlist','autosave', 'charmap', 'code', 'codesample', 'hr', 'print', 'searchreplace', 'toc', 'wordcount','quickbars','visualblocks'
                    ],
                    toolbar: 'styleselect fontsizeselect bold italic forecolor bullist|blocks|hr charmap codesample toc|blocks|' +
                        'print code searchreplace|blocks|',
                    content_style: 'body { font-family:Helvetica,Arial,sans-serif; font-size:14px}',
                    save_enablewhendirty: false,
                    // save_onsavecallback: function () { doSave()}
                }}
            />
        </>
    );
}
export default XS_Editor;