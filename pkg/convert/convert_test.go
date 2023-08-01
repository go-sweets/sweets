package convert

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	var str StringTo
	str = "100"
	fmt.Println("string: ", str.String())
	fmt.Println(str.Int())
	fmt.Println(str.MustInt())
	fmt.Println(str.UInt32())
	fmt.Println(str.MustUInt32())
	fmt.Println(str.Int64())
	fmt.Println(str.MustInt64())
	str = "100.123"
	fmt.Println(str.Float64())
	fmt.Println(str.MustFloat64())
}
