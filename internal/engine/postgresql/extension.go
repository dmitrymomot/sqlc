// Code generated by sqlc-pg-gen. DO NOT EDIT.

package postgresql

import (
	"github.com/dmitrymomot/sqlc/internal/engine/postgresql/contrib"
	"github.com/dmitrymomot/sqlc/internal/sql/catalog"
)

func loadExtension(name string) *catalog.Schema {
	switch name {
	case "adminpack":
		return contrib.Adminpack()
	case "amcheck":
		return contrib.Amcheck()
	case "btree_gin":
		return contrib.BtreeGin()
	case "btree_gist":
		return contrib.BtreeGist()
	case "citext":
		return contrib.Citext()
	case "cube":
		return contrib.Cube()
	case "dblink":
		return contrib.Dblink()
	case "earthdistance":
		return contrib.Earthdistance()
	case "file_fdw":
		return contrib.FileFdw()
	case "fuzzystrmatch":
		return contrib.Fuzzystrmatch()
	case "hstore":
		return contrib.Hstore()
	case "intagg":
		return contrib.Intagg()
	case "intarray":
		return contrib.Intarray()
	case "isn":
		return contrib.Isn()
	case "lo":
		return contrib.Lo()
	case "ltree":
		return contrib.Ltree()
	case "pageinspect":
		return contrib.Pageinspect()
	case "pg_buffercache":
		return contrib.PgBuffercache()
	case "pgcrypto":
		return contrib.Pgcrypto()
	case "pg_freespacemap":
		return contrib.PgFreespacemap()
	case "pg_prewarm":
		return contrib.PgPrewarm()
	case "pg_stat_statements":
		return contrib.PgStatStatements()
	case "pgstattuple":
		return contrib.Pgstattuple()
	case "pg_trgm":
		return contrib.PgTrgm()
	case "pg_visibility":
		return contrib.PgVisibility()
	case "postgres_fdw":
		return contrib.PostgresFdw()
	case "seg":
		return contrib.Seg()
	case "sslinfo":
		return contrib.Sslinfo()
	case "tablefunc":
		return contrib.Tablefunc()
	case "tcn":
		return contrib.Tcn()
	case "unaccent":
		return contrib.Unaccent()
	case "uuid-ossp":
		return contrib.UuidOssp()
	case "xml2":
		return contrib.Xml2()
	}
	return nil
}
