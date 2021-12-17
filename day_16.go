package main

import (
	_ "embed"
	"log"
)

var (
	//go:embed inputs/day_16.txt
	packetsFile string
)

var (
	versionCount int
)

// Version and TypeID positions.
const (
	V0 = 0
	V1 = 3
	T0 = 3
	T1 = 6
)

// Packet Type IDs.
const (
	Sum = iota
	Product
	Minimum
	Maximum
	LiteralValue
	GreaterThan
	LessThan
	EqualTo
)

func packetDecoderVersionCount(bits string) int {
	versionCount = 0
	src := []byte(bits)
	dst := make([]byte, len(src)/2)
	decodeHex(dst, src)
	decodePacket(bytesToBitSlice(dst))
	return versionCount
}

func packetDecoder(bits string) int {
	src := []byte(bits)
	dst := make([]byte, len(src)/2)
	decodeHex(dst, src)
	value, _ := decodePacket(bytesToBitSlice(dst))
	return value
}

func decodePacket(bits []int) (value int, rem []int) {
	version, typeID := bitsToInt(bits[V0:V1]), bitsToInt(bits[T0:T1])
	versionCount += version
	bits = bits[T1:]

	if typeID == LiteralValue {
		return decodeLiteral(bits)
	}

	var values []int
	values, bits = decodeOperator(bits)

	switch typeID {
	case Sum:
		return SumOp(values), bits
	case Product:
		return ProductOp(values), bits
	case Minimum:
		return MinOp(values), bits
	case Maximum:
		return MaxOp(values), bits
	case GreaterThan:
		return GtOp(values), bits
	case LessThan:
		return LtOp(values), bits
	case EqualTo:
		return EqOp(values), bits
	default:
		log.Fatalf("Unknown operator type: %d", typeID)
		return 0, bits
	}
}

func decodeOperator(bits []int) (values, rem []int) {
	var lengthTypeID int
	lengthTypeID, bits = bits[0], bits[1:]
	switch lengthTypeID {
	case 0:
		return decodeOM15(bits)
	case 1:
		return decodeOM11(bits)
	}
	return nil, bits
}

func decodeOM15(bits []int) (values, rem []int) {
	var field []int
	field, bits = bits[:15], bits[15:]
	subPktLen := bitsToInt(field)
	i := 0
	for i < subPktLen {
		prevBitLen := len(bits)
		var v int
		v, bits = decodePacket(bits)
		values = append(values, v)
		i += prevBitLen - len(bits)
	}
	return values, bits
}

func decodeOM11(bits []int) (values, rem []int) {
	var field []int
	field, bits = bits[:11], bits[11:]
	subPktCnt := bitsToInt(field)
	for i := 0; i < subPktCnt; i++ {
		var v int
		v, bits = decodePacket(bits)
		values = append(values, v)
	}
	return values, bits
}

func decodeLiteral(bits []int) (int, []int) {
	num := 0
	lastByte := false
	i := 0
	for ; i < len(bits); i++ {
		if i%5 == 0 {
			if bits[i] == 0 {
				lastByte = true
			}
			continue
		}
		num = (num << 1) + bits[i]
		if lastByte && (i+1)%5 == 0 {
			break
		}
	}
	return num, bits[i+1:]
}

func bytesToBitSlice(b []byte) []int {
	bits := make([]int, 0, len(b)*8)

	for i := 0; i < len(b); i++ {
		for j := 7; j >= 0; j-- {
			mask := byte(1 << j)
			bits = append(bits, int(b[i]&mask/mask))
		}
	}
	return bits
}

func decodeHex(dst, src []byte) {
	i, j := 0, 1
	for ; j < len(src); j += 2 {
		a := fromHexChar(src[j-1])
		b := fromHexChar(src[j])
		dst[i] = (a << 4) | b
		i++
	}
}

func fromHexChar(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func bitsToInt(b []int) int {
	num := 0
	for i := 0; i < len(b); i++ {
		num = (num << 1) + b[i]
	}
	return num
}
