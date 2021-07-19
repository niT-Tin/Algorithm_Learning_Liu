package main

import (
	"fmt"
	"testing"
)

func BenchmarkMaxValue(b *testing.B) {

	ms := GenerateMap()
	fmt.Printf("The map is \n")
	showMap(ms)
	maxVal_v1(0, 0, &ms)
	b.Logf("The max value is %d\n", Sum) // 第五个地图就开始爆内存

}
