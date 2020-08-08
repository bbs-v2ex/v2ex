var app = new Vue({
    el: '#vue-app',
    delimiters:['${','}'],
    data() {
        return {
            user_info:'',
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

    }
});