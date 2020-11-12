package main

import (
	"golang-basic/lang/container/map/nonrepeating"
	"testing"
)

/*
1.代码覆盖测试
* 使用编译器的覆盖率测试可以看到每个文件夹下每个文件的的代码覆盖率
* run "ping" with coverage
* >go ping -coverprofile=c.out
* >go tool cover -html=c.out

2.性能测试
* 使用*testing.B来做性能测试
* >go ping -bench .

3.性能调优
* 1.go ping -bench . -cpuprofile cpu.out #获取性能数据
* 2.go tool pprof cpu.out  web   #使用pprof查看具体的性能数据，并进行代码优化
* 3.分析慢在哪里

4.优化代码
* 根据pprof分析报告进行代码调优
*/

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		//Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		//Edge cases
		{"", 0},
		{"a", 1},
		{"bbbbbbbbbb", 1},
		{"abcabcabcabcd", 4},

		//Chinese support
		{"一二三四一二三", 4},
		{"黑化肥挥发发灰会花飞黑化肥挥发发灰会花飞", 9},
	}
	for _, tt := range tests {
		if actual := nonrepeating.LengthOfNonRepeatingSubStr(tt.s); actual != tt.ans {
			t.Errorf("got %d for input %s;"+"expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发灰会花飞黑化肥挥发灰会花飞"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	ans := 9

	b.Logf("len(s):%d", len(s))
	b.ResetTimer()

	//循环次数由系统设定
	for i := 0; i < b.N; i++ {
		if actual := nonrepeating.LengthOfNonRepeatingSubStr(s); actual != ans {
			b.Errorf("got %d for input %s;"+"expected %d", actual, s, ans)
		}
	}
}
