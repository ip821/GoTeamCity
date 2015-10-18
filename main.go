// GoJiraNotifier project main.go
package main

import (
	tc "GoTeamCity/teamcity"
	"fmt"
	"strings"
)

func main() {
	settings := Settings{}
	err := settings.Load()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	loginData := strings.Split(settings.LoginData, ":")
	connection := tc.NewConnection(settings.Url, loginData[0], loginData[1])
	manager := NewBuildStatusManager(connection, loginData[0])

	buildTypes := strings.Split(settings.BuildTypes, ",")
	fmt.Println(buildTypes)

	for _, item := range buildTypes {
		fmt.Println(item)
		buildStatus, err := manager.GetBuildStatus(item)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(buildStatus)
	}

	//	fmt.Scanln()
}
