// var myModal =, {
//     keyboard: false
// })
var app = new Vue({
    el: '#vue-app',
    delimiters: ['${', '}'],
    data() {
        return {
            comment: comment,
            user_info: '',
            edit_root: {
                txt: '',
                did: '',
            },
            wait_loading: false,
            ajax_message: '',
            //评论
            discuss: {
                edit_child: {
                    rid: '000000000000000000000000',
                    pid: '000000000000000000000000',
                    txt: '',
                },
                ajax_message:'',
                wait_loading: false,
                message: '',
                list: [],
            }
        }
    },
    created() {
        const userInfo = getUserInfo();
        try {
            this.user_info = userInfo
        } catch (e) {

        }
        console.log(userInfo)
    },
    methods: {
        discuss_show(_id) {
            this.discuss.message = '加载中';
            this.discuss.edit_child.rid = _id;
            discuss_modal.show();
            this.discuss_reload()
        },
        discuss_reload(){
          post('/article/comment_child_list',{rid:this.discuss.rid}).then(res => {
              if (res.data.length === 0){
                  this.discuss.message = '无评论，快去添加你的神评吧'
              }
              this.discuss.list = res.data
          }).catch(err =>{
              this.discuss.message = '接口错误,拉取失败'
          })
        },
        zan(index, _id) {
            let add_zan = false;
            try {
                add_zan = this.comment[index].zan_user.includes(this.user_info.mid);
            } catch (e) {
                add_zan = false
            }
            let zan_url = '/article/zan_add';
            if (add_zan) {
                zan_url = '/article/zan_del'
            }

            post(zan_url, {'_id': this.comment[index]._id}).then(res => {
                if (res.code === 1) {
                    this.comment[index].zan_user = res.data;
                    this.comment[index].zan = res.data.length;
                }
            })
        },
        is_user(user_list = []) {
            try {
                return user_list.includes(this.user_info.mid)
            } catch (e) {
                return false
            }

        },
        j(u) {
            window.location.href = u
        },
        just_login() {
            document.querySelector('#navbarCollapse > div > ul > li:nth-child(1) > a').click()
        },
        submit_comment_root() {
            this.wait_loading = true;
            post('/article/comment_root_add', this.edit_root).then(res => {
                this.ajax_message = res.message;
                if (res.code === 1) {

                }
            }).catch(err => {
                this.ajax_message = '接口错误';

            }).finally(() => {

                this.wait_loading = false

            })
        },

        submit_comment_child() {
            this.discuss.wait_loading = true;
            post('/article/comment_child_add', this.discuss.edit_child).then(res => {
                this.discuss.ajax_message = res.message;
                if (res.code === 1) {
                    //清空输入框
                    this.discuss.txt = '';
                    //重新拉取数据
                    this.discuss_reload()
                }
            }).catch(err => {
                this.discuss.ajax_message = '接口错误';

            }).finally(() => {

                this.discuss.wait_loading = false

            })
        }
    }
});