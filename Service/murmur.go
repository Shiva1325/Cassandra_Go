package service

import "encoding/binary"

const (
	c1    int64 = -8663945395140668459 // 0x87c37b91114253d5
	c2    int64 = 5545529020109919103  // 0x4cf5ad432745937f
	fmix1 int64 = -49064778989728563   // 0xff51afd7ed558ccd
	fmix2 int64 = -4265267296055464877 // 0xc4ceb9fe1a85ec53
)

func fmix(n int64) int64 {
	n ^= int64(uint64(n) >> 33)
	n *= fmix1
	n ^= int64(uint64(n) >> 33)
	n *= fmix2
	n ^= int64(uint64(n) >> 33)

	return n
}

func block(p byte) int64 {
	return int64(int8(p))
}

func rotl(x int64, r uint8) int64 {
	return (x << r) | (int64)((uint64(x) >> (64 - r)))
}

func getBlock(data []byte, n int) (int64, int64) {
	k1 := int64(binary.LittleEndian.Uint64(data[n*16:]))
	k2 := int64(binary.LittleEndian.Uint64(data[(n*16)+8:]))
	return k1, k2
}
