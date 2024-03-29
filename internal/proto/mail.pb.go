// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: internal/proto/mail.proto

package notifier

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

type EmailRequest_BodyFormat int32

const (
	EmailRequest__UNSPECIFIED EmailRequest_BodyFormat = 0
	EmailRequest_TEXT         EmailRequest_BodyFormat = 1
	EmailRequest_HTML         EmailRequest_BodyFormat = 2
)

// Enum value maps for EmailRequest_BodyFormat.
var (
	EmailRequest_BodyFormat_name = map[int32]string{
		0: "_UNSPECIFIED",
		1: "TEXT",
		2: "HTML",
	}
	EmailRequest_BodyFormat_value = map[string]int32{
		"_UNSPECIFIED": 0,
		"TEXT":         1,
		"HTML":         2,
	}
)

func (x EmailRequest_BodyFormat) Enum() *EmailRequest_BodyFormat {
	p := new(EmailRequest_BodyFormat)
	*p = x
	return p
}

func (x EmailRequest_BodyFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EmailRequest_BodyFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_proto_mail_proto_enumTypes[0].Descriptor()
}

func (EmailRequest_BodyFormat) Type() protoreflect.EnumType {
	return &file_internal_proto_mail_proto_enumTypes[0]
}

func (x EmailRequest_BodyFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EmailRequest_BodyFormat.Descriptor instead.
func (EmailRequest_BodyFormat) EnumDescriptor() ([]byte, []int) {
	return file_internal_proto_mail_proto_rawDescGZIP(), []int{0, 0}
}

type EmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To       []string                `protobuf:"bytes,1,rep,name=to,proto3" json:"to,omitempty"`
	Subject  string                  `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Body     string                  `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	BodyType EmailRequest_BodyFormat `protobuf:"varint,4,opt,name=body_type,json=bodyType,proto3,enum=notifier.v1.EmailRequest_BodyFormat" json:"body_type,omitempty"` // deprecated
}

func (x *EmailRequest) Reset() {
	*x = EmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_mail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailRequest) ProtoMessage() {}

func (x *EmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_mail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailRequest.ProtoReflect.Descriptor instead.
func (*EmailRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_mail_proto_rawDescGZIP(), []int{0}
}

func (x *EmailRequest) GetTo() []string {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *EmailRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *EmailRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *EmailRequest) GetBodyType() EmailRequest_BodyFormat {
	if x != nil {
		return x.BodyType
	}
	return EmailRequest__UNSPECIFIED
}

type SmsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SmsRequest) Reset() {
	*x = SmsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_mail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmsRequest) ProtoMessage() {}

func (x *SmsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_mail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmsRequest.ProtoReflect.Descriptor instead.
func (*SmsRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_mail_proto_rawDescGZIP(), []int{1}
}

func (x *SmsRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *SmsRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_internal_proto_mail_proto protoreflect.FileDescriptor

var file_internal_proto_mail_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x01, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x12, 0x41, 0x0a, 0x09, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x08, 0x62,
	0x6f, 0x64, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0x32, 0x0a, 0x0a, 0x42, 0x6f, 0x64, 0x79, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x0c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x45, 0x58, 0x54, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x54, 0x4d, 0x4c, 0x10, 0x02, 0x22, 0x36, 0x0a, 0x0a, 0x53,
	0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x32, 0x85, 0x01, 0x0a, 0x0b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x3c, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x38, 0x0a, 0x03, 0x53, 0x6d, 0x73, 0x12, 0x17, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_mail_proto_rawDescOnce sync.Once
	file_internal_proto_mail_proto_rawDescData = file_internal_proto_mail_proto_rawDesc
)

func file_internal_proto_mail_proto_rawDescGZIP() []byte {
	file_internal_proto_mail_proto_rawDescOnce.Do(func() {
		file_internal_proto_mail_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_mail_proto_rawDescData)
	})
	return file_internal_proto_mail_proto_rawDescData
}

var file_internal_proto_mail_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_proto_mail_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_proto_mail_proto_goTypes = []interface{}{
	(EmailRequest_BodyFormat)(0), // 0: notifier.v1.EmailRequest.BodyFormat
	(*EmailRequest)(nil),         // 1: notifier.v1.EmailRequest
	(*SmsRequest)(nil),           // 2: notifier.v1.SmsRequest
	(*emptypb.Empty)(nil),        // 3: google.protobuf.Empty
}
var file_internal_proto_mail_proto_depIdxs = []int32{
	0, // 0: notifier.v1.EmailRequest.body_type:type_name -> notifier.v1.EmailRequest.BodyFormat
	1, // 1: notifier.v1.Notificator.Email:input_type -> notifier.v1.EmailRequest
	2, // 2: notifier.v1.Notificator.Sms:input_type -> notifier.v1.SmsRequest
	3, // 3: notifier.v1.Notificator.Email:output_type -> google.protobuf.Empty
	3, // 4: notifier.v1.Notificator.Sms:output_type -> google.protobuf.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_proto_mail_proto_init() }
func file_internal_proto_mail_proto_init() {
	if File_internal_proto_mail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_mail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailRequest); i {
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
		file_internal_proto_mail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmsRequest); i {
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
			RawDescriptor: file_internal_proto_mail_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_mail_proto_goTypes,
		DependencyIndexes: file_internal_proto_mail_proto_depIdxs,
		EnumInfos:         file_internal_proto_mail_proto_enumTypes,
		MessageInfos:      file_internal_proto_mail_proto_msgTypes,
	}.Build()
	File_internal_proto_mail_proto = out.File
	file_internal_proto_mail_proto_rawDesc = nil
	file_internal_proto_mail_proto_goTypes = nil
	file_internal_proto_mail_proto_depIdxs = nil
}
