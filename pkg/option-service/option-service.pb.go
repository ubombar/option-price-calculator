// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: pkg/option-service/option-service.proto

package option_server

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

type OptionPriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockPrice           float32 `protobuf:"fixed32,1,opt,name=stock_price,json=stockPrice,proto3" json:"stock_price,omitempty"`
	ExercisePrice        float32 `protobuf:"fixed32,2,opt,name=exercise_price,json=exercisePrice,proto3" json:"exercise_price,omitempty"`
	Volatility           float32 `protobuf:"fixed32,3,opt,name=volatility,proto3" json:"volatility,omitempty"`
	ExpirationTime       float32 `protobuf:"fixed32,4,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	RiskFreeInterestRate float32 `protobuf:"fixed32,5,opt,name=risk_free_interest_rate,json=riskFreeInterestRate,proto3" json:"risk_free_interest_rate,omitempty"`
}

func (x *OptionPriceRequest) Reset() {
	*x = OptionPriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_option_service_option_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionPriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionPriceRequest) ProtoMessage() {}

func (x *OptionPriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_option_service_option_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionPriceRequest.ProtoReflect.Descriptor instead.
func (*OptionPriceRequest) Descriptor() ([]byte, []int) {
	return file_pkg_option_service_option_service_proto_rawDescGZIP(), []int{0}
}

func (x *OptionPriceRequest) GetStockPrice() float32 {
	if x != nil {
		return x.StockPrice
	}
	return 0
}

func (x *OptionPriceRequest) GetExercisePrice() float32 {
	if x != nil {
		return x.ExercisePrice
	}
	return 0
}

func (x *OptionPriceRequest) GetVolatility() float32 {
	if x != nil {
		return x.Volatility
	}
	return 0
}

func (x *OptionPriceRequest) GetExpirationTime() float32 {
	if x != nil {
		return x.ExpirationTime
	}
	return 0
}

func (x *OptionPriceRequest) GetRiskFreeInterestRate() float32 {
	if x != nil {
		return x.RiskFreeInterestRate
	}
	return 0
}

type OptionPriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockPrice           float32 `protobuf:"fixed32,1,opt,name=stock_price,json=stockPrice,proto3" json:"stock_price,omitempty"`
	ExercisePrice        float32 `protobuf:"fixed32,2,opt,name=exercise_price,json=exercisePrice,proto3" json:"exercise_price,omitempty"`
	Volatility           float32 `protobuf:"fixed32,3,opt,name=volatility,proto3" json:"volatility,omitempty"`
	ExpirationTime       float32 `protobuf:"fixed32,4,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	RiskFreeInterestRate float32 `protobuf:"fixed32,5,opt,name=risk_free_interest_rate,json=riskFreeInterestRate,proto3" json:"risk_free_interest_rate,omitempty"`
	OptionPrice          float32 `protobuf:"fixed32,6,opt,name=option_price,json=optionPrice,proto3" json:"option_price,omitempty"`
}

func (x *OptionPriceResponse) Reset() {
	*x = OptionPriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_option_service_option_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionPriceResponse) ProtoMessage() {}

func (x *OptionPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_option_service_option_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionPriceResponse.ProtoReflect.Descriptor instead.
func (*OptionPriceResponse) Descriptor() ([]byte, []int) {
	return file_pkg_option_service_option_service_proto_rawDescGZIP(), []int{1}
}

func (x *OptionPriceResponse) GetStockPrice() float32 {
	if x != nil {
		return x.StockPrice
	}
	return 0
}

func (x *OptionPriceResponse) GetExercisePrice() float32 {
	if x != nil {
		return x.ExercisePrice
	}
	return 0
}

func (x *OptionPriceResponse) GetVolatility() float32 {
	if x != nil {
		return x.Volatility
	}
	return 0
}

func (x *OptionPriceResponse) GetExpirationTime() float32 {
	if x != nil {
		return x.ExpirationTime
	}
	return 0
}

func (x *OptionPriceResponse) GetRiskFreeInterestRate() float32 {
	if x != nil {
		return x.RiskFreeInterestRate
	}
	return 0
}

func (x *OptionPriceResponse) GetOptionPrice() float32 {
	if x != nil {
		return x.OptionPrice
	}
	return 0
}

var File_pkg_option_service_option_service_proto protoreflect.FileDescriptor

var file_pkg_option_service_option_service_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x01, 0x0a, 0x12, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x65, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x6f, 0x6c, 0x61,
	0x74, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x76, 0x6f,
	0x6c, 0x61, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x35, 0x0a, 0x17, 0x72, 0x69, 0x73, 0x6b, 0x5f, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x14, 0x72, 0x69, 0x73, 0x6b, 0x46, 0x72, 0x65, 0x65, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x65, 0x73, 0x74, 0x52, 0x61, 0x74, 0x65, 0x22, 0x80, 0x02, 0x0a, 0x13, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x65, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x6f, 0x6c, 0x61,
	0x74, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x76, 0x6f,
	0x6c, 0x61, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x35, 0x0a, 0x17, 0x72, 0x69, 0x73, 0x6b, 0x5f, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x14, 0x72, 0x69, 0x73, 0x6b, 0x46, 0x72, 0x65, 0x65, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x65, 0x73, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x32, 0x49, 0x0a, 0x0d, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0b,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x13, 0x2e, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x62, 0x6f, 0x6d, 0x62, 0x61, 0x72, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2d, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2d, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_option_service_option_service_proto_rawDescOnce sync.Once
	file_pkg_option_service_option_service_proto_rawDescData = file_pkg_option_service_option_service_proto_rawDesc
)

func file_pkg_option_service_option_service_proto_rawDescGZIP() []byte {
	file_pkg_option_service_option_service_proto_rawDescOnce.Do(func() {
		file_pkg_option_service_option_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_option_service_option_service_proto_rawDescData)
	})
	return file_pkg_option_service_option_service_proto_rawDescData
}

var file_pkg_option_service_option_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_option_service_option_service_proto_goTypes = []interface{}{
	(*OptionPriceRequest)(nil),  // 0: OptionPriceRequest
	(*OptionPriceResponse)(nil), // 1: OptionPriceResponse
}
var file_pkg_option_service_option_service_proto_depIdxs = []int32{
	0, // 0: OptionService.OptionPrice:input_type -> OptionPriceRequest
	1, // 1: OptionService.OptionPrice:output_type -> OptionPriceResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_option_service_option_service_proto_init() }
func file_pkg_option_service_option_service_proto_init() {
	if File_pkg_option_service_option_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_option_service_option_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionPriceRequest); i {
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
		file_pkg_option_service_option_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionPriceResponse); i {
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
			RawDescriptor: file_pkg_option_service_option_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_option_service_option_service_proto_goTypes,
		DependencyIndexes: file_pkg_option_service_option_service_proto_depIdxs,
		MessageInfos:      file_pkg_option_service_option_service_proto_msgTypes,
	}.Build()
	File_pkg_option_service_option_service_proto = out.File
	file_pkg_option_service_option_service_proto_rawDesc = nil
	file_pkg_option_service_option_service_proto_goTypes = nil
	file_pkg_option_service_option_service_proto_depIdxs = nil
}
