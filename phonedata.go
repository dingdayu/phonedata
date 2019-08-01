package phonedata

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"phonedata/data"
)

const (
	CMCC             byte = iota + 0x01 // 中国移动
	CUCC                                // 中国联通
	CTCC                                // 中国电信
	CTCCV                               // 电信虚拟运营商
	CUCCV                               // 联通虚拟运营商
	CMCCV                               // 移动虚拟运营商
	IntLen           = 4
	CharLen          = 1
	HeadLength       = 8
	PhoneIndexLength = 9
	PhoneDatFile     = "data/phone.dat"
)

type PhoneRecord struct {
	PhoneNum string
	Province string
	City     string
	ZipCode  string
	AreaZone string
	CardType string
}

var (
	content  []byte
	CardType = map[byte]string{
		CMCC:  "中国移动",
		CUCC:  "中国联通",
		CTCC:  "中国电信",
		CTCCV: "中国电信虚拟运营商",
		CUCCV: "中国联通虚拟运营商",
		CMCCV: "中国移动虚拟运营商",
	}
	totalLen, firstOffset int32
)

func init() {
	var err error
	file := os.Getenv("PHONE_DATA_FILE")
	if file != "" {
		content, err = ioutil.ReadFile(file)
	} else {
		content, err = data.Asset(PhoneDatFile)
	}
	if err != nil {
		panic(err)
	}
	totalLen = int32(len(content))
	firstOffset = toInt32(content[IntLen:HeadLength])
}

func LoadDataFile(file string) (err error) {
	content, err = ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	totalLen = int32(len(content))
	firstOffset = toInt32(content[IntLen:HeadLength])
	return nil
}

// PhoneRecord 格式化输出归属地信息
func (pr PhoneRecord) String() string {
	return fmt.Sprintf("PhoneNum: %s\nAreaZone: %s\nCardType: %s\nCity: %s\nZipCode: %s\nProvince: %s\n", pr.PhoneNum, pr.AreaZone, pr.CardType, pr.City, pr.ZipCode, pr.Province)
}

func toInt32(b []byte) int32 {
	if len(b) < 4 {
		return 0
	}
	return int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16 | int32(b[3])<<24
}

// Version phone.dat version
// 获取记录文件头部版本号
func Version() string {
	return string(content[0:IntLen])
}

// TotalRecord phone.dat record total
// 	通过 索引长度 / 每个索引的长度 = 索引记录数
func TotalRecord() int32 {
	return (int32(len(content)) - firstRecordOffset()) / PhoneIndexLength
}

func firstRecordOffset() int32 {
	return toInt32(content[IntLen:HeadLength])
}

// IsPhone check phone is right
func IsPhone(v string) bool {
	reg := `^1([358][0-9]|14[57]|17[0678])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(v)
}

// Find search phone info
// 通过对索引区域进行二分查找，得到记录偏移位置和卡类型
// 再通过记录偏移获取号码详细信息
func Find(phoneNum string) (pr *PhoneRecord, err error) {
	phoneNum = strings.TrimPrefix(phoneNum, "+")
	if len(phoneNum) < 7 || len(phoneNum) > 11 {
		return nil, errors.New("illegal phone length")
	}

	var left int32
	phoneSevenInt, err := strconv.Atoi(phoneNum[0:7])
	if err != nil {
		return nil, errors.New("illegal phone number")
	}
	phoneSevenInt32 := int32(phoneSevenInt)
	right := (totalLen - firstOffset) / PhoneIndexLength
	for {
		if left > right {
			break
		}
		mid := (left + right) / 2
		offset := firstOffset + mid*PhoneIndexLength
		if offset >= totalLen {
			break
		}
		curPhone := toInt32(content[offset : offset+IntLen])
		recordOffset := toInt32(content[offset+IntLen : offset+IntLen*2])
		cardType := content[offset+IntLen*2 : offset+IntLen*2+CharLen][0]
		switch {
		case curPhone > phoneSevenInt32:
			right = mid - 1
		case curPhone < phoneSevenInt32:
			left = mid + 1
		default:
			cbyte := content[recordOffset:]
			endOffset := int32(bytes.Index(cbyte, []byte("\000")))
			cdata := bytes.Split(cbyte[:endOffset], []byte("|"))
			cardStr, ok := CardType[cardType]
			if !ok {
				cardStr = "未知运营商"
			}
			pr = &PhoneRecord{
				PhoneNum: phoneNum,
				Province: string(cdata[0]),
				City:     string(cdata[1]),
				ZipCode:  string(cdata[2]),
				AreaZone: string(cdata[3]),
				CardType: cardStr,
			}
			return
		}
	}
	return nil, errors.New("phone's data not found")
}
