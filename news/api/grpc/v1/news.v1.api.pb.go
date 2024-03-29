// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: news.v1.api.proto

package v1

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

type NewsCategoryDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticleId     int64  `protobuf:"varint,1,opt,name=ArticleId,proto3" json:"ArticleId,omitempty"`
	Title         string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	FromName      string `protobuf:"bytes,3,opt,name=FromName,proto3" json:"FromName,omitempty"`
	CreateTime    int64  `protobuf:"varint,4,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	Intro         string `protobuf:"bytes,5,opt,name=Intro,proto3" json:"Intro,omitempty"`
	AuthorId      int64  `protobuf:"varint,6,opt,name=AuthorId,proto3" json:"AuthorId,omitempty"`
	Status        int32  `protobuf:"varint,7,opt,name=Status,proto3" json:"Status,omitempty"`
	RowPicUrl     string `protobuf:"bytes,8,opt,name=RowPicUrl,proto3" json:"RowPicUrl,omitempty"`
	ColPicUrl     string `protobuf:"bytes,9,opt,name=ColPicUrl,proto3" json:"ColPicUrl,omitempty"`
	PageUrl       string `protobuf:"bytes,10,opt,name=PageUrl,proto3" json:"PageUrl,omitempty"`
	AuthorUid     int64  `protobuf:"varint,11,opt,name=AuthorUid,proto3" json:"AuthorUid,omitempty"`
	Cover         string `protobuf:"bytes,12,opt,name=Cover,proto3" json:"Cover,omitempty"`
	Nickname      string `protobuf:"bytes,13,opt,name=Nickname,proto3" json:"Nickname,omitempty"`
	MoodAmount    int64  `protobuf:"varint,14,opt,name=MoodAmount,proto3" json:"MoodAmount,omitempty"`
	CommentAmount int64  `protobuf:"varint,15,opt,name=CommentAmount,proto3" json:"CommentAmount,omitempty"`
	TagId         int64  `protobuf:"varint,16,opt,name=TagId,proto3" json:"TagId,omitempty"`
}

func (x *NewsCategoryDetail) Reset() {
	*x = NewsCategoryDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_v1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsCategoryDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsCategoryDetail) ProtoMessage() {}

func (x *NewsCategoryDetail) ProtoReflect() protoreflect.Message {
	mi := &file_news_v1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsCategoryDetail.ProtoReflect.Descriptor instead.
func (*NewsCategoryDetail) Descriptor() ([]byte, []int) {
	return file_news_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *NewsCategoryDetail) GetArticleId() int64 {
	if x != nil {
		return x.ArticleId
	}
	return 0
}

func (x *NewsCategoryDetail) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NewsCategoryDetail) GetFromName() string {
	if x != nil {
		return x.FromName
	}
	return ""
}

func (x *NewsCategoryDetail) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *NewsCategoryDetail) GetIntro() string {
	if x != nil {
		return x.Intro
	}
	return ""
}

func (x *NewsCategoryDetail) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *NewsCategoryDetail) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *NewsCategoryDetail) GetRowPicUrl() string {
	if x != nil {
		return x.RowPicUrl
	}
	return ""
}

func (x *NewsCategoryDetail) GetColPicUrl() string {
	if x != nil {
		return x.ColPicUrl
	}
	return ""
}

func (x *NewsCategoryDetail) GetPageUrl() string {
	if x != nil {
		return x.PageUrl
	}
	return ""
}

func (x *NewsCategoryDetail) GetAuthorUid() int64 {
	if x != nil {
		return x.AuthorUid
	}
	return 0
}

func (x *NewsCategoryDetail) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *NewsCategoryDetail) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *NewsCategoryDetail) GetMoodAmount() int64 {
	if x != nil {
		return x.MoodAmount
	}
	return 0
}

func (x *NewsCategoryDetail) GetCommentAmount() int64 {
	if x != nil {
		return x.CommentAmount
	}
	return 0
}

func (x *NewsCategoryDetail) GetTagId() int64 {
	if x != nil {
		return x.TagId
	}
	return 0
}

type NewsCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagId   int64  `protobuf:"varint,1,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
	TagName string `protobuf:"bytes,2,opt,name=tag_name,json=tagName,proto3" json:"tag_name,omitempty"`
}

func (x *NewsCategory) Reset() {
	*x = NewsCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_v1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsCategory) ProtoMessage() {}

func (x *NewsCategory) ProtoReflect() protoreflect.Message {
	mi := &file_news_v1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsCategory.ProtoReflect.Descriptor instead.
func (*NewsCategory) Descriptor() ([]byte, []int) {
	return file_news_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *NewsCategory) GetTagId() int64 {
	if x != nil {
		return x.TagId
	}
	return 0
}

func (x *NewsCategory) GetTagName() string {
	if x != nil {
		return x.TagName
	}
	return ""
}

type ListNewsCategoryDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagId  int32 `protobuf:"varint,1,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
	Sort   int32 `protobuf:"varint,2,opt,name=sort,proto3" json:"sort,omitempty"` // 0:人气排序 1:更新排序
	Limit  int32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListNewsCategoryDetailRequest) Reset() {
	*x = ListNewsCategoryDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_v1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNewsCategoryDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNewsCategoryDetailRequest) ProtoMessage() {}

func (x *ListNewsCategoryDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_news_v1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNewsCategoryDetailRequest.ProtoReflect.Descriptor instead.
func (*ListNewsCategoryDetailRequest) Descriptor() ([]byte, []int) {
	return file_news_v1_api_proto_rawDescGZIP(), []int{2}
}

func (x *ListNewsCategoryDetailRequest) GetTagId() int32 {
	if x != nil {
		return x.TagId
	}
	return 0
}

func (x *ListNewsCategoryDetailRequest) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *ListNewsCategoryDetailRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListNewsCategoryDetailRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListNewsCategoryDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details []*NewsCategoryDetail `protobuf:"bytes,1,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *ListNewsCategoryDetailResponse) Reset() {
	*x = ListNewsCategoryDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_v1_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNewsCategoryDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNewsCategoryDetailResponse) ProtoMessage() {}

func (x *ListNewsCategoryDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_news_v1_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNewsCategoryDetailResponse.ProtoReflect.Descriptor instead.
func (*ListNewsCategoryDetailResponse) Descriptor() ([]byte, []int) {
	return file_news_v1_api_proto_rawDescGZIP(), []int{3}
}

func (x *ListNewsCategoryDetailResponse) GetDetails() []*NewsCategoryDetail {
	if x != nil {
		return x.Details
	}
	return nil
}

type ListNewsCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListNewsCategoryRequest) Reset() {
	*x = ListNewsCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_v1_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNewsCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNewsCategoryRequest) ProtoMessage() {}

func (x *ListNewsCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_news_v1_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNewsCategoryRequest.ProtoReflect.Descriptor instead.
func (*ListNewsCategoryRequest) Descriptor() ([]byte, []int) {
	return file_news_v1_api_proto_rawDescGZIP(), []int{4}
}

type ListNewsCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListNewsCategoryResponse) Reset() {
	*x = ListNewsCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_v1_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNewsCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNewsCategoryResponse) ProtoMessage() {}

func (x *ListNewsCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_news_v1_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNewsCategoryResponse.ProtoReflect.Descriptor instead.
func (*ListNewsCategoryResponse) Descriptor() ([]byte, []int) {
	return file_news_v1_api_proto_rawDescGZIP(), []int{5}
}

var File_news_v1_api_proto protoreflect.FileDescriptor

var file_news_v1_api_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x22, 0xd0, 0x03, 0x0a,
	0x12, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x72, 0x6f, 0x6d, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x72, 0x6f, 0x6d, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x52, 0x6f, 0x77, 0x50, 0x69, 0x63, 0x55, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x52, 0x6f, 0x77, 0x50, 0x69, 0x63, 0x55, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x43,
	0x6f, 0x6c, 0x50, 0x69, 0x63, 0x55, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x43, 0x6f, 0x6c, 0x50, 0x69, 0x63, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x61, 0x67, 0x65,
	0x55, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x55, 0x69, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x55, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x6f, 0x6f, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4d, 0x6f, 0x6f, 0x64, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x61, 0x67,
	0x49, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x54, 0x61, 0x67, 0x49, 0x64, 0x22,
	0x40, 0x0a, 0x0c, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x15, 0x0a, 0x06, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x67, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x78, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x57, 0x0a, 0x1e, 0x4c,
	0x69, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x07, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x22, 0x19, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x73,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x1a, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xd1, 0x01, 0x0a, 0x0b,
	0x4e, 0x65, 0x77, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x16, 0x4c,
	0x69, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x26, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x73,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65,
	0x77, 0x73, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x2e, 0x6e, 0x65, 0x77,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6e,
	0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_news_v1_api_proto_rawDescOnce sync.Once
	file_news_v1_api_proto_rawDescData = file_news_v1_api_proto_rawDesc
)

func file_news_v1_api_proto_rawDescGZIP() []byte {
	file_news_v1_api_proto_rawDescOnce.Do(func() {
		file_news_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_news_v1_api_proto_rawDescData)
	})
	return file_news_v1_api_proto_rawDescData
}

var file_news_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_news_v1_api_proto_goTypes = []interface{}{
	(*NewsCategoryDetail)(nil),             // 0: news.v1.NewsCategoryDetail
	(*NewsCategory)(nil),                   // 1: news.v1.NewsCategory
	(*ListNewsCategoryDetailRequest)(nil),  // 2: news.v1.ListNewsCategoryDetailRequest
	(*ListNewsCategoryDetailResponse)(nil), // 3: news.v1.ListNewsCategoryDetailResponse
	(*ListNewsCategoryRequest)(nil),        // 4: news.v1.ListNewsCategoryRequest
	(*ListNewsCategoryResponse)(nil),       // 5: news.v1.ListNewsCategoryResponse
}
var file_news_v1_api_proto_depIdxs = []int32{
	0, // 0: news.v1.ListNewsCategoryDetailResponse.details:type_name -> news.v1.NewsCategoryDetail
	2, // 1: news.v1.NewsService.ListNewsCategoryDetail:input_type -> news.v1.ListNewsCategoryDetailRequest
	4, // 2: news.v1.NewsService.ListNewsCategory:input_type -> news.v1.ListNewsCategoryRequest
	3, // 3: news.v1.NewsService.ListNewsCategoryDetail:output_type -> news.v1.ListNewsCategoryDetailResponse
	5, // 4: news.v1.NewsService.ListNewsCategory:output_type -> news.v1.ListNewsCategoryResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_news_v1_api_proto_init() }
func file_news_v1_api_proto_init() {
	if File_news_v1_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_news_v1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsCategoryDetail); i {
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
		file_news_v1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsCategory); i {
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
		file_news_v1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNewsCategoryDetailRequest); i {
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
		file_news_v1_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNewsCategoryDetailResponse); i {
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
		file_news_v1_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNewsCategoryRequest); i {
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
		file_news_v1_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNewsCategoryResponse); i {
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
			RawDescriptor: file_news_v1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_news_v1_api_proto_goTypes,
		DependencyIndexes: file_news_v1_api_proto_depIdxs,
		MessageInfos:      file_news_v1_api_proto_msgTypes,
	}.Build()
	File_news_v1_api_proto = out.File
	file_news_v1_api_proto_rawDesc = nil
	file_news_v1_api_proto_goTypes = nil
	file_news_v1_api_proto_depIdxs = nil
}
