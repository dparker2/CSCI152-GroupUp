const socket = new WebSocket('ws://localhost:3000');

(function() {
    const Login = {
        template: '#tmpl-login',
        data: function() {
            return {
                username: "",
                password: "",
            }
        },
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
        template: `
        <div id="recoveruser">
            :D
        </div>`,
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