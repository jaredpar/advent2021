package day16

import (
	"fmt"
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

type PacketData struct {
	Version int
	TypeId  int
	Payload string
}

type Packet interface {
	TypeId() int
}

type LiteralPacket struct {
	typeId int
	Value  int
}

func (p *LiteralPacket) TypeId() int {
	return p.typeId
}

func NewPacketData(version, typeId int, payload string) *PacketData {
	return &PacketData{Version: version, TypeId: typeId, Payload: payload}
}

func ParsePacketData(text string) (*PacketData, error) {
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

	return NewPacketData(int(version), int(typeId), binary), nil
}

func ParsePacketLiteral(data *PacketData) (*LiteralPacket, error) {
	runes := []rune(data.Payload)
	var sb strings.Builder

	for len(runes) >= 5 {
		var isLast = runes[0] == '0'
		section := runes[1:5]
		runes = runes[5:]
		for _, r := range section {
			sb.WriteRune(r)
		}

		if isLast {
			break
		}
	}

	value, err := strconv.ParseInt(sb.String(), 2, 32)
	if err != nil {
		return nil, err
	}

	return &LiteralPacket{typeId: data.TypeId, Value: int(value)}, nil
}

func ParsePacket(text string) (Packet, error) {
	data, err := ParsePacketData(text)
	if err != nil {
		return nil, err
	}

	switch data.TypeId {
	case 4:
		return ParsePacketLiteral(data)
	default:
		return nil, fmt.Errorf("bad typeid: %d", data.TypeId)
	}
}
