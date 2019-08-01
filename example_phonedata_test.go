package phonedata_test

import (
	"fmt"
	"phonedata"
)

// 查询一个手机号的归属地信息
func Example() {
	info, _ := phonedata.Find("13298181000")
	fmt.Println(info)
	// Output:
	// PhoneNum: 13298181000
	// AreaZone: 0371
	// CardType: 中国联通
	// City: 郑州
	// ZipCode: 450000
	// Province: 河南
}
