// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: habitTracker.proto

package generated

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

type HabitId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HabbitId string `protobuf:"bytes,1,opt,name=HabbitId,proto3" json:"HabbitId,omitempty"`
}

func (x *HabitId) Reset() {
	*x = HabitId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_habitTracker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HabitId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HabitId) ProtoMessage() {}

func (x *HabitId) ProtoReflect() protoreflect.Message {
	mi := &file_habitTracker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HabitId.ProtoReflect.Descriptor instead.
func (*HabitId) Descriptor() ([]byte, []int) {
	return file_habitTracker_proto_rawDescGZIP(), []int{0}
}

func (x *HabitId) GetHabbitId() string {
	if x != nil {
		return x.HabbitId
	}
	return ""
}

type Habit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HabbitId    string `protobuf:"bytes,1,opt,name=habbitId,proto3" json:"habbitId,omitempty"`
	UserId      string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Discription string `protobuf:"bytes,4,opt,name=discription,proto3" json:"discription,omitempty"`
	Frequency   string `protobuf:"bytes,5,opt,name=frequency,proto3" json:"frequency,omitempty"`
	CreatedAt   string `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Habit) Reset() {
	*x = Habit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_habitTracker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Habit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Habit) ProtoMessage() {}

func (x *Habit) ProtoReflect() protoreflect.Message {
	mi := &file_habitTracker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Habit.ProtoReflect.Descriptor instead.
func (*Habit) Descriptor() ([]byte, []int) {
	return file_habitTracker_proto_rawDescGZIP(), []int{1}
}

func (x *Habit) GetHabbitId() string {
	if x != nil {
		return x.HabbitId
	}
	return ""
}

func (x *Habit) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Habit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Habit) GetDiscription() string {
	if x != nil {
		return x.Discription
	}
	return ""
}

func (x *Habit) GetFrequency() string {
	if x != nil {
		return x.Frequency
	}
	return ""
}

func (x *Habit) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type HabitLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	HabbitId string `protobuf:"bytes,2,opt,name=habbitId,proto3" json:"habbitId,omitempty"`
	LoggedAt string `protobuf:"bytes,3,opt,name=loggedAt,proto3" json:"loggedAt,omitempty"`
	Notes    string `protobuf:"bytes,4,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (x *HabitLog) Reset() {
	*x = HabitLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_habitTracker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HabitLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HabitLog) ProtoMessage() {}

func (x *HabitLog) ProtoReflect() protoreflect.Message {
	mi := &file_habitTracker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HabitLog.ProtoReflect.Descriptor instead.
func (*HabitLog) Descriptor() ([]byte, []int) {
	return file_habitTracker_proto_rawDescGZIP(), []int{2}
}

func (x *HabitLog) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *HabitLog) GetHabbitId() string {
	if x != nil {
		return x.HabbitId
	}
	return ""
}

func (x *HabitLog) GetLoggedAt() string {
	if x != nil {
		return x.LoggedAt
	}
	return ""
}

func (x *HabitLog) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_habitTracker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_habitTracker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_habitTracker_proto_rawDescGZIP(), []int{3}
}

func (x *Status) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type UserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UserId) Reset() {
	*x = UserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_habitTracker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserId) ProtoMessage() {}

func (x *UserId) ProtoReflect() protoreflect.Message {
	mi := &file_habitTracker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserId.ProtoReflect.Descriptor instead.
func (*UserId) Descriptor() ([]byte, []int) {
	return file_habitTracker_proto_rawDescGZIP(), []int{4}
}

func (x *UserId) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Req) Reset() {
	*x = Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_habitTracker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req) ProtoMessage() {}

func (x *Req) ProtoReflect() protoreflect.Message {
	mi := &file_habitTracker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req.ProtoReflect.Descriptor instead.
func (*Req) Descriptor() ([]byte, []int) {
	return file_habitTracker_proto_rawDescGZIP(), []int{5}
}

type UserHabits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Habbits []*Habit `protobuf:"bytes,1,rep,name=habbits,proto3" json:"habbits,omitempty"`

func (x *UserHabits) Reset() {
	*x = UserHabits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_habitTracker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserHabits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserHabits) ProtoMessage() {}

func (x *UserHabits) ProtoReflect() protoreflect.Message {
	mi := &file_habitTracker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserHabits.ProtoReflect.Descriptor instead.
func (*UserHabits) Descriptor() ([]byte, []int) {
	return file_habitTracker_proto_rawDescGZIP(), []int{6}
}

func (x *UserHabits) GetHabbits() []*Habit {
	if x != nil {
		return x.Habbits
	}
	return nil
}

type Habits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Habbits []*HabitLog `protobuf:"bytes,1,rep,name=habbits,proto3" json:"habbits,omitempty"`
}

func (x *Habits) Reset() {
	*x = Habits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_habitTracker_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Habits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Habits) ProtoMessage() {}

func (x *Habits) ProtoReflect() protoreflect.Message {
	mi := &file_habitTracker_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Habits.ProtoReflect.Descriptor instead.
func (*Habits) Descriptor() ([]byte, []int) {
	return file_habitTracker_proto_rawDescGZIP(), []int{7}
}

func (x *Habits) GetHabbits() []*HabitLog {

	if x != nil {
		return x.Habbits
	}
	return nil
}

var File_habitTracker_proto protoreflect.FileDescriptor

var file_habitTracker_proto_rawDesc = []byte{
	0x0a, 0x12, 0x68, 0x61, 0x62, 0x69, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x68, 0x61, 0x62, 0x62, 0x69, 0x74, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x22, 0x25, 0x0a, 0x07, 0x48, 0x61, 0x62, 0x69, 0x74, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x48, 0x61, 0x62, 0x62, 0x69, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x48, 0x61, 0x62, 0x62, 0x69, 0x74, 0x49, 0x64, 0x22, 0xad, 0x01, 0x0a, 0x05, 0x48,
	0x61, 0x62, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x61, 0x62, 0x62, 0x69, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x61, 0x62, 0x62, 0x69, 0x74, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x69, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x68, 0x0a, 0x08, 0x48, 0x61,
	0x62, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x61, 0x62, 0x62, 0x69, 0x74,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x61, 0x62, 0x62, 0x69, 0x74,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e,
	0x6f, 0x74, 0x65, 0x73, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x20, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x05, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x22,
	0x3c, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x48, 0x61, 0x62, 0x69, 0x74, 0x73, 0x12, 0x2e, 0x0a,
	0x07, 0x68, 0x61, 0x62, 0x62, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x68, 0x61, 0x62, 0x62, 0x69, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x48,
	0x61, 0x62, 0x69, 0x74, 0x52, 0x07, 0x68, 0x61, 0x62, 0x62, 0x69, 0x74, 0x73, 0x22, 0x3b, 0x0a,
	0x06, 0x48, 0x61, 0x62, 0x69, 0x74, 0x73, 0x12, 0x31, 0x0a, 0x07, 0x68, 0x61, 0x62, 0x62, 0x69,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x68, 0x61, 0x62, 0x62, 0x69,
	0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x48, 0x61, 0x62, 0x69, 0x74, 0x4c, 0x6f,
	0x67, 0x52, 0x07, 0x68, 0x61, 0x62, 0x62, 0x69, 0x74, 0x73, 0x32, 0x91, 0x04, 0x0a, 0x0c, 0x48,
	0x61, 0x62, 0x69, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x61, 0x62, 0x69, 0x74, 0x12, 0x14, 0x2e, 0x68, 0x61, 0x62,
	0x62, 0x69, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x48, 0x61, 0x62, 0x69, 0x74,
	0x1a, 0x15, 0x2e, 0x68, 0x61, 0x62, 0x62, 0x69, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x48, 0x61, 0x62, 0x69, 0x74, 0x12, 0x16, 0x2e, 0x68, 0x61, 0x62, 0x62, 0x69, 0x74, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x48, 0x61, 0x62, 0x69, 0x74, 0x49, 0x64, 0x1a, 0x14, 0x2e,
	0x68, 0x61, 0x62, 0x62, 0x69, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x48, 0x61,
	0x62, 0x69, 0x74, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48,
	0x61, 0x62, 0x69, 0x74, 0x12, 0x14, 0x2e, 0x68, 0x61, 0x62, 0x62, 0x69, 0x74, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x2e, 0x48, 0x61, 0x62, 0x69, 0x74, 0x1a, 0x15, 0x2e, 0x68, 0x61, 0x62,
	0x62, 0x69, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x61, 0x62,
	0x69, 0x74, 0x12, 0x16, 0x2e, 0x68, 0x61, 0x62, 0x62, 0x69, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x72, 0x2e, 0x48, 0x61, 0x62, 0x69, 0x74, 0x49, 0x64, 0x1a, 0x15, 0x2e, 0x68, 0x61, 0x62,
	0x62, 0x69, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x48, 0x61,
	0x62, 0x69, 0x74, 0x73, 0x12, 0x15, 0x2e, 0x68, 0x61, 0x62, 0x62, 0x69, 0x74, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x14, 0x2e, 0x68, 0x61,
	0x62, 0x62, 0x69, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x48, 0x61, 0x62, 0x69,
	0x74, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x61, 0x62,
	0x69, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x17, 0x2e, 0x68, 0x61, 0x62, 0x62, 0x69, 0x74, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x48, 0x61, 0x62, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x1a, 0x15,
	0x2e, 0x68, 0x61, 0x62, 0x62, 0x69, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x48, 0x61,
	0x62, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x16, 0x2e, 0x68, 0x61, 0x62, 0x62, 0x69, 0x74,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x48, 0x61, 0x62, 0x69, 0x74, 0x49, 0x64, 0x1a,
	0x17, 0x2e, 0x68, 0x61, 0x62, 0x62, 0x69, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e,
	0x48, 0x61, 0x62, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x48, 0x61, 0x62, 0x69, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x12, 0x2e, 0x68, 0x61, 0x62, 0x62, 0x69, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x68, 0x61, 0x62, 0x62, 0x69, 0x74, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x48, 0x61, 0x62, 0x69, 0x74, 0x73, 0x22, 0x00, 0x42, 0x0c,
	0x5a, 0x0a, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_habitTracker_proto_rawDescOnce sync.Once
	file_habitTracker_proto_rawDescData = file_habitTracker_proto_rawDesc
)

func file_habitTracker_proto_rawDescGZIP() []byte {
	file_habitTracker_proto_rawDescOnce.Do(func() {
		file_habitTracker_proto_rawDescData = protoimpl.X.CompressGZIP(file_habitTracker_proto_rawDescData)
	})
	return file_habitTracker_proto_rawDescData
}

var file_habitTracker_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_habitTracker_proto_goTypes = []any{
	(*HabitId)(nil),    // 0: habbitTracker.HabitId
	(*Habit)(nil),      // 1: habbitTracker.Habit
	(*HabitLog)(nil),   // 2: habbitTracker.HabitLog
	(*Status)(nil),     // 3: habbitTracker.Status
	(*UserId)(nil),     // 4: habbitTracker.UserId
	(*Req)(nil),        // 5: habbitTracker.Req
	(*UserHabits)(nil), // 6: habbitTracker.UserHabits
	(*Habits)(nil),     // 7: habbitTracker.Habits
}
var file_habitTracker_proto_depIdxs = []int32{
	1,  // 0: habbitTracker.UserHabits.habbits:type_name -> habbitTracker.Habit
	2,  // 1: habbitTracker.Habits.habbits:type_name -> habbitTracker.HabitLog
	1,  // 2: habbitTracker.HabitTracker.CreateHabit:input_type -> habbitTracker.Habit
	0,  // 3: habbitTracker.HabitTracker.GetHabit:input_type -> habbitTracker.HabitId
	1,  // 4: habbitTracker.HabitTracker.UpdateHabit:input_type -> habbitTracker.Habit
	0,  // 5: habbitTracker.HabitTracker.DeleteHabit:input_type -> habbitTracker.HabitId
	4,  // 6: habbitTracker.HabitTracker.GetUserHabits:input_type -> habbitTracker.UserId
	2,  // 7: habbitTracker.HabitTracker.CreateHabitLog:input_type -> habbitTracker.HabitLog
	0,  // 8: habbitTracker.HabitTracker.GetHabitLogs:input_type -> habbitTracker.HabitId
	5,  // 9: habbitTracker.HabitTracker.GetHabitSuggestions:input_type -> habbitTracker.Req
	3,  // 10: habbitTracker.HabitTracker.CreateHabit:output_type -> habbitTracker.Status
	1,  // 11: habbitTracker.HabitTracker.GetHabit:output_type -> habbitTracker.Habit
	3,  // 12: habbitTracker.HabitTracker.UpdateHabit:output_type -> habbitTracker.Status
	3,  // 13: habbitTracker.HabitTracker.DeleteHabit:output_type -> habbitTracker.Status
	1,  // 14: habbitTracker.HabitTracker.GetUserHabits:output_type -> habbitTracker.Habit
	3,  // 15: habbitTracker.HabitTracker.CreateHabitLog:output_type -> habbitTracker.Status
	2,  // 16: habbitTracker.HabitTracker.GetHabitLogs:output_type -> habbitTracker.HabitLog
	7,  // 17: habbitTracker.HabitTracker.GetHabitSuggestions:output_type -> habbitTracker.Habits
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_habitTracker_proto_init() }
func file_habitTracker_proto_init() {
	if File_habitTracker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_habitTracker_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*HabitId); i {
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
		file_habitTracker_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Habit); i {
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
		file_habitTracker_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*HabitLog); i {
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
		file_habitTracker_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Status); i {
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
		file_habitTracker_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UserId); i {
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
		file_habitTracker_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Req); i {
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
		file_habitTracker_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UserHabits); i {
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
		file_habitTracker_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Habits); i {
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
			RawDescriptor: file_habitTracker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_habitTracker_proto_goTypes,
		DependencyIndexes: file_habitTracker_proto_depIdxs,
		MessageInfos:      file_habitTracker_proto_msgTypes,
	}.Build()
	File_habitTracker_proto = out.File
	file_habitTracker_proto_rawDesc = nil
	file_habitTracker_proto_goTypes = nil
	file_habitTracker_proto_depIdxs = nil
}
