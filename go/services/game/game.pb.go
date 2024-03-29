// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/services/game/game.proto

package game

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Game []*Game `protobuf:"bytes,1,rep,name=Game,proto3" json:"Game,omitempty"`
}

func (x *GameResponse) Reset() {
	*x = GameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_game_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameResponse) ProtoMessage() {}

func (x *GameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_game_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameResponse.ProtoReflect.Descriptor instead.
func (*GameResponse) Descriptor() ([]byte, []int) {
	return file_proto_services_game_game_proto_rawDescGZIP(), []int{0}
}

func (x *GameResponse) GetGame() []*Game {
	if x != nil {
		return x.Game
	}
	return nil
}

type Game struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         string      `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Code       string      `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
	GroundID   string      `protobuf:"bytes,3,opt,name=GroundID,proto3" json:"GroundID,omitempty"`
	EventID    string      `protobuf:"bytes,4,opt,name=EventID,proto3" json:"EventID,omitempty"`
	Name       string      `protobuf:"bytes,5,opt,name=Name,proto3" json:"Name,omitempty"`
	StartDate  string      `protobuf:"bytes,6,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	EndDate    string      `protobuf:"bytes,7,opt,name=EndDate,proto3" json:"EndDate,omitempty"`
	ModifiedBy *ModifiedBy `protobuf:"bytes,8,opt,name=ModifiedBy,proto3" json:"ModifiedBy,omitempty"`
	Metadata   *Metadata   `protobuf:"bytes,9,opt,name=Metadata,proto3" json:"Metadata,omitempty"`
	CreatedAt  string      `protobuf:"bytes,10,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt  string      `protobuf:"bytes,11,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Game) Reset() {
	*x = Game{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_game_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Game) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Game) ProtoMessage() {}

func (x *Game) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_game_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Game.ProtoReflect.Descriptor instead.
func (*Game) Descriptor() ([]byte, []int) {
	return file_proto_services_game_game_proto_rawDescGZIP(), []int{1}
}

func (x *Game) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Game) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Game) GetGroundID() string {
	if x != nil {
		return x.GroundID
	}
	return ""
}

func (x *Game) GetEventID() string {
	if x != nil {
		return x.EventID
	}
	return ""
}

func (x *Game) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Game) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *Game) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *Game) GetModifiedBy() *ModifiedBy {
	if x != nil {
		return x.ModifiedBy
	}
	return nil
}

func (x *Game) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Game) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Game) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type GameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroundID   string      `protobuf:"bytes,1,opt,name=GroundID,proto3" json:"GroundID,omitempty"`
	EventID    string      `protobuf:"bytes,2,opt,name=EventID,proto3" json:"EventID,omitempty"`
	Name       string      `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	StartDate  string      `protobuf:"bytes,4,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	EndDate    string      `protobuf:"bytes,5,opt,name=EndDate,proto3" json:"EndDate,omitempty"`
	ModifiedBy *ModifiedBy `protobuf:"bytes,6,opt,name=ModifiedBy,proto3" json:"ModifiedBy,omitempty"`
	Metadata   *Metadata   `protobuf:"bytes,7,opt,name=Metadata,proto3" json:"Metadata,omitempty"`
}

func (x *GameRequest) Reset() {
	*x = GameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_game_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameRequest) ProtoMessage() {}

func (x *GameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_game_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameRequest.ProtoReflect.Descriptor instead.
func (*GameRequest) Descriptor() ([]byte, []int) {
	return file_proto_services_game_game_proto_rawDescGZIP(), []int{2}
}

func (x *GameRequest) GetGroundID() string {
	if x != nil {
		return x.GroundID
	}
	return ""
}

func (x *GameRequest) GetEventID() string {
	if x != nil {
		return x.EventID
	}
	return ""
}

func (x *GameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GameRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *GameRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *GameRequest) GetModifiedBy() *ModifiedBy {
	if x != nil {
		return x.ModifiedBy
	}
	return nil
}

func (x *GameRequest) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageURL string `protobuf:"bytes,1,opt,name=ImageURL,proto3" json:"ImageURL,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_game_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_game_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_proto_services_game_game_proto_rawDescGZIP(), []int{3}
}

func (x *Metadata) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

type ModifiedBy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Role string `protobuf:"bytes,2,opt,name=Role,proto3" json:"Role,omitempty"`
}

func (x *ModifiedBy) Reset() {
	*x = ModifiedBy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_game_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifiedBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifiedBy) ProtoMessage() {}

func (x *ModifiedBy) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_game_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifiedBy.ProtoReflect.Descriptor instead.
func (*ModifiedBy) Descriptor() ([]byte, []int) {
	return file_proto_services_game_game_proto_rawDescGZIP(), []int{4}
}

func (x *ModifiedBy) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ModifiedBy) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

var File_proto_services_game_game_proto protoreflect.FileDescriptor

var file_proto_services_game_game_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x0c, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x47,
	0x61, 0x6d, 0x65, 0x22, 0xc6, 0x02, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x47, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x47, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x42, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x42, 0x79, 0x52, 0x0a, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x42, 0x79, 0x12, 0x2a, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xed, 0x01, 0x0a,
	0x0b, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x47, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x47, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x30,
	0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x42, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x42, 0x79, 0x52, 0x0a, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x42, 0x79,
	0x12, 0x2a, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x26, 0x0a, 0x08,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x52, 0x4c, 0x22, 0x30, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x32, 0x7d, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x2b, 0x5a, 0x29, 0x66, 0x61, 0x6c, 0x74, 0x61, 0x72, 0x2d,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x3b, 0x67, 0x61,
	0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_services_game_game_proto_rawDescOnce sync.Once
	file_proto_services_game_game_proto_rawDescData = file_proto_services_game_game_proto_rawDesc
)

func file_proto_services_game_game_proto_rawDescGZIP() []byte {
	file_proto_services_game_game_proto_rawDescOnce.Do(func() {
		file_proto_services_game_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_services_game_game_proto_rawDescData)
	})
	return file_proto_services_game_game_proto_rawDescData
}

var file_proto_services_game_game_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_services_game_game_proto_goTypes = []interface{}{
	(*GameResponse)(nil),  // 0: post.GameResponse
	(*Game)(nil),          // 1: post.Game
	(*GameRequest)(nil),   // 2: post.GameRequest
	(*Metadata)(nil),      // 3: post.Metadata
	(*ModifiedBy)(nil),    // 4: post.ModifiedBy
	(*emptypb.Empty)(nil), // 5: google.protobuf.Empty
}
var file_proto_services_game_game_proto_depIdxs = []int32{
	1, // 0: post.GameResponse.Game:type_name -> post.Game
	4, // 1: post.Game.ModifiedBy:type_name -> post.ModifiedBy
	3, // 2: post.Game.Metadata:type_name -> post.Metadata
	4, // 3: post.GameRequest.ModifiedBy:type_name -> post.ModifiedBy
	3, // 4: post.GameRequest.Metadata:type_name -> post.Metadata
	5, // 5: post.GameService.GetGame:input_type -> google.protobuf.Empty
	2, // 6: post.GameService.CreateGame:input_type -> post.GameRequest
	0, // 7: post.GameService.GetGame:output_type -> post.GameResponse
	5, // 8: post.GameService.CreateGame:output_type -> google.protobuf.Empty
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_services_game_game_proto_init() }
func file_proto_services_game_game_proto_init() {
	if File_proto_services_game_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_services_game_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameResponse); i {
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
		file_proto_services_game_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Game); i {
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
		file_proto_services_game_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameRequest); i {
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
		file_proto_services_game_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_proto_services_game_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifiedBy); i {
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
			RawDescriptor: file_proto_services_game_game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_services_game_game_proto_goTypes,
		DependencyIndexes: file_proto_services_game_game_proto_depIdxs,
		MessageInfos:      file_proto_services_game_game_proto_msgTypes,
	}.Build()
	File_proto_services_game_game_proto = out.File
	file_proto_services_game_game_proto_rawDesc = nil
	file_proto_services_game_game_proto_goTypes = nil
	file_proto_services_game_game_proto_depIdxs = nil
}
