// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     record.avsc
 */
package avro

import (
	"io"
	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)

  
type Integer struct {

	
	
		Precision int32
	

	
	
		Value string
	

}

func NewInteger() (*Integer) {
	return &Integer{}
}

func DeserializeInteger(r io.Reader) (*Integer, error) {
	t := NewInteger()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err	
	}
	return t, err
}

func DeserializeIntegerFromSchema(r io.Reader, schema string) (*Integer, error) {
	t := NewInteger()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err	
	}
	return t, err
}

func writeInteger(r *Integer, w io.Writer) error {
	var err error
	
	err = vm.WriteInt( r.Precision, w)
	if err != nil {
		return err			
	}
	
	err = vm.WriteString( r.Value, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *Integer) Serialize(w io.Writer) error {
	return writeInteger(r, w)
}

func (r *Integer) Schema() string {
	return "{\"fields\":[{\"name\":\"precision\",\"type\":\"int\"},{\"name\":\"value\",\"type\":\"string\"}],\"name\":\"com.alibaba.dts.formats.avro.Integer\",\"type\":\"record\"}"
}

func (r *Integer) SchemaName() string {
	return "com.alibaba.dts.formats.avro.Integer"
}

func (_ *Integer) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *Integer) SetInt(v int32) { panic("Unsupported operation") }
func (_ *Integer) SetLong(v int64) { panic("Unsupported operation") }
func (_ *Integer) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *Integer) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *Integer) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *Integer) SetString(v string) { panic("Unsupported operation") }
func (_ *Integer) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Integer) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
			return (*types.Int)(&r.Precision)
		
	
	case 1:
		
		
			return (*types.String)(&r.Value)
		
	
	}
	panic("Unknown field index")
}

func (r *Integer) SetDefault(i int) {
	switch (i) {
	
        
	
        
	
	}
	panic("Unknown field index")
}

func (_ *Integer) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Integer) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *Integer) Finalize() { }
