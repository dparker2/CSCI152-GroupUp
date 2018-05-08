
function Flashcards(ws) {
    return {
        template: '#tmpl-flashcards',
        created: function() {
            
        },
        data: function() {
            return { 
                isShowingCardText: true,
                isEditingCardText: false,
                isEditingFrontListView: false,
                isEditingBackListView: false,
                cardEditingInList: 0,
                currentCard: 1,
                view: 'list',
                textareaFront: '',
                textareaBack: '',
                textareaListView: '',
                deckSize: 0,
            }
        },
        methods: {   
            flipCard: function() {
                document.getElementById('flipContainer').classList.toggle('flip');
            },
            toggleCardView: function() {
                this.isShowingCardText = !this.isShowingCardText; 
                this.isEditingCardText = !this.isEditingCardText;
                this.showExitEditButton = !this.showExitEditButton;
            },
            editListView: function(index, text, isFront) {
                this.textareaEdit = text;
                this.cardEditingInList = index;
                if (isFront) {
                    this.isEditingFrontListView = true;
                } else {
                    this.isEditingBackListView = true;
                }
            },
            normalListView: function(index, isFront, sendfunc) {
                sendfunc(index);
                this.cardEditingInList = 0;
                this.currentCard = index;
                if (isFront) {
                    this.isEditingFrontListView = false;
                } else {
                    this.isEditingBackListView = false;
                }
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
            sendFront: function (index) {
                text = { text: this.textareaEdit };
                card = (index) ? { index: index } : { index: this.currentCard };
                this.textareaEdit = '';
                ws.send(JSON.stringify({
                    code: "group/flashcards/editfront",
                    groupid: this.$parent.groupid,
                    front: text.text,
                    index: card.index.toString(),
                }));
            },
            sendBack: function (index) {
                text = { text: this.textareaEdit };
                card = (index) ? { index: index } : { index: this.currentCard };
                this.textareaEdit = '';
                ws.send(JSON.stringify({
                    code: "group/flashcards/editback",
                    groupid: this.$parent.groupid,
                    back: text.text,
                    index: card.index.toString(),
                }));
            }
        },
        watch: {
            cardEditingInList: function () {
                this.$nextTick(function() {
                    jQuery('#activeTextarea').focus();
                })
            }
        }
    }
}
