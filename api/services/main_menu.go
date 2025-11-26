package services

import (
	"context"

	dbquery "github.com/SONEsee/go-echo/pkg/db-pkg/db-query"
	dbschema "github.com/SONEsee/go-echo/pkg/db-pkg/db-schema"
)

func GetMainMenuByID(ctx context.Context, id int) (*dbschema.MainMenuDGSchema, error) {
	result, err := dbquery.GetMainMenuByID(ctx, id)
	return result, err
}

func GetAllMainMenusService(ctx context.Context) ([]dbschema.MainMenuDGSchema, error) {

	result, err := dbquery.GetAllMainMenus(ctx)
	return result, err
}

func GetMainTester(ctx context.Context) ([]dbschema.MainMenuDGSchema, error) {
	result, err := dbquery.GetTestMainmenu(ctx)
	return result, err
}
