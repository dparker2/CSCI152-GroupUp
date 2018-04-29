
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
        },
        data: function() {
            return {

            }
        },
        methods: {
            
            drawWB: function(data){
                var oldDrawing = groupBoard.isDrawing;
                var oldCoords = groupBoard.coords;
                var oldColor = groupBoard.color;
                var oldMode = groupBoard.mode;

                groupBoard.isDrawing = true;
                groupBoard.coords = JSON.parse(data.whiteboardCoords);
                groupBoard.setColor(data.whiteboardColor);
                groupBoard.setMode(data.whiteboardMode, true);
                console.log(groupBoard.coords);
                /*if(groupBoard.getMode() == 'filler'){
                    groupBoard.fill(groupBoard);
                }*/
                groupBoard.draw();

                groupBoard.isDrawing = oldDrawing;
                groupBoard.coords = oldCoords
                groupBoard.setColor(oldColor);
                groupBoard.setMode(oldMode, true);
            },

            sendWB: function() {
                setTimeout(function () {
                    if (groupBoard.isDrawing) {
                        ws.send(JSON.stringify({
                            code: "group/whiteboard",
                            groupid: this.$parent.groupid,
                            whiteboardCoords: JSON.stringify(groupBoard.coords),
                            whiteboardColor: groupBoard.color,
                            whiteboardMode: groupBoard.getMode(),                    
                        }));
                    }
                }.bind(this), 0);
            },
            /*
            stopDrawing: function(){                
                groupBoard.isDrawing = false;
            }
            */
        }      
    }
};