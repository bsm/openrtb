package openrtb

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func BenchmarkBidRequest_Unmarshal(b *testing.B) {
	data, err := os.ReadFile(filepath.Join("testdata", "breq.video.json"))
	if err != nil {
		b.Fatal(err.Error())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var req *BidRequest
		if err := json.Unmarshal(data, &req); err != nil {
			b.Fatal(err.Error())
		}
	}
}

func BenchmarkBidRequest_Marshal(b *testing.B) {
	data, err := os.ReadFile(filepath.Join("testdata", "breq.video.json"))
	if err != nil {
		b.Fatal(err.Error())
	}

	var req *BidRequest
	if err := json.Unmarshal(data, &req); err != nil {
		b.Fatal(err.Error())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(req); err != nil {
			b.Fatal(err.Error())
		}
	}
}
