/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/runtime/generated.proto

/*
	Package runtime is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/runtime/generated.proto

	It has these top-level messages:
		RawExtension
		TypeMeta
		Unknown
*/
package runtime

import "github.com/gogo/protobuf/proto"
import "fmt"
import "math"

import "strings"
import "reflect"

import "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *RawExtension) Reset()                    { *m = RawExtension{} }
func (*RawExtension) ProtoMessage()               {}
func (*RawExtension) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *TypeMeta) Reset()                    { *m = TypeMeta{} }
func (*TypeMeta) ProtoMessage()               {}
func (*TypeMeta) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *Unknown) Reset()                    { *m = Unknown{} }
func (*Unknown) ProtoMessage()               {}
func (*Unknown) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func init() {
	proto.RegisterType((*RawExtension)(nil), "k8s.io.apimachinery.pkg.runtime.RawExtension")
	proto.RegisterType((*TypeMeta)(nil), "k8s.io.apimachinery.pkg.runtime.TypeMeta")
	proto.RegisterType((*Unknown)(nil), "k8s.io.apimachinery.pkg.runtime.Unknown")
}
func (m *RawExtension) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RawExtension) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Raw != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Raw)))
		i += copy(dAtA[i:], m.Raw)
	}
	return i, nil
}

func (m *TypeMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TypeMeta) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.APIVersion)))
	i += copy(dAtA[i:], m.APIVersion)
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Kind)))
	i += copy(dAtA[i:], m.Kind)
	return i, nil
}

func (m *Unknown) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Unknown) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.TypeMeta.Size()))
	n1, err := m.TypeMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.Raw != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Raw)))
		i += copy(dAtA[i:], m.Raw)
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.ContentEncoding)))
	i += copy(dAtA[i:], m.ContentEncoding)
	dAtA[i] = 0x22
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.ContentType)))
	i += copy(dAtA[i:], m.ContentType)
	return i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RawExtension) Size() (n int) {
	var l int
	_ = l
	if m.Raw != nil {
		l = len(m.Raw)
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *TypeMeta) Size() (n int) {
	var l int
	_ = l
	l = len(m.APIVersion)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Kind)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *Unknown) Size() (n int) {
	var l int
	_ = l
	l = m.TypeMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if m.Raw != nil {
		l = len(m.Raw)
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.ContentEncoding)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.ContentType)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *RawExtension) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RawExtension{`,
		`Raw:` + valueToStringGenerated(this.Raw) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TypeMeta) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TypeMeta{`,
		`APIVersion:` + fmt.Sprintf("%v", this.APIVersion) + `,`,
		`Kind:` + fmt.Sprintf("%v", this.Kind) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Unknown) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Unknown{`,
		`TypeMeta:` + strings.Replace(strings.Replace(this.TypeMeta.String(), "TypeMeta", "TypeMeta", 1), `&`, ``, 1) + `,`,
		`Raw:` + valueToStringGenerated(this.Raw) + `,`,
		`ContentEncoding:` + fmt.Sprintf("%v", this.ContentEncoding) + `,`,
		`ContentType:` + fmt.Sprintf("%v", this.ContentType) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *RawExtension) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RawExtension: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RawExtension: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Raw", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Raw = append(m.Raw[:0], dAtA[iNdEx:postIndex]...)
			if m.Raw == nil {
				m.Raw = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TypeMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TypeMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TypeMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field APIVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.APIVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kind = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Unknown) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Unknown: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Unknown: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TypeMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Raw", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Raw = append(m.Raw[:0], dAtA[iNdEx:postIndex]...)
			if m.Raw == nil {
				m.Raw = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentEncoding", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContentEncoding = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContentType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/runtime/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0x4f, 0xab, 0x13, 0x31,
	0x14, 0xc5, 0x27, 0xaf, 0x85, 0x3e, 0xd3, 0xc2, 0x93, 0xb8, 0x70, 0x74, 0x91, 0x79, 0x74, 0xe5,
	0x5b, 0xbc, 0x04, 0x1e, 0x08, 0x6e, 0x3b, 0xa5, 0xa0, 0x88, 0x20, 0xc1, 0x3f, 0xe0, 0xca, 0x74,
	0x26, 0x4e, 0xc3, 0xd0, 0x9b, 0x21, 0xcd, 0x38, 0x76, 0xe7, 0x47, 0xf0, 0x63, 0x75, 0xd9, 0x65,
	0x57, 0xc5, 0x8e, 0x1f, 0xc2, 0xad, 0x34, 0x4d, 0x6b, 0xd5, 0x85, 0xbb, 0xe4, 0x9e, 0xf3, 0x3b,
	0xf7, 0x1e, 0xfc, 0xbc, 0x7c, 0xb6, 0x60, 0xda, 0xf0, 0xb2, 0x9e, 0x2a, 0x0b, 0xca, 0xa9, 0x05,
	0xff, 0xac, 0x20, 0x37, 0x96, 0x07, 0x41, 0x56, 0x7a, 0x2e, 0xb3, 0x99, 0x06, 0x65, 0x97, 0xbc,
	0x2a, 0x0b, 0x6e, 0x6b, 0x70, 0x7a, 0xae, 0x78, 0xa1, 0x40, 0x59, 0xe9, 0x54, 0xce, 0x2a, 0x6b,
	0x9c, 0x21, 0xc9, 0x01, 0x60, 0xe7, 0x00, 0xab, 0xca, 0x82, 0x05, 0xe0, 0xf1, 0x6d, 0xa1, 0xdd,
	0xac, 0x9e, 0xb2, 0xcc, 0xcc, 0x79, 0x61, 0x0a, 0xc3, 0x3d, 0x37, 0xad, 0x3f, 0xf9, 0x9f, 0xff,
	0xf8, 0xd7, 0x21, 0x6f, 0x78, 0x83, 0x07, 0x42, 0x36, 0x93, 0x2f, 0x4e, 0xc1, 0x42, 0x1b, 0x20,
	0x8f, 0x70, 0xc7, 0xca, 0x26, 0x46, 0xd7, 0xe8, 0xc9, 0x20, 0xed, 0xb5, 0xdb, 0xa4, 0x23, 0x64,
	0x23, 0xf6, 0xb3, 0xe1, 0x47, 0x7c, 0xf9, 0x66, 0x59, 0xa9, 0x57, 0xca, 0x49, 0x72, 0x87, 0xb1,
	0xac, 0xf4, 0x3b, 0x65, 0xf7, 0x90, 0x77, 0xdf, 0x4b, 0xc9, 0x6a, 0x9b, 0x44, 0xed, 0x36, 0xc1,
	0xa3, 0xd7, 0x2f, 0x82, 0x22, 0xce, 0x5c, 0xe4, 0x1a, 0x77, 0x4b, 0x0d, 0x79, 0x7c, 0xe1, 0xdd,
	0x83, 0xe0, 0xee, 0xbe, 0xd4, 0x90, 0x0b, 0xaf, 0x0c, 0x7f, 0x22, 0xdc, 0x7b, 0x0b, 0x25, 0x98,
	0x06, 0xc8, 0x7b, 0x7c, 0xe9, 0xc2, 0x36, 0x9f, 0xdf, 0xbf, 0xbb, 0x61, 0xff, 0xe9, 0xce, 0x8e,
	0xe7, 0xa5, 0xf7, 0x43, 0xf8, 0xe9, 0x60, 0x71, 0x0a, 0x3b, 0x36, 0xbc, 0xf8, 0xb7, 0x21, 0x19,
	0xe1, 0xab, 0xcc, 0x80, 0x53, 0xe0, 0x26, 0x90, 0x99, 0x5c, 0x43, 0x11, 0x77, 0xfc, 0xb1, 0x0f,
	0x43, 0xde, 0xd5, 0xf8, 0x4f, 0x59, 0xfc, 0xed, 0x27, 0x4f, 0x71, 0x3f, 0x8c, 0xf6, 0xab, 0xe3,
	0xae, 0xc7, 0x1f, 0x04, 0xbc, 0x3f, 0xfe, 0x2d, 0x89, 0x73, 0x5f, 0x7a, 0xbb, 0xda, 0xd1, 0x68,
	0xbd, 0xa3, 0xd1, 0x66, 0x47, 0xa3, 0xaf, 0x2d, 0x45, 0xab, 0x96, 0xa2, 0x75, 0x4b, 0xd1, 0xa6,
	0xa5, 0xe8, 0x7b, 0x4b, 0xd1, 0xb7, 0x1f, 0x34, 0xfa, 0xd0, 0x0b, 0x45, 0x7f, 0x05, 0x00, 0x00,
	0xff, 0xff, 0xe3, 0x33, 0x18, 0x0b, 0x50, 0x02, 0x00, 0x00,
}
