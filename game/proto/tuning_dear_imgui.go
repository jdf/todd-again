// Code generated by protoc-gen-imgui. DO NOT EDIT.
package proto

import "github.com/inkyblackness/imgui-go/v4"

var (
	tmpFloat32 float32
)

func RenderColor(p *Color) {

	tmpFloat32 = p.GetRed()

	imgui.SliderFloat("Red", &tmpFloat32, .5*Default_Color_Red, 2*Default_Color_Red)
	if p.Red == nil {
		var f float32
		p.Red = &f
	}
	*p.Red = tmpFloat32

	tmpFloat32 = p.GetGreen()

	imgui.SliderFloat("Green", &tmpFloat32, .5*Default_Color_Green, 2*Default_Color_Green)
	if p.Green == nil {
		var f float32
		p.Green = &f
	}
	*p.Green = tmpFloat32

	tmpFloat32 = p.GetBlue()

	imgui.SliderFloat("Blue", &tmpFloat32, .5*Default_Color_Blue, 2*Default_Color_Blue)
	if p.Blue == nil {
		var f float32
		p.Blue = &f
	}
	*p.Blue = tmpFloat32

	tmpFloat32 = p.GetAlpha()

	imgui.SliderFloat("Alpha", &tmpFloat32, .5*Default_Color_Alpha, 2*Default_Color_Alpha)
	if p.Alpha == nil {
		var f float32
		p.Alpha = &f
	}
	*p.Alpha = tmpFloat32
}
func RenderTuning(p *Tuning) {

	tmpFloat32 = p.GetToddSideLength()

	imgui.SliderFloat("ToddSideLength", &tmpFloat32, .5*Default_Tuning_ToddSideLength, 2*Default_Tuning_ToddSideLength)
	if p.ToddSideLength == nil {
		var f float32
		p.ToddSideLength = &f
	}
	*p.ToddSideLength = tmpFloat32

	tmpFloat32 = p.GetFriction()

	imgui.SliderFloat("Friction", &tmpFloat32, .5*Default_Tuning_Friction, 2*Default_Tuning_Friction)
	if p.Friction == nil {
		var f float32
		p.Friction = &f
	}
	*p.Friction = tmpFloat32

	tmpFloat32 = p.GetBearingFriction()

	imgui.SliderFloat("BearingFriction", &tmpFloat32, .5*Default_Tuning_BearingFriction, 2*Default_Tuning_BearingFriction)
	if p.BearingFriction == nil {
		var f float32
		p.BearingFriction = &f
	}
	*p.BearingFriction = tmpFloat32

	tmpFloat32 = p.GetGravity()

	imgui.SliderFloat("Gravity", &tmpFloat32, .5*Default_Tuning_Gravity, 2*Default_Tuning_Gravity)
	if p.Gravity == nil {
		var f float32
		p.Gravity = &f
	}
	*p.Gravity = tmpFloat32

	tmpFloat32 = p.GetMaxVelocity()

	imgui.SliderFloat("MaxVelocity", &tmpFloat32, .5*Default_Tuning_MaxVelocity, 2*Default_Tuning_MaxVelocity)
	if p.MaxVelocity == nil {
		var f float32
		p.MaxVelocity = &f
	}
	*p.MaxVelocity = tmpFloat32

	tmpFloat32 = p.GetAcceleration()

	imgui.SliderFloat("Acceleration", &tmpFloat32, .5*Default_Tuning_Acceleration, 2*Default_Tuning_Acceleration)
	if p.Acceleration == nil {
		var f float32
		p.Acceleration = &f
	}
	*p.Acceleration = tmpFloat32

	tmpFloat32 = p.GetAirBending()

	imgui.SliderFloat("AirBending", &tmpFloat32, .5*Default_Tuning_AirBending, 2*Default_Tuning_AirBending)
	if p.AirBending == nil {
		var f float32
		p.AirBending = &f
	}
	*p.AirBending = tmpFloat32

	tmpFloat32 = p.GetBearingAcceleration()

	imgui.SliderFloat("BearingAcceleration", &tmpFloat32, .5*Default_Tuning_BearingAcceleration, 2*Default_Tuning_BearingAcceleration)
	if p.BearingAcceleration == nil {
		var f float32
		p.BearingAcceleration = &f
	}
	*p.BearingAcceleration = tmpFloat32

	tmpFloat32 = p.GetJumpImpulse()

	imgui.SliderFloat("JumpImpulse", &tmpFloat32, .5*Default_Tuning_JumpImpulse, 2*Default_Tuning_JumpImpulse)
	if p.JumpImpulse == nil {
		var f float32
		p.JumpImpulse = &f
	}
	*p.JumpImpulse = tmpFloat32

	tmpFloat32 = p.GetMaxSquishVelocity()

	imgui.SliderFloat("MaxSquishVelocity", &tmpFloat32, .5*Default_Tuning_MaxSquishVelocity, 2*Default_Tuning_MaxSquishVelocity)
	if p.MaxSquishVelocity == nil {
		var f float32
		p.MaxSquishVelocity = &f
	}
	*p.MaxSquishVelocity = tmpFloat32

	tmpFloat32 = p.GetJumpTerminalVelocity()

	imgui.SliderFloat("JumpTerminalVelocity", &tmpFloat32, .5*Default_Tuning_JumpTerminalVelocity, 2*Default_Tuning_JumpTerminalVelocity)
	if p.JumpTerminalVelocity == nil {
		var f float32
		p.JumpTerminalVelocity = &f
	}
	*p.JumpTerminalVelocity = tmpFloat32

	tmpFloat32 = p.GetTerminalVelocity()

	imgui.SliderFloat("TerminalVelocity", &tmpFloat32, .5*Default_Tuning_TerminalVelocity, 2*Default_Tuning_TerminalVelocity)
	if p.TerminalVelocity == nil {
		var f float32
		p.TerminalVelocity = &f
	}
	*p.TerminalVelocity = tmpFloat32

	tmpFloat32 = p.GetBlinkOdds()

	imgui.SliderFloat("BlinkOdds", &tmpFloat32, .5*Default_Tuning_BlinkOdds, 2*Default_Tuning_BlinkOdds)
	if p.BlinkOdds == nil {
		var f float32
		p.BlinkOdds = &f
	}
	*p.BlinkOdds = tmpFloat32

	tmpFloat32 = p.GetBlinkCycleSeconds()

	imgui.SliderFloat("BlinkCycleSeconds", &tmpFloat32, .5*Default_Tuning_BlinkCycleSeconds, 2*Default_Tuning_BlinkCycleSeconds)
	if p.BlinkCycleSeconds == nil {
		var f float32
		p.BlinkCycleSeconds = &f
	}
	*p.BlinkCycleSeconds = tmpFloat32

	tmpFloat32 = p.GetEyeCenteringDurationSeconds()

	imgui.SliderFloat("EyeCenteringDurationSeconds", &tmpFloat32, .5*Default_Tuning_EyeCenteringDurationSeconds, 2*Default_Tuning_EyeCenteringDurationSeconds)
	if p.EyeCenteringDurationSeconds == nil {
		var f float32
		p.EyeCenteringDurationSeconds = &f
	}
	*p.EyeCenteringDurationSeconds = tmpFloat32

	tmpFloat32 = p.GetJumpStateGravityFactor()

	imgui.SliderFloat("JumpStateGravityFactor", &tmpFloat32, .5*Default_Tuning_JumpStateGravityFactor, 2*Default_Tuning_JumpStateGravityFactor)
	if p.JumpStateGravityFactor == nil {
		var f float32
		p.JumpStateGravityFactor = &f
	}
	*p.JumpStateGravityFactor = tmpFloat32

	tmpFloat32 = p.GetCameraTiltSeconds()

	imgui.SliderFloat("CameraTiltSeconds", &tmpFloat32, .5*Default_Tuning_CameraTiltSeconds, 2*Default_Tuning_CameraTiltSeconds)
	if p.CameraTiltSeconds == nil {
		var f float32
		p.CameraTiltSeconds = &f
	}
	*p.CameraTiltSeconds = tmpFloat32

	tmpFloat32 = p.GetJumpRequestSlopSeconds()

	imgui.SliderFloat("JumpRequestSlopSeconds", &tmpFloat32, .5*Default_Tuning_JumpRequestSlopSeconds, 2*Default_Tuning_JumpRequestSlopSeconds)
	if p.JumpRequestSlopSeconds == nil {
		var f float32
		p.JumpRequestSlopSeconds = &f
	}
	*p.JumpRequestSlopSeconds = tmpFloat32

	tmpFloat32 = p.GetGroundingSlopSeconds()

	imgui.SliderFloat("GroundingSlopSeconds", &tmpFloat32, .5*Default_Tuning_GroundingSlopSeconds, 2*Default_Tuning_GroundingSlopSeconds)
	if p.GroundingSlopSeconds == nil {
		var f float32
		p.GroundingSlopSeconds = &f
	}
	*p.GroundingSlopSeconds = tmpFloat32

	tmpFloat32 = p.GetAirFriction()

	imgui.SliderFloat("AirFriction", &tmpFloat32, .5*Default_Tuning_AirFriction, 2*Default_Tuning_AirFriction)
	if p.AirFriction == nil {
		var f float32
		p.AirFriction = &f
	}
	*p.AirFriction = tmpFloat32
}
