package m

import (
	"bytes"
	"github.com/golang/protobuf/jsonpb"
	"strings"
	"testing"
)

func Benchmark_Json(b *testing.B) {
	msg := new(MyMessage)
	s := "{" + strings.Repeat(`"msg":{"@type":"MyMessage",`, 1000) + `"msg":{"@type":"MyMessage"}` + strings.Repeat("}", 1000) + "}"
	for i := 0; i < b.N; i++ {
		buf := bytes.NewBufferString(s)
		if err := jsonpb.Unmarshal(buf, msg); err != nil {
			panic(err)
		}
	}
}
