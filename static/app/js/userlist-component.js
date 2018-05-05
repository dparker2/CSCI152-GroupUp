
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
                if (data.username && data.status === "0") {
                    var index = this.inactiveUsers.indexOf(data.username);
                    if (index > -1) { // In inactive users
                        this.inactiveUsers.splice(index, 1);
                        this.activeUsers.push(data.username)
                    } else {
                        this.inactiveUsers.push(data.username)
                    }
                } else {
                    this.activeUsers.push(data.username)
                    var index = this.inactiveUsers.indexOf(data.username);
                    if (index > -1) { // In inactive users
                        this.inactiveUsers.splice(index, 1);
                    }
                }
            }.bind(this));

            ws.addEventListener('message', function (event) {
                data = event.data;
                if (!data)
                    return;
                data = JSON.parse(data);
                code = data.code;
                if (!code || code !== "group/leave")
                    return;
                if (data.username) {
                    var index = this.activeUsers.indexOf(data.username);
                    if (index > -1) { // In active users
                      this.activeUsers.splice(index, 1);
                      this.inactiveUsers.push(data.username)
                    }
                }
            }.bind(this));

            ws.addEventListener('message', function (event) {
                data = event.data;
                if (!data)
                    return;
                data = JSON.parse(data);
                code = data.code;
                if (!code || code !== "group/remove")
                    return;
                if (data.username) {
                    var activeIndex = this.activeUsers.indexOf(data.username);
                    var inactiveIndex = this.inactiveUsers.indexOf(data.username);
                    if (activeIndex > -1) { // In active users
                        this.activeUsers.splice(activeIndex, 1);
                    }
                    if (inactiveIndex > -1) {
                        this.inactiveUsers.splice(inactiveIndex, 1);
                    }
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
