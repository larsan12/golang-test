// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: checkout-service.proto

package checkout

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddToCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  int64  `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	Sku   uint32 `protobuf:"varint,2,opt,name=sku,proto3" json:"sku,omitempty"`
	Count uint32 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *AddToCartRequest) Reset() {
	*x = AddToCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddToCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartRequest) ProtoMessage() {}

func (x *AddToCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartRequest.ProtoReflect.Descriptor instead.
func (*AddToCartRequest) Descriptor() ([]byte, []int) {
	return file_checkout_service_proto_rawDescGZIP(), []int{0}
}

func (x *AddToCartRequest) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *AddToCartRequest) GetSku() uint32 {
	if x != nil {
		return x.Sku
	}
	return 0
}

func (x *AddToCartRequest) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type DeleteFromCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  int64  `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	Sku   uint32 `protobuf:"varint,2,opt,name=sku,proto3" json:"sku,omitempty"`
	Count uint32 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *DeleteFromCartRequest) Reset() {
	*x = DeleteFromCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFromCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFromCartRequest) ProtoMessage() {}

func (x *DeleteFromCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFromCartRequest.ProtoReflect.Descriptor instead.
func (*DeleteFromCartRequest) Descriptor() ([]byte, []int) {
	return file_checkout_service_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteFromCartRequest) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *DeleteFromCartRequest) GetSku() uint32 {
	if x != nil {
		return x.Sku
	}
	return 0
}

func (x *DeleteFromCartRequest) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type AddToCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddToCartResponse) Reset() {
	*x = AddToCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddToCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartResponse) ProtoMessage() {}

func (x *AddToCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartResponse.ProtoReflect.Descriptor instead.
func (*AddToCartResponse) Descriptor() ([]byte, []int) {
	return file_checkout_service_proto_rawDescGZIP(), []int{2}
}

func (x *AddToCartResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items      []*CartItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	TotalPrice uint32      `protobuf:"varint,2,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
}

func (x *ListCartResponse) Reset() {
	*x = ListCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCartResponse) ProtoMessage() {}

func (x *ListCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCartResponse.ProtoReflect.Descriptor instead.
func (*ListCartResponse) Descriptor() ([]byte, []int) {
	return file_checkout_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListCartResponse) GetItems() []*CartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListCartResponse) GetTotalPrice() uint32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

type ListCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User int64 `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ListCartRequest) Reset() {
	*x = ListCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCartRequest) ProtoMessage() {}

func (x *ListCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCartRequest.ProtoReflect.Descriptor instead.
func (*ListCartRequest) Descriptor() ([]byte, []int) {
	return file_checkout_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListCartRequest) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku   uint32 `protobuf:"varint,1,opt,name=sku,proto3" json:"sku,omitempty"`
	Count uint32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Price uint32 `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_checkout_service_proto_rawDescGZIP(), []int{5}
}

func (x *CartItem) GetSku() uint32 {
	if x != nil {
		return x.Sku
	}
	return 0
}

func (x *CartItem) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CartItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CartItem) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type PuchaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User int64 `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *PuchaseRequest) Reset() {
	*x = PuchaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PuchaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PuchaseRequest) ProtoMessage() {}

func (x *PuchaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PuchaseRequest.ProtoReflect.Descriptor instead.
func (*PuchaseRequest) Descriptor() ([]byte, []int) {
	return file_checkout_service_proto_rawDescGZIP(), []int{6}
}

func (x *PuchaseRequest) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

var File_checkout_service_proto protoreflect.FileDescriptor

var file_checkout_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f,
	0x75, 0x74, 0x5f, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b,
	0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x53, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x6b,
	0x75, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x54, 0x6f,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x60, 0x0a, 0x10,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x61,
	0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x25,
	0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x5c, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x73, 0x6b, 0x75, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x22, 0x24, 0x0a, 0x0e, 0x50, 0x75, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0xa7, 0x02, 0x0a, 0x0a, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x56, 0x31, 0x12, 0x42, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x54,
	0x6f, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1d, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74,
	0x5f, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4c, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x12, 0x22,
	0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x47, 0x0a, 0x08, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1c, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75,
	0x74, 0x5f, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x5f,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x50, 0x75, 0x63, 0x68, 0x61, 0x73, 0x65, 0x12, 0x1b,
	0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x63,
	0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x13, 0x5a, 0x11, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_checkout_service_proto_rawDescOnce sync.Once
	file_checkout_service_proto_rawDescData = file_checkout_service_proto_rawDesc
)

func file_checkout_service_proto_rawDescGZIP() []byte {
	file_checkout_service_proto_rawDescOnce.Do(func() {
		file_checkout_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_checkout_service_proto_rawDescData)
	})
	return file_checkout_service_proto_rawDescData
}

var file_checkout_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_checkout_service_proto_goTypes = []interface{}{
	(*AddToCartRequest)(nil),      // 0: checkout_v1.AddToCartRequest
	(*DeleteFromCartRequest)(nil), // 1: checkout_v1.DeleteFromCartRequest
	(*AddToCartResponse)(nil),     // 2: checkout_v1.AddToCartResponse
	(*ListCartResponse)(nil),      // 3: checkout_v1.ListCartResponse
	(*ListCartRequest)(nil),       // 4: checkout_v1.ListCartRequest
	(*CartItem)(nil),              // 5: checkout_v1.CartItem
	(*PuchaseRequest)(nil),        // 6: checkout_v1.PuchaseRequest
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_checkout_service_proto_depIdxs = []int32{
	5, // 0: checkout_v1.ListCartResponse.items:type_name -> checkout_v1.CartItem
	0, // 1: checkout_v1.CheckoutV1.AddToCart:input_type -> checkout_v1.AddToCartRequest
	1, // 2: checkout_v1.CheckoutV1.DeleteFromCart:input_type -> checkout_v1.DeleteFromCartRequest
	4, // 3: checkout_v1.CheckoutV1.ListCart:input_type -> checkout_v1.ListCartRequest
	6, // 4: checkout_v1.CheckoutV1.Puchase:input_type -> checkout_v1.PuchaseRequest
	7, // 5: checkout_v1.CheckoutV1.AddToCart:output_type -> google.protobuf.Empty
	7, // 6: checkout_v1.CheckoutV1.DeleteFromCart:output_type -> google.protobuf.Empty
	3, // 7: checkout_v1.CheckoutV1.ListCart:output_type -> checkout_v1.ListCartResponse
	7, // 8: checkout_v1.CheckoutV1.Puchase:output_type -> google.protobuf.Empty
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_checkout_service_proto_init() }
func file_checkout_service_proto_init() {
	if File_checkout_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_checkout_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddToCartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_checkout_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFromCartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_checkout_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddToCartResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_checkout_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCartResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_checkout_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_checkout_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_checkout_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PuchaseRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_checkout_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_checkout_service_proto_goTypes,
		DependencyIndexes: file_checkout_service_proto_depIdxs,
		MessageInfos:      file_checkout_service_proto_msgTypes,
	}.Build()
	File_checkout_service_proto = out.File
	file_checkout_service_proto_rawDesc = nil
	file_checkout_service_proto_goTypes = nil
	file_checkout_service_proto_depIdxs = nil
}
