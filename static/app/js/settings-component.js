
const Settings = {
    template: '#tmpl-settings',
    data: function() {
        return {
            message: '',
        }
    },
    methods: {
        sendmsg: function() {
            socket.send(this.message);
        }
    },
}