// Code generated by protoc-gen-go.
// source: object.proto
// DO NOT EDIT!

/*
Package spreadproto is a generated protocol buffer package.

It is generated from these files:
	object.proto

It has these top-level messages:
	Field
	Object
	Array
	SRI
	Link
	Document
	DocumentInfo
*/
package spreadproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// Field represents a field of an object.
type Field struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Field_Number
	//	*Field_Str
	//	*Field_Boolean
	//	*Field_Object
	//	*Field_Array
	//	*Field_Link
	Value isField_Value `protobuf_oneof:"value"`
}

func (m *Field) Reset()                    { *m = Field{} }
func (m *Field) String() string            { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()               {}
func (*Field) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isField_Value interface {
	isField_Value()
}

type Field_Number struct {
	Number float64 `protobuf:"fixed64,2,opt,name=number,oneof"`
}
type Field_Str struct {
	Str string `protobuf:"bytes,3,opt,name=str,oneof"`
}
type Field_Boolean struct {
	Boolean bool `protobuf:"varint,4,opt,name=boolean,oneof"`
}
type Field_Object struct {
	Object *Object `protobuf:"bytes,5,opt,name=object,oneof"`
}
type Field_Array struct {
	Array *Array `protobuf:"bytes,6,opt,name=array,oneof"`
}
type Field_Link struct {
	Link *Link `protobuf:"bytes,7,opt,name=link,oneof"`
}

func (*Field_Number) isField_Value()  {}
func (*Field_Str) isField_Value()     {}
func (*Field_Boolean) isField_Value() {}
func (*Field_Object) isField_Value()  {}
func (*Field_Array) isField_Value()   {}
func (*Field_Link) isField_Value()    {}

func (m *Field) GetValue() isField_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Field) GetNumber() float64 {
	if x, ok := m.GetValue().(*Field_Number); ok {
		return x.Number
	}
	return 0
}

func (m *Field) GetStr() string {
	if x, ok := m.GetValue().(*Field_Str); ok {
		return x.Str
	}
	return ""
}

func (m *Field) GetBoolean() bool {
	if x, ok := m.GetValue().(*Field_Boolean); ok {
		return x.Boolean
	}
	return false
}

func (m *Field) GetObject() *Object {
	if x, ok := m.GetValue().(*Field_Object); ok {
		return x.Object
	}
	return nil
}

func (m *Field) GetArray() *Array {
	if x, ok := m.GetValue().(*Field_Array); ok {
		return x.Array
	}
	return nil
}

func (m *Field) GetLink() *Link {
	if x, ok := m.GetValue().(*Field_Link); ok {
		return x.Link
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Field) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Field_OneofMarshaler, _Field_OneofUnmarshaler, _Field_OneofSizer, []interface{}{
		(*Field_Number)(nil),
		(*Field_Str)(nil),
		(*Field_Boolean)(nil),
		(*Field_Object)(nil),
		(*Field_Array)(nil),
		(*Field_Link)(nil),
	}
}

func _Field_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Field)
	// value
	switch x := m.Value.(type) {
	case *Field_Number:
		b.EncodeVarint(2<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.Number))
	case *Field_Str:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Str)
	case *Field_Boolean:
		t := uint64(0)
		if x.Boolean {
			t = 1
		}
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Field_Object:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Object); err != nil {
			return err
		}
	case *Field_Array:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Array); err != nil {
			return err
		}
	case *Field_Link:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Link); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Field.Value has unexpected type %T", x)
	}
	return nil
}

func _Field_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Field)
	switch tag {
	case 2: // value.number
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &Field_Number{math.Float64frombits(x)}
		return true, err
	case 3: // value.str
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &Field_Str{x}
		return true, err
	case 4: // value.boolean
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Field_Boolean{x != 0}
		return true, err
	case 5: // value.object
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Object)
		err := b.DecodeMessage(msg)
		m.Value = &Field_Object{msg}
		return true, err
	case 6: // value.array
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Array)
		err := b.DecodeMessage(msg)
		m.Value = &Field_Array{msg}
		return true, err
	case 7: // value.link
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Link)
		err := b.DecodeMessage(msg)
		m.Value = &Field_Link{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Field_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Field)
	// value
	switch x := m.Value.(type) {
	case *Field_Number:
		n += proto.SizeVarint(2<<3 | proto.WireFixed64)
		n += 8
	case *Field_Str:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Str)))
		n += len(x.Str)
	case *Field_Boolean:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += 1
	case *Field_Object:
		s := proto.Size(x.Object)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Field_Array:
		s := proto.Size(x.Array)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Field_Link:
		s := proto.Size(x.Link)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Object represents a map with strings for keys.
type Object struct {
	Items map[string]*Field `protobuf:"bytes,1,rep,name=items" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Object) Reset()                    { *m = Object{} }
func (m *Object) String() string            { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()               {}
func (*Object) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Object) GetItems() map[string]*Field {
	if m != nil {
		return m.Items
	}
	return nil
}

// Array represents an array.
type Array struct {
	Items []*Field `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *Array) Reset()                    { *m = Array{} }
func (m *Array) String() string            { return proto.CompactTextString(m) }
func (*Array) ProtoMessage()               {}
func (*Array) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Array) GetItems() []*Field {
	if m != nil {
		return m.Items
	}
	return nil
}

// A SRI represents a parsed Spread Resource Identifier (SRI), a globally unique address for an object or field stored within a repository.
type SRI struct {
	Treeish string `protobuf:"bytes,1,opt,name=treeish" json:"treeish,omitempty"`
	Path    string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Field   string `protobuf:"bytes,3,opt,name=field" json:"field,omitempty"`
}

func (m *SRI) Reset()                    { *m = SRI{} }
func (m *SRI) String() string            { return proto.CompactTextString(m) }
func (*SRI) ProtoMessage()               {}
func (*SRI) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Link represents a relationship to another field.
type Link struct {
	PackageName string `protobuf:"bytes,1,opt,name=packageName" json:"packageName,omitempty"`
	Target      *SRI   `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
	Override    bool   `protobuf:"varint,3,opt,name=override" json:"override,omitempty"`
}

func (m *Link) Reset()                    { *m = Link{} }
func (m *Link) String() string            { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()               {}
func (*Link) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Link) GetTarget() *SRI {
	if m != nil {
		return m.Target
	}
	return nil
}

// Document is the root of Spread data stored in a Git blob. It has field stored at it's root, typically with an object as it's value.
type Document struct {
	Name string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Info *DocumentInfo `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
	Root *Field        `protobuf:"bytes,3,opt,name=root" json:"root,omitempty"`
}

func (m *Document) Reset()                    { *m = Document{} }
func (m *Document) String() string            { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()               {}
func (*Document) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Document) GetInfo() *DocumentInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Document) GetRoot() *Field {
	if m != nil {
		return m.Root
	}
	return nil
}

// DocumentInfo provides metadata about an document.
type DocumentInfo struct {
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
}

func (m *DocumentInfo) Reset()                    { *m = DocumentInfo{} }
func (m *DocumentInfo) String() string            { return proto.CompactTextString(m) }
func (*DocumentInfo) ProtoMessage()               {}
func (*DocumentInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*Field)(nil), "spread.Field")
	proto.RegisterType((*Object)(nil), "spread.Object")
	proto.RegisterType((*Array)(nil), "spread.Array")
	proto.RegisterType((*SRI)(nil), "spread.SRI")
	proto.RegisterType((*Link)(nil), "spread.Link")
	proto.RegisterType((*Document)(nil), "spread.Document")
	proto.RegisterType((*DocumentInfo)(nil), "spread.DocumentInfo")
}

var fileDescriptor0 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x52, 0xcf, 0x4f, 0xe2, 0x40,
	0x18, 0xa5, 0xf4, 0x27, 0x5f, 0xcb, 0x2e, 0xdb, 0xdd, 0x43, 0x77, 0x97, 0x18, 0xd2, 0xc4, 0x84,
	0x53, 0x0f, 0x72, 0x31, 0x7a, 0x92, 0xa8, 0x01, 0x63, 0x34, 0xd1, 0x93, 0xde, 0xa6, 0xf0, 0x01,
	0x95, 0xb6, 0xd3, 0x4c, 0x07, 0x12, 0xfe, 0x2c, 0xff, 0x43, 0x67, 0xa6, 0x6d, 0xa0, 0xf1, 0xd4,
	0x7c, 0xef, 0xbd, 0x79, 0xef, 0x7b, 0x5f, 0x0a, 0x1e, 0x8d, 0x3f, 0x70, 0xc1, 0xa3, 0x82, 0x51,
	0x4e, 0x7d, 0xab, 0x2c, 0x18, 0x92, 0x65, 0xf8, 0xa9, 0x81, 0x79, 0x9f, 0x60, 0xba, 0xf4, 0x5d,
	0xd0, 0xb7, 0x78, 0x08, 0xb4, 0x91, 0x36, 0xee, 0xf9, 0x03, 0xb0, 0xf2, 0x5d, 0x16, 0x23, 0x0b,
	0xba, 0x62, 0xd6, 0x66, 0x1d, 0xbf, 0x0f, 0x7a, 0xc9, 0x59, 0xa0, 0x4b, 0x5a, 0x8c, 0xbf, 0xc0,
	0x8e, 0x29, 0x4d, 0x91, 0xe4, 0x81, 0x21, 0x20, 0x47, 0x40, 0x23, 0xb0, 0xaa, 0x88, 0xc0, 0x14,
	0x88, 0x7b, 0xf1, 0x23, 0xaa, 0x32, 0xa2, 0x67, 0x85, 0x0a, 0xc5, 0x19, 0x98, 0x84, 0x31, 0x72,
	0x08, 0x2c, 0x25, 0xe8, 0x37, 0x82, 0x1b, 0x09, 0x0a, 0x7e, 0x08, 0x46, 0x9a, 0xe4, 0xdb, 0xc0,
	0x56, 0xb4, 0xd7, 0xd0, 0x8f, 0x02, 0x9b, 0x75, 0xa6, 0x36, 0x98, 0x7b, 0x92, 0xee, 0x30, 0xa4,
	0x60, 0x55, 0x96, 0xfe, 0x18, 0xcc, 0x84, 0x63, 0x56, 0x8a, 0xad, 0x75, 0xf1, 0xe2, 0x6f, 0x3b,
	0x31, 0x9a, 0x4b, 0xee, 0x2e, 0xe7, 0xec, 0xf0, 0xef, 0x1a, 0xe0, 0x38, 0xb5, 0xbb, 0x0e, 0x6b,
	0x5f, 0x55, 0xf5, 0x64, 0x2b, 0x75, 0x96, 0xab, 0xee, 0xa5, 0x16, 0x9e, 0x83, 0xa9, 0x56, 0x94,
	0xd2, 0xd3, 0xbc, 0xb6, 0x34, 0x9c, 0x80, 0xfe, 0xfa, 0x32, 0xf7, 0x7f, 0x82, 0xcd, 0x19, 0x62,
	0x52, 0x6e, 0xea, 0x00, 0x0f, 0x8c, 0x82, 0xf0, 0x8d, 0xf2, 0xef, 0x89, 0x43, 0x9a, 0x2b, 0x29,
	0xaf, 0x4e, 0x19, 0x3e, 0x80, 0x21, 0xfb, 0xf9, 0xbf, 0xc1, 0x2d, 0xc8, 0x62, 0x4b, 0xd6, 0xf8,
	0x44, 0x32, 0xac, 0x5f, 0xfe, 0x07, 0x8b, 0x13, 0xb6, 0x46, 0x5e, 0xef, 0xe6, 0x36, 0x81, 0x32,
	0x67, 0x00, 0x0e, 0xdd, 0x23, 0x63, 0xc9, 0x12, 0x95, 0x97, 0x13, 0xbe, 0x81, 0x73, 0x4b, 0x17,
	0xbb, 0x0c, 0x73, 0x2e, 0x43, 0xf3, 0xa3, 0x51, 0x08, 0x46, 0x92, 0xaf, 0x68, 0x6d, 0xf3, 0xa7,
	0xb1, 0x69, 0xd4, 0x73, 0xc1, 0x89, 0x30, 0x83, 0x51, 0xca, 0x95, 0xd7, 0xb7, 0x6e, 0x43, 0xf0,
	0x5a, 0xe2, 0xa6, 0x93, 0xb2, 0x9f, 0xf6, 0xdf, 0xdd, 0x4a, 0xad, 0x7e, 0xae, 0xd8, 0x52, 0x9f,
	0xc9, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x01, 0x5c, 0x94, 0x66, 0x73, 0x02, 0x00, 0x00,
}
