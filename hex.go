package marshaler

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"

	"github.com/go-passwd/hasher"
)

// HexMarshaler stores password in HEX
type HexMarshaler struct {
	Separator string
}

// Marshal hasher.Hasher to string
func (m *HexMarshaler) Marshal(h hasher.Hasher) (string, error) {
	var params templateParams
	switch h.Code() {
	case hasher.TypePlain:
		hh := h.(*hasher.PlainHasher)
		params = templateParams{
			Code:       h.Code(),
			Iterations: 0,
			Salt:       "",
			Password:   hex.EncodeToString(*hh.Password),
		}
	case hasher.TypeMD5:
		hh := h.(*hasher.MD5Hasher)
		params = templateParams{
			Code:       h.Code(),
			Iterations: *hh.Iter,
			Salt:       *hh.Salt,
			Password:   hex.EncodeToString(*hh.Password),
		}
	case hasher.TypeSHA1:
		hh := h.(*hasher.SHA1Hasher)
		params = templateParams{
			Code:       h.Code(),
			Iterations: *hh.Iter,
			Salt:       *hh.Salt,
			Password:   hex.EncodeToString(*hh.Password),
		}
	case hasher.TypeSHA224:
		hh := h.(*hasher.SHA224Hasher)
		params = templateParams{
			Code:       h.Code(),
			Iterations: *hh.Iter,
			Salt:       *hh.Salt,
			Password:   hex.EncodeToString(*hh.Password),
		}
	case hasher.TypeSHA256:
		hh := h.(*hasher.SHA256Hasher)
		params = templateParams{
			Code:       h.Code(),
			Iterations: *hh.Iter,
			Salt:       *hh.Salt,
			Password:   hex.EncodeToString(*hh.Password),
		}
	case hasher.TypeSHA384:
		hh := h.(*hasher.SHA384Hasher)
		params = templateParams{
			Code:       h.Code(),
			Iterations: *hh.Iter,
			Salt:       *hh.Salt,
			Password:   hex.EncodeToString(*hh.Password),
		}
	case hasher.TypeSHA512:
		hh := h.(*hasher.SHA512Hasher)
		params = templateParams{
			Code:       h.Code(),
			Iterations: *hh.Iter,
			Salt:       *hh.Salt,
			Password:   hex.EncodeToString(*hh.Password),
		}
	case hasher.TypeSHA512_224:
		hh := h.(*hasher.SHA512_224Hasher)
		params = templateParams{
			Code:       h.Code(),
			Iterations: *hh.Iter,
			Salt:       *hh.Salt,
			Password:   hex.EncodeToString(*hh.Password),
		}
	case hasher.TypeSHA512_256:
		hh := h.(*hasher.SHA512_256Hasher)
		params = templateParams{
			Code:       h.Code(),
			Iterations: *hh.Iter,
			Salt:       *hh.Salt,
			Password:   hex.EncodeToString(*hh.Password),
		}
	}
	params.Separator = m.Separator
	buf := bytes.NewBufferString("")
	err := marshalTemplate.ExecuteTemplate(buf, "marshalTemplate", params)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// Unmarshal string to Hasher
func (m *HexMarshaler) Unmarshal(s string) (hasher.Hasher, error) {
	buf := bytes.NewBufferString("")
	params := templateParams{Separator: m.Separator}
	err := unmarshalPattern.ExecuteTemplate(buf, "unmarshalPattern", params)
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(buf.String())
	submatch := re.FindStringSubmatch(s)
	if submatch == nil {
		return nil, fmt.Errorf("cannot unmarshal string %s", s)
	}

	password, err := hex.DecodeString(submatch[4])
	if err != nil {
		return nil, err
	}
	iter, err := strconv.Atoi(submatch[2])
	if err != nil {
		return nil, err
	}

	switch submatch[1] {
	case hasher.TypePlain:
		return &hasher.PlainHasher{
			Password: &password,
		}, nil
	case hasher.TypeMD5:
		return &hasher.MD5Hasher{
			Iter:     &iter,
			Salt:     &submatch[3],
			Password: &password,
		}, nil
	case hasher.TypeSHA1:
		return &hasher.SHA1Hasher{
			Iter:     &iter,
			Salt:     &submatch[3],
			Password: &password,
		}, nil
	case hasher.TypeSHA224:
		return &hasher.SHA224Hasher{
			Iter:     &iter,
			Salt:     &submatch[3],
			Password: &password,
		}, nil
	case hasher.TypeSHA256:
		return &hasher.SHA256Hasher{
			Iter:     &iter,
			Salt:     &submatch[3],
			Password: &password,
		}, nil
	case hasher.TypeSHA384:
		return &hasher.SHA384Hasher{
			Iter:     &iter,
			Salt:     &submatch[3],
			Password: &password,
		}, nil
	case hasher.TypeSHA512:
		return &hasher.SHA512Hasher{
			Iter:     &iter,
			Salt:     &submatch[3],
			Password: &password,
		}, nil
	case hasher.TypeSHA512_224:
		return &hasher.SHA512_224Hasher{
			Iter:     &iter,
			Salt:     &submatch[3],
			Password: &password,
		}, nil
	case hasher.TypeSHA512_256:
		return &hasher.SHA512_256Hasher{
			Iter:     &iter,
			Salt:     &submatch[3],
			Password: &password,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported hasher %s", submatch[1])
	}
}