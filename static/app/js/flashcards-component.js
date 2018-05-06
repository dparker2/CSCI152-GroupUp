
function Flashcards(ws) {
    return {
        template: '#tmpl-flashcards',
        created: function() {
            this.setDeck();   /** TEST **/
            
            ws.addEventListener('message', function(event) {
                data = event.data;
                if (!data)
                    return;
                data = JSON.parse(data);
                code = data.code;
                if (!code || code !== "group/flashcard")
                    return;
                
                if(data.front){
                    
                }
                if(data.front){

                }
                this.addCard(data.index, data.front, data.back)

                // do something important group/flashcard/new
            }.bind(this));
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
                deck: [],
                view: 'card',
                deckSize: 0,
            }
        },
        methods: {
            addCard: function(index, front, back) {
                

            },
            updateCard: function() {
                
            },
            setDeck: function() {
        
                this.deck.push({
                    front: "michelle",
                    back: "salomon",
                    index: 4
                });

                this.deck.push({
                    front: "group",
                    back: "up",
                    index: 2
                });

                this.deck.push({
                    front: "hellow",
                    back: "world",
                    index: 1
                });
                this.deck.push({
                    front: "number",
                    back: "three",
                    index: 3
                });
                this.deckSize = this.deck.length;
                
            }, 
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
            editCard: function() {
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
                    index: -1
                }));
            },
            sendFront: function () {
                ws.send(JSON.stringify({
                    code: "group/flashcards/editFront",
                    groupid: this.$parent.groupid,
                    front: this.frontText,
                    index: this.cardIndex
                }));
            },
            sendBack: function () {
                ws.send(JSON.stringify({
                    code: "group/flashcards/editBack",
                    groupid: this.$parent.groupid,
                    back: this.backText,
                    index: this.cardIndex
                }));
            }
        }
    }
}
