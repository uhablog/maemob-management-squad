// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: pkg/api/proto/squad-service.proto

package __

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

type CreateSquadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId   string `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	PlayerId string `protobuf:"bytes,2,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *CreateSquadRequest) Reset() {
	*x = CreateSquadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_proto_squad_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSquadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSquadRequest) ProtoMessage() {}

func (x *CreateSquadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_proto_squad_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSquadRequest.ProtoReflect.Descriptor instead.
func (*CreateSquadRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_proto_squad_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSquadRequest) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *CreateSquadRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type CreateSquadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SquadId string `protobuf:"bytes,1,opt,name=squad_id,proto3" json:"squad_id,omitempty"`
}

func (x *CreateSquadResponse) Reset() {
	*x = CreateSquadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_proto_squad_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSquadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSquadResponse) ProtoMessage() {}

func (x *CreateSquadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_proto_squad_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSquadResponse.ProtoReflect.Descriptor instead.
func (*CreateSquadResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_proto_squad_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSquadResponse) GetSquadId() string {
	if x != nil {
		return x.SquadId
	}
	return ""
}

var File_pkg_api_proto_squad_service_proto protoreflect.FileDescriptor

var file_pkg_api_proto_squad_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x71, 0x75, 0x61, 0x64, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x71, 0x75, 0x61, 0x64, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x4a, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x71, 0x75, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x22, 0x31, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x71, 0x75, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x71, 0x75, 0x61, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x71, 0x75, 0x61, 0x64, 0x5f, 0x69, 0x64,
	0x32, 0x62, 0x0a, 0x0c, 0x53, 0x71, 0x75, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x52, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x71, 0x75, 0x61, 0x64, 0x12,
	0x20, 0x2e, 0x73, 0x71, 0x75, 0x61, 0x64, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x71, 0x75, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x73, 0x71, 0x75, 0x61, 0x64, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x71, 0x75, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_api_proto_squad_service_proto_rawDescOnce sync.Once
	file_pkg_api_proto_squad_service_proto_rawDescData = file_pkg_api_proto_squad_service_proto_rawDesc
)

func file_pkg_api_proto_squad_service_proto_rawDescGZIP() []byte {
	file_pkg_api_proto_squad_service_proto_rawDescOnce.Do(func() {
		file_pkg_api_proto_squad_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_api_proto_squad_service_proto_rawDescData)
	})
	return file_pkg_api_proto_squad_service_proto_rawDescData
}

var file_pkg_api_proto_squad_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_api_proto_squad_service_proto_goTypes = []any{
	(*CreateSquadRequest)(nil),  // 0: squadservice.CreateSquadRequest
	(*CreateSquadResponse)(nil), // 1: squadservice.CreateSquadResponse
}
var file_pkg_api_proto_squad_service_proto_depIdxs = []int32{
	0, // 0: squadservice.SquadService.CreateSquad:input_type -> squadservice.CreateSquadRequest
	1, // 1: squadservice.SquadService.CreateSquad:output_type -> squadservice.CreateSquadResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_api_proto_squad_service_proto_init() }
func file_pkg_api_proto_squad_service_proto_init() {
	if File_pkg_api_proto_squad_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_api_proto_squad_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateSquadRequest); i {
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
		file_pkg_api_proto_squad_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateSquadResponse); i {
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
			RawDescriptor: file_pkg_api_proto_squad_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_api_proto_squad_service_proto_goTypes,
		DependencyIndexes: file_pkg_api_proto_squad_service_proto_depIdxs,
		MessageInfos:      file_pkg_api_proto_squad_service_proto_msgTypes,
	}.Build()
	File_pkg_api_proto_squad_service_proto = out.File
	file_pkg_api_proto_squad_service_proto_rawDesc = nil
	file_pkg_api_proto_squad_service_proto_goTypes = nil
	file_pkg_api_proto_squad_service_proto_depIdxs = nil
}
