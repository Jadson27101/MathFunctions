package Functions

import (
	"strconv"
	"strings"
)

func Inv(val uint64) uint64 {
	var stringVal = strconv.FormatUint(val, 2)
	var flag = 0
	var result uint64 = 0
	//fmt.Println(stringVal)
	if val == 0 {
		stringVal = "1111111111111111111111111111111111111111111111111111111111111111"
		res, _ := strconv.ParseUint(stringVal, 10, len(stringVal))
		result = res
	} else {
		m := strings.NewReplacer("1", "0", "0", "1")
		for i := 0; i < len(stringVal); i++ {
			if string([]rune(m.Replace(stringVal))[i]) == "0" {
				flag++
			} else {
				break
			}
		}
		//fmt.Println(m.Replace(stringVal))
		a, _ := strconv.ParseUint(m.Replace(stringVal), 2, len(m.Replace(stringVal))-flag)
		result = a
	}
	return result
}
