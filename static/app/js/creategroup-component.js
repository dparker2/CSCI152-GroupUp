
function CreateGroup() {
    return {
        template: '#tmpl-creategroup',
        data: function() {
            return {
                createGroupName: '',
            }
        },
        methods: {
            createGroup: function() {
                alert(this.createGroupName);
            }
        }
    }
}