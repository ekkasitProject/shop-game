// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: modules/player/playerPb/playerPb.proto

package shop_game

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

// Structures
type PlayerProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Username  string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	RoleCode  string `protobuf:"bytes,4,opt,name=roleCode,proto3" json:"roleCode,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *PlayerProfile) Reset() {
	*x = PlayerProfile{}
	mi := &file_modules_player_playerPb_playerPb_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlayerProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerProfile) ProtoMessage() {}

func (x *PlayerProfile) ProtoReflect() protoreflect.Message {
	mi := &file_modules_player_playerPb_playerPb_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerProfile.ProtoReflect.Descriptor instead.
func (*PlayerProfile) Descriptor() ([]byte, []int) {
	return file_modules_player_playerPb_playerPb_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerProfile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PlayerProfile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PlayerProfile) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *PlayerProfile) GetRoleCode() string {
	if x != nil {
		return x.RoleCode
	}
	return ""
}

func (x *PlayerProfile) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PlayerProfile) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CredentialSearchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CredentialSearchReq) Reset() {
	*x = CredentialSearchReq{}
	mi := &file_modules_player_playerPb_playerPb_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CredentialSearchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CredentialSearchReq) ProtoMessage() {}

func (x *CredentialSearchReq) ProtoReflect() protoreflect.Message {
	mi := &file_modules_player_playerPb_playerPb_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CredentialSearchReq.ProtoReflect.Descriptor instead.
func (*CredentialSearchReq) Descriptor() ([]byte, []int) {
	return file_modules_player_playerPb_playerPb_proto_rawDescGZIP(), []int{1}
}

func (x *CredentialSearchReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CredentialSearchReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type FindOnePlayerProfileToRefreshReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
}

func (x *FindOnePlayerProfileToRefreshReq) Reset() {
	*x = FindOnePlayerProfileToRefreshReq{}
	mi := &file_modules_player_playerPb_playerPb_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindOnePlayerProfileToRefreshReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOnePlayerProfileToRefreshReq) ProtoMessage() {}

func (x *FindOnePlayerProfileToRefreshReq) ProtoReflect() protoreflect.Message {
	mi := &file_modules_player_playerPb_playerPb_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOnePlayerProfileToRefreshReq.ProtoReflect.Descriptor instead.
func (*FindOnePlayerProfileToRefreshReq) Descriptor() ([]byte, []int) {
	return file_modules_player_playerPb_playerPb_proto_rawDescGZIP(), []int{2}
}

func (x *FindOnePlayerProfileToRefreshReq) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type GetPlayerSavingAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
}

func (x *GetPlayerSavingAccountReq) Reset() {
	*x = GetPlayerSavingAccountReq{}
	mi := &file_modules_player_playerPb_playerPb_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPlayerSavingAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerSavingAccountReq) ProtoMessage() {}

func (x *GetPlayerSavingAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_modules_player_playerPb_playerPb_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerSavingAccountReq.ProtoReflect.Descriptor instead.
func (*GetPlayerSavingAccountReq) Descriptor() ([]byte, []int) {
	return file_modules_player_playerPb_playerPb_proto_rawDescGZIP(), []int{3}
}

func (x *GetPlayerSavingAccountReq) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type GetPlayerSavingAccountRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string  `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Balance  float64 `protobuf:"fixed64,2,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *GetPlayerSavingAccountRes) Reset() {
	*x = GetPlayerSavingAccountRes{}
	mi := &file_modules_player_playerPb_playerPb_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPlayerSavingAccountRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerSavingAccountRes) ProtoMessage() {}

func (x *GetPlayerSavingAccountRes) ProtoReflect() protoreflect.Message {
	mi := &file_modules_player_playerPb_playerPb_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerSavingAccountRes.ProtoReflect.Descriptor instead.
func (*GetPlayerSavingAccountRes) Descriptor() ([]byte, []int) {
	return file_modules_player_playerPb_playerPb_proto_rawDescGZIP(), []int{4}
}

func (x *GetPlayerSavingAccountRes) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *GetPlayerSavingAccountRes) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

var File_modules_player_playerPb_playerPb_proto protoreflect.FileDescriptor

var file_modules_player_playerPb_playerPb_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x62, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x50, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x0d, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x6f, 0x6c, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x6f, 0x6c, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x47, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x3e, 0x0a, 0x20, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x6f, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x37, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x61, 0x76, 0x69,
	0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x32, 0xf3, 0x01, 0x0a, 0x11,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x38, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x52, 0x0a, 0x1d, 0x46,
	0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x54, 0x6f, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x21, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x54, 0x6f, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x71, 0x1a,
	0x0e, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x50, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x61, 0x76, 0x69,
	0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6b, 0x6b, 0x61, 0x73, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x73,
	0x68, 0x6f, 0x70, 0x2d, 0x67, 0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_modules_player_playerPb_playerPb_proto_rawDescOnce sync.Once
	file_modules_player_playerPb_playerPb_proto_rawDescData = file_modules_player_playerPb_playerPb_proto_rawDesc
)

func file_modules_player_playerPb_playerPb_proto_rawDescGZIP() []byte {
	file_modules_player_playerPb_playerPb_proto_rawDescOnce.Do(func() {
		file_modules_player_playerPb_playerPb_proto_rawDescData = protoimpl.X.CompressGZIP(file_modules_player_playerPb_playerPb_proto_rawDescData)
	})
	return file_modules_player_playerPb_playerPb_proto_rawDescData
}

var file_modules_player_playerPb_playerPb_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_modules_player_playerPb_playerPb_proto_goTypes = []any{
	(*PlayerProfile)(nil),                    // 0: PlayerProfile
	(*CredentialSearchReq)(nil),              // 1: CredentialSearchReq
	(*FindOnePlayerProfileToRefreshReq)(nil), // 2: FindOnePlayerProfileToRefreshReq
	(*GetPlayerSavingAccountReq)(nil),        // 3: GetPlayerSavingAccountReq
	(*GetPlayerSavingAccountRes)(nil),        // 4: GetPlayerSavingAccountRes
}
var file_modules_player_playerPb_playerPb_proto_depIdxs = []int32{
	1, // 0: PlayerGrpcService.CredentialSearch:input_type -> CredentialSearchReq
	2, // 1: PlayerGrpcService.FindOnePlayerProfileToRefresh:input_type -> FindOnePlayerProfileToRefreshReq
	3, // 2: PlayerGrpcService.GetPlayerSavingAccount:input_type -> GetPlayerSavingAccountReq
	0, // 3: PlayerGrpcService.CredentialSearch:output_type -> PlayerProfile
	0, // 4: PlayerGrpcService.FindOnePlayerProfileToRefresh:output_type -> PlayerProfile
	4, // 5: PlayerGrpcService.GetPlayerSavingAccount:output_type -> GetPlayerSavingAccountRes
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_modules_player_playerPb_playerPb_proto_init() }
func file_modules_player_playerPb_playerPb_proto_init() {
	if File_modules_player_playerPb_playerPb_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_modules_player_playerPb_playerPb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_modules_player_playerPb_playerPb_proto_goTypes,
		DependencyIndexes: file_modules_player_playerPb_playerPb_proto_depIdxs,
		MessageInfos:      file_modules_player_playerPb_playerPb_proto_msgTypes,
	}.Build()
	File_modules_player_playerPb_playerPb_proto = out.File
	file_modules_player_playerPb_playerPb_proto_rawDesc = nil
	file_modules_player_playerPb_playerPb_proto_goTypes = nil
	file_modules_player_playerPb_playerPb_proto_depIdxs = nil
}
