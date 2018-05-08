package models

import (
	"testing"

	"github.com/gorilla/websocket"
)

func TestGroupExists(t *testing.T) {
	nonexistentGroup := "noname_1234"
	if GroupExists(nonexistentGroup) {
		t.Error("GroupExists did not find the group correctly")
	}

	if !GroupExists("hello_9797") {
		t.Error("GroupExists did not find the group correctly")
	}
}

func TestUserExistsInGroup(t *testing.T) {
	users["testtoken"] = &user{
		Name:  "testUsername",
		Token: "testtoken",
	}
	usrList := userList{}
	testuser := users["testtoken"]
	usrList.add(testuser)
	groups["testgroup"] = &group{
		Name:  "testgroup",
		Users: usrList,
	}
	if !UserExistsInGroup("testtoken", "testgroup") {
		t.Error("UserExistsInGroup failed to correctly find user")
	}
}

func TestAddGroup(t *testing.T) {
	users["testtoken"] = &user{
		Name: "testUsername",
	}
	groupid := AddGroup("testGroup", "testtoken")
	if !GroupExists(groupid) {
		t.Error("TestAddGroup failed to add group")
	}
}

func TestAddUserToGroupMap(t *testing.T) {
	users["testtoken"] = &user{
		Name:  "testUsername",
		Token: "testtoken",
	}
	groups["testgroup"] = &group{
		Name: "testgroup",
	}
	AddUserToGroupMap("nonexistenttoken", "testgroup")
	AddUserToGroupMap("testtoken", "nonexistentgroup")
	AddUserToGroupMap("testtoken", "testgroup")
	AddUserToGroupMap("testtoken", "testgroup")

	if !UserExistsInGroup("testtoken", "testgroup") {
		t.Error("TestAddUserToGroup failed to add user to group")
	}
}

func TestRemoveUserFromGroup(t *testing.T) {
	users["testtoken"] = &user{
		Name:  "testUsername",
		Token: "testtoken",
	}
	groups["testgroup"] = &group{
		Name: "testgroup",
	}
	//Used non DB part of test since we want to test on server functionality
	newUser := users["testtoken"]
	grp := groups["testgroup"]
	grp.addUser(newUser)
	RemoveUserFromGroup("nonexistenttoken", "testgroup")
	RemoveUserFromGroup("testtoken", "nonexistentgroup")
	RemoveUserFromGroup("testtoken", "testgroup")
	RemoveUserFromGroup("testtoken", "testgroup")
	if UserExistsInGroup("testtoken", "testgroup") {
		t.Error("TestRemoveUserFromGroup failed to remove user from group")
	}
}

func TestGetConnectionsInGroup(t *testing.T) {
	conn1 := &websocket.Conn{}
	users["testtoken1"] = &user{
		WsConn: conn1,
		Name:   "testUser1",
		Token:  "testtoken1",
	}
	conn2 := &websocket.Conn{}
	users["testtoken2"] = &user{
		WsConn: conn2,
		Name:   "testUser2",
		Token:  "testtoken2",
	}
	groups["testgroup"] = &group{
		Name: "testgroup",
	}
	grp := groups["testgroup"]
	testuser1 := users["testtoken1"]
	testuser2 := users["testtoken2"]
	grp.addUser(testuser1)
	grp.addUser(testuser2)
	connections := GetConnectionsInGroup("testgroup")
	if testuser1.WsConn != connections[0] {
		t.Error("GetConnectionsInGroup returned wrong values.")
	}
	if testuser2.WsConn != connections[1] {
		t.Error("GetConnectionsInGroup returned wrong values.")
	}
}

func TestGetOtherConnectionsInGroup(t *testing.T) {
	conn1 := &websocket.Conn{}
	users["testtoken1"] = &user{
		WsConn: conn1,
		Name:   "testUser1",
		Token:  "testtoken1",
	}
	conn2 := &websocket.Conn{}
	users["testtoken2"] = &user{
		WsConn: conn2,
		Name:   "testUser2",
		Token:  "testtoken2",
	}
	groups["testgroup"] = &group{
		Name: "testgroup",
	}
	grp := groups["testgroup"]
	testuser1 := users["testtoken1"]
	testuser2 := users["testtoken2"]
	grp.addUser(testuser1)
	grp.addUser(testuser2)
	connections := GetOtherConnectionsInGroup("testtoken1", "testgroup")
	if testuser2.WsConn != connections[0] {
		t.Error("GetOtherConnectionsInGroup return wrong values.")
	}
}

func TestGAddUser(t *testing.T) {
	grp := &group{
		Name: "testgroup",
	}
	testuser := &user{
		Name: "testuser",
	}
	grp.addUser(testuser)
	if !grp.Users.containsUsername("testuser") {
		t.Error("AddUser failed to add user to group struct")
	}
}

func TestGRemoveUser(t *testing.T) {
	grp := &group{
		Name: "testgroup",
	}
	testuser := &user{
		Name: "testuser",
	}
	grp.addUser(testuser)
	grp.removeUser(testuser)
	if grp.Users.containsUsername("testuser") {
		t.Error("AddUser failed to add user to group struct")
	}
}

func TestGLAdd(t *testing.T) {
	grpLst := groupList{}
	grp := &group{
		Name: "testgroup",
	}
	grpLst.add(grp)
	if !grpLst.contains(grp) {
		t.Error("GL add failed to add group to group list")
	}
}

func TestGLRemove(t *testing.T) {
	grpLst := groupList{}
	grp := &group{
		Name: "testgroup",
	}
	grpLst.add(grp)
	grpLst.remove(grp)
	if grpLst.contains(grp) {
		t.Error("GL add failed to add group to group list")
	}
}
