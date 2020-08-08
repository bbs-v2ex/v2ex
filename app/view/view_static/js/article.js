var app = new Vue({
    el: '#vue-app',
    delimiters:['${','}'],
    data() {
        return {
            user_info:'',
            edit_root:{
                txt:'',
                did:'',
            },
            wait_loading:false,
            ajax_message:'',
        }
    },
    created(){
        const userInfo = getUserInfo();
        try {
            this.user_info = userInfo
        }catch (e) {

        }
        console.log(userInfo)
    },
    methods: {
        just_login(){
            document.querySelector('#navbarCollapse > div > ul > li:nth-child(1) > a').click()
        },
        submit_comment_root(){
            // if (this.edit_root.length < 10){
            //     this.ajax_message = '不要灌水';
            //     return
            // }
            this.wait_loading = true;
            post('/article/comment_root_add',this.edit_root).then(res => {
                this.ajax_message = res.message;
                if (res.code ===1 ){

                }
            }).catch(err =>{
                this.ajax_message = '接口错误';

            }).finally(() => {

                this.wait_loading = false

            })
        }
    }
});