
function Whiteboard(ws) {
    return {
        template: '#tmpl-whiteboard',
        mounted: function(){
            console.log("first");
            groupBoard = new DrawingBoard.Board("wb");
            groupBoard.ev.bind('board:stopDrawing', this.sendWB);            
            
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
                groupBoard.coords = JSON.parse(data.whiteboardCoords);
                groupBoard.setColor(data.whiteboardColor);
                groupBoard.setMode(data.whiteboardMode, true);
                if(groupBoard.getMode() == 'filler'){
                    groupBoard.fill(groupBoard);
                }
                groupBoard.isDrawing = true;
                groupBoard.draw();
            },

            sendWB: function() {
                ws.send(JSON.stringify({
                    code: "group/whiteboard",
                    groupid: this.$parent.groupid,
                    whiteboardCoords: JSON.stringify(groupBoard.coords),
                    whiteboardColor: groupBoard.color,
                    whiteboardMode: groupBoard.getMode(),                    
                }));
            },
            /*
            stopDrawing: function(){                
                groupBoard.isDrawing = false;
            }
            */
        }      
    }
};