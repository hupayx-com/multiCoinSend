// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: multicoinsend/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgCoinSend struct {
	Creator   string     `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Receivers []Receiver `protobuf:"bytes,2,rep,name=receivers,proto3" json:"receivers"`
}

func (m *MsgCoinSend) Reset()         { *m = MsgCoinSend{} }
func (m *MsgCoinSend) String() string { return proto.CompactTextString(m) }
func (*MsgCoinSend) ProtoMessage()    {}
func (*MsgCoinSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3f878001e3eb274, []int{0}
}
func (m *MsgCoinSend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCoinSend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCoinSend.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCoinSend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCoinSend.Merge(m, src)
}
func (m *MsgCoinSend) XXX_Size() int {
	return m.Size()
}
func (m *MsgCoinSend) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCoinSend.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCoinSend proto.InternalMessageInfo

func (m *MsgCoinSend) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCoinSend) GetReceivers() []Receiver {
	if m != nil {
		return m.Receivers
	}
	return nil
}

type Receiver struct {
	To     string                                   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Amount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *Receiver) Reset()         { *m = Receiver{} }
func (m *Receiver) String() string { return proto.CompactTextString(m) }
func (*Receiver) ProtoMessage()    {}
func (*Receiver) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3f878001e3eb274, []int{1}
}
func (m *Receiver) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Receiver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Receiver.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Receiver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Receiver.Merge(m, src)
}
func (m *Receiver) XXX_Size() int {
	return m.Size()
}
func (m *Receiver) XXX_DiscardUnknown() {
	xxx_messageInfo_Receiver.DiscardUnknown(m)
}

var xxx_messageInfo_Receiver proto.InternalMessageInfo

func (m *Receiver) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Receiver) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

type MsgCoinSendResponse struct {
	CntValue    string `protobuf:"bytes,1,opt,name=cntValue,proto3" json:"cntValue,omitempty"`
	TotalAmount string `protobuf:"bytes,2,opt,name=totalAmount,proto3" json:"totalAmount,omitempty"`
}

func (m *MsgCoinSendResponse) Reset()         { *m = MsgCoinSendResponse{} }
func (m *MsgCoinSendResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCoinSendResponse) ProtoMessage()    {}
func (*MsgCoinSendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3f878001e3eb274, []int{2}
}
func (m *MsgCoinSendResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCoinSendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCoinSendResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCoinSendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCoinSendResponse.Merge(m, src)
}
func (m *MsgCoinSendResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCoinSendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCoinSendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCoinSendResponse proto.InternalMessageInfo

func (m *MsgCoinSendResponse) GetCntValue() string {
	if m != nil {
		return m.CntValue
	}
	return ""
}

func (m *MsgCoinSendResponse) GetTotalAmount() string {
	if m != nil {
		return m.TotalAmount
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgCoinSend)(nil), "hupayxcom.multicoinsend.multicoinsend.MsgCoinSend")
	proto.RegisterType((*Receiver)(nil), "hupayxcom.multicoinsend.multicoinsend.Receiver")
	proto.RegisterType((*MsgCoinSendResponse)(nil), "hupayxcom.multicoinsend.multicoinsend.MsgCoinSendResponse")
}

func init() { proto.RegisterFile("multicoinsend/tx.proto", fileDescriptor_e3f878001e3eb274) }

var fileDescriptor_e3f878001e3eb274 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x4e, 0xf2, 0x40,
	0x14, 0x6d, 0xe1, 0x0b, 0x1f, 0x0c, 0x89, 0x8b, 0xd1, 0x18, 0xec, 0xa2, 0x10, 0x12, 0x13, 0x36,
	0xcc, 0x08, 0xae, 0x74, 0x27, 0xae, 0xd9, 0xb4, 0x89, 0x0b, 0x77, 0xed, 0x30, 0x29, 0x8d, 0xb4,
	0xb7, 0xe9, 0x4c, 0x09, 0x18, 0x13, 0xe3, 0x1b, 0xf8, 0x1c, 0x3e, 0x09, 0x4b, 0x96, 0xae, 0xd4,
	0xc0, 0x8b, 0x98, 0xe9, 0x0f, 0x16, 0x57, 0xb8, 0xea, 0xbd, 0xbd, 0x73, 0xce, 0x3d, 0xe7, 0xe4,
	0xa2, 0xd3, 0x20, 0x99, 0x49, 0x9f, 0x81, 0x1f, 0x0a, 0x1e, 0x4e, 0xa8, 0x5c, 0x90, 0x28, 0x06,
	0x09, 0xf8, 0x7c, 0x9a, 0x44, 0xce, 0x72, 0xc1, 0x20, 0x20, 0x7b, 0x2f, 0xf6, 0x3b, 0xe3, 0xc4,
	0x03, 0x0f, 0x52, 0x04, 0x55, 0x55, 0x06, 0x36, 0x4c, 0x06, 0x22, 0x00, 0x41, 0x5d, 0x47, 0x70,
	0x3a, 0x1f, 0xb8, 0x5c, 0x3a, 0x03, 0xaa, 0x30, 0xd9, 0xbc, 0xfb, 0x84, 0x9a, 0x63, 0xe1, 0xdd,
	0x82, 0x1f, 0xda, 0x3c, 0x9c, 0xe0, 0x16, 0xfa, 0xcf, 0x62, 0xee, 0x48, 0x88, 0x5b, 0x7a, 0x47,
	0xef, 0x35, 0xac, 0xa2, 0xc5, 0x36, 0x6a, 0xc4, 0x9c, 0x71, 0x7f, 0xce, 0x63, 0xd1, 0xaa, 0x74,
	0xaa, 0xbd, 0xe6, 0x90, 0x92, 0x83, 0x94, 0x11, 0x2b, 0xc7, 0x8d, 0xfe, 0xad, 0x3e, 0xda, 0x9a,
	0xf5, 0xc3, 0xd3, 0x7d, 0x46, 0xf5, 0x62, 0x88, 0x8f, 0x50, 0x45, 0x42, 0xbe, 0xb5, 0x22, 0x01,
	0x33, 0x54, 0x73, 0x02, 0x48, 0x42, 0x99, 0x6f, 0x3b, 0x23, 0x99, 0x15, 0xa2, 0xac, 0x90, 0xdc,
	0x0a, 0x51, 0xca, 0x47, 0x17, 0x8a, 0xf7, 0xed, 0xb3, 0xdd, 0xf3, 0x7c, 0x39, 0x4d, 0x5c, 0xc2,
	0x20, 0xa0, 0xb9, 0xef, 0xec, 0xd3, 0x17, 0x93, 0x07, 0x2a, 0x97, 0x11, 0x17, 0x29, 0x40, 0x58,
	0x39, 0x75, 0xd7, 0x46, 0xc7, 0x25, 0xfb, 0x16, 0x17, 0x11, 0x84, 0x82, 0x63, 0x03, 0xd5, 0x59,
	0x28, 0xef, 0x9c, 0x59, 0xc2, 0x73, 0x45, 0xbb, 0x1e, 0x77, 0x50, 0x53, 0x82, 0x74, 0x66, 0x37,
	0x85, 0x38, 0x35, 0x2e, 0xff, 0x1a, 0xbe, 0xe8, 0xa8, 0x3a, 0x16, 0x1e, 0x7e, 0x44, 0xf5, 0x5d,
	0xb0, 0xc3, 0x03, 0xb3, 0x2a, 0xa9, 0x31, 0xae, 0xff, 0x8e, 0x29, 0x1c, 0x8c, 0xec, 0xd5, 0xc6,
	0xd4, 0xd7, 0x1b, 0x53, 0xff, 0xda, 0x98, 0xfa, 0xeb, 0xd6, 0xd4, 0xd6, 0x5b, 0x53, 0x7b, 0xdf,
	0x9a, 0xda, 0xfd, 0x55, 0x29, 0xa4, 0x8c, 0xbf, 0xaf, 0xca, 0x94, 0xb2, 0xa0, 0xa1, 0x0b, 0xfa,
	0xeb, 0x18, 0x55, 0x76, 0x6e, 0x2d, 0xbd, 0x99, 0xcb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc7,
	0xdf, 0xf6, 0x17, 0xaa, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CoinSend(ctx context.Context, in *MsgCoinSend, opts ...grpc.CallOption) (*MsgCoinSendResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CoinSend(ctx context.Context, in *MsgCoinSend, opts ...grpc.CallOption) (*MsgCoinSendResponse, error) {
	out := new(MsgCoinSendResponse)
	err := c.cc.Invoke(ctx, "/hupayxcom.multicoinsend.multicoinsend.Msg/CoinSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CoinSend(context.Context, *MsgCoinSend) (*MsgCoinSendResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CoinSend(ctx context.Context, req *MsgCoinSend) (*MsgCoinSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoinSend not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CoinSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCoinSend)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CoinSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hupayxcom.multicoinsend.multicoinsend.Msg/CoinSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CoinSend(ctx, req.(*MsgCoinSend))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hupayxcom.multicoinsend.multicoinsend.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CoinSend",
			Handler:    _Msg_CoinSend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "multicoinsend/tx.proto",
}

func (m *MsgCoinSend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCoinSend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCoinSend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Receivers) > 0 {
		for iNdEx := len(m.Receivers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Receivers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Receiver) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Receiver) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Receiver) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintTx(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCoinSendResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCoinSendResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCoinSendResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TotalAmount) > 0 {
		i -= len(m.TotalAmount)
		copy(dAtA[i:], m.TotalAmount)
		i = encodeVarintTx(dAtA, i, uint64(len(m.TotalAmount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CntValue) > 0 {
		i -= len(m.CntValue)
		copy(dAtA[i:], m.CntValue)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CntValue)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCoinSend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Receivers) > 0 {
		for _, e := range m.Receivers {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *Receiver) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgCoinSendResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CntValue)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.TotalAmount)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCoinSend) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCoinSend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCoinSend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receivers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receivers = append(m.Receivers, Receiver{})
			if err := m.Receivers[len(m.Receivers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *Receiver) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: Receiver: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Receiver: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCoinSendResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCoinSendResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCoinSendResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CntValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CntValue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
