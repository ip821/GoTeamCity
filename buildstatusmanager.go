// buildstatusmanager
package main

import (
	tc "GoTeamCity/teamcity"
	//	"fmt"
	. "github.com/ahmetalpbalkan/go-linq"
)

type BuildStatusManager struct {
	Connection tc.Connection
	UserName   string
}

type BuildStatus struct {
	BuildType                 string
	LastFailedBuildChanges    []tc.Change
	BuildTypeInvestigations   []tc.Investigation
	IsBroken                  bool
	IsPotentiallyBrokenByUser bool
	IsAssigned                bool
}

var emptyBuildStatus BuildStatus = BuildStatus{}

func NewBuildStatusManager(c tc.Connection, userName string) BuildStatusManager {
	return BuildStatusManager{Connection: c, UserName: userName}
}

func (m *BuildStatusManager) GetBuildStatus(buildType string) (BuildStatus, error) {
	builds, err := m.Connection.GetBuilds(buildType)
	if err != nil {
		return emptyBuildStatus, err
	}

	buildStatus := BuildStatus{BuildType: buildType}

	if len(builds.Builds) == 0 {
		return buildStatus, nil
	}

	buildStatus.IsBroken = true

	firstFailedBuild, _, _ := From(builds.Builds).Last()
	changes, err := m.Connection.GetChanges(firstFailedBuild.(tc.Build).Id)
	buildStatus.LastFailedBuildChanges = changes.Changes

	containsUserNameInChanges, err := From(changes.Changes).Where(func(s T) (bool, error) { return s.(tc.Change).UserName == m.UserName, nil }).Any()
	buildStatus.IsPotentiallyBrokenByUser = containsUserNameInChanges
	if err != nil {
		return emptyBuildStatus, err
	}

	investigations, err := m.Connection.GetInvestigations(buildType)
	if err != nil {
		return emptyBuildStatus, err
	}

	buildStatus.BuildTypeInvestigations = investigations.Investigations

	containsUserNameInInvestiagations, err := From(investigations.Investigations).Where(func(s T) (bool, error) { return s.(tc.Investigation).Assignee.UserName == m.UserName, nil }).Any()
	buildStatus.IsAssigned = containsUserNameInInvestiagations
	return buildStatus, nil
}
