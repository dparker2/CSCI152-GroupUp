const socket = WebSocket;

socket.onmessage = function (event) {
    console.log(event);
}

(function() {
    const Home = {
        template: '#tmpl-home',
        data: function() {
            return {

            }
        },
    }

    const CreateGroup = {
        template: '#tmpl-creategroup',
        data: function() {
            return {

            }
        },
    }

    const JoinGroup = {
        template: '#tmpl-joingroup',
        data: function() {
            return {

            }
        },
    }

    const Settings = {
        template: '#tmpl-settings',
        data: function() {
            return {

            }
        }
    }

    const router = new VueRouter({
        routes: [
            { path: '/', component: Home },
            { path: '/group/create', component: CreateGroup },
            { path: '/group/join', component: JoinGroup },
            { path: '/settings', component: Settings },
        ],
        linkExactActiveClass: "active-button",
    })

    var vm = new Vue({
        el: "#app", 
        router,
        data: {
            showMenu: false,
        },
     })
}());