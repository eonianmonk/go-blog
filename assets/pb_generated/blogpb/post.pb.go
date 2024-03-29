// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: post.proto

package blogpb

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

type State int32

const (
	State_Readable State = 0
	State_Deleted  State = 1
	State_Hidden   State = 2
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "Readable",
		1: "Deleted",
		2: "Hidden",
	}
	State_value = map[string]int32{
		"Readable": 0,
		"Deleted":  1,
		"Hidden":   2,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_post_proto_enumTypes[0].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_post_proto_enumTypes[0]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{0}
}

type PostId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PostId) Reset() {
	*x = PostId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostId) ProtoMessage() {}

func (x *PostId) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostId.ProtoReflect.Descriptor instead.
func (*PostId) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{0}
}

func (x *PostId) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// post metatdata
type PostHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId   int32      `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Author   string     `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Title    string     `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	State    State      `protobuf:"varint,4,opt,name=state,proto3,enum=blog.State" json:"state,omitempty"`
	Body     *string    `protobuf:"bytes,5,opt,name=body,proto3,oneof" json:"body,omitempty"`
	Comments []*Comment `protobuf:"bytes,6,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *PostHeader) Reset() {
	*x = PostHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostHeader) ProtoMessage() {}

func (x *PostHeader) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostHeader.ProtoReflect.Descriptor instead.
func (*PostHeader) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{1}
}

func (x *PostHeader) GetPostId() int32 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *PostHeader) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *PostHeader) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PostHeader) GetState() State {
	if x != nil {
		return x.State
	}
	return State_Readable
}

func (x *PostHeader) GetBody() string {
	if x != nil && x.Body != nil {
		return *x.Body
	}
	return ""
}

func (x *PostHeader) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type PostBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *PostBody) Reset() {
	*x = PostBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostBody) ProtoMessage() {}

func (x *PostBody) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostBody.ProtoReflect.Descriptor instead.
func (*PostBody) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{2}
}

func (x *PostBody) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type RequestHeaders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"` // amount of headers
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"` // offset for pagination
}

func (x *RequestHeaders) Reset() {
	*x = RequestHeaders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHeaders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHeaders) ProtoMessage() {}

func (x *RequestHeaders) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHeaders.ProtoReflect.Descriptor instead.
func (*RequestHeaders) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{3}
}

func (x *RequestHeaders) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *RequestHeaders) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ReturnHeaders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Headers []*PostHeader `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
}

func (x *ReturnHeaders) Reset() {
	*x = ReturnHeaders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReturnHeaders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnHeaders) ProtoMessage() {}

func (x *ReturnHeaders) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnHeaders.ProtoReflect.Descriptor instead.
func (*ReturnHeaders) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{4}
}

func (x *ReturnHeaders) GetHeaders() []*PostHeader {
	if x != nil {
		return x.Headers
	}
	return nil
}

var File_post_proto protoreflect.FileDescriptor

var file_post_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x6c,
	0x6f, 0x67, 0x1a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x18, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0xc3, 0x01, 0x0a, 0x0a,
	0x50, 0x6f, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6f, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x21, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0b, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a,
	0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x62, 0x6f, 0x64,
	0x79, 0x22, 0x24, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x40, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x3b, 0x0a, 0x0d, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x07, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x07, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2a, 0x2e, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x0c, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x69,
	0x64, 0x64, 0x65, 0x6e, 0x10, 0x02, 0x32, 0xf1, 0x01, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x67, 0x12,
	0x2c, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x10, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a,
	0x0c, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a,
	0x08, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0c, 0x2e, 0x62, 0x6c, 0x6f, 0x67,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x0c, 0x2e, 0x62, 0x6c, 0x6f, 0x67,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0c, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x37, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x14, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x13, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x62, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_post_proto_rawDescOnce sync.Once
	file_post_proto_rawDescData = file_post_proto_rawDesc
)

func file_post_proto_rawDescGZIP() []byte {
	file_post_proto_rawDescOnce.Do(func() {
		file_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_post_proto_rawDescData)
	})
	return file_post_proto_rawDescData
}

var file_post_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_post_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_post_proto_goTypes = []interface{}{
	(State)(0),             // 0: blog.State
	(*PostId)(nil),         // 1: blog.PostId
	(*PostHeader)(nil),     // 2: blog.PostHeader
	(*PostBody)(nil),       // 3: blog.PostBody
	(*RequestHeaders)(nil), // 4: blog.RequestHeaders
	(*ReturnHeaders)(nil),  // 5: blog.ReturnHeaders
	(*Comment)(nil),        // 6: blog.Comment
}
var file_post_proto_depIdxs = []int32{
	0, // 0: blog.PostHeader.state:type_name -> blog.State
	6, // 1: blog.PostHeader.comments:type_name -> blog.Comment
	2, // 2: blog.ReturnHeaders.headers:type_name -> blog.PostHeader
	2, // 3: blog.Blog.CreatePost:input_type -> blog.PostHeader
	1, // 4: blog.Blog.ReadPost:input_type -> blog.PostId
	2, // 5: blog.Blog.UpdatePost:input_type -> blog.PostHeader
	1, // 6: blog.Blog.DeletePost:input_type -> blog.PostId
	4, // 7: blog.Blog.GetHeaders:input_type -> blog.RequestHeaders
	1, // 8: blog.Blog.CreatePost:output_type -> blog.PostId
	2, // 9: blog.Blog.ReadPost:output_type -> blog.PostHeader
	1, // 10: blog.Blog.UpdatePost:output_type -> blog.PostId
	1, // 11: blog.Blog.DeletePost:output_type -> blog.PostId
	5, // 12: blog.Blog.GetHeaders:output_type -> blog.ReturnHeaders
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_post_proto_init() }
func file_post_proto_init() {
	if File_post_proto != nil {
		return
	}
	file_comment_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostId); i {
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
		file_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostHeader); i {
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
		file_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostBody); i {
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
		file_post_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHeaders); i {
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
		file_post_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReturnHeaders); i {
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
	file_post_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_post_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_post_proto_goTypes,
		DependencyIndexes: file_post_proto_depIdxs,
		EnumInfos:         file_post_proto_enumTypes,
		MessageInfos:      file_post_proto_msgTypes,
	}.Build()
	File_post_proto = out.File
	file_post_proto_rawDesc = nil
	file_post_proto_goTypes = nil
	file_post_proto_depIdxs = nil
}
