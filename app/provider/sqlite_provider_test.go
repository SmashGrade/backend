package provider

import (
	"testing"
)

func TestProvider(t *testing.T) {
	prov := &SqliteProvider{}
	prov.Connect()
	t.Fatalf("Hello")
}
