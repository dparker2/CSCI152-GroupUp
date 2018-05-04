onmessage = function(message) {
    var data = message.data;
    if (data[0] === "sendWB") {
        var coords = JSON.stringify(data[1]);
        self.postMessage([data[0], JSON.stringify({
            code: "group/whiteboard",
            groupid: data[4],
            whiteboardCoords: coords,
            whiteboardColor: data[2],
            whiteboardMode: data[3],  
        })]);
    }
}