package clash_test

import (
	"testing"
	"time"
	"github.com/overwolfmobile/go-clash"
	"github.com/stretchr/testify/assert"
)

// test that our time layout is right, since we use this to convert time values to an object.
func TestTimeParse(t *testing.T) {
	tm, _ := time.Parse(clash.TimeLayout, "20180712T110230.000Z")
	assert.Equal(t, int64(1531393350), tm.Unix())
}
