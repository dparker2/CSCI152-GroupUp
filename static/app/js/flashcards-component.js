
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
                saveIcon: false,
                navArrow: true,
                frontText: 'Double click to edit',
                backText: 'Double click to edit',
                deck: this.$parent.deck,
                view: 'card',
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
                this.saveIcon = true;
                this.navArrow = false;
            },
            nextCard: function () {
                // check if index not out of bounds
                while(this.deck[this.currentCard].index != this.currentCard){
                    this.currentCard =  (1 + this.currentCard) % this.deck.length;
                }
            },
            prevCard: function () {
                // check if index not out of bounds
                while(this.deck[this.currentCard].index != this.currentCard){
                    this.currentCard =  (this.currentCard - 1) % this.deck.length;
                }
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
                    front: this.frontText,
                    index: this.currentCard.toString()
                }));
            },
            sendBack: function () {
                ws.send(JSON.stringify({
                    code: "group/flashcards/editback",
                    groupid: this.$parent.groupid,
                    back: this.backText,
                    index: this.currentCard
                }));
            }
        }
    }
}
