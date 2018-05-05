
function Home(ws) {
    return {
        template: '#tmpl-home',
        created: function() {
            // Need to send after socket has connected
            if (ws.readyState === 1) {
                ws.send(JSON.stringify({
                    code: "home",
                }));
            } else {
                ws.onopen = function() {
                    ws.send(JSON.stringify({
                        code: "home",
                    }));
                }.bind(this);
            }
        },
        data: function() {
            return {
                currentGroups: this.$parent.currentGroups,
            }
        },
        methods: {
            removeCurrentGroup: function(groupid) {
                ws.send(JSON.stringify({
                    code: "group/remove",
                    groupid: groupid,
                }));
            }
        },
        components: {
            'chat-box': Chatbox(),
        },
    }
}