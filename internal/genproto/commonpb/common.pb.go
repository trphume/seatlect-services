// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: common.proto

package commonpb

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Dob      string   `protobuf:"bytes,2,opt,name=dob,proto3" json:"dob,omitempty"`
	Avatar   string   `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Favorite []string `protobuf:"bytes,4,rep,name=favorite,proto3" json:"favorite,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetFavorite() []string {
	if x != nil {
		return x.Favorite
	}
	return nil
}

type Business struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId          string      `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	Name         string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type         string      `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Tags         []string    `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	Description  string      `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Location     *Latlng     `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	Address      string      `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	DisplayImage string      `protobuf:"bytes,8,opt,name=displayImage,proto3" json:"displayImage,omitempty"`
	Images       []string    `protobuf:"bytes,9,rep,name=images,proto3" json:"images,omitempty"`
	Menu         []*MenuItem `protobuf:"bytes,10,rep,name=menu,proto3" json:"menu,omitempty"`
}

func (x *Business) Reset() {
	*x = Business{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Business) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Business) ProtoMessage() {}

func (x *Business) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Business.ProtoReflect.Descriptor instead.
func (*Business) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *Business) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *Business) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Business) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Business) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Business) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Business) GetLocation() *Latlng {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Business) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Business) GetDisplayImage() string {
	if x != nil {
		return x.DisplayImage
	}
	return ""
}

func (x *Business) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *Business) GetMenu() []*MenuItem {
	if x != nil {
		return x.Menu
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId           string       `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	ReservationId string       `protobuf:"bytes,2,opt,name=reservationId,proto3" json:"reservationId,omitempty"`
	Business      string       `protobuf:"bytes,3,opt,name=business,proto3" json:"business,omitempty"`
	Start         string       `protobuf:"bytes,4,opt,name=start,proto3" json:"start,omitempty"`
	End           string       `protobuf:"bytes,5,opt,name=end,proto3" json:"end,omitempty"`
	Seats         []*OrderSeat `protobuf:"bytes,6,rep,name=seats,proto3" json:"seats,omitempty"`
	Status        string       `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Image         string       `protobuf:"bytes,8,opt,name=image,proto3" json:"image,omitempty"`
	ExtraSpace    int32        `protobuf:"varint,9,opt,name=extraSpace,proto3" json:"extraSpace,omitempty"`
	Name          string       `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *Order) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *Order) GetReservationId() string {
	if x != nil {
		return x.ReservationId
	}
	return ""
}

func (x *Order) GetBusiness() string {
	if x != nil {
		return x.Business
	}
	return ""
}

func (x *Order) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *Order) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

func (x *Order) GetSeats() []*OrderSeat {
	if x != nil {
		return x.Seats
	}
	return nil
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Order) GetExtraSpace() int32 {
	if x != nil {
		return x.ExtraSpace
	}
	return 0
}

func (x *Order) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Latlng struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Latlng) Reset() {
	*x = Latlng{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Latlng) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Latlng) ProtoMessage() {}

func (x *Latlng) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Latlng.ProtoReflect.Descriptor instead.
func (*Latlng) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *Latlng) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Latlng) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type MenuItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Image       string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Price       string `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *MenuItem) Reset() {
	*x = MenuItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuItem) ProtoMessage() {}

func (x *MenuItem) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuItem.ProtoReflect.Descriptor instead.
func (*MenuItem) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

func (x *MenuItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MenuItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MenuItem) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *MenuItem) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

type ReservationPlacement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width  int32              `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height int32              `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Seats  []*ReservationSeat `protobuf:"bytes,3,rep,name=seats,proto3" json:"seats,omitempty"`
}

func (x *ReservationPlacement) Reset() {
	*x = ReservationPlacement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationPlacement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationPlacement) ProtoMessage() {}

func (x *ReservationPlacement) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationPlacement.ProtoReflect.Descriptor instead.
func (*ReservationPlacement) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{5}
}

func (x *ReservationPlacement) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *ReservationPlacement) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *ReservationPlacement) GetSeats() []*ReservationSeat {
	if x != nil {
		return x.Seats
	}
	return nil
}

type OrderSeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Space int32  `protobuf:"varint,4,opt,name=space,proto3" json:"space,omitempty"`
}

func (x *OrderSeat) Reset() {
	*x = OrderSeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderSeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderSeat) ProtoMessage() {}

func (x *OrderSeat) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderSeat.ProtoReflect.Descriptor instead.
func (*OrderSeat) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{6}
}

func (x *OrderSeat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrderSeat) GetSpace() int32 {
	if x != nil {
		return x.Space
	}
	return 0
}

type ReservationSeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Floor    int32   `protobuf:"varint,2,opt,name=floor,proto3" json:"floor,omitempty"`
	Type     string  `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Space    int32   `protobuf:"varint,4,opt,name=space,proto3" json:"space,omitempty"`
	User     string  `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	Status   string  `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	X        float64 `protobuf:"fixed64,7,opt,name=x,proto3" json:"x,omitempty"`
	Y        float64 `protobuf:"fixed64,8,opt,name=y,proto3" json:"y,omitempty"`
	Width    float64 `protobuf:"fixed64,9,opt,name=width,proto3" json:"width,omitempty"`
	Height   float64 `protobuf:"fixed64,10,opt,name=height,proto3" json:"height,omitempty"`
	Rotation float64 `protobuf:"fixed64,11,opt,name=rotation,proto3" json:"rotation,omitempty"`
}

func (x *ReservationSeat) Reset() {
	*x = ReservationSeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationSeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationSeat) ProtoMessage() {}

func (x *ReservationSeat) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationSeat.ProtoReflect.Descriptor instead.
func (*ReservationSeat) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{7}
}

func (x *ReservationSeat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReservationSeat) GetFloor() int32 {
	if x != nil {
		return x.Floor
	}
	return 0
}

func (x *ReservationSeat) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ReservationSeat) GetSpace() int32 {
	if x != nil {
		return x.Space
	}
	return 0
}

func (x *ReservationSeat) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ReservationSeat) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ReservationSeat) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *ReservationSeat) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *ReservationSeat) GetWidth() float64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *ReservationSeat) GetHeight() float64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *ReservationSeat) GetRotation() float64 {
	if x != nil {
		return x.Rotation
	}
	return 0
}

type Reservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BusinessId string                `protobuf:"bytes,2,opt,name=businessId,proto3" json:"businessId,omitempty"`
	Name       string                `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Start      string                `protobuf:"bytes,4,opt,name=start,proto3" json:"start,omitempty"`
	End        string                `protobuf:"bytes,5,opt,name=end,proto3" json:"end,omitempty"`
	Placement  *ReservationPlacement `protobuf:"bytes,6,opt,name=placement,proto3" json:"placement,omitempty"`
	Image      string                `protobuf:"bytes,7,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *Reservation) Reset() {
	*x = Reservation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reservation) ProtoMessage() {}

func (x *Reservation) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reservation.ProtoReflect.Descriptor instead.
func (*Reservation) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{8}
}

func (x *Reservation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Reservation) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *Reservation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Reservation) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *Reservation) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

func (x *Reservation) GetPlacement() *ReservationPlacement {
	if x != nil {
		return x.Placement
	}
	return nil
}

func (x *Reservation) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x73, 0x65, 0x61, 0x74, 0x6c, 0x65, 0x63, 0x74, 0x22, 0x68, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x64, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x22, 0xa5, 0x02, 0x0a, 0x08, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12,
	0x0f, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x4c, 0x61, 0x74, 0x6c,
	0x6e, 0x67, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x26, 0x0a, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x4d, 0x65, 0x6e, 0x75,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x22, 0x8f, 0x02, 0x0a, 0x05, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x0f, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12,
	0x29, 0x0a, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x73, 0x65, 0x61, 0x74, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x65, 0x61, 0x74, 0x52, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x53, 0x70, 0x61, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x42, 0x0a, 0x06,
	0x4c, 0x61, 0x74, 0x6c, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x22, 0x6c, 0x0a, 0x08, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x75,
	0x0a, 0x14, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x2f, 0x0a, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x61, 0x74, 0x52, 0x05,
	0x73, 0x65, 0x61, 0x74, 0x73, 0x22, 0x35, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65,
	0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0xf7, 0x01, 0x0a,
	0x0f, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x61, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c,
	0x0a, 0x01, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x72, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xcd, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65,
	0x6e, 0x64, 0x12, 0x3c, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x6c, 0x65, 0x63, 0x74,
	0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x70, 0x68, 0x75, 0x6d, 0x65, 0x2f, 0x73, 0x65, 0x61, 0x74,
	0x6c, 0x65, 0x63, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_common_proto_goTypes = []interface{}{
	(*User)(nil),                 // 0: seatlect.User
	(*Business)(nil),             // 1: seatlect.Business
	(*Order)(nil),                // 2: seatlect.Order
	(*Latlng)(nil),               // 3: seatlect.Latlng
	(*MenuItem)(nil),             // 4: seatlect.MenuItem
	(*ReservationPlacement)(nil), // 5: seatlect.ReservationPlacement
	(*OrderSeat)(nil),            // 6: seatlect.OrderSeat
	(*ReservationSeat)(nil),      // 7: seatlect.ReservationSeat
	(*Reservation)(nil),          // 8: seatlect.Reservation
}
var file_common_proto_depIdxs = []int32{
	3, // 0: seatlect.Business.location:type_name -> seatlect.Latlng
	4, // 1: seatlect.Business.menu:type_name -> seatlect.MenuItem
	6, // 2: seatlect.Order.seats:type_name -> seatlect.OrderSeat
	7, // 3: seatlect.ReservationPlacement.seats:type_name -> seatlect.ReservationSeat
	5, // 4: seatlect.Reservation.placement:type_name -> seatlect.ReservationPlacement
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Business); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Latlng); i {
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
		file_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuItem); i {
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
		file_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationPlacement); i {
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
		file_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderSeat); i {
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
		file_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationSeat); i {
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
		file_common_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reservation); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
