package domain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test validators
func TestValidators(t *testing.T) {
	d := Domain{}
	failures := map[int][][]string{
		1: {},
		2: {{}},
		3: {{"ATL", "EWR"}, {"SFO", "SFO"}},
		4: {{"EWR", "IND"}, {"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}},
		5: {{"ATL", "EWR"}, {"SFO", "IND", "ATL"}},
		6: {{"ATL", "EWR"}, {"SFO"}},
	}
	for i, v := range failures {
		res, err := d.Sort(v)
		assert.Empty(t, res)
		assert.Error(t, err, fmt.Sprint("failed on ", i))
	}
}

func TestExpectedInputOutput(t *testing.T) {
	d := Domain{}
	successes := map[int][][]string{
		1: {{"SFO", "EWR"}},
		2: {{"ATL", "EWR"}, {"SFO", "ATL"}},
		3: {{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}},
	}
	expected := []string{"SFO", "EWR"}
	for i, v := range successes {
		res, err := d.Sort(v)
		assert.NoError(t, err, fmt.Sprint("failed on ", i))
		assert.NotEmpty(t, res, fmt.Sprint("failed on ", i))
		assert.EqualValues(t, expected, res, fmt.Sprint("failed on ", i))
	}
}
