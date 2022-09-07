// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: journalpb/journal.proto

package journalpb

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

type CreateJournalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScheduleId string `protobuf:"bytes,1,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
	StudentId  string `protobuf:"bytes,2,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	Attended   bool   `protobuf:"varint,3,opt,name=attended,proto3" json:"attended,omitempty"`
	Mark       int32  `protobuf:"varint,4,opt,name=mark,proto3" json:"mark,omitempty"`
}

func (x *CreateJournalRequest) Reset() {
	*x = CreateJournalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journalpb_journal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJournalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJournalRequest) ProtoMessage() {}

func (x *CreateJournalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_journalpb_journal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJournalRequest.ProtoReflect.Descriptor instead.
func (*CreateJournalRequest) Descriptor() ([]byte, []int) {
	return file_journalpb_journal_proto_rawDescGZIP(), []int{0}
}

func (x *CreateJournalRequest) GetScheduleId() string {
	if x != nil {
		return x.ScheduleId
	}
	return ""
}

func (x *CreateJournalRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *CreateJournalRequest) GetAttended() bool {
	if x != nil {
		return x.Attended
	}
	return false
}

func (x *CreateJournalRequest) GetMark() int32 {
	if x != nil {
		return x.Mark
	}
	return 0
}

type Journal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ScheduleId string `protobuf:"bytes,2,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
	StudentId  string `protobuf:"bytes,3,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	Attended   bool   `protobuf:"varint,4,opt,name=attended,proto3" json:"attended,omitempty"`
	Mark       int32  `protobuf:"varint,5,opt,name=mark,proto3" json:"mark,omitempty"`
}

func (x *Journal) Reset() {
	*x = Journal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journalpb_journal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Journal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Journal) ProtoMessage() {}

func (x *Journal) ProtoReflect() protoreflect.Message {
	mi := &file_journalpb_journal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Journal.ProtoReflect.Descriptor instead.
func (*Journal) Descriptor() ([]byte, []int) {
	return file_journalpb_journal_proto_rawDescGZIP(), []int{1}
}

func (x *Journal) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Journal) GetScheduleId() string {
	if x != nil {
		return x.ScheduleId
	}
	return ""
}

func (x *Journal) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *Journal) GetAttended() bool {
	if x != nil {
		return x.Attended
	}
	return false
}

func (x *Journal) GetMark() int32 {
	if x != nil {
		return x.Mark
	}
	return 0
}

type GetJournalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JournalId string `protobuf:"bytes,1,opt,name=journal_id,json=journalId,proto3" json:"journal_id,omitempty"`
}

func (x *GetJournalRequest) Reset() {
	*x = GetJournalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journalpb_journal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJournalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJournalRequest) ProtoMessage() {}

func (x *GetJournalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_journalpb_journal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJournalRequest.ProtoReflect.Descriptor instead.
func (*GetJournalRequest) Descriptor() ([]byte, []int) {
	return file_journalpb_journal_proto_rawDescGZIP(), []int{2}
}

func (x *GetJournalRequest) GetJournalId() string {
	if x != nil {
		return x.JournalId
	}
	return ""
}

type DeleteJournalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JournalId string `protobuf:"bytes,1,opt,name=journal_id,json=journalId,proto3" json:"journal_id,omitempty"`
}

func (x *DeleteJournalRequest) Reset() {
	*x = DeleteJournalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journalpb_journal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJournalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJournalRequest) ProtoMessage() {}

func (x *DeleteJournalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_journalpb_journal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJournalRequest.ProtoReflect.Descriptor instead.
func (*DeleteJournalRequest) Descriptor() ([]byte, []int) {
	return file_journalpb_journal_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteJournalRequest) GetJournalId() string {
	if x != nil {
		return x.JournalId
	}
	return ""
}

var File_journalpb_journal_proto protoreflect.FileDescriptor

var file_journalpb_journal_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x70, 0x62, 0x2f, 0x6a, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6a, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6c, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x86, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x74,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x74,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x89, 0x01, 0x0a, 0x07, 0x4a,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64,
	0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x32, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x75,
	0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6a,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x14, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x49,
	0x64, 0x32, 0x99, 0x02, 0x0a, 0x0e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x1f, 0x2e, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x70,
	0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c,
	0x70, 0x62, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x3e, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x1c, 0x2e, 0x6a, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6c, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c,
	0x70, 0x62, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x37, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x12, 0x2e, 0x6a, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6c, 0x70, 0x62, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x1a,
	0x12, 0x2e, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x70, 0x62, 0x2e, 0x4a, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6c, 0x12, 0x48, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x75,
	0x72, 0x6e, 0x61, 0x6c, 0x12, 0x1f, 0x2e, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x70, 0x62,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x2a, 0x5a,
	0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x6b, 0x74,
	0x6f, 0x73, 0x68, 0x30, 0x33, 0x2f, 0x63, 0x72, 0x6d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_journalpb_journal_proto_rawDescOnce sync.Once
	file_journalpb_journal_proto_rawDescData = file_journalpb_journal_proto_rawDesc
)

func file_journalpb_journal_proto_rawDescGZIP() []byte {
	file_journalpb_journal_proto_rawDescOnce.Do(func() {
		file_journalpb_journal_proto_rawDescData = protoimpl.X.CompressGZIP(file_journalpb_journal_proto_rawDescData)
	})
	return file_journalpb_journal_proto_rawDescData
}

var file_journalpb_journal_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_journalpb_journal_proto_goTypes = []interface{}{
	(*CreateJournalRequest)(nil), // 0: journalpb.CreateJournalRequest
	(*Journal)(nil),              // 1: journalpb.Journal
	(*GetJournalRequest)(nil),    // 2: journalpb.GetJournalRequest
	(*DeleteJournalRequest)(nil), // 3: journalpb.DeleteJournalRequest
	(*emptypb.Empty)(nil),        // 4: google.protobuf.Empty
}
var file_journalpb_journal_proto_depIdxs = []int32{
	0, // 0: journalpb.JournalService.CreateJournal:input_type -> journalpb.CreateJournalRequest
	2, // 1: journalpb.JournalService.GetJournal:input_type -> journalpb.GetJournalRequest
	1, // 2: journalpb.JournalService.UpdateJournal:input_type -> journalpb.Journal
	3, // 3: journalpb.JournalService.DeleteJournal:input_type -> journalpb.DeleteJournalRequest
	1, // 4: journalpb.JournalService.CreateJournal:output_type -> journalpb.Journal
	1, // 5: journalpb.JournalService.GetJournal:output_type -> journalpb.Journal
	1, // 6: journalpb.JournalService.UpdateJournal:output_type -> journalpb.Journal
	4, // 7: journalpb.JournalService.DeleteJournal:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_journalpb_journal_proto_init() }
func file_journalpb_journal_proto_init() {
	if File_journalpb_journal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_journalpb_journal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJournalRequest); i {
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
		file_journalpb_journal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Journal); i {
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
		file_journalpb_journal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJournalRequest); i {
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
		file_journalpb_journal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteJournalRequest); i {
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
			RawDescriptor: file_journalpb_journal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_journalpb_journal_proto_goTypes,
		DependencyIndexes: file_journalpb_journal_proto_depIdxs,
		MessageInfos:      file_journalpb_journal_proto_msgTypes,
	}.Build()
	File_journalpb_journal_proto = out.File
	file_journalpb_journal_proto_rawDesc = nil
	file_journalpb_journal_proto_goTypes = nil
	file_journalpb_journal_proto_depIdxs = nil
}
