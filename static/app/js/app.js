
(function() {
    var socket;

    function start(websocketServerLocation){
        socket = new WebSocket(websocketServerLocation);
        socket.onclose = function(){
            // Try to reconnect in 5 seconds
            setTimeout(function(){start(websocketServerLocation)}, 5000);
        };
        socket.onopen = function() {
            socket.send(JSON.stringify({
                code: "home",
            }));
        };
    }
    start("ws://" + document.location.host + "/app/ws");

    socket.addEventListener('message', function(event) {
        console.log(event);
    });

    socket.addEventListener('message', function (event) {
        data = event.data;
        if (!data)
            return;
        data = JSON.parse(data);
        code = data.code;
        if (!code || !code.startsWith("app/current"))
            return;
        var index = vm.currentGroups.indexOf(data.groupid);
        if (code.endsWith("add")) {
            if (index === -1) {
                vm.currentGroups.push(data.groupid)
            }
        } else if (code.endsWith("remove")) {
            if (index !== -1) {
                vm.currentGroups.splice(index, 1);
            }
        }
    }.bind(vm));

    const router = new VueRouter({
        routes: [
            { path: '/', component: Home(socket) },
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
            currentGroups: [],
        },
     })
}());