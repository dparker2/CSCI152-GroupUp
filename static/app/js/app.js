
(function() {
    const socket = new WebSocket("ws://" + document.location.host + "/app/ws")

    const router = new VueRouter({
        routes: [
            { path: '/', component: Home() },
            { path: '/group/create', component: CreateGroup() },
            { path: '/group/join', component: JoinGroup() },
            { path: '/settings', component: Settings() },
            { path: '/group/:groupid/', component: Group(socket) }
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