package db_test

import (
	"reflect"
	"testing"

	c "github.com/SmashGrade/backend/app/config"
	"github.com/SmashGrade/backend/app/db"
)

// Test the migration of the database with the in memory provider
func TestProviderMigration(t *testing.T) {
	// Create a new provider
	provider := db.NewMockProvider()
	// Migrate the database
	err := db.Migrate(provider)
	// Check if the migration was successful
	if err != nil {
		t.Errorf("Migration failed: %s", err)
	}
}

// Tests the selection of the database provider based on the connection string prefix
func TestProviderDBSelection(t *testing.T) {
	tests := []struct {
		name             string
		connectionString string
		want             string
	}{
		{
			name:             "SQLite",
			connectionString: "sqlite://:memory:",
			want:             "*db.SQLiteProvider",
		},
		{
			name:             "Postgres",
			connectionString: "postgres://localhost:5432/db",
			want:             "*db.PostgresProvider",
		},
		{
			name:             "MySQL",
			connectionString: "mysql://localhost:3306/db",
			want:             "*db.MySQLProvider",
		},
	}
	for _, testData := range tests {
		t.Run(testData.name, func(t *testing.T) {
			config := c.NewAPIConfig()
			config.DBConnectionStr = testData.connectionString
			config.Connect = false
			got := db.NewProvider(config)
			if reflect.TypeOf(got).String() != testData.want {
				t.Errorf("NewProvider() = %v, want %v", reflect.TypeOf(got), testData.want)
			}
		})
	}
}
