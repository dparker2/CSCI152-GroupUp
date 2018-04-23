function Whiteboard(ws) {
    var myBoard;
    var test = new Array();
    return {
        template: '#tmpl-whiteboard',
        mounted: function(){
            myBoard = new DrawingBoard.Board("wb");
            alert(document.getElementById('wb'));
            var x = "hi michelle salomon";
            myBoard.ev.bind('board:stopDrawing', this.getWB);
            myBoard.ev.bind('board:reset', this.sendWB);

        },
        data: function() {
            return {

            }
        },
        methods: {
            sendWB: function(){
              myBoard._onCanvasDrop();


               //alert(x);
            },
            getWB: function(){
                //alert("GETWB" + myBoard);
                var storage = myBoard._getStorage();
                myBoard.clearWebStorage();
                test.push(storage);
            }
        }      
    }
};