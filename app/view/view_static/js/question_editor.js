var E = window.wangEditor;
var editor = new E('#f1-content');
editor.customConfig.menus =
    [
        'head',  // 标题
        'bold',  // 粗体
        'strikeThrough',  // 删除线
        'foreColor',  // 文字颜色
        'backColor',  // 背景颜色
        'link',  // 插入链接
        'list',  // 列表
        'justify',  // 对齐方式
        'quote',  // 引用
        'image',  // 插入图片
        'code',  // 插入代码
        'undo',  // 撤销
        'redo'  // 重复
    ];
editor.customConfig.uploadImgServer = UploadServer +'/upload_img_123';
editor.customConfig.uploadImgMaxLength = 1;
editor.customConfig.uploadFileName = 'file';

editor.customConfig.linkImgCallback = function (url) {
    console.log('copy',url) // url 即插入图片的地址
};

editor.customConfig.uploadImgHooks = {
    customInsert: function (insertImg, result, editor) {
        if (result.code !== 1){
            alert(result.message);
            return
        }
        var url = UploadServer + '/'+result.url;
        insertImg(url)
    }
};
editor.create();
var jiance = false;
setInterval(function () {
    check()
},500);
async function check() {
    if (jiance){
        return
    }
    jiance = true;

    let list = document.querySelectorAll('#f1-content img');

    for (item of  list){
        let src = item.src;
        if (src === ""){

            item.remove();
            continue
        }
        if (src.startsWith(UploadServer) ){
            continue;
        }
        if (src.indexOf("/static/tmp/") === -1 || src.startsWith('data:image/')){
            try {
                let  data  =  await post('/download_temp_img',{u:src});
                if (data.code === 1){
                    item.src = data.data;
                }else {

                    item.remove();
                }
            }catch (e) {

                item.remove();
            }
            continue
        }
        if (src.indexOf('/static/tmp/') > -1  ){
            //上传到服务器
            try {

                let base64_img =  await getBase64(src);

                var forms = new FormData();
                var configs = {
                    headers:{'Content-Type':'multipart/form-data'}
                };
                let blob= dataURLtoFile(base64_img,'image/jpeg');
                let fileOfBlob = new File([blob], new Date()+'.jpg'); // 重命名了
                forms.append("file", fileOfBlob);

                let dddd = await axios.post(UploadServer+"/upload_img_123",forms ,configs);
                let data = dddd.data;
                if (data.code === 1){
                    item.outerHTML = `<img src="${UploadServer}/${data.url}"/>`
                }else{

                    item.remove();
                }
            }catch (e) {
                console.log(e);

                item.remove();

            }
        }
    }
    jiance = false;
}