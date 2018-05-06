
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
    
                this.deck.push({
                    front: data.front,
                    back: data.back,
                    index: data.index
                });
                

                // do something important group/flashcard/new
            }.bind(this));
        },
        data: function() {
            return { 
                cardText: false,
                cardLabel: true,
                index: 0,
                saveIcon: false,
                navArrow: true,
                frontText: 'Double click to edit',
                backText: 'Double click to edit',
                deck: [],
                view: 'card',
                deckSize: 0
            }
        },
        methods: {
            setDeck: function() {
                this.deck.push({
                    front: "michelle",
                    back: "salomon",
                    index: 1
                });

                this.deck.push({
                    front: "group",
                    back: "up",
                    index: 2
                });

                this.deck.push({
                    front: "hellow",
                    back: "world",
                    index: 3
                });
                this.deckSize = this.deck.length;
                
            }, 
            flipCard: function() {
                document.getElementById('flipContainer').classList.toggle('flip');
            }, 
            saveEdit: function() {
                this.cardLabel = !this.cardLabel; 
                this.cardText = !this.cardText;
                this.saveIcon = false;
                this.navArrow = true;
            },
            editCard: function() {
                this.cardLabel = !this.cardLabel; 
                this.cardText = !this.cardText;
                this.saveIcon = true;
                this.navArrow = false;
            },
            nextCard: function () {
                // check if index not out of bounds
                if(this.index < this.deck.length){
                    this.frontText = this.deck[this.index].front;
                    this.backText = this.deck[this.index].back;
                    this.index = ++this.index;
                }
            },
            prevCard: function () {
                // check if index not out of bounds
                if(this.index > 0){
                    this.index = --this.index;                
                    this.frontText =this.deck[this.index].front;
                    this.backText = this.deck[this.index].back;
                }
            },
            sendCard: function() {
                ws.send(JSON.stringify({
                    code: "group/flashcards/new",
                    groupid: this.$parent.groupid,
                }));

                ws.send(JSON.stringify({
                    code: "group/flashcards/editFront",
                    groupid: this.$parent.groupid,
                    front: this.frontText,
                    index: this.cardIndex
                }));

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
