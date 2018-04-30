
function Whiteboard(ws) {
    return {
        template: '#tmpl-whiteboard',
        mounted: function(){
            console.log("first");
            groupBoard = new DrawingBoard.Board("wb");
            groupBoard.ev.bind('board:drawing', this.sendWB);            
            
            ws.addEventListener('message', function(event){
                data = event.data;
                if(!data)
                    return;
                data = JSON.parse(data);
                code = data.code;
                if(!code || code!=="group/whiteboard")
                    return;   
                this.drawWB(data); 
            }.bind(this));

            this.worker = new Worker('/js/app/whiteboard-worker.js');

            this.worker.onmessage = function(message) {
                var data = message.data;
                if (data[0] === "sendWB") {
                    ws.send(data[1]);
                }
            }
        },
        data: function() {
            return {

            }
        },
        methods: {
            
            drawWB: function(data){
                var oldState = {
                    color: groupBoard.color,
                    mode: groupBoard.getMode(),
                }

                //groupBoard.setColor(data.whiteboardColor);
                groupBoard.setMode(data.whiteboardMode, true);
                if(groupBoard.getMode() == 'filler'){
                    groupBoard.fill(groupBoard);
                }

                data.whiteboardCoords = JSON.parse(data.whiteboardCoords);

                var currentMid = {
                    x: data.whiteboardCoords.old.x + data.whiteboardCoords.current.x>>1,
                    y: data.whiteboardCoords.old.y + data.whiteboardCoords.current.y>>1,
                }
                groupBoard.ctx.beginPath();
                groupBoard.ctx.moveTo(currentMid.x, currentMid.y);
                groupBoard.ctx.quadraticCurveTo(data.whiteboardCoords.old.x, 
                    data.whiteboardCoords.old.y, 
                    data.whiteboardCoords.oldMid.x, 
                    data.whiteboardCoords.oldMid.y);
                groupBoard.ctx.stroke();

                groupBoard.setColor(oldState.color);
                groupBoard.setMode(oldState.mode, true);
            },

            sendWB: function() {
                if (groupBoard.isDrawing) {
                    this.worker.postMessage(["sendWB", groupBoard.coords, groupBoard.color, groupBoard.getMode(), this.$parent.groupid]);
                }
            },
            /*
            stopDrawing: function(){                
                groupBoard.isDrawing = false;
            }
            */
        }      
    }
};