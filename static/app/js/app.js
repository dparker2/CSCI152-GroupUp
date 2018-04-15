(function() {
    const socket = new WebSocket("ws://" + document.location.host + "/app/ws");

    socket.onmessage = function (event) {
        console.log(event.data);
        wsData = JSON.parse(event.data)
        if (wsData.code == "CHAT") {
            $('.chat-box').append("<p>" + wsData.username + ": " + wsData.chat + "</p>");
        }
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
                createGroupName: '',
            }
        },
        methods: {
            createGroup: function() {
                alert(this.createGroupName);
            }
        }
    }

    const JoinGroup = {
        template: '#tmpl-joingroup',
        data: function() {
            return {
                joinGroupName: '',
            }
        },
        methods: {
            joinGroup: function() {
                alert(this.joinGroupName);
            }
        }
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
                drawing: false,
                inputMessage: '',
            }
        },
        methods: {
            startDraw: function(e) {
                if (!this.drawing) {
                    this.drawing = true;
                    var x = e.x - e.target.offsetLeft;
                    var y = e.y - e.target.offsetTop;
                    whiteboardCtx.moveTo(x, y);
                    whiteboardCtx.beginPath();
                }
            },
            draw: function(e) {
                if (this.drawing) {
                    console.log(e);
                    var x = e.x - e.target.offsetLeft;
                    var y = e.y - e.target.offsetTop;
                    whiteboardCtx.lineTo(x, y);
                    whiteboardCtx.stroke();
                }
            },
            endDraw: function() {
                if (this.drawing) {
                    this.drawing = false;
                    whiteboardCtx.closePath();
                }
            },
            sendChat: function() {
                socket.send(JSON.stringify({
                    code: "CHAT",
                    groupid: this.$route.params.groupid,
                    chat: this.inputMessage,
                    username: this.$route.params.username,
                }));
                this.inputMessage = '';
            },
        },
        created: function() {
            var me = this;
            // Need to send after socket has connected
            function send_join() {
                socket.send(JSON.stringify({
                    code: "JOIN GROUP",
                    groupid: me.$route.params.groupid,
                    username: me.$route.params.username,
                }));
            }
            // Make send_join() send after socket is opened, or now if it already is
            if (socket.readyState !== 1) {
                socket.onopen = send_join;
            } else {
                send_join();
            }
        }
    }

    const router = new VueRouter({
        routes: [
            { path: '/', component: Home },
            { path: '/group/create', component: CreateGroup },
            { path: '/group/join', component: JoinGroup },
            { path: '/settings', component: Settings },
            { path: '/group/:groupid/:username', component: Group }
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