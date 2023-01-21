package broid

import "strconv"

// BrowserID contains browser characteristics.
type BrowserID []uint8

// FromString returns a BrowserID parsed from the string.
func FromString(idStr string) (BrowserID, error) {
	if len(idStr)%2 != 0 {
		return nil, strconv.ErrSyntax
	}

	id := make([]uint8, 0, len(idStr)/2)

	for i := 0; i < len(idStr)/2; i++ {
		f := idStr[i*2 : i*2+2]
		v, err := strconv.ParseUint(f, 16, 8)
		if err != nil {
			return nil, err
		}

		id = append(id, uint8(v))
	}

	return id, nil
}

// String stringifies the BrowserID.
func (id BrowserID) String() string {
	const digits = "0123456789abcdef"

	str := make([]byte, 0, len(id)*2)

	for _, f := range id {
		v1 := digits[f>>4]
		str = append(str, v1)
		v2 := digits[f&0xf]
		str = append(str, v2)
	}

	return string(str)
}

// FromUint64 returns a BrowserID parsed from an uint64.
func FromUint64(n uint64) BrowserID {
	id := make([]uint8, 0, 8)

	for n > 0 {
		id = append(id, uint8(n&0xff))
		n >>= 8
	}

	for i, j := 0, len(id)-1; i < j; i, j = i+1, j-1 {
		id[i], id[j] = id[j], id[i]
	}

	return id
}

// Uint64 represents the BrowserID as an uint64.
func (id BrowserID) Uint64() uint64 {
	var n uint64

	for _, f := range id {
		n = (n << 8) + uint64(f)
	}

	return n
}

// Compare returns the number of distinct field values between two given BrowserIDs.
func (id BrowserID) Compare(other BrowserID) uint8 {
	var diff uint8

	minLen := len(id)
	if minLen > len(other) {
		minLen = len(other)
		diff = uint8(len(id) - len(other))
	} else if minLen < len(other) {
		diff = uint8(len(other) - len(id))
	}

	for i := 0; i < minLen; i++ {
		if id[i] != other[i] {
			diff++
		}
	}

	return diff
}
