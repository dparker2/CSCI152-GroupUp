
function Flashcards(ws) {
    return {
        template: '#tmpl-flashcards',
        created: function() {
            
        },
        data: function() {
            return { 
                isShowingCardText: true,
                isEditingCardText: false,
                currentCard: 1,
                view: 'card',
                textareaFront: '',
                textareaBack: '',
                deckSize: 0,
            }
        },
        methods: {   
            flipCard: function() {
                document.getElementById('flipContainer').classList.toggle('flip');
            },
            showEditView: function() {
                this.isShowingCardText = !this.isShowingCardText; 
                this.isEditingCardText = !this.isEditingCardText;
            },
            nextCard: function () {
                this.currentCard = this.currentCard % (this.$parent.deck.length) + 1;
                document.getElementById('flipContainer').classList.remove('flip');
            },
            prevCard: function () {
                this.currentCard = (this.currentCard === 1) ? this.$parent.deck.length : this.currentCard - 1;
                document.getElementById('flipContainer').classList.remove('flip');
            },
            sendNewCard: function() {
                ws.send(JSON.stringify({
                    code: "group/flashcards/new",
                    groupid: this.$parent.groupid,
                }));
            },
            sendFront: function () {
                this.isEditingCardText = !this.isEditingCardText;
                this.isShowingCardText = !this.isShowingCardText;
                ws.send(JSON.stringify({
                    code: "group/flashcards/editfront",
                    groupid: this.$parent.groupid,
                    front: this.textareaFront,
                    index: this.currentCard.toString(),
                }));
                this.textareaFront = '';
            },
            sendBack: function () {
                this.isEditingCardText = !this.isEditingCardText;
                this.isShowingCardText = !this.isShowingCardText;
                ws.send(JSON.stringify({
                    code: "group/flashcards/editback",
                    groupid: this.$parent.groupid,
                    back: this.textareaBack,
                    index: this.currentCard.toString(),
                }));
                this.textareaBack = '';
            }
        }
    }
}
