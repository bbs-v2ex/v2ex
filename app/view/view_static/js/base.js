//将base64转换为blob
function dataURLtoFile(dataURI, type) {
    let binary = atob(dataURI.split(',')[1]);
    let array = [];
    for(let i = 0; i < binary.length; i++) {
        array.push(binary.charCodeAt(i));
    }
    return new Blob([new Uint8Array(array)], {type:type });
}


/**
 * 将以base64的图片url数据转换为Blob
 * @param urlData
 *            用url方式表示的base64图片数据
 */
function convertBase64UrlToBlob(urlData){

    var bytes=window.atob(urlData.split(',')[1]);        //去掉url的头，并转换为byte

    //处理异常,将ascii码小于0的转换为大于0
    var ab = new ArrayBuffer(bytes.length);
    var ia = new Uint8Array(ab);
    for (var i = 0; i < bytes.length; i++) {
        ia[i] = bytes.charCodeAt(i);
    }

    return new Blob( [ab] , {type : 'image/png'});
}

/**
 *  传入图片路径，返回base64
 * @param img
 * @returns {any}
 */

function getBase64(img){
    return  new Promise(function(reslove,reject) {
        try {
            function getBase64Image(img,width,height) {//width、height调用时传入具体像素值，控制大小 ,不传则默认图像大小
                var canvas = document.createElement("canvas");
                canvas.width = width ? width : img.width;
                canvas.height = height ? height : img.height;

                var ctx = canvas.getContext("2d");
                ctx.drawImage(img, 0, 0, canvas.width, canvas.height);
                var dataURL = canvas.toDataURL();
                return dataURL;
            }
            var image = new Image();
            image.crossOrigin = '';
            image.src = img;
            if(img){
                image.onload =function (){
                   return  reslove(getBase64Image(image));//将base64传给done上传处理
                };
            }
        }catch (e) {
            return reject(e)
        }
    })
}

/**
 * 监听加载
 * @param h
 * @returns {boolean}
 */

function bodyScrollLoad(h = 0){
    var pageHeight = Math.max(document.body.scrollHeight,document.body.offsetHeight);
    var viewportHeight = window.innerHeight ||
        document.documentElement.clientHeight ||
        document.body.clientHeight || 0;
    var scrollHeight = window.pageYOffset ||
        document.documentElement.scrollTop ||
        document.body.scrollTop || 0;
    return pageHeight - viewportHeight - scrollHeight < h;  // 通过 真实内容高度 - 视窗高度 - 上面隐藏的高度 < 20，作为加载的触发条件
}


var login_app = new Vue({
    el: '.vue-user-show',
    delimiters:['${','}'],
    data() {
        return {
            user_info:{}
        }
    },
    created(){
        this.user_info = getUserInfo();
        if (this.user_info.mid == undefined){
            this.user_info.mid = 0
        }
        console.log(this.user_info);
        document.querySelectorAll('.seo-html').forEach(function (z,i) {
            // console.dir(z)
            z.style.display="none";
        })
    },

    methods:{
        loginout(){
            post('/loginout').then(res =>{
                localStorage.removeItem(APITOKEN);
                window.location.href = "?"
            })
        },
        goMemberCentre(){
            window.location.href = "/_/member/z/";
        }
    }
});