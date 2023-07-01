// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.6
// source: block.proto

package proto

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

type TextualType int32

const (
	TextualType_TEXTUAL_TYPE_UNKNOWN  TextualType = 0
	TextualType_TEXTUAL_TYPE_TEXT     TextualType = 1
	TextualType_TEXTUAL_TYPE_CODE     TextualType = 2
	TextualType_TEXTUAL_TYPE_HTML     TextualType = 3
	TextualType_TEXTUAL_TYPE_MARKDOWN TextualType = 4
)

// Enum value maps for TextualType.
var (
	TextualType_name = map[int32]string{
		0: "TEXTUAL_TYPE_UNKNOWN",
		1: "TEXTUAL_TYPE_TEXT",
		2: "TEXTUAL_TYPE_CODE",
		3: "TEXTUAL_TYPE_HTML",
		4: "TEXTUAL_TYPE_MARKDOWN",
	}
	TextualType_value = map[string]int32{
		"TEXTUAL_TYPE_UNKNOWN":  0,
		"TEXTUAL_TYPE_TEXT":     1,
		"TEXTUAL_TYPE_CODE":     2,
		"TEXTUAL_TYPE_HTML":     3,
		"TEXTUAL_TYPE_MARKDOWN": 4,
	}
)

func (x TextualType) Enum() *TextualType {
	p := new(TextualType)
	*p = x
	return p
}

func (x TextualType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TextualType) Descriptor() protoreflect.EnumDescriptor {
	return file_block_proto_enumTypes[0].Descriptor()
}

func (TextualType) Type() protoreflect.EnumType {
	return &file_block_proto_enumTypes[0]
}

func (x TextualType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TextualType.Descriptor instead.
func (TextualType) EnumDescriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{0}
}

type MediaType int32

const (
	MediaType_MEDIA_TYPE_UNKNOWN MediaType = 0
	MediaType_MEDIA_TYPE_IMAGE   MediaType = 1
	MediaType_MEDIA_TYPE_VIDEO   MediaType = 2
	MediaType_MEDIA_TYPE_AUDIO   MediaType = 3
	MediaType_MEDIA_TYPE_FILE    MediaType = 4
)

// Enum value maps for MediaType.
var (
	MediaType_name = map[int32]string{
		0: "MEDIA_TYPE_UNKNOWN",
		1: "MEDIA_TYPE_IMAGE",
		2: "MEDIA_TYPE_VIDEO",
		3: "MEDIA_TYPE_AUDIO",
		4: "MEDIA_TYPE_FILE",
	}
	MediaType_value = map[string]int32{
		"MEDIA_TYPE_UNKNOWN": 0,
		"MEDIA_TYPE_IMAGE":   1,
		"MEDIA_TYPE_VIDEO":   2,
		"MEDIA_TYPE_AUDIO":   3,
		"MEDIA_TYPE_FILE":    4,
	}
)

func (x MediaType) Enum() *MediaType {
	p := new(MediaType)
	*p = x
	return p
}

func (x MediaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MediaType) Descriptor() protoreflect.EnumDescriptor {
	return file_block_proto_enumTypes[1].Descriptor()
}

func (MediaType) Type() protoreflect.EnumType {
	return &file_block_proto_enumTypes[1]
}

func (x MediaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MediaType.Descriptor instead.
func (MediaType) EnumDescriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{1}
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId     string          `protobuf:"bytes,1,opt,name=block_id,json=id,proto3" json:"block_id,omitempty"`
	Name        string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type        string          `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Lang        string          `protobuf:"bytes,4,opt,name=lang,proto3" json:"lang,omitempty"`
	Version     int32           `protobuf:"varint,5,opt,name=version,proto3" json:"version,omitempty"`
	Tags        []string        `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
	Categories  []string        `protobuf:"bytes,7,rep,name=categories,proto3" json:"categories,omitempty"`
	Authors     []*Author       `protobuf:"bytes,8,rep,name=authors,proto3" json:"authors,omitempty"`
	Content     []*ElementType  `protobuf:"bytes,9,rep,name=content,proto3" json:"content,omitempty"`
	Children    []*BlockContent `protobuf:"bytes,10,rep,name=children,proto3" json:"children,omitempty"`
	Rules       *BlockRules     `protobuf:"bytes,11,opt,name=rules,proto3" json:"rules,omitempty"`
	UpdatedAt   string          `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedAt   string          `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Description string          `protobuf:"bytes,14,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

func (x *Block) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Block) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Block) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *Block) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Block) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Block) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *Block) GetAuthors() []*Author {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *Block) GetContent() []*ElementType {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Block) GetChildren() []*BlockContent {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *Block) GetRules() *BlockRules {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *Block) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Block) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Block) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type BlockContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId  string          `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Name     string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Lang     string          `protobuf:"bytes,3,opt,name=lang,proto3" json:"lang,omitempty"`
	Version  int32           `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	Content  []*ElementType  `protobuf:"bytes,5,rep,name=content,proto3" json:"content,omitempty"`
	Children []*BlockContent `protobuf:"bytes,6,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *BlockContent) Reset() {
	*x = BlockContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockContent) ProtoMessage() {}

func (x *BlockContent) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockContent.ProtoReflect.Descriptor instead.
func (*BlockContent) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{1}
}

func (x *BlockContent) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

func (x *BlockContent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BlockContent) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *BlockContent) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *BlockContent) GetContent() []*ElementType {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *BlockContent) GetChildren() []*BlockContent {
	if x != nil {
		return x.Children
	}
	return nil
}

type BlockMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId     string    `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Name        string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type        string    `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Tags        []string  `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	Categories  []string  `protobuf:"bytes,5,rep,name=categories,proto3" json:"categories,omitempty"`
	Authors     []*Author `protobuf:"bytes,6,rep,name=authors,proto3" json:"authors,omitempty"`
	Description string    `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *BlockMeta) Reset() {
	*x = BlockMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockMeta) ProtoMessage() {}

func (x *BlockMeta) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockMeta.ProtoReflect.Descriptor instead.
func (*BlockMeta) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{2}
}

func (x *BlockMeta) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

func (x *BlockMeta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BlockMeta) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *BlockMeta) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *BlockMeta) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *BlockMeta) GetAuthors() []*Author {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *BlockMeta) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Image string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{3}
}

func (x *Author) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Author) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type BlockRules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RuleName          string `protobuf:"bytes,1,opt,name=rule_name,json=ruleName,proto3" json:"rule_name,omitempty"`
	Description       string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Nested            bool   `protobuf:"varint,3,opt,name=nested,proto3" json:"nested,omitempty"`
	HasLikes          bool   `protobuf:"varint,4,opt,name=has_likes,json=hasLikes,proto3" json:"has_likes,omitempty"`
	HasComments       bool   `protobuf:"varint,5,opt,name=has_comments,json=hasComments,proto3" json:"has_comments,omitempty"`
	CommentsHasLikes  bool   `protobuf:"varint,7,opt,name=comments_has_likes,json=commentsHasLikes,proto3" json:"comments_has_likes,omitempty"`
	CommentsEditable  bool   `protobuf:"varint,8,opt,name=comments_editable,json=commentsEditable,proto3" json:"comments_editable,omitempty"`
	CommentsMaxNested int32  `protobuf:"varint,9,opt,name=comments_max_nested,json=commentsMaxNested,proto3" json:"comments_max_nested,omitempty"`
}

func (x *BlockRules) Reset() {
	*x = BlockRules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockRules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockRules) ProtoMessage() {}

func (x *BlockRules) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockRules.ProtoReflect.Descriptor instead.
func (*BlockRules) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{4}
}

func (x *BlockRules) GetRuleName() string {
	if x != nil {
		return x.RuleName
	}
	return ""
}

func (x *BlockRules) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BlockRules) GetNested() bool {
	if x != nil {
		return x.Nested
	}
	return false
}

func (x *BlockRules) GetHasLikes() bool {
	if x != nil {
		return x.HasLikes
	}
	return false
}

func (x *BlockRules) GetHasComments() bool {
	if x != nil {
		return x.HasComments
	}
	return false
}

func (x *BlockRules) GetCommentsHasLikes() bool {
	if x != nil {
		return x.CommentsHasLikes
	}
	return false
}

func (x *BlockRules) GetCommentsEditable() bool {
	if x != nil {
		return x.CommentsEditable
	}
	return false
}

func (x *BlockRules) GetCommentsMaxNested() int32 {
	if x != nil {
		return x.CommentsMaxNested
	}
	return 0
}

type ElementType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Element:
	//
	//	*ElementType_Media
	//	*ElementType_Text
	Element isElementType_Element `protobuf_oneof:"Element"`
}

func (x *ElementType) Reset() {
	*x = ElementType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElementType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElementType) ProtoMessage() {}

func (x *ElementType) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElementType.ProtoReflect.Descriptor instead.
func (*ElementType) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{5}
}

func (m *ElementType) GetElement() isElementType_Element {
	if m != nil {
		return m.Element
	}
	return nil
}

func (x *ElementType) GetMedia() *Media {
	if x, ok := x.GetElement().(*ElementType_Media); ok {
		return x.Media
	}
	return nil
}

func (x *ElementType) GetText() *Textual {
	if x, ok := x.GetElement().(*ElementType_Text); ok {
		return x.Text
	}
	return nil
}

type isElementType_Element interface {
	isElementType_Element()
}

type ElementType_Media struct {
	Media *Media `protobuf:"bytes,1,opt,name=media,proto3,oneof"`
}

type ElementType_Text struct {
	Text *Textual `protobuf:"bytes,2,opt,name=text,proto3,oneof"`
}

func (*ElementType_Media) isElementType_Element() {}

func (*ElementType_Text) isElementType_Element() {}

type Media struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string    `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Type  MediaType `protobuf:"varint,2,opt,name=type,proto3,enum=proto.MediaType" json:"type,omitempty"`
	File  string    `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
	Alt   string    `protobuf:"bytes,4,opt,name=alt,proto3" json:"alt,omitempty"`
}

func (x *Media) Reset() {
	*x = Media{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Media) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Media) ProtoMessage() {}

func (x *Media) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Media.ProtoReflect.Descriptor instead.
func (*Media) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{6}
}

func (x *Media) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Media) GetType() MediaType {
	if x != nil {
		return x.Type
	}
	return MediaType_MEDIA_TYPE_UNKNOWN
}

func (x *Media) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *Media) GetAlt() string {
	if x != nil {
		return x.Alt
	}
	return ""
}

type Textual struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type    TextualType `protobuf:"varint,2,opt,name=type,proto3,enum=proto.TextualType" json:"type,omitempty"`
	Content string      `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Hint    string      `protobuf:"bytes,4,opt,name=hint,proto3" json:"hint,omitempty"`
}

func (x *Textual) Reset() {
	*x = Textual{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Textual) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Textual) ProtoMessage() {}

func (x *Textual) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Textual.ProtoReflect.Descriptor instead.
func (*Textual) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{7}
}

func (x *Textual) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Textual) GetType() TextualType {
	if x != nil {
		return x.Type
	}
	return TextualType_TEXTUAL_TYPE_UNKNOWN
}

func (x *Textual) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Textual) GetHint() string {
	if x != nil {
		return x.Hint
	}
	return ""
}

var File_block_proto protoreflect.FileDescriptor

var file_block_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x03, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x14,
	0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x61, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x27,
	0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x07,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65,
	0x6e, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x27, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0xca, 0x01, 0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c,
	0x61, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x08, 0x63,
	0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0xcd, 0x01, 0x0a,
	0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x27, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x42, 0x0a, 0x06,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x22, 0xae, 0x02, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x5f, 0x6c, 0x69,
	0x6b, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x61, 0x73, 0x4c, 0x69,
	0x6b, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x61, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x68, 0x61, 0x73, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x5f, 0x68, 0x61, 0x73, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x48, 0x61, 0x73, 0x4c,
	0x69, 0x6b, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x5f, 0x65, 0x64, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x64, 0x69, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x6d, 0x61,
	0x78, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4d, 0x61, 0x78, 0x4e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x22, 0x64, 0x0a, 0x0b, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x24, 0x0a, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x48, 0x00, 0x52,
	0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x78,
	0x74, 0x75, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x42, 0x09, 0x0a, 0x07,
	0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x69, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61,
	0x6c, 0x74, 0x22, 0x73, 0x0a, 0x07, 0x54, 0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x2a, 0x87, 0x01, 0x0a, 0x0b, 0x54, 0x65, 0x78, 0x74,
	0x75, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x45, 0x58, 0x54, 0x55,
	0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x45, 0x58, 0x54, 0x55, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x54, 0x45, 0x58, 0x54, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x45, 0x58, 0x54,
	0x55, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x02, 0x12,
	0x15, 0x0a, 0x11, 0x54, 0x45, 0x58, 0x54, 0x55, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x48, 0x54, 0x4d, 0x4c, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x45, 0x58, 0x54, 0x55, 0x41,
	0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x44, 0x4f, 0x57, 0x4e, 0x10,
	0x04, 0x2a, 0x7a, 0x0a, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x12, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10,
	0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f,
	0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x45, 0x44, 0x49,
	0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x04, 0x42, 0x0f, 0x5a,
	0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_block_proto_rawDescOnce sync.Once
	file_block_proto_rawDescData = file_block_proto_rawDesc
)

func file_block_proto_rawDescGZIP() []byte {
	file_block_proto_rawDescOnce.Do(func() {
		file_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_block_proto_rawDescData)
	})
	return file_block_proto_rawDescData
}

var file_block_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_block_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_block_proto_goTypes = []interface{}{
	(TextualType)(0),     // 0: proto.TextualType
	(MediaType)(0),       // 1: proto.MediaType
	(*Block)(nil),        // 2: proto.Block
	(*BlockContent)(nil), // 3: proto.BlockContent
	(*BlockMeta)(nil),    // 4: proto.BlockMeta
	(*Author)(nil),       // 5: proto.Author
	(*BlockRules)(nil),   // 6: proto.BlockRules
	(*ElementType)(nil),  // 7: proto.ElementType
	(*Media)(nil),        // 8: proto.Media
	(*Textual)(nil),      // 9: proto.Textual
}
var file_block_proto_depIdxs = []int32{
	5,  // 0: proto.Block.authors:type_name -> proto.Author
	7,  // 1: proto.Block.content:type_name -> proto.ElementType
	3,  // 2: proto.Block.children:type_name -> proto.BlockContent
	6,  // 3: proto.Block.rules:type_name -> proto.BlockRules
	7,  // 4: proto.BlockContent.content:type_name -> proto.ElementType
	3,  // 5: proto.BlockContent.children:type_name -> proto.BlockContent
	5,  // 6: proto.BlockMeta.authors:type_name -> proto.Author
	8,  // 7: proto.ElementType.media:type_name -> proto.Media
	9,  // 8: proto.ElementType.text:type_name -> proto.Textual
	1,  // 9: proto.Media.type:type_name -> proto.MediaType
	0,  // 10: proto.Textual.type:type_name -> proto.TextualType
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_block_proto_init() }
func file_block_proto_init() {
	if File_block_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockContent); i {
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
		file_block_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockMeta); i {
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
		file_block_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
		file_block_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockRules); i {
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
		file_block_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElementType); i {
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
		file_block_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Media); i {
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
		file_block_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Textual); i {
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
	file_block_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*ElementType_Media)(nil),
		(*ElementType_Text)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_block_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_block_proto_goTypes,
		DependencyIndexes: file_block_proto_depIdxs,
		EnumInfos:         file_block_proto_enumTypes,
		MessageInfos:      file_block_proto_msgTypes,
	}.Build()
	File_block_proto = out.File
	file_block_proto_rawDesc = nil
	file_block_proto_goTypes = nil
	file_block_proto_depIdxs = nil
}
