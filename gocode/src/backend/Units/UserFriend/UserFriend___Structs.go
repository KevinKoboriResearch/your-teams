package UserFriend

type UserFriendRepository struct {
}

type UserFriendController struct {
	UserFriendRepository UserFriendRepository
}

type UserFriend struct {
	User1       string `json:"user1" validate:"required,username-exist"`
	User2       string `json:"user2"`
	WhoBlock    string `json:"whoblock"`
	User1Enable bool   `json:"user1enable"`
	User2Enable bool   `json:"user2enable"`
	Operator    int    `json:"operator"`
}

type UserFriendUpdate struct {
	User1       string `json:"user1" bson:"user1,omitempty"`
	User2       string `json:"user2" bson:"user2,omitempty"`
	WhoBlock    string `json:"whoblock" bson:"whoblock,omitempty"`
	User1Enable bool   `json:"user1enable" bson:"user1enable,omitempty"`
	User2Enable bool   `json:"user2enable" bson:"user2enable,omitempty"`
	Operator    int    `json:"operator" bson:"operator,omitempty"`
}

type Userfirends []UserFriend
