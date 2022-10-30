package storage

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"testing"

	"github.com/ungtb10d/gdu/v5/pkg/storage/itempb"
	"google.golang.org/protobuf/proto"
)

func BenchmarkSerializeToPb(b *testing.B) {
	dirpb := &itempb.Dir{
		Item: &itempb.Item{
			Path: "/xxx/zzz",
			Size: 1e3,
		},
	}

	proto.Marshal(dirpb)
}

func BenchmarkSerializeToGob(b *testing.B) {
	dirpb := &itempb.Dir{
		Item: &itempb.Item{
			Path: "/xxx/zzz",
			Size: 1e3,
		},
	}

	encoder := gob.NewEncoder(bytes.NewBufferString(""))

	b.ResetTimer()

	encoder.Encode(dirpb)
}

func BenchmarkSerializeToJson(b *testing.B) {
	dirpb := &itempb.Dir{
		Item: &itempb.Item{
			Path: "/xxx/zzz",
			Size: 1e3,
		},
	}

	encoder := json.NewEncoder(bytes.NewBufferString(""))

	b.ResetTimer()

	encoder.Encode(dirpb)
}
