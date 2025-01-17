// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package timestream

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Type for the dimension.
type ScheduledQueryDimensionValueType string

const (
	ScheduledQueryDimensionValueTypeVarchar = ScheduledQueryDimensionValueType("VARCHAR")
)

func (ScheduledQueryDimensionValueType) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledQueryDimensionValueType)(nil)).Elem()
}

func (e ScheduledQueryDimensionValueType) ToScheduledQueryDimensionValueTypeOutput() ScheduledQueryDimensionValueTypeOutput {
	return pulumi.ToOutput(e).(ScheduledQueryDimensionValueTypeOutput)
}

func (e ScheduledQueryDimensionValueType) ToScheduledQueryDimensionValueTypeOutputWithContext(ctx context.Context) ScheduledQueryDimensionValueTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScheduledQueryDimensionValueTypeOutput)
}

func (e ScheduledQueryDimensionValueType) ToScheduledQueryDimensionValueTypePtrOutput() ScheduledQueryDimensionValueTypePtrOutput {
	return e.ToScheduledQueryDimensionValueTypePtrOutputWithContext(context.Background())
}

func (e ScheduledQueryDimensionValueType) ToScheduledQueryDimensionValueTypePtrOutputWithContext(ctx context.Context) ScheduledQueryDimensionValueTypePtrOutput {
	return ScheduledQueryDimensionValueType(e).ToScheduledQueryDimensionValueTypeOutputWithContext(ctx).ToScheduledQueryDimensionValueTypePtrOutputWithContext(ctx)
}

func (e ScheduledQueryDimensionValueType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScheduledQueryDimensionValueType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScheduledQueryDimensionValueType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScheduledQueryDimensionValueType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScheduledQueryDimensionValueTypeOutput struct{ *pulumi.OutputState }

func (ScheduledQueryDimensionValueTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledQueryDimensionValueType)(nil)).Elem()
}

func (o ScheduledQueryDimensionValueTypeOutput) ToScheduledQueryDimensionValueTypeOutput() ScheduledQueryDimensionValueTypeOutput {
	return o
}

func (o ScheduledQueryDimensionValueTypeOutput) ToScheduledQueryDimensionValueTypeOutputWithContext(ctx context.Context) ScheduledQueryDimensionValueTypeOutput {
	return o
}

func (o ScheduledQueryDimensionValueTypeOutput) ToScheduledQueryDimensionValueTypePtrOutput() ScheduledQueryDimensionValueTypePtrOutput {
	return o.ToScheduledQueryDimensionValueTypePtrOutputWithContext(context.Background())
}

func (o ScheduledQueryDimensionValueTypeOutput) ToScheduledQueryDimensionValueTypePtrOutputWithContext(ctx context.Context) ScheduledQueryDimensionValueTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScheduledQueryDimensionValueType) *ScheduledQueryDimensionValueType {
		return &v
	}).(ScheduledQueryDimensionValueTypePtrOutput)
}

func (o ScheduledQueryDimensionValueTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScheduledQueryDimensionValueTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScheduledQueryDimensionValueType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScheduledQueryDimensionValueTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScheduledQueryDimensionValueTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScheduledQueryDimensionValueType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScheduledQueryDimensionValueTypePtrOutput struct{ *pulumi.OutputState }

func (ScheduledQueryDimensionValueTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledQueryDimensionValueType)(nil)).Elem()
}

func (o ScheduledQueryDimensionValueTypePtrOutput) ToScheduledQueryDimensionValueTypePtrOutput() ScheduledQueryDimensionValueTypePtrOutput {
	return o
}

func (o ScheduledQueryDimensionValueTypePtrOutput) ToScheduledQueryDimensionValueTypePtrOutputWithContext(ctx context.Context) ScheduledQueryDimensionValueTypePtrOutput {
	return o
}

func (o ScheduledQueryDimensionValueTypePtrOutput) Elem() ScheduledQueryDimensionValueTypeOutput {
	return o.ApplyT(func(v *ScheduledQueryDimensionValueType) ScheduledQueryDimensionValueType {
		if v != nil {
			return *v
		}
		var ret ScheduledQueryDimensionValueType
		return ret
	}).(ScheduledQueryDimensionValueTypeOutput)
}

func (o ScheduledQueryDimensionValueTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScheduledQueryDimensionValueTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScheduledQueryDimensionValueType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ScheduledQueryDimensionValueTypeInput is an input type that accepts ScheduledQueryDimensionValueTypeArgs and ScheduledQueryDimensionValueTypeOutput values.
// You can construct a concrete instance of `ScheduledQueryDimensionValueTypeInput` via:
//
//          ScheduledQueryDimensionValueTypeArgs{...}
type ScheduledQueryDimensionValueTypeInput interface {
	pulumi.Input

	ToScheduledQueryDimensionValueTypeOutput() ScheduledQueryDimensionValueTypeOutput
	ToScheduledQueryDimensionValueTypeOutputWithContext(context.Context) ScheduledQueryDimensionValueTypeOutput
}

var scheduledQueryDimensionValueTypePtrType = reflect.TypeOf((**ScheduledQueryDimensionValueType)(nil)).Elem()

type ScheduledQueryDimensionValueTypePtrInput interface {
	pulumi.Input

	ToScheduledQueryDimensionValueTypePtrOutput() ScheduledQueryDimensionValueTypePtrOutput
	ToScheduledQueryDimensionValueTypePtrOutputWithContext(context.Context) ScheduledQueryDimensionValueTypePtrOutput
}

type scheduledQueryDimensionValueTypePtr string

func ScheduledQueryDimensionValueTypePtr(v string) ScheduledQueryDimensionValueTypePtrInput {
	return (*scheduledQueryDimensionValueTypePtr)(&v)
}

func (*scheduledQueryDimensionValueTypePtr) ElementType() reflect.Type {
	return scheduledQueryDimensionValueTypePtrType
}

func (in *scheduledQueryDimensionValueTypePtr) ToScheduledQueryDimensionValueTypePtrOutput() ScheduledQueryDimensionValueTypePtrOutput {
	return pulumi.ToOutput(in).(ScheduledQueryDimensionValueTypePtrOutput)
}

func (in *scheduledQueryDimensionValueTypePtr) ToScheduledQueryDimensionValueTypePtrOutputWithContext(ctx context.Context) ScheduledQueryDimensionValueTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScheduledQueryDimensionValueTypePtrOutput)
}

// Encryption at rest options for the error reports. If no encryption option is specified, Timestream will choose SSE_S3 as default.
type ScheduledQueryEncryptionOption string

const (
	ScheduledQueryEncryptionOptionSseS3  = ScheduledQueryEncryptionOption("SSE_S3")
	ScheduledQueryEncryptionOptionSseKms = ScheduledQueryEncryptionOption("SSE_KMS")
)

func (ScheduledQueryEncryptionOption) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledQueryEncryptionOption)(nil)).Elem()
}

func (e ScheduledQueryEncryptionOption) ToScheduledQueryEncryptionOptionOutput() ScheduledQueryEncryptionOptionOutput {
	return pulumi.ToOutput(e).(ScheduledQueryEncryptionOptionOutput)
}

func (e ScheduledQueryEncryptionOption) ToScheduledQueryEncryptionOptionOutputWithContext(ctx context.Context) ScheduledQueryEncryptionOptionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScheduledQueryEncryptionOptionOutput)
}

func (e ScheduledQueryEncryptionOption) ToScheduledQueryEncryptionOptionPtrOutput() ScheduledQueryEncryptionOptionPtrOutput {
	return e.ToScheduledQueryEncryptionOptionPtrOutputWithContext(context.Background())
}

func (e ScheduledQueryEncryptionOption) ToScheduledQueryEncryptionOptionPtrOutputWithContext(ctx context.Context) ScheduledQueryEncryptionOptionPtrOutput {
	return ScheduledQueryEncryptionOption(e).ToScheduledQueryEncryptionOptionOutputWithContext(ctx).ToScheduledQueryEncryptionOptionPtrOutputWithContext(ctx)
}

func (e ScheduledQueryEncryptionOption) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScheduledQueryEncryptionOption) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScheduledQueryEncryptionOption) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScheduledQueryEncryptionOption) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScheduledQueryEncryptionOptionOutput struct{ *pulumi.OutputState }

func (ScheduledQueryEncryptionOptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledQueryEncryptionOption)(nil)).Elem()
}

func (o ScheduledQueryEncryptionOptionOutput) ToScheduledQueryEncryptionOptionOutput() ScheduledQueryEncryptionOptionOutput {
	return o
}

func (o ScheduledQueryEncryptionOptionOutput) ToScheduledQueryEncryptionOptionOutputWithContext(ctx context.Context) ScheduledQueryEncryptionOptionOutput {
	return o
}

func (o ScheduledQueryEncryptionOptionOutput) ToScheduledQueryEncryptionOptionPtrOutput() ScheduledQueryEncryptionOptionPtrOutput {
	return o.ToScheduledQueryEncryptionOptionPtrOutputWithContext(context.Background())
}

func (o ScheduledQueryEncryptionOptionOutput) ToScheduledQueryEncryptionOptionPtrOutputWithContext(ctx context.Context) ScheduledQueryEncryptionOptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScheduledQueryEncryptionOption) *ScheduledQueryEncryptionOption {
		return &v
	}).(ScheduledQueryEncryptionOptionPtrOutput)
}

func (o ScheduledQueryEncryptionOptionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScheduledQueryEncryptionOptionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScheduledQueryEncryptionOption) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScheduledQueryEncryptionOptionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScheduledQueryEncryptionOptionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScheduledQueryEncryptionOption) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScheduledQueryEncryptionOptionPtrOutput struct{ *pulumi.OutputState }

func (ScheduledQueryEncryptionOptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledQueryEncryptionOption)(nil)).Elem()
}

func (o ScheduledQueryEncryptionOptionPtrOutput) ToScheduledQueryEncryptionOptionPtrOutput() ScheduledQueryEncryptionOptionPtrOutput {
	return o
}

func (o ScheduledQueryEncryptionOptionPtrOutput) ToScheduledQueryEncryptionOptionPtrOutputWithContext(ctx context.Context) ScheduledQueryEncryptionOptionPtrOutput {
	return o
}

func (o ScheduledQueryEncryptionOptionPtrOutput) Elem() ScheduledQueryEncryptionOptionOutput {
	return o.ApplyT(func(v *ScheduledQueryEncryptionOption) ScheduledQueryEncryptionOption {
		if v != nil {
			return *v
		}
		var ret ScheduledQueryEncryptionOption
		return ret
	}).(ScheduledQueryEncryptionOptionOutput)
}

func (o ScheduledQueryEncryptionOptionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScheduledQueryEncryptionOptionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScheduledQueryEncryptionOption) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ScheduledQueryEncryptionOptionInput is an input type that accepts ScheduledQueryEncryptionOptionArgs and ScheduledQueryEncryptionOptionOutput values.
// You can construct a concrete instance of `ScheduledQueryEncryptionOptionInput` via:
//
//          ScheduledQueryEncryptionOptionArgs{...}
type ScheduledQueryEncryptionOptionInput interface {
	pulumi.Input

	ToScheduledQueryEncryptionOptionOutput() ScheduledQueryEncryptionOptionOutput
	ToScheduledQueryEncryptionOptionOutputWithContext(context.Context) ScheduledQueryEncryptionOptionOutput
}

var scheduledQueryEncryptionOptionPtrType = reflect.TypeOf((**ScheduledQueryEncryptionOption)(nil)).Elem()

type ScheduledQueryEncryptionOptionPtrInput interface {
	pulumi.Input

	ToScheduledQueryEncryptionOptionPtrOutput() ScheduledQueryEncryptionOptionPtrOutput
	ToScheduledQueryEncryptionOptionPtrOutputWithContext(context.Context) ScheduledQueryEncryptionOptionPtrOutput
}

type scheduledQueryEncryptionOptionPtr string

func ScheduledQueryEncryptionOptionPtr(v string) ScheduledQueryEncryptionOptionPtrInput {
	return (*scheduledQueryEncryptionOptionPtr)(&v)
}

func (*scheduledQueryEncryptionOptionPtr) ElementType() reflect.Type {
	return scheduledQueryEncryptionOptionPtrType
}

func (in *scheduledQueryEncryptionOptionPtr) ToScheduledQueryEncryptionOptionPtrOutput() ScheduledQueryEncryptionOptionPtrOutput {
	return pulumi.ToOutput(in).(ScheduledQueryEncryptionOptionPtrOutput)
}

func (in *scheduledQueryEncryptionOptionPtr) ToScheduledQueryEncryptionOptionPtrOutputWithContext(ctx context.Context) ScheduledQueryEncryptionOptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScheduledQueryEncryptionOptionPtrOutput)
}

// Type of the value that is to be read from SourceColumn. If the mapping is for MULTI, use MeasureValueType.MULTI.
type ScheduledQueryMixedMeasureMappingMeasureValueType string

const (
	ScheduledQueryMixedMeasureMappingMeasureValueTypeBigint  = ScheduledQueryMixedMeasureMappingMeasureValueType("BIGINT")
	ScheduledQueryMixedMeasureMappingMeasureValueTypeBoolean = ScheduledQueryMixedMeasureMappingMeasureValueType("BOOLEAN")
	ScheduledQueryMixedMeasureMappingMeasureValueTypeDouble  = ScheduledQueryMixedMeasureMappingMeasureValueType("DOUBLE")
	ScheduledQueryMixedMeasureMappingMeasureValueTypeVarchar = ScheduledQueryMixedMeasureMappingMeasureValueType("VARCHAR")
	ScheduledQueryMixedMeasureMappingMeasureValueTypeMulti   = ScheduledQueryMixedMeasureMappingMeasureValueType("MULTI")
)

func (ScheduledQueryMixedMeasureMappingMeasureValueType) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledQueryMixedMeasureMappingMeasureValueType)(nil)).Elem()
}

func (e ScheduledQueryMixedMeasureMappingMeasureValueType) ToScheduledQueryMixedMeasureMappingMeasureValueTypeOutput() ScheduledQueryMixedMeasureMappingMeasureValueTypeOutput {
	return pulumi.ToOutput(e).(ScheduledQueryMixedMeasureMappingMeasureValueTypeOutput)
}

func (e ScheduledQueryMixedMeasureMappingMeasureValueType) ToScheduledQueryMixedMeasureMappingMeasureValueTypeOutputWithContext(ctx context.Context) ScheduledQueryMixedMeasureMappingMeasureValueTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScheduledQueryMixedMeasureMappingMeasureValueTypeOutput)
}

func (e ScheduledQueryMixedMeasureMappingMeasureValueType) ToScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput() ScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput {
	return e.ToScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutputWithContext(context.Background())
}

func (e ScheduledQueryMixedMeasureMappingMeasureValueType) ToScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutputWithContext(ctx context.Context) ScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput {
	return ScheduledQueryMixedMeasureMappingMeasureValueType(e).ToScheduledQueryMixedMeasureMappingMeasureValueTypeOutputWithContext(ctx).ToScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutputWithContext(ctx)
}

func (e ScheduledQueryMixedMeasureMappingMeasureValueType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScheduledQueryMixedMeasureMappingMeasureValueType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScheduledQueryMixedMeasureMappingMeasureValueType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScheduledQueryMixedMeasureMappingMeasureValueType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScheduledQueryMixedMeasureMappingMeasureValueTypeOutput struct{ *pulumi.OutputState }

func (ScheduledQueryMixedMeasureMappingMeasureValueTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledQueryMixedMeasureMappingMeasureValueType)(nil)).Elem()
}

func (o ScheduledQueryMixedMeasureMappingMeasureValueTypeOutput) ToScheduledQueryMixedMeasureMappingMeasureValueTypeOutput() ScheduledQueryMixedMeasureMappingMeasureValueTypeOutput {
	return o
}

func (o ScheduledQueryMixedMeasureMappingMeasureValueTypeOutput) ToScheduledQueryMixedMeasureMappingMeasureValueTypeOutputWithContext(ctx context.Context) ScheduledQueryMixedMeasureMappingMeasureValueTypeOutput {
	return o
}

func (o ScheduledQueryMixedMeasureMappingMeasureValueTypeOutput) ToScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput() ScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput {
	return o.ToScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutputWithContext(context.Background())
}

func (o ScheduledQueryMixedMeasureMappingMeasureValueTypeOutput) ToScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutputWithContext(ctx context.Context) ScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScheduledQueryMixedMeasureMappingMeasureValueType) *ScheduledQueryMixedMeasureMappingMeasureValueType {
		return &v
	}).(ScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput)
}

func (o ScheduledQueryMixedMeasureMappingMeasureValueTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScheduledQueryMixedMeasureMappingMeasureValueTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScheduledQueryMixedMeasureMappingMeasureValueType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScheduledQueryMixedMeasureMappingMeasureValueTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScheduledQueryMixedMeasureMappingMeasureValueTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScheduledQueryMixedMeasureMappingMeasureValueType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput struct{ *pulumi.OutputState }

func (ScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledQueryMixedMeasureMappingMeasureValueType)(nil)).Elem()
}

func (o ScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput) ToScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput() ScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput {
	return o
}

func (o ScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput) ToScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutputWithContext(ctx context.Context) ScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput {
	return o
}

func (o ScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput) Elem() ScheduledQueryMixedMeasureMappingMeasureValueTypeOutput {
	return o.ApplyT(func(v *ScheduledQueryMixedMeasureMappingMeasureValueType) ScheduledQueryMixedMeasureMappingMeasureValueType {
		if v != nil {
			return *v
		}
		var ret ScheduledQueryMixedMeasureMappingMeasureValueType
		return ret
	}).(ScheduledQueryMixedMeasureMappingMeasureValueTypeOutput)
}

func (o ScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScheduledQueryMixedMeasureMappingMeasureValueType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ScheduledQueryMixedMeasureMappingMeasureValueTypeInput is an input type that accepts ScheduledQueryMixedMeasureMappingMeasureValueTypeArgs and ScheduledQueryMixedMeasureMappingMeasureValueTypeOutput values.
// You can construct a concrete instance of `ScheduledQueryMixedMeasureMappingMeasureValueTypeInput` via:
//
//          ScheduledQueryMixedMeasureMappingMeasureValueTypeArgs{...}
type ScheduledQueryMixedMeasureMappingMeasureValueTypeInput interface {
	pulumi.Input

	ToScheduledQueryMixedMeasureMappingMeasureValueTypeOutput() ScheduledQueryMixedMeasureMappingMeasureValueTypeOutput
	ToScheduledQueryMixedMeasureMappingMeasureValueTypeOutputWithContext(context.Context) ScheduledQueryMixedMeasureMappingMeasureValueTypeOutput
}

var scheduledQueryMixedMeasureMappingMeasureValueTypePtrType = reflect.TypeOf((**ScheduledQueryMixedMeasureMappingMeasureValueType)(nil)).Elem()

type ScheduledQueryMixedMeasureMappingMeasureValueTypePtrInput interface {
	pulumi.Input

	ToScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput() ScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput
	ToScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutputWithContext(context.Context) ScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput
}

type scheduledQueryMixedMeasureMappingMeasureValueTypePtr string

func ScheduledQueryMixedMeasureMappingMeasureValueTypePtr(v string) ScheduledQueryMixedMeasureMappingMeasureValueTypePtrInput {
	return (*scheduledQueryMixedMeasureMappingMeasureValueTypePtr)(&v)
}

func (*scheduledQueryMixedMeasureMappingMeasureValueTypePtr) ElementType() reflect.Type {
	return scheduledQueryMixedMeasureMappingMeasureValueTypePtrType
}

func (in *scheduledQueryMixedMeasureMappingMeasureValueTypePtr) ToScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput() ScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput {
	return pulumi.ToOutput(in).(ScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput)
}

func (in *scheduledQueryMixedMeasureMappingMeasureValueTypePtr) ToScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutputWithContext(ctx context.Context) ScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput)
}

// Value type of the measure value column to be read from the query result.
type ScheduledQueryMultiMeasureAttributeMappingMeasureValueType string

const (
	ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeBigint  = ScheduledQueryMultiMeasureAttributeMappingMeasureValueType("BIGINT")
	ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeBoolean = ScheduledQueryMultiMeasureAttributeMappingMeasureValueType("BOOLEAN")
	ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeDouble  = ScheduledQueryMultiMeasureAttributeMappingMeasureValueType("DOUBLE")
	ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeVarchar = ScheduledQueryMultiMeasureAttributeMappingMeasureValueType("VARCHAR")
)

func (ScheduledQueryMultiMeasureAttributeMappingMeasureValueType) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledQueryMultiMeasureAttributeMappingMeasureValueType)(nil)).Elem()
}

func (e ScheduledQueryMultiMeasureAttributeMappingMeasureValueType) ToScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput() ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput {
	return pulumi.ToOutput(e).(ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput)
}

func (e ScheduledQueryMultiMeasureAttributeMappingMeasureValueType) ToScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutputWithContext(ctx context.Context) ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput)
}

func (e ScheduledQueryMultiMeasureAttributeMappingMeasureValueType) ToScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput() ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput {
	return e.ToScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutputWithContext(context.Background())
}

func (e ScheduledQueryMultiMeasureAttributeMappingMeasureValueType) ToScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutputWithContext(ctx context.Context) ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput {
	return ScheduledQueryMultiMeasureAttributeMappingMeasureValueType(e).ToScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutputWithContext(ctx).ToScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutputWithContext(ctx)
}

func (e ScheduledQueryMultiMeasureAttributeMappingMeasureValueType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScheduledQueryMultiMeasureAttributeMappingMeasureValueType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScheduledQueryMultiMeasureAttributeMappingMeasureValueType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScheduledQueryMultiMeasureAttributeMappingMeasureValueType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput struct{ *pulumi.OutputState }

func (ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledQueryMultiMeasureAttributeMappingMeasureValueType)(nil)).Elem()
}

func (o ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput) ToScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput() ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput {
	return o
}

func (o ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput) ToScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutputWithContext(ctx context.Context) ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput {
	return o
}

func (o ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput) ToScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput() ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput {
	return o.ToScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutputWithContext(context.Background())
}

func (o ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput) ToScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutputWithContext(ctx context.Context) ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScheduledQueryMultiMeasureAttributeMappingMeasureValueType) *ScheduledQueryMultiMeasureAttributeMappingMeasureValueType {
		return &v
	}).(ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput)
}

func (o ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScheduledQueryMultiMeasureAttributeMappingMeasureValueType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScheduledQueryMultiMeasureAttributeMappingMeasureValueType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput struct{ *pulumi.OutputState }

func (ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledQueryMultiMeasureAttributeMappingMeasureValueType)(nil)).Elem()
}

func (o ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput) ToScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput() ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput {
	return o
}

func (o ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput) ToScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutputWithContext(ctx context.Context) ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput {
	return o
}

func (o ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput) Elem() ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput {
	return o.ApplyT(func(v *ScheduledQueryMultiMeasureAttributeMappingMeasureValueType) ScheduledQueryMultiMeasureAttributeMappingMeasureValueType {
		if v != nil {
			return *v
		}
		var ret ScheduledQueryMultiMeasureAttributeMappingMeasureValueType
		return ret
	}).(ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput)
}

func (o ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScheduledQueryMultiMeasureAttributeMappingMeasureValueType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeInput is an input type that accepts ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeArgs and ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput values.
// You can construct a concrete instance of `ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeInput` via:
//
//          ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeArgs{...}
type ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeInput interface {
	pulumi.Input

	ToScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput() ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput
	ToScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutputWithContext(context.Context) ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput
}

var scheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrType = reflect.TypeOf((**ScheduledQueryMultiMeasureAttributeMappingMeasureValueType)(nil)).Elem()

type ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrInput interface {
	pulumi.Input

	ToScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput() ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput
	ToScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutputWithContext(context.Context) ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput
}

type scheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtr string

func ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtr(v string) ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrInput {
	return (*scheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtr)(&v)
}

func (*scheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtr) ElementType() reflect.Type {
	return scheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrType
}

func (in *scheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtr) ToScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput() ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput {
	return pulumi.ToOutput(in).(ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput)
}

func (in *scheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtr) ToScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutputWithContext(ctx context.Context) ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledQueryDimensionValueTypeInput)(nil)).Elem(), ScheduledQueryDimensionValueType("VARCHAR"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledQueryDimensionValueTypePtrInput)(nil)).Elem(), ScheduledQueryDimensionValueType("VARCHAR"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledQueryEncryptionOptionInput)(nil)).Elem(), ScheduledQueryEncryptionOption("SSE_S3"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledQueryEncryptionOptionPtrInput)(nil)).Elem(), ScheduledQueryEncryptionOption("SSE_S3"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledQueryMixedMeasureMappingMeasureValueTypeInput)(nil)).Elem(), ScheduledQueryMixedMeasureMappingMeasureValueType("BIGINT"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledQueryMixedMeasureMappingMeasureValueTypePtrInput)(nil)).Elem(), ScheduledQueryMixedMeasureMappingMeasureValueType("BIGINT"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeInput)(nil)).Elem(), ScheduledQueryMultiMeasureAttributeMappingMeasureValueType("BIGINT"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrInput)(nil)).Elem(), ScheduledQueryMultiMeasureAttributeMappingMeasureValueType("BIGINT"))
	pulumi.RegisterOutputType(ScheduledQueryDimensionValueTypeOutput{})
	pulumi.RegisterOutputType(ScheduledQueryDimensionValueTypePtrOutput{})
	pulumi.RegisterOutputType(ScheduledQueryEncryptionOptionOutput{})
	pulumi.RegisterOutputType(ScheduledQueryEncryptionOptionPtrOutput{})
	pulumi.RegisterOutputType(ScheduledQueryMixedMeasureMappingMeasureValueTypeOutput{})
	pulumi.RegisterOutputType(ScheduledQueryMixedMeasureMappingMeasureValueTypePtrOutput{})
	pulumi.RegisterOutputType(ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypeOutput{})
	pulumi.RegisterOutputType(ScheduledQueryMultiMeasureAttributeMappingMeasureValueTypePtrOutput{})
}
