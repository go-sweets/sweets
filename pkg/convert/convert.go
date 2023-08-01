package convert

import "strconv"

type StringTo string

func (s StringTo) String() string {
	return string(s)
}

func (s StringTo) Int() int {
	v, _ := strconv.ParseInt(s.String(), 10, 16)
	return int(v)
}

func (s StringTo) MustInt() int {
	return s.Int()
}

func (s StringTo) UInt32() uint32 {
	v, _ := strconv.ParseUint(s.String(), 10, 32)
	return uint32(v)
}

func (s StringTo) MustUInt32() uint32 {
	return s.UInt32()
}

func (s StringTo) Int64() int64 {
	v, _ := strconv.ParseInt(s.String(), 10, 64)
	return v
}

func (s StringTo) MustInt64() int64 {
	return s.Int64()
}

func (s StringTo) Float64() float64 {
	v, _ := strconv.ParseFloat(s.String(), 64)
	return v
}

func (s StringTo) MustFloat64() float64 {
	return s.Float64()
}
