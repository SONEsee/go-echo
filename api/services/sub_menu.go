package services

import (
	"context"

	dbquery "github.com/SONEsee/go-echo/pkg/db-pkg/db-query"
	dbschema "github.com/SONEsee/go-echo/pkg/db-pkg/db-schema"
)

func GateAllWhitSubmenu(ctx context.Context) ([]dbschema.SubMenuSchema, error) {
	result, err := dbquery.GetSubmenuWhitAll(ctx)
	return result, err
}
