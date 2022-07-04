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
	Red   *float32 `protobuf:"fixed32,1,opt,name=red,def=0" json:"red,omitempty"`
	Green *float32 `protobuf:"fixed32,2,opt,name=green,def=0" json:"green,omitempty"`
	Blue  *float32 `protobuf:"fixed32,3,opt,name=blue,def=0" json:"blue,omitempty"`
	Alpha *float32 `protobuf:"fixed32,4,opt,name=alpha,def=1" json:"alpha,omitempty"`
}

// Default values for Color fields.
const (
	Default_Color_Red   = float32(0)
	Default_Color_Green = float32(0)
	Default_Color_Blue  = float32(0)
	Default_Color_Alpha = float32(1)
)

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

func (x *Color) GetRed() float32 {
	if x != nil && x.Red != nil {
		return *x.Red
	}
	return Default_Color_Red
}

func (x *Color) GetGreen() float32 {
	if x != nil && x.Green != nil {
		return *x.Green
	}
	return Default_Color_Green
}

func (x *Color) GetBlue() float32 {
	if x != nil && x.Blue != nil {
		return *x.Blue
	}
	return Default_Color_Blue
}

func (x *Color) GetAlpha() float32 {
	if x != nil && x.Alpha != nil {
		return *x.Alpha
	}
	return Default_Color_Alpha
}

// All distances are in world coordinates, all durations are in seconds.
type Tuning struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Negative is down.
	Gravity        *float32 `protobuf:"fixed32,1,opt,name=gravity,def=-1200" json:"gravity,omitempty"`
	Bg             *Color   `protobuf:"bytes,2,opt,name=bg" json:"bg,omitempty"`
	ToddSideLength *float32 `protobuf:"fixed32,3,opt,name=todd_side_length,json=toddSideLength,def=30" json:"todd_side_length,omitempty"`
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
	BlinkOdds            *float32 `protobuf:"fixed32,14,opt,name=blink_odds,json=blinkOdds,def=0.0003" json:"blink_odds,omitempty"`
	BlinkCycleSeconds    *float32 `protobuf:"fixed32,15,opt,name=blink_cycle_seconds,json=blinkCycleSeconds,def=0.25" json:"blink_cycle_seconds,omitempty"`
	// Eye centering speed for tumbling/landing.
	EyeCenteringDurationSeconds *float32 `protobuf:"fixed32,16,opt,name=eye_centering_duration_seconds,json=eyeCenteringDurationSeconds,def=1" json:"eye_centering_duration_seconds,omitempty"`
	JumpStateGravityFactor      *float32 `protobuf:"fixed32,17,opt,name=jump_state_gravity_factor,json=jumpStateGravityFactor,def=0.55" json:"jump_state_gravity_factor,omitempty"`
	CameraTiltSeconds           *float32 `protobuf:"fixed32,18,opt,name=camera_tilt_seconds,json=cameraTiltSeconds,def=0.2" json:"camera_tilt_seconds,omitempty"`
	JumpRequestSlopSeconds      *float32 `protobuf:"fixed32,19,opt,name=jump_request_slop_seconds,json=jumpRequestSlopSeconds,def=0.2" json:"jump_request_slop_seconds,omitempty"`
	GroundingSlopSeconds        *float32 `protobuf:"fixed32,20,opt,name=grounding_slop_seconds,json=groundingSlopSeconds,def=0.2" json:"grounding_slop_seconds,omitempty"`
	// A unitless constant that we scale velocity by while in the air.
	// Thanks, Jonah!
	AirFriction *float32 `protobuf:"fixed32,21,opt,name=air_friction,json=airFriction,def=0.99" json:"air_friction,omitempty"`
}

// Default values for Tuning fields.
const (
	Default_Tuning_Gravity                     = float32(-1200)
	Default_Tuning_ToddSideLength              = float32(30)
	Default_Tuning_Friction                    = float32(0.9700000286102295)
	Default_Tuning_BearingFriction             = float32(0.949999988079071)
	Default_Tuning_MaxVelocity                 = float32(240)
	Default_Tuning_Acceleration                = float32(900)
	Default_Tuning_AirBending                  = float32(575)
	Default_Tuning_BearingAcceleration         = float32(1200)
	Default_Tuning_JumpImpulse                 = float32(350)
	Default_Tuning_MaxSquishVelocity           = float32(60)
	Default_Tuning_JumpTerminalVelocity        = float32(-350)
	Default_Tuning_TerminalVelocity            = float32(-550)
	Default_Tuning_BlinkOdds                   = float32(0.0003000000142492354)
	Default_Tuning_BlinkCycleSeconds           = float32(0.25)
	Default_Tuning_EyeCenteringDurationSeconds = float32(1)
	Default_Tuning_JumpStateGravityFactor      = float32(0.550000011920929)
	Default_Tuning_CameraTiltSeconds           = float32(0.20000000298023224)
	Default_Tuning_JumpRequestSlopSeconds      = float32(0.20000000298023224)
	Default_Tuning_GroundingSlopSeconds        = float32(0.20000000298023224)
	Default_Tuning_AirFriction                 = float32(0.9900000095367432)
)

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

func (x *Tuning) GetGravity() float32 {
	if x != nil && x.Gravity != nil {
		return *x.Gravity
	}
	return Default_Tuning_Gravity
}

func (x *Tuning) GetBg() *Color {
	if x != nil {
		return x.Bg
	}
	return nil
}

func (x *Tuning) GetToddSideLength() float32 {
	if x != nil && x.ToddSideLength != nil {
		return *x.ToddSideLength
	}
	return Default_Tuning_ToddSideLength
}

func (x *Tuning) GetFriction() float32 {
	if x != nil && x.Friction != nil {
		return *x.Friction
	}
	return Default_Tuning_Friction
}

func (x *Tuning) GetBearingFriction() float32 {
	if x != nil && x.BearingFriction != nil {
		return *x.BearingFriction
	}
	return Default_Tuning_BearingFriction
}

func (x *Tuning) GetMaxVelocity() float32 {
	if x != nil && x.MaxVelocity != nil {
		return *x.MaxVelocity
	}
	return Default_Tuning_MaxVelocity
}

func (x *Tuning) GetAcceleration() float32 {
	if x != nil && x.Acceleration != nil {
		return *x.Acceleration
	}
	return Default_Tuning_Acceleration
}

func (x *Tuning) GetAirBending() float32 {
	if x != nil && x.AirBending != nil {
		return *x.AirBending
	}
	return Default_Tuning_AirBending
}

func (x *Tuning) GetBearingAcceleration() float32 {
	if x != nil && x.BearingAcceleration != nil {
		return *x.BearingAcceleration
	}
	return Default_Tuning_BearingAcceleration
}

func (x *Tuning) GetJumpImpulse() float32 {
	if x != nil && x.JumpImpulse != nil {
		return *x.JumpImpulse
	}
	return Default_Tuning_JumpImpulse
}

func (x *Tuning) GetMaxSquishVelocity() float32 {
	if x != nil && x.MaxSquishVelocity != nil {
		return *x.MaxSquishVelocity
	}
	return Default_Tuning_MaxSquishVelocity
}

func (x *Tuning) GetJumpTerminalVelocity() float32 {
	if x != nil && x.JumpTerminalVelocity != nil {
		return *x.JumpTerminalVelocity
	}
	return Default_Tuning_JumpTerminalVelocity
}

func (x *Tuning) GetTerminalVelocity() float32 {
	if x != nil && x.TerminalVelocity != nil {
		return *x.TerminalVelocity
	}
	return Default_Tuning_TerminalVelocity
}

func (x *Tuning) GetBlinkOdds() float32 {
	if x != nil && x.BlinkOdds != nil {
		return *x.BlinkOdds
	}
	return Default_Tuning_BlinkOdds
}

func (x *Tuning) GetBlinkCycleSeconds() float32 {
	if x != nil && x.BlinkCycleSeconds != nil {
		return *x.BlinkCycleSeconds
	}
	return Default_Tuning_BlinkCycleSeconds
}

func (x *Tuning) GetEyeCenteringDurationSeconds() float32 {
	if x != nil && x.EyeCenteringDurationSeconds != nil {
		return *x.EyeCenteringDurationSeconds
	}
	return Default_Tuning_EyeCenteringDurationSeconds
}

func (x *Tuning) GetJumpStateGravityFactor() float32 {
	if x != nil && x.JumpStateGravityFactor != nil {
		return *x.JumpStateGravityFactor
	}
	return Default_Tuning_JumpStateGravityFactor
}

func (x *Tuning) GetCameraTiltSeconds() float32 {
	if x != nil && x.CameraTiltSeconds != nil {
		return *x.CameraTiltSeconds
	}
	return Default_Tuning_CameraTiltSeconds
}

func (x *Tuning) GetJumpRequestSlopSeconds() float32 {
	if x != nil && x.JumpRequestSlopSeconds != nil {
		return *x.JumpRequestSlopSeconds
	}
	return Default_Tuning_JumpRequestSlopSeconds
}

func (x *Tuning) GetGroundingSlopSeconds() float32 {
	if x != nil && x.GroundingSlopSeconds != nil {
		return *x.GroundingSlopSeconds
	}
	return Default_Tuning_GroundingSlopSeconds
}

func (x *Tuning) GetAirFriction() float32 {
	if x != nil && x.AirFriction != nil {
		return *x.AirFriction
	}
	return Default_Tuning_AirFriction
}

var file_tuning_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50000,
		Name:          "game.section",
		Tag:           "bytes,50000,opt,name=section",
		Filename:      "tuning.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*float32)(nil),
		Field:         50001,
		Name:          "game.min",
		Tag:           "fixed32,50001,opt,name=min",
		Filename:      "tuning.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*float32)(nil),
		Field:         50002,
		Name:          "game.max",
		Tag:           "fixed32,50002,opt,name=max",
		Filename:      "tuning.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional string section = 50000;
	E_Section = &file_tuning_proto_extTypes[0]
	// optional float min = 50001;
	E_Min = &file_tuning_proto_extTypes[1]
	// optional float max = 50002;
	E_Max = &file_tuning_proto_extTypes[2]
)

var File_tuning_proto protoreflect.FileDescriptor

var file_tuning_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x67, 0x61, 0x6d, 0x65, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12,
	0x13, 0x0a, 0x03, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x01, 0x30, 0x52,
	0x03, 0x72, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x05, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x3a, 0x01, 0x30, 0x52, 0x05, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x12, 0x15, 0x0a,
	0x04, 0x62, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x01, 0x30, 0x52, 0x04,
	0x62, 0x6c, 0x75, 0x65, 0x12, 0x17, 0x0a, 0x05, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x3a, 0x01, 0x31, 0x52, 0x05, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x22, 0xde, 0x0b,
	0x0a, 0x06, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x38, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x76,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x05, 0x2d, 0x31, 0x32, 0x30, 0x30,
	0x42, 0x17, 0x82, 0xb5, 0x18, 0x05, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x8d, 0xb5, 0x18, 0x00, 0x00,
	0xfa, 0xc4, 0x95, 0xb5, 0x18, 0x00, 0x00, 0xc8, 0xc2, 0x52, 0x07, 0x67, 0x72, 0x61, 0x76, 0x69,
	0x74, 0x79, 0x12, 0x25, 0x0a, 0x02, 0x62, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x42, 0x08, 0x82, 0xb5, 0x18,
	0x04, 0x54, 0x6f, 0x64, 0x64, 0x52, 0x02, 0x62, 0x67, 0x12, 0x36, 0x0a, 0x10, 0x74, 0x6f, 0x64,
	0x64, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x3a, 0x02, 0x33, 0x30, 0x42, 0x08, 0x82, 0xb5, 0x18, 0x04, 0x54, 0x6f, 0x64,
	0x64, 0x52, 0x0e, 0x74, 0x6f, 0x64, 0x64, 0x53, 0x69, 0x64, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x12, 0x38, 0x0a, 0x08, 0x66, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x3a, 0x04, 0x30, 0x2e, 0x39, 0x37, 0x42, 0x16, 0x82, 0xb5, 0x18, 0x04, 0x54,
	0x6f, 0x64, 0x64, 0x8d, 0xb5, 0x18, 0x66, 0x66, 0x66, 0x3f, 0x95, 0xb5, 0x18, 0x00, 0x00, 0x80,
	0x3f, 0x52, 0x08, 0x66, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x10, 0x62,
	0x65, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x04, 0x30, 0x2e, 0x39, 0x35, 0x42, 0x16, 0x82, 0xb5, 0x18,
	0x04, 0x54, 0x6f, 0x64, 0x64, 0x8d, 0xb5, 0x18, 0x66, 0x66, 0x66, 0x3f, 0x95, 0xb5, 0x18, 0x00,
	0x00, 0x80, 0x3f, 0x52, 0x0f, 0x62, 0x65, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x72, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x65, 0x6c, 0x6f,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x03, 0x32, 0x34, 0x30, 0x42,
	0x16, 0x82, 0xb5, 0x18, 0x04, 0x54, 0x6f, 0x64, 0x64, 0x8d, 0xb5, 0x18, 0x00, 0x00, 0xc8, 0x42,
	0x95, 0xb5, 0x18, 0x00, 0x00, 0xaf, 0x43, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x56, 0x65, 0x6c, 0x6f,
	0x63, 0x69, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x03, 0x39, 0x30, 0x30, 0x42,
	0x16, 0x82, 0xb5, 0x18, 0x04, 0x54, 0x6f, 0x64, 0x64, 0x8d, 0xb5, 0x18, 0x00, 0x00, 0x48, 0x43,
	0x95, 0xb5, 0x18, 0x00, 0x00, 0xfa, 0x44, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x0b, 0x61, 0x69, 0x72, 0x5f, 0x62, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x03, 0x35, 0x37, 0x35, 0x42,
	0x16, 0x82, 0xb5, 0x18, 0x04, 0x54, 0x6f, 0x64, 0x64, 0x8d, 0xb5, 0x18, 0x00, 0x00, 0x48, 0x43,
	0x95, 0xb5, 0x18, 0x00, 0x00, 0x48, 0x44, 0x52, 0x0a, 0x61, 0x69, 0x72, 0x42, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x12, 0x4f, 0x0a, 0x14, 0x62, 0x65, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x61,
	0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x02, 0x3a, 0x04, 0x31, 0x32, 0x30, 0x30, 0x42, 0x16, 0x82, 0xb5, 0x18, 0x04, 0x54, 0x6f, 0x64,
	0x64, 0x8d, 0xb5, 0x18, 0x00, 0x00, 0x48, 0x43, 0x95, 0xb5, 0x18, 0x00, 0x00, 0xfa, 0x44, 0x52,
	0x13, 0x62, 0x65, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0c, 0x6a, 0x75, 0x6d, 0x70, 0x5f, 0x69, 0x6d, 0x70,
	0x75, 0x6c, 0x73, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x03, 0x33, 0x35, 0x30, 0x42,
	0x16, 0x82, 0xb5, 0x18, 0x04, 0x54, 0x6f, 0x64, 0x64, 0x8d, 0xb5, 0x18, 0x00, 0x00, 0xc8, 0x42,
	0x95, 0xb5, 0x18, 0x00, 0x00, 0x48, 0x44, 0x52, 0x0b, 0x6a, 0x75, 0x6d, 0x70, 0x49, 0x6d, 0x70,
	0x75, 0x6c, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x71, 0x75, 0x69,
	0x73, 0x68, 0x5f, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x02, 0x3a, 0x02, 0x36, 0x30, 0x42, 0x16, 0x82, 0xb5, 0x18, 0x04, 0x54, 0x6f, 0x64, 0x64, 0x8d,
	0xb5, 0x18, 0x00, 0x00, 0xa0, 0x41, 0x95, 0xb5, 0x18, 0x00, 0x00, 0xf0, 0x42, 0x52, 0x11, 0x6d,
	0x61, 0x78, 0x53, 0x71, 0x75, 0x69, 0x73, 0x68, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79,
	0x12, 0x52, 0x0a, 0x16, 0x6a, 0x75, 0x6d, 0x70, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61,
	0x6c, 0x5f, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02,
	0x3a, 0x04, 0x2d, 0x33, 0x35, 0x30, 0x42, 0x16, 0x82, 0xb5, 0x18, 0x04, 0x54, 0x6f, 0x64, 0x64,
	0x8d, 0xb5, 0x18, 0x00, 0x00, 0x16, 0xc4, 0x95, 0xb5, 0x18, 0x00, 0x00, 0xc8, 0xc2, 0x52, 0x14,
	0x6a, 0x75, 0x6d, 0x70, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x56, 0x65, 0x6c, 0x6f,
	0x63, 0x69, 0x74, 0x79, 0x12, 0x49, 0x0a, 0x11, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c,
	0x5f, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x3a,
	0x04, 0x2d, 0x35, 0x35, 0x30, 0x42, 0x16, 0x82, 0xb5, 0x18, 0x04, 0x54, 0x6f, 0x64, 0x64, 0x8d,
	0xb5, 0x18, 0x00, 0x00, 0x7a, 0xc4, 0x95, 0xb5, 0x18, 0x00, 0x00, 0xc8, 0xc2, 0x52, 0x10, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x3d, 0x0a, 0x0a, 0x62, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x6f, 0x64, 0x64, 0x73, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x02, 0x3a, 0x06, 0x30, 0x2e, 0x30, 0x30, 0x30, 0x33, 0x42, 0x16, 0x82, 0xb5, 0x18,
	0x04, 0x54, 0x6f, 0x64, 0x64, 0x8d, 0xb5, 0x18, 0xac, 0xc5, 0x27, 0x37, 0x95, 0xb5, 0x18, 0x6f,
	0x12, 0x83, 0x3a, 0x52, 0x09, 0x62, 0x6c, 0x69, 0x6e, 0x6b, 0x4f, 0x64, 0x64, 0x73, 0x12, 0x4c,
	0x0a, 0x13, 0x62, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x04, 0x30, 0x2e, 0x32,
	0x35, 0x42, 0x16, 0x82, 0xb5, 0x18, 0x04, 0x54, 0x6f, 0x64, 0x64, 0x8d, 0xb5, 0x18, 0x00, 0x00,
	0x00, 0x00, 0x95, 0xb5, 0x18, 0x00, 0x00, 0x00, 0x40, 0x52, 0x11, 0x62, 0x6c, 0x69, 0x6e, 0x6b,
	0x43, 0x79, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x5e, 0x0a, 0x1e,
	0x65, 0x79, 0x65, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x02, 0x3a, 0x01, 0x31, 0x42, 0x16, 0x82, 0xb5, 0x18, 0x04, 0x54, 0x6f, 0x64,
	0x64, 0x8d, 0xb5, 0x18, 0x00, 0x00, 0x00, 0x00, 0x95, 0xb5, 0x18, 0x00, 0x00, 0x40, 0x40, 0x52,
	0x1b, 0x65, 0x79, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x57, 0x0a, 0x19,
	0x6a, 0x75, 0x6d, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x76, 0x69,
	0x74, 0x79, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x02, 0x3a,
	0x04, 0x30, 0x2e, 0x35, 0x35, 0x42, 0x16, 0x82, 0xb5, 0x18, 0x04, 0x54, 0x6f, 0x64, 0x64, 0x8d,
	0xb5, 0x18, 0xcd, 0xcc, 0x4c, 0x3e, 0x95, 0xb5, 0x18, 0x00, 0x00, 0x80, 0x3f, 0x52, 0x16, 0x6a,
	0x75, 0x6d, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x47, 0x72, 0x61, 0x76, 0x69, 0x74, 0x79, 0x46,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x4b, 0x0a, 0x13, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x5f,
	0x74, 0x69, 0x6c, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x02, 0x3a, 0x03, 0x30, 0x2e, 0x32, 0x42, 0x16, 0x82, 0xb5, 0x18, 0x04, 0x54, 0x6f, 0x64,
	0x64, 0x8d, 0xb5, 0x18, 0x00, 0x00, 0x00, 0x00, 0x95, 0xb5, 0x18, 0xcd, 0xcc, 0x4c, 0x3f, 0x52,
	0x11, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x54, 0x69, 0x6c, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x12, 0x56, 0x0a, 0x19, 0x6a, 0x75, 0x6d, 0x70, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x73, 0x6c, 0x6f, 0x70, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x03, 0x30, 0x2e, 0x32, 0x42, 0x16, 0x82, 0xb5, 0x18, 0x04,
	0x54, 0x6f, 0x64, 0x64, 0x8d, 0xb5, 0x18, 0x00, 0x00, 0x00, 0x00, 0x95, 0xb5, 0x18, 0x00, 0x00,
	0x00, 0x3f, 0x52, 0x16, 0x6a, 0x75, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53,
	0x6c, 0x6f, 0x70, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x51, 0x0a, 0x16, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x6c, 0x6f, 0x70, 0x5f, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x03, 0x30, 0x2e, 0x32, 0x42,
	0x16, 0x82, 0xb5, 0x18, 0x04, 0x54, 0x6f, 0x64, 0x64, 0x8d, 0xb5, 0x18, 0x00, 0x00, 0x00, 0x00,
	0x95, 0xb5, 0x18, 0x00, 0x00, 0x00, 0x3f, 0x52, 0x14, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x53, 0x6c, 0x6f, 0x70, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x3f, 0x0a,
	0x0c, 0x61, 0x69, 0x72, 0x5f, 0x66, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x02, 0x3a, 0x04, 0x30, 0x2e, 0x39, 0x39, 0x42, 0x16, 0x82, 0xb5, 0x18, 0x04, 0x54,
	0x6f, 0x64, 0x64, 0x8d, 0xb5, 0x18, 0x9a, 0x99, 0x59, 0x3f, 0x95, 0xb5, 0x18, 0x00, 0x00, 0x80,
	0x3f, 0x52, 0x0b, 0x61, 0x69, 0x72, 0x46, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x39,
	0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x31, 0x0a, 0x03, 0x6d, 0x69, 0x6e,
	0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xd1, 0x86, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x3a, 0x31, 0x0a, 0x03,
	0x6d, 0x61, 0x78, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xd2, 0x86, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x42,
	0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x64,
	0x66, 0x2f, 0x74, 0x6f, 0x64, 0x64, 0x2d, 0x61, 0x67, 0x61, 0x69, 0x6e, 0x2f, 0x67, 0x61, 0x6d,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
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

var file_tuning_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tuning_proto_goTypes = []interface{}{
	(*Color)(nil),                     // 0: game.Color
	(*Tuning)(nil),                    // 1: game.Tuning
	(*descriptorpb.FieldOptions)(nil), // 2: google.protobuf.FieldOptions
}
var file_tuning_proto_depIdxs = []int32{
	0, // 0: game.Tuning.bg:type_name -> game.Color
	2, // 1: game.section:extendee -> google.protobuf.FieldOptions
	2, // 2: game.min:extendee -> google.protobuf.FieldOptions
	2, // 3: game.max:extendee -> google.protobuf.FieldOptions
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	1, // [1:4] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tuning_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
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
