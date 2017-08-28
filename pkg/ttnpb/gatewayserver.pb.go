// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/TheThingsNetwork/ttn/api/gatewayserver.proto

package ttnpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import google_protobuf4 "github.com/gogo/protobuf/types"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// GatewayUp may contain zero or more uplink messages and/or a status message for the gateway.
type GatewayUp struct {
	UplinkMessage []*UplinkMessage `protobuf:"bytes,1,rep,name=uplink_message,json=uplinkMessage" json:"uplink_message,omitempty"`
	GatewayStatus *GatewayStatus   `protobuf:"bytes,2,opt,name=gateway_status,json=gatewayStatus" json:"gateway_status,omitempty"`
}

func (m *GatewayUp) Reset()                    { *m = GatewayUp{} }
func (*GatewayUp) ProtoMessage()               {}
func (*GatewayUp) Descriptor() ([]byte, []int) { return fileDescriptorGatewayserver, []int{0} }

func (m *GatewayUp) GetUplinkMessage() []*UplinkMessage {
	if m != nil {
		return m.UplinkMessage
	}
	return nil
}

func (m *GatewayUp) GetGatewayStatus() *GatewayStatus {
	if m != nil {
		return m.GatewayStatus
	}
	return nil
}

// GatewayDown contains downlink messages for the gateway
type GatewayDown struct {
	DownlinkMessage *DownlinkMessage `protobuf:"bytes,1,opt,name=downlink_message,json=downlinkMessage" json:"downlink_message,omitempty"`
}

func (m *GatewayDown) Reset()                    { *m = GatewayDown{} }
func (*GatewayDown) ProtoMessage()               {}
func (*GatewayDown) Descriptor() ([]byte, []int) { return fileDescriptorGatewayserver, []int{1} }

func (m *GatewayDown) GetDownlinkMessage() *DownlinkMessage {
	if m != nil {
		return m.DownlinkMessage
	}
	return nil
}

func init() {
	proto.RegisterType((*GatewayUp)(nil), "ttn.v3.GatewayUp")
	proto.RegisterType((*GatewayDown)(nil), "ttn.v3.GatewayDown")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GtwGs service

type GtwGsClient interface {
	// Link the gateway to the gateway server. The authentication information will
	// be used to determine the gateway ID. If no authentication information is present,
	// this gateway may not be used for downlink.
	Link(ctx context.Context, opts ...grpc.CallOption) (GtwGs_LinkClient, error)
}

type gtwGsClient struct {
	cc *grpc.ClientConn
}

func NewGtwGsClient(cc *grpc.ClientConn) GtwGsClient {
	return &gtwGsClient{cc}
}

func (c *gtwGsClient) Link(ctx context.Context, opts ...grpc.CallOption) (GtwGs_LinkClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GtwGs_serviceDesc.Streams[0], c.cc, "/ttn.v3.GtwGs/Link", opts...)
	if err != nil {
		return nil, err
	}
	x := &gtwGsLinkClient{stream}
	return x, nil
}

type GtwGs_LinkClient interface {
	Send(*GatewayUp) error
	Recv() (*GatewayDown, error)
	grpc.ClientStream
}

type gtwGsLinkClient struct {
	grpc.ClientStream
}

func (x *gtwGsLinkClient) Send(m *GatewayUp) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gtwGsLinkClient) Recv() (*GatewayDown, error) {
	m := new(GatewayDown)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GtwGs service

type GtwGsServer interface {
	// Link the gateway to the gateway server. The authentication information will
	// be used to determine the gateway ID. If no authentication information is present,
	// this gateway may not be used for downlink.
	Link(GtwGs_LinkServer) error
}

func RegisterGtwGsServer(s *grpc.Server, srv GtwGsServer) {
	s.RegisterService(&_GtwGs_serviceDesc, srv)
}

func _GtwGs_Link_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GtwGsServer).Link(&gtwGsLinkServer{stream})
}

type GtwGs_LinkServer interface {
	Send(*GatewayDown) error
	Recv() (*GatewayUp, error)
	grpc.ServerStream
}

type gtwGsLinkServer struct {
	grpc.ServerStream
}

func (x *gtwGsLinkServer) Send(m *GatewayDown) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gtwGsLinkServer) Recv() (*GatewayUp, error) {
	m := new(GatewayUp)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GtwGs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.v3.GtwGs",
	HandlerType: (*GtwGsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Link",
			Handler:       _GtwGs_Link_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/TheThingsNetwork/ttn/api/gatewayserver.proto",
}

// Client API for NsGs service

type NsGsClient interface {
	// ScheduleDownlink instructs the gateway server to schedule a downlink message.
	// The gateway server may refuse if there are any conflicts in the schedule or
	// if a duty cycle prevents the gateway from transmitting.
	ScheduleDownlink(ctx context.Context, in *DownlinkMessage, opts ...grpc.CallOption) (*google_protobuf4.Empty, error)
}

type nsGsClient struct {
	cc *grpc.ClientConn
}

func NewNsGsClient(cc *grpc.ClientConn) NsGsClient {
	return &nsGsClient{cc}
}

func (c *nsGsClient) ScheduleDownlink(ctx context.Context, in *DownlinkMessage, opts ...grpc.CallOption) (*google_protobuf4.Empty, error) {
	out := new(google_protobuf4.Empty)
	err := grpc.Invoke(ctx, "/ttn.v3.NsGs/ScheduleDownlink", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NsGs service

type NsGsServer interface {
	// ScheduleDownlink instructs the gateway server to schedule a downlink message.
	// The gateway server may refuse if there are any conflicts in the schedule or
	// if a duty cycle prevents the gateway from transmitting.
	ScheduleDownlink(context.Context, *DownlinkMessage) (*google_protobuf4.Empty, error)
}

func RegisterNsGsServer(s *grpc.Server, srv NsGsServer) {
	s.RegisterService(&_NsGs_serviceDesc, srv)
}

func _NsGs_ScheduleDownlink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownlinkMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsGsServer).ScheduleDownlink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.v3.NsGs/ScheduleDownlink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsGsServer).ScheduleDownlink(ctx, req.(*DownlinkMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _NsGs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.v3.NsGs",
	HandlerType: (*NsGsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScheduleDownlink",
			Handler:    _NsGs_ScheduleDownlink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/TheThingsNetwork/ttn/api/gatewayserver.proto",
}

// Client API for Gs service

type GsClient interface {
	GetGatewayObservations(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*GatewayObservations, error)
}

type gsClient struct {
	cc *grpc.ClientConn
}

func NewGsClient(cc *grpc.ClientConn) GsClient {
	return &gsClient{cc}
}

func (c *gsClient) GetGatewayObservations(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*GatewayObservations, error) {
	out := new(GatewayObservations)
	err := grpc.Invoke(ctx, "/ttn.v3.Gs/GetGatewayObservations", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Gs service

type GsServer interface {
	GetGatewayObservations(context.Context, *GatewayIdentifiers) (*GatewayObservations, error)
}

func RegisterGsServer(s *grpc.Server, srv GsServer) {
	s.RegisterService(&_Gs_serviceDesc, srv)
}

func _Gs_GetGatewayObservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GsServer).GetGatewayObservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.v3.Gs/GetGatewayObservations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GsServer).GetGatewayObservations(ctx, req.(*GatewayIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.v3.Gs",
	HandlerType: (*GsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGatewayObservations",
			Handler:    _Gs_GetGatewayObservations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/TheThingsNetwork/ttn/api/gatewayserver.proto",
}

func (m *GatewayUp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GatewayUp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UplinkMessage) > 0 {
		for _, msg := range m.UplinkMessage {
			dAtA[i] = 0xa
			i++
			i = encodeVarintGatewayserver(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.GatewayStatus != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGatewayserver(dAtA, i, uint64(m.GatewayStatus.Size()))
		n1, err := m.GatewayStatus.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *GatewayDown) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GatewayDown) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DownlinkMessage != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGatewayserver(dAtA, i, uint64(m.DownlinkMessage.Size()))
		n2, err := m.DownlinkMessage.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeFixed64Gatewayserver(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Gatewayserver(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGatewayserver(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GatewayUp) Size() (n int) {
	var l int
	_ = l
	if len(m.UplinkMessage) > 0 {
		for _, e := range m.UplinkMessage {
			l = e.Size()
			n += 1 + l + sovGatewayserver(uint64(l))
		}
	}
	if m.GatewayStatus != nil {
		l = m.GatewayStatus.Size()
		n += 1 + l + sovGatewayserver(uint64(l))
	}
	return n
}

func (m *GatewayDown) Size() (n int) {
	var l int
	_ = l
	if m.DownlinkMessage != nil {
		l = m.DownlinkMessage.Size()
		n += 1 + l + sovGatewayserver(uint64(l))
	}
	return n
}

func sovGatewayserver(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGatewayserver(x uint64) (n int) {
	return sovGatewayserver(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GatewayUp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GatewayUp{`,
		`UplinkMessage:` + strings.Replace(fmt.Sprintf("%v", this.UplinkMessage), "UplinkMessage", "UplinkMessage", 1) + `,`,
		`GatewayStatus:` + strings.Replace(fmt.Sprintf("%v", this.GatewayStatus), "GatewayStatus", "GatewayStatus", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GatewayDown) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GatewayDown{`,
		`DownlinkMessage:` + strings.Replace(fmt.Sprintf("%v", this.DownlinkMessage), "DownlinkMessage", "DownlinkMessage", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGatewayserver(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GatewayUp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGatewayserver
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
			return fmt.Errorf("proto: GatewayUp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GatewayUp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UplinkMessage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatewayserver
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
				return ErrInvalidLengthGatewayserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UplinkMessage = append(m.UplinkMessage, &UplinkMessage{})
			if err := m.UplinkMessage[len(m.UplinkMessage)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GatewayStatus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatewayserver
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
				return ErrInvalidLengthGatewayserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GatewayStatus == nil {
				m.GatewayStatus = &GatewayStatus{}
			}
			if err := m.GatewayStatus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGatewayserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGatewayserver
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
func (m *GatewayDown) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGatewayserver
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
			return fmt.Errorf("proto: GatewayDown: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GatewayDown: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DownlinkMessage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatewayserver
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
				return ErrInvalidLengthGatewayserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DownlinkMessage == nil {
				m.DownlinkMessage = &DownlinkMessage{}
			}
			if err := m.DownlinkMessage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGatewayserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGatewayserver
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
func skipGatewayserver(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGatewayserver
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
					return 0, ErrIntOverflowGatewayserver
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
					return 0, ErrIntOverflowGatewayserver
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
				return 0, ErrInvalidLengthGatewayserver
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGatewayserver
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
				next, err := skipGatewayserver(dAtA[start:])
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
	ErrInvalidLengthGatewayserver = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGatewayserver   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/gatewayserver.proto", fileDescriptorGatewayserver)
}

var fileDescriptorGatewayserver = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbf, 0xae, 0xd3, 0x30,
	0x18, 0xc5, 0xaf, 0x2f, 0x97, 0x2b, 0xe1, 0xab, 0x42, 0x31, 0xa2, 0x54, 0xa9, 0x14, 0x55, 0x9d,
	0xca, 0x80, 0x03, 0xe9, 0xc0, 0x00, 0x53, 0x29, 0x8a, 0x10, 0x50, 0xd4, 0x7f, 0x42, 0x62, 0xa9,
	0x92, 0xc6, 0x75, 0xac, 0xb4, 0xb6, 0x15, 0x3b, 0x8d, 0xd8, 0xd8, 0x78, 0x35, 0x46, 0x46, 0x46,
	0xc8, 0xc4, 0xc8, 0x23, 0xa0, 0x24, 0x8e, 0x08, 0x91, 0x10, 0x30, 0x25, 0x9f, 0x8f, 0xce, 0x39,
	0x3f, 0x7f, 0x32, 0x7c, 0x4c, 0x99, 0x8e, 0xd2, 0x00, 0xef, 0xc4, 0xd1, 0x59, 0x47, 0x64, 0x1d,
	0x31, 0x4e, 0xd5, 0x9c, 0xe8, 0x4c, 0x24, 0xb1, 0xa3, 0x35, 0x77, 0x7c, 0xc9, 0x1c, 0xea, 0x6b,
	0x92, 0xf9, 0xef, 0x15, 0x49, 0x4e, 0x24, 0xc1, 0x32, 0x11, 0x5a, 0xa0, 0x4b, 0xad, 0x39, 0x3e,
	0x4d, 0xac, 0x07, 0x8d, 0x00, 0x2a, 0xa8, 0x70, 0x4a, 0x39, 0x48, 0xf7, 0xe5, 0x54, 0x0e, 0xe5,
	0x5f, 0x65, 0xb3, 0x1e, 0xfd, 0x47, 0x9f, 0xb1, 0xb8, 0xff, 0x62, 0x39, 0x12, 0xa5, 0x7c, 0x4a,
	0x94, 0xf1, 0x0c, 0xa8, 0x10, 0xf4, 0x40, 0x7e, 0xc1, 0x90, 0xa3, 0xd4, 0x26, 0x70, 0xf4, 0x11,
	0xc0, 0x1b, 0x5e, 0x55, 0xb1, 0x91, 0xe8, 0x29, 0xbc, 0x99, 0xca, 0x03, 0xe3, 0xf1, 0xd6, 0x64,
	0xf4, 0xc1, 0xf0, 0xda, 0xf8, 0xca, 0xbd, 0x8b, 0xab, 0x1b, 0xe2, 0x4d, 0xa9, 0xbe, 0xae, 0xc4,
	0x65, 0x27, 0x6d, 0x8e, 0x85, 0xdb, 0xd0, 0x6e, 0x95, 0xf6, 0x75, 0xaa, 0xfa, 0xe7, 0x43, 0xd0,
	0x74, 0x9b, 0xa2, 0x55, 0x29, 0x2e, 0x3b, 0xb4, 0x39, 0x8e, 0x16, 0xf0, 0xca, 0xe8, 0x33, 0x91,
	0x71, 0x34, 0x85, 0xdd, 0x50, 0x64, 0xbc, 0x05, 0x53, 0xc4, 0xdd, 0xab, 0xe3, 0x66, 0x46, 0xaf,
	0x71, 0x6e, 0x85, 0xbf, 0x1f, 0xb8, 0x4f, 0xe0, 0x75, 0x4f, 0x67, 0x9e, 0x42, 0x2e, 0xbc, 0x78,
	0xc5, 0x78, 0x8c, 0x6e, 0xb7, 0x48, 0x36, 0xd2, 0xba, 0xd3, 0x3a, 0x2a, 0x42, 0xc7, 0xe0, 0x21,
	0x70, 0x5f, 0xc2, 0x8b, 0xb9, 0xf2, 0x14, 0x7a, 0x06, 0xbb, 0xab, 0x5d, 0x44, 0xc2, 0xf4, 0x40,
	0xea, 0x42, 0xf4, 0x27, 0x04, 0xab, 0x87, 0xab, 0x65, 0xe3, 0x7a, 0xd9, 0xf8, 0x79, 0xb1, 0x6c,
	0xf7, 0x2d, 0x3c, 0xf7, 0x14, 0x5a, 0xc0, 0x9e, 0x47, 0xb4, 0x29, 0x7a, 0x13, 0x14, 0x6f, 0xc8,
	0xd7, 0x4c, 0x70, 0x85, 0xac, 0x16, 0xc5, 0x8b, 0x90, 0x70, 0xcd, 0xf6, 0x8c, 0x24, 0xca, 0x1a,
	0xb4, 0xb4, 0xa6, 0x71, 0xea, 0x7d, 0xf9, 0x66, 0x9f, 0x7d, 0xc8, 0x6d, 0xf0, 0x29, 0xb7, 0xc1,
	0xe7, 0xdc, 0x06, 0x5f, 0x73, 0x1b, 0x7c, 0xcf, 0xed, 0xb3, 0x1f, 0xb9, 0x0d, 0xde, 0xdd, 0xff,
	0xdb, 0x73, 0x91, 0x31, 0x2d, 0xbe, 0x32, 0x08, 0x2e, 0x4b, 0xe2, 0xc9, 0xcf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xfe, 0xf7, 0x22, 0x76, 0x05, 0x03, 0x00, 0x00,
}
