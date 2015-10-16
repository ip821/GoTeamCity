// JiraIssue.go
package teamcity

import(
	"fmt"
)

type Build struct{
	Id int`json:"id"`
	BuildTypeId string`json:"buildTypeId"`
	Number string`json:"number"`
	Status string`json:"status"`
	State string`json:"state"`
}

func(b Build) String() string{
	return fmt.Sprintf("\n{Id: %v, BuildTypeId: %v, Number: %v, Status: %v, State: %v}", b.Id, b.BuildTypeId, b.Number, b.Status, b.State)
}


