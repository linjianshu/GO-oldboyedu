package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("abcb", "b")
	want := []string{"a", "c", " "}
	if !reflect.DeepEqual(got, want) {
		fmt.Println("测试用例失败")
		t.Errorf("want %v but got %v\n", want, got)
	}
}

func TestSplit2(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v but got %v\n", want, got)
	}
}

//测试组
func TestSplitGroup(t *testing.T) {
	type TestCase struct {
		str  string
		sep  string
		got  string
		want []string
	}
	testCase := []TestCase{
		{
			str:  "abcb",
			sep:  "b",
			want: []string{"a", "c", ""},
		},
		{
			str:  "a:b:c",
			sep:  ":",
			want: []string{"a", "b", "c"},
		},
	}
	for _, v := range testCase {
		if !reflect.DeepEqual(Split(v.str, v.sep), v.want) {
			t.Errorf("want %v but got %v\n", v.want, Split(v.str, v.sep))
		}
	}
}

func TestSingle(t *testing.T) {
	type TestCase struct {
		str  string
		sep  string
		got  string
		want []string
	}
	testCase := map[string]TestCase{
		"case1": {
			str:  "abcb",
			sep:  "b",
			want: []string{"a", "c", ""},
		},
		"case2": {
			str:  "a:b:c",
			sep:  ":",
			want: []string{"a", "b", "c"},
		},
	}

	for name, v := range testCase {
		t.Run(name, func(t *testing.T) {
			got := Split(v.str, v.sep)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("want %v but got %v\n", v.want, got)
			}
		})
	}
}

//基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}
