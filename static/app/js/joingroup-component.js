
function JoinGroup(ws) {
    return {
        template: '#tmpl-joingroup',
        data: function() {
            return {
                joinGroupName: '',
            }
        },
        methods: {
            joinGroup: function() {
                this.$router.push(this.joinGroupName)
            }
        }
    }
}