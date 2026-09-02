package spatial

import (
	"math"
	"testing"
)

func TestDownmixToBinaural(t *testing.T) {
	l, r := DownmixToBinaural(SurroundChannels{FL: 1.0, FR: 1.0, FC: 1.0})
	if math.Abs(l-1.707) > 0.01 || math.Abs(r-1.707) > 0.01 {
		t.Errorf("unexpected binaural downmix: %f, %f", l, r)
	}
}
