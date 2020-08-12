var app = new Vue({
    el: '#vue-app',
    delimiters: ['${', '}'],

    data() {
        return {
            next: '',
            message: '',
            wait: false,
            stop: false,
            load_list: [],
        }
    },
    created() {
        this.load_more()
    },
    methods: {
        load_more() {
            var that = this;
            document.onscroll = function () {
                if (that.wait || that.stop || that.next === '') {
                    return false
                }
                if (!bodyScrollLoad(20)) {
                    return false
                }
                that.wait = true;
                that.message = '加载中';
                get(that.next + "?xx=nohead", {}, true).then(res => {
                    if (res.code == 1) {
                        that.message = '';
                        that.load_list.push(res.data);
                        that.next = res.next;
                        if (that.next === '') {
                            that.message = '没有更多数据了';
                            that.stop = true
                        }
                    }
                }).finally(() => {
                    that.wait = false
                })

            }
        }
    }
});