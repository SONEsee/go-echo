package dbschema

type MainMenuDGSchema struct {
	ID       int    `db:"id" json:"id"`
	NameMenu string `db:"mame_menu" json:"name_menu"`
	IconMenu string `db:"icon_menu" json:"icon_menu"`
}

type MainMenuWhitSubMenuSchema struct {
	ID       int             `db:"id" json:"id"`
	NameMenu string          `db:"mame_menu" json:"name_menu"`
	IconMenu string          `db:"icon_menu" json:"icon_menu"`
	SubMenu  []SubMenuSchema `json:"sub_menus"`
}
