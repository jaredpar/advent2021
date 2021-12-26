package day16

import (
	"fmt"
	"strconv"
	"unicode"
)

type Kind int

func trimStart(text string, count int) (start, rest string) {
	runes := []rune(text)
	start = string(runes[0:count])
	rest = string(runes[count:])
	return
}

func hexToBinaryRunes(hex []rune) ([]rune, error) {
	binary := make([]rune, 0)
	for _, r := range hex {
		r = unicode.ToLower(r)
		value, err := strconv.ParseInt(string(r), 16, 32)
		if err != nil {
			return nil, err
		}

		str := strconv.FormatInt(value, 2)
		for len(str) < 4 {
			str = "0" + str
		}

		curRunes := []rune(str)
		binary = append(binary, curRunes...)
	}

	return binary, nil
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
	Payload []rune
}

func NewPacketData(version, typeId int, payload []rune) *PacketData {
	return &PacketData{Version: version, TypeId: typeId, Payload: payload}
}

func parsePacketDataCore(payload []rune) (*PacketData, error) {
	version, err := parseBinaryInt(payload[0:3])
	if err != nil {
		return nil, err
	}

	typeId, err := parseBinaryInt(payload[3:6])
	if err != nil {
		return nil, err
	}

	return NewPacketData(int(version), int(typeId), payload[6:]), nil
}

func ParsePacketData(text string) (*PacketData, error) {
	payload, err := hexToBinaryRunes([]rune(text))
	if err != nil {
		return nil, err
	}
	return parsePacketDataCore(payload)
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
	parts := make([]rune, 0, len(payload))
	for len(payload) >= 5 {
		var isLast = payload[0] == '0'
		section := payload[1:5]
		payload = payload[5:]
		parts = append(parts, section...)
		if isLast {
			break
		}
	}

	value, err := parseBinaryInt(parts)
	if err != nil {
		return nil, nil, err
	}

	return &LiteralPacket{Value: int(value)}, payload, nil
}

func parseOperatorPacket(payload []rune) (*OperatorPacket, []rune, error) {
	if len(payload) < 1 {
		return nil, nil, fmt.Errorf("bad payload")
	}

	var children []Packet
	var remaining []rune
	if payload[0] == '0' {
		payload := payload[1:]
		length, err := parseBinaryInt(payload[:15])
		if err != nil {
			return nil, nil, err
		}

		children = make([]Packet, 0)
		payload = payload[15:]
		consumed := 0
		for {
			child, remaining, err := parsePacketCore(payload)
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
	} else if payload[0] == '1' {
		payload := payload[1:]
		length, err := parseBinaryInt(payload[:11])
		if err != nil {
			return nil, nil, err
		}

		children = make([]Packet, length)
		payload = payload[11:]
		for i := 0; i < length; i++ {
			child, remaining, err := parsePacketCore(payload)
			if err != nil {
				return nil, nil, err
			}

			children[i] = child
			payload = remaining
		}
	} else {
		return nil, nil, fmt.Errorf("bad payload kind: %d", payload[0])
	}

	return &OperatorPacket{Children: children}, remaining, nil
}

func parsePacketCore(payload []rune) (Packet, []rune, error) {
	data, err := parsePacketDataCore(payload)
	if err != nil {
		return nil, nil, err
	}

	payload = data.Payload
	var packet Packet
	if data.TypeId == 4 {
		packet, payload, err = parseLiteralPacket(payload)
	} else {
		packet, payload, err = parseOperatorPacket(payload)
	}

	return packet, payload, err
}

func ParsePacket(text string) (Packet, error) {
	payload, err := hexToBinaryRunes([]rune(text))
	if err != nil {
		return nil, err
	}

	packet, _, err := parsePacketCore(payload)
	return packet, err
}
