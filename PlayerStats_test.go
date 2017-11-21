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

	rating, tier := stats.HighestSkillRating()
	assert.NotZero(t, rating)
	assert.NotEmpty(t, tier)
}
