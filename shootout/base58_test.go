package shootout

import (
	"crypto/rand"
	"testing"

	btcutil "github.com/btcsuite/btcutil/base58"
	caleb "github.com/calebcase/base58"
	mrtron "github.com/mr-tron/base58"
)

var decoded []byte
var encoded string

func BenchmarkCalebEncode1K(b *testing.B) {
	src := make([]byte, 1024)
	rand.Read(src)

	var r string
	for n := 0; n < b.N; n++ {
		r = caleb.Encode(src)
	}

	encoded = r
}

func BenchmarkCalebEncode10K(b *testing.B) {
	src := make([]byte, 10240)
	rand.Read(src)

	var r string
	for n := 0; n < b.N; n++ {
		r = caleb.Encode(src)
	}

	encoded = r
}

func BenchmarkCalebEncode100K(b *testing.B) {
	src := make([]byte, 102400)
	rand.Read(src)

	var r string
	for n := 0; n < b.N; n++ {
		r = caleb.Encode(src)
	}

	encoded = r
}

func BenchmarkMrTronEncode1K(b *testing.B) {
	src := make([]byte, 1024)
	rand.Read(src)

	var r string
	for n := 0; n < b.N; n++ {
		r = mrtron.Encode(src)
	}

	encoded = r
}

func BenchmarkMrTronEncode10K(b *testing.B) {
	src := make([]byte, 10240)
	rand.Read(src)

	var r string
	for n := 0; n < b.N; n++ {
		r = mrtron.Encode(src)
	}

	encoded = r
}

func BenchmarkMrTronEncode100K(b *testing.B) {
	src := make([]byte, 102400)
	rand.Read(src)

	var r string
	for n := 0; n < b.N; n++ {
		r = mrtron.Encode(src)
	}

	encoded = r
}

func BenchmarkBtcutilEncode1K(b *testing.B) {
	src := make([]byte, 1024)
	rand.Read(src)

	var r string
	for n := 0; n < b.N; n++ {
		r = btcutil.Encode(src)
	}

	encoded = r
}

func BenchmarkBtcutilEncode10K(b *testing.B) {
	src := make([]byte, 10240)
	rand.Read(src)

	var r string
	for n := 0; n < b.N; n++ {
		r = btcutil.Encode(src)
	}

	encoded = r
}

func BenchmarkBtcutilEncode100K(b *testing.B) {
	src := make([]byte, 102400)
	rand.Read(src)

	var r string
	for n := 0; n < b.N; n++ {
		r = btcutil.Encode(src)
	}

	encoded = r
}
