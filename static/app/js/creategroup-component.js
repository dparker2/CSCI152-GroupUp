
function CreateGroup(ws) {
    return {
        template: '#tmpl-creategroup',
        created: function() {
            ws.addEventListener('message', function(event) {
                data = event.data;
                if (!data)
                    return;
                data = JSON.parse(data);
                code = data.code;
                if (!code || code !== "group/create")
                    return;
                groupid = data.groupid;
                this.$router.push(groupid)
            }.bind(this));
        },
        data: function() {
            return {
                createGroupName: '',
            }
        },
        methods: {
            createGroup: function() {
                ws.send(JSON.stringify({
                    code: "group/create",
                    groupid: this.createGroupName,
                }));
            }
        }
    }
}