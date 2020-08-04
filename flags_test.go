package translate

import (
	"strconv"
	"testing"
)

func TestFlagsMapToSupportedLanguages(t *testing.T) {
	for k, v := range flags {
		if _, ok := languages[v]; !ok {
			t.Errorf("%v: %q not found in supported languages", strconv.QuoteToASCII(k), v)
		}
	}
}
