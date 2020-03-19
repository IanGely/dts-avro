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

  
type Decimal struct {

	
	
		Value string
	

	
	
		Precision int32
	

	
	
		Scale int32
	

}

func NewDecimal() (*Decimal) {
	return &Decimal{}
}

func DeserializeDecimal(r io.Reader) (*Decimal, error) {
	t := NewDecimal()
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

func DeserializeDecimalFromSchema(r io.Reader, schema string) (*Decimal, error) {
	t := NewDecimal()

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

func writeDecimal(r *Decimal, w io.Writer) error {
	var err error
	
	err = vm.WriteString( r.Value, w)
	if err != nil {
		return err			
	}
	
	err = vm.WriteInt( r.Precision, w)
	if err != nil {
		return err			
	}
	
	err = vm.WriteInt( r.Scale, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *Decimal) Serialize(w io.Writer) error {
	return writeDecimal(r, w)
}

func (r *Decimal) Schema() string {
	return "{\"fields\":[{\"name\":\"value\",\"type\":\"string\"},{\"name\":\"precision\",\"type\":\"int\"},{\"name\":\"scale\",\"type\":\"int\"}],\"name\":\"com.alibaba.dts.formats.avro.Decimal\",\"type\":\"record\"}"
}

func (r *Decimal) SchemaName() string {
	return "com.alibaba.dts.formats.avro.Decimal"
}

func (_ *Decimal) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *Decimal) SetInt(v int32) { panic("Unsupported operation") }
func (_ *Decimal) SetLong(v int64) { panic("Unsupported operation") }
func (_ *Decimal) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *Decimal) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *Decimal) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *Decimal) SetString(v string) { panic("Unsupported operation") }
func (_ *Decimal) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Decimal) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
			return (*types.String)(&r.Value)
		
	
	case 1:
		
		
			return (*types.Int)(&r.Precision)
		
	
	case 2:
		
		
			return (*types.Int)(&r.Scale)
		
	
	}
	panic("Unknown field index")
}

func (r *Decimal) SetDefault(i int) {
	switch (i) {
	
        
	
        
	
        
	
	}
	panic("Unknown field index")
}

func (_ *Decimal) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Decimal) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *Decimal) Finalize() { }
