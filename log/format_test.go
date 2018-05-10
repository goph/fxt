package fxlog

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseFormat(t *testing.T) {
	formats := []format{JsonFormat, LogfmtFormat}

	for _, f := range formats {
		t.Run(f.String(), func(t *testing.T) {
			assert.Contains(t, formatMap, f)
			assert.Contains(t, formatNameMap, f.String())

			pf, err := ParseFormat(f.String())
			require.NoError(t, err)
			assert.Equal(t, f, pf)
		})
	}
}
