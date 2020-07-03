package bowling

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResult(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input   string
		want    int
		wantErr string
	}{
		{"--------------------", 0, ""},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()

			got, err := Result(tt.input)
			assert.Equal(t, tt.want, got)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
				return
			}
			assert.NoError(t, err)
		})
	}
}
