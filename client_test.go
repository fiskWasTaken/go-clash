package clash_test

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
	"github.com/fiskie/go-clash"
)

// test that our time layout is right, since we use this to convert time values to an object.
func TestTimeParsing(t *testing.T) {
	tm, _ := time.Parse(clash.TimeLayout, "20180712T110230.000Z")
	assert.Equal(t, int64(1531393350), tm.Unix())
}
