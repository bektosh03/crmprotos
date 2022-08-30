// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: teacherpb/teacher.proto

package teacherpb

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

type Teacher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName   string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName    string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	SubjectId   string `protobuf:"bytes,6,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
}

func (x *Teacher) Reset() {
	*x = Teacher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacherpb_teacher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Teacher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Teacher) ProtoMessage() {}

func (x *Teacher) ProtoReflect() protoreflect.Message {
	mi := &file_teacherpb_teacher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Teacher.ProtoReflect.Descriptor instead.
func (*Teacher) Descriptor() ([]byte, []int) {
	return file_teacherpb_teacher_proto_rawDescGZIP(), []int{0}
}

func (x *Teacher) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Teacher) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Teacher) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Teacher) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Teacher) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Teacher) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

type RegisterTeacherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName   string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName    string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	SubjectId   string `protobuf:"bytes,5,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
}

func (x *RegisterTeacherRequest) Reset() {
	*x = RegisterTeacherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacherpb_teacher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterTeacherRequest) ProtoMessage() {}

func (x *RegisterTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacherpb_teacher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterTeacherRequest.ProtoReflect.Descriptor instead.
func (*RegisterTeacherRequest) Descriptor() ([]byte, []int) {
	return file_teacherpb_teacher_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterTeacherRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *RegisterTeacherRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *RegisterTeacherRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterTeacherRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *RegisterTeacherRequest) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

type Subject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Subject) Reset() {
	*x = Subject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacherpb_teacher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subject) ProtoMessage() {}

func (x *Subject) ProtoReflect() protoreflect.Message {
	mi := &file_teacherpb_teacher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subject.ProtoReflect.Descriptor instead.
func (*Subject) Descriptor() ([]byte, []int) {
	return file_teacherpb_teacher_proto_rawDescGZIP(), []int{2}
}

func (x *Subject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Subject) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Subject) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateSubjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateSubjectRequest) Reset() {
	*x = CreateSubjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacherpb_teacher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSubjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubjectRequest) ProtoMessage() {}

func (x *CreateSubjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacherpb_teacher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubjectRequest.ProtoReflect.Descriptor instead.
func (*CreateSubjectRequest) Descriptor() ([]byte, []int) {
	return file_teacherpb_teacher_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSubjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSubjectRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetTeacherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeacherId string `protobuf:"bytes,1,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"`
}

func (x *GetTeacherRequest) Reset() {
	*x = GetTeacherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacherpb_teacher_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTeacherRequest) ProtoMessage() {}

func (x *GetTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacherpb_teacher_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTeacherRequest.ProtoReflect.Descriptor instead.
func (*GetTeacherRequest) Descriptor() ([]byte, []int) {
	return file_teacherpb_teacher_proto_rawDescGZIP(), []int{4}
}

func (x *GetTeacherRequest) GetTeacherId() string {
	if x != nil {
		return x.TeacherId
	}
	return ""
}

type GetSubjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubjectId string `protobuf:"bytes,1,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
}

func (x *GetSubjectRequest) Reset() {
	*x = GetSubjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacherpb_teacher_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubjectRequest) ProtoMessage() {}

func (x *GetSubjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacherpb_teacher_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubjectRequest.ProtoReflect.Descriptor instead.
func (*GetSubjectRequest) Descriptor() ([]byte, []int) {
	return file_teacherpb_teacher_proto_rawDescGZIP(), []int{5}
}

func (x *GetSubjectRequest) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

type DeleteTeacherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeacherId string `protobuf:"bytes,1,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"`
}

func (x *DeleteTeacherRequest) Reset() {
	*x = DeleteTeacherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacherpb_teacher_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTeacherRequest) ProtoMessage() {}

func (x *DeleteTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacherpb_teacher_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTeacherRequest.ProtoReflect.Descriptor instead.
func (*DeleteTeacherRequest) Descriptor() ([]byte, []int) {
	return file_teacherpb_teacher_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteTeacherRequest) GetTeacherId() string {
	if x != nil {
		return x.TeacherId
	}
	return ""
}

type DeleteSubjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubjectId string `protobuf:"bytes,1,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
}

func (x *DeleteSubjectRequest) Reset() {
	*x = DeleteSubjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacherpb_teacher_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSubjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSubjectRequest) ProtoMessage() {}

func (x *DeleteSubjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacherpb_teacher_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSubjectRequest.ProtoReflect.Descriptor instead.
func (*DeleteSubjectRequest) Descriptor() ([]byte, []int) {
	return file_teacherpb_teacher_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteSubjectRequest) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

type ListTeachersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListTeachersRequest) Reset() {
	*x = ListTeachersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacherpb_teacher_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTeachersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTeachersRequest) ProtoMessage() {}

func (x *ListTeachersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacherpb_teacher_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTeachersRequest.ProtoReflect.Descriptor instead.
func (*ListTeachersRequest) Descriptor() ([]byte, []int) {
	return file_teacherpb_teacher_proto_rawDescGZIP(), []int{8}
}

func (x *ListTeachersRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListTeachersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListTeachersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Teachers []*Teacher `protobuf:"bytes,1,rep,name=teachers,proto3" json:"teachers,omitempty"`
}

func (x *ListTeachersResponse) Reset() {
	*x = ListTeachersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacherpb_teacher_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTeachersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTeachersResponse) ProtoMessage() {}

func (x *ListTeachersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teacherpb_teacher_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTeachersResponse.ProtoReflect.Descriptor instead.
func (*ListTeachersResponse) Descriptor() ([]byte, []int) {
	return file_teacherpb_teacher_proto_rawDescGZIP(), []int{9}
}

func (x *ListTeachersResponse) GetTeachers() []*Teacher {
	if x != nil {
		return x.Teachers
	}
	return nil
}

type ListSubjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit  int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *ListSubjectsRequest) Reset() {
	*x = ListSubjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacherpb_teacher_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubjectsRequest) ProtoMessage() {}

func (x *ListSubjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacherpb_teacher_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubjectsRequest.ProtoReflect.Descriptor instead.
func (*ListSubjectsRequest) Descriptor() ([]byte, []int) {
	return file_teacherpb_teacher_proto_rawDescGZIP(), []int{10}
}

func (x *ListSubjectsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListSubjectsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListSubjectsRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type ListSubjectsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subjects []*Subject `protobuf:"bytes,1,rep,name=subjects,proto3" json:"subjects,omitempty"`
}

func (x *ListSubjectsResponse) Reset() {
	*x = ListSubjectsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacherpb_teacher_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubjectsResponse) ProtoMessage() {}

func (x *ListSubjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teacherpb_teacher_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubjectsResponse.ProtoReflect.Descriptor instead.
func (*ListSubjectsResponse) Descriptor() ([]byte, []int) {
	return file_teacherpb_teacher_proto_rawDescGZIP(), []int{11}
}

func (x *ListSubjectsResponse) GetSubjects() []*Subject {
	if x != nil {
		return x.Subjects
	}
	return nil
}

var File_teacherpb_teacher_proto protoreflect.FileDescriptor

var file_teacherpb_teacher_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x74, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xad, 0x01, 0x0a, 0x07, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x22, 0xac, 0x01, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x22, 0x4f, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x4c, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x32, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x22, 0x35,
	0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x46, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e,
	0x0a, 0x08, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x54, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x52, 0x08, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x22, 0x57,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x46, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2e, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x53, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x08, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x32,
	0xd6, 0x04, 0x0a, 0x0e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x48, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x70, 0x62, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1f, 0x2e,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x1c, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x12, 0x3e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x1c, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x48, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x48, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1f, 0x2e,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4f, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x1e, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x6b, 0x74, 0x6f, 0x73, 0x68, 0x30, 0x33,
	0x2f, 0x63, 0x72, 0x6d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teacherpb_teacher_proto_rawDescOnce sync.Once
	file_teacherpb_teacher_proto_rawDescData = file_teacherpb_teacher_proto_rawDesc
)

func file_teacherpb_teacher_proto_rawDescGZIP() []byte {
	file_teacherpb_teacher_proto_rawDescOnce.Do(func() {
		file_teacherpb_teacher_proto_rawDescData = protoimpl.X.CompressGZIP(file_teacherpb_teacher_proto_rawDescData)
	})
	return file_teacherpb_teacher_proto_rawDescData
}

var file_teacherpb_teacher_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_teacherpb_teacher_proto_goTypes = []interface{}{
	(*Teacher)(nil),                // 0: teacherpb.Teacher
	(*RegisterTeacherRequest)(nil), // 1: teacherpb.RegisterTeacherRequest
	(*Subject)(nil),                // 2: teacherpb.Subject
	(*CreateSubjectRequest)(nil),   // 3: teacherpb.CreateSubjectRequest
	(*GetTeacherRequest)(nil),      // 4: teacherpb.GetTeacherRequest
	(*GetSubjectRequest)(nil),      // 5: teacherpb.GetSubjectRequest
	(*DeleteTeacherRequest)(nil),   // 6: teacherpb.DeleteTeacherRequest
	(*DeleteSubjectRequest)(nil),   // 7: teacherpb.DeleteSubjectRequest
	(*ListTeachersRequest)(nil),    // 8: teacherpb.ListTeachersRequest
	(*ListTeachersResponse)(nil),   // 9: teacherpb.ListTeachersResponse
	(*ListSubjectsRequest)(nil),    // 10: teacherpb.ListSubjectsRequest
	(*ListSubjectsResponse)(nil),   // 11: teacherpb.ListSubjectsResponse
	(*emptypb.Empty)(nil),          // 12: google.protobuf.Empty
}
var file_teacherpb_teacher_proto_depIdxs = []int32{
	0,  // 0: teacherpb.ListTeachersResponse.teachers:type_name -> teacherpb.Teacher
	2,  // 1: teacherpb.ListSubjectsResponse.subjects:type_name -> teacherpb.Subject
	1,  // 2: teacherpb.TeacherService.RegisterTeacher:input_type -> teacherpb.RegisterTeacherRequest
	3,  // 3: teacherpb.TeacherService.CreateSubject:input_type -> teacherpb.CreateSubjectRequest
	4,  // 4: teacherpb.TeacherService.GetTeacher:input_type -> teacherpb.GetTeacherRequest
	5,  // 5: teacherpb.TeacherService.GetSubject:input_type -> teacherpb.GetSubjectRequest
	6,  // 6: teacherpb.TeacherService.DeleteTeacher:input_type -> teacherpb.DeleteTeacherRequest
	7,  // 7: teacherpb.TeacherService.DeleteSubject:input_type -> teacherpb.DeleteSubjectRequest
	8,  // 8: teacherpb.TeacherService.ListTeachers:input_type -> teacherpb.ListTeachersRequest
	10, // 9: teacherpb.TeacherService.ListSubjects:input_type -> teacherpb.ListSubjectsRequest
	0,  // 10: teacherpb.TeacherService.RegisterTeacher:output_type -> teacherpb.Teacher
	2,  // 11: teacherpb.TeacherService.CreateSubject:output_type -> teacherpb.Subject
	0,  // 12: teacherpb.TeacherService.GetTeacher:output_type -> teacherpb.Teacher
	2,  // 13: teacherpb.TeacherService.GetSubject:output_type -> teacherpb.Subject
	12, // 14: teacherpb.TeacherService.DeleteTeacher:output_type -> google.protobuf.Empty
	12, // 15: teacherpb.TeacherService.DeleteSubject:output_type -> google.protobuf.Empty
	9,  // 16: teacherpb.TeacherService.ListTeachers:output_type -> teacherpb.ListTeachersResponse
	11, // 17: teacherpb.TeacherService.ListSubjects:output_type -> teacherpb.ListSubjectsResponse
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_teacherpb_teacher_proto_init() }
func file_teacherpb_teacher_proto_init() {
	if File_teacherpb_teacher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teacherpb_teacher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Teacher); i {
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
		file_teacherpb_teacher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterTeacherRequest); i {
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
		file_teacherpb_teacher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subject); i {
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
		file_teacherpb_teacher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSubjectRequest); i {
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
		file_teacherpb_teacher_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTeacherRequest); i {
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
		file_teacherpb_teacher_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubjectRequest); i {
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
		file_teacherpb_teacher_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTeacherRequest); i {
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
		file_teacherpb_teacher_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSubjectRequest); i {
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
		file_teacherpb_teacher_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTeachersRequest); i {
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
		file_teacherpb_teacher_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTeachersResponse); i {
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
		file_teacherpb_teacher_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSubjectsRequest); i {
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
		file_teacherpb_teacher_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSubjectsResponse); i {
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
			RawDescriptor: file_teacherpb_teacher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teacherpb_teacher_proto_goTypes,
		DependencyIndexes: file_teacherpb_teacher_proto_depIdxs,
		MessageInfos:      file_teacherpb_teacher_proto_msgTypes,
	}.Build()
	File_teacherpb_teacher_proto = out.File
	file_teacherpb_teacher_proto_rawDesc = nil
	file_teacherpb_teacher_proto_goTypes = nil
	file_teacherpb_teacher_proto_depIdxs = nil
}