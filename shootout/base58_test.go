package shootout

import (
	"crypto/rand"
	"os"
	"testing"

	btcutil "github.com/btcsuite/btcutil/base58"
	caleb "github.com/calebcase/base58"
	mrtron "github.com/mr-tron/base58"
)

var encoded string

type EncodeCase struct {
	Name string
	Src  []byte
}

var EncodeCases []EncodeCase

func TestMain(m *testing.M) {
	KB1 := make([]byte, 1024)
	rand.Read(KB1)

	KB10 := make([]byte, 10240)
	rand.Read(KB10)

	KB100 := make([]byte, 102400)
	rand.Read(KB100)

	EncodeCases = []EncodeCase{
		EncodeCase{
			Name: "1K",
			Src:  KB1,
		},
		EncodeCase{
			Name: "10K",
			Src:  KB10,
		},
		EncodeCase{
			Name: "100K",
			Src:  KB100,
		},
	}

	os.Exit(m.Run())
}

func BenchmarkCalebEncode(b *testing.B) {
	for _, tc := range EncodeCases {
		b.Run(tc.Name, func(b *testing.B) {
			var r string
			for n := 0; n < b.N; n++ {
				r = caleb.Encode(tc.Src)
			}

			encoded = r
		})
	}
}

func BenchmarkMrTronEncode(b *testing.B) {
	for _, tc := range EncodeCases {
		b.Run(tc.Name, func(b *testing.B) {
			var r string
			for n := 0; n < b.N; n++ {
				r = mrtron.Encode(tc.Src)
			}

			encoded = r
		})
	}
}

func BenchmarkBtcutilEncode(b *testing.B) {
	for _, tc := range EncodeCases {
		b.Run(tc.Name, func(b *testing.B) {
			var r string
			for n := 0; n < b.N; n++ {
				r = btcutil.Encode(tc.Src)
			}

			encoded = r
		})
	}
}
