// function fuzhi() {
//     try {
//         document.querySelectorAll('span[data-zan-user]').forEach(function (z,i) {
//             let v =  JSON.parse(z.getAttribute('data-zan-user'));
//             if (v.length > 0){
//                 console.log(v)
//             }
//         })
//     }catch (e) {
//
//     }
//
// }
//
// fuzhi();
var app = new Vue({
    el: '#vue-app',
    delimiters:['${','}'],
    data() {
        return {
            comment:comment,
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
        zan(_id){
            console.log(_id)
        },
        is_user(user_list =[]){
           return user_list.includes(this.user_info.mid)
        },
        j(u){
            window.location.href = u
        },
        just_login(){
            document.querySelector('#navbarCollapse > div > ul > li:nth-child(1) > a').click()
        },
        submit_comment_root(){
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