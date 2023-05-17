// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: user_details.proto

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Country string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Hashkey int64  `protobuf:"varint,4,opt,name=hashkey,proto3" json:"hashkey,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_details_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_details_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_details_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *User) GetHashkey() int64 {
	if x != nil {
		return x.Hashkey
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_details_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_user_details_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_user_details_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type GetUserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key         int64  `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Countryname string `protobuf:"bytes,2,opt,name=countryname,proto3" json:"countryname,omitempty"`
}

func (x *GetUserData) Reset() {
	*x = GetUserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_details_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserData) ProtoMessage() {}

func (x *GetUserData) ProtoReflect() protoreflect.Message {
	mi := &file_user_details_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserData.ProtoReflect.Descriptor instead.
func (*GetUserData) Descriptor() ([]byte, []int) {
	return file_user_details_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserData) GetKey() int64 {
	if x != nil {
		return x.Key
	}
	return 0
}

func (x *GetUserData) GetCountryname() string {
	if x != nil {
		return x.Countryname
	}
	return ""
}

type SavedRecords struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records map[uint32]*User `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SavedRecords) Reset() {
	*x = SavedRecords{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_details_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavedRecords) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavedRecords) ProtoMessage() {}

func (x *SavedRecords) ProtoReflect() protoreflect.Message {
	mi := &file_user_details_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavedRecords.ProtoReflect.Descriptor instead.
func (*SavedRecords) Descriptor() ([]byte, []int) {
	return file_user_details_proto_rawDescGZIP(), []int{3}
}

func (x *SavedRecords) GetRecords() map[uint32]*User {
	if x != nil {
		return x.Records
	}
	return nil
}

var File_user_details_proto protoreflect.FileDescriptor

var file_user_details_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x73, 0x68, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x68, 0x61, 0x73, 0x68, 0x6b, 0x65, 0x79, 0x22, 0x22, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x41, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x0c, 0x53, 0x61, 0x76, 0x65, 0x64, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x12, 0x3a, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x76,
	0x65, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x1a,
	0x47, 0x0a, 0x0c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x74, 0x0a, 0x10, 0x43, 0x61, 0x73, 0x73,
	0x61, 0x6e, 0x64, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x08,
	0x53, 0x61, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x61, 0x76, 0x65, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x00, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_details_proto_rawDescOnce sync.Once
	file_user_details_proto_rawDescData = file_user_details_proto_rawDesc
)

func file_user_details_proto_rawDescGZIP() []byte {
	file_user_details_proto_rawDescOnce.Do(func() {
		file_user_details_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_details_proto_rawDescData)
	})
	return file_user_details_proto_rawDescData
}

var file_user_details_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_user_details_proto_goTypes = []interface{}{
	(*User)(nil),         // 0: proto.User
	(*Response)(nil),     // 1: proto.Response
	(*GetUserData)(nil),  // 2: proto.GetUserData
	(*SavedRecords)(nil), // 3: proto.SavedRecords
	nil,                  // 4: proto.SavedRecords.RecordsEntry
}
var file_user_details_proto_depIdxs = []int32{
	4, // 0: proto.SavedRecords.records:type_name -> proto.SavedRecords.RecordsEntry
	0, // 1: proto.SavedRecords.RecordsEntry.value:type_name -> proto.User
	0, // 2: proto.CassandraService.SaveData:input_type -> proto.User
	2, // 3: proto.CassandraService.GetData:input_type -> proto.GetUserData
	1, // 4: proto.CassandraService.SaveData:output_type -> proto.Response
	3, // 5: proto.CassandraService.GetData:output_type -> proto.SavedRecords
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_details_proto_init() }
func file_user_details_proto_init() {
	if File_user_details_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_details_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_user_details_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_user_details_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserData); i {
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
		file_user_details_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavedRecords); i {
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
			RawDescriptor: file_user_details_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_details_proto_goTypes,
		DependencyIndexes: file_user_details_proto_depIdxs,
		MessageInfos:      file_user_details_proto_msgTypes,
	}.Build()
	File_user_details_proto = out.File
	file_user_details_proto_rawDesc = nil
	file_user_details_proto_goTypes = nil
	file_user_details_proto_depIdxs = nil
}