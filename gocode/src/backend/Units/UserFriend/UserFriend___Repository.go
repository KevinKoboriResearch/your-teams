package UserFriend

import (
	"backend/HyperText"
	"backend/Interface"
	"sort"
	"strings"
)

//___________________________ INSERT _________________________________________//
func (r UserFriendRepository) InsertUserFriend(uf UserFriend) int {

	ue.Username = uf.User2

	interf := HyperText.StructValidate(&ue)

	if interf != nil {
		return 0
	}

	uf.User1 = strings.ToLower(uf.User1)
	uf.User1Enable = true
	uf.User2Enable = false
	uf.Operator = 1
	strs := []string{uf.User1, uf.User2}
	strsAfter := []string{uf.User1, uf.User2}
	sort.Strings(strsAfter)

	if strs[0] != strsAfter[0] {
		uf.User1 = strsAfter[0]
		uf.User2 = strsAfter[1]
		uf.User1Enable = false
		uf.User2Enable = true
		uf.Operator = 2
	}

	operator := uf.Operator

	err := Interface.FindUnitDB(DOCNAME, &ufind, "user1", uf.User1, "user2", uf.User2)

	if err == nil {
		if operator != ufind.Operator {
			uf.User1Enable = true
			uf.User2Enable = true
			Interface.UpdateUnitDB(DOCNAME, uf, "user1", uf.User1, "user2", uf.User2)
			return 2
		}
		return 0
	}

	err = Interface.InsertDB(DOCNAME, uf)

	return 1
}

//___________________________ DELETE _________________________________________//
func (r UserFriendRepository) DeleteUserFriend(uf UserFriend) (err error) {
	uf.User1 = strings.ToLower(uf.User1)

	err = Interface.DeleteUnitDB(DOCNAME, uf, "user1", uf.User1, "user2", uf.User2)

	return
}
