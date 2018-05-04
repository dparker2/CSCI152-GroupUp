
function Userlist(ws) {
    return {
        template: '#tmpl-userlist',
        created: function() {

            ws.addEventListener('message', function (event) {
                data = event.data;
                if (!data)
                    return;
                data = JSON.parse(data);
                code = data.code;
                if (!code || code !== "group/join")
                    return;
                if (data.status == 0) {
                    this.inactiveUsers.push(data.username)
                } else {
                    this.activeUsers.push(data.username)
                }
            }.bind(this));
        },
        data: function() {
            return {
                activeUsers: [],
                inactiveUsers: [],
                parentName: this.$parent.groupid,
            }
        },
        methods: {
        },
    }
};
