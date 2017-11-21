package overwatch

import (
	"strings"

	"github.com/aerogo/http/client"
)

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

// HighestSkillRating returns a player's highest skill rating in any region.
func (stats *PlayerStats) HighestSkillRating() (highest int, highestTier string) {
	if stats.EU.Stats.Competitive.OverallStats.Comprank > highest {
		highest = stats.EU.Stats.Competitive.OverallStats.Comprank
		highestTier = stats.EU.Stats.Competitive.OverallStats.Tier
	}

	if stats.US.Stats.Competitive.OverallStats.Comprank > highest {
		highest = stats.US.Stats.Competitive.OverallStats.Comprank
		highestTier = stats.US.Stats.Competitive.OverallStats.Tier
	}

	if stats.KR.Stats.Competitive.OverallStats.Comprank > highest {
		highest = stats.KR.Stats.Competitive.OverallStats.Comprank
		highestTier = stats.KR.Stats.Competitive.OverallStats.Tier
	}

	return highest, highestTier
}

// GetPlayerStats returns a player's statistics grouped by region.
func GetPlayerStats(battleTag string) (*PlayerStats, error) {
	battleTag = strings.Replace(battleTag, "#", "-", 1)
	playerStats := &PlayerStats{}
	_, err := client.Get("https://owapi.net/api/v3/u/" + battleTag + "/stats").EndStruct(playerStats)
	return playerStats, err
}
