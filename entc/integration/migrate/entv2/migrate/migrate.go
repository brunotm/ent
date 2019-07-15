// Code generated (@generated) by entc, DO NOT EDIT.

package migrate

import (
	"context"
	"fmt"

	"fbc/ent/dialect"
	"fbc/ent/dialect/sql/schema"
)

// Schema is the API for creating, migrating and dropping a schema.
type Schema struct {
	drv         dialect.Driver
	universalID bool
}

// NewSchema creates a new schema client.
func NewSchema(drv dialect.Driver) *Schema { return &Schema{drv: drv} }

// Create creates all schema resources.
func (s *Schema) Create(ctx context.Context, opts ...schema.MigrateOption) error {
	migrate, err := schema.NewMigrate(s.drv, opts...)
	if err != nil {
		return fmt.Errorf("ent/migrate: %v", err)
	}
	return migrate.Create(ctx, Tables...)
}
