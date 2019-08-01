package phonedata_test

import (
	"fmt"
	"phonedata"
)

// 此注释将会被展示在页面上
// 此函数将被展示在OverView区域
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
