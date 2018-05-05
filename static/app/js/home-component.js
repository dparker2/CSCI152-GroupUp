
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

            ws.addEventListener('message', function (event) {
                data = event.data;
                if (!data)
                    return;
                data = JSON.parse(data);
                code = data.code;
                if (!code || !code.startsWith("app/friends"))
                    return;
                if (code.endsWith("online")) {
                    var index = this.offlineFriends.indexOf(data.username)
                    if (index !== -1) {
                        this.offlineFriends.splice(index, 1)
                    }
                    this.onlineFriends.push(data.username)
                } else if (code.endsWith("offline")) {
                    var index = this.onlineFriends.indexOf(data.username)
                    if (index !== -1) {
                        this.onlineFriends.splice(index, 1)
                    }
                    this.offlineFriends.push(data.username)
                }
            }.bind(this));

            ws.addEventListener('message', function (event) {
                data = event.data;
                if (!data)
                    return;
                data = JSON.parse(data);
                code = data.code;
                if (!code || !code.startsWith("app/search"))
                    return;
                if (code.endsWith("users") && data.query && data.query === this.searchQuery) {
                    this.friendsResults.push(data.username)
                }
            }.bind(this));
        },
        data: function() {
            return {
                currentGroups: this.$parent.currentGroups,
                onlineFriends: [],
                offlineFriends: [],
                friendsResults: [],
                searchQuery: "",
            }
        },
        methods: {
            removeCurrentGroup: function(groupid) {
                ws.send(JSON.stringify({
                    code: "group/remove",
                    groupid: groupid,
                }));
            },
            searchUsers: function() {
                this.friendsResults = []
                if (this.searchQuery.length > 2) {
                    ws.send(JSON.stringify({
                        code: "app/search/users",
                        query: this.searchQuery,
                    }))
                }
            }
        },
        components: {
        },
    }
}