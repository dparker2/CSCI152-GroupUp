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
                security_question1: "",
                security_answer1: "",
                security_question2: "",
                security_answer2: "",
                security_question3: "",
                security_answer3: "",
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

    const About = {
        template: '#templ-about'
    }

    const Help = {
        template: '#templ-help'
    }

    const FAQ = {
        template: ''
    }

    const Contacts = {
        template: ''
    }

    const Privacy = {
        template: ''
    }

    const News = {
        template: ''
    }

    const router = new VueRouter({
        routes: [
            { path: '/', component: Login },
            { path: '/register', component: Register },
            { path: '/recover/user', component: RecoverUser },
            { path: '/recover/pass/', component: RecoverPass },
            { path: '/about', component: About },
            { path: '/help', component: Help },
            { path: '/FAQ', component: FAQ },
            { path: '/contacts', componenet: Contacts },
            { path: '/privacy', component: Privacy },
            { path: '/news', component: News },
        ]
    })

    var vm = new Vue({
        el: "#portal", 
        router,
     })
}());