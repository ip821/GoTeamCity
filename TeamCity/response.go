package teamcity

type BuildsResponse struct {
	Builds []Build `json:"build"`
}

type ChangesResponse struct {
	Changes []Change `json:"change"`
}

type InvestigationsResponse struct {
	Investigations []Investigation`json:"investigation"`
}
