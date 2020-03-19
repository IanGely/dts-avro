// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     Record.avsc
 */
package avro

import (
	"io"
	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)

  
type Field struct {

	
	
		Name string
	

	
	
		DataTypeNumber int32
	

}

func NewField() (*Field) {
	return &Field{}
}

func DeserializeField(r io.Reader) (*Field, error) {
	t := NewField()
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

func DeserializeFieldFromSchema(r io.Reader, schema string) (*Field, error) {
	t := NewField()

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

func writeField(r *Field, w io.Writer) error {
	var err error
	
	err = vm.WriteString( r.Name, w)
	if err != nil {
		return err			
	}
	
	err = vm.WriteInt( r.DataTypeNumber, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *Field) Serialize(w io.Writer) error {
	return writeField(r, w)
}

func (r *Field) Schema() string {
	return "{\"fields\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataTypeNumber\",\"type\":\"int\"}],\"name\":\"com.alibaba.dts.formats.avro.Field\",\"type\":\"record\"}"
}

func (r *Field) SchemaName() string {
	return "com.alibaba.dts.formats.avro.Field"
}

func (_ *Field) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *Field) SetInt(v int32) { panic("Unsupported operation") }
func (_ *Field) SetLong(v int64) { panic("Unsupported operation") }
func (_ *Field) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *Field) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *Field) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *Field) SetString(v string) { panic("Unsupported operation") }
func (_ *Field) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Field) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
			return (*types.String)(&r.Name)
		
	
	case 1:
		
		
			return (*types.Int)(&r.DataTypeNumber)
		
	
	}
	panic("Unknown field index")
}

func (r *Field) SetDefault(i int) {
	switch (i) {
	
        
	
        
	
	}
	panic("Unknown field index")
}

func (_ *Field) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Field) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *Field) Finalize() { }
