// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package stepfunctions

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StateMachineLoggingConfigurationLevel string

const (
	StateMachineLoggingConfigurationLevelAll   = StateMachineLoggingConfigurationLevel("ALL")
	StateMachineLoggingConfigurationLevelError = StateMachineLoggingConfigurationLevel("ERROR")
	StateMachineLoggingConfigurationLevelFatal = StateMachineLoggingConfigurationLevel("FATAL")
	StateMachineLoggingConfigurationLevelOff   = StateMachineLoggingConfigurationLevel("OFF")
)

func (StateMachineLoggingConfigurationLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*StateMachineLoggingConfigurationLevel)(nil)).Elem()
}

func (e StateMachineLoggingConfigurationLevel) ToStateMachineLoggingConfigurationLevelOutput() StateMachineLoggingConfigurationLevelOutput {
	return pulumi.ToOutput(e).(StateMachineLoggingConfigurationLevelOutput)
}

func (e StateMachineLoggingConfigurationLevel) ToStateMachineLoggingConfigurationLevelOutputWithContext(ctx context.Context) StateMachineLoggingConfigurationLevelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StateMachineLoggingConfigurationLevelOutput)
}

func (e StateMachineLoggingConfigurationLevel) ToStateMachineLoggingConfigurationLevelPtrOutput() StateMachineLoggingConfigurationLevelPtrOutput {
	return e.ToStateMachineLoggingConfigurationLevelPtrOutputWithContext(context.Background())
}

func (e StateMachineLoggingConfigurationLevel) ToStateMachineLoggingConfigurationLevelPtrOutputWithContext(ctx context.Context) StateMachineLoggingConfigurationLevelPtrOutput {
	return StateMachineLoggingConfigurationLevel(e).ToStateMachineLoggingConfigurationLevelOutputWithContext(ctx).ToStateMachineLoggingConfigurationLevelPtrOutputWithContext(ctx)
}

func (e StateMachineLoggingConfigurationLevel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StateMachineLoggingConfigurationLevel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StateMachineLoggingConfigurationLevel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StateMachineLoggingConfigurationLevel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StateMachineLoggingConfigurationLevelOutput struct{ *pulumi.OutputState }

func (StateMachineLoggingConfigurationLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StateMachineLoggingConfigurationLevel)(nil)).Elem()
}

func (o StateMachineLoggingConfigurationLevelOutput) ToStateMachineLoggingConfigurationLevelOutput() StateMachineLoggingConfigurationLevelOutput {
	return o
}

func (o StateMachineLoggingConfigurationLevelOutput) ToStateMachineLoggingConfigurationLevelOutputWithContext(ctx context.Context) StateMachineLoggingConfigurationLevelOutput {
	return o
}

func (o StateMachineLoggingConfigurationLevelOutput) ToStateMachineLoggingConfigurationLevelPtrOutput() StateMachineLoggingConfigurationLevelPtrOutput {
	return o.ToStateMachineLoggingConfigurationLevelPtrOutputWithContext(context.Background())
}

func (o StateMachineLoggingConfigurationLevelOutput) ToStateMachineLoggingConfigurationLevelPtrOutputWithContext(ctx context.Context) StateMachineLoggingConfigurationLevelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StateMachineLoggingConfigurationLevel) *StateMachineLoggingConfigurationLevel {
		return &v
	}).(StateMachineLoggingConfigurationLevelPtrOutput)
}

func (o StateMachineLoggingConfigurationLevelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StateMachineLoggingConfigurationLevelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StateMachineLoggingConfigurationLevel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StateMachineLoggingConfigurationLevelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StateMachineLoggingConfigurationLevelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StateMachineLoggingConfigurationLevel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StateMachineLoggingConfigurationLevelPtrOutput struct{ *pulumi.OutputState }

func (StateMachineLoggingConfigurationLevelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StateMachineLoggingConfigurationLevel)(nil)).Elem()
}

func (o StateMachineLoggingConfigurationLevelPtrOutput) ToStateMachineLoggingConfigurationLevelPtrOutput() StateMachineLoggingConfigurationLevelPtrOutput {
	return o
}

func (o StateMachineLoggingConfigurationLevelPtrOutput) ToStateMachineLoggingConfigurationLevelPtrOutputWithContext(ctx context.Context) StateMachineLoggingConfigurationLevelPtrOutput {
	return o
}

func (o StateMachineLoggingConfigurationLevelPtrOutput) Elem() StateMachineLoggingConfigurationLevelOutput {
	return o.ApplyT(func(v *StateMachineLoggingConfigurationLevel) StateMachineLoggingConfigurationLevel {
		if v != nil {
			return *v
		}
		var ret StateMachineLoggingConfigurationLevel
		return ret
	}).(StateMachineLoggingConfigurationLevelOutput)
}

func (o StateMachineLoggingConfigurationLevelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StateMachineLoggingConfigurationLevelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StateMachineLoggingConfigurationLevel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// StateMachineLoggingConfigurationLevelInput is an input type that accepts StateMachineLoggingConfigurationLevelArgs and StateMachineLoggingConfigurationLevelOutput values.
// You can construct a concrete instance of `StateMachineLoggingConfigurationLevelInput` via:
//
//          StateMachineLoggingConfigurationLevelArgs{...}
type StateMachineLoggingConfigurationLevelInput interface {
	pulumi.Input

	ToStateMachineLoggingConfigurationLevelOutput() StateMachineLoggingConfigurationLevelOutput
	ToStateMachineLoggingConfigurationLevelOutputWithContext(context.Context) StateMachineLoggingConfigurationLevelOutput
}

var stateMachineLoggingConfigurationLevelPtrType = reflect.TypeOf((**StateMachineLoggingConfigurationLevel)(nil)).Elem()

type StateMachineLoggingConfigurationLevelPtrInput interface {
	pulumi.Input

	ToStateMachineLoggingConfigurationLevelPtrOutput() StateMachineLoggingConfigurationLevelPtrOutput
	ToStateMachineLoggingConfigurationLevelPtrOutputWithContext(context.Context) StateMachineLoggingConfigurationLevelPtrOutput
}

type stateMachineLoggingConfigurationLevelPtr string

func StateMachineLoggingConfigurationLevelPtr(v string) StateMachineLoggingConfigurationLevelPtrInput {
	return (*stateMachineLoggingConfigurationLevelPtr)(&v)
}

func (*stateMachineLoggingConfigurationLevelPtr) ElementType() reflect.Type {
	return stateMachineLoggingConfigurationLevelPtrType
}

func (in *stateMachineLoggingConfigurationLevelPtr) ToStateMachineLoggingConfigurationLevelPtrOutput() StateMachineLoggingConfigurationLevelPtrOutput {
	return pulumi.ToOutput(in).(StateMachineLoggingConfigurationLevelPtrOutput)
}

func (in *stateMachineLoggingConfigurationLevelPtr) ToStateMachineLoggingConfigurationLevelPtrOutputWithContext(ctx context.Context) StateMachineLoggingConfigurationLevelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StateMachineLoggingConfigurationLevelPtrOutput)
}

type StateMachineType string

const (
	StateMachineTypeStandard = StateMachineType("STANDARD")
	StateMachineTypeExpress  = StateMachineType("EXPRESS")
)

func (StateMachineType) ElementType() reflect.Type {
	return reflect.TypeOf((*StateMachineType)(nil)).Elem()
}

func (e StateMachineType) ToStateMachineTypeOutput() StateMachineTypeOutput {
	return pulumi.ToOutput(e).(StateMachineTypeOutput)
}

func (e StateMachineType) ToStateMachineTypeOutputWithContext(ctx context.Context) StateMachineTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StateMachineTypeOutput)
}

func (e StateMachineType) ToStateMachineTypePtrOutput() StateMachineTypePtrOutput {
	return e.ToStateMachineTypePtrOutputWithContext(context.Background())
}

func (e StateMachineType) ToStateMachineTypePtrOutputWithContext(ctx context.Context) StateMachineTypePtrOutput {
	return StateMachineType(e).ToStateMachineTypeOutputWithContext(ctx).ToStateMachineTypePtrOutputWithContext(ctx)
}

func (e StateMachineType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StateMachineType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StateMachineType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StateMachineType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StateMachineTypeOutput struct{ *pulumi.OutputState }

func (StateMachineTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StateMachineType)(nil)).Elem()
}

func (o StateMachineTypeOutput) ToStateMachineTypeOutput() StateMachineTypeOutput {
	return o
}

func (o StateMachineTypeOutput) ToStateMachineTypeOutputWithContext(ctx context.Context) StateMachineTypeOutput {
	return o
}

func (o StateMachineTypeOutput) ToStateMachineTypePtrOutput() StateMachineTypePtrOutput {
	return o.ToStateMachineTypePtrOutputWithContext(context.Background())
}

func (o StateMachineTypeOutput) ToStateMachineTypePtrOutputWithContext(ctx context.Context) StateMachineTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StateMachineType) *StateMachineType {
		return &v
	}).(StateMachineTypePtrOutput)
}

func (o StateMachineTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StateMachineTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StateMachineType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StateMachineTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StateMachineTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StateMachineType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StateMachineTypePtrOutput struct{ *pulumi.OutputState }

func (StateMachineTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StateMachineType)(nil)).Elem()
}

func (o StateMachineTypePtrOutput) ToStateMachineTypePtrOutput() StateMachineTypePtrOutput {
	return o
}

func (o StateMachineTypePtrOutput) ToStateMachineTypePtrOutputWithContext(ctx context.Context) StateMachineTypePtrOutput {
	return o
}

func (o StateMachineTypePtrOutput) Elem() StateMachineTypeOutput {
	return o.ApplyT(func(v *StateMachineType) StateMachineType {
		if v != nil {
			return *v
		}
		var ret StateMachineType
		return ret
	}).(StateMachineTypeOutput)
}

func (o StateMachineTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StateMachineTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StateMachineType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// StateMachineTypeInput is an input type that accepts StateMachineTypeArgs and StateMachineTypeOutput values.
// You can construct a concrete instance of `StateMachineTypeInput` via:
//
//          StateMachineTypeArgs{...}
type StateMachineTypeInput interface {
	pulumi.Input

	ToStateMachineTypeOutput() StateMachineTypeOutput
	ToStateMachineTypeOutputWithContext(context.Context) StateMachineTypeOutput
}

var stateMachineTypePtrType = reflect.TypeOf((**StateMachineType)(nil)).Elem()

type StateMachineTypePtrInput interface {
	pulumi.Input

	ToStateMachineTypePtrOutput() StateMachineTypePtrOutput
	ToStateMachineTypePtrOutputWithContext(context.Context) StateMachineTypePtrOutput
}

type stateMachineTypePtr string

func StateMachineTypePtr(v string) StateMachineTypePtrInput {
	return (*stateMachineTypePtr)(&v)
}

func (*stateMachineTypePtr) ElementType() reflect.Type {
	return stateMachineTypePtrType
}

func (in *stateMachineTypePtr) ToStateMachineTypePtrOutput() StateMachineTypePtrOutput {
	return pulumi.ToOutput(in).(StateMachineTypePtrOutput)
}

func (in *stateMachineTypePtr) ToStateMachineTypePtrOutputWithContext(ctx context.Context) StateMachineTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StateMachineTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StateMachineLoggingConfigurationLevelInput)(nil)).Elem(), StateMachineLoggingConfigurationLevel("ALL"))
	pulumi.RegisterInputType(reflect.TypeOf((*StateMachineLoggingConfigurationLevelPtrInput)(nil)).Elem(), StateMachineLoggingConfigurationLevel("ALL"))
	pulumi.RegisterInputType(reflect.TypeOf((*StateMachineTypeInput)(nil)).Elem(), StateMachineType("STANDARD"))
	pulumi.RegisterInputType(reflect.TypeOf((*StateMachineTypePtrInput)(nil)).Elem(), StateMachineType("STANDARD"))
	pulumi.RegisterOutputType(StateMachineLoggingConfigurationLevelOutput{})
	pulumi.RegisterOutputType(StateMachineLoggingConfigurationLevelPtrOutput{})
	pulumi.RegisterOutputType(StateMachineTypeOutput{})
	pulumi.RegisterOutputType(StateMachineTypePtrOutput{})
}
