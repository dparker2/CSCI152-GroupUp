
function Home(ws) {
    var throttle = function(func, limit) {
        var lastFunc
        var lastRan
        return function() {
          const context = this
          const args = arguments
          if (!lastRan) {
            func.apply(context, args)
            lastRan = Date.now()
          } else {
            clearTimeout(lastFunc)
            lastFunc = setTimeout(function() {
              if ((Date.now() - lastRan) >= limit) {
                func.apply(context, args)
                lastRan = Date.now()
              }
            }, limit - (Date.now() - lastRan))
          }
        }
      }
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
                previousSearchQuery: "",
                groupSearchQuery: "",
                groupsResults: [],
                previousGroupSearchQuery: "",
            }
        },
        methods: {
            removeCurrentGroup: function(groupid) {
                ws.send(JSON.stringify({
                    code: "group/remove",
                    groupid: groupid,
                }));
            },
            searchUsers: throttle(function() {
                if (this.searchQuery.length > 2 && this.previousSearchQuery !== this.searchQuery) {
                    this.previousSearchQuery = this.searchQuery
                    ws.send(JSON.stringify({
                        code: "app/search/users",
                        query: this.searchQuery,
                    }))
                    this.friendsResults = []
                } else if (this.searchQuery.length <= 2) {
                    this.previousSearchQuery = this.searchQuery
                    this.friendsResults = []
                }
            }, 250),
            searchGroups: throttle(function() {
                if (this.groupSearchQuery.length > 2 && this.previousGroupSearchQuery !== this.groupSearchQuery) {
                    this.previousGroupSearchQuery = this.groupSearchQuery
                    ws.send(JSON.stringify({
                        code: "app/search/groups",
                        query: this.groupSearchQuery,
                    }))
                    this.groupsResults = []
                } else if (this.groupSearchQuery.length <= 2) {
                    this.previousGroupSearchQuery = this.groupSearchQuery
                    this.groupsResults = []
                }
            }, 250),
            addFriend: function(username) {
                ws.send(JSON.stringify({
                    code: "app/friends/add",
                    username: username,
                }))
            },

        },
        components: {
        },
    }
}