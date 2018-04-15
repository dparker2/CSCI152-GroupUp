(function() {
    const Login = {
        template: '#tmpl-login',
        data: function() {
            return {
                username: "",
                password: "",
            }
        },
        methods: {
            login: function() {
                this.$http.post('/portal', { // Data
                    username: this.username,
                    password: this.password
                }, { // Config
                    emulateJSON: true
                }).then(function(response) { // Success
                    console.log(response); 
                }, function(response) { // Error
                    console.log(response); 
                });
            }
        }
    }

    const Register = {
        template: '#tmpl-register',
        data: function() {
            return {
                reg_email: "",
                reg_username: "",
                reg_password1: "",
                reg_password2: "",
            }
        },
    }

    const RecoverUser = {
        template: '#tmpl-recoveruser',
        data: function() {
            return {
                reg_recover_email: "Username",
            }
        },
    }

    const RecoverPass = {
        template: `
        <div id="recoverpass">
            D:
        </div>`,
    }

    const router = new VueRouter({
        routes: [
            { path: '/', component: Login },
            { path: '/register', component: Register },
            { path: '/recover/user', component: RecoverUser },
            { path: '/recover/pass/', component: RecoverPass },
        ]
    })

    var vm = new Vue({
        el: "#portal", 
        router,
     })
}());