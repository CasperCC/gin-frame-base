package sqids_tool

import (
	"gin-frame-base/internal/constant"
	"github.com/sqids/sqids-go"
	"strings"
)

func id2string(id uint64, minLength uint8, alphabet string) (idString string) {
	opt := sqids.Options{
		MinLength: minLength,
		Alphabet:  alphabet,
	}
	s, _ := sqids.New(opt)
	idString, _ = s.Encode([]uint64{id})
	return
}

func string2id(idString string, minLength uint8, alphabet string) (id uint64) {
	opt := sqids.Options{
		MinLength: minLength,
		Alphabet:  alphabet,
	}
	idString = strings.TrimPrefix(idString, constant.SqidsPrefixUser)
	s, _ := sqids.New(opt)
	ids := s.Decode(idString)
	id = ids[0]
	return
}

func UserIdEncode(id uint64) string {
	return constant.SqidsPrefixUser + id2string(id, constant.SqidsMinLengthUser, constant.SqidsShufflingUser)
}

func UserIdDecode(idString string) uint64 {
	idString = strings.TrimPrefix(idString, constant.SqidsPrefixUser)
	return string2id(idString, constant.SqidsMinLengthUser, constant.SqidsShufflingUser)
}

func fileIdEncode(id uint64) string {
	return constant.SqidsPrefixFile + id2string(id, constant.SqidsMinLengthFile, constant.SqidsShufflingFile)
}

func fileIdDecode(idString string) uint64 {
	idString = strings.TrimPrefix(idString, constant.SqidsPrefixFile)
	return string2id(idString, constant.SqidsMinLengthFile, constant.SqidsShufflingFile)
}
