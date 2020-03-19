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

  
type TextObject struct {

	
	
		Type string
	

	
	
		Value string
	

}

func NewTextObject() (*TextObject) {
	return &TextObject{}
}

func DeserializeTextObject(r io.Reader) (*TextObject, error) {
	t := NewTextObject()
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

func DeserializeTextObjectFromSchema(r io.Reader, schema string) (*TextObject, error) {
	t := NewTextObject()

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

func writeTextObject(r *TextObject, w io.Writer) error {
	var err error
	
	err = vm.WriteString( r.Type, w)
	if err != nil {
		return err			
	}
	
	err = vm.WriteString( r.Value, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *TextObject) Serialize(w io.Writer) error {
	return writeTextObject(r, w)
}

func (r *TextObject) Schema() string {
	return "{\"fields\":[{\"name\":\"type\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"string\"}],\"name\":\"com.alibaba.dts.formats.avro.TextObject\",\"type\":\"record\"}"
}

func (r *TextObject) SchemaName() string {
	return "com.alibaba.dts.formats.avro.TextObject"
}

func (_ *TextObject) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *TextObject) SetInt(v int32) { panic("Unsupported operation") }
func (_ *TextObject) SetLong(v int64) { panic("Unsupported operation") }
func (_ *TextObject) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *TextObject) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *TextObject) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *TextObject) SetString(v string) { panic("Unsupported operation") }
func (_ *TextObject) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *TextObject) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
			return (*types.String)(&r.Type)
		
	
	case 1:
		
		
			return (*types.String)(&r.Value)
		
	
	}
	panic("Unknown field index")
}

func (r *TextObject) SetDefault(i int) {
	switch (i) {
	
        
	
        
	
	}
	panic("Unknown field index")
}

func (_ *TextObject) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *TextObject) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *TextObject) Finalize() { }