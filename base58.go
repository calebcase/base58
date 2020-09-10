package base58

import (
	"bytes"
	"errors"
	"math/big"
	"strings"
)

var ErrInvalidSymbol = errors.New("invalid symbol")

func to58(b byte) byte {
	switch b {
	case '0':
		return '1'
	case '1':
		return '2'
	case '2':
		return '3'
	case '3':
		return '4'
	case '4':
		return '5'
	case '5':
		return '6'
	case '6':
		return '7'
	case '7':
		return '8'
	case '8':
		return '9'
	case '9':
		return 'A'
	case 'a':
		return 'B'
	case 'b':
		return 'C'
	case 'c':
		return 'D'
	case 'd':
		return 'E'
	case 'e':
		return 'F'
	case 'f':
		return 'G'
	case 'g':
		return 'H'
	case 'h':
		return 'J'
	case 'i':
		return 'K'
	case 'j':
		return 'L'
	case 'k':
		return 'M'
	case 'l':
		return 'N'
	case 'm':
		return 'P'
	case 'n':
		return 'Q'
	case 'o':
		return 'R'
	case 'p':
		return 'S'
	case 'q':
		return 'T'
	case 'r':
		return 'U'
	case 's':
		return 'V'
	case 't':
		return 'W'
	case 'u':
		return 'X'
	case 'v':
		return 'Y'
	case 'w':
		return 'Z'
	case 'x':
		return 'a'
	case 'y':
		return 'b'
	case 'z':
		return 'c'
	case 'A':
		return 'd'
	case 'B':
		return 'e'
	case 'C':
		return 'f'
	case 'D':
		return 'g'
	case 'E':
		return 'h'
	case 'F':
		return 'i'
	case 'G':
		return 'j'
	case 'H':
		return 'k'
	case 'I':
		return 'm'
	case 'J':
		return 'n'
	case 'K':
		return 'o'
	case 'L':
		return 'p'
	case 'M':
		return 'q'
	case 'N':
		return 'r'
	case 'O':
		return 's'
	case 'P':
		return 't'
	case 'Q':
		return 'u'
	case 'R':
		return 'v'
	case 'S':
		return 'w'
	case 'T':
		return 'x'
	case 'U':
		return 'y'
	case 'V':
		return 'z'
	default:
		return '!'
	}
}

func from58(b byte) (byte, error) {
	switch b {
	case '1':
		return '0', nil
	case '2':
		return '1', nil
	case '3':
		return '2', nil
	case '4':
		return '3', nil
	case '5':
		return '4', nil
	case '6':
		return '5', nil
	case '7':
		return '6', nil
	case '8':
		return '7', nil
	case '9':
		return '8', nil
	case 'A':
		return '9', nil
	case 'B':
		return 'a', nil
	case 'C':
		return 'b', nil
	case 'D':
		return 'c', nil
	case 'E':
		return 'd', nil
	case 'F':
		return 'e', nil
	case 'G':
		return 'f', nil
	case 'H':
		return 'g', nil
	case 'J':
		return 'h', nil
	case 'K':
		return 'i', nil
	case 'L':
		return 'j', nil
	case 'M':
		return 'k', nil
	case 'N':
		return 'l', nil
	case 'P':
		return 'm', nil
	case 'Q':
		return 'n', nil
	case 'R':
		return 'o', nil
	case 'S':
		return 'p', nil
	case 'T':
		return 'q', nil
	case 'U':
		return 'r', nil
	case 'V':
		return 's', nil
	case 'W':
		return 't', nil
	case 'X':
		return 'u', nil
	case 'Y':
		return 'v', nil
	case 'Z':
		return 'w', nil
	case 'a':
		return 'x', nil
	case 'b':
		return 'y', nil
	case 'c':
		return 'z', nil
	case 'd':
		return 'A', nil
	case 'e':
		return 'B', nil
	case 'f':
		return 'C', nil
	case 'g':
		return 'D', nil
	case 'h':
		return 'E', nil
	case 'i':
		return 'F', nil
	case 'j':
		return 'G', nil
	case 'k':
		return 'H', nil
	case 'm':
		return 'I', nil
	case 'n':
		return 'J', nil
	case 'o':
		return 'K', nil
	case 'p':
		return 'L', nil
	case 'q':
		return 'M', nil
	case 'r':
		return 'N', nil
	case 's':
		return 'O', nil
	case 't':
		return 'P', nil
	case 'u':
		return 'Q', nil
	case 'v':
		return 'R', nil
	case 'w':
		return 'S', nil
	case 'x':
		return 'T', nil
	case 'y':
		return 'U', nil
	case 'z':
		return 'V', nil
	default:
		return '!', ErrInvalidSymbol
	}
}

func Encode(src []byte) string {
	var dst strings.Builder

	for _, b := range src {
		if b == 0 {
			dst.WriteByte('1')
		} else {
			break
		}
	}

	i := new(big.Int).SetBytes(src)

	for _, b := range []byte(i.Text(58)) {
		dst.WriteByte(to58(b))
	}

	return dst.String()
}

func Decode(src string) (_ []byte, err error) {
	var dst bytes.Buffer

	tmp := []byte(src)

	leading := true
	for i, b := range tmp {
		if leading {
			if b == '1' {
				dst.WriteByte(0)
			} else {
				leading = false
			}
		}

		tmp[i], err = from58(b)
		if err != nil {
			return nil, err
		}
	}

	i, ok := new(big.Int).SetString(string(tmp), 58)
	if !ok {
		return nil, errors.New("cannot decode input")
	}

	dst.Write(i.Bytes())

	return dst.Bytes(), nil
}
