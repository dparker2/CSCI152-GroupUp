
function Flashcards(ws) {
    return {
        template: '#tmpl-flashcards',
        created: function() {
            
        },
        data: function() {
            return { 
                cardText: false,
                cardLabel: true,
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
            saveEdit: function(side) {
                this.cardLabel = !this.cardLabel; 
                this.cardText = !this.cardText;
                this.saveIcon = false;
                this.navArrow = true;

                if(side == 'front'){
                    this.sendFront();
                }
                else{
                    this.sendBack();
                }
            },
            showEditView: function() {
                this.cardLabel = !this.cardLabel; 
                this.cardText = !this.cardText;
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
                ws.send(JSON.stringify({
                    code: "group/flashcards/editfront",
                    groupid: this.$parent.groupid,
                    front: this.frontText,
                    index: this.cardIndex
                }));
            },
            sendBack: function () {
                ws.send(JSON.stringify({
                    code: "group/flashcards/editback",
                    groupid: this.$parent.groupid,
                    back: this.backText,
                    index: this.cardIndex
                }));
            }
        }
    }
}
