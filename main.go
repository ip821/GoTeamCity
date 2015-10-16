// GoJiraNotifier project main.go
package main

import (
	"fmt"
	tc "GoTeamCity/teamcity"
)

func main() {
	settings := Settings{}
	err := settings.Load()
	if err != nil{
		fmt.Println(err.Error())
		return
	}	

	connection := tc.NewConnection(settings.Url, settings.LoginData)

	buildsObject, err := connection.GetBuilds()
	if err != nil{
		fmt.Println(err.Error())
		return
	}	
	fmt.Println(buildsObject)

//	fmt.Scanln()
}
