package models

import "fmt"

func createGroupInDB(groupid string) (err error) {
	fmt.Println("Creating group " + groupid)
	_, err = db.Exec("CREATE TABLE " + groupid + " (Admin varchar(50), userList varchar(20), ipAddress varchar(50), User varchar(20), Clock datetime, Message varchar(255), Whiteboard LONGTEXT)")
	if err != nil {
		panic(err)
	}
	return
}

func putAdminInGroupDB(groupid string, admin string) (err error) {
	adminstmt, err := db.Prepare("INSERT INTO " + groupid + " (Admin) VALUES (?)")
	if err != nil {
		panic(err)
	}
	_, err = adminstmt.Exec(admin)

	return
}

func WriteChatToDB(groupid string, username string, msg string) (err error) {

	chatstmt, err := db.Prepare("INSERT INTO " + groupid + " (user, Clock, Message) VALUES (?, current_timestamp(),  ?)")
	fmt.Println("Inserting chat to DB...")
	if err != nil {
		panic(err)
	}
	_, err = chatstmt.Exec(username, msg)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return
}
