// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: tuning.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Color struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// all values [0,1]
	C []float32 `protobuf:"fixed32,1,rep,packed,name=c" json:"c,omitempty"`
}

func (x *Color) Reset() {
	*x = Color{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tuning_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Color) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Color) ProtoMessage() {}

func (x *Color) ProtoReflect() protoreflect.Message {
	mi := &file_tuning_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Color.ProtoReflect.Descriptor instead.
func (*Color) Descriptor() ([]byte, []int) {
	return file_tuning_proto_rawDescGZIP(), []int{0}
}

func (x *Color) GetC() []float32 {
	if x != nil {
		return x.C
	}
	return nil
}

// All distances are in world coordinates, all durations are in seconds.
type Tuning struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	World  *World  `protobuf:"bytes,1,opt,name=world" json:"world,omitempty"`
	Todd   *Todd   `protobuf:"bytes,2,opt,name=todd" json:"todd,omitempty"`
	Camera *Camera `protobuf:"bytes,3,opt,name=camera" json:"camera,omitempty"`
}

func (x *Tuning) Reset() {
	*x = Tuning{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tuning_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tuning) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tuning) ProtoMessage() {}

func (x *Tuning) ProtoReflect() protoreflect.Message {
	mi := &file_tuning_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tuning.ProtoReflect.Descriptor instead.
func (*Tuning) Descriptor() ([]byte, []int) {
	return file_tuning_proto_rawDescGZIP(), []int{1}
}

func (x *Tuning) GetWorld() *World {
	if x != nil {
		return x.World
	}
	return nil
}

func (x *Tuning) GetTodd() *Todd {
	if x != nil {
		return x.Todd
	}
	return nil
}

func (x *Tuning) GetCamera() *Camera {
	if x != nil {
		return x.Camera
	}
	return nil
}

type Camera struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TiltSeconds *float32 `protobuf:"fixed32,3,opt,name=tilt_seconds,json=tiltSeconds,def=0.2" json:"tilt_seconds,omitempty"`
}

// Default values for Camera fields.
const (
	Default_Camera_TiltSeconds = float32(0.20000000298023224)
)

func (x *Camera) Reset() {
	*x = Camera{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tuning_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Camera) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Camera) ProtoMessage() {}

func (x *Camera) ProtoReflect() protoreflect.Message {
	mi := &file_tuning_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Camera.ProtoReflect.Descriptor instead.
func (*Camera) Descriptor() ([]byte, []int) {
	return file_tuning_proto_rawDescGZIP(), []int{2}
}

func (x *Camera) GetTiltSeconds() float32 {
	if x != nil && x.TiltSeconds != nil {
		return *x.TiltSeconds
	}
	return Default_Camera_TiltSeconds
}

type World struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Negative is down.
	Gravity *float32 `protobuf:"fixed32,1,opt,name=gravity,def=-1200" json:"gravity,omitempty"`
	Bg      *Color   `protobuf:"bytes,2,opt,name=bg" json:"bg,omitempty"`
}

// Default values for World fields.
const (
	Default_World_Gravity = float32(-1200)
)

func (x *World) Reset() {
	*x = World{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tuning_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *World) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*World) ProtoMessage() {}

func (x *World) ProtoReflect() protoreflect.Message {
	mi := &file_tuning_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use World.ProtoReflect.Descriptor instead.
func (*World) Descriptor() ([]byte, []int) {
	return file_tuning_proto_rawDescGZIP(), []int{3}
}

func (x *World) GetGravity() float32 {
	if x != nil && x.Gravity != nil {
		return *x.Gravity
	}
	return Default_World_Gravity
}

func (x *World) GetBg() *Color {
	if x != nil {
		return x.Bg
	}
	return nil
}

type Blink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Odds         *float32 `protobuf:"fixed32,14,opt,name=odds,def=0.0003" json:"odds,omitempty"`
	CycleSeconds *float32 `protobuf:"fixed32,15,opt,name=cycle_seconds,json=cycleSeconds,def=0.25" json:"cycle_seconds,omitempty"`
}

// Default values for Blink fields.
const (
	Default_Blink_Odds         = float32(0.0003000000142492354)
	Default_Blink_CycleSeconds = float32(0.25)
)

func (x *Blink) Reset() {
	*x = Blink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tuning_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blink) ProtoMessage() {}

func (x *Blink) ProtoReflect() protoreflect.Message {
	mi := &file_tuning_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blink.ProtoReflect.Descriptor instead.
func (*Blink) Descriptor() ([]byte, []int) {
	return file_tuning_proto_rawDescGZIP(), []int{4}
}

func (x *Blink) GetOdds() float32 {
	if x != nil && x.Odds != nil {
		return *x.Odds
	}
	return Default_Blink_Odds
}

func (x *Blink) GetCycleSeconds() float32 {
	if x != nil && x.CycleSeconds != nil {
		return *x.CycleSeconds
	}
	return Default_Blink_CycleSeconds
}

type Todd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color      *Color   `protobuf:"bytes,1,opt,name=color" json:"color,omitempty"`
	SideLength *float32 `protobuf:"fixed32,3,opt,name=side_length,json=sideLength,def=30" json:"side_length,omitempty"`
	// A unitless constant that we scale velocity by while on the ground.
	Friction *float32 `protobuf:"fixed32,4,opt,name=friction,def=0.97" json:"friction,omitempty"`
	// A unitless constant that we scale bearing by when not accelerating.
	BearingFriction     *float32 `protobuf:"fixed32,5,opt,name=bearing_friction,json=bearingFriction,def=0.95" json:"bearing_friction,omitempty"`
	MaxVelocity         *float32 `protobuf:"fixed32,6,opt,name=max_velocity,json=maxVelocity,def=240" json:"max_velocity,omitempty"`
	Acceleration        *float32 `protobuf:"fixed32,7,opt,name=acceleration,def=900" json:"acceleration,omitempty"`
	AirBending          *float32 `protobuf:"fixed32,8,opt,name=air_bending,json=airBending,def=575" json:"air_bending,omitempty"`
	BearingAcceleration *float32 `protobuf:"fixed32,9,opt,name=bearing_acceleration,json=bearingAcceleration,def=1200" json:"bearing_acceleration,omitempty"`
	JumpImpulse         *float32 `protobuf:"fixed32,10,opt,name=jump_impulse,json=jumpImpulse,def=350" json:"jump_impulse,omitempty"`
	MaxSquishVelocity   *float32 `protobuf:"fixed32,11,opt,name=max_squish_velocity,json=maxSquishVelocity,def=60" json:"max_squish_velocity,omitempty"`
	// Max vertical velocity while holding down jump.
	JumpTerminalVelocity *float32 `protobuf:"fixed32,12,opt,name=jump_terminal_velocity,json=jumpTerminalVelocity,def=-350" json:"jump_terminal_velocity,omitempty"`
	TerminalVelocity     *float32 `protobuf:"fixed32,13,opt,name=terminal_velocity,json=terminalVelocity,def=-550" json:"terminal_velocity,omitempty"`
	// Eye centering speed for tumbling/landing.
	EyeCenteringDurationSeconds *float32 `protobuf:"fixed32,16,opt,name=eye_centering_duration_seconds,json=eyeCenteringDurationSeconds,def=1" json:"eye_centering_duration_seconds,omitempty"`
	JumpStateGravityFactor      *float32 `protobuf:"fixed32,17,opt,name=jump_state_gravity_factor,json=jumpStateGravityFactor,def=0.55" json:"jump_state_gravity_factor,omitempty"`
	// A unitless constant that we scale velocity by while in the air.
	// Thanks, Jonah!
	AirFriction            *float32 `protobuf:"fixed32,18,opt,name=air_friction,json=airFriction,def=0.99" json:"air_friction,omitempty"`
	JumpRequestSlopSeconds *float32 `protobuf:"fixed32,19,opt,name=jump_request_slop_seconds,json=jumpRequestSlopSeconds,def=0.2" json:"jump_request_slop_seconds,omitempty"`
	GroundingSlopSeconds   *float32 `protobuf:"fixed32,20,opt,name=grounding_slop_seconds,json=groundingSlopSeconds,def=0.2" json:"grounding_slop_seconds,omitempty"`
	Blink                  *Blink   `protobuf:"bytes,21,opt,name=blink" json:"blink,omitempty"`
}

// Default values for Todd fields.
const (
	Default_Todd_SideLength                  = float32(30)
	Default_Todd_Friction                    = float32(0.9700000286102295)
	Default_Todd_BearingFriction             = float32(0.949999988079071)
	Default_Todd_MaxVelocity                 = float32(240)
	Default_Todd_Acceleration                = float32(900)
	Default_Todd_AirBending                  = float32(575)
	Default_Todd_BearingAcceleration         = float32(1200)
	Default_Todd_JumpImpulse                 = float32(350)
	Default_Todd_MaxSquishVelocity           = float32(60)
	Default_Todd_JumpTerminalVelocity        = float32(-350)
	Default_Todd_TerminalVelocity            = float32(-550)
	Default_Todd_EyeCenteringDurationSeconds = float32(1)
	Default_Todd_JumpStateGravityFactor      = float32(0.550000011920929)
	Default_Todd_AirFriction                 = float32(0.9900000095367432)
	Default_Todd_JumpRequestSlopSeconds      = float32(0.20000000298023224)
	Default_Todd_GroundingSlopSeconds        = float32(0.20000000298023224)
)

func (x *Todd) Reset() {
	*x = Todd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tuning_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todd) ProtoMessage() {}

func (x *Todd) ProtoReflect() protoreflect.Message {
	mi := &file_tuning_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todd.ProtoReflect.Descriptor instead.
func (*Todd) Descriptor() ([]byte, []int) {
	return file_tuning_proto_rawDescGZIP(), []int{5}
}

func (x *Todd) GetColor() *Color {
	if x != nil {
		return x.Color
	}
	return nil
}

func (x *Todd) GetSideLength() float32 {
	if x != nil && x.SideLength != nil {
		return *x.SideLength
	}
	return Default_Todd_SideLength
}

func (x *Todd) GetFriction() float32 {
	if x != nil && x.Friction != nil {
		return *x.Friction
	}
	return Default_Todd_Friction
}

func (x *Todd) GetBearingFriction() float32 {
	if x != nil && x.BearingFriction != nil {
		return *x.BearingFriction
	}
	return Default_Todd_BearingFriction
}

func (x *Todd) GetMaxVelocity() float32 {
	if x != nil && x.MaxVelocity != nil {
		return *x.MaxVelocity
	}
	return Default_Todd_MaxVelocity
}

func (x *Todd) GetAcceleration() float32 {
	if x != nil && x.Acceleration != nil {
		return *x.Acceleration
	}
	return Default_Todd_Acceleration
}

func (x *Todd) GetAirBending() float32 {
	if x != nil && x.AirBending != nil {
		return *x.AirBending
	}
	return Default_Todd_AirBending
}

func (x *Todd) GetBearingAcceleration() float32 {
	if x != nil && x.BearingAcceleration != nil {
		return *x.BearingAcceleration
	}
	return Default_Todd_BearingAcceleration
}

func (x *Todd) GetJumpImpulse() float32 {
	if x != nil && x.JumpImpulse != nil {
		return *x.JumpImpulse
	}
	return Default_Todd_JumpImpulse
}

func (x *Todd) GetMaxSquishVelocity() float32 {
	if x != nil && x.MaxSquishVelocity != nil {
		return *x.MaxSquishVelocity
	}
	return Default_Todd_MaxSquishVelocity
}

func (x *Todd) GetJumpTerminalVelocity() float32 {
	if x != nil && x.JumpTerminalVelocity != nil {
		return *x.JumpTerminalVelocity
	}
	return Default_Todd_JumpTerminalVelocity
}

func (x *Todd) GetTerminalVelocity() float32 {
	if x != nil && x.TerminalVelocity != nil {
		return *x.TerminalVelocity
	}
	return Default_Todd_TerminalVelocity
}

func (x *Todd) GetEyeCenteringDurationSeconds() float32 {
	if x != nil && x.EyeCenteringDurationSeconds != nil {
		return *x.EyeCenteringDurationSeconds
	}
	return Default_Todd_EyeCenteringDurationSeconds
}

func (x *Todd) GetJumpStateGravityFactor() float32 {
	if x != nil && x.JumpStateGravityFactor != nil {
		return *x.JumpStateGravityFactor
	}
	return Default_Todd_JumpStateGravityFactor
}

func (x *Todd) GetAirFriction() float32 {
	if x != nil && x.AirFriction != nil {
		return *x.AirFriction
	}
	return Default_Todd_AirFriction
}

func (x *Todd) GetJumpRequestSlopSeconds() float32 {
	if x != nil && x.JumpRequestSlopSeconds != nil {
		return *x.JumpRequestSlopSeconds
	}
	return Default_Todd_JumpRequestSlopSeconds
}

func (x *Todd) GetGroundingSlopSeconds() float32 {
	if x != nil && x.GroundingSlopSeconds != nil {
		return *x.GroundingSlopSeconds
	}
	return Default_Todd_GroundingSlopSeconds
}

func (x *Todd) GetBlink() *Blink {
	if x != nil {
		return x.Blink
	}
	return nil
}

var file_tuning_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50000,
		Name:          "game.top_level",
		Tag:           "varint,50000,opt,name=top_level",
		Filename:      "tuning.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*float32)(nil),
		Field:         50000,
		Name:          "game.min",
		Tag:           "fixed32,50000,opt,name=min",
		Filename:      "tuning.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*float32)(nil),
		Field:         50001,
		Name:          "game.max",
		Tag:           "fixed32,50001,opt,name=max",
		Filename:      "tuning.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional bool top_level = 50000;
	E_TopLevel = &file_tuning_proto_extTypes[0]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional float min = 50000;
	E_Min = &file_tuning_proto_extTypes[1]
	// optional float max = 50001;
	E_Max = &file_tuning_proto_extTypes[2]
)

var File_tuning_proto protoreflect.FileDescriptor

var file_tuning_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x67, 0x61, 0x6d, 0x65, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x19, 0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12,
	0x10, 0x0a, 0x01, 0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x02, 0x42, 0x02, 0x10, 0x01, 0x52, 0x01,
	0x63, 0x22, 0x77, 0x0a, 0x06, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x05, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x05, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x1e,
	0x0a, 0x04, 0x74, 0x6f, 0x64, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x54, 0x6f, 0x64, 0x64, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x64, 0x12, 0x24,
	0x0a, 0x06, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x52, 0x06, 0x63, 0x61,
	0x6d, 0x65, 0x72, 0x61, 0x3a, 0x04, 0x80, 0xb5, 0x18, 0x01, 0x22, 0x40, 0x0a, 0x06, 0x43, 0x61,
	0x6d, 0x65, 0x72, 0x61, 0x12, 0x36, 0x0a, 0x0c, 0x74, 0x69, 0x6c, 0x74, 0x5f, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x03, 0x30, 0x2e, 0x32, 0x42,
	0x0e, 0x85, 0xb5, 0x18, 0x00, 0x00, 0x00, 0x00, 0x8d, 0xb5, 0x18, 0xcd, 0xcc, 0x4c, 0x3f, 0x52,
	0x0b, 0x74, 0x69, 0x6c, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x55, 0x0a, 0x05,
	0x57, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x2f, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x05, 0x2d, 0x31, 0x32, 0x30, 0x30, 0x42, 0x0e, 0x85,
	0xb5, 0x18, 0x00, 0x00, 0xfa, 0xc4, 0x8d, 0xb5, 0x18, 0x00, 0x00, 0xc8, 0xc2, 0x52, 0x07, 0x67,
	0x72, 0x61, 0x76, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x02, 0x62, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52,
	0x02, 0x62, 0x67, 0x22, 0x6e, 0x0a, 0x05, 0x42, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x2a, 0x0a, 0x04,
	0x6f, 0x64, 0x64, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x06, 0x30, 0x2e, 0x30, 0x30,
	0x30, 0x33, 0x42, 0x0e, 0x85, 0xb5, 0x18, 0xac, 0xc5, 0x27, 0x37, 0x8d, 0xb5, 0x18, 0x6f, 0x12,
	0x83, 0x3a, 0x52, 0x04, 0x6f, 0x64, 0x64, 0x73, 0x12, 0x39, 0x0a, 0x0d, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x02, 0x3a,
	0x04, 0x30, 0x2e, 0x32, 0x35, 0x42, 0x0e, 0x85, 0xb5, 0x18, 0x00, 0x00, 0x00, 0x00, 0x8d, 0xb5,
	0x18, 0x00, 0x00, 0x00, 0x40, 0x52, 0x0c, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x22, 0xdc, 0x08, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x64, 0x12, 0x21, 0x0a, 0x05,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12,
	0x23, 0x0a, 0x0b, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x3a, 0x02, 0x33, 0x30, 0x52, 0x0a, 0x73, 0x69, 0x64, 0x65, 0x4c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x12, 0x30, 0x0a, 0x08, 0x66, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x04, 0x30, 0x2e, 0x39, 0x37, 0x42, 0x0e, 0x85, 0xb5,
	0x18, 0x66, 0x66, 0x66, 0x3f, 0x8d, 0xb5, 0x18, 0x00, 0x00, 0x80, 0x3f, 0x52, 0x08, 0x66, 0x72,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x10, 0x62, 0x65, 0x61, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x66, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02,
	0x3a, 0x04, 0x30, 0x2e, 0x39, 0x35, 0x42, 0x0e, 0x85, 0xb5, 0x18, 0x66, 0x66, 0x66, 0x3f, 0x8d,
	0xb5, 0x18, 0x00, 0x00, 0x80, 0x3f, 0x52, 0x0f, 0x62, 0x65, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x46,
	0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x76,
	0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x03, 0x32,
	0x34, 0x30, 0x42, 0x0e, 0x85, 0xb5, 0x18, 0x00, 0x00, 0xc8, 0x42, 0x8d, 0xb5, 0x18, 0x00, 0x00,
	0xaf, 0x43, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x37, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x03, 0x39, 0x30, 0x30, 0x42, 0x0e, 0x85, 0xb5, 0x18, 0x00,
	0x00, 0x48, 0x43, 0x8d, 0xb5, 0x18, 0x00, 0x00, 0xfa, 0x44, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65,
	0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x0b, 0x61, 0x69, 0x72, 0x5f,
	0x62, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x03, 0x35,
	0x37, 0x35, 0x42, 0x0e, 0x85, 0xb5, 0x18, 0x00, 0x00, 0x48, 0x43, 0x8d, 0xb5, 0x18, 0x00, 0x00,
	0x48, 0x44, 0x52, 0x0a, 0x61, 0x69, 0x72, 0x42, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x47,
	0x0a, 0x14, 0x62, 0x65, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x04, 0x31, 0x32,
	0x30, 0x30, 0x42, 0x0e, 0x85, 0xb5, 0x18, 0x00, 0x00, 0x48, 0x43, 0x8d, 0xb5, 0x18, 0x00, 0x00,
	0xfa, 0x44, 0x52, 0x13, 0x62, 0x65, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x65, 0x6c,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x0c, 0x6a, 0x75, 0x6d, 0x70, 0x5f,
	0x69, 0x6d, 0x70, 0x75, 0x6c, 0x73, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x03, 0x33,
	0x35, 0x30, 0x42, 0x0e, 0x85, 0xb5, 0x18, 0x00, 0x00, 0xc8, 0x42, 0x8d, 0xb5, 0x18, 0x00, 0x00,
	0x48, 0x44, 0x52, 0x0b, 0x6a, 0x75, 0x6d, 0x70, 0x49, 0x6d, 0x70, 0x75, 0x6c, 0x73, 0x65, 0x12,
	0x42, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x71, 0x75, 0x69, 0x73, 0x68, 0x5f, 0x76, 0x65,
	0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x02, 0x36, 0x30,
	0x42, 0x0e, 0x85, 0xb5, 0x18, 0x00, 0x00, 0xa0, 0x41, 0x8d, 0xb5, 0x18, 0x00, 0x00, 0xf0, 0x42,
	0x52, 0x11, 0x6d, 0x61, 0x78, 0x53, 0x71, 0x75, 0x69, 0x73, 0x68, 0x56, 0x65, 0x6c, 0x6f, 0x63,
	0x69, 0x74, 0x79, 0x12, 0x4a, 0x0a, 0x16, 0x6a, 0x75, 0x6d, 0x70, 0x5f, 0x74, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x02, 0x3a, 0x04, 0x2d, 0x33, 0x35, 0x30, 0x42, 0x0e, 0x85, 0xb5, 0x18, 0x00, 0x00,
	0x16, 0xc4, 0x8d, 0xb5, 0x18, 0x00, 0x00, 0xc8, 0xc2, 0x52, 0x14, 0x6a, 0x75, 0x6d, 0x70, 0x54,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x41, 0x0a, 0x11, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x76, 0x65, 0x6c, 0x6f,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x04, 0x2d, 0x35, 0x35, 0x30,
	0x42, 0x0e, 0x85, 0xb5, 0x18, 0x00, 0x00, 0x7a, 0xc4, 0x8d, 0xb5, 0x18, 0x00, 0x00, 0xc8, 0xc2,
	0x52, 0x10, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x56, 0x0a, 0x1e, 0x65, 0x79, 0x65, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x01, 0x31, 0x42, 0x0e, 0x85,
	0xb5, 0x18, 0x00, 0x00, 0x00, 0x00, 0x8d, 0xb5, 0x18, 0x00, 0x00, 0x40, 0x40, 0x52, 0x1b, 0x65,
	0x79, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x4f, 0x0a, 0x19, 0x6a, 0x75,
	0x6d, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x79,
	0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x04, 0x30,
	0x2e, 0x35, 0x35, 0x42, 0x0e, 0x85, 0xb5, 0x18, 0xcd, 0xcc, 0x4c, 0x3e, 0x8d, 0xb5, 0x18, 0x00,
	0x00, 0x80, 0x3f, 0x52, 0x16, 0x6a, 0x75, 0x6d, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x47, 0x72,
	0x61, 0x76, 0x69, 0x74, 0x79, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x37, 0x0a, 0x0c, 0x61,
	0x69, 0x72, 0x5f, 0x66, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x02, 0x3a, 0x04, 0x30, 0x2e, 0x39, 0x39, 0x42, 0x0e, 0x85, 0xb5, 0x18, 0x9a, 0x99, 0x59, 0x3f,
	0x8d, 0xb5, 0x18, 0x00, 0x00, 0x80, 0x3f, 0x52, 0x0b, 0x61, 0x69, 0x72, 0x46, 0x72, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x19, 0x6a, 0x75, 0x6d, 0x70, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x6c, 0x6f, 0x70, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x03, 0x30, 0x2e, 0x32, 0x42, 0x0e, 0x85, 0xb5,
	0x18, 0x00, 0x00, 0x00, 0x00, 0x8d, 0xb5, 0x18, 0x00, 0x00, 0x00, 0x3f, 0x52, 0x16, 0x6a, 0x75,
	0x6d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x6c, 0x6f, 0x70, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x12, 0x49, 0x0a, 0x16, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x6c, 0x6f, 0x70, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x02, 0x3a, 0x03, 0x30, 0x2e, 0x32, 0x42, 0x0e, 0x85, 0xb5, 0x18, 0x00, 0x00,
	0x00, 0x00, 0x8d, 0xb5, 0x18, 0x00, 0x00, 0x00, 0x3f, 0x52, 0x14, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x53, 0x6c, 0x6f, 0x70, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12,
	0x21, 0x0a, 0x05, 0x62, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x42, 0x6c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x62, 0x6c, 0x69,
	0x6e, 0x6b, 0x3a, 0x3e, 0x0a, 0x09, 0x74, 0x6f, 0x70, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x74, 0x6f, 0x70, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x3a, 0x31, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x03, 0x6d, 0x69, 0x6e, 0x3a, 0x31, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd1, 0x86, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x64, 0x66, 0x2f, 0x74, 0x6f, 0x64, 0x64, 0x2d,
	0x61, 0x67, 0x61, 0x69, 0x6e, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_tuning_proto_rawDescOnce sync.Once
	file_tuning_proto_rawDescData = file_tuning_proto_rawDesc
)

func file_tuning_proto_rawDescGZIP() []byte {
	file_tuning_proto_rawDescOnce.Do(func() {
		file_tuning_proto_rawDescData = protoimpl.X.CompressGZIP(file_tuning_proto_rawDescData)
	})
	return file_tuning_proto_rawDescData
}

var file_tuning_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tuning_proto_goTypes = []interface{}{
	(*Color)(nil),                       // 0: game.Color
	(*Tuning)(nil),                      // 1: game.Tuning
	(*Camera)(nil),                      // 2: game.Camera
	(*World)(nil),                       // 3: game.World
	(*Blink)(nil),                       // 4: game.Blink
	(*Todd)(nil),                        // 5: game.Todd
	(*descriptorpb.MessageOptions)(nil), // 6: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 7: google.protobuf.FieldOptions
}
var file_tuning_proto_depIdxs = []int32{
	3, // 0: game.Tuning.world:type_name -> game.World
	5, // 1: game.Tuning.todd:type_name -> game.Todd
	2, // 2: game.Tuning.camera:type_name -> game.Camera
	0, // 3: game.World.bg:type_name -> game.Color
	0, // 4: game.Todd.color:type_name -> game.Color
	4, // 5: game.Todd.blink:type_name -> game.Blink
	6, // 6: game.top_level:extendee -> google.protobuf.MessageOptions
	7, // 7: game.min:extendee -> google.protobuf.FieldOptions
	7, // 8: game.max:extendee -> google.protobuf.FieldOptions
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	6, // [6:9] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_tuning_proto_init() }
func file_tuning_proto_init() {
	if File_tuning_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tuning_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Color); i {
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
		file_tuning_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tuning); i {
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
		file_tuning_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Camera); i {
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
		file_tuning_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*World); i {
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
		file_tuning_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blink); i {
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
		file_tuning_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todd); i {
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
			RawDescriptor: file_tuning_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_tuning_proto_goTypes,
		DependencyIndexes: file_tuning_proto_depIdxs,
		MessageInfos:      file_tuning_proto_msgTypes,
		ExtensionInfos:    file_tuning_proto_extTypes,
	}.Build()
	File_tuning_proto = out.File
	file_tuning_proto_rawDesc = nil
	file_tuning_proto_goTypes = nil
	file_tuning_proto_depIdxs = nil
}
