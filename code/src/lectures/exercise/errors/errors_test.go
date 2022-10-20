package timeparse

import "testing"

func TestParseTime(t *testing.T) {
	// we're gonna testing multiple inputs, so we're gonna make a test table:
	table := []struct {
		time string
		ok   bool
	}{
		{"19:00:12", true},
	}

	for _, data := range table {
		_, err := ParseTime(data.time)
		if data.ok && err != nil {
			t.Errorf("%v: %v, error should be nill", data.time, err)
		}
	}
}
