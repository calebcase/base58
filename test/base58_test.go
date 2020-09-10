package base58

import (
	"os"
	"testing"

	"github.com/calebcase/base58"
	"github.com/stretchr/testify/require"
)

type TestCase struct {
	Name string
	Raw  []byte
	Enc  string
}

var TestCases []TestCase

func TestMain(m *testing.M) {
	TestCases = []TestCase{
		TestCase{
			Name: "hello",
			Raw:  []byte("Hello World!"),
			Enc:  "2NEpo7TZRRrLZSi2U",
		},
		TestCase{
			Name: "quick",
			Raw:  []byte("The quick brown fox jumps over the lazy dog."),
			Enc:  "USm3fpXnKG5EUBx2ndxBDMPVciP5hGey2Jh4NDv6gmeo1LkMeiKrLJUUBk6Z",
		},
		TestCase{
			Name: "hex",
			Raw:  []byte{0x00, 0x00, 0x00, 0x28, 0x7f, 0xb4, 0xcd},
			Enc:  "111233QC4",
		},
	}

	os.Exit(m.Run())
}

func TestEncode(t *testing.T) {
	for _, tc := range TestCases {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			require.Equal(t, tc.Enc, base58.Encode(tc.Raw))
		})
	}
}

func TestDecode(t *testing.T) {
	for _, tc := range TestCases {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			raw, err := base58.Decode(tc.Enc)
			require.NoError(t, err)
			require.Equal(t, tc.Raw, raw)
		})
	}
}
