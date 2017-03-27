// Code generated by protoc-gen-go.
// source: http_request.proto
// DO NOT EDIT!

package gwproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HttpRequest_NetType int32

const (
	HttpRequest_UNKNOW HttpRequest_NetType = 0
	HttpRequest_WIFI   HttpRequest_NetType = 1
	HttpRequest_NT2G   HttpRequest_NetType = 2
	HttpRequest_NT3G   HttpRequest_NetType = 3
	HttpRequest_NT4G   HttpRequest_NetType = 4
	HttpRequest_NT5G   HttpRequest_NetType = 5
)

var HttpRequest_NetType_name = map[int32]string{
	0: "UNKNOW",
	1: "WIFI",
	2: "NT2G",
	3: "NT3G",
	4: "NT4G",
	5: "NT5G",
}
var HttpRequest_NetType_value = map[string]int32{
	"UNKNOW": 0,
	"WIFI":   1,
	"NT2G":   2,
	"NT3G":   3,
	"NT4G":   4,
	"NT5G":   5,
}

func (x HttpRequest_NetType) String() string {
	return proto.EnumName(HttpRequest_NetType_name, int32(x))
}
func (HttpRequest_NetType) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

type HttpRequest_HttpMethod int32

const (
	HttpRequest_UNKNOWN HttpRequest_HttpMethod = 0
	HttpRequest_GET     HttpRequest_HttpMethod = 1
	HttpRequest_POST    HttpRequest_HttpMethod = 2
	HttpRequest_HEAD    HttpRequest_HttpMethod = 3
	HttpRequest_PUT     HttpRequest_HttpMethod = 4
	HttpRequest_DELETE  HttpRequest_HttpMethod = 5
	HttpRequest_TRACE   HttpRequest_HttpMethod = 6
	HttpRequest_CONNECT HttpRequest_HttpMethod = 7
)

var HttpRequest_HttpMethod_name = map[int32]string{
	0: "UNKNOWN",
	1: "GET",
	2: "POST",
	3: "HEAD",
	4: "PUT",
	5: "DELETE",
	6: "TRACE",
	7: "CONNECT",
}
var HttpRequest_HttpMethod_value = map[string]int32{
	"UNKNOWN": 0,
	"GET":     1,
	"POST":    2,
	"HEAD":    3,
	"PUT":     4,
	"DELETE":  5,
	"TRACE":   6,
	"CONNECT": 7,
}

func (x HttpRequest_HttpMethod) String() string {
	return proto.EnumName(HttpRequest_HttpMethod_name, int32(x))
}
func (HttpRequest_HttpMethod) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 1} }

type HttpRequest struct {
	ReqId     string                 `protobuf:"bytes,1,opt,name=req_id,json=reqId" json:"req_id,omitempty"`
	Net       HttpRequest_NetType    `protobuf:"varint,2,opt,name=net,enum=HttpRequest_NetType" json:"net,omitempty"`
	Time      int64                  `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	Host      string                 `protobuf:"bytes,4,opt,name=host" json:"host,omitempty"`
	Url       string                 `protobuf:"bytes,5,opt,name=url" json:"url,omitempty"`
	Port      int32                  `protobuf:"varint,6,opt,name=port" json:"port,omitempty"`
	Method    HttpRequest_HttpMethod `protobuf:"varint,7,opt,name=method,enum=HttpRequest_HttpMethod" json:"method,omitempty"`
	ReqLen    int64                  `protobuf:"varint,8,opt,name=req_len,json=reqLen" json:"req_len,omitempty"`
	Query     string                 `protobuf:"bytes,9,opt,name=query" json:"query,omitempty"`
	BeginTime int64                  `protobuf:"varint,10,opt,name=begin_time,json=beginTime" json:"begin_time,omitempty"`
	DnsStart  int64                  `protobuf:"varint,11,opt,name=dns_start,json=dnsStart" json:"dns_start,omitempty"`
	DnsEnd    int64                  `protobuf:"varint,12,opt,name=dns_end,json=dnsEnd" json:"dns_end,omitempty"`
	TcpStart  int64                  `protobuf:"varint,13,opt,name=tcp_start,json=tcpStart" json:"tcp_start,omitempty"`
	TcpEnd    int64                  `protobuf:"varint,14,opt,name=tcp_end,json=tcpEnd" json:"tcp_end,omitempty"`
	SslStart  int64                  `protobuf:"varint,15,opt,name=ssl_start,json=sslStart" json:"ssl_start,omitempty"`
	SslEnd    int64                  `protobuf:"varint,16,opt,name=ssl_end,json=sslEnd" json:"ssl_end,omitempty"`
	ReqStart  int64                  `protobuf:"varint,17,opt,name=req_start,json=reqStart" json:"req_start,omitempty"`
	ReqEnd    int64                  `protobuf:"varint,18,opt,name=req_end,json=reqEnd" json:"req_end,omitempty"`
	RespStart int64                  `protobuf:"varint,19,opt,name=resp_start,json=respStart" json:"resp_start,omitempty"`
	RespEnd   int64                  `protobuf:"varint,20,opt,name=resp_end,json=respEnd" json:"resp_end,omitempty"`
	// Types that are valid to be assigned to Record:
	//	*HttpRequest_Complete
	//	*HttpRequest_NetErr
	Record isHttpRequest_Record `protobuf_oneof:"record"`
}

func (m *HttpRequest) Reset()                    { *m = HttpRequest{} }
func (m *HttpRequest) String() string            { return proto.CompactTextString(m) }
func (*HttpRequest) ProtoMessage()               {}
func (*HttpRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type isHttpRequest_Record interface {
	isHttpRequest_Record()
}

type HttpRequest_Complete struct {
	Complete *Complete `protobuf:"bytes,21,opt,name=complete,oneof"`
}
type HttpRequest_NetErr struct {
	NetErr *NetErr `protobuf:"bytes,22,opt,name=net_err,json=netErr,oneof"`
}

func (*HttpRequest_Complete) isHttpRequest_Record() {}
func (*HttpRequest_NetErr) isHttpRequest_Record()   {}

func (m *HttpRequest) GetRecord() isHttpRequest_Record {
	if m != nil {
		return m.Record
	}
	return nil
}

func (m *HttpRequest) GetReqId() string {
	if m != nil {
		return m.ReqId
	}
	return ""
}

func (m *HttpRequest) GetNet() HttpRequest_NetType {
	if m != nil {
		return m.Net
	}
	return HttpRequest_UNKNOW
}

func (m *HttpRequest) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *HttpRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *HttpRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *HttpRequest) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *HttpRequest) GetMethod() HttpRequest_HttpMethod {
	if m != nil {
		return m.Method
	}
	return HttpRequest_UNKNOWN
}

func (m *HttpRequest) GetReqLen() int64 {
	if m != nil {
		return m.ReqLen
	}
	return 0
}

func (m *HttpRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *HttpRequest) GetBeginTime() int64 {
	if m != nil {
		return m.BeginTime
	}
	return 0
}

func (m *HttpRequest) GetDnsStart() int64 {
	if m != nil {
		return m.DnsStart
	}
	return 0
}

func (m *HttpRequest) GetDnsEnd() int64 {
	if m != nil {
		return m.DnsEnd
	}
	return 0
}

func (m *HttpRequest) GetTcpStart() int64 {
	if m != nil {
		return m.TcpStart
	}
	return 0
}

func (m *HttpRequest) GetTcpEnd() int64 {
	if m != nil {
		return m.TcpEnd
	}
	return 0
}

func (m *HttpRequest) GetSslStart() int64 {
	if m != nil {
		return m.SslStart
	}
	return 0
}

func (m *HttpRequest) GetSslEnd() int64 {
	if m != nil {
		return m.SslEnd
	}
	return 0
}

func (m *HttpRequest) GetReqStart() int64 {
	if m != nil {
		return m.ReqStart
	}
	return 0
}

func (m *HttpRequest) GetReqEnd() int64 {
	if m != nil {
		return m.ReqEnd
	}
	return 0
}

func (m *HttpRequest) GetRespStart() int64 {
	if m != nil {
		return m.RespStart
	}
	return 0
}

func (m *HttpRequest) GetRespEnd() int64 {
	if m != nil {
		return m.RespEnd
	}
	return 0
}

func (m *HttpRequest) GetComplete() *Complete {
	if x, ok := m.GetRecord().(*HttpRequest_Complete); ok {
		return x.Complete
	}
	return nil
}

func (m *HttpRequest) GetNetErr() *NetErr {
	if x, ok := m.GetRecord().(*HttpRequest_NetErr); ok {
		return x.NetErr
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HttpRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HttpRequest_OneofMarshaler, _HttpRequest_OneofUnmarshaler, _HttpRequest_OneofSizer, []interface{}{
		(*HttpRequest_Complete)(nil),
		(*HttpRequest_NetErr)(nil),
	}
}

func _HttpRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HttpRequest)
	// record
	switch x := m.Record.(type) {
	case *HttpRequest_Complete:
		b.EncodeVarint(21<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Complete); err != nil {
			return err
		}
	case *HttpRequest_NetErr:
		b.EncodeVarint(22<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NetErr); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("HttpRequest.Record has unexpected type %T", x)
	}
	return nil
}

func _HttpRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HttpRequest)
	switch tag {
	case 21: // record.complete
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Complete)
		err := b.DecodeMessage(msg)
		m.Record = &HttpRequest_Complete{msg}
		return true, err
	case 22: // record.net_err
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NetErr)
		err := b.DecodeMessage(msg)
		m.Record = &HttpRequest_NetErr{msg}
		return true, err
	default:
		return false, nil
	}
}

func _HttpRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HttpRequest)
	// record
	switch x := m.Record.(type) {
	case *HttpRequest_Complete:
		s := proto.Size(x.Complete)
		n += proto.SizeVarint(21<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *HttpRequest_NetErr:
		s := proto.Size(x.NetErr)
		n += proto.SizeVarint(22<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*HttpRequest)(nil), "HttpRequest")
	proto.RegisterEnum("HttpRequest_NetType", HttpRequest_NetType_name, HttpRequest_NetType_value)
	proto.RegisterEnum("HttpRequest_HttpMethod", HttpRequest_HttpMethod_name, HttpRequest_HttpMethod_value)
}

func init() { proto.RegisterFile("http_request.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x93, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0x9b, 0xa5, 0xf9, 0x3a, 0x65, 0xc5, 0x98, 0x8e, 0x19, 0xd0, 0xa4, 0xaa, 0x17, 0xd0,
	0xab, 0x22, 0x75, 0xf0, 0x00, 0xa3, 0x0b, 0x69, 0xc5, 0x48, 0xa7, 0x2c, 0xd3, 0x2e, 0x2b, 0x96,
	0x58, 0xb4, 0x52, 0xea, 0xa4, 0xb6, 0x7b, 0xb1, 0xb7, 0xe2, 0x11, 0xd1, 0xb1, 0xb3, 0x00, 0x77,
	0xff, 0x73, 0xfe, 0xfe, 0x9d, 0x8f, 0xe8, 0x04, 0xe8, 0x56, 0xeb, 0x66, 0x23, 0xf9, 0xe1, 0xc8,
	0x95, 0x9e, 0x35, 0xb2, 0xd6, 0xf5, 0xbb, 0x61, 0x51, 0xef, 0x9b, 0x8a, 0x6b, 0xde, 0xc6, 0xa7,
	0x82, 0xeb, 0x0d, 0x97, 0xd2, 0x86, 0x93, 0xdf, 0x3e, 0x0c, 0x96, 0x5a, 0x37, 0x99, 0x85, 0xe8,
	0x19, 0xf8, 0x92, 0x1f, 0x36, 0xbb, 0x92, 0x39, 0x63, 0x67, 0x1a, 0x65, 0x9e, 0xe4, 0x87, 0x55,
	0x49, 0x3f, 0x80, 0x2b, 0xb8, 0x66, 0x27, 0x63, 0x67, 0x3a, 0x9c, 0x8f, 0x66, 0xff, 0x10, 0xb3,
	0x94, 0xeb, 0xfc, 0xa9, 0xe1, 0x19, 0x3e, 0xa0, 0x14, 0xfa, 0x7a, 0xb7, 0xe7, 0xcc, 0x1d, 0x3b,
	0x53, 0x37, 0x33, 0x1a, 0x73, 0xdb, 0x5a, 0x69, 0xd6, 0x37, 0x05, 0x8d, 0xa6, 0x04, 0xdc, 0xa3,
	0xac, 0x98, 0x67, 0x52, 0x28, 0xf1, 0x55, 0x53, 0x4b, 0xcd, 0xfc, 0xb1, 0x33, 0xf5, 0x32, 0xa3,
	0xe9, 0x27, 0xf0, 0xf7, 0x5c, 0x6f, 0xeb, 0x92, 0x05, 0xa6, 0xf1, 0xf9, 0x7f, 0x8d, 0x51, 0xff,
	0x30, 0x76, 0xd6, 0x3e, 0xa3, 0xe7, 0x10, 0xe0, 0xf4, 0x15, 0x17, 0x2c, 0x34, 0x13, 0xe0, 0x32,
	0x37, 0x5c, 0xd0, 0x11, 0x78, 0x87, 0x23, 0x97, 0x4f, 0x2c, 0xb2, 0x5b, 0x99, 0x80, 0x5e, 0x00,
	0x3c, 0xf2, 0x5f, 0x3b, 0xb1, 0x31, 0x33, 0x83, 0x21, 0x22, 0x93, 0xc9, 0x71, 0xf0, 0xf7, 0x10,
	0x95, 0x42, 0x6d, 0x94, 0xfe, 0x29, 0x35, 0x1b, 0x18, 0x37, 0x2c, 0x85, 0xba, 0xc3, 0x18, 0x5b,
	0xa1, 0xc9, 0x45, 0xc9, 0x5e, 0xd8, 0x56, 0xa5, 0x50, 0xb1, 0x28, 0x91, 0xd2, 0x45, 0xd3, 0x52,
	0xa7, 0x96, 0xd2, 0x45, 0xd3, 0x51, 0x68, 0x22, 0x35, 0xb4, 0x94, 0x2e, 0x9a, 0x96, 0x52, 0xaa,
	0x6a, 0xa9, 0x97, 0x96, 0x52, 0xaa, 0xea, 0x28, 0x34, 0x91, 0x22, 0x96, 0x52, 0xaa, 0x6a, 0x29,
	0xdc, 0xd7, 0x52, 0xaf, 0x2c, 0x25, 0xf9, 0xa1, 0xa3, 0xd0, 0x44, 0x8a, 0x76, 0x1f, 0x03, 0xa9,
	0x0b, 0x00, 0xc9, 0xd5, 0xf3, 0x88, 0xaf, 0xed, 0xda, 0x98, 0xb1, 0xdc, 0x5b, 0x08, 0x8d, 0x8d,
	0xe0, 0xc8, 0x98, 0x01, 0xc6, 0x48, 0x7e, 0x84, 0xf0, 0xf9, 0x9c, 0xd8, 0xd9, 0xd8, 0x99, 0x0e,
	0xe6, 0xd1, 0x6c, 0xd1, 0x26, 0x96, 0xbd, 0xac, 0x33, 0xe9, 0x04, 0x82, 0xf6, 0xce, 0xd8, 0x1b,
	0xf3, 0x2e, 0xc0, 0x3b, 0x89, 0xa5, 0x5c, 0xf6, 0x32, 0x5f, 0x18, 0x35, 0x49, 0x20, 0x68, 0x6f,
	0x87, 0x02, 0xf8, 0xf7, 0xe9, 0xf7, 0x74, 0xfd, 0x40, 0x7a, 0x34, 0x84, 0xfe, 0xc3, 0xea, 0xdb,
	0x8a, 0x38, 0xa8, 0xd2, 0x7c, 0x9e, 0x90, 0x13, 0xab, 0x2e, 0x13, 0xe2, 0x5a, 0xf5, 0x39, 0x21,
	0x7d, 0xab, 0xbe, 0x24, 0xc4, 0x9b, 0x14, 0x00, 0x7f, 0x6f, 0x81, 0x0e, 0x20, 0xb0, 0xb5, 0x52,
	0xd2, 0xa3, 0x01, 0xb8, 0x49, 0x9c, 0xdb, 0x5a, 0xb7, 0xeb, 0xbb, 0xdc, 0xd6, 0x5a, 0xc6, 0x57,
	0xd7, 0xc4, 0x45, 0xf3, 0xf6, 0x3e, 0x27, 0x7d, 0x6c, 0x7f, 0x1d, 0xdf, 0xc4, 0x79, 0x4c, 0x3c,
	0x1a, 0x81, 0x97, 0x67, 0x57, 0x8b, 0x98, 0xf8, 0x58, 0x69, 0xb1, 0x4e, 0xd3, 0x78, 0x91, 0x93,
	0xe0, 0x6b, 0x88, 0x3f, 0x46, 0x51, 0xcb, 0xf2, 0xd1, 0x37, 0x7f, 0xce, 0xe5, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x5b, 0x4a, 0x2c, 0x9a, 0x6e, 0x03, 0x00, 0x00,
}
