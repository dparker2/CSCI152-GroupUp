
function Group(ws) {
    return {
        template: '#tmpl-group',
        data: function() {
            return {
                drawing: false,
                groupid: this.$route.params.groupid,
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
        },
        components: {
            'chat-box': Chatbox(ws),
        },
        created: function() {
            var me = this;
            // Need to send after socket has connected
            ws.send(JSON.stringify({
                code: "group/join",
                groupid: me.groupid,
            }));
        }
    }
}