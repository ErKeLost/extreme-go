package UnitTest

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	res := add(123, 5)
	if res != 128 {
		t.Errorf("expect %d, actual %d", 128, res)
	}
}

func TestAdd2(t *testing.T) {
	if testing.Short() {
		t.Skip("short模式下跳过")
	}
	re := add(1, 5)
	if re != 6 {
		t.Errorf("expect %d, actual %d", 6, re)
	}
}

// 表格驱动测试编码

func TestAd(t *testing.T) {
	var dataset = []struct {
		a   int
		b   int
		out int
	}{
		{1, 2, 434},
		{1, 2, 343214},
		{1, 2, 3},
		{1, 4565, 434},
	}

	for _, value := range dataset {
		re := add(value.a, value.b)
		if re != value.out {
			t.Error("error with add func")
		}
	}
}

// 性能测试

func BenchmarkAdd(bb *testing.B) {
	var a, b, c int
	a = 999
	b = 8888
	c = 7777

	for i := 0; i < bb.N; i++ {
		if actual := add(a, b); actual != c {
			fmt.Printf("%d + %d, except: %d, actual:%d", a, b, c, actual)
		}
	}
}
