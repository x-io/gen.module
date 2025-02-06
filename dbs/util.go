package dbs

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//CreateHistory CreateHistory
func CreateHistory(staffID, t string) JSON {
	return []byte(fmt.Sprintf(`{"staffid":"%s","type":"%s","time":%d}`, staffID, t, time.Now().Unix()))
}

//CheckPage CheckPage
func CheckPage(pageSize int, pageIndex int) int {
	begin := (pageIndex - 1) * pageSize
	if begin < 0 {
		begin = 0
	}
	return begin
}

//GenOrderID 生成日期型ID
func GenOrderID() string {
	return strings.Replace(time.Now().Format("20060102150405.000000"), ".", "", 14) //+ string(krand(9, KC_RAND_KIND_NUM))
}

//Krand 随机字符串
func krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	isall := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isall {
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}
