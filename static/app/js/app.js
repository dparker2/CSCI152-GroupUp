const socket = new WebSocket("ws://localhost:3000/app/ws");

(function() {
    socket.onmessage = function (event) {
        console.log(event.data);
    }

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
                message: '',
            }
        },
        methods: {
            sendmsg: function() {
                socket.send(this.message);
            }
        },
    }

    const Group = {
        template: '#tmpl-group',
        data: function() {
            return {

            }
        },
        methods: {

        },
    }

    const router = new VueRouter({
        routes: [
            { path: '/', component: Home },
            { path: '/group/create', component: CreateGroup },
            { path: '/group/join', component: JoinGroup },
            { path: '/settings', component: Settings },
            { path: '/group/', component: Group }
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