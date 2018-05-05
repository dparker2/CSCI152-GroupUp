
function Home(ws) {
    return {
        template: '#tmpl-home',
        created: function() {
            // Need to send after socket has connected
            console.log(ws.readyState)
            if (ws.readyState === 1) {
                ws.send(JSON.stringify({
                    code: "home",
                }));
            }

            ws.addEventListener('message', function (event) {
                data = event.data;
                if (!data)
                    return;
                data = JSON.parse(data);
                code = data.code;
                if (!code || !code.startsWith("app/friends"))
                    return;
                var offlineIndex = this.offlineFriends.indexOf(data.username)
                var onlineIndex = this.onlineFriends.indexOf(data.username)
                if (code.endsWith("online")) {
                    if (offlineIndex !== -1) {
                        this.offlineFriends.splice(offlineIndex, 1)
                    }
                    if (onlineIndex === -1) {
                        this.onlineFriends.push(data.username)
                    }
                } else if (code.endsWith("offline")) {
                    if (onlineIndex !== -1) {
                        this.onlineFriends.splice(onlineIndex, 1)
                    }
                    if (offlineIndex === -1) {
                        this.offlineFriends.push(data.username)
                    }
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
                if (code.endsWith("groups") && data.query && data.query === this.groupSearchQuery) {
                    this.groupsResults.push({
                        groupid: data.groupid,
                        users: data.status,
                        creator: data.username,
                    })
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
                groupSearchQuery: "",
                groupsResults: [],
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
            },
            searchGroups: function() {
                this.groupsResults = []
                if (this.groupSearchQuery.length > 2) {
                    ws.send(JSON.stringify({
                        code: "app/search/groups",
                        query: this.groupSearchQuery,
                    }))
                }
            },
            addFriend: function(username) {
                ws.send(JSON.stringify({
                    code: "app/friends/add",
                    username: username,
                }))
            }
        },
        components: {
        },
    }
}