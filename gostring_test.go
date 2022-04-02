package gostringbenchmark

import (
	"testing"
)

func BenchmarkStringPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringPlus()
	}
}
func BenchmarkStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringJoin()
	}
}
func BenchmarkStringBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBuffer()
	}
}
func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBuilder()
	}
}

// -----------------------------------------------------------------
const STRING = "GO字符串链接速度测试!"

func initStrings(N int) []string {
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = STRING
	}
	return s
}

func BenchmarkStringPlus2(b *testing.B) {
	p := initStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringPlus2(p)
	}
}

func BenchmarkStringJoin2(b *testing.B) {
	p := initStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringJoin2(p)
	}
}

func BenchmarkStringBuffer2(b *testing.B) {
	p := initStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringBuffer2(p)
	}
}

func BenchmarkStringBuilder2(b *testing.B) {
	p := initStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringBuilder2(p)
	}
}
