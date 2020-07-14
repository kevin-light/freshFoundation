package split

import (
	"reflect"
	"testing"
)

//	>go test -v 运行详细信息  go test运行  go test -cover 测试覆盖率
// $ go test -cover -coverprofile=c.out  覆盖率相关的信息输出到当前文件夹下面的c.out文件中， 然后我们执行go tool cover -html=c.out，使用cover工具来处理生成的记录信息，该命令会打开本地的浏览器窗口生成一个HTML报告。
// go test -bench=Split -benchmem

//func TestSplit(t *testing.T) {
//	got := Split("我爱你", "爱")
//	want := []string{"我","你"}
//	//got := Split("a:b:c",":")
//	//want := []string{"a","b","c"}
//	if !reflect.DeepEqual(got,want){
//		t.Errorf("want:%v got:%v",want,got)
//	}
//}

func TestSplit(t *testing.T) {
	type test struct {
		input string
		sep string
		want []string
	}
	tests := map[string]test{
		"simple":test{input: "我爱你",sep: "爱", want: []string{"我","你"}},
		"mult sep":test{input: "a:b:c",sep: ":",want: []string{"a","b","c"}},
		"mulr sep2":test{input: "abcd",sep: "bc",want: []string{"a","d"}},

	}
	for name,tc := range tests{
		t.Run(name, func(t *testing.T) {

		got := Split(tc.input,tc.sep)
		if !reflect.DeepEqual(got,tc.want) {
			t.Errorf("name:%v want:%v got:%v", name, tc.want, got)
		}
		})

	}
	}

func BenchmarkSplit(b *testing.B) {
	// b.N 不是固定的数
	for i:=0;i<b.N;i++{
		Split("大家好哇哈哈啊哈","哇")
	}
}