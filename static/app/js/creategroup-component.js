
const CreateGroup = {
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