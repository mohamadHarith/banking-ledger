// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.2
// source: shared/proto/transaction_processor_proto/transaction_processor_service.proto

package transaction_processor_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateAccountRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	UserId         *string                `protobuf:"bytes,1,req,name=userId" json:"userId,omitempty"`
	InitialBalance *uint32                `protobuf:"varint,2,req,name=initialBalance" json:"initialBalance,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	mi := &file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccountRequest) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

func (x *CreateAccountRequest) GetInitialBalance() uint32 {
	if x != nil && x.InitialBalance != nil {
		return *x.InitialBalance
	}
	return 0
}

type CreateAccountResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Account       *Account               `protobuf:"bytes,1,req,name=account" json:"account,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAccountResponse) Reset() {
	*x = CreateAccountResponse{}
	mi := &file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountResponse) ProtoMessage() {}

func (x *CreateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountResponse.ProtoReflect.Descriptor instead.
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccountResponse) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

type WithdrawRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Amount        *uint32                `protobuf:"varint,1,req,name=amount" json:"amount,omitempty"`
	UserId        *string                `protobuf:"bytes,2,req,name=userId" json:"userId,omitempty"`
	AccountId     *string                `protobuf:"bytes,3,req,name=accountId" json:"accountId,omitempty"`
	Description   *string                `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WithdrawRequest) Reset() {
	*x = WithdrawRequest{}
	mi := &file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WithdrawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawRequest) ProtoMessage() {}

func (x *WithdrawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawRequest.ProtoReflect.Descriptor instead.
func (*WithdrawRequest) Descriptor() ([]byte, []int) {
	return file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_rawDescGZIP(), []int{2}
}

func (x *WithdrawRequest) GetAmount() uint32 {
	if x != nil && x.Amount != nil {
		return *x.Amount
	}
	return 0
}

func (x *WithdrawRequest) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

func (x *WithdrawRequest) GetAccountId() string {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return ""
}

func (x *WithdrawRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

type DepositRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Amount        *uint32                `protobuf:"varint,1,req,name=amount" json:"amount,omitempty"`
	UserId        *string                `protobuf:"bytes,2,req,name=userId" json:"userId,omitempty"`
	AccountId     *string                `protobuf:"bytes,3,req,name=accountId" json:"accountId,omitempty"`
	Description   *string                `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DepositRequest) Reset() {
	*x = DepositRequest{}
	mi := &file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DepositRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositRequest) ProtoMessage() {}

func (x *DepositRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositRequest.ProtoReflect.Descriptor instead.
func (*DepositRequest) Descriptor() ([]byte, []int) {
	return file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_rawDescGZIP(), []int{3}
}

func (x *DepositRequest) GetAmount() uint32 {
	if x != nil && x.Amount != nil {
		return *x.Amount
	}
	return 0
}

func (x *DepositRequest) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

func (x *DepositRequest) GetAccountId() string {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return ""
}

func (x *DepositRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

type TransferRequest struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	Amount               *uint32                `protobuf:"varint,1,req,name=amount" json:"amount,omitempty"`
	OriginUserId         *string                `protobuf:"bytes,2,req,name=originUserId" json:"originUserId,omitempty"`
	OriginAccountId      *string                `protobuf:"bytes,3,req,name=originAccountId" json:"originAccountId,omitempty"`
	DestinationUserId    *string                `protobuf:"bytes,4,req,name=destinationUserId" json:"destinationUserId,omitempty"`
	DestinationAccountId *string                `protobuf:"bytes,5,req,name=destinationAccountId" json:"destinationAccountId,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *TransferRequest) Reset() {
	*x = TransferRequest{}
	mi := &file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferRequest) ProtoMessage() {}

func (x *TransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferRequest.ProtoReflect.Descriptor instead.
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_rawDescGZIP(), []int{4}
}

func (x *TransferRequest) GetAmount() uint32 {
	if x != nil && x.Amount != nil {
		return *x.Amount
	}
	return 0
}

func (x *TransferRequest) GetOriginUserId() string {
	if x != nil && x.OriginUserId != nil {
		return *x.OriginUserId
	}
	return ""
}

func (x *TransferRequest) GetOriginAccountId() string {
	if x != nil && x.OriginAccountId != nil {
		return *x.OriginAccountId
	}
	return ""
}

func (x *TransferRequest) GetDestinationUserId() string {
	if x != nil && x.DestinationUserId != nil {
		return *x.DestinationUserId
	}
	return ""
}

func (x *TransferRequest) GetDestinationAccountId() string {
	if x != nil && x.DestinationAccountId != nil {
		return *x.DestinationAccountId
	}
	return ""
}

type Account struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *string                `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	UserId        *string                `protobuf:"bytes,2,req,name=userId" json:"userId,omitempty"`
	Balance       *uint32                `protobuf:"varint,3,req,name=balance" json:"balance,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,req,name=createdAt" json:"createdAt,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,req,name=updatedAt" json:"updatedAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Account) Reset() {
	*x = Account{}
	mi := &file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_rawDescGZIP(), []int{5}
}

func (x *Account) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Account) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

func (x *Account) GetBalance() uint32 {
	if x != nil && x.Balance != nil {
		return *x.Balance
	}
	return 0
}

func (x *Account) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Account) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_shared_proto_transaction_processor_proto_transaction_processor_service_proto protoreflect.FileDescriptor

var file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x0e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x22, 0x3b, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x81, 0x01, 0x0a, 0x0f, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x80, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd9, 0x01, 0x0a, 0x0f, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x0f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x02, 0x28, 0x09, 0x52, 0x11, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x32, 0x0a, 0x14, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20, 0x02, 0x28, 0x09, 0x52, 0x14,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0xbf, 0x01, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x04, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x85, 0x02, 0x0a, 0x1b, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x08, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x12, 0x10, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x34, 0x0a, 0x07, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x12, 0x0f, 0x2e, 0x44, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x12, 0x10, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x51,
	0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x68,
	0x61, 0x6d, 0x61, 0x64, 0x48, 0x61, 0x72, 0x69, 0x74, 0x68, 0x2f, 0x62, 0x61, 0x6b, 0x69, 0x6e,
	0x67, 0x2d, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f,
}

var (
	file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_rawDescOnce sync.Once
	file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_rawDescData = file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_rawDesc
)

func file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_rawDescGZIP() []byte {
	file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_rawDescOnce.Do(func() {
		file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_rawDescData)
	})
	return file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_rawDescData
}

var file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_goTypes = []any{
	(*CreateAccountRequest)(nil),  // 0: CreateAccountRequest
	(*CreateAccountResponse)(nil), // 1: CreateAccountResponse
	(*WithdrawRequest)(nil),       // 2: WithdrawRequest
	(*DepositRequest)(nil),        // 3: DepositRequest
	(*TransferRequest)(nil),       // 4: TransferRequest
	(*Account)(nil),               // 5: Account
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_depIdxs = []int32{
	5, // 0: CreateAccountResponse.account:type_name -> Account
	6, // 1: Account.createdAt:type_name -> google.protobuf.Timestamp
	6, // 2: Account.updatedAt:type_name -> google.protobuf.Timestamp
	0, // 3: TransactionProcessorService.CreateAccount:input_type -> CreateAccountRequest
	2, // 4: TransactionProcessorService.Withdraw:input_type -> WithdrawRequest
	3, // 5: TransactionProcessorService.Deposit:input_type -> DepositRequest
	4, // 6: TransactionProcessorService.Transfer:input_type -> TransferRequest
	1, // 7: TransactionProcessorService.CreateAccount:output_type -> CreateAccountResponse
	7, // 8: TransactionProcessorService.Withdraw:output_type -> google.protobuf.Empty
	7, // 9: TransactionProcessorService.Deposit:output_type -> google.protobuf.Empty
	7, // 10: TransactionProcessorService.Transfer:output_type -> google.protobuf.Empty
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_init() }
func file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_init() {
	if File_shared_proto_transaction_processor_proto_transaction_processor_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_goTypes,
		DependencyIndexes: file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_depIdxs,
		MessageInfos:      file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_msgTypes,
	}.Build()
	File_shared_proto_transaction_processor_proto_transaction_processor_service_proto = out.File
	file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_rawDesc = nil
	file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_goTypes = nil
	file_shared_proto_transaction_processor_proto_transaction_processor_service_proto_depIdxs = nil
}
