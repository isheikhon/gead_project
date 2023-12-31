// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.4
// source: analytics/analytics.proto

package analytics

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TotalSales struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sales int32 `protobuf:"varint,3,opt,name=sales,proto3" json:"sales,omitempty"`
}

func (x *TotalSales) Reset() {
	*x = TotalSales{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_analytics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TotalSales) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TotalSales) ProtoMessage() {}

func (x *TotalSales) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_analytics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TotalSales.ProtoReflect.Descriptor instead.
func (*TotalSales) Descriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{0}
}

func (x *TotalSales) GetSales() int32 {
	if x != nil {
		return x.Sales
	}
	return 0
}

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Sales int32  `protobuf:"varint,3,opt,name=sales,proto3" json:"sales,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_analytics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_analytics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{1}
}

func (x *Customer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Customer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Customer) GetSales() int32 {
	if x != nil {
		return x.Sales
	}
	return 0
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Sales int32  `protobuf:"varint,3,opt,name=sales,proto3" json:"sales,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_analytics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_analytics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{2}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetSales() int32 {
	if x != nil {
		return x.Sales
	}
	return 0
}

type GetTotalSalesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTotalSalesRequest) Reset() {
	*x = GetTotalSalesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_analytics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTotalSalesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTotalSalesRequest) ProtoMessage() {}

func (x *GetTotalSalesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_analytics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTotalSalesRequest.ProtoReflect.Descriptor instead.
func (*GetTotalSalesRequest) Descriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{3}
}

type GetSalesByProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSalesByProductRequest) Reset() {
	*x = GetSalesByProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_analytics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSalesByProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSalesByProductRequest) ProtoMessage() {}

func (x *GetSalesByProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_analytics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSalesByProductRequest.ProtoReflect.Descriptor instead.
func (*GetSalesByProductRequest) Descriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{4}
}

type GetTopCustomersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTopCustomersRequest) Reset() {
	*x = GetTopCustomersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_analytics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopCustomersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopCustomersRequest) ProtoMessage() {}

func (x *GetTopCustomersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_analytics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopCustomersRequest.ProtoReflect.Descriptor instead.
func (*GetTopCustomersRequest) Descriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{5}
}

type GetTotalSalesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sales *TotalSales `protobuf:"bytes,1,opt,name=sales,proto3" json:"sales,omitempty"`
}

func (x *GetTotalSalesResponse) Reset() {
	*x = GetTotalSalesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_analytics_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTotalSalesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTotalSalesResponse) ProtoMessage() {}

func (x *GetTotalSalesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_analytics_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTotalSalesResponse.ProtoReflect.Descriptor instead.
func (*GetTotalSalesResponse) Descriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{6}
}

func (x *GetTotalSalesResponse) GetSales() *TotalSales {
	if x != nil {
		return x.Sales
	}
	return nil
}

type GetTopCustomersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customers []*Customer `protobuf:"bytes,1,rep,name=customers,proto3" json:"customers,omitempty"`
}

func (x *GetTopCustomersResponse) Reset() {
	*x = GetTopCustomersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_analytics_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopCustomersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopCustomersResponse) ProtoMessage() {}

func (x *GetTopCustomersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_analytics_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopCustomersResponse.ProtoReflect.Descriptor instead.
func (*GetTopCustomersResponse) Descriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{7}
}

func (x *GetTopCustomersResponse) GetCustomers() []*Customer {
	if x != nil {
		return x.Customers
	}
	return nil
}

type GetSalesByProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *GetSalesByProductResponse) Reset() {
	*x = GetSalesByProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_analytics_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSalesByProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSalesByProductResponse) ProtoMessage() {}

func (x *GetSalesByProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_analytics_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSalesByProductResponse.ProtoReflect.Descriptor instead.
func (*GetSalesByProductResponse) Descriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{8}
}

func (x *GetSalesByProductResponse) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_analytics_analytics_proto protoreflect.FileDescriptor

var file_analytics_analytics_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x22, 0x22, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53,
	0x61, 0x6c, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x22, 0x44, 0x0a, 0x08, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x61,
	0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x61, 0x6c, 0x65, 0x73,
	0x22, 0x43, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x73, 0x61, 0x6c, 0x65, 0x73, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1a, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x70, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x44, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53,
	0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x05,
	0x73, 0x61, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x61, 0x6c,
	0x65, 0x73, 0x52, 0x05, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x22, 0x4c, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x70, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x09, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x22, 0x4b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x61,
	0x6c, 0x65, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x32, 0x9f, 0x02, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x61,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53,
	0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x23, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x12,
	0x21, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x6f, 0x70, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x70, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_analytics_analytics_proto_rawDescOnce sync.Once
	file_analytics_analytics_proto_rawDescData = file_analytics_analytics_proto_rawDesc
)

func file_analytics_analytics_proto_rawDescGZIP() []byte {
	file_analytics_analytics_proto_rawDescOnce.Do(func() {
		file_analytics_analytics_proto_rawDescData = protoimpl.X.CompressGZIP(file_analytics_analytics_proto_rawDescData)
	})
	return file_analytics_analytics_proto_rawDescData
}

var file_analytics_analytics_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_analytics_analytics_proto_goTypes = []interface{}{
	(*TotalSales)(nil),                // 0: analytics.TotalSales
	(*Customer)(nil),                  // 1: analytics.Customer
	(*Product)(nil),                   // 2: analytics.Product
	(*GetTotalSalesRequest)(nil),      // 3: analytics.GetTotalSalesRequest
	(*GetSalesByProductRequest)(nil),  // 4: analytics.GetSalesByProductRequest
	(*GetTopCustomersRequest)(nil),    // 5: analytics.GetTopCustomersRequest
	(*GetTotalSalesResponse)(nil),     // 6: analytics.GetTotalSalesResponse
	(*GetTopCustomersResponse)(nil),   // 7: analytics.GetTopCustomersResponse
	(*GetSalesByProductResponse)(nil), // 8: analytics.GetSalesByProductResponse
}
var file_analytics_analytics_proto_depIdxs = []int32{
	0, // 0: analytics.GetTotalSalesResponse.sales:type_name -> analytics.TotalSales
	1, // 1: analytics.GetTopCustomersResponse.customers:type_name -> analytics.Customer
	2, // 2: analytics.GetSalesByProductResponse.products:type_name -> analytics.Product
	3, // 3: analytics.CustomerService.GetTotalSales:input_type -> analytics.GetTotalSalesRequest
	4, // 4: analytics.CustomerService.GetSalesByProduct:input_type -> analytics.GetSalesByProductRequest
	5, // 5: analytics.CustomerService.GetTopCustomers:input_type -> analytics.GetTopCustomersRequest
	6, // 6: analytics.CustomerService.GetTotalSales:output_type -> analytics.GetTotalSalesResponse
	8, // 7: analytics.CustomerService.GetSalesByProduct:output_type -> analytics.GetSalesByProductResponse
	7, // 8: analytics.CustomerService.GetTopCustomers:output_type -> analytics.GetTopCustomersResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_analytics_analytics_proto_init() }
func file_analytics_analytics_proto_init() {
	if File_analytics_analytics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_analytics_analytics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TotalSales); i {
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
		file_analytics_analytics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
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
		file_analytics_analytics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_analytics_analytics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTotalSalesRequest); i {
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
		file_analytics_analytics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSalesByProductRequest); i {
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
		file_analytics_analytics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopCustomersRequest); i {
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
		file_analytics_analytics_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTotalSalesResponse); i {
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
		file_analytics_analytics_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopCustomersResponse); i {
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
		file_analytics_analytics_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSalesByProductResponse); i {
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
			RawDescriptor: file_analytics_analytics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_analytics_analytics_proto_goTypes,
		DependencyIndexes: file_analytics_analytics_proto_depIdxs,
		MessageInfos:      file_analytics_analytics_proto_msgTypes,
	}.Build()
	File_analytics_analytics_proto = out.File
	file_analytics_analytics_proto_rawDesc = nil
	file_analytics_analytics_proto_goTypes = nil
	file_analytics_analytics_proto_depIdxs = nil
}
