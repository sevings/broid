package broid

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFromString(t *testing.T) {
	req := require.New(t)

	var id BrowserID
	var err error

	id, err = FromString("")
	req.Nil(err)
	req.Empty(id)

	id, err = FromString("f")
	req.NotNil(err)

	id, err = FromString("fac")
	req.NotNil(err)

	id, err = FromString("fafafs")
	req.NotNil(err)

	id, err = FromString("deadbeef12deadbeef12deadbeef12deadbeef12")
	req.Nil(err)

	for i := 0; i < 4; i++ {
		req.Equal(uint8(0xde), id[i*5])
		req.Equal(uint8(0xad), id[i*5+1])
		req.Equal(uint8(0xbe), id[i*5+2])
		req.Equal(uint8(0xef), id[i*5+3])
		req.Equal(uint8(0x12), id[i*5+4])
	}
}

func TestString(t *testing.T) {
	var id BrowserID

	for i := 0; i < 4; i++ {
		id = append(id, 0xde)
		id = append(id, 0xad)
		id = append(id, 0xbe)
		id = append(id, 0xef)
		id = append(id, 0x12)
	}

	require.Equal(t, "deadbeef12deadbeef12deadbeef12deadbeef12", id.String())
}

func TestFromUint64(t *testing.T) {
	req := require.New(t)

	id := FromUint64(0xdeadbeef01234567)

	req.Equal(uint8(0xde), id[0])
	req.Equal(uint8(0xad), id[1])
	req.Equal(uint8(0xbe), id[2])
	req.Equal(uint8(0xef), id[3])
	req.Equal(uint8(0x01), id[4])
	req.Equal(uint8(0x23), id[5])
	req.Equal(uint8(0x45), id[6])
	req.Equal(uint8(0x67), id[7])
}

func TestUint64(t *testing.T) {
	var id BrowserID

	id = append(id, 0xde)
	id = append(id, 0xad)
	id = append(id, 0xbe)
	id = append(id, 0xef)
	id = append(id, 0x01)
	id = append(id, 0x23)
	id = append(id, 0x45)
	id = append(id, 0x67)

	require.Equal(t, uint64(0xdeadbeef01234567), id.Uint64())
}

func TestCompare(t *testing.T) {
	var id1, id2 BrowserID

	id1 = append(id1, 0xde)
	id1 = append(id1, 0xad)
	id1 = append(id1, 0xbe)
	id1 = append(id1, 0xef)

	require.Zero(t, id1.Compare(id1))
	require.Zero(t, id2.Compare(id2))

	require.Equal(t, uint8(4), id1.Compare(id2))
	require.Equal(t, uint8(4), id2.Compare(id1))

	id2 = append(id2, 0xde)
	id2 = append(id2, 0xad)
	id2 = append(id2, 0xde)
	id2 = append(id2, 0xef)
	id2 = append(id2, 0x12)

	require.Equal(t, uint8(2), id1.Compare(id2))
	require.Equal(t, uint8(2), id2.Compare(id1))
}

func BenchmarkConvert(b *testing.B) {
	var n uint64 = 0xdeadbeef

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = FromUint64(n).String()
		n += 255
	}
}
