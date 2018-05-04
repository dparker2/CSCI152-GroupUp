
function Flashcards(ws) {
    return {
        template: '#tmpl-flashcards',
        created: function() {
            ws.addEventListener('message', function(event) {
                data = event.data;
                if (!data)
                    return;
                data = JSON.parse(data);
                code = data.code;
                if (!code || code !== "group/flashcard")
                    return;

                // do something
            }.bind(this));
        },
        data: function() {
            return { 
                cardText: false,
                cardLabel: true,
                saveIcon: false,
                frontText: 'Double click to edit',
                backText: 'Double click to edit',
            }
        },
        methods: {
            flipCard: function() {
                document.getElementById('flipContainer').classList.toggle('flip');
            }, 
            saveEdit: function() {
                this.cardLabel = !this.cardLabel; 
                this.cardText = !this.cardText;
                this.saveIcon = false;
            },
            editCard: function() {
                this.cardLabel = !this.cardLabel; 
                this.cardText = !this.cardText;
                this.saveIcon = true;
            }
        }
    }
}
