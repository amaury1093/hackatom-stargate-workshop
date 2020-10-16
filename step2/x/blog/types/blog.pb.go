// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blog/v1/blog.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgPost struct {
	Id      string                                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty"`
	Title   string                                        `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body    string                                        `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *MsgPost) Reset()         { *m = MsgPost{} }
func (m *MsgPost) String() string { return proto.CompactTextString(m) }
func (*MsgPost) ProtoMessage()    {}
func (*MsgPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cd44f2a4ee163de, []int{0}
}
func (m *MsgPost) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPost.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPost.Merge(m, src)
}
func (m *MsgPost) XXX_Size() int {
	return m.Size()
}
func (m *MsgPost) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPost.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPost proto.InternalMessageInfo

func (m *MsgPost) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgPost) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *MsgPost) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MsgPost) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type MsgComment struct {
	Id      string                                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty"`
	PostID  string                                        `protobuf:"bytes,3,opt,name=postID,proto3" json:"postID,omitempty"`
	Body    string                                        `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *MsgComment) Reset()         { *m = MsgComment{} }
func (m *MsgComment) String() string { return proto.CompactTextString(m) }
func (*MsgComment) ProtoMessage()    {}
func (*MsgComment) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cd44f2a4ee163de, []int{1}
}
func (m *MsgComment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgComment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgComment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgComment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgComment.Merge(m, src)
}
func (m *MsgComment) XXX_Size() int {
	return m.Size()
}
func (m *MsgComment) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgComment.DiscardUnknown(m)
}

var xxx_messageInfo_MsgComment proto.InternalMessageInfo

func (m *MsgComment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgComment) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *MsgComment) GetPostID() string {
	if m != nil {
		return m.PostID
	}
	return ""
}

func (m *MsgComment) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgPost)(nil), "blog.v1.MsgPost")
	proto.RegisterType((*MsgComment)(nil), "blog.v1.MsgComment")
}

func init() { proto.RegisterFile("blog/v1/blog.proto", fileDescriptor_2cd44f2a4ee163de) }

var fileDescriptor_2cd44f2a4ee163de = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xca, 0xc9, 0x4f,
	0xd7, 0x2f, 0x33, 0xd4, 0x07, 0xd1, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x60, 0x76,
	0x99, 0xa1, 0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x58, 0x4c, 0x1f, 0xc4, 0x82, 0x48, 0x2b, 0x4d,
	0x60, 0xe4, 0x62, 0xf7, 0x2d, 0x4e, 0x0f, 0xc8, 0x2f, 0x2e, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c,
	0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0xf2, 0xe6, 0x62, 0x4f, 0x2e,
	0x4a, 0x4d, 0x2c, 0xc9, 0x2f, 0x92, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x71, 0x32, 0xfc, 0x75, 0x4f,
	0x5e, 0x37, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x39, 0xbf, 0x38,
	0x37, 0xbf, 0x18, 0x4a, 0xe9, 0x16, 0xa7, 0x64, 0xeb, 0x97, 0x54, 0x16, 0xa4, 0x16, 0xeb, 0x39,
	0x26, 0x27, 0x3b, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x07, 0xc1, 0x4c, 0x10, 0x12, 0xe1, 0x62,
	0x2d, 0xc9, 0x2c, 0xc9, 0x49, 0x95, 0x60, 0x06, 0x9b, 0x0f, 0xe1, 0x08, 0x09, 0x71, 0xb1, 0x24,
	0xe5, 0xa7, 0x54, 0x4a, 0xb0, 0x80, 0x05, 0xc1, 0x6c, 0xa5, 0xa9, 0x8c, 0x5c, 0x5c, 0xbe, 0xc5,
	0xe9, 0xce, 0xf9, 0xb9, 0xb9, 0xa9, 0x79, 0x34, 0x76, 0x95, 0x18, 0x17, 0x5b, 0x41, 0x7e, 0x71,
	0x89, 0xa7, 0x0b, 0xd4, 0x59, 0x50, 0x1e, 0x36, 0x77, 0x39, 0xb9, 0x9e, 0x78, 0x24, 0xc7, 0x78,
	0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7,
	0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x36, 0x92, 0xed, 0x89, 0xb9, 0x89, 0xa5, 0x45, 0x95, 0xb9,
	0x89, 0x45, 0x25, 0x99, 0x79, 0x95, 0xfa, 0xc5, 0x25, 0xa9, 0x05, 0x46, 0xfa, 0x15, 0xe0, 0x08,
	0x81, 0x38, 0x23, 0x89, 0x0d, 0x1c, 0xf0, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x96, 0xd6,
	0xf2, 0xd2, 0xad, 0x01, 0x00, 0x00,
}

func (m *MsgPost) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPost) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPost) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintBlog(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintBlog(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintBlog(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintBlog(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgComment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgComment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgComment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintBlog(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PostID) > 0 {
		i -= len(m.PostID)
		copy(dAtA[i:], m.PostID)
		i = encodeVarintBlog(dAtA, i, uint64(len(m.PostID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintBlog(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintBlog(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlog(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlog(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgPost) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovBlog(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovBlog(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovBlog(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovBlog(uint64(l))
	}
	return n
}

func (m *MsgComment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovBlog(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovBlog(uint64(l))
	}
	l = len(m.PostID)
	if l > 0 {
		n += 1 + l + sovBlog(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovBlog(uint64(l))
	}
	return n
}

func sovBlog(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlog(x uint64) (n int) {
	return sovBlog(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgPost) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgPost: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPost: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBlog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBlog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlog
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBlog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBlog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBlog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBlog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlog
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBlog
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
func (m *MsgComment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgComment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgComment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBlog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBlog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlog
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBlog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBlog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBlog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBlog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlog
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBlog
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
func skipBlog(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlog
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
					return 0, ErrIntOverflowBlog
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBlog
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
			if length < 0 {
				return 0, ErrInvalidLengthBlog
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlog
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlog
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlog        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlog          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlog = fmt.Errorf("proto: unexpected end of group")
)