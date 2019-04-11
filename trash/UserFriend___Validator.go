package UserFriend

import (
	//"be/HyperText"
	"be/Interface"
	//	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2/bson"
	//"strings"
)

//__ FRIEND EXIST ____________________________________________________________//
func ValidateFriendUsed(c *UserFriendController, uf UserFriend) bool {
	Interface.OpenSession(DOCNAME).Find(bson.M{"user1": uf.User1, "user2": uf.User2}).One(&uf)
	return uf.User1 == ""
}
