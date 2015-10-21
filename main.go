// GoJiraNotifier project main.go
package main

import (
	"fmt"
	"strings"
	"time"
	. "GoTeamCity/macos"
	. "GoTeamCity/windows"
)

func main() {
	InitMacOsPackage()
	InitWindowsPackage()

	settings := Settings{}
	err := settings.Load()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	urlToOpen := strings.Replace(settings.Url, "httpAuth/app/rest", "", -1)
	StartAppUi(urlToOpen)
	//	ticker := time.NewTicker(time.Minute * 1)
	ticker := time.NewTicker(time.Second * 10)
	go UpdateRoutine(settings, ticker.C, UpdateCallbackFunc)
	RunApp()
}

func UpdateCallbackFunc(status Status){
	UpdateAppUi(int(status.BuildsStatus))
}