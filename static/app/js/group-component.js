
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
                code: "group/chat",
                groupid: this.$route.params.groupid,
                chat: this.inputMessage,
            }));
            this.inputMessage = '';
        },
    },
    created: function() {
        var me = this;
        // Need to send after socket has connected
        function send_join() {
            socket.send(JSON.stringify({
                code: "group/join",
                groupid: me.$route.params.groupid,
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