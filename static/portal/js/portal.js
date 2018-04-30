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
                    if (response.data) {
                        console.log(response.data);
                        window.location.href = response.data["redirect-path"];
                    } else {
                        console.log("Incorrect Password")
                    }
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
        template: '#tmpl-recoverpass',
        data: function() {
            return {
                reg_recover_email: "",
            }
        },
    }

    const check = {
        template: '#tmpl-check',
        data: function() {
            return {
                security_question1: "",
                security_answer1: "",
                security_question2: "",
                security_answer2: "",
                security_question3: "",
                security_answer3: "",
            }
        },
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
            { path: '/recover/pass', component: RecoverPass },
            { path: '/recover/pass/check/', component: check },
            { path: '/about', component: About },
            { path: '/help', component: Help },
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