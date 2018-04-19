
function Chatbox(ws) {
    return {
        template: '#tmpl-chatbox',
        created: function() {
            function readChat(parentName, msgs) {
                return function(event) { // Closure to capture parentName
                    data = event.data;
                    if (!data)
                        return;
                    data = JSON.parse(data);
                    code = data.code;
                    if (!code || code !== "group/chat")
                        return;
                    groupid = data.groupid;
                    if (!groupid || groupid !== parentName)
                        return;
                    chat = data.chat;
                    username = data.username;
                    if (chat && username) {
                        msgs.push({
                            user: username,
                            msg: chat,
                        });
                    }
                };
            }

            ws.onmessage = readChat(this.parentName, this.messages);
        },
        data: function() {
            return {
                messages: [],
                inputMessage: '',
                parentName: this.$parent.groupid,
            }
        },
        methods: {
            sendChat: function() {
                ws.send(JSON.stringify({
                    code: "group/chat",
                    groupid: this.parentName,
                    chat: this.inputMessage,
                }));
                this.inputMessage = '';
            },
        },
    }
};
