package dbs

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/x-io/gen"
)

type M map[string]string

type Combine struct {
	c *gen.Context

	params map[string]string

	where map[string]*clause
	query []string
	sort  string // field, sort

	bPage bool
	page  int
	size  int
	stamp int64

	debug bool
}

// ToCombine
func ToCombine(c *gen.Context) *Combine {
	param, _ := url.QueryUnescape(c.Request.URL.RawQuery)

	return initCombine(c, param)
	// return initCombine(c.Request.URL.RawQuery)
}

type HookFunc func(c *gen.Context) (string, bool)

var _resFuncs = make(map[string]HookFunc)

// 注册资源函数
func RegisterHook(name string, fn HookFunc) {
	_resFuncs[name] = fn
}

func initCombine(c *gen.Context, urlParams string) *Combine {
	paramMap := make(map[string]string)
	var key, val string
	var index int
	for _, value := range strings.Split(urlParams, "&") {
		key, val = value, ""
		index = strings.Index(value, "=")
		if index != -1 {
			key = value[:index]
			val = value[index+1:]
		}
		if len(val) != 0 {
			paramMap[key] = val
		}
	}

	return &Combine{
		c:      c,
		params: paramMap,
		where:  make(map[string]*clause),
		query:  []string{},
		page:   0,
		size:   15,
		stamp:  time.Now().Unix(),
	}
}

/*
key|keyFormat|#symbol#$dbtype$
*/
func (c *Combine) Filter(labes ...string) *Combine {
	for _, label := range labes {
		key, keyFormat, symbol, dbType := analyzeLabel(label)
		value, ok := c.params[key]
		if !ok {
			continue
		}

		if strings.Contains(keyFormat, "$") {
			symbol = "$"
		}

		if len(keyFormat) == 0 {
			keyFormat = key
		}
		//fmt.Println(111, label, keyFormat, value, symbol, 44, dbType)
		c.where[key] = createClause(keyFormat, value, symbol, dbType)
	}
	return c
}

func (c *Combine) SetQuery(key, value string) *Combine {
	c.params[key] = value
	return c
}

func (c *Combine) Resource(label string) *Combine {
	key, keyFormat, symbol, dbType := analyzeLabel(label)

	var value string
	var ignore bool
	if fn, ok := _resFuncs[key]; ok && c != nil {
		value, ignore = fn(c.c)
		if ignore {
			return c
		}
	}

	if strings.Contains(keyFormat, "$") {
		symbol = "$"
	}

	if len(keyFormat) == 0 {
		keyFormat = key
	}
	c.where[key] = createClause(keyFormat, value, symbol, dbType)
	return c
}

/*
	c.In("type", dbs.M{
		"sold":   "type='sold'",
		"lease":  "type='lease'",
		"xlease": "type='xlease'",
		"repair": "type='repair'",
		"@":   "1=1"
		"*":""
	})

xxx 精准匹配
* 任意匹配
@ 如果字段为控制 默认字段
*/
func (c *Combine) In(label string, values M) *Combine {
	val, ok := c.params[label]
	for key, sql := range values {
		if !ok || !strings.EqualFold(strings.ToUpper(key), strings.ToUpper(val)) {
			continue
		}

		c.query = append(c.query, "("+sql+")")
		return c
	}

	//label有值，但是没有匹配上, * 匹配任意值
	if sql, ok2 := values["*"]; ok2 && ok {
		c.query = append(c.query, "("+sql+")")
		return c
	}

	sql, ok := values["@"]
	if !ok {
		return c
	}
	c.query = append(c.query, "("+sql+")")

	return c
}

func (c *Combine) Sort(values M) *Combine {
	for key, sql := range values {
		val, ok := c.params["sort"]
		if !ok || !strings.EqualFold(strings.ToUpper(key), strings.ToUpper(val)) {
			continue
		}
		c.sort = sql
		return c
	}
	c.sort = values["@"]
	return c
}

func (c *Combine) Page(label string) *Combine {
	value, ok := c.params["stamp"]
	if !ok {
		return c
	}

	vals := strings.Split(value, "_")

	if i, err := strconv.ParseInt(vals[0], 10, 64); err == nil {
		c.stamp = i + 1
	}

	if len(vals) > 1 {
		if i, err := strconv.Atoi(vals[1]); err == nil {
			c.page = i
		}
	}
	if len(vals) > 2 {
		if i, err := strconv.Atoi(vals[2]); err == nil {
			c.size = i
		}
	}

	c.bPage = true
	if len(c.sort) == 0 {
		c.sort = label + " desc "
	}
	c.where[label] = createClause(label, strconv.Itoa(int(c.stamp)), "<=", "time")
	return c
}

func (c *Combine) Size(size int) *Combine {
	c.size = size
	return c
}

func (c *Combine) SQL(start int) (string, []interface{}) {

	var sql string
	values := make([]interface{}, 0)

	for _, cal := range c.where {
		filterSQL, items := cal.SQL(&start)
		sql += filterSQL
		values = append(values, items...)
	}

	if len(c.query) > 0 {
		sql += " AND " + strings.Join(c.query, " AND ")
	}
	sql += "\n"
	if len(c.sort) != 0 {
		sql += " ORDER BY " + c.sort + "\n"
	}

	if c.bPage {
		sql += fmt.Sprintf(" LIMIT $%d OFFSET $%d", start, start+1)
		values = append(values, c.Limit(), c.Offset())
	}

	return sql, values
}

func (c *Combine) Limit() int {
	return c.size
}

func (c *Combine) Offset() int {
	begin := (c.page - 1) * c.size
	if begin < 0 {
		begin = 0
	}
	return begin
}

func (c *Combine) Debug() *Combine {
	c.debug = true
	return c
}

func (c *Combine) Query(query string, args ...interface{}) (*sql.Rows, error) {
	sql, values := c.SQL(len(args) + 1)
	query += sql
	args = append(args, values...)
	if c.debug {
		log.Println(query, args)
	}
	return db.Query(query, args...)
}
func (c *Combine) QueryRow(query string, args ...interface{}) *sql.Row {
	sql, values := c.SQL(len(args) + 1)
	query += sql
	args = append(args, values...)
	if c.debug {
		log.Println(query, args)
	}
	return db.QueryRow(query, args...)
}

func (c *Combine) QueryPageData(query string, args ...interface{}) (*sql.Rows, *int64, error) {
	if c.params["total"] == "true" {
		c.bPage = false
		c.sort = ""
		sql, values := c.SQL(len(args) + 1)

		query = fmt.Sprintf("SELECT count(1) FROM (%s %s) v", query, sql)
		args = append(args, values...)
		if c.debug {
			log.Println(query, args)
		}
		row := db.QueryRow(query, args...)
		var count int64
		return nil, &count, row.Scan(&count)
	}

	c.bPage = true
	rows, err := c.Query(query, args...)
	return rows, nil, err
}

func (c *Combine) Exec(query string, args ...interface{}) (sql.Result, error) {

	sql, values := c.SQL(len(args) + 1)
	query += sql
	args = append(args, values...)
	if c.debug {
		log.Println(query, args)
	}

	return db.Exec(query, args...)
}

var verticalRegex = regexp.MustCompile(`\|(.*?)\|`)
var poundRegex = regexp.MustCompile(`\#(.*?)\#`)
var dollarRegex = regexp.MustCompile(`\$(.*?)\$`)

// time|created_at|#desc# return key, vertical, pound, dollar
func analyzeLabel(label string) (string, string, string, string) {
	key := label
	if index := strings.IndexFunc(key, func(value rune) bool {
		if value == '|' || value == '#' || value == '$' {
			return true
		}
		return false
	}); index != -1 {
		key = label[:index]
	}

	keyFormat := regexValue(verticalRegex, label)
	symbol := regexValue(poundRegex, label)
	dbType := regexValue(dollarRegex, label)
	return key, keyFormat, symbol, dbType
}

func regexValue(reg *regexp.Regexp, args string) string {
	val := reg.FindStringSubmatch(args)
	if len(val) > 1 {
		return val[1]
	}
	return ""
}
