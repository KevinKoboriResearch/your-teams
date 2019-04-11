package test

const (
	UserEntityInsertEXAMPLE1 = `{
									"username":"EXAMPLE1",
									"email":"example1@company.com",
									"password":"123456789",
									"image":"http://www.xxx.jpg"
								}`
	UserEntityInsertEXAMPLE2 = `{
									"username":"EXAMPLE2",
									"email":"example2@company.com",
									"password":"123456789",
									"image":"http://www.xxx.jpg"
								}`
	UserEntityVerifyEXAMPLE1 = `{
									"username":"EXAMPLE1",
									"password":"123456789"
								}`
	UserEntityVerifyEXAMPLE2 = `{
									"username":"EXAMPLE2",
									"password":"123456789"
								}`
	UserFriendRequestWRONGBODY = `{
									"user1""EXAMPLE1"
								}`
	UserFriendRequestEXAMPLE1 = `{
									"user1":"EXAMPLE1"
								}`
	UserFriendAcceptEXAMPLE2 = `{
									"user1":"EXAMPLE2"
								}`
	UserFriendRequestALREADYEXIST = `{
									"user1":"EXAMPLE2",
									"operator":1
								}`
	UserFriendRequestNOTEXIST = `{
									"user1":"BLABLABLA"
								}`
)
