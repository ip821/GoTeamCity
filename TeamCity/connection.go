package teamcity

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"fmt"
	"strings"
)

type Connection struct {
	Url        string
	LoginData  string
	HttpClient http.Client
}

const methodGetBuilds = "/builds"

var EmptyBuildsResponse BuildsResponse = BuildsResponse{}

func NewConnection(url string, loginData string) Connection {
	return Connection{
		HttpClient: http.Client{},
		Url:        url,
		LoginData:  loginData,
	}
}

func deserializeBuildsRequestFromReader(reader io.Reader) (BuildsResponse, error) {
	bodyBuffer, err := ioutil.ReadAll(reader)
	if err != nil{
		return EmptyBuildsResponse, err
	}

	var buildsResponse BuildsResponse
	err = json.Unmarshal(bodyBuffer, &buildsResponse)	
	if err != nil{
		return EmptyBuildsResponse, err
	}
	
	return buildsResponse, nil
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
	
	loginData := strings.Split(connection.LoginData, ":")
	request.SetBasicAuth(loginData[0], loginData[1])
	request.Header.Add("Accept", "application/json")
		
	fmt.Println(loginData[0])
	fmt.Println(loginData[1])

	response, err := connection.HttpClient.Do(request)

	if err != nil{
		return nil, err
	}
	
	fmt.Print("Response: ")
	fmt.Println(response) 
	return response, nil
}

func (connection *Connection) GetBuilds() (BuildsResponse, error) {
	response, err := connection.doRequest(methodGetBuilds)
	if err != nil{
		return EmptyBuildsResponse, err
	}
	
	defer response.Body.Close()
	data, err := deserializeBuildsRequestFromReader(response.Body)
	if err != nil{
		return EmptyBuildsResponse, err
	}
		
	return data, nil
}
