var app = new Vue({
    el: '#vue-app',
    delimiters: ['${', '}'],

    data() {
        return {
            collect: {
                status: false,
                txt: '收藏',
                ajax_txt: '',
            },
            load_data: {
                stop: false,
                wait: false,
                message: '点击,加载更多... ',
            },
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
                active_reply_user: {},
                edit_child: {
                    rid: '000000000000000000000000',
                    pid: '000000000000000000000000',
                    txt: '',
                },
                ajax_message: '',
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
        post('/member/is_collect', {did: DID}).then(res => {
            this.collect.status = res.data;
            if (this.collect.status) {
                this.collect.txt = '已收藏'
            }
        });
        setTimeout(function () {
            post('/show',{did: DID})
        },1500)
    },
    methods: {

        collect_toggle() {
            if (this.collect.status) {
                console.log("取消收藏")
                post('/member/collect_del', {did: DID}).then(res => {
                    if (res.code === 1) {

                        this.collect.status = false;
                        this.collect.txt = '收藏';
                        this.collect.ajax_txt = '';
                    }
                });
            } else {
                post('/member/collect_add', {did: DID}).then(res => {
                    if (res.code === 1) {
                        if (res.data) {
                            this.collect.status = true;
                            this.collect.txt = '已收藏';
                            this.collect.ajax_txt = '';
                        } else {
                            this.collect.ajax_txt = res.message;
                        }
                    } else {
                        this.collect.ajax_txt = res.message;
                    }
                });
                console.log("添加收藏")
            }
        },

        //加载更多数据
        load_more() {
            if (this.load_data.stop || this.load_data.wait) {
                return
            }
            this.load_data.wait = true;
            let rid = this.comment.slice(-1)[0]._id;
            post('/article/comment_root_list', {did: this.edit_root.did, rid: rid}).then(res => {
                if (res.code === 1) {
                    this.load_data.message = '点击加载数据...';
                    this.comment.push(...res.data);
                    if (res.data.length < 10) {
                        this.load_data.message = '已加载所有评论';
                        this.load_data.stop = true
                    }
                } else {
                    this.load_data.message = '加载失败刷新重新'
                }
            }).finally(() => {
                this.load_data.wait = false
            });
            console.log("load 更多数据")
        },

        active_reply(item) {
            this.discuss.active_reply_user = item;
        },
        active_reply_close() {
            this.discuss.active_reply_user = {};
        },
        discuss_avatar(src, w) {
            return src + "?&w=" + w
        },
        discuss_show(_id) {
            this.discuss.message = '加载中';
            this.discuss.edit_child.rid = _id;
            discuss_modal.show();
            this.discuss_reload()
        },
        discuss_reload() {
            post('/article/comment_child_list', {rid: this.discuss.edit_child.rid}).then(res => {
                if (res.code === 1) {
                    this.discuss.message = '';
                    if (res.data.length === 0)
                        this.discuss.message = '无评论，快去添加你的神评吧'

                } else {
                    this.discuss.message = res.message
                }

                this.discuss.list = res.data
            }).catch(err => {
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
        m(u) {
            window.location.href = u_list.member + u
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

            try {
                if (this.discuss.active_reply_user._id.length === 24) {
                    this.discuss.edit_child.pid = this.discuss.active_reply_user._id
                }
            } catch (e) {

            }

            post('/article/comment_child_add', this.discuss.edit_child).then(res => {
                this.discuss.ajax_message = res.message;
                if (res.code === 1) {
                    //清空输入框
                    this.discuss.edit_child.txt = '';
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