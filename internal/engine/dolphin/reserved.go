package dolphin

import "strings"

// https://dev.mysql.com/doc/refman/8.0/en/keywords.html
func (p *Parser) IsReservedKeyword(s string) bool {
	switch strings.ToLower(s) {
	case "accessible":
	case "add":
	case "all":
	case "alter":
	case "analyze":
	case "and":
	case "as":
	case "asc":
	case "asensitive":
	case "before":
	case "between":
	case "bigint":
	case "binary":
	case "blob":
	case "both":
	case "by":
	case "call":
	case "cascade":
	case "case":
	case "change":
	case "char":
	case "character":
	case "check":
	case "collate":
	case "column":
	case "condition":
	case "constraint":
	case "continue":
	case "convert":
	case "create":
	case "cross":
	case "cube":
	case "cume_dist":
	case "current_date":
	case "current_time":
	case "current_timestamp":
	case "current_user":
	case "cursor":
	case "database":
	case "databases":
	case "day_hour":
	case "day_microsecond":
	case "day_minute":
	case "day_second":
	case "dec":
	case "decimal":
	case "declare":
	case "default":
	case "delayed":
	case "delete":
	case "dense_rank":
	case "desc":
	case "describe":
	case "deterministic":
	case "distinct":
	case "distinctrow":
	case "div":
	case "double":
	case "drop":
	case "dual":
	case "each":
	case "else":
	case "elseif":
	case "empty":
	case "enclosed":
	case "escaped":
	case "except":
	case "exists":
	case "exit":
	case "explain":
	case "false":
	case "fetch":
	case "first_value":
	case "float":
	case "float4":
	case "float8":
	case "for":
	case "force":
	case "foreign":
	case "from":
	case "fulltext":
	case "function":
	case "generated":
	case "get":
	case "grant":
	case "group":
	case "grouping":
	case "groups":
	case "having":
	case "high_priority":
	case "hour_microsecond":
	case "hour_minute":
	case "hour_second":
	case "if":
	case "ignore":
	case "in":
	case "index":
	case "infile":
	case "inner":
	case "inout":
	case "insensitive":
	case "insert":
	case "int":
	case "int1":
	case "int2":
	case "int3":
	case "int4":
	case "int8":
	case "integer":
	case "interval":
	case "into":
	case "io_after_gtids":
	case "io_before_gtids":
	case "is":
	case "iterate":
	case "join":
	case "json_table":
	case "key":
	case "keys":
	case "kill":
	case "lag":
	case "last_value":
	case "lateral":
	case "lead":
	case "leading":
	case "leave":
	case "left":
	case "like":
	case "limit":
	case "linear":
	case "lines":
	case "load":
	case "localtime":
	case "localtimestamp":
	case "lock":
	case "long":
	case "longblob":
	case "longtext":
	case "loop":
	case "low_priority":
	case "master_bind":
	case "master_ssl_verify_server_cert":
	case "match":
	case "maxvalue":
	case "mediumblob":
	case "mediumint":
	case "mediumtext":
	case "middleint":
	case "minute_microsecond":
	case "minute_second":
	case "mod":
	case "modifies":
	case "natural":
	case "not":
	case "no_write_to_binlog":
	case "nth_value":
	case "ntile":
	case "null":
	case "numeric":
	case "of":
	case "on":
	case "optimize":
	case "optimizer_costs":
	case "option":
	case "optionally":
	case "or":
	case "order":
	case "out":
	case "outer":
	case "outfile":
	case "over":
	case "partition":
	case "percent_rank":
	case "precision":
	case "primary":
	case "procedure":
	case "purge":
	case "range":
	case "rank":
	case "read":
	case "reads":
	case "read_write":
	case "real":
	case "recursive":
	case "references":
	case "regexp":
	case "release":
	case "rename":
	case "repeat":
	case "replace":
	case "require":
	case "resignal":
	case "restrict":
	case "return":
	case "revoke":
	case "right":
	case "rlike":
	case "row":
	case "rows":
	case "row_number":
	case "schema":
	case "schemas":
	case "second_microsecond":
	case "select":
	case "sensitive":
	case "separator":
	case "set":
	case "show":
	case "signal":
	case "smallint":
	case "spatial":
	case "specific":
	case "sql":
	case "sqlexception":
	case "sqlstate":
	case "sqlwarning":
	case "sql_big_result":
	case "sql_calc_found_rows":
	case "sql_small_result":
	case "ssl":
	case "starting":
	case "stored":
	case "straight_join":
	case "system":
	case "table":
	case "terminated":
	case "then":
	case "tinyblob":
	case "tinyint":
	case "tinytext":
	case "to":
	case "trailing":
	case "trigger":
	case "true":
	case "undo":
	case "union":
	case "unique":
	case "unlock":
	case "unsigned":
	case "update":
	case "usage":
	case "use":
	case "using":
	case "utc_date":
	case "utc_time":
	case "utc_timestamp":
	case "values":
	case "varbinary":
	case "varchar":
	case "varcharacter":
	case "varying":
	case "virtual":
	case "when":
	case "where":
	case "while":
	case "window":
	case "with":
	case "write":
	case "xor":
	case "year_month":
	case "zerofill":
	default:
		return false
	}
	return true
}
