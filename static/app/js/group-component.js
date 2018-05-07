
function Group(ws) {
    return {
        template: '#tmpl-group',
        mounted: function() {

            ws.addEventListener('message', function (event) {
                    data = event.data;
                    if (!data)
                        return;
                    data = JSON.parse(data);
                    code = data.code;
                    if (!code || code !== "group")
                        return;
                    if (data.groupid === "") {
                        this.showError = true
                    } else {
                        this.showGroup = true
                    }
                }.bind(this));

                ws.addEventListener('message', function(event) {
                    data = event.data;
                    if (!data)
                        return;
                    data = JSON.parse(data);
                    code = data.code;
                    if (!code || !code.startsWith("group/flashcards"))
                        return;    
                    
                    if(code.endsWith("new")){       
                        this.addCard(data.index, data.front, data.back)
                    }
                    if(code.endsWith("editfront")){
                        this.editFront(data.index, data.front)
                    }
                    if(code.endsWith("editback")){
                        this.editBack(data.index, data.back)
                    }
                
    
                    // do something important group/flashcard/new
                }.bind(this));
            
            // Need to send after socket has connected
            if (ws.readyState === 1) {
                ws.send(JSON.stringify({
                    code: "group/join",
                    groupid: this.groupid,
                }));
            } else {
                ws.onopen = function() {
                    ws.send(JSON.stringify({
                        code: "group/join",
                        groupid: this.groupid,
                    }));
                }.bind(this);
            }
        },
        beforeDestroy: function() {
            ws.send(JSON.stringify({
                code: "group/leave",
                groupid: this.groupid,
            }));
        },
        data: function() {
            return {
                drawing: false,
                groupid: this.$route.params.groupid,
                showGroup: false,
                showError: false,
                studyMode: 'whiteboard',
                deck: []
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
            addCard: function(index) {
                this.deck.push({
                    front: "",
                    back: "",
                    index: index
                })
            },
            editFront: function(index, front) {
                var card_location = this.findCard(index);
                if(card_location !== -1){
                    this.deck[card_location].front = front;
                }
            },
            editBack: function(index, back) {
                var card_location = this.findCard(index);
                if(card_location !== -1){
                    this.deck[card_location].back = back;
                }
            },
            findCard: function(index){
                
                for(var card_location = 0; card_location < this.deck.length; card_location++){
                    if(this.deck[card_location].index == index){
                        return card_location;
                    }
                }
                return -1;
            }, 
        },
        components: {
            'chat-box': Chatbox(ws),
            'white-board': Whiteboard(ws),
            'flash-cards': Flashcards(ws),
            'user-list': Userlist(ws),
        },
    }
}