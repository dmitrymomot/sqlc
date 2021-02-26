package sqlite

import "github.com/dmitrymomot/sqlc/internal/sql/catalog"

func NewCatalog() *catalog.Catalog {
	c := catalog.New("main")
	return c
}
