package overwatch

// PlayerStats represents a response from https://owapi.net/api/v3/u/:battletag/stats
type PlayerStats struct {
	KR      PlayerRegion `json:"kr"`
	EU      PlayerRegion `json:"eu"`
	US      PlayerRegion `json:"us"`
	Any     interface{}  `json:"any"`
	Request struct {
		APIVer int    `json:"api_ver"`
		Route  string `json:"route"`
	} `json:"_request"`
}
