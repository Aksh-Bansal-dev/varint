package main

const mask = int64((1 << 7) - 1) 

func EncodeInt64(n int64) []byte {
	res := []byte{}
	for {
		res = append(res, byte(n&mask))
		n >>= 7
		if n > 0 {
			res[len(res)-1] |= 1 << 7
		} else {
			break
		}
	}
	return res
}

func DecodeInt64(arr []byte) int64 {
	res := int64(0)
	for i, e := range arr {
		res |= ((int64(e) & mask) << (7 * i))
	}
	return res
}
