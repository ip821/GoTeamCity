// GoJiraNotifier project main.go
package main

import (
	"fmt"
	jira "GoJiraNotifier/Jira"
)

func main() {
	settings := Settings{}
	err := settings.Load()
	if err != nil{
		fmt.Println(err.Error())
		return
	}	

	connection := jira.NewConnection(settings.Url, settings.LoginData)

	userObject, err := connection.GetCurrentUser()
	if err != nil{
		fmt.Println(err.Error())
		return
	}	
	fmt.Println(userObject)

	issuesArray, err := connection.GetIssuesByCriteria(settings.FilterSearchString)
	if err != nil{
		fmt.Println(err.Error())
		return
	}	
	
	fmt.Println(issuesArray)
//	fmt.Scanln()
}
