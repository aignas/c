package dbinsert

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func BenchmarkBulk(b *testing.B) {
	data := make([]Row, 1000)
	for i := range data {
		data[i] = Row{
			ID:   uuid.Must(uuid.NewV4()),
			Name: uuid.Must(uuid.NewV4()).String(),
		}
	}

	db, err := sqlx.ConnectContext(context.Background(), "mysql", "test:bench@/testbulk")
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		err := InsertBulk(context.Background(), db, data)
		require.NoError(b, err)
	}
}

func BenchmarkSingle(b *testing.B) {
	data := make([]Row, 1000)
	for i := range data {
		data[i] = Row{
			ID:   uuid.Must(uuid.NewV4()),
			Name: uuid.Must(uuid.NewV4()).String(),
		}
	}

	db, err := sqlx.ConnectContext(context.Background(), "mysql", "test:bench@/testbulk")
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		err := InsertCached(context.Background(), db, data)
		require.NoError(b, err)
	}
}
