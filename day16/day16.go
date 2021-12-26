package day16

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Kind int

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

func parseBinaryInt(runes []rune) (int, error) {
	value := 0
	for i, r := range runes {
		if r == '1' {
			shift := len(runes) - (i + 1)
			cur := 1 << shift
			value += cur
		} else if r != '0' {
			return 0, fmt.Errorf("not a valid binary digit: %s", string(r))
		}
	}

	return value, nil
}

type PacketData struct {
	Version int
	TypeId  int
	Payload string
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

const (
	KindLiteral Kind = iota
	KindOperator
)

type Packet interface {
	Kind() Kind
}

type LiteralPacket struct {
	Value int
}

func (p *LiteralPacket) Kind() Kind {
	return KindLiteral
}

type OperatorPacket struct {
	Children []Packet
}

func (p *OperatorPacket) Kind() Kind {
	return KindOperator
}

func parseLiteralPacket(payload []rune) (*LiteralPacket, []rune, error) {
	var sb strings.Builder
	for len(payload) >= 5 {
		var isLast = payload[0] == '0'
		section := payload[1:5]
		payload = payload[5:]
		for _, r := range section {
			sb.WriteRune(r)
		}

		if isLast {
			break
		}
	}

	value, err := strconv.ParseInt(sb.String(), 2, 32)
	if err != nil {
		return nil, nil, err
	}

	return &LiteralPacket{Value: int(value)}, payload, nil
}

func parseOperatorPacket(payload []rune) (*OperatorPacket, []rune, error) {
	if len(payload) < 16 {
		return nil, nil, fmt.Errorf("bad payload")
	}

	length, err := parseBinaryInt(payload[:15])
	if err != nil {
		return nil, nil, err
	}

	children := make([]Packet, 0)
	payload = payload[15:]
	consumed := 0
	for {
		child, remaining, err := parseLiteralPacket(payload)
		if err != nil {
			return nil, nil, err
		}

		children = append(children, child)
		consumed += len(payload) - len(remaining)
		payload = remaining
		if consumed >= length {
			break
		}
	}

	return &OperatorPacket{Children: children}, payload, nil
}

func ParsePacket(text string) (Packet, error) {
	data, err := ParsePacketData(text)
	if err != nil {
		return nil, err
	}

	payload := []rune(data.Payload)
	var packet Packet
	switch data.TypeId {
	case 4:
		packet, _, err = parseLiteralPacket(payload)
	case 6:
		packet, _, err = parseOperatorPacket(payload)
	default:
		err = fmt.Errorf("bad typeid: %d", data.TypeId)
	}

	return packet, err
}
