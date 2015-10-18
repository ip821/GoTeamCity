package teamcity

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	//	"strings"
)

type Connection struct {
	Url        string
	UserName   string
	Password   string
	HttpClient http.Client
}

const methodGetBuilds = "/builds/"
const methodGetChanges = "/changes/"
const methodGetInvestigations = "/investigations/"

var EmptyBuildsResponse BuildsResponse = BuildsResponse{}
var EmptyChangesResponse ChangesResponse = ChangesResponse{}
var EmptyInvestigationsResponse InvestigationsResponse = InvestigationsResponse{}

func NewConnection(url string, userName string, password string) Connection {
	return Connection{
		HttpClient: http.Client{},
		Url:        url,
		UserName:   userName,
		Password:   password,
	}
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

	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(connection.UserName, connection.Password)
	request.Header.Add("Accept", "application/json")

	response, err := connection.HttpClient.Do(request)

	if err != nil {
		return nil, err
	}

	fmt.Print("Response: ")
	fmt.Println(response)
	return response, nil
}

func deserializeBuildsRequestFromReader(reader io.Reader) (BuildsResponse, error) {
	bodyBuffer, err := ioutil.ReadAll(reader)
	if err != nil {
		return EmptyBuildsResponse, err
	}

	var buildsResponse BuildsResponse
	err = json.Unmarshal(bodyBuffer, &buildsResponse)
	if err != nil {
		return EmptyBuildsResponse, err
	}

	return buildsResponse, nil
}

func (connection *Connection) GetBuilds(buildType string) (BuildsResponse, error) {
	buildTypeLocator := fmt.Sprintf("buildType:%v", buildType)
	response, err := connection.doRequest(methodGetBuilds, "locator=sinceBuild(status:SUCCESS),branch:default:any,"+buildTypeLocator)
	if err != nil {
		return EmptyBuildsResponse, err
	}

	defer response.Body.Close()
	data, err := deserializeBuildsRequestFromReader(response.Body)
	if err != nil {
		return EmptyBuildsResponse, err
	}

	return data, nil
}

func deserializeChangesRequestFromReader(reader io.Reader) (ChangesResponse, error) {
	bodyBuffer, err := ioutil.ReadAll(reader)
	if err != nil {
		return EmptyChangesResponse, err
	}

	var changesResponse ChangesResponse
	err = json.Unmarshal(bodyBuffer, &changesResponse)
	if err != nil {
		return EmptyChangesResponse, err
	}

	return changesResponse, nil
}

func (connection *Connection) GetChanges(buildId int) (ChangesResponse, error) {
	response, err := connection.doRequest(methodGetChanges, "locator=build:"+strconv.Itoa(buildId))
	if err != nil {
		return EmptyChangesResponse, err
	}

	defer response.Body.Close()
	data, err := deserializeChangesRequestFromReader(response.Body)
	if err != nil {
		return EmptyChangesResponse, err
	}

	return data, nil
}

func deserializeInvestigationsRequestFromReader(reader io.Reader) (InvestigationsResponse, error) {
	bodyBuffer, err := ioutil.ReadAll(reader)
	if err != nil {
		return EmptyInvestigationsResponse, err
	}

	var investigationsResponse InvestigationsResponse
	err = json.Unmarshal(bodyBuffer, &investigationsResponse)
	if err != nil {
		return EmptyInvestigationsResponse, err
	}

	return investigationsResponse, nil
}

func (connection *Connection) GetInvestigations(buildType string) (InvestigationsResponse, error) {
	response, err := connection.doRequest(methodGetInvestigations, "locator=buildType:"+buildType)
	if err != nil {
		return EmptyInvestigationsResponse, err
	}

	defer response.Body.Close()
	data, err := deserializeInvestigationsRequestFromReader(response.Body)
	if err != nil {
		return EmptyInvestigationsResponse, err
	}

	return data, nil
}
