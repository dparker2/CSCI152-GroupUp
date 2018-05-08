
function Chatbox(ws) {
    return {
        template: '#tmpl-chatbox',
        created: function() {

            ws.addEventListener('message', function(event) {
                data = event.data;
                if (!data)
                    return;
                data = JSON.parse(data);
                code = data.code;
                if (!code || code !== "group/chat")
                    return;
                groupid = data.groupid;
                if (!groupid || groupid !== this.parentName)
                    return;
                
                chat = data.chat;
                username = data.username;
                d = new Date(data.timestamp);
                tstamp = d.toLocaleString()

                if (chat && username) {
                    this.messages.push({
                        user: username,
                        msg: chat,
                        timestamp: tstamp
                    });
                }
            }.bind(this));
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
