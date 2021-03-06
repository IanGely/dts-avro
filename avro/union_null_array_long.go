// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     Record.avsc
 */
package avro

import (
	"io"
	"fmt"

	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
)


type UnionNullArrayLongTypeEnum int
const (

	 UnionNullArrayLongTypeEnumNull UnionNullArrayLongTypeEnum = 0

	 UnionNullArrayLongTypeEnumArrayLong UnionNullArrayLongTypeEnum = 1

)

type UnionNullArrayLong struct {

	Null *types.NullVal

	ArrayLong []int64

	UnionType UnionNullArrayLongTypeEnum
}

func writeUnionNullArrayLong(r *UnionNullArrayLong, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case UnionNullArrayLongTypeEnumNull:
		return vm.WriteNull(r.Null, w)
        
	case UnionNullArrayLongTypeEnumArrayLong:
		return writeArrayLong(r.ArrayLong, w)
        
	}
	return fmt.Errorf("invalid value for *UnionNullArrayLong")
}

func NewUnionNullArrayLong() *UnionNullArrayLong {
	return &UnionNullArrayLong{}
}

func (_ *UnionNullArrayLong) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *UnionNullArrayLong) SetInt(v int32) { panic("Unsupported operation") }
func (_ *UnionNullArrayLong) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *UnionNullArrayLong) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullArrayLong) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *UnionNullArrayLong) SetString(v string) { panic("Unsupported operation") }
func (r *UnionNullArrayLong) SetLong(v int64) { 
	r.UnionType = (UnionNullArrayLongTypeEnum)(v)
}
func (r *UnionNullArrayLong) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return r.Null
		
	
	case 1:
		
		r.ArrayLong = make([]int64, 0)
		
		
		return (*ArrayLongWrapper)(&r.ArrayLong)
		
	
	}
	panic("Unknown field index")
}
func (_ *UnionNullArrayLong) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullArrayLong) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullArrayLong) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullArrayLong) Finalize()  { }
