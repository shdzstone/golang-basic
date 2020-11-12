package main

import (
	"golang-basic/lang/basic/basic/var"
	"testing"
)

/*
1.传统测试
* 测试数据和测试逻辑写在一起
* 出错信息不明确
* 一旦一个数据出错测试全部结束
2.表格驱动测试
* 分离的测试数据和测试逻辑
* 明确的出错信息
* 可以部分失败
* go语言的语法使得我们更容易实践表格驱动测试
3.测试实例
* testing.T的使用:t.Errorf提供了一个报错的接口
* 运行测试：使用命令行go ping .
4.debug
* 直接使用goland的debug功能即可
*/

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

//func TestSum(t *testing.T) {
//	assert.Equal(3, sum(3, 4))
//	assert.Equal(3, sum(3, 4))
//	assert.Equal(3, sum(3, 4))
//	assert.Equal(3, sum(3, 4))
//	assert.Equal(3, sum(3, 4))
//
//}

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}
	for _, tt := range tests {
		if actual := _var.CalcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("CalcTriangle(%d,%d);"+"got %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}
