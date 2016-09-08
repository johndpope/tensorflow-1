// Code generated by protoc-gen-go.
// source: tensorflow/core/framework/function.proto
// DO NOT EDIT!

package config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A library is a set of named functions.
type FunctionDefLibrary struct {
	Function []*FunctionDef `protobuf:"bytes,1,rep,name=function" json:"function,omitempty"`
	Gradient []*GradientDef `protobuf:"bytes,2,rep,name=gradient" json:"gradient,omitempty"`
}

func (m *FunctionDefLibrary) Reset()                    { *m = FunctionDefLibrary{} }
func (m *FunctionDefLibrary) String() string            { return proto.CompactTextString(m) }
func (*FunctionDefLibrary) ProtoMessage()               {}
func (*FunctionDefLibrary) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *FunctionDefLibrary) GetFunction() []*FunctionDef {
	if m != nil {
		return m.Function
	}
	return nil
}

func (m *FunctionDefLibrary) GetGradient() []*GradientDef {
	if m != nil {
		return m.Gradient
	}
	return nil
}

// A function can be instantiated when the runtime can bind every attr
// with a value. When a GraphDef has a call to a function, it must
// have binding for every attr defined in the signature.
//
// TODO(zhifengc):
//   * device spec, etc.
type FunctionDef struct {
	// The definition of the function's name, arguments, return values,
	// attrs etc.
	Signature *OpDef `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
	// The body of the function.
	Node []*FunctionDef_Node `protobuf:"bytes,2,rep,name=node" json:"node,omitempty"`
}

func (m *FunctionDef) Reset()                    { *m = FunctionDef{} }
func (m *FunctionDef) String() string            { return proto.CompactTextString(m) }
func (*FunctionDef) ProtoMessage()               {}
func (*FunctionDef) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *FunctionDef) GetSignature() *OpDef {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *FunctionDef) GetNode() []*FunctionDef_Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// A node is a multi-value assignment:
//   (ret[0], ret[1], ...) = func(arg[0], arg[1], ...)
//
// By convention, "func" is resolved by consulting with a user-defined
// library first. If not resolved, "func" is assumed to be a builtin op.
type FunctionDef_Node struct {
	// This node produces multiple outputs. They are named ret[0],
	// ret[1], ..., etc.
	//
	// REQUIRES: function.node.ret[*] are unique across all nodes.
	// REQUIRES: ret.size == func/op def's number of output args.
	Ret []string `protobuf:"bytes,1,rep,name=ret" json:"ret,omitempty"`
	// The op/function name.
	Op string `protobuf:"bytes,2,opt,name=op" json:"op,omitempty"`
	// Arguments passed to this func/op.
	//
	// arg[i] must be either one of
	// function.signature.input_args[*].name or one of
	// function.node[*].ret[*].
	//
	// REQUIRES: arg.size == func/op def's number of input args.
	Arg []string `protobuf:"bytes,3,rep,name=arg" json:"arg,omitempty"`
	// Control dependencies.
	//
	// dep[i] must be one of function.node[*].ret[*] or one of
	// function.signature.input_args[*].name.
	Dep []string `protobuf:"bytes,4,rep,name=dep" json:"dep,omitempty"`
	// Attrs.
	//
	// 'attr' maps names defined by 'func's attr defs to attr values.
	// attr values may have placeholders which are substituted
	// recursively by concrete values when this node is instantiated.
	// These placeholders must name an attr listed in the FunctionDef's
	// signature.
	Attr map[string]*AttrValue `protobuf:"bytes,5,rep,name=attr" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FunctionDef_Node) Reset()                    { *m = FunctionDef_Node{} }
func (m *FunctionDef_Node) String() string            { return proto.CompactTextString(m) }
func (*FunctionDef_Node) ProtoMessage()               {}
func (*FunctionDef_Node) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1, 0} }

func (m *FunctionDef_Node) GetAttr() map[string]*AttrValue {
	if m != nil {
		return m.Attr
	}
	return nil
}

// GradientDef defines the gradient function of a function defined in
// a function library.
//
// A gradient function g (specified by gradient_func) for a function f
// (specified by function_name) must follow the following:
//
// The function 'f' must be a numerical function which takes N inputs
// and produces M outputs. Its gradient function 'g', which is a
// function taking N + M inputs and produces N outputs.
//
// I.e. if we have
//    (y1, y2, ..., y_M) = f(x1, x2, ..., x_N),
// then, g is
//    (dL/dx1, dL/dx2, ..., dL/dx_N) = g(x1, x2, ..., x_N,
//                                      dL/dy1, dL/dy2, ..., dL/dy_M),
// where L is a scalar-value function of (x1, x2, ..., xN) (e.g., the
// loss function). dL/dx_i is the partial derivative of L with respect
// to x_i.
type GradientDef struct {
	FunctionName string `protobuf:"bytes,1,opt,name=function_name,json=functionName" json:"function_name,omitempty"`
	GradientFunc string `protobuf:"bytes,2,opt,name=gradient_func,json=gradientFunc" json:"gradient_func,omitempty"`
}

func (m *GradientDef) Reset()                    { *m = GradientDef{} }
func (m *GradientDef) String() string            { return proto.CompactTextString(m) }
func (*GradientDef) ProtoMessage()               {}
func (*GradientDef) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func init() {
	proto.RegisterType((*FunctionDefLibrary)(nil), "tensorflow.FunctionDefLibrary")
	proto.RegisterType((*FunctionDef)(nil), "tensorflow.FunctionDef")
	proto.RegisterType((*FunctionDef_Node)(nil), "tensorflow.FunctionDef.Node")
	proto.RegisterType((*GradientDef)(nil), "tensorflow.GradientDef")
}

func init() { proto.RegisterFile("tensorflow/core/framework/function.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0xab, 0xd3, 0x40,
	0x10, 0xc7, 0xd9, 0x34, 0x4f, 0xcc, 0xe4, 0xf9, 0xd0, 0x05, 0x71, 0x09, 0x1e, 0x4a, 0x85, 0x47,
	0x50, 0x48, 0xa4, 0xef, 0x22, 0xbd, 0x59, 0xfc, 0x71, 0x91, 0x5a, 0x72, 0xd0, 0x63, 0xd8, 0x36,
	0x93, 0x10, 0xda, 0xee, 0x86, 0xe9, 0xd6, 0xd2, 0x8b, 0xff, 0xa0, 0xff, 0x80, 0x7f, 0x8a, 0x47,
	0xd9, 0x4d, 0xd2, 0x06, 0xa4, 0xde, 0x92, 0x99, 0xcf, 0xf7, 0x3b, 0x33, 0x3b, 0x03, 0xb1, 0x41,
	0xb5, 0xd7, 0x54, 0x6e, 0xf5, 0x31, 0x5d, 0x6b, 0xc2, 0xb4, 0x24, 0xb9, 0xc3, 0xa3, 0xa6, 0x4d,
	0x5a, 0x1e, 0xd4, 0xda, 0xd4, 0x5a, 0x25, 0x0d, 0x69, 0xa3, 0x39, 0x5c, 0xc8, 0xe8, 0xf5, 0x75,
	0x95, 0x34, 0x86, 0xf2, 0x1f, 0x72, 0x7b, 0xc0, 0x56, 0x17, 0xdd, 0x5f, 0x67, 0x75, 0x93, 0x17,
	0x58, 0xb6, 0xdc, 0xe4, 0x27, 0xf0, 0x4f, 0x5d, 0xc5, 0x0f, 0x58, 0x7e, 0xa9, 0x57, 0x24, 0xe9,
	0xc4, 0x1f, 0xe0, 0x71, 0xdf, 0x87, 0x60, 0xe3, 0x51, 0x1c, 0x4e, 0x5f, 0x24, 0x17, 0xc3, 0x64,
	0xa0, 0xc8, 0xce, 0xa0, 0x15, 0x55, 0x24, 0x8b, 0x1a, 0x95, 0x11, 0xde, 0xbf, 0xa2, 0xcf, 0x5d,
	0xce, 0x89, 0x7a, 0x70, 0xf2, 0xcb, 0x83, 0x70, 0x60, 0xc7, 0x53, 0x08, 0xf6, 0x75, 0xa5, 0xa4,
	0x39, 0x10, 0x0a, 0x36, 0x66, 0x71, 0x38, 0x7d, 0x36, 0x74, 0xf9, 0xda, 0x58, 0xfd, 0x85, 0xe1,
	0x6f, 0xc1, 0x57, 0xba, 0xc0, 0xae, 0xe2, 0xcb, 0x2b, 0x6d, 0x26, 0x0b, 0x5d, 0x60, 0xe6, 0xc8,
	0xe8, 0x37, 0x03, 0xdf, 0xfe, 0xf2, 0xa7, 0x30, 0x22, 0x34, 0x6e, 0xc0, 0x20, 0xb3, 0x9f, 0xfc,
	0x0e, 0x3c, 0xdd, 0x08, 0x6f, 0xcc, 0xe2, 0x20, 0xf3, 0x74, 0x63, 0x09, 0x49, 0x95, 0x18, 0xb5,
	0x84, 0xa4, 0xca, 0x46, 0x0a, 0x6c, 0x84, 0xdf, 0x46, 0x0a, 0x6c, 0xf8, 0x0c, 0x7c, 0xfb, 0xfa,
	0xe2, 0xc6, 0x35, 0x70, 0xff, 0xbf, 0x06, 0x92, 0xf7, 0xc6, 0xd0, 0x47, 0x65, 0xe8, 0x94, 0x39,
	0x4d, 0xb4, 0x80, 0xe0, 0x1c, 0xb2, 0xd6, 0x1b, 0x3c, 0xb9, 0xa1, 0x83, 0xcc, 0x7e, 0xf2, 0x37,
	0x70, 0xe3, 0x76, 0xea, 0x3a, 0x0a, 0xa7, 0xcf, 0x87, 0xde, 0x56, 0xf7, 0xcd, 0x26, 0xb3, 0x96,
	0x99, 0x79, 0xef, 0xd8, 0xe4, 0x3b, 0x84, 0x83, 0x67, 0xe6, 0xaf, 0xe0, 0x49, 0xbf, 0x9d, 0x5c,
	0xc9, 0x1d, 0x76, 0xde, 0xb7, 0x7d, 0x70, 0x21, 0x77, 0x68, 0xa1, 0x7e, 0x1b, 0xb9, 0x4d, 0x74,
	0xe3, 0xdf, 0xf6, 0x41, 0x3b, 0xc4, 0x3c, 0x05, 0xa1, 0xa9, 0x1a, 0xd6, 0x3f, 0xdf, 0xd3, 0xfc,
	0xae, 0x1f, 0x73, 0x69, 0x2f, 0x6a, 0xbf, 0x64, 0x7f, 0x18, 0x5b, 0x3d, 0x72, 0xe7, 0xf5, 0xf0,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x90, 0x85, 0x02, 0xaf, 0xea, 0x02, 0x00, 0x00,
}
