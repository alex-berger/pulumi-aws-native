// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ivs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Channel latency mode.
type ChannelLatencyMode string

const (
	ChannelLatencyModeNormal = ChannelLatencyMode("NORMAL")
	ChannelLatencyModeLow    = ChannelLatencyMode("LOW")
)

func (ChannelLatencyMode) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelLatencyMode)(nil)).Elem()
}

func (e ChannelLatencyMode) ToChannelLatencyModeOutput() ChannelLatencyModeOutput {
	return pulumi.ToOutput(e).(ChannelLatencyModeOutput)
}

func (e ChannelLatencyMode) ToChannelLatencyModeOutputWithContext(ctx context.Context) ChannelLatencyModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ChannelLatencyModeOutput)
}

func (e ChannelLatencyMode) ToChannelLatencyModePtrOutput() ChannelLatencyModePtrOutput {
	return e.ToChannelLatencyModePtrOutputWithContext(context.Background())
}

func (e ChannelLatencyMode) ToChannelLatencyModePtrOutputWithContext(ctx context.Context) ChannelLatencyModePtrOutput {
	return ChannelLatencyMode(e).ToChannelLatencyModeOutputWithContext(ctx).ToChannelLatencyModePtrOutputWithContext(ctx)
}

func (e ChannelLatencyMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ChannelLatencyMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ChannelLatencyMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ChannelLatencyMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ChannelLatencyModeOutput struct{ *pulumi.OutputState }

func (ChannelLatencyModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelLatencyMode)(nil)).Elem()
}

func (o ChannelLatencyModeOutput) ToChannelLatencyModeOutput() ChannelLatencyModeOutput {
	return o
}

func (o ChannelLatencyModeOutput) ToChannelLatencyModeOutputWithContext(ctx context.Context) ChannelLatencyModeOutput {
	return o
}

func (o ChannelLatencyModeOutput) ToChannelLatencyModePtrOutput() ChannelLatencyModePtrOutput {
	return o.ToChannelLatencyModePtrOutputWithContext(context.Background())
}

func (o ChannelLatencyModeOutput) ToChannelLatencyModePtrOutputWithContext(ctx context.Context) ChannelLatencyModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ChannelLatencyMode) *ChannelLatencyMode {
		return &v
	}).(ChannelLatencyModePtrOutput)
}

func (o ChannelLatencyModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ChannelLatencyModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ChannelLatencyMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ChannelLatencyModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ChannelLatencyModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ChannelLatencyMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ChannelLatencyModePtrOutput struct{ *pulumi.OutputState }

func (ChannelLatencyModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ChannelLatencyMode)(nil)).Elem()
}

func (o ChannelLatencyModePtrOutput) ToChannelLatencyModePtrOutput() ChannelLatencyModePtrOutput {
	return o
}

func (o ChannelLatencyModePtrOutput) ToChannelLatencyModePtrOutputWithContext(ctx context.Context) ChannelLatencyModePtrOutput {
	return o
}

func (o ChannelLatencyModePtrOutput) Elem() ChannelLatencyModeOutput {
	return o.ApplyT(func(v *ChannelLatencyMode) ChannelLatencyMode {
		if v != nil {
			return *v
		}
		var ret ChannelLatencyMode
		return ret
	}).(ChannelLatencyModeOutput)
}

func (o ChannelLatencyModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ChannelLatencyModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ChannelLatencyMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ChannelLatencyModeInput is an input type that accepts ChannelLatencyModeArgs and ChannelLatencyModeOutput values.
// You can construct a concrete instance of `ChannelLatencyModeInput` via:
//
//          ChannelLatencyModeArgs{...}
type ChannelLatencyModeInput interface {
	pulumi.Input

	ToChannelLatencyModeOutput() ChannelLatencyModeOutput
	ToChannelLatencyModeOutputWithContext(context.Context) ChannelLatencyModeOutput
}

var channelLatencyModePtrType = reflect.TypeOf((**ChannelLatencyMode)(nil)).Elem()

type ChannelLatencyModePtrInput interface {
	pulumi.Input

	ToChannelLatencyModePtrOutput() ChannelLatencyModePtrOutput
	ToChannelLatencyModePtrOutputWithContext(context.Context) ChannelLatencyModePtrOutput
}

type channelLatencyModePtr string

func ChannelLatencyModePtr(v string) ChannelLatencyModePtrInput {
	return (*channelLatencyModePtr)(&v)
}

func (*channelLatencyModePtr) ElementType() reflect.Type {
	return channelLatencyModePtrType
}

func (in *channelLatencyModePtr) ToChannelLatencyModePtrOutput() ChannelLatencyModePtrOutput {
	return pulumi.ToOutput(in).(ChannelLatencyModePtrOutput)
}

func (in *channelLatencyModePtr) ToChannelLatencyModePtrOutputWithContext(ctx context.Context) ChannelLatencyModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ChannelLatencyModePtrOutput)
}

// Channel type, which determines the allowable resolution and bitrate. If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately.
type ChannelType string

const (
	ChannelTypeStandard = ChannelType("STANDARD")
	ChannelTypeBasic    = ChannelType("BASIC")
)

func (ChannelType) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelType)(nil)).Elem()
}

func (e ChannelType) ToChannelTypeOutput() ChannelTypeOutput {
	return pulumi.ToOutput(e).(ChannelTypeOutput)
}

func (e ChannelType) ToChannelTypeOutputWithContext(ctx context.Context) ChannelTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ChannelTypeOutput)
}

func (e ChannelType) ToChannelTypePtrOutput() ChannelTypePtrOutput {
	return e.ToChannelTypePtrOutputWithContext(context.Background())
}

func (e ChannelType) ToChannelTypePtrOutputWithContext(ctx context.Context) ChannelTypePtrOutput {
	return ChannelType(e).ToChannelTypeOutputWithContext(ctx).ToChannelTypePtrOutputWithContext(ctx)
}

func (e ChannelType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ChannelType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ChannelType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ChannelType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ChannelTypeOutput struct{ *pulumi.OutputState }

func (ChannelTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelType)(nil)).Elem()
}

func (o ChannelTypeOutput) ToChannelTypeOutput() ChannelTypeOutput {
	return o
}

func (o ChannelTypeOutput) ToChannelTypeOutputWithContext(ctx context.Context) ChannelTypeOutput {
	return o
}

func (o ChannelTypeOutput) ToChannelTypePtrOutput() ChannelTypePtrOutput {
	return o.ToChannelTypePtrOutputWithContext(context.Background())
}

func (o ChannelTypeOutput) ToChannelTypePtrOutputWithContext(ctx context.Context) ChannelTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ChannelType) *ChannelType {
		return &v
	}).(ChannelTypePtrOutput)
}

func (o ChannelTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ChannelTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ChannelType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ChannelTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ChannelTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ChannelType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ChannelTypePtrOutput struct{ *pulumi.OutputState }

func (ChannelTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ChannelType)(nil)).Elem()
}

func (o ChannelTypePtrOutput) ToChannelTypePtrOutput() ChannelTypePtrOutput {
	return o
}

func (o ChannelTypePtrOutput) ToChannelTypePtrOutputWithContext(ctx context.Context) ChannelTypePtrOutput {
	return o
}

func (o ChannelTypePtrOutput) Elem() ChannelTypeOutput {
	return o.ApplyT(func(v *ChannelType) ChannelType {
		if v != nil {
			return *v
		}
		var ret ChannelType
		return ret
	}).(ChannelTypeOutput)
}

func (o ChannelTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ChannelTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ChannelType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ChannelTypeInput is an input type that accepts ChannelTypeArgs and ChannelTypeOutput values.
// You can construct a concrete instance of `ChannelTypeInput` via:
//
//          ChannelTypeArgs{...}
type ChannelTypeInput interface {
	pulumi.Input

	ToChannelTypeOutput() ChannelTypeOutput
	ToChannelTypeOutputWithContext(context.Context) ChannelTypeOutput
}

var channelTypePtrType = reflect.TypeOf((**ChannelType)(nil)).Elem()

type ChannelTypePtrInput interface {
	pulumi.Input

	ToChannelTypePtrOutput() ChannelTypePtrOutput
	ToChannelTypePtrOutputWithContext(context.Context) ChannelTypePtrOutput
}

type channelTypePtr string

func ChannelTypePtr(v string) ChannelTypePtrInput {
	return (*channelTypePtr)(&v)
}

func (*channelTypePtr) ElementType() reflect.Type {
	return channelTypePtrType
}

func (in *channelTypePtr) ToChannelTypePtrOutput() ChannelTypePtrOutput {
	return pulumi.ToOutput(in).(ChannelTypePtrOutput)
}

func (in *channelTypePtr) ToChannelTypePtrOutputWithContext(ctx context.Context) ChannelTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ChannelTypePtrOutput)
}

// Recording Configuration State.
type RecordingConfigurationStateEnum string

const (
	RecordingConfigurationStateEnumCreating     = RecordingConfigurationStateEnum("CREATING")
	RecordingConfigurationStateEnumCreateFailed = RecordingConfigurationStateEnum("CREATE_FAILED")
	RecordingConfigurationStateEnumActive       = RecordingConfigurationStateEnum("ACTIVE")
)

type RecordingConfigurationStateEnumOutput struct{ *pulumi.OutputState }

func (RecordingConfigurationStateEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecordingConfigurationStateEnum)(nil)).Elem()
}

func (o RecordingConfigurationStateEnumOutput) ToRecordingConfigurationStateEnumOutput() RecordingConfigurationStateEnumOutput {
	return o
}

func (o RecordingConfigurationStateEnumOutput) ToRecordingConfigurationStateEnumOutputWithContext(ctx context.Context) RecordingConfigurationStateEnumOutput {
	return o
}

func (o RecordingConfigurationStateEnumOutput) ToRecordingConfigurationStateEnumPtrOutput() RecordingConfigurationStateEnumPtrOutput {
	return o.ToRecordingConfigurationStateEnumPtrOutputWithContext(context.Background())
}

func (o RecordingConfigurationStateEnumOutput) ToRecordingConfigurationStateEnumPtrOutputWithContext(ctx context.Context) RecordingConfigurationStateEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RecordingConfigurationStateEnum) *RecordingConfigurationStateEnum {
		return &v
	}).(RecordingConfigurationStateEnumPtrOutput)
}

func (o RecordingConfigurationStateEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RecordingConfigurationStateEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecordingConfigurationStateEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RecordingConfigurationStateEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecordingConfigurationStateEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecordingConfigurationStateEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RecordingConfigurationStateEnumPtrOutput struct{ *pulumi.OutputState }

func (RecordingConfigurationStateEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecordingConfigurationStateEnum)(nil)).Elem()
}

func (o RecordingConfigurationStateEnumPtrOutput) ToRecordingConfigurationStateEnumPtrOutput() RecordingConfigurationStateEnumPtrOutput {
	return o
}

func (o RecordingConfigurationStateEnumPtrOutput) ToRecordingConfigurationStateEnumPtrOutputWithContext(ctx context.Context) RecordingConfigurationStateEnumPtrOutput {
	return o
}

func (o RecordingConfigurationStateEnumPtrOutput) Elem() RecordingConfigurationStateEnumOutput {
	return o.ApplyT(func(v *RecordingConfigurationStateEnum) RecordingConfigurationStateEnum {
		if v != nil {
			return *v
		}
		var ret RecordingConfigurationStateEnum
		return ret
	}).(RecordingConfigurationStateEnumOutput)
}

func (o RecordingConfigurationStateEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecordingConfigurationStateEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RecordingConfigurationStateEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// Thumbnail Recording Mode, which determines whether thumbnails are recorded at an interval or are disabled.
type RecordingConfigurationThumbnailConfigurationRecordingMode string

const (
	RecordingConfigurationThumbnailConfigurationRecordingModeInterval = RecordingConfigurationThumbnailConfigurationRecordingMode("INTERVAL")
	RecordingConfigurationThumbnailConfigurationRecordingModeDisabled = RecordingConfigurationThumbnailConfigurationRecordingMode("DISABLED")
)

func (RecordingConfigurationThumbnailConfigurationRecordingMode) ElementType() reflect.Type {
	return reflect.TypeOf((*RecordingConfigurationThumbnailConfigurationRecordingMode)(nil)).Elem()
}

func (e RecordingConfigurationThumbnailConfigurationRecordingMode) ToRecordingConfigurationThumbnailConfigurationRecordingModeOutput() RecordingConfigurationThumbnailConfigurationRecordingModeOutput {
	return pulumi.ToOutput(e).(RecordingConfigurationThumbnailConfigurationRecordingModeOutput)
}

func (e RecordingConfigurationThumbnailConfigurationRecordingMode) ToRecordingConfigurationThumbnailConfigurationRecordingModeOutputWithContext(ctx context.Context) RecordingConfigurationThumbnailConfigurationRecordingModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RecordingConfigurationThumbnailConfigurationRecordingModeOutput)
}

func (e RecordingConfigurationThumbnailConfigurationRecordingMode) ToRecordingConfigurationThumbnailConfigurationRecordingModePtrOutput() RecordingConfigurationThumbnailConfigurationRecordingModePtrOutput {
	return e.ToRecordingConfigurationThumbnailConfigurationRecordingModePtrOutputWithContext(context.Background())
}

func (e RecordingConfigurationThumbnailConfigurationRecordingMode) ToRecordingConfigurationThumbnailConfigurationRecordingModePtrOutputWithContext(ctx context.Context) RecordingConfigurationThumbnailConfigurationRecordingModePtrOutput {
	return RecordingConfigurationThumbnailConfigurationRecordingMode(e).ToRecordingConfigurationThumbnailConfigurationRecordingModeOutputWithContext(ctx).ToRecordingConfigurationThumbnailConfigurationRecordingModePtrOutputWithContext(ctx)
}

func (e RecordingConfigurationThumbnailConfigurationRecordingMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecordingConfigurationThumbnailConfigurationRecordingMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecordingConfigurationThumbnailConfigurationRecordingMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RecordingConfigurationThumbnailConfigurationRecordingMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RecordingConfigurationThumbnailConfigurationRecordingModeOutput struct{ *pulumi.OutputState }

func (RecordingConfigurationThumbnailConfigurationRecordingModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecordingConfigurationThumbnailConfigurationRecordingMode)(nil)).Elem()
}

func (o RecordingConfigurationThumbnailConfigurationRecordingModeOutput) ToRecordingConfigurationThumbnailConfigurationRecordingModeOutput() RecordingConfigurationThumbnailConfigurationRecordingModeOutput {
	return o
}

func (o RecordingConfigurationThumbnailConfigurationRecordingModeOutput) ToRecordingConfigurationThumbnailConfigurationRecordingModeOutputWithContext(ctx context.Context) RecordingConfigurationThumbnailConfigurationRecordingModeOutput {
	return o
}

func (o RecordingConfigurationThumbnailConfigurationRecordingModeOutput) ToRecordingConfigurationThumbnailConfigurationRecordingModePtrOutput() RecordingConfigurationThumbnailConfigurationRecordingModePtrOutput {
	return o.ToRecordingConfigurationThumbnailConfigurationRecordingModePtrOutputWithContext(context.Background())
}

func (o RecordingConfigurationThumbnailConfigurationRecordingModeOutput) ToRecordingConfigurationThumbnailConfigurationRecordingModePtrOutputWithContext(ctx context.Context) RecordingConfigurationThumbnailConfigurationRecordingModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RecordingConfigurationThumbnailConfigurationRecordingMode) *RecordingConfigurationThumbnailConfigurationRecordingMode {
		return &v
	}).(RecordingConfigurationThumbnailConfigurationRecordingModePtrOutput)
}

func (o RecordingConfigurationThumbnailConfigurationRecordingModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RecordingConfigurationThumbnailConfigurationRecordingModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecordingConfigurationThumbnailConfigurationRecordingMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RecordingConfigurationThumbnailConfigurationRecordingModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecordingConfigurationThumbnailConfigurationRecordingModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecordingConfigurationThumbnailConfigurationRecordingMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RecordingConfigurationThumbnailConfigurationRecordingModePtrOutput struct{ *pulumi.OutputState }

func (RecordingConfigurationThumbnailConfigurationRecordingModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecordingConfigurationThumbnailConfigurationRecordingMode)(nil)).Elem()
}

func (o RecordingConfigurationThumbnailConfigurationRecordingModePtrOutput) ToRecordingConfigurationThumbnailConfigurationRecordingModePtrOutput() RecordingConfigurationThumbnailConfigurationRecordingModePtrOutput {
	return o
}

func (o RecordingConfigurationThumbnailConfigurationRecordingModePtrOutput) ToRecordingConfigurationThumbnailConfigurationRecordingModePtrOutputWithContext(ctx context.Context) RecordingConfigurationThumbnailConfigurationRecordingModePtrOutput {
	return o
}

func (o RecordingConfigurationThumbnailConfigurationRecordingModePtrOutput) Elem() RecordingConfigurationThumbnailConfigurationRecordingModeOutput {
	return o.ApplyT(func(v *RecordingConfigurationThumbnailConfigurationRecordingMode) RecordingConfigurationThumbnailConfigurationRecordingMode {
		if v != nil {
			return *v
		}
		var ret RecordingConfigurationThumbnailConfigurationRecordingMode
		return ret
	}).(RecordingConfigurationThumbnailConfigurationRecordingModeOutput)
}

func (o RecordingConfigurationThumbnailConfigurationRecordingModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecordingConfigurationThumbnailConfigurationRecordingModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RecordingConfigurationThumbnailConfigurationRecordingMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// RecordingConfigurationThumbnailConfigurationRecordingModeInput is an input type that accepts RecordingConfigurationThumbnailConfigurationRecordingModeArgs and RecordingConfigurationThumbnailConfigurationRecordingModeOutput values.
// You can construct a concrete instance of `RecordingConfigurationThumbnailConfigurationRecordingModeInput` via:
//
//          RecordingConfigurationThumbnailConfigurationRecordingModeArgs{...}
type RecordingConfigurationThumbnailConfigurationRecordingModeInput interface {
	pulumi.Input

	ToRecordingConfigurationThumbnailConfigurationRecordingModeOutput() RecordingConfigurationThumbnailConfigurationRecordingModeOutput
	ToRecordingConfigurationThumbnailConfigurationRecordingModeOutputWithContext(context.Context) RecordingConfigurationThumbnailConfigurationRecordingModeOutput
}

var recordingConfigurationThumbnailConfigurationRecordingModePtrType = reflect.TypeOf((**RecordingConfigurationThumbnailConfigurationRecordingMode)(nil)).Elem()

type RecordingConfigurationThumbnailConfigurationRecordingModePtrInput interface {
	pulumi.Input

	ToRecordingConfigurationThumbnailConfigurationRecordingModePtrOutput() RecordingConfigurationThumbnailConfigurationRecordingModePtrOutput
	ToRecordingConfigurationThumbnailConfigurationRecordingModePtrOutputWithContext(context.Context) RecordingConfigurationThumbnailConfigurationRecordingModePtrOutput
}

type recordingConfigurationThumbnailConfigurationRecordingModePtr string

func RecordingConfigurationThumbnailConfigurationRecordingModePtr(v string) RecordingConfigurationThumbnailConfigurationRecordingModePtrInput {
	return (*recordingConfigurationThumbnailConfigurationRecordingModePtr)(&v)
}

func (*recordingConfigurationThumbnailConfigurationRecordingModePtr) ElementType() reflect.Type {
	return recordingConfigurationThumbnailConfigurationRecordingModePtrType
}

func (in *recordingConfigurationThumbnailConfigurationRecordingModePtr) ToRecordingConfigurationThumbnailConfigurationRecordingModePtrOutput() RecordingConfigurationThumbnailConfigurationRecordingModePtrOutput {
	return pulumi.ToOutput(in).(RecordingConfigurationThumbnailConfigurationRecordingModePtrOutput)
}

func (in *recordingConfigurationThumbnailConfigurationRecordingModePtr) ToRecordingConfigurationThumbnailConfigurationRecordingModePtrOutputWithContext(ctx context.Context) RecordingConfigurationThumbnailConfigurationRecordingModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RecordingConfigurationThumbnailConfigurationRecordingModePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ChannelLatencyModeInput)(nil)).Elem(), ChannelLatencyMode("NORMAL"))
	pulumi.RegisterInputType(reflect.TypeOf((*ChannelLatencyModePtrInput)(nil)).Elem(), ChannelLatencyMode("NORMAL"))
	pulumi.RegisterInputType(reflect.TypeOf((*ChannelTypeInput)(nil)).Elem(), ChannelType("STANDARD"))
	pulumi.RegisterInputType(reflect.TypeOf((*ChannelTypePtrInput)(nil)).Elem(), ChannelType("STANDARD"))
	pulumi.RegisterInputType(reflect.TypeOf((*RecordingConfigurationThumbnailConfigurationRecordingModeInput)(nil)).Elem(), RecordingConfigurationThumbnailConfigurationRecordingMode("INTERVAL"))
	pulumi.RegisterInputType(reflect.TypeOf((*RecordingConfigurationThumbnailConfigurationRecordingModePtrInput)(nil)).Elem(), RecordingConfigurationThumbnailConfigurationRecordingMode("INTERVAL"))
	pulumi.RegisterOutputType(ChannelLatencyModeOutput{})
	pulumi.RegisterOutputType(ChannelLatencyModePtrOutput{})
	pulumi.RegisterOutputType(ChannelTypeOutput{})
	pulumi.RegisterOutputType(ChannelTypePtrOutput{})
	pulumi.RegisterOutputType(RecordingConfigurationStateEnumOutput{})
	pulumi.RegisterOutputType(RecordingConfigurationStateEnumPtrOutput{})
	pulumi.RegisterOutputType(RecordingConfigurationThumbnailConfigurationRecordingModeOutput{})
	pulumi.RegisterOutputType(RecordingConfigurationThumbnailConfigurationRecordingModePtrOutput{})
}
