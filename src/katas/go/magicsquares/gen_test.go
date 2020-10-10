package magicsquares

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	Magic3 = Matrix{
		{4, 9, 2},
		{3, 5, 7},
		{8, 1, 6},
	}
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		msg  string
		in   int
		want []Matrix
	}{
		{"empty: 1", 1, nil},
		{"empty: 2", 2, nil},
		{"8 squares: 3", 3, []Matrix{
			Magic3,
		}},
	}

	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			got := Generate(tt.in)
			require.Len(t, got, len(tt.want)*8)
			for i, m := range tt.want {
				assert.Contains(t, got, m, i)
			}
		})
	}
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate(3)
	}
}

func TestCheck(t *testing.T) {
	tests := []struct {
		msg  string
		in   Matrix
		want bool
	}{
		{"empty", nil, false},
		{"size 1", Matrix{{1}}, false},
		{"size 2", Matrix{
			{1, 2},
			{3, 4},
		}, false},
		{"smallest", Matrix{
			{6, 7, 2},
			{1, 5, 9},
			{8, 3, 4},
		}, true},
	}

	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			assert.Equal(t, tt.want, Check(tt.in))
		})
	}
}
