// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: newGoal_msg.proto

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

type NewGoalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Purpose1 string `protobuf:"bytes,2,opt,name=purpose1,proto3" json:"purpose1,omitempty"`
	Purpose2 string `protobuf:"bytes,3,opt,name=purpose2,proto3" json:"purpose2,omitempty"`
}

func (x *NewGoalRequest) Reset() {
	*x = NewGoalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_newGoal_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewGoalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewGoalRequest) ProtoMessage() {}

func (x *NewGoalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_newGoal_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewGoalRequest.ProtoReflect.Descriptor instead.
func (*NewGoalRequest) Descriptor() ([]byte, []int) {
	return file_newGoal_msg_proto_rawDescGZIP(), []int{0}
}

func (x *NewGoalRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NewGoalRequest) GetPurpose1() string {
	if x != nil {
		return x.Purpose1
	}
	return ""
}

func (x *NewGoalRequest) GetPurpose2() string {
	if x != nil {
		return x.Purpose2
	}
	return ""
}

type NewGoalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *NewGoalResponse) Reset() {
	*x = NewGoalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_newGoal_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewGoalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewGoalResponse) ProtoMessage() {}

func (x *NewGoalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_newGoal_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewGoalResponse.ProtoReflect.Descriptor instead.
func (*NewGoalResponse) Descriptor() ([]byte, []int) {
	return file_newGoal_msg_proto_rawDescGZIP(), []int{1}
}

func (x *NewGoalResponse) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_newGoal_msg_proto protoreflect.FileDescriptor

var file_newGoal_msg_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6e, 0x65, 0x77, 0x47, 0x6f, 0x61, 0x6c, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x5e, 0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x47, 0x6f,
	0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x31, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x32, 0x22, 0x27, 0x0a, 0x0f, 0x6e, 0x65, 0x77, 0x47, 0x6f,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_newGoal_msg_proto_rawDescOnce sync.Once
	file_newGoal_msg_proto_rawDescData = file_newGoal_msg_proto_rawDesc
)

func file_newGoal_msg_proto_rawDescGZIP() []byte {
	file_newGoal_msg_proto_rawDescOnce.Do(func() {
		file_newGoal_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_newGoal_msg_proto_rawDescData)
	})
	return file_newGoal_msg_proto_rawDescData
}

var file_newGoal_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_newGoal_msg_proto_goTypes = []interface{}{
	(*NewGoalRequest)(nil),  // 0: pb.newGoalRequest
	(*NewGoalResponse)(nil), // 1: pb.newGoalResponse
}
var file_newGoal_msg_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_newGoal_msg_proto_init() }
func file_newGoal_msg_proto_init() {
	if File_newGoal_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_newGoal_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewGoalRequest); i {
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
		file_newGoal_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewGoalResponse); i {
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
			RawDescriptor: file_newGoal_msg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_newGoal_msg_proto_goTypes,
		DependencyIndexes: file_newGoal_msg_proto_depIdxs,
		MessageInfos:      file_newGoal_msg_proto_msgTypes,
	}.Build()
	File_newGoal_msg_proto = out.File
	file_newGoal_msg_proto_rawDesc = nil
	file_newGoal_msg_proto_goTypes = nil
	file_newGoal_msg_proto_depIdxs = nil
}
