(function() {
    const socket = new WebSocket("ws://" + document.location.host + "/app/ws");

    socket.onmessage = function (event) {
        console.log(event.data);
        wsData = JSON.parse(event.data)
        if (wsData.code == "group/chat") {
            $('.chat-box').append("<p>" + wsData.username + ": " + wsData.chat + "</p>");
        }
    }

    const router = new VueRouter({
        routes: [
            { path: '/', component: Home },
            { path: '/group/create', component: CreateGroup },
            { path: '/group/join', component: JoinGroup },
            { path: '/settings', component: Settings },
            { path: '/group/:groupid/', component: Group }
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


whiteboardCtx = document.getElementById('whiteboard').getContext('2d');