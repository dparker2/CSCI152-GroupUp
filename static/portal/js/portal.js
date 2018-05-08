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
                errorList: [],
            }
        },
        methods: {
            register: function() {
                this.errorList = []
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
                            window.location.href = response.data["redirect-path"];
                        } else {
                            console.log("Something went wrong")
                        }
                    }, function(response) { // Error
                        console.log(response);
                    });
                } 
            },
            validateRegister: function() {
                var validPass = this.validatePass();
                var validUser = this.validateUsername();
                var validEmail = this.validateEmail();
                return validPass && validUser && validEmail
            },
            validatePass: function() {
                if(this.reg_password1 === "" || this.hasWhitespace(this.reg_password1) || (this.reg_password1 !== this.reg_password2)) {
                    console.log(this.errorList)
                    this.errorList.push("Invalid password. Your password must exist and not contain any spaces.")
                    console.log(this.errorList)
                    return false
                }
                return true
            },
            validateUsername: function() {
                if(this.reg_username === "" || this.hasWhitespace(this.reg_username)) {
                    this.errorList.push("Invalid username. Your username must exist and not contain any spaces.")
                    return false
                }
                return true
            },
            validateEmail: function() {
                if(this.reg_email === "" || this.hasWhitespace(this.reg_email)){
                    this.errorList.push("Invalid email. Your email must exist and not contain any spaces.")
                    return false
                }
                return true
            },
            hasWhitespace: function(myString) {
                return myString.indexOf(" ") >= 0
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