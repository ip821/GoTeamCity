package tc

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"fmt"
	"net/url"
)

type Connection struct {
	Url        string
	LoginData  string
	HttpClient http.Client
}

const getCurrebtUserMethod = "/2/myself"
const search = "/2/search"

var EmptyJiraUser = JiraUser{}

func NewConnection(url string, loginData string) Connection {
	return Connection{
		HttpClient: http.Client{},
		Url:        url,
		LoginData:  loginData,
	}
}

func deserializeUserRequestFromReader(reader io.Reader) (jiraUserResponse, error) {
	bodyBuffer, err := ioutil.ReadAll(reader)
	if err != nil{
		return EmptyJiraUserResponse, err
	}

	var userResponse jiraUserResponse
	err = json.Unmarshal(bodyBuffer, &userResponse)	
	if err != nil{
		return EmptyJiraUserResponse, err
	}
	
	return userResponse, nil
}

func (connection *Connection) doRequest(method string, params ...string) (*http.Response, error) {
	fullUrl := connection.Url + method

	if params != nil {
		for i, item := range params {
			if i == 0 {
				fullUrl += "?"
			} else {
				fullUrl += "&"
			}
			fullUrl += item
		}
	}

	fmt.Println("Request url: " + fullUrl)

	request, err := http.NewRequest("GET", fullUrl, nil)

	if err != nil{
		return nil, err
	}
	
	request.Header.Add("Authorization: Basic ", connection.LoginData)

	response, err := connection.HttpClient.Do(request)

	if err != nil{
		return nil, err
	}
	
	fmt.Print("Response: ")
	fmt.Println(response) 
	return response, nil
}

func (connection *Connection) GetCurrentUser() (JiraUser, error) {
	response, err := connection.doRequest(getCurrebtUserMethod)
	if err != nil{
		return EmptyJiraUser, err
	}
	
	defer response.Body.Close()
	data, err := deserializeUserRequestFromReader(response.Body)
	if err != nil{
		return EmptyJiraUser, err
	}
		
	return data.JiraUser, nil
}

func (connection *Connection) GetIssuesByCriteria(query string) ([]JiraIssue, error) {
	response, err := connection.doRequest(search, "fields=id,key", "jql="+url.QueryEscape(query))
	if err != nil{
		return nil, err
	}
	
	defer response.Body.Close()
	searchResult, err := deserializeSearchRequestFromReader(response.Body)
	if err != nil{
		return nil, err
	}

	return searchResult.Issues, nil
}
