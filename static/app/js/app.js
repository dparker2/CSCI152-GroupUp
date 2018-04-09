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
                drawing: false
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

    whiteboardCtx = document.getElementById('whiteboard').getContext('2d');
}());