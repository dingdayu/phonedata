package phonedata

import (
	"testing"
)

func TestIsPhone(t *testing.T) {
	if !IsPhone("13398829873") {
		t.Error("测试失败")
	}
	if IsPhone("1339882983") {
		t.Error("测试失败")
	}
}

func TestTotalRecord(t *testing.T) {
	if TotalRecord() == 0 {
		t.Error("测试失败")
	}
}

func TestVersion(t *testing.T) {
	if Version() == "" {
		t.Error("测试失败")
	}
}

func TestLoadDataFile(t *testing.T) {
	err := LoadDataFile(PhoneDatFile)
	if err != nil {
		t.Error("测试失败")
	}
	is := IsPhone("13398829873")
	if !is {
		t.Error("测试失败")
	}
}

func BenchmarkIsPhone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		is := IsPhone("13398829873")
		if !is {
			b.Error("测试失败")
		}
	}
}

func BenchmarkVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Version()
	}
}

func BenchmarkFind(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Find("13398829873")
		if err != nil {
			b.Error("测试失败")
		}
	}
}
