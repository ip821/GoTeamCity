// GoJiraNotifier project main.go
package main

import (
	. "GoTeamCity/macos"
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

	urlToOpen := strings.Replace(settings.Url, "httpAuth/app/rest", "", -1)
	StartAppUi(urlToOpen)

	loginData := strings.Split(settings.LoginData, ":")
	user := loginData[0]
	pwd := loginData[1]
	connection := tc.NewConnection(settings.Url, user, pwd)
	manager := NewBuildStatusManager(connection, user)

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
