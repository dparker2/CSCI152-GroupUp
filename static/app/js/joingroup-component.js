
const JoinGroup = {
    template: '#tmpl-joingroup',
    data: function() {
        return {
            joinGroupName: '',
        }
    },
    methods: {
        joinGroup: function() {
            alert(this.joinGroupName);
        }
    }
}