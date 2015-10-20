// GoJiraNotifier project main.go
package main

import (
	. "GoTeamCity/macos"
	"fmt"
	"strings"
	"time"
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
	//	Update(settings)
	//	ticker := time.NewTicker(time.Minute * 1)
	ticker := time.NewTicker(time.Second * 10)
	go UpdateRoutine(settings, ticker.C)
	RunApp()
}
