
function Home() {
    return {
        template: '#tmpl-home',
        data: function() {
            return {
                
            }
        },
        components: {
            'chat-box': Chatbox(),
        }, 
    }
}