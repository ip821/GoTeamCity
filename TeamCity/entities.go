// JiraIssue.go
package teamcity

type Build struct {
	Id          int    `json:"id"`
	BuildTypeId string `json:"buildTypeId"`
	Number      string `json:"number"`
	Status      string `json:"status"`
	State       string `json:"state"`
	Branch      string `json:"branchName"`
}

type Change struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
}

type Investigation struct {
	UserName string `json:"username"`
}
