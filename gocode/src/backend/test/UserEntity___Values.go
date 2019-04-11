package test

const (
	UserEntityInsertWRONGBODY = `{
									"username":"testing",
									"email":"usercompany.com",
									"password":"123456789",
									"image":"http://www.xxx.jpg",
									"admin":true,
									"enable":false
								}`
	UserEntityInsertSUCCESS = `{
									"username":"testing",
									"email":"user@company.com",
									"password":"123456789",
									"image":"http://www.xxx.jpg",
									"admin":true,
									"enable":false
								}`
	UserEntityVerifyWRONGBODY = `{
									"username":"testing"
									"password":"123456789"
								}`
	UserEntityVerifyWRONGPASSWORD = `{
									"username":"testing",
									"password":"12345678910"
								}`
	UserEntityVerifySUCCESS = `{
									"username":"testing",
									"password":"123456789"
								}`
	UserEntityUpdateSingleWRONGBODY = `{
									"password":"123456789",
									"position":"email"
									"value":"testing@company.com"
								}`
	UserEntityUpdateSingleWRONGEMAIL = `{
									"password":"123456789",
									"position":"email",
									"value":"testingcompany.com"
								}`
	UserEntityUpdateSingleSUCCESS = `{
									"password":"123456789",
									"position":"image",
									"value":"http://www.xxx.jpg"
								}`
	UserEntityUpdateWRONGBODY = `{
									"email":"testing@company.com"
									"image":"http://www.xxx.jpg"
								}`
	UserEntityUpdateWRONGIMAGE = `{
									"email":"testing@company.com",
									"image":"httww.xxx.jpg"
								}`
	UserEntityUpdateSUCCESS = `{
									"email":"testing@company.com",
									"image":"http://www.xxx.jpg"
								}`
	UserEntityDisableWRONGBODY = `{
									"password""123456789"
								}`
	UserEntityDisableWRONGPASSWORD = `{
									"password":"1234568201"
								}`
	UserEntityDisableSUCCESS = `{
									"password":"123456789"
								}`
	UserEntityDeleteSUCCESS = `{
									"password":"123456789"
								}`

	ResponseUserEntityGetUnique  = `{"username":"testing","email":"testing@company.com","image":"http://www.xxx.jpg","admin":true,"enable":false}`
	ResponseUserEntityAllEnabled = `null`
	ResponseUserEntityGetAll     = `[{"username":"testing","email":"testing@company.com","image":"http://www.xxx.jpg","admin":true,"enable":false}]`
)
