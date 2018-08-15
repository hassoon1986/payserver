// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

/*
Package crpc is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	EmptyRequest
	EmptyResponse
	CreateReceiptRequest
	CreateReceiptResponse
	BalanceRequest
	BalanceResponse
	ValidateReceiptRequest
	EstimateFeeRequest
	EstimateFeeResponse
	SendPaymentRequest
	PaymentByIDRequest
	PaymentsByReceiptRequest
	PaymentsByReceiptResponse
	ListPaymentsRequest
	ListPaymentsResponse
	Payment
*/
package crpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Asset is the list of a trading assets which are available in the exchange
// platform.
type Asset int32

const (
	Asset_ASSET_NONE Asset = 0
	//
	// Bitcoin
	Asset_BTC Asset = 1
	//
	// Bitcoin Cash
	Asset_BCH Asset = 2
	//
	// Ethereum
	Asset_ETH Asset = 3
	//
	// Litecoin
	Asset_LTC Asset = 4
	// Dash
	Asset_DASH Asset = 5
)

var Asset_name = map[int32]string{
	0: "ASSET_NONE",
	1: "BTC",
	2: "BCH",
	3: "ETH",
	4: "LTC",
	5: "DASH",
}
var Asset_value = map[string]int32{
	"ASSET_NONE": 0,
	"BTC":        1,
	"BCH":        2,
	"ETH":        3,
	"LTC":        4,
	"DASH":       5,
}

func (x Asset) String() string {
	return proto.EnumName(Asset_name, int32(x))
}
func (Asset) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Media is a list of possible media types. Media is a type of technology which
// is used to transport value of underlying asset.
type Media int32

const (
	Media_MEDIA_NONE Media = 0
	//
	// BLOCKCHAIN means that blockchain direct used for making the payments.
	Media_BLOCKCHAIN Media = 1
	//
	// LIGHTNING means that second layer on top of the blockchain is used for
	// making the payments.
	Media_LIGHTNING Media = 2
)

var Media_name = map[int32]string{
	0: "MEDIA_NONE",
	1: "BLOCKCHAIN",
	2: "LIGHTNING",
}
var Media_value = map[string]int32{
	"MEDIA_NONE": 0,
	"BLOCKCHAIN": 1,
	"LIGHTNING":  2,
}

func (x Media) String() string {
	return proto.EnumName(Media_name, int32(x))
}
func (Media) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// PaymentStatus denotes the stage of the processing the payment.
type PaymentStatus int32

const (
	PaymentStatus_STATUS_NONE PaymentStatus = 0
	//
	// WAITING means that payment has been created and waiting to be approved
	// for sending.
	PaymentStatus_WAITING PaymentStatus = 1
	//
	// PENDING means that service is seeing the payment, but it not yet approved
	// from the its POV.
	PaymentStatus_PENDING PaymentStatus = 2
	//
	// COMPLETED in case of outgoing/incoming payment this means that we
	// sent/received the transaction in/from the network and it was confirmed
	// number of times service believe sufficient. In case of the forward
	// transaction it means that we succesfully routed it through and
	// earned fee for that.
	PaymentStatus_COMPLETED PaymentStatus = 3
	//
	// FAILED means that services has tryied to send payment for couple of
	// times, but without success, and now service gave up.
	PaymentStatus_FAILED PaymentStatus = 4
)

var PaymentStatus_name = map[int32]string{
	0: "STATUS_NONE",
	1: "WAITING",
	2: "PENDING",
	3: "COMPLETED",
	4: "FAILED",
}
var PaymentStatus_value = map[string]int32{
	"STATUS_NONE": 0,
	"WAITING":     1,
	"PENDING":     2,
	"COMPLETED":   3,
	"FAILED":      4,
}

func (x PaymentStatus) String() string {
	return proto.EnumName(PaymentStatus_name, int32(x))
}
func (PaymentStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// PaymentType denotes the direction of the payment.
type PaymentType int32

const (
	PaymentType_TYPE_NONE PaymentType = 0
	//
	// INCOMING type of payment which service has received from someone else
	// in the media.
	PaymentType_INCOMING PaymentType = 1
	//
	// OUTGOING type of payment which service has sent to someone else in the
	// media.
	PaymentType_OUTGOING PaymentType = 2
)

var PaymentType_name = map[int32]string{
	0: "TYPE_NONE",
	1: "INCOMING",
	2: "OUTGOING",
}
var PaymentType_value = map[string]int32{
	"TYPE_NONE": 0,
	"INCOMING":  1,
	"OUTGOING":  2,
}

func (x PaymentType) String() string {
	return proto.EnumName(PaymentType_name, int32(x))
}
func (PaymentType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type EmptyRequest struct {
}

func (m *EmptyRequest) Reset()                    { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string            { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()               {}
func (*EmptyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type EmptyResponse struct {
}

func (m *EmptyResponse) Reset()                    { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string            { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()               {}
func (*EmptyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CreateReceiptRequest struct {
	//
	// Asset is an acronim of the crypto currency.
	Asset Asset `protobuf:"varint,1,opt,name=asset,enum=crpc.Asset" json:"asset,omitempty"`
	//
	// Media is a type of technology which is used to transport value of
	// underlying asset.
	Media Media `protobuf:"varint,2,opt,name=media,enum=crpc.Media" json:"media,omitempty"`
	//
	// (optional) Amount is the amount which should be received on this
	// receipt.
	Amount string `protobuf:"bytes,3,opt,name=amount" json:"amount,omitempty"`
}

func (m *CreateReceiptRequest) Reset()                    { *m = CreateReceiptRequest{} }
func (m *CreateReceiptRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateReceiptRequest) ProtoMessage()               {}
func (*CreateReceiptRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateReceiptRequest) GetAsset() Asset {
	if m != nil {
		return m.Asset
	}
	return Asset_ASSET_NONE
}

func (m *CreateReceiptRequest) GetMedia() Media {
	if m != nil {
		return m.Media
	}
	return Media_MEDIA_NONE
}

func (m *CreateReceiptRequest) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

type CreateReceiptResponse struct {
	//
	// Receipt represent either blockchains address or lightning network invoice,
	// depending on the type of the request.
	Receipt string `protobuf:"bytes,2,opt,name=receipt" json:"receipt,omitempty"`
}

func (m *CreateReceiptResponse) Reset()                    { *m = CreateReceiptResponse{} }
func (m *CreateReceiptResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateReceiptResponse) ProtoMessage()               {}
func (*CreateReceiptResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CreateReceiptResponse) GetReceipt() string {
	if m != nil {
		return m.Receipt
	}
	return ""
}

type BalanceRequest struct {
	//
	// Asset is an acronim of the crypto currency.
	Asset Asset `protobuf:"varint,1,opt,name=asset,enum=crpc.Asset" json:"asset,omitempty"`
	//
	// Media is a type of technology which is used to transport value of
	// underlying asset.
	Media Media `protobuf:"varint,2,opt,name=media,enum=crpc.Media" json:"media,omitempty"`
}

func (m *BalanceRequest) Reset()                    { *m = BalanceRequest{} }
func (m *BalanceRequest) String() string            { return proto.CompactTextString(m) }
func (*BalanceRequest) ProtoMessage()               {}
func (*BalanceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BalanceRequest) GetAsset() Asset {
	if m != nil {
		return m.Asset
	}
	return Asset_ASSET_NONE
}

func (m *BalanceRequest) GetMedia() Media {
	if m != nil {
		return m.Media
	}
	return Media_MEDIA_NONE
}

type BalanceResponse struct {
	//
	// Available is the number of funds which could be used by this account
	// to send funds to someone else within the specified media.
	Available string `protobuf:"bytes,1,opt,name=available" json:"available,omitempty"`
	//
	// Pending is the number of funds are in the state of confirmation. In
	// case of blockchain media it is the transactions which are not
	// confirmed. In case of lightning media it is funds in pending payment
	// channels.
	Pending string `protobuf:"bytes,2,opt,name=pending" json:"pending,omitempty"`
}

func (m *BalanceResponse) Reset()                    { *m = BalanceResponse{} }
func (m *BalanceResponse) String() string            { return proto.CompactTextString(m) }
func (*BalanceResponse) ProtoMessage()               {}
func (*BalanceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *BalanceResponse) GetAvailable() string {
	if m != nil {
		return m.Available
	}
	return ""
}

func (m *BalanceResponse) GetPending() string {
	if m != nil {
		return m.Pending
	}
	return ""
}

type ValidateReceiptRequest struct {
	//
	// Receipt is the blockchain address in case of blockchain media and
	// lightning network invoice in case of lightning media.
	Receipt string `protobuf:"bytes,1,opt,name=receipt" json:"receipt,omitempty"`
	//
	// Asset is an acronim of the crypto currency.
	Asset Asset `protobuf:"varint,2,opt,name=asset,enum=crpc.Asset" json:"asset,omitempty"`
	//
	// Media is a type of technology which is used to transport value of
	// underlying asset.
	Media Media `protobuf:"varint,3,opt,name=media,enum=crpc.Media" json:"media,omitempty"`
	//
	// (optional) Amount is the amount which should be received on this
	// receipt.
	Amount string `protobuf:"bytes,4,opt,name=amount" json:"amount,omitempty"`
}

func (m *ValidateReceiptRequest) Reset()                    { *m = ValidateReceiptRequest{} }
func (m *ValidateReceiptRequest) String() string            { return proto.CompactTextString(m) }
func (*ValidateReceiptRequest) ProtoMessage()               {}
func (*ValidateReceiptRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ValidateReceiptRequest) GetReceipt() string {
	if m != nil {
		return m.Receipt
	}
	return ""
}

func (m *ValidateReceiptRequest) GetAsset() Asset {
	if m != nil {
		return m.Asset
	}
	return Asset_ASSET_NONE
}

func (m *ValidateReceiptRequest) GetMedia() Media {
	if m != nil {
		return m.Media
	}
	return Media_MEDIA_NONE
}

func (m *ValidateReceiptRequest) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

type EstimateFeeRequest struct {
	//
	// Asset is an acronim of the crypto currency.
	Asset Asset `protobuf:"varint,1,opt,name=asset,enum=crpc.Asset" json:"asset,omitempty"`
	//
	// Media is a type of technology which is used to transport value of
	// underlying asset.
	Media Media `protobuf:"varint,2,opt,name=media,enum=crpc.Media" json:"media,omitempty"`
	//
	// Amount is number of money which should be given to the another entity.
	Amount string `protobuf:"bytes,3,opt,name=amount" json:"amount,omitempty"`
}

func (m *EstimateFeeRequest) Reset()                    { *m = EstimateFeeRequest{} }
func (m *EstimateFeeRequest) String() string            { return proto.CompactTextString(m) }
func (*EstimateFeeRequest) ProtoMessage()               {}
func (*EstimateFeeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *EstimateFeeRequest) GetAsset() Asset {
	if m != nil {
		return m.Asset
	}
	return Asset_ASSET_NONE
}

func (m *EstimateFeeRequest) GetMedia() Media {
	if m != nil {
		return m.Media
	}
	return Media_MEDIA_NONE
}

func (m *EstimateFeeRequest) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

type EstimateFeeResponse struct {
	//
	// MediaFee is the fee which is taken by the blockchain or lightning
	// network in order to propagate the payment.
	MediaFee string `protobuf:"bytes,1,opt,name=media_fee,json=mediaFee" json:"media_fee,omitempty"`
}

func (m *EstimateFeeResponse) Reset()                    { *m = EstimateFeeResponse{} }
func (m *EstimateFeeResponse) String() string            { return proto.CompactTextString(m) }
func (*EstimateFeeResponse) ProtoMessage()               {}
func (*EstimateFeeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *EstimateFeeResponse) GetMediaFee() string {
	if m != nil {
		return m.MediaFee
	}
	return ""
}

type SendPaymentRequest struct {
	//
	// SystemPaymentID is the payment id which was created by service itself,
	// for unified identification of the payment.
	SystemPaymentId string `protobuf:"bytes,1,opt,name=system_payment_id,json=systemPaymentId" json:"system_payment_id,omitempty"`
}

func (m *SendPaymentRequest) Reset()                    { *m = SendPaymentRequest{} }
func (m *SendPaymentRequest) String() string            { return proto.CompactTextString(m) }
func (*SendPaymentRequest) ProtoMessage()               {}
func (*SendPaymentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *SendPaymentRequest) GetSystemPaymentId() string {
	if m != nil {
		return m.SystemPaymentId
	}
	return ""
}

type PaymentByIDRequest struct {
	//
	// SystemPaymentID is the payment id which was created by service itself,
	// for unified identification of the payment.
	SystemPaymentId string `protobuf:"bytes,1,opt,name=system_payment_id,json=systemPaymentId" json:"system_payment_id,omitempty"`
}

func (m *PaymentByIDRequest) Reset()                    { *m = PaymentByIDRequest{} }
func (m *PaymentByIDRequest) String() string            { return proto.CompactTextString(m) }
func (*PaymentByIDRequest) ProtoMessage()               {}
func (*PaymentByIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *PaymentByIDRequest) GetSystemPaymentId() string {
	if m != nil {
		return m.SystemPaymentId
	}
	return ""
}

type PaymentsByReceiptRequest struct {
	//
	// Receipt represent either blockchains address or lightning
	// network invoice, depending on the type of the request.
	Receipt string `protobuf:"bytes,1,opt,name=receipt" json:"receipt,omitempty"`
}

func (m *PaymentsByReceiptRequest) Reset()                    { *m = PaymentsByReceiptRequest{} }
func (m *PaymentsByReceiptRequest) String() string            { return proto.CompactTextString(m) }
func (*PaymentsByReceiptRequest) ProtoMessage()               {}
func (*PaymentsByReceiptRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *PaymentsByReceiptRequest) GetReceipt() string {
	if m != nil {
		return m.Receipt
	}
	return ""
}

type PaymentsByReceiptResponse struct {
	Payments []*Payment `protobuf:"bytes,1,rep,name=payments" json:"payments,omitempty"`
}

func (m *PaymentsByReceiptResponse) Reset()                    { *m = PaymentsByReceiptResponse{} }
func (m *PaymentsByReceiptResponse) String() string            { return proto.CompactTextString(m) }
func (*PaymentsByReceiptResponse) ProtoMessage()               {}
func (*PaymentsByReceiptResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *PaymentsByReceiptResponse) GetPayments() []*Payment {
	if m != nil {
		return m.Payments
	}
	return nil
}

type ListPaymentsRequest struct {
	//
	// (optional) Status denotes the stage of the processing the payment.
	Status PaymentStatus `protobuf:"varint,1,opt,name=status,enum=crpc.PaymentStatus" json:"status,omitempty"`
	//
	// (optional) Type denotes the direction of the payment.
	Type PaymentType `protobuf:"varint,2,opt,name=type,enum=crpc.PaymentType" json:"type,omitempty"`
	//
	// (optional) Asset is an acronim of the crypto currency.
	Asset Asset `protobuf:"varint,3,opt,name=asset,enum=crpc.Asset" json:"asset,omitempty"`
	//
	// (optional) Media is a type of technology which is used to transport
	// value of underlying asset.
	Media Media `protobuf:"varint,4,opt,name=media,enum=crpc.Media" json:"media,omitempty"`
}

func (m *ListPaymentsRequest) Reset()                    { *m = ListPaymentsRequest{} }
func (m *ListPaymentsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListPaymentsRequest) ProtoMessage()               {}
func (*ListPaymentsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *ListPaymentsRequest) GetStatus() PaymentStatus {
	if m != nil {
		return m.Status
	}
	return PaymentStatus_STATUS_NONE
}

func (m *ListPaymentsRequest) GetType() PaymentType {
	if m != nil {
		return m.Type
	}
	return PaymentType_TYPE_NONE
}

func (m *ListPaymentsRequest) GetAsset() Asset {
	if m != nil {
		return m.Asset
	}
	return Asset_ASSET_NONE
}

func (m *ListPaymentsRequest) GetMedia() Media {
	if m != nil {
		return m.Media
	}
	return Media_MEDIA_NONE
}

type ListPaymentsResponse struct {
	Payments []*Payment `protobuf:"bytes,1,rep,name=payments" json:"payments,omitempty"`
}

func (m *ListPaymentsResponse) Reset()                    { *m = ListPaymentsResponse{} }
func (m *ListPaymentsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListPaymentsResponse) ProtoMessage()               {}
func (*ListPaymentsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *ListPaymentsResponse) GetPayments() []*Payment {
	if m != nil {
		return m.Payments
	}
	return nil
}

type Payment struct {
	//
	// SystemPaymentID is the payment id which was created by service itself,
	// for unified identification of the payment.
	SystemPaymentId string `protobuf:"bytes,1,opt,name=system_payment_id,json=systemPaymentId" json:"system_payment_id,omitempty"`
	//
	// MediaPaymentID in case of blockchain media payment id is the
	// transaction id, in case of lightning media it is the payment hash.
	MediaPaymentId string `protobuf:"bytes,2,opt,name=media_payment_id,json=mediaPaymentId" json:"media_payment_id,omitempty"`
	//
	// Status denotes the stage of the processing the payment.
	Status PaymentStatus `protobuf:"varint,3,opt,name=status,enum=crpc.PaymentStatus" json:"status,omitempty"`
	//
	// Type denotes the direction of the payment.
	Type PaymentType `protobuf:"varint,4,opt,name=type,enum=crpc.PaymentType" json:"type,omitempty"`
	//
	// Asset is an acronim of the crypto currency.
	Asset Asset `protobuf:"varint,5,opt,name=asset,enum=crpc.Asset" json:"asset,omitempty"`
	//
	// Media is a type of technology which is used to transport value of
	// underlying asset.
	Media Media `protobuf:"varint,6,opt,name=media,enum=crpc.Media" json:"media,omitempty"`
	//
	// Amount is the number of funds which receiver gets at the end.
	Amount int64 `protobuf:"varint,7,opt,name=amount" json:"amount,omitempty"`
	//
	// MediaFee is the fee which is taken by the blockchain or lightning
	// network in order to propagate the payment.
	MediaFee string `protobuf:"bytes,8,opt,name=media_fee,json=mediaFee" json:"media_fee,omitempty"`
}

func (m *Payment) Reset()                    { *m = Payment{} }
func (m *Payment) String() string            { return proto.CompactTextString(m) }
func (*Payment) ProtoMessage()               {}
func (*Payment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *Payment) GetSystemPaymentId() string {
	if m != nil {
		return m.SystemPaymentId
	}
	return ""
}

func (m *Payment) GetMediaPaymentId() string {
	if m != nil {
		return m.MediaPaymentId
	}
	return ""
}

func (m *Payment) GetStatus() PaymentStatus {
	if m != nil {
		return m.Status
	}
	return PaymentStatus_STATUS_NONE
}

func (m *Payment) GetType() PaymentType {
	if m != nil {
		return m.Type
	}
	return PaymentType_TYPE_NONE
}

func (m *Payment) GetAsset() Asset {
	if m != nil {
		return m.Asset
	}
	return Asset_ASSET_NONE
}

func (m *Payment) GetMedia() Media {
	if m != nil {
		return m.Media
	}
	return Media_MEDIA_NONE
}

func (m *Payment) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Payment) GetMediaFee() string {
	if m != nil {
		return m.MediaFee
	}
	return ""
}

func init() {
	proto.RegisterType((*EmptyRequest)(nil), "crpc.EmptyRequest")
	proto.RegisterType((*EmptyResponse)(nil), "crpc.EmptyResponse")
	proto.RegisterType((*CreateReceiptRequest)(nil), "crpc.CreateReceiptRequest")
	proto.RegisterType((*CreateReceiptResponse)(nil), "crpc.CreateReceiptResponse")
	proto.RegisterType((*BalanceRequest)(nil), "crpc.BalanceRequest")
	proto.RegisterType((*BalanceResponse)(nil), "crpc.BalanceResponse")
	proto.RegisterType((*ValidateReceiptRequest)(nil), "crpc.ValidateReceiptRequest")
	proto.RegisterType((*EstimateFeeRequest)(nil), "crpc.EstimateFeeRequest")
	proto.RegisterType((*EstimateFeeResponse)(nil), "crpc.EstimateFeeResponse")
	proto.RegisterType((*SendPaymentRequest)(nil), "crpc.SendPaymentRequest")
	proto.RegisterType((*PaymentByIDRequest)(nil), "crpc.PaymentByIDRequest")
	proto.RegisterType((*PaymentsByReceiptRequest)(nil), "crpc.PaymentsByReceiptRequest")
	proto.RegisterType((*PaymentsByReceiptResponse)(nil), "crpc.PaymentsByReceiptResponse")
	proto.RegisterType((*ListPaymentsRequest)(nil), "crpc.ListPaymentsRequest")
	proto.RegisterType((*ListPaymentsResponse)(nil), "crpc.ListPaymentsResponse")
	proto.RegisterType((*Payment)(nil), "crpc.Payment")
	proto.RegisterEnum("crpc.Asset", Asset_name, Asset_value)
	proto.RegisterEnum("crpc.Media", Media_name, Media_value)
	proto.RegisterEnum("crpc.PaymentStatus", PaymentStatus_name, PaymentStatus_value)
	proto.RegisterEnum("crpc.PaymentType", PaymentType_name, PaymentType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PayServer service

type PayServerClient interface {
	//
	// CreateReceipt is used to create blockchain deposit address in
	// case of blockchain media, and lightning network invoice in
	// case of the lightning media, which will be used to receive money from
	// external entity.
	CreateReceipt(ctx context.Context, in *CreateReceiptRequest, opts ...grpc.CallOption) (*CreateReceiptResponse, error)
	//
	// ValidateReceipt is used to validate receipt for given asset and media.
	ValidateReceipt(ctx context.Context, in *ValidateReceiptRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	//
	// Balance is used to determine balance.
	Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
	//
	// EstimateFee estimates the fee of the payment.
	EstimateFee(ctx context.Context, in *EstimateFeeRequest, opts ...grpc.CallOption) (*EstimateFeeResponse, error)
	//
	// SendPayment sends payment to the given recipient,
	// ensures in the validity of the receipt as well as the
	// account has enough money for doing that.
	SendPayment(ctx context.Context, in *SendPaymentRequest, opts ...grpc.CallOption) (*Payment, error)
	//
	// PaymentByID is used to fetch the information about payment, by the
	// given system payment id.
	PaymentByID(ctx context.Context, in *PaymentByIDRequest, opts ...grpc.CallOption) (*Payment, error)
	//
	// PaymentsByReceipt is used to fetch the information about payment, by the
	// given receipt.
	PaymentsByReceipt(ctx context.Context, in *PaymentsByReceiptRequest, opts ...grpc.CallOption) (*PaymentsByReceiptResponse, error)
	//
	// ListPayments returnes list of payment which were registered by the
	// system.
	ListPayments(ctx context.Context, in *ListPaymentsRequest, opts ...grpc.CallOption) (*ListPaymentsResponse, error)
}

type payServerClient struct {
	cc *grpc.ClientConn
}

func NewPayServerClient(cc *grpc.ClientConn) PayServerClient {
	return &payServerClient{cc}
}

func (c *payServerClient) CreateReceipt(ctx context.Context, in *CreateReceiptRequest, opts ...grpc.CallOption) (*CreateReceiptResponse, error) {
	out := new(CreateReceiptResponse)
	err := grpc.Invoke(ctx, "/crpc.PayServer/CreateReceipt", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payServerClient) ValidateReceipt(ctx context.Context, in *ValidateReceiptRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := grpc.Invoke(ctx, "/crpc.PayServer/ValidateReceipt", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payServerClient) Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	out := new(BalanceResponse)
	err := grpc.Invoke(ctx, "/crpc.PayServer/Balance", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payServerClient) EstimateFee(ctx context.Context, in *EstimateFeeRequest, opts ...grpc.CallOption) (*EstimateFeeResponse, error) {
	out := new(EstimateFeeResponse)
	err := grpc.Invoke(ctx, "/crpc.PayServer/EstimateFee", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payServerClient) SendPayment(ctx context.Context, in *SendPaymentRequest, opts ...grpc.CallOption) (*Payment, error) {
	out := new(Payment)
	err := grpc.Invoke(ctx, "/crpc.PayServer/SendPayment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payServerClient) PaymentByID(ctx context.Context, in *PaymentByIDRequest, opts ...grpc.CallOption) (*Payment, error) {
	out := new(Payment)
	err := grpc.Invoke(ctx, "/crpc.PayServer/PaymentByID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payServerClient) PaymentsByReceipt(ctx context.Context, in *PaymentsByReceiptRequest, opts ...grpc.CallOption) (*PaymentsByReceiptResponse, error) {
	out := new(PaymentsByReceiptResponse)
	err := grpc.Invoke(ctx, "/crpc.PayServer/PaymentsByReceipt", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payServerClient) ListPayments(ctx context.Context, in *ListPaymentsRequest, opts ...grpc.CallOption) (*ListPaymentsResponse, error) {
	out := new(ListPaymentsResponse)
	err := grpc.Invoke(ctx, "/crpc.PayServer/ListPayments", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PayServer service

type PayServerServer interface {
	//
	// CreateReceipt is used to create blockchain deposit address in
	// case of blockchain media, and lightning network invoice in
	// case of the lightning media, which will be used to receive money from
	// external entity.
	CreateReceipt(context.Context, *CreateReceiptRequest) (*CreateReceiptResponse, error)
	//
	// ValidateReceipt is used to validate receipt for given asset and media.
	ValidateReceipt(context.Context, *ValidateReceiptRequest) (*EmptyResponse, error)
	//
	// Balance is used to determine balance.
	Balance(context.Context, *BalanceRequest) (*BalanceResponse, error)
	//
	// EstimateFee estimates the fee of the payment.
	EstimateFee(context.Context, *EstimateFeeRequest) (*EstimateFeeResponse, error)
	//
	// SendPayment sends payment to the given recipient,
	// ensures in the validity of the receipt as well as the
	// account has enough money for doing that.
	SendPayment(context.Context, *SendPaymentRequest) (*Payment, error)
	//
	// PaymentByID is used to fetch the information about payment, by the
	// given system payment id.
	PaymentByID(context.Context, *PaymentByIDRequest) (*Payment, error)
	//
	// PaymentsByReceipt is used to fetch the information about payment, by the
	// given receipt.
	PaymentsByReceipt(context.Context, *PaymentsByReceiptRequest) (*PaymentsByReceiptResponse, error)
	//
	// ListPayments returnes list of payment which were registered by the
	// system.
	ListPayments(context.Context, *ListPaymentsRequest) (*ListPaymentsResponse, error)
}

func RegisterPayServerServer(s *grpc.Server, srv PayServerServer) {
	s.RegisterService(&_PayServer_serviceDesc, srv)
}

func _PayServer_CreateReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReceiptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServerServer).CreateReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crpc.PayServer/CreateReceipt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServerServer).CreateReceipt(ctx, req.(*CreateReceiptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayServer_ValidateReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateReceiptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServerServer).ValidateReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crpc.PayServer/ValidateReceipt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServerServer).ValidateReceipt(ctx, req.(*ValidateReceiptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayServer_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServerServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crpc.PayServer/Balance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServerServer).Balance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayServer_EstimateFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EstimateFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServerServer).EstimateFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crpc.PayServer/EstimateFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServerServer).EstimateFee(ctx, req.(*EstimateFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayServer_SendPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServerServer).SendPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crpc.PayServer/SendPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServerServer).SendPayment(ctx, req.(*SendPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayServer_PaymentByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServerServer).PaymentByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crpc.PayServer/PaymentByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServerServer).PaymentByID(ctx, req.(*PaymentByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayServer_PaymentsByReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentsByReceiptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServerServer).PaymentsByReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crpc.PayServer/PaymentsByReceipt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServerServer).PaymentsByReceipt(ctx, req.(*PaymentsByReceiptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayServer_ListPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPaymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServerServer).ListPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crpc.PayServer/ListPayments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServerServer).ListPayments(ctx, req.(*ListPaymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PayServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crpc.PayServer",
	HandlerType: (*PayServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReceipt",
			Handler:    _PayServer_CreateReceipt_Handler,
		},
		{
			MethodName: "ValidateReceipt",
			Handler:    _PayServer_ValidateReceipt_Handler,
		},
		{
			MethodName: "Balance",
			Handler:    _PayServer_Balance_Handler,
		},
		{
			MethodName: "EstimateFee",
			Handler:    _PayServer_EstimateFee_Handler,
		},
		{
			MethodName: "SendPayment",
			Handler:    _PayServer_SendPayment_Handler,
		},
		{
			MethodName: "PaymentByID",
			Handler:    _PayServer_PaymentByID_Handler,
		},
		{
			MethodName: "PaymentsByReceipt",
			Handler:    _PayServer_PaymentsByReceipt_Handler,
		},
		{
			MethodName: "ListPayments",
			Handler:    _PayServer_ListPayments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 802 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x5b, 0x6f, 0x9b, 0x48,
	0x14, 0x5e, 0x0c, 0xbe, 0x70, 0x7c, 0x23, 0x93, 0x8b, 0x1c, 0x12, 0xed, 0x7a, 0x91, 0x56, 0xf2,
	0x7a, 0xa5, 0x48, 0xeb, 0x56, 0x51, 0x1f, 0x8b, 0x31, 0x8e, 0x51, 0x7d, 0x13, 0x90, 0x54, 0x7d,
	0xb2, 0x88, 0x3d, 0xad, 0x90, 0x7c, 0xa1, 0x30, 0x89, 0xc4, 0xbf, 0xe8, 0x8f, 0xe8, 0x73, 0xfb,
	0x17, 0x2b, 0x60, 0xb0, 0xc1, 0x76, 0x14, 0xa7, 0xaa, 0xfa, 0xc6, 0xb9, 0x7d, 0x73, 0xce, 0x77,
	0xe6, 0x9c, 0x01, 0x78, 0xd7, 0x99, 0x5e, 0x39, 0xee, 0x8a, 0xac, 0x10, 0x37, 0x75, 0x9d, 0xa9,
	0x54, 0x81, 0x92, 0xba, 0x70, 0x88, 0xaf, 0xe3, 0xcf, 0x0f, 0xd8, 0x23, 0x52, 0x15, 0xca, 0x54,
	0xf6, 0x9c, 0xd5, 0xd2, 0xc3, 0x12, 0x81, 0x13, 0xc5, 0xc5, 0x16, 0xc1, 0x3a, 0x9e, 0x62, 0xdb,
	0x21, 0xd4, 0x11, 0xfd, 0x0d, 0x59, 0xcb, 0xf3, 0x30, 0xa9, 0x31, 0x75, 0xa6, 0x51, 0x69, 0x15,
	0xaf, 0x02, 0xb8, 0x2b, 0x39, 0x50, 0xe9, 0x91, 0x25, 0x70, 0x59, 0xe0, 0x99, 0x6d, 0xd5, 0x32,
	0x49, 0x97, 0x41, 0xa0, 0xd2, 0x23, 0x0b, 0x3a, 0x83, 0x9c, 0xb5, 0x58, 0x3d, 0x2c, 0x49, 0x8d,
	0xad, 0x33, 0x0d, 0x5e, 0xa7, 0x92, 0xf4, 0x3f, 0x9c, 0x6e, 0x9d, 0x1a, 0xa5, 0x83, 0x6a, 0x90,
	0x77, 0x23, 0x55, 0x88, 0xca, 0xeb, 0xb1, 0x28, 0xdd, 0x41, 0xa5, 0x6d, 0xcd, 0xad, 0xe5, 0x14,
	0xff, 0xd2, 0x14, 0x25, 0x0d, 0xaa, 0x6b, 0x5c, 0x9a, 0xc4, 0x25, 0xf0, 0xd6, 0xa3, 0x65, 0xcf,
	0xad, 0xfb, 0x39, 0x0e, 0xc1, 0x79, 0x7d, 0xa3, 0x08, 0x52, 0x74, 0xf0, 0x72, 0x66, 0x2f, 0x3f,
	0xc5, 0x29, 0x52, 0x51, 0xfa, 0xc2, 0xc0, 0xd9, 0x9d, 0x35, 0xb7, 0x67, 0xbb, 0x74, 0x26, 0xea,
	0x62, 0x52, 0x75, 0x6d, 0xaa, 0xc8, 0x3c, 0x5f, 0x05, 0x7b, 0x00, 0xd1, 0x5c, 0x8a, 0x68, 0x17,
	0x90, 0xea, 0x11, 0x7b, 0x61, 0x11, 0xdc, 0xc5, 0xf8, 0xf7, 0x34, 0xb7, 0x05, 0xc7, 0xa9, 0x33,
	0x29, 0xab, 0x17, 0xc0, 0x87, 0x71, 0x93, 0x8f, 0x38, 0x66, 0xb5, 0x10, 0x2a, 0xba, 0x18, 0x4b,
	0x6f, 0x01, 0x19, 0x78, 0x39, 0x1b, 0x5b, 0xfe, 0x02, 0x2f, 0xd7, 0xac, 0x35, 0xe1, 0xc8, 0xf3,
	0x3d, 0x82, 0x17, 0x13, 0x27, 0x32, 0x4c, 0xec, 0x19, 0x0d, 0xad, 0x46, 0x06, 0x1a, 0xa0, 0xcd,
	0x02, 0x04, 0x2a, 0xb4, 0x7d, 0xad, 0xf3, 0x33, 0x08, 0xaf, 0xa1, 0x46, 0x05, 0xaf, 0xed, 0x1f,
	0xda, 0x3f, 0xa9, 0x0b, 0xe7, 0x7b, 0xa2, 0x68, 0xcd, 0xff, 0x42, 0x81, 0x9e, 0xeb, 0xd5, 0x98,
	0x3a, 0xdb, 0x28, 0xb6, 0xca, 0x11, 0x91, 0x71, 0xa1, 0x6b, 0xb3, 0xf4, 0x9d, 0x81, 0xe3, 0xbe,
	0xed, 0x91, 0x18, 0x2c, 0x3e, 0xf9, 0x3f, 0xc8, 0x79, 0xc4, 0x22, 0x0f, 0x1e, 0x6d, 0xd6, 0x71,
	0x0a, 0xc0, 0x08, 0x4d, 0x3a, 0x75, 0x41, 0xff, 0x00, 0x47, 0x7c, 0x07, 0xd3, 0xa6, 0x1d, 0xa5,
	0x5c, 0x4d, 0xdf, 0xc1, 0x7a, 0x68, 0xde, 0xf4, 0x9f, 0x7d, 0xbe, 0xff, 0xdc, 0x93, 0x93, 0x23,
	0xc3, 0x49, 0x3a, 0xe1, 0x97, 0x17, 0xfd, 0x2d, 0x03, 0x79, 0xaa, 0x7d, 0x49, 0xab, 0x50, 0x03,
	0x84, 0xe8, 0x2e, 0x25, 0x5c, 0xa3, 0x61, 0xac, 0x84, 0xfa, 0x8d, 0xe7, 0x86, 0x3e, 0xf6, 0x70,
	0xfa, 0xb8, 0x03, 0xe9, 0xcb, 0x3e, 0x4f, 0x5f, 0xee, 0x80, 0xf1, 0xc9, 0xd7, 0x99, 0x06, 0x1b,
	0x8f, 0x4f, 0x7a, 0x4e, 0x0a, 0xe9, 0x39, 0x69, 0xaa, 0x90, 0x0d, 0xcf, 0x41, 0x15, 0x00, 0xd9,
	0x30, 0x54, 0x73, 0x32, 0x1c, 0x0d, 0x55, 0xe1, 0x0f, 0x94, 0x07, 0xb6, 0x6d, 0x2a, 0x02, 0x13,
	0x7e, 0x28, 0x3d, 0x21, 0x13, 0x7c, 0xa8, 0x66, 0x4f, 0x60, 0x83, 0x8f, 0xbe, 0xa9, 0x08, 0x1c,
	0x2a, 0x00, 0xd7, 0x91, 0x8d, 0x9e, 0x90, 0x6d, 0x5e, 0x43, 0x36, 0xcc, 0x25, 0x80, 0x19, 0xa8,
	0x1d, 0x4d, 0x8e, 0x61, 0x2a, 0x00, 0xed, 0xfe, 0x48, 0x79, 0xa7, 0xf4, 0x64, 0x6d, 0x28, 0x30,
	0xa8, 0x0c, 0x7c, 0x5f, 0xbb, 0xe9, 0x99, 0x43, 0x6d, 0x78, 0x23, 0x64, 0x9a, 0xb7, 0x50, 0x4e,
	0x31, 0x87, 0xaa, 0x50, 0x34, 0x4c, 0xd9, 0xbc, 0x35, 0x62, 0x80, 0x22, 0xe4, 0xdf, 0xcb, 0x9a,
	0x19, 0xb8, 0x33, 0x81, 0x30, 0x56, 0x87, 0x9d, 0x30, 0x36, 0x80, 0x52, 0x46, 0x83, 0x71, 0x5f,
	0x35, 0xd5, 0x8e, 0xc0, 0x22, 0x80, 0x5c, 0x57, 0xd6, 0xfa, 0x6a, 0x47, 0xe0, 0x9a, 0x6f, 0xa0,
	0x98, 0x60, 0x39, 0xf0, 0x34, 0x3f, 0x8c, 0xd5, 0x18, 0xb2, 0x04, 0x05, 0x6d, 0xa8, 0x8c, 0x06,
	0x11, 0x66, 0x09, 0x0a, 0xa3, 0x5b, 0xf3, 0x66, 0x14, 0x82, 0xb6, 0xbe, 0x72, 0xc0, 0x8f, 0x2d,
	0xdf, 0xc0, 0xee, 0x23, 0x76, 0x51, 0x0f, 0xca, 0xa9, 0x67, 0x05, 0x89, 0x11, 0xef, 0xfb, 0x5e,
	0x38, 0xf1, 0x62, 0xaf, 0x8d, 0xde, 0xe1, 0x0e, 0x54, 0xb7, 0x36, 0x39, 0xba, 0x8c, 0xfc, 0xf7,
	0x2f, 0x78, 0x91, 0xde, 0xab, 0xd4, 0xe3, 0x8a, 0xae, 0x21, 0x4f, 0xdf, 0x16, 0x74, 0x12, 0xd9,
	0xd3, 0x4f, 0x98, 0x78, 0xba, 0xa5, 0xa5, 0x71, 0x6d, 0x28, 0x26, 0x36, 0x28, 0xaa, 0x51, 0xec,
	0x9d, 0x45, 0x2e, 0x9e, 0xef, 0xb1, 0xac, 0xcf, 0x2e, 0x26, 0x36, 0x6a, 0x8c, 0xb1, 0xbb, 0x64,
	0xc5, 0xf4, 0x70, 0x06, 0x71, 0x89, 0x3d, 0x1a, 0xc7, 0xed, 0xae, 0xd6, 0xed, 0x38, 0x13, 0x8e,
	0x76, 0xf6, 0x20, 0xfa, 0x33, 0xe5, 0xb3, 0xb3, 0x56, 0xc5, 0xbf, 0x9e, 0xb4, 0xd3, 0x2a, 0x54,
	0x28, 0x25, 0x77, 0x0c, 0xa2, 0x05, 0xef, 0x59, 0x94, 0xa2, 0xb8, 0xcf, 0x14, 0xc1, 0xdc, 0xe7,
	0xc2, 0x7f, 0xa2, 0x57, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x24, 0xc2, 0xbf, 0x20, 0x09,
	0x00, 0x00,
}
