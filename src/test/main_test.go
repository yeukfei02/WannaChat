package test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const rootURL = "https://wanna-chat-api.herokuapp.com/api"
const userSignupURL = rootURL + "/user/signup"
const userLoginURL = rootURL + "/user/login"
const getAllUsersURL = rootURL + "/user"
const getUserByIDURL = rootURL + "/user/1"
const createGroupURL = rootURL + "/group/create-group"
const getAllGroupsURL = rootURL + "/group"
const getGroupByIDURL = rootURL + "/group/2"
const createMembershipURL = rootURL + "/membership/create-membership"
const getAllMembershipsURL = rootURL + "/membership"
const getMembershipByGroupIDURL = rootURL + "/membership/get-membership-by-group-id"

const email = "test@test.com"
const password = "test"
const groupLabel = "groupLabel-test"

var token string

var responseBody []byte

type LoginResponse struct {
	Message string
	Token   string
}

func TestUserSignup(t *testing.T) {
	requestBody := getUserRequestBody()
	assertResult(t, "POST", userSignupURL, requestBody)
}

func TestUserLogin(t *testing.T) {
	assert := assert.New(t)

	requestBody := getUserRequestBody()
	assertResult(t, "POST", userLoginURL, requestBody)

	var loginResponse LoginResponse
	json.Unmarshal([]byte(responseBody), &loginResponse)
	assert.NotNil(loginResponse.Message)
	assert.NotNil(loginResponse.Token)
	token = loginResponse.Token
}

func TestGetAllUsers(t *testing.T) {
	assertResult(t, "GET", getAllUsersURL, nil)
}

func TestGetUserByID(t *testing.T) {
	assertResult(t, "GET", getUserByIDURL, nil)
}

func TestCreateGroup(t *testing.T) {
	requestBody := getGroupRequestBody()
	assertResult(t, "POST", createGroupURL, requestBody)
}

func TestGetAllGroups(t *testing.T) {
	assertResult(t, "GET", getAllGroupsURL, nil)
}

func TestGetGroupByID(t *testing.T) {
	assertResult(t, "GET", getGroupByIDURL, nil)
}

func TestCreateMembership(t *testing.T) {
	requestBody := getMemberRequestBody()
	assertResult(t, "POST", createMembershipURL, requestBody)
}

func TestGetAllMemberships(t *testing.T) {
	assertResult(t, "GET", getAllMembershipsURL, nil)
}

func TestGetMembershipByGroupID(t *testing.T) {
	assert := assert.New(t)

	request, err := http.NewRequest("GET", getMembershipByGroupIDURL, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	queryParams := request.URL.Query()
	queryParams.Add("groupId", "2")
	request.URL.RawQuery = queryParams.Encode()

	assert.Nil(err)

	client := &http.Client{}
	response, responseErr := client.Do(request)
	assert.Nil(responseErr)

	t.Log("response.Status = ", response.Status)
	assert.NotNil(response.Status)

	body, _ := ioutil.ReadAll(response.Body)
	responseBody = body
	t.Log("response.Body = ", string(body))
	assert.NotNil(string(body))
}

func getUserRequestBody() io.Reader {
	values := map[string]string{"email": email, "password": password}
	jsonValue, _ := json.Marshal(values)
	return bytes.NewBuffer(jsonValue)
}

func getGroupRequestBody() io.Reader {
	values := map[string]string{"groupLabel": groupLabel}
	jsonValue, _ := json.Marshal(values)
	return bytes.NewBuffer(jsonValue)
}

func getMemberRequestBody() io.Reader {
	values := map[string]int{"userFk": 1, "groupFk": 2}
	jsonValue, _ := json.Marshal(values)
	return bytes.NewBuffer(jsonValue)
}

func assertResult(t *testing.T, requestType string, url string, requestBody io.Reader) {
	assert := assert.New(t)

	request, err := http.NewRequest(requestType, url, requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	assert.Nil(err)

	client := &http.Client{}
	response, responseErr := client.Do(request)
	assert.Nil(responseErr)

	t.Log("response.Status = ", response.Status)
	assert.NotNil(response.Status)

	body, _ := ioutil.ReadAll(response.Body)
	responseBody = body
	t.Log("response.Body = ", string(body))
	assert.NotNil(string(body))
}
