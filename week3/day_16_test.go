package week3

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func getBits(t *testing.T, hex string) []int {
	t.Helper()
	src := []byte(hex)
	dst := make([]byte, len(src)/2)
	decodeHex(dst, src)
	return bytesToBitSlice(dst)
}

func bitsStr(t *testing.T, bits []int) string {
	t.Helper()
	bs := ""
	for _, b := range bits {
		bs += fmt.Sprint(b)
	}
	return bs
}

func TestDecodeLiteral(t *testing.T) {
	bits := getBits(t, "D2FE28")
	wantInt := 2021
	wantRem := []int{0, 0, 0}
	gotInt, gotRem := decodeLiteral(bits[T1:])
	if gotInt != wantInt {
		t.Errorf("decodeLiteral(%q) = (%d, %v), want = (%d, %v)", bitsStr(t, bits), gotInt, gotRem, wantInt, wantRem)
	}
	if diff := cmp.Diff(wantRem, gotRem); diff != "" {
		t.Errorf("decodeLiteral(%q) returned wrong rem bits (-want +got):\n%s", bitsStr(t, bits), diff)
	}
}

func TestDecodeOM15(t *testing.T) {
	bits := getBits(t, "38006F45291200")
	wantVals := []int{10, 20}
	wantRem := []int{0, 0, 0, 0, 0, 0, 0}
	gotVals, gotRem := decodeOM15(bits[T1+1:])
	if diff := cmp.Diff(wantVals, gotVals); diff != "" {
		t.Errorf("decodeOM15(%q) returned wrong values (-want +got):\n%s", bitsStr(t, bits[T1:]), diff)
	}
	if diff := cmp.Diff(wantRem, gotRem); diff != "" {
		t.Errorf("decodeOM15(%q) returned wrong rem bits (-want +got):\n%s", bitsStr(t, bits[T1:]), diff)
	}
}

func TestDecodeOM11(t *testing.T) {
	bits := getBits(t, "EE00D40C823060")
	wantVals := []int{1, 2, 3}
	wantRem := []int{0, 0, 0, 0, 0}
	gotVals, gotRem := decodeOM11(bits[T1+1:])
	if diff := cmp.Diff(wantVals, gotVals); diff != "" {
		t.Errorf("[OM11] decodeOM11(%q) returned wrong values (-want +got):\n%s", bitsStr(t, bits[T1:]), diff)
	}
	if diff := cmp.Diff(wantRem, gotRem); diff != "" {
		t.Errorf("[OM11] decodeOM11(%q) returned wrong rem bits (-want +got):\n%s", bitsStr(t, bits[T1:]), diff)
	}
}

func TestPacketDecoderVersionCount(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"8A004A801A8002F478", 16},
		{"620080001611562C8802118E34", 12},
		{"C0015000016115A2E0802F182340", 23},
		{"A0016C880162017C3686B18A3D4780", 31},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			if got := PacketDecoderVersionCount(tc.input); got != tc.want {
				t.Errorf("packetDecoder(%s) = %d, want = %d", tc.input, got, tc.want)
			}
		})
	}
}

func TestPacketDecoder(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{
			"sum",
			"C200B40A82",
			3,
		},
		{
			"product",
			"04005AC33890",
			54,
		},
		{
			"min",
			"880086C3E88112",
			7,
		},
		{
			"max",
			"CE00C43D881120",
			9,
		},
		{
			"less than",
			"D8005AC2A8F0",
			1,
		},
		{
			"greater than",
			"F600BC2D8F",
			0,
		},
		{
			"equal",
			"9C005AC2F8F0",
			0,
		},
		{
			"sum equal product",
			"9C0141080250320F1802104A08",
			1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := PacketDecoder(tc.input); got != tc.want {
				t.Errorf("packetDecoder(%q) = %d, want = %d", tc.input, got, tc.want)
			}
		})
	}
}
