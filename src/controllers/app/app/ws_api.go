package app

type wsMessage struct {
	Code      string `json:"code"`
	Groupid   string `json:"groupid"`
	Chat      string `json:"chat"`
	Timestamp string `json:"timestamp"`
	Username  string `json:"username"`
	Coords    string `json:"whiteboardCoords"`
	Color     string `json:"whiteboardColor"`
	Mode      string `json:"whiteboardMode"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	Query     string `json:"query"`
	Index     string `json:"index"`
	Front     string `json:"front"`
	Back      string `json:"back"`
}

type wsAPIstruct struct {
	UserToken string
	Msg       *wsMessage
}

var wsAPI map[string]func(wsAPIstruct) error

func setupAPI() {
	// To add another API code:
	//   Add a line here corresponding to the name of the code (eg settings/update friends/add etc)
	//   Follow the pattern and set it equal to the name of the function (that exists within
	//    this package) that is used to handle that code
	//   The wsAPIstruct is passed to every API function, this can be expanded (mostly through wsMessage)
	//    to support more data parameters sent from the client, add those if needed. They'll be automatically
	//    decoded and added to the struct when sent.
	wsAPI = make(map[string]func(wsAPIstruct) error)
	wsAPI["home"] = home
	wsAPI["group/create"] = groupCreate
	wsAPI["group/remove"] = groupRemove
	wsAPI["group/join"] = groupJoin
	wsAPI["group/leave"] = groupLeave
	wsAPI["group/chat"] = groupChat
	wsAPI["group/whiteboard"] = groupWhiteboard
	wsAPI["app/search/users"] = searchUsers
	wsAPI["app/search/groups"] = searchGroups
	wsAPI["app/friends/add"] = friendsAdd
	wsAPI["app/friends/remove"] = friendsRemove
	wsAPI["group/flashcards/new"] = groupFlashcardNew
	wsAPI["group/flashcards/editfront"] = groupFlashcardEditFront
	wsAPI["group/flashcards/editback"] = groupFlashcardEditBack
}
