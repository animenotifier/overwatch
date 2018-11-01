package overwatch_test

import (
	"testing"

	"github.com/animenotifier/overwatch"
	"github.com/stretchr/testify/assert"
)

func TestPlayerStats(t *testing.T) {
	stats, err := overwatch.GetPlayerStats("Aky#2908")
	assert.NoError(t, err)
	assert.NotNil(t, stats)

	// This can actually return 0-values for inactive players,
	// so we're not checking the return values.
	stats.HighestSkillRating()
}
