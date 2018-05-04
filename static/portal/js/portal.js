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
        methods: {
            register: function() {
                if( this.validateRegister() ) {
                    this.$http.post('/register', { // Data
                        reg_email: this.reg_email,
                        reg_username: this.reg_username,
                        reg_password1: this.reg_password1,
                        reg_password2: this.reg_password2,
                        security_question1: this.security_question1,
                        security_answer1: this.security_answer1,
                        security_question2: this.security_question2,
                        security_answer2: this.security_answer2,
                        security_question3: this.security_question3,
                        security_answer3: this.security_answer3
                    }, { // Config
                        emulateJSON: true
                    }).then(function(response) { // Success
                        console.log(response);
                        if (response.data) {
                            console.log(response.data);
                            alert("Account registered!")
                            window.location.href = response.data["redirect-path"];
                        } else {
                            console.log("Something went wrong")
                        }
                    }, function(response) { // Error
                        console.log(response);
                    });
                } else {
                    alert("Please fill in the entire form.")
                }
            },
            validateRegister: function() {
                return this.validatePass() && this.validateUsername() && this.validateEmail()
            },
            validatePass: function() {
                if(this.reg_password1 === ""){
                    return false
                }
                return (this.reg_password1 === this.reg_password2)
            },
            validateUsername: function() {
                if(this.reg_username === ""){
                    return false
                }
                return true
            },
            validateEmail: function() {
                if(this.reg_email === ""){
                    return false
                }
                return true
            }
        }
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
            { path: '/recover/pass', component: RecoverPass,
                children: [{ path: '/check/', component: check}]
            },
            
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