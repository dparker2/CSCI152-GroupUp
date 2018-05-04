
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
    });

    socket.addEventListener('message', function (event) {
        data = event.data;
        if (!data)
            return;
        data = JSON.parse(data);
        code = data.code;
        if (!code || !code.startsWith("app/current"))
            return;
        if (code.endsWith("add")) {
            console.log(data)
            vm.currentGroups.push(data.groupid)
        } else if (code.endsWith("remove")) {
            var index = vm.currentGroups.indexOf(data.groupid);
            if (index > -1) {
                vm.currentGroups.slice(index, 1);
            }
        }
    }.bind(vm));
    
    socket.addEventListener('message', function (event) {
        data = event.data;
        if (!data)
            return;
        data = JSON.parse(data);
        code = data.code;
        if (!code || !code.startsWith("app/previous"))
            return;
            if (code.endsWith("add")) {
                vm.previousGroups.push(data.groupid)
            } else if (code.endsWith("remove")) {
                var index = vm.previousGroups.indexOf(data.groupid);
                if (index > -1) {
                    vm.previousGroups.slice(index, 1);
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
            previousGroups: [],
        },
     })
}());