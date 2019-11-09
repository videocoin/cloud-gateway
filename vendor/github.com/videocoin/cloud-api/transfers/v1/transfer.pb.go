// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transfers/v1/transfer.proto

package v1

import (
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TransferKind int32

const (
	TransferKindNone     TransferKind = 0
	TransferKindWithdraw TransferKind = 1
)

var TransferKind_name = map[int32]string{
	0: "TRANSFER_KIND_NONE",
	1: "TRANSFER_KIND_WITHDRAW",
}

var TransferKind_value = map[string]int32{
	"TRANSFER_KIND_NONE":     0,
	"TRANSFER_KIND_WITHDRAW": 1,
}

func (x TransferKind) String() string {
	return proto.EnumName(TransferKind_name, int32(x))
}

func (TransferKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_01623e3fc229fa7c, []int{0}
}

type TransferStatus int32

const (
	TransferStatusNone           TransferStatus = 0
	TransferStatusNew            TransferStatus = 1
	TransferStatusPendingNative  TransferStatus = 2
	TransferStatusExecutedNative TransferStatus = 3
	TransferStatusPendingErc     TransferStatus = 4
	TransferStatusCompleted      TransferStatus = 5
	TransferStatusFailed         TransferStatus = 6
)

var TransferStatus_name = map[int32]string{
	0: "TRANSFER_STATUS_NONE",
	1: "TRANSFER_STATUS_NEW",
	2: "TRANSFER_STATUS_PENDING_NATIVE",
	3: "TRANSFER_STATUS_EXECUTED_NATIVE",
	4: "TRANSFER_STATUS_PENDING_ERC",
	5: "TRANSFER_STATUS_COMPLETED",
	6: "TRANSFER_STATUS_FAILED",
}

var TransferStatus_value = map[string]int32{
	"TRANSFER_STATUS_NONE":            0,
	"TRANSFER_STATUS_NEW":             1,
	"TRANSFER_STATUS_PENDING_NATIVE":  2,
	"TRANSFER_STATUS_EXECUTED_NATIVE": 3,
	"TRANSFER_STATUS_PENDING_ERC":     4,
	"TRANSFER_STATUS_COMPLETED":       5,
	"TRANSFER_STATUS_FAILED":          6,
}

func (x TransferStatus) String() string {
	return proto.EnumName(TransferStatus_name, int32(x))
}

func (TransferStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_01623e3fc229fa7c, []int{1}
}

type Transfer struct {
	Id                   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" gorm:"type:varchar(36);primary_key"`
	UserId               string         `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" gorm:"type:varchar(36)"`
	Pin                  string         `protobuf:"bytes,3,opt,name=pin,proto3" json:"pin,omitempty" gorm:"type:varchar(6)"`
	Kind                 TransferKind   `protobuf:"varint,4,opt,name=kind,proto3,enum=cloud.api.transfers.v1.TransferKind" json:"kind,omitempty"`
	Status               TransferStatus `protobuf:"varint,5,opt,name=status,proto3,enum=cloud.api.transfers.v1.TransferStatus" json:"status,omitempty"`
	ToAddress            string         `protobuf:"bytes,6,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Amount               []byte         `protobuf:"bytes,7,opt,name=amount,proto3" json:"amount,omitempty"`
	TxNativeId           []byte         `protobuf:"bytes,10,opt,name=tx_native_id,json=txNativeId,proto3" json:"tx_native_id,omitempty"`
	TxErc20Id            []byte         `protobuf:"bytes,11,opt,name=tx_erc20_id,json=txErc20Id,proto3" json:"tx_erc20_id,omitempty"`
	CreatedAt            *time.Time     `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at,omitempty"`
	ExpiresAt            *time.Time     `protobuf:"bytes,15,opt,name=expires_at,json=expiresAt,proto3,stdtime" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Transfer) Reset()         { *m = Transfer{} }
func (m *Transfer) String() string { return proto.CompactTextString(m) }
func (*Transfer) ProtoMessage()    {}
func (*Transfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_01623e3fc229fa7c, []int{0}
}
func (m *Transfer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Transfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Transfer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Transfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transfer.Merge(m, src)
}
func (m *Transfer) XXX_Size() int {
	return m.Size()
}
func (m *Transfer) XXX_DiscardUnknown() {
	xxx_messageInfo_Transfer.DiscardUnknown(m)
}

var xxx_messageInfo_Transfer proto.InternalMessageInfo

func (m *Transfer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Transfer) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Transfer) GetPin() string {
	if m != nil {
		return m.Pin
	}
	return ""
}

func (m *Transfer) GetKind() TransferKind {
	if m != nil {
		return m.Kind
	}
	return TransferKindNone
}

func (m *Transfer) GetStatus() TransferStatus {
	if m != nil {
		return m.Status
	}
	return TransferStatusNone
}

func (m *Transfer) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *Transfer) GetAmount() []byte {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *Transfer) GetTxNativeId() []byte {
	if m != nil {
		return m.TxNativeId
	}
	return nil
}

func (m *Transfer) GetTxErc20Id() []byte {
	if m != nil {
		return m.TxErc20Id
	}
	return nil
}

func (m *Transfer) GetCreatedAt() *time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Transfer) GetExpiresAt() *time.Time {
	if m != nil {
		return m.ExpiresAt
	}
	return nil
}

func (*Transfer) XXX_MessageName() string {
	return "cloud.api.transfers.v1.Transfer"
}
func init() {
	proto.RegisterEnum("cloud.api.transfers.v1.TransferKind", TransferKind_name, TransferKind_value)
	golang_proto.RegisterEnum("cloud.api.transfers.v1.TransferKind", TransferKind_name, TransferKind_value)
	proto.RegisterEnum("cloud.api.transfers.v1.TransferStatus", TransferStatus_name, TransferStatus_value)
	golang_proto.RegisterEnum("cloud.api.transfers.v1.TransferStatus", TransferStatus_name, TransferStatus_value)
	proto.RegisterType((*Transfer)(nil), "cloud.api.transfers.v1.Transfer")
	golang_proto.RegisterType((*Transfer)(nil), "cloud.api.transfers.v1.Transfer")
}

func init() { proto.RegisterFile("transfers/v1/transfer.proto", fileDescriptor_01623e3fc229fa7c) }
func init() { golang_proto.RegisterFile("transfers/v1/transfer.proto", fileDescriptor_01623e3fc229fa7c) }

var fileDescriptor_01623e3fc229fa7c = []byte{
	// 746 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6e, 0xe3, 0x44,
	0x1c, 0xc6, 0xd7, 0x49, 0x36, 0x4b, 0xa6, 0x55, 0x09, 0x43, 0xc8, 0x1a, 0xa7, 0x38, 0x56, 0x40,
	0x10, 0x56, 0x8b, 0xb3, 0xed, 0xae, 0x00, 0x15, 0x01, 0x72, 0x93, 0x29, 0x58, 0xbb, 0x78, 0x2b,
	0xc7, 0x4b, 0x11, 0x17, 0x6b, 0x9a, 0x99, 0xa6, 0xa3, 0x26, 0x1e, 0xcb, 0x9e, 0xa4, 0xe9, 0x1b,
	0x54, 0x3e, 0xf1, 0x00, 0xf8, 0x04, 0x4f, 0xc1, 0x89, 0x1b, 0x3d, 0xf2, 0x04, 0x05, 0xb5, 0x6f,
	0xd0, 0x27, 0x40, 0x76, 0x9c, 0x52, 0x87, 0x20, 0xb8, 0xf9, 0x3f, 0xf3, 0xfd, 0xbe, 0xf9, 0x3c,
	0xdf, 0x80, 0x86, 0x08, 0xb0, 0x17, 0x1e, 0xd1, 0x20, 0xec, 0x4c, 0xb7, 0x3a, 0x8b, 0x41, 0xf7,
	0x03, 0x2e, 0x38, 0xac, 0x0f, 0x46, 0x7c, 0x42, 0x74, 0xec, 0x33, 0xfd, 0x56, 0xa6, 0x4f, 0xb7,
	0x94, 0xe6, 0x90, 0xf3, 0xe1, 0x88, 0x76, 0x52, 0xd5, 0xe1, 0xe4, 0xa8, 0x23, 0xd8, 0x98, 0x86,
	0x02, 0x8f, 0xfd, 0x39, 0xa8, 0x6c, 0x66, 0x02, 0xec, 0xb3, 0x0e, 0xf6, 0x3c, 0x2e, 0xb0, 0x60,
	0xdc, 0x0b, 0xb3, 0xdd, 0x8f, 0x86, 0x4c, 0x1c, 0x4f, 0x0e, 0xf5, 0x01, 0x1f, 0x77, 0x86, 0x7c,
	0xc8, 0xff, 0xf6, 0x49, 0xa6, 0x74, 0x48, 0xbf, 0xe6, 0xf2, 0xd6, 0x8f, 0x25, 0xf0, 0x9a, 0x93,
	0x1d, 0x0f, 0x3f, 0x01, 0x05, 0x46, 0x64, 0x49, 0x93, 0xda, 0x95, 0xdd, 0x0f, 0x6e, 0x2e, 0x9b,
	0xef, 0x0e, 0x79, 0x30, 0xde, 0x69, 0x89, 0x33, 0x9f, 0xee, 0x4c, 0x71, 0x30, 0x38, 0xc6, 0x41,
	0xfb, 0xe9, 0xc7, 0x1f, 0x7e, 0xe6, 0x07, 0x6c, 0x8c, 0x83, 0x33, 0xf7, 0x84, 0x9e, 0xb5, 0xec,
	0x02, 0x23, 0xf0, 0x19, 0x78, 0x30, 0x09, 0x69, 0xe0, 0x32, 0x22, 0x17, 0x52, 0xba, 0x71, 0x73,
	0xd9, 0x7c, 0xb8, 0x9a, 0x6e, 0xd9, 0xe5, 0x44, 0x6b, 0x12, 0xf8, 0x18, 0x14, 0x7d, 0xe6, 0xc9,
	0xc5, 0x94, 0x50, 0x6e, 0x2e, 0x9b, 0xf5, 0x15, 0x44, 0x02, 0x24, 0x32, 0xf8, 0x29, 0x28, 0x9d,
	0x30, 0x8f, 0xc8, 0x25, 0x4d, 0x6a, 0x6f, 0x6c, 0xbf, 0xa7, 0xaf, 0xbe, 0x3e, 0x7d, 0xf1, 0x33,
	0xcf, 0x99, 0x47, 0xec, 0x94, 0x80, 0x5f, 0x80, 0x72, 0x28, 0xb0, 0x98, 0x84, 0xf2, 0xfd, 0x94,
	0x7d, 0xff, 0xbf, 0xd8, 0x7e, 0xaa, 0xb6, 0x33, 0x0a, 0xbe, 0x03, 0x80, 0xe0, 0x2e, 0x26, 0x24,
	0xa0, 0x61, 0x28, 0x97, 0x93, 0xb8, 0x76, 0x45, 0x70, 0x63, 0xbe, 0x00, 0xeb, 0xa0, 0x8c, 0xc7,
	0x7c, 0xe2, 0x09, 0xf9, 0x81, 0x26, 0xb5, 0xd7, 0xed, 0x6c, 0x82, 0x1a, 0x58, 0x17, 0x33, 0xd7,
	0xc3, 0x82, 0x4d, 0x69, 0x72, 0x33, 0x20, 0xdd, 0x05, 0x62, 0x66, 0xa5, 0x4b, 0x26, 0x81, 0x2a,
	0x58, 0x13, 0x33, 0x97, 0x06, 0x83, 0xed, 0x27, 0x89, 0x60, 0x2d, 0x15, 0x54, 0xc4, 0x0c, 0x25,
	0x2b, 0x26, 0x81, 0x5f, 0x02, 0x30, 0x08, 0x28, 0x16, 0x94, 0xb8, 0x58, 0xc8, 0x1b, 0x9a, 0xd4,
	0x5e, 0xdb, 0x56, 0xf4, 0x79, 0xfd, 0xfa, 0xa2, 0x57, 0xdd, 0x59, 0xbc, 0x8f, 0xdd, 0xd2, 0x0f,
	0x7f, 0x34, 0x25, 0xbb, 0x92, 0x31, 0x86, 0x48, 0x0c, 0xe8, 0xcc, 0x67, 0x01, 0x0d, 0x13, 0x83,
	0xd7, 0xff, 0xaf, 0x41, 0xc6, 0x18, 0xe2, 0xd1, 0xb9, 0x04, 0xd6, 0xef, 0xde, 0x28, 0x7c, 0x0c,
	0xa0, 0x63, 0x1b, 0x56, 0x7f, 0x0f, 0xd9, 0xee, 0x73, 0xd3, 0xea, 0xb9, 0xd6, 0x4b, 0x0b, 0x55,
	0xef, 0x29, 0xb5, 0x28, 0xd6, 0xaa, 0x77, 0x95, 0x16, 0xf7, 0x28, 0x7c, 0x06, 0xea, 0x79, 0xf5,
	0x81, 0xe9, 0x7c, 0xdd, 0xb3, 0x8d, 0x83, 0xaa, 0xa4, 0xc8, 0x51, 0xac, 0xd5, 0xee, 0x12, 0x07,
	0x4c, 0x1c, 0x93, 0x00, 0x9f, 0x2a, 0xb5, 0xf3, 0x9f, 0xd4, 0x7b, 0xbf, 0xfc, 0xac, 0xe6, 0x4e,
	0x7e, 0xf4, 0x5b, 0x11, 0x6c, 0xe4, 0x0b, 0x82, 0x4f, 0x40, 0xed, 0xd6, 0xbe, 0xef, 0x18, 0xce,
	0xab, 0xfe, 0x22, 0x4e, 0x3d, 0x8a, 0x35, 0x98, 0x57, 0xa7, 0x81, 0x74, 0xf0, 0xe6, 0x3f, 0x08,
	0x94, 0xa4, 0x79, 0x2b, 0x8a, 0xb5, 0x37, 0x96, 0x00, 0x7a, 0x0a, 0xbb, 0x40, 0x5d, 0xd6, 0xef,
	0x23, 0xab, 0x67, 0x5a, 0x5f, 0xb9, 0x96, 0xe1, 0x98, 0xdf, 0xa2, 0x6a, 0x41, 0x69, 0x46, 0xb1,
	0xd6, 0xc8, 0xa3, 0xfb, 0xd4, 0x23, 0xcc, 0x1b, 0xce, 0x8b, 0x86, 0x08, 0x34, 0x97, 0x4d, 0xd0,
	0x77, 0xa8, 0xfb, 0xca, 0x41, 0xbd, 0x85, 0x4b, 0x51, 0xd1, 0xa2, 0x58, 0xdb, 0xcc, 0xbb, 0xa0,
	0x19, 0x1d, 0x4c, 0x04, 0x25, 0x99, 0xcd, 0xe7, 0xa0, 0xf1, 0x6f, 0x59, 0x90, 0xdd, 0xad, 0x96,
	0x94, 0xcd, 0x28, 0xd6, 0xe4, 0x95, 0x41, 0x50, 0x30, 0x80, 0x3b, 0xe0, 0xed, 0x65, 0xbc, 0xfb,
	0xf2, 0x9b, 0xfd, 0x17, 0xc8, 0x41, 0xbd, 0xea, 0x7d, 0xa5, 0x11, 0xc5, 0xda, 0xc3, 0x3c, 0xdc,
	0xe5, 0x63, 0x7f, 0x44, 0x05, 0x25, 0xb9, 0x1e, 0x33, 0x76, 0xcf, 0x30, 0x5f, 0xa0, 0x5e, 0xb5,
	0x9c, 0xef, 0x71, 0x0e, 0xee, 0x61, 0x36, 0xa2, 0x44, 0xa9, 0x67, 0x3d, 0x2e, 0xd5, 0xb6, 0x2b,
	0x5f, 0x5c, 0xa9, 0xd2, 0xef, 0x57, 0xaa, 0xf4, 0xe7, 0x95, 0x2a, 0xfd, 0x7a, 0xad, 0x4a, 0x17,
	0xd7, 0xaa, 0xf4, 0x7d, 0x61, 0xba, 0x75, 0x58, 0x4e, 0xdf, 0xe4, 0xd3, 0xbf, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xef, 0xf0, 0xe6, 0x70, 0x39, 0x05, 0x00, 0x00,
}

func (m *Transfer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Transfer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Transfer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ExpiresAt != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.ExpiresAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.ExpiresAt):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintTransfer(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x7a
	}
	if m.CreatedAt != nil {
		n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreatedAt):])
		if err2 != nil {
			return 0, err2
		}
		i -= n2
		i = encodeVarintTransfer(dAtA, i, uint64(n2))
		i--
		dAtA[i] = 0x72
	}
	if len(m.TxErc20Id) > 0 {
		i -= len(m.TxErc20Id)
		copy(dAtA[i:], m.TxErc20Id)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.TxErc20Id)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.TxNativeId) > 0 {
		i -= len(m.TxNativeId)
		copy(dAtA[i:], m.TxNativeId)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.TxNativeId)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ToAddress) > 0 {
		i -= len(m.ToAddress)
		copy(dAtA[i:], m.ToAddress)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.ToAddress)))
		i--
		dAtA[i] = 0x32
	}
	if m.Status != 0 {
		i = encodeVarintTransfer(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if m.Kind != 0 {
		i = encodeVarintTransfer(dAtA, i, uint64(m.Kind))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Pin) > 0 {
		i -= len(m.Pin)
		copy(dAtA[i:], m.Pin)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.Pin)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.UserId) > 0 {
		i -= len(m.UserId)
		copy(dAtA[i:], m.UserId)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.UserId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTransfer(dAtA []byte, offset int, v uint64) int {
	offset -= sovTransfer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Transfer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	l = len(m.Pin)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	if m.Kind != 0 {
		n += 1 + sovTransfer(uint64(m.Kind))
	}
	if m.Status != 0 {
		n += 1 + sovTransfer(uint64(m.Status))
	}
	l = len(m.ToAddress)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	l = len(m.TxNativeId)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	l = len(m.TxErc20Id)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	if m.CreatedAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreatedAt)
		n += 1 + l + sovTransfer(uint64(l))
	}
	if m.ExpiresAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.ExpiresAt)
		n += 1 + l + sovTransfer(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTransfer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTransfer(x uint64) (n int) {
	return sovTransfer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Transfer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransfer
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
			return fmt.Errorf("proto: Transfer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Transfer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
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
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
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
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
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
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			m.Kind = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Kind |= TransferKind(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= TransferStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
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
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
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
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount[:0], dAtA[iNdEx:postIndex]...)
			if m.Amount == nil {
				m.Amount = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxNativeId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
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
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxNativeId = append(m.TxNativeId[:0], dAtA[iNdEx:postIndex]...)
			if m.TxNativeId == nil {
				m.TxNativeId = []byte{}
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxErc20Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
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
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxErc20Id = append(m.TxErc20Id[:0], dAtA[iNdEx:postIndex]...)
			if m.TxErc20Id == nil {
				m.TxErc20Id = []byte{}
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
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
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreatedAt == nil {
				m.CreatedAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiresAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
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
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExpiresAt == nil {
				m.ExpiresAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.ExpiresAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTransfer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTransfer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTransfer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransfer
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
					return 0, ErrIntOverflowTransfer
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
					return 0, ErrIntOverflowTransfer
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
				return 0, ErrInvalidLengthTransfer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTransfer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTransfer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTransfer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransfer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTransfer = fmt.Errorf("proto: unexpected end of group")
)
