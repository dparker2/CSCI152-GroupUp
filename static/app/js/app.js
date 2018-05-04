
(function() {

    function start(websocketServerLocation){
        socket = new WebSocket(websocketServerLocation);
        socket.onclose = function(){
            // Try to reconnect in 5 seconds
            setTimeout(function(){start(websocketServerLocation)}, 5000);
        };
    }
    start("ws://" + document.location.host + "/app/ws");

    socket.addEventListener('message', function(event) {
        console.log(event);
    })

    const router = new VueRouter({
        routes: [
            { path: '/', component: Home() },
            { path: '/group/create', component: CreateGroup(socket) },
            { path: '/group/join', component: JoinGroup(socket) },
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