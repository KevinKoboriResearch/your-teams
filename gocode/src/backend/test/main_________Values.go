package test

import (
	"backend/HyperText"
	"bytes"
	"gopkg.in/mgo.v2"
	"net/http"
	"testing"
)

const (
	MONGO_IP       = "localhost"
	MONGO_PORT     = ":27017"
	MONGO_HOST     = MONGO_IP + MONGO_PORT
	APPJASON       = "application/json"
	CHARSET_UTF_8  = "charset=UTF-8"
	APPJASON_UTF_8 = APPJASON + "; " + CHARSET_UTF_8
)

var (
	client  http.Client
	session *mgo.Session
	auth    *http.Response
	success = HyperText.TestResponses["success"]
)

//__ SEND POST _______________________________________________________________//
func sendPost(path string, typeReq string, entity string) (resp *http.Response) {
	r := bytes.NewReader([]byte(entity))
	resp, _ = http.Post(path, typeReq, r)
	return
}

//__ SEND PUT ________________________________________________________________//
func sendPut(path string, request string, token string) (resp *http.Response, err error) {
	c := &client
	req, _ := http.NewRequest("PUT", path, stringToReader(request))
	if token != "" {
		req.Header.Set("Authorization", token)
	}
	return c.Do(req)
}

//__ SEND GET ________________________________________________________________//
func sendGet(path string) (req *http.Response) {
	req, _ = http.Get(path)
	return req
}

//__ SEND DELETE ______________________________________________________________//
func sendDelete(path string, request string, token string) (resp *http.Response, err error) {
	c := &client
	req, _ := http.NewRequest("DELETE", path, stringToReader(request))
	if token != "" {
		req.Header.Set("Authorization", token)
	}
	return c.Do(req)
}

//__ RESPONSE TO STRING ______________________________________________________//
func responseToString(resp *http.Response) (response string) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	response = buf.String()
	return
}

//__ REQUEST TO STRING _______________________________________________________//
func requestToString(req *http.Request) (response string) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	response = buf.String()
	return
}

//__ STRING TO READER ________________________________________________________//
func stringToReader(s string) *bytes.Reader {
	return bytes.NewReader([]byte(s))
}

//__ DELETE COLLECTION _______________________________________________________//
func compareResults(t *testing.T, response string, expected string) {
	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}
}

//__ DELETE COLLECTION _______________________________________________________//
func DeleteCollection(collection string) (err error) {
	session, err = mgo.Dial(MONGO_HOST)
	err = session.DB("your_teams").C(collection).Remove(nil)
	return
}
