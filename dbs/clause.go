package dbs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lib/pq"
)

type clause struct {
	keyFormat string //
	dbType    string // int string uuid point time json
	symbol    string // = like in between
	value     string
	sqlTpl    string
}

/*
	value:
		1|2|3_in =>  in ($,$,$)
		a_% => like '%'||$1||'%'
		4|5_[] => between $ and $
	dbType: int string uuid point time json
*/
//
func createClause(key, value, symbol, dbType string) *clause {

	c := &clause{
		keyFormat: key,
		value:     value,
		symbol:    symbol,
		dbType:    dbType,
	}
	return c
}

func (c *clause) SetSQLTpl(sql string) {
	c.sqlTpl = sql
}

func (c *clause) SQL(start *int) (string, []interface{}) {

	sql, symbol, value := "", c.symbol, c.value
	index := strings.Index(c.value, "_") //a|b|c_%
	if index != -1 && len(c.symbol) == 0 {
		symbol = c.value[index+1:]
		value = c.value[:index]
	}

	items := make([]interface{}, 0)
	for _, v := range strings.Split(value, ",") {
		items = append(items, convertValue(c.dbType, v))
	}

	switch strings.ToLower(symbol) {
	case "%", "like":
		sql = "("
		for i := 0; i < len(items); i++ {
			if i != 0 {
				sql += " OR "
			}

			sql += "UPPER(" + c.keyFormat + ") like '%'||UPPER($" + strconv.Itoa(*start) + ")||'%' "
			*start++
		}
		sql += ")"
	case "lower", "LOWER":
		sql = fmt.Sprintf(`LOWER(%s) = $%d`, c.keyFormat, *start)
		*start++
	case "upper", "UPPER":
		sql = fmt.Sprintf(`UPPER(%s) = $%d`, c.keyFormat, *start)
		*start++
	case "[]", "between":
		sql = fmt.Sprintf(`%s between $%d AND $%d`, c.keyFormat, *start, *start+1)
		*start += 2
	case ">", ">=", "<", "<=", "!=":
		sql = fmt.Sprintf(`%s %s $%d`, c.keyFormat, symbol, *start)
		*start++
	case "extend": // sku ~* '版本:(v1|v2),颜色:(红|蓝)'
		sqls := make([]string, 0, 10)
		items = []interface{}{}
		//定制:12芯哈喽,颜色:白
		for _, val := range strings.Split(value, ",") {
			sqls = append(sqls, c.keyFormat+" ~* $"+strconv.Itoa(*start))
			// items = append(items, val[strings.Index(val, ":")+1:]+"(/|$)")
			items = append(items, val[strings.Index(val, ":")+1:])

			*start++
		}
		sql = strings.Join(sqls, " AND ")
	case "array_contains":
		sql = fmt.Sprintf(`%s @> $%d`, c.keyFormat, *start)
		*start++
		items = []interface{}{pq.Array(items)}
	case "any":
		sql = fmt.Sprintf(`%s = any($%d)`, c.keyFormat, *start)
		*start++
		items = []interface{}{pq.Array(items)}
	case "nin":
		if len(items) == 1 {
			sql = fmt.Sprintf(`%s != $%d`, c.keyFormat, *start)
			*start++
			break
		}

		items = []interface{}{pq.Array(items)}
		sql = fmt.Sprintf(`not %s = any($%d)`, c.keyFormat, *start)
		*start++
	case "?":
		sql = fmt.Sprintf(`%s ? $%d`, c.keyFormat, *start)
		*start++
	case "$":
		values := []interface{}{}
		sql = ""
		tpl := c.keyFormat
		for {
			idx := strings.Index(tpl, "$")
			if idx == -1 {
				sql += tpl
				break
			}
			sql += tpl[:idx]
			sql += "$" + strconv.Itoa(*start)
			tpl = tpl[idx+1:]
			values = append(values, items[0])
			*start++
		}

		items = values
	case "=", "in":
		fallthrough
	default:
		if len(items) == 1 {
			sql = fmt.Sprintf(`%s = ($%d)`, c.keyFormat, *start)
			*start++
			break
		}

		items = []interface{}{pq.Array(items)}
		sql = fmt.Sprintf(`%s = any($%d)`, c.keyFormat, *start)
		*start++
	}

	return " AND (" + sql + ")", items
}

func convertValue(dbType, value string) interface{} {
	switch strings.ToLower(dbType) {
	case "bool":
		if value == "true" {
			return Bool(true)
		}
		return Bool(false)
	case "int":
		i, _ := strconv.Atoi(value)
		return Int(i)
	case "json":
		return JSON([]byte(value))
	case "point":
		return Point(value)
	case "string":
		return String(value)
	case "time":
		i, _ := strconv.ParseInt(value, 10, 64)
		return Time(i)
	case "uuid":
		return String(value)
	default:
		return value
	}
}
