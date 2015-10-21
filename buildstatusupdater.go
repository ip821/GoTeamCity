// buildstatusupdater.go
package main

import (
	tc "GoTeamCity/teamcity"
	"fmt"
	"strings"
	"time"
)

type StatusEnum int

const (
	BSuccess                   StatusEnum = 1
	BFailed                               = 2
	BFailedAndUserChangesFound            = 3
	BAssigned                             = 4
)

type Status struct {
	BuildsStatus StatusEnum
}

type UpdateCallback func(statuc Status)

func UpdateRoutine(settings Settings, tickerC <-chan time.Time, updateCallback UpdateCallback) {
	for range tickerC {
		Update(settings, updateCallback)
	}
}

func Update(settings Settings, updateCallback UpdateCallback) {
	loginData := strings.Split(settings.LoginData, ":")
	user := loginData[0]
	pwd := loginData[1]
	connection := tc.NewConnection(settings.Url, user, pwd)
	manager := NewBuildStatusManager(connection, user)

	buildTypes := strings.Split(settings.BuildTypes, ",")
	fmt.Println(buildTypes)

	overallStatus := Status{}
	overallStatus.BuildsStatus = BSuccess
	for _, item := range buildTypes {
		fmt.Println(item)
		buildStatus, err := manager.GetBuildStatus(item)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(buildStatus)

		if buildStatus.IsAssigned {
			overallStatus.BuildsStatus = BAssigned
			break
		}

		if buildStatus.IsPotentiallyBrokenByUser {
			overallStatus.BuildsStatus = BFailedAndUserChangesFound
		}

		if buildStatus.IsBroken && overallStatus.BuildsStatus != BFailedAndUserChangesFound {
			overallStatus.BuildsStatus = BFailed
		}
	}
	fmt.Println(overallStatus)
	updateCallback(overallStatus)
}
