package gostringbenchmark

import (
	"bytes"
	"strings"
)

func StringPlus() string {
	return "GO" + "字符串" + "链接" + "速度" + "测试" + "!"
}

func StringJoin() string {
	stringList := []string{"GO", "字符串", "链接", "速度", "测试", "!"}
	return strings.Join(stringList, "")
}

func StringBuffer() string {
	var b bytes.Buffer
	b.WriteString("GO")
	b.WriteString("字符串")
	b.WriteString("链接")
	b.WriteString("速度")
	b.WriteString("测试")
	b.WriteString("!")
	return b.String()
}

func StringBuilder() string {
	var b strings.Builder
	b.WriteString("GO")
	b.WriteString("字符串")
	b.WriteString("链接")
	b.WriteString("速度")
	b.WriteString("测试")
	b.WriteString("!")
	return b.String()
}

// ----------------------------------------------

func StringPlus2(p []string) string {
	var s string
	l := len(p)
	for i := 0; i < l; i++ {
		s += p[i]
	}
	return s
}

func StringJoin2(p []string) string {
	return strings.Join(p, "")
}

func StringBuffer2(p []string) string {
	var b bytes.Buffer
	l := len(p)
	for i := 0; i < l; i++ {
		b.WriteString(p[i])
	}
	return b.String()
}

func StringBuilder2(p []string) string {
	var b strings.Builder
	l := len(p)
	for i := 0; i < l; i++ {
		b.WriteString(p[i])
	}
	return b.String()
}
