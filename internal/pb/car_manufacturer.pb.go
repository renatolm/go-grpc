// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/car_manufacturer.proto

package pb

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

type Manufacturer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Origincountry string `protobuf:"bytes,3,opt,name=origincountry,proto3" json:"origincountry,omitempty"`
}

func (x *Manufacturer) Reset() {
	*x = Manufacturer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_car_manufacturer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Manufacturer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Manufacturer) ProtoMessage() {}

func (x *Manufacturer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_car_manufacturer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Manufacturer.ProtoReflect.Descriptor instead.
func (*Manufacturer) Descriptor() ([]byte, []int) {
	return file_proto_car_manufacturer_proto_rawDescGZIP(), []int{0}
}

func (x *Manufacturer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Manufacturer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Manufacturer) GetOrigincountry() string {
	if x != nil {
		return x.Origincountry
	}
	return ""
}

type CreateManufacturerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Origincountry string `protobuf:"bytes,2,opt,name=origincountry,proto3" json:"origincountry,omitempty"`
}

func (x *CreateManufacturerRequest) Reset() {
	*x = CreateManufacturerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_car_manufacturer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateManufacturerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateManufacturerRequest) ProtoMessage() {}

func (x *CreateManufacturerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_car_manufacturer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateManufacturerRequest.ProtoReflect.Descriptor instead.
func (*CreateManufacturerRequest) Descriptor() ([]byte, []int) {
	return file_proto_car_manufacturer_proto_rawDescGZIP(), []int{1}
}

func (x *CreateManufacturerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateManufacturerRequest) GetOrigincountry() string {
	if x != nil {
		return x.Origincountry
	}
	return ""
}

type ManufacturerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Manufacturer *Manufacturer `protobuf:"bytes,1,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
}

func (x *ManufacturerResponse) Reset() {
	*x = ManufacturerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_car_manufacturer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManufacturerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManufacturerResponse) ProtoMessage() {}

func (x *ManufacturerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_car_manufacturer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManufacturerResponse.ProtoReflect.Descriptor instead.
func (*ManufacturerResponse) Descriptor() ([]byte, []int) {
	return file_proto_car_manufacturer_proto_rawDescGZIP(), []int{2}
}

func (x *ManufacturerResponse) GetManufacturer() *Manufacturer {
	if x != nil {
		return x.Manufacturer
	}
	return nil
}

var File_proto_car_manufacturer_proto protoreflect.FileDescriptor

var file_proto_car_manufacturer_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x75,
	0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x58, 0x0a, 0x0c, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x55, 0x0a, 0x19,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x22, 0x4c, 0x0a, 0x14, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x6d,
	0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x72, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x72, 0x32, 0x5e, 0x0a, 0x13, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x1d,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x70, 0x62, 0x2e, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x22,
	0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_car_manufacturer_proto_rawDescOnce sync.Once
	file_proto_car_manufacturer_proto_rawDescData = file_proto_car_manufacturer_proto_rawDesc
)

func file_proto_car_manufacturer_proto_rawDescGZIP() []byte {
	file_proto_car_manufacturer_proto_rawDescOnce.Do(func() {
		file_proto_car_manufacturer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_car_manufacturer_proto_rawDescData)
	})
	return file_proto_car_manufacturer_proto_rawDescData
}

var file_proto_car_manufacturer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_car_manufacturer_proto_goTypes = []interface{}{
	(*Manufacturer)(nil),              // 0: pb.Manufacturer
	(*CreateManufacturerRequest)(nil), // 1: pb.CreateManufacturerRequest
	(*ManufacturerResponse)(nil),      // 2: pb.ManufacturerResponse
}
var file_proto_car_manufacturer_proto_depIdxs = []int32{
	0, // 0: pb.ManufacturerResponse.manufacturer:type_name -> pb.Manufacturer
	1, // 1: pb.ManufacturerService.CreateManufacturer:input_type -> pb.CreateManufacturerRequest
	0, // 2: pb.ManufacturerService.CreateManufacturer:output_type -> pb.Manufacturer
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_car_manufacturer_proto_init() }
func file_proto_car_manufacturer_proto_init() {
	if File_proto_car_manufacturer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_car_manufacturer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Manufacturer); i {
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
		file_proto_car_manufacturer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateManufacturerRequest); i {
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
		file_proto_car_manufacturer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManufacturerResponse); i {
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
			RawDescriptor: file_proto_car_manufacturer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_car_manufacturer_proto_goTypes,
		DependencyIndexes: file_proto_car_manufacturer_proto_depIdxs,
		MessageInfos:      file_proto_car_manufacturer_proto_msgTypes,
	}.Build()
	File_proto_car_manufacturer_proto = out.File
	file_proto_car_manufacturer_proto_rawDesc = nil
	file_proto_car_manufacturer_proto_goTypes = nil
	file_proto_car_manufacturer_proto_depIdxs = nil
}
