package dbs

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
)

func TestRowsC(t *testing.T) {
	// sort.Strings(arr)
	// key := strings.Join(arr, ";")
	// fmt.Println(arr, key)
	str := "time=1619280000%7C1619366399"

	fmt.Println(url.QueryUnescape(str))
	// Init("postgres://zcs:123456@127.0.0.1:5432/zcs?sslmode=disable")

	// c := initCombine("name_小,sort_100,time_1518819832|1618819832", "state", "time_desc", &Page{Index: 1, Size: 10})
	// c.Filters("name#like#", "sort#>=#", "time#between#|created_at|$time$")
	// c.Query("state", "status = 'ONLINE'")
	// c.Sort("time|created_at|")
	// // c.Sort("time|created_at|#asc#")
	// c.Page().Debug()
	// rows, err := QueryC(c, "SELECT name FROM shop WHERE 1=1 ")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// for rows.Next() {
	// 	var name string
	// 	if err := rows.Scan(&name); err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println(name)
	// }
}

func TestClause(t *testing.T) {
	//  sku ~* '版本:(v1|v2)'
	start := 0
	keyFormat := "sku"
	value := "默认;颜色:红,黑;版本:新车"
	sqls := make([]string, 0, 10)
	items := make([]string, 0, 10)
	for _, val := range strings.Split(value, ";") {
		tmp := strings.Split(val, ":")
		if len(tmp) == 1 {
			sqls = append(sqls, fmt.Sprintf("%s = $%d", keyFormat, start))
			items = append(items, val)
		} else if len(tmp) == 2 {
			sqls = append(sqls, fmt.Sprintf("%s ~* $%d", keyFormat, start))
			items = append(items, fmt.Sprintf("%s:(%s)", tmp[0], strings.ReplaceAll(tmp[1], ",", "|")))
		}
		start++
	}

	fmt.Println("SQL:", sqls)
	fmt.Println("items:", items)
}

func TestRules(t *testing.T) {
	Init("postgres://zcs:123456@127.0.0.1:5432/zcs3?sslmode=disable")
	// c := initCombine("name=小&sort=100&time=1518819832|1618819832&state=online&created_at&sort=asc&stamp=1619227756_2")
	c := initCombine(nil, "name=小&sort=100&time=1518819832|1618819832&state=online|offline&created_at&sort=asc&stamp=1619227756_2")
	// c.Debug()
	c.Filter("name#like#", "state|status|")
	c.In("name", M{
		"online":  "status='ONLINE'",
		"offline": "status!='ONLINE'",
		"null":    "status is null",
		"*":       "1=1",
	})
	c.Sort(M{
		"asc": "sort asc",
		"@":   "created_at desc",
	})
	// c.Page("created_at")
	rows, err := c.Query("SELECT name FROM shop WHERE 1=1 and 100 > $1", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(name)
	}
	return
}
