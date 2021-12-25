package day16

import (
	"strconv"
	"strings"
	"unicode"
)

func trimStart(text string, count int) (start, rest string) {
	runes := []rune(text)
	start = string(runes[0:count])
	rest = string(runes[count:])
	return
}

func toBinaryString(r rune) (string, error) {
	r = unicode.ToLower(r)
	value, err := strconv.ParseInt(string(r), 16, 32)
	if err != nil {
		return "", err
	}

	str := strconv.FormatInt(value, 2)
	for len(str) < 4 {
		str = "0" + str
	}

	return str, nil
}

type Packet struct {
	Version int
	TypeId  int
	Payload string
}

func NewPacket(version, typeId int, payload string) *Packet {
	return &Packet{Version: version, TypeId: typeId, Payload: payload}
}

func ParsePacket(text string) (*Packet, error) {
	var sb strings.Builder
	for _, r := range text {
		str, err := toBinaryString(r)
		if err != nil {
			return nil, err
		}

		sb.WriteString(str)
	}

	binary := sb.String()
	versionStr, binary := trimStart(binary, 3)
	version, err := strconv.ParseInt(versionStr, 2, 32)
	if err != nil {
		return nil, err
	}

	typeIdStr, binary := trimStart(binary, 3)
	typeId, err := strconv.ParseInt(typeIdStr, 2, 32)
	if err != nil {
		return nil, err
	}

	return NewPacket(int(version), int(typeId), binary), nil
}
