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

type Packet interface {
	Version() int
	Children() []Packet
}

type LiteralPacket struct {
	version int
	Value   int
}

func (p *LiteralPacket) Version() int {
	return p.version
}

func (p *LiteralPacket) Children() []Packet {
	return make([]Packet, 0)
}

type OperatorPacket struct {
	version  int
	children []Packet
}

func (p *OperatorPacket) Version() int {
	return p.version
}

func (p *OperatorPacket) Children() []Packet {
	return p.children
}

func parseLiteralPacket(version int, payload []rune) (*LiteralPacket, []rune, error) {
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

	return &LiteralPacket{version: version, Value: int(value)}, payload, nil
}

func parseOperatorPacket(version int, payload []rune) (*OperatorPacket, []rune, error) {
	if len(payload) < 1 {
		return nil, nil, fmt.Errorf("bad payload")
	}

	if payload[0] == '0' {
		payload := payload[1:]
		length, err := parseBinaryInt(payload[:15])
		if err != nil {
			return nil, nil, err
		}

		children := make([]Packet, 0)
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
				return &OperatorPacket{children: children, version: version}, remaining, nil
			}
		}
	} else if payload[0] == '1' {
		payload := payload[1:]
		length, err := parseBinaryInt(payload[:11])
		if err != nil {
			return nil, nil, err
		}

		children := make([]Packet, length)
		payload = payload[11:]
		for i := 0; i < length; i++ {
			child, remaining, err := parsePacketCore(payload)
			if err != nil {
				return nil, nil, err
			}

			children[i] = child
			payload = remaining
		}
		return &OperatorPacket{children: children, version: version}, payload, nil
	} else {
		return nil, nil, fmt.Errorf("bad payload kind: %d", payload[0])
	}
}

func parsePacketCore(payload []rune) (Packet, []rune, error) {
	data, err := parsePacketDataCore(payload)
	if err != nil {
		return nil, nil, err
	}

	payload = data.Payload
	var packet Packet
	if data.TypeId == 4 {
		packet, payload, err = parseLiteralPacket(data.Version, payload)
	} else {
		packet, payload, err = parseOperatorPacket(data.Version, payload)
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

func Part1(text string) (int, error) {
	packet, err := ParsePacket(text)
	if err != nil {
		return 0, err
	}

	sum := 0
	toVisit := []Packet{packet}
	for len(toVisit) > 0 {
		cur := toVisit[0]
		sum += cur.Version()
		toVisit = append(toVisit[1:], cur.Children()...)
	}

	return sum, nil
}
