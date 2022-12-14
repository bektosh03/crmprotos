// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: schedulepb/schedule.proto

package schedulepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Weekday int32

const (
	Weekday_MONDAY    Weekday = 0
	Weekday_TUESDAY   Weekday = 1
	Weekday_WEDNESDAY Weekday = 2
	Weekday_THURSDAY  Weekday = 3
	Weekday_FRIDAY    Weekday = 4
	Weekday_SATURDAY  Weekday = 5
	Weekday_SUNDAY    Weekday = 6
)

// Enum value maps for Weekday.
var (
	Weekday_name = map[int32]string{
		0: "MONDAY",
		1: "TUESDAY",
		2: "WEDNESDAY",
		3: "THURSDAY",
		4: "FRIDAY",
		5: "SATURDAY",
		6: "SUNDAY",
	}
	Weekday_value = map[string]int32{
		"MONDAY":    0,
		"TUESDAY":   1,
		"WEDNESDAY": 2,
		"THURSDAY":  3,
		"FRIDAY":    4,
		"SATURDAY":  5,
		"SUNDAY":    6,
	}
)

func (x Weekday) Enum() *Weekday {
	p := new(Weekday)
	*p = x
	return p
}

func (x Weekday) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Weekday) Descriptor() protoreflect.EnumDescriptor {
	return file_schedulepb_schedule_proto_enumTypes[0].Descriptor()
}

func (Weekday) Type() protoreflect.EnumType {
	return &file_schedulepb_schedule_proto_enumTypes[0]
}

func (x Weekday) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Weekday.Descriptor instead.
func (Weekday) EnumDescriptor() ([]byte, []int) {
	return file_schedulepb_schedule_proto_rawDescGZIP(), []int{0}
}

type CreateScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId      string  `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	SubjectId    string  `protobuf:"bytes,2,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	TeacherId    string  `protobuf:"bytes,3,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"`
	Weekday      Weekday `protobuf:"varint,4,opt,name=weekday,proto3,enum=schedulepb.Weekday" json:"weekday,omitempty"`
	LessonNumber int32   `protobuf:"varint,5,opt,name=lesson_number,json=lessonNumber,proto3" json:"lesson_number,omitempty"`
}

func (x *CreateScheduleRequest) Reset() {
	*x = CreateScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedulepb_schedule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScheduleRequest) ProtoMessage() {}

func (x *CreateScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schedulepb_schedule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScheduleRequest.ProtoReflect.Descriptor instead.
func (*CreateScheduleRequest) Descriptor() ([]byte, []int) {
	return file_schedulepb_schedule_proto_rawDescGZIP(), []int{0}
}

func (x *CreateScheduleRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *CreateScheduleRequest) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

func (x *CreateScheduleRequest) GetTeacherId() string {
	if x != nil {
		return x.TeacherId
	}
	return ""
}

func (x *CreateScheduleRequest) GetWeekday() Weekday {
	if x != nil {
		return x.Weekday
	}
	return Weekday_MONDAY
}

func (x *CreateScheduleRequest) GetLessonNumber() int32 {
	if x != nil {
		return x.LessonNumber
	}
	return 0
}

type Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	GroupId      string  `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	SubjectId    string  `protobuf:"bytes,3,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	TeacherId    string  `protobuf:"bytes,4,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"`
	Weekday      Weekday `protobuf:"varint,5,opt,name=weekday,proto3,enum=schedulepb.Weekday" json:"weekday,omitempty"`
	LessonNumber int32   `protobuf:"varint,6,opt,name=lesson_number,json=lessonNumber,proto3" json:"lesson_number,omitempty"`
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedulepb_schedule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_schedulepb_schedule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_schedulepb_schedule_proto_rawDescGZIP(), []int{1}
}

func (x *Schedule) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Schedule) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *Schedule) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

func (x *Schedule) GetTeacherId() string {
	if x != nil {
		return x.TeacherId
	}
	return ""
}

func (x *Schedule) GetWeekday() Weekday {
	if x != nil {
		return x.Weekday
	}
	return Weekday_MONDAY
}

func (x *Schedule) GetLessonNumber() int32 {
	if x != nil {
		return x.LessonNumber
	}
	return 0
}

type GetScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScheduleId string `protobuf:"bytes,1,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
}

func (x *GetScheduleRequest) Reset() {
	*x = GetScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedulepb_schedule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScheduleRequest) ProtoMessage() {}

func (x *GetScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schedulepb_schedule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScheduleRequest.ProtoReflect.Descriptor instead.
func (*GetScheduleRequest) Descriptor() ([]byte, []int) {
	return file_schedulepb_schedule_proto_rawDescGZIP(), []int{2}
}

func (x *GetScheduleRequest) GetScheduleId() string {
	if x != nil {
		return x.ScheduleId
	}
	return ""
}

type DeleteScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScheduleId string `protobuf:"bytes,1,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
}

func (x *DeleteScheduleRequest) Reset() {
	*x = DeleteScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedulepb_schedule_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteScheduleRequest) ProtoMessage() {}

func (x *DeleteScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schedulepb_schedule_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteScheduleRequest.ProtoReflect.Descriptor instead.
func (*DeleteScheduleRequest) Descriptor() ([]byte, []int) {
	return file_schedulepb_schedule_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteScheduleRequest) GetScheduleId() string {
	if x != nil {
		return x.ScheduleId
	}
	return ""
}

type GetFullScheduleForGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *GetFullScheduleForGroupRequest) Reset() {
	*x = GetFullScheduleForGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedulepb_schedule_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFullScheduleForGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFullScheduleForGroupRequest) ProtoMessage() {}

func (x *GetFullScheduleForGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schedulepb_schedule_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFullScheduleForGroupRequest.ProtoReflect.Descriptor instead.
func (*GetFullScheduleForGroupRequest) Descriptor() ([]byte, []int) {
	return file_schedulepb_schedule_proto_rawDescGZIP(), []int{4}
}

func (x *GetFullScheduleForGroupRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type GetSpecificDateScheduleForGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string                 `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Date    *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *GetSpecificDateScheduleForGroupRequest) Reset() {
	*x = GetSpecificDateScheduleForGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedulepb_schedule_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSpecificDateScheduleForGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSpecificDateScheduleForGroupRequest) ProtoMessage() {}

func (x *GetSpecificDateScheduleForGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schedulepb_schedule_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSpecificDateScheduleForGroupRequest.ProtoReflect.Descriptor instead.
func (*GetSpecificDateScheduleForGroupRequest) Descriptor() ([]byte, []int) {
	return file_schedulepb_schedule_proto_rawDescGZIP(), []int{5}
}

func (x *GetSpecificDateScheduleForGroupRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *GetSpecificDateScheduleForGroupRequest) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type ScheduleList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schedules []*Schedule `protobuf:"bytes,1,rep,name=schedules,proto3" json:"schedules,omitempty"`
}

func (x *ScheduleList) Reset() {
	*x = ScheduleList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedulepb_schedule_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleList) ProtoMessage() {}

func (x *ScheduleList) ProtoReflect() protoreflect.Message {
	mi := &file_schedulepb_schedule_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleList.ProtoReflect.Descriptor instead.
func (*ScheduleList) Descriptor() ([]byte, []int) {
	return file_schedulepb_schedule_proto_rawDescGZIP(), []int{6}
}

func (x *ScheduleList) GetSchedules() []*Schedule {
	if x != nil {
		return x.Schedules
	}
	return nil
}

type GetFullScheduleForTeacherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeacherId string `protobuf:"bytes,1,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"`
}

func (x *GetFullScheduleForTeacherRequest) Reset() {
	*x = GetFullScheduleForTeacherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedulepb_schedule_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFullScheduleForTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFullScheduleForTeacherRequest) ProtoMessage() {}

func (x *GetFullScheduleForTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schedulepb_schedule_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFullScheduleForTeacherRequest.ProtoReflect.Descriptor instead.
func (*GetFullScheduleForTeacherRequest) Descriptor() ([]byte, []int) {
	return file_schedulepb_schedule_proto_rawDescGZIP(), []int{7}
}

func (x *GetFullScheduleForTeacherRequest) GetTeacherId() string {
	if x != nil {
		return x.TeacherId
	}
	return ""
}

type GetSpecificDateScheduleForTeacherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeacherId string                 `protobuf:"bytes,1,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"`
	Date      *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *GetSpecificDateScheduleForTeacherRequest) Reset() {
	*x = GetSpecificDateScheduleForTeacherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedulepb_schedule_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSpecificDateScheduleForTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSpecificDateScheduleForTeacherRequest) ProtoMessage() {}

func (x *GetSpecificDateScheduleForTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schedulepb_schedule_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSpecificDateScheduleForTeacherRequest.ProtoReflect.Descriptor instead.
func (*GetSpecificDateScheduleForTeacherRequest) Descriptor() ([]byte, []int) {
	return file_schedulepb_schedule_proto_rawDescGZIP(), []int{8}
}

func (x *GetSpecificDateScheduleForTeacherRequest) GetTeacherId() string {
	if x != nil {
		return x.TeacherId
	}
	return ""
}

func (x *GetSpecificDateScheduleForTeacherRequest) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

var File_schedulepb_schedule_proto protoreflect.FileDescriptor

var file_schedulepb_schedule_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x70, 0x62, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x07, 0x77, 0x65, 0x65, 0x6b,
	0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x64, 0x61, 0x79, 0x52, 0x07,
	0x77, 0x65, 0x65, 0x6b, 0x64, 0x61, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xc7, 0x01, 0x0a,
	0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x2d, 0x0a, 0x07, 0x77, 0x65, 0x65, 0x6b, 0x64, 0x61, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x70, 0x62,
	0x2e, 0x57, 0x65, 0x65, 0x6b, 0x64, 0x61, 0x79, 0x52, 0x07, 0x77, 0x65, 0x65, 0x6b, 0x64, 0x61,
	0x79, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x35, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x38, 0x0a,
	0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x46, 0x75,
	0x6c, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x22, 0x73, 0x0a, 0x26, 0x47, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x44, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x46,
	0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x42, 0x0a, 0x0c, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x09, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x52, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x41, 0x0a,
	0x20, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x46, 0x6f, 0x72, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x79, 0x0a, 0x28, 0x47, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x44,
	0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x2a, 0x65, 0x0a, 0x07, 0x57,
	0x65, 0x65, 0x6b, 0x64, 0x61, 0x79, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x4f, 0x4e, 0x44, 0x41, 0x59,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x55, 0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x01, 0x12,
	0x0d, 0x0a, 0x09, 0x57, 0x45, 0x44, 0x4e, 0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x02, 0x12, 0x0c,
	0x0a, 0x08, 0x54, 0x48, 0x55, 0x52, 0x53, 0x44, 0x41, 0x59, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06,
	0x46, 0x52, 0x49, 0x44, 0x41, 0x59, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x41, 0x54, 0x55,
	0x52, 0x44, 0x41, 0x59, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x55, 0x4e, 0x44, 0x41, 0x59,
	0x10, 0x06, 0x32, 0xd8, 0x05, 0x0a, 0x0f, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x21, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x12, 0x43, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x12, 0x1e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x1a, 0x14,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x12, 0x4b, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x21, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x5f, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2a, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c,
	0x6c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x6f, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x44, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x46, 0x6f, 0x72,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x32, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x44, 0x61,
	0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x63, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x2c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x46, 0x75, 0x6c, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x46, 0x6f, 0x72,
	0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x73, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x53,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x44, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x34, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x44, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x46, 0x6f, 0x72, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x70, 0x62,
	0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x2b, 0x5a,
	0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x6b, 0x74,
	0x6f, 0x73, 0x68, 0x30, 0x33, 0x2f, 0x63, 0x72, 0x6d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_schedulepb_schedule_proto_rawDescOnce sync.Once
	file_schedulepb_schedule_proto_rawDescData = file_schedulepb_schedule_proto_rawDesc
)

func file_schedulepb_schedule_proto_rawDescGZIP() []byte {
	file_schedulepb_schedule_proto_rawDescOnce.Do(func() {
		file_schedulepb_schedule_proto_rawDescData = protoimpl.X.CompressGZIP(file_schedulepb_schedule_proto_rawDescData)
	})
	return file_schedulepb_schedule_proto_rawDescData
}

var file_schedulepb_schedule_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_schedulepb_schedule_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_schedulepb_schedule_proto_goTypes = []interface{}{
	(Weekday)(0),                                     // 0: schedulepb.Weekday
	(*CreateScheduleRequest)(nil),                    // 1: schedulepb.CreateScheduleRequest
	(*Schedule)(nil),                                 // 2: schedulepb.Schedule
	(*GetScheduleRequest)(nil),                       // 3: schedulepb.GetScheduleRequest
	(*DeleteScheduleRequest)(nil),                    // 4: schedulepb.DeleteScheduleRequest
	(*GetFullScheduleForGroupRequest)(nil),           // 5: schedulepb.GetFullScheduleForGroupRequest
	(*GetSpecificDateScheduleForGroupRequest)(nil),   // 6: schedulepb.GetSpecificDateScheduleForGroupRequest
	(*ScheduleList)(nil),                             // 7: schedulepb.ScheduleList
	(*GetFullScheduleForTeacherRequest)(nil),         // 8: schedulepb.GetFullScheduleForTeacherRequest
	(*GetSpecificDateScheduleForTeacherRequest)(nil), // 9: schedulepb.GetSpecificDateScheduleForTeacherRequest
	(*timestamppb.Timestamp)(nil),                    // 10: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                            // 11: google.protobuf.Empty
}
var file_schedulepb_schedule_proto_depIdxs = []int32{
	0,  // 0: schedulepb.CreateScheduleRequest.weekday:type_name -> schedulepb.Weekday
	0,  // 1: schedulepb.Schedule.weekday:type_name -> schedulepb.Weekday
	10, // 2: schedulepb.GetSpecificDateScheduleForGroupRequest.date:type_name -> google.protobuf.Timestamp
	2,  // 3: schedulepb.ScheduleList.schedules:type_name -> schedulepb.Schedule
	10, // 4: schedulepb.GetSpecificDateScheduleForTeacherRequest.date:type_name -> google.protobuf.Timestamp
	1,  // 5: schedulepb.ScheduleService.CreateSchedule:input_type -> schedulepb.CreateScheduleRequest
	3,  // 6: schedulepb.ScheduleService.GetSchedule:input_type -> schedulepb.GetScheduleRequest
	2,  // 7: schedulepb.ScheduleService.UpdateSchedule:input_type -> schedulepb.Schedule
	4,  // 8: schedulepb.ScheduleService.DeleteSchedule:input_type -> schedulepb.DeleteScheduleRequest
	5,  // 9: schedulepb.ScheduleService.GetFullScheduleForGroup:input_type -> schedulepb.GetFullScheduleForGroupRequest
	6,  // 10: schedulepb.ScheduleService.GetSpecificDateScheduleForGroup:input_type -> schedulepb.GetSpecificDateScheduleForGroupRequest
	8,  // 11: schedulepb.ScheduleService.GetFullScheduleForTeacher:input_type -> schedulepb.GetFullScheduleForTeacherRequest
	9,  // 12: schedulepb.ScheduleService.GetSpecificDateScheduleForTeacher:input_type -> schedulepb.GetSpecificDateScheduleForTeacherRequest
	2,  // 13: schedulepb.ScheduleService.CreateSchedule:output_type -> schedulepb.Schedule
	2,  // 14: schedulepb.ScheduleService.GetSchedule:output_type -> schedulepb.Schedule
	2,  // 15: schedulepb.ScheduleService.UpdateSchedule:output_type -> schedulepb.Schedule
	11, // 16: schedulepb.ScheduleService.DeleteSchedule:output_type -> google.protobuf.Empty
	7,  // 17: schedulepb.ScheduleService.GetFullScheduleForGroup:output_type -> schedulepb.ScheduleList
	7,  // 18: schedulepb.ScheduleService.GetSpecificDateScheduleForGroup:output_type -> schedulepb.ScheduleList
	7,  // 19: schedulepb.ScheduleService.GetFullScheduleForTeacher:output_type -> schedulepb.ScheduleList
	7,  // 20: schedulepb.ScheduleService.GetSpecificDateScheduleForTeacher:output_type -> schedulepb.ScheduleList
	13, // [13:21] is the sub-list for method output_type
	5,  // [5:13] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_schedulepb_schedule_proto_init() }
func file_schedulepb_schedule_proto_init() {
	if File_schedulepb_schedule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schedulepb_schedule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateScheduleRequest); i {
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
		file_schedulepb_schedule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schedule); i {
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
		file_schedulepb_schedule_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetScheduleRequest); i {
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
		file_schedulepb_schedule_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteScheduleRequest); i {
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
		file_schedulepb_schedule_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFullScheduleForGroupRequest); i {
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
		file_schedulepb_schedule_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSpecificDateScheduleForGroupRequest); i {
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
		file_schedulepb_schedule_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleList); i {
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
		file_schedulepb_schedule_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFullScheduleForTeacherRequest); i {
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
		file_schedulepb_schedule_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSpecificDateScheduleForTeacherRequest); i {
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
			RawDescriptor: file_schedulepb_schedule_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_schedulepb_schedule_proto_goTypes,
		DependencyIndexes: file_schedulepb_schedule_proto_depIdxs,
		EnumInfos:         file_schedulepb_schedule_proto_enumTypes,
		MessageInfos:      file_schedulepb_schedule_proto_msgTypes,
	}.Build()
	File_schedulepb_schedule_proto = out.File
	file_schedulepb_schedule_proto_rawDesc = nil
	file_schedulepb_schedule_proto_goTypes = nil
	file_schedulepb_schedule_proto_depIdxs = nil
}
