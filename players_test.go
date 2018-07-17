package clash_test

import (
	"testing"
	"github.com/fiskie/go-clash"
	"github.com/stretchr/testify/assert"
)

var teamWin = clash.Battle{
	Team: []clash.BattlePlayer{
		{Tag: "#111", Crowns: 2},
		{Tag: "#112", Crowns: 2},
	},
	Opponent: []clash.BattlePlayer{
		{Tag: "#113", Crowns: 1},
		{Tag: "#114", Crowns: 1},
	},
}

var draw = clash.Battle{
	Team: []clash.BattlePlayer{
		{Tag: "#111", Crowns: 1},
		{Tag: "#112", Crowns: 1},
	},
	Opponent: []clash.BattlePlayer{
		{Tag: "#113", Crowns: 1},
		{Tag: "#114", Crowns: 1},
	},
}

func TestBattle_PlayerByTag(t *testing.T) {
	player4, err := teamWin.PlayerByTag("#114")
	assert.Nil(t, err)
	assert.Equal(t, "#114", player4.Tag)

	_, err = teamWin.PlayerByTag("#115")
	assert.NotNil(t, err)
}

func TestBattle_Outcome(t *testing.T) {
	outcome := teamWin.Outcome()

	assert.False(t, outcome.IsDraw)
	assert.Equal(t, outcome.Winners[0].Tag, teamWin.Team[0].Tag)
	assert.Equal(t, outcome.Losers[0].Tag, teamWin.Opponent[0].Tag)

	drawOutcome := draw.Outcome()
	assert.True(t, drawOutcome.IsDraw)
}
