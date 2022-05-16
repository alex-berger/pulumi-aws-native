// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotsitewise

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssetModelDataType string

const (
	AssetModelDataTypeString  = AssetModelDataType("STRING")
	AssetModelDataTypeInteger = AssetModelDataType("INTEGER")
	AssetModelDataTypeDouble  = AssetModelDataType("DOUBLE")
	AssetModelDataTypeBoolean = AssetModelDataType("BOOLEAN")
	AssetModelDataTypeStruct  = AssetModelDataType("STRUCT")
)

func (AssetModelDataType) ElementType() reflect.Type {
	return reflect.TypeOf((*AssetModelDataType)(nil)).Elem()
}

func (e AssetModelDataType) ToAssetModelDataTypeOutput() AssetModelDataTypeOutput {
	return pulumi.ToOutput(e).(AssetModelDataTypeOutput)
}

func (e AssetModelDataType) ToAssetModelDataTypeOutputWithContext(ctx context.Context) AssetModelDataTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AssetModelDataTypeOutput)
}

func (e AssetModelDataType) ToAssetModelDataTypePtrOutput() AssetModelDataTypePtrOutput {
	return e.ToAssetModelDataTypePtrOutputWithContext(context.Background())
}

func (e AssetModelDataType) ToAssetModelDataTypePtrOutputWithContext(ctx context.Context) AssetModelDataTypePtrOutput {
	return AssetModelDataType(e).ToAssetModelDataTypeOutputWithContext(ctx).ToAssetModelDataTypePtrOutputWithContext(ctx)
}

func (e AssetModelDataType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssetModelDataType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssetModelDataType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AssetModelDataType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AssetModelDataTypeOutput struct{ *pulumi.OutputState }

func (AssetModelDataTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssetModelDataType)(nil)).Elem()
}

func (o AssetModelDataTypeOutput) ToAssetModelDataTypeOutput() AssetModelDataTypeOutput {
	return o
}

func (o AssetModelDataTypeOutput) ToAssetModelDataTypeOutputWithContext(ctx context.Context) AssetModelDataTypeOutput {
	return o
}

func (o AssetModelDataTypeOutput) ToAssetModelDataTypePtrOutput() AssetModelDataTypePtrOutput {
	return o.ToAssetModelDataTypePtrOutputWithContext(context.Background())
}

func (o AssetModelDataTypeOutput) ToAssetModelDataTypePtrOutputWithContext(ctx context.Context) AssetModelDataTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssetModelDataType) *AssetModelDataType {
		return &v
	}).(AssetModelDataTypePtrOutput)
}

func (o AssetModelDataTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AssetModelDataTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssetModelDataType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AssetModelDataTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssetModelDataTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssetModelDataType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AssetModelDataTypePtrOutput struct{ *pulumi.OutputState }

func (AssetModelDataTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssetModelDataType)(nil)).Elem()
}

func (o AssetModelDataTypePtrOutput) ToAssetModelDataTypePtrOutput() AssetModelDataTypePtrOutput {
	return o
}

func (o AssetModelDataTypePtrOutput) ToAssetModelDataTypePtrOutputWithContext(ctx context.Context) AssetModelDataTypePtrOutput {
	return o
}

func (o AssetModelDataTypePtrOutput) Elem() AssetModelDataTypeOutput {
	return o.ApplyT(func(v *AssetModelDataType) AssetModelDataType {
		if v != nil {
			return *v
		}
		var ret AssetModelDataType
		return ret
	}).(AssetModelDataTypeOutput)
}

func (o AssetModelDataTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssetModelDataTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AssetModelDataType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AssetModelDataTypeInput is an input type that accepts AssetModelDataTypeArgs and AssetModelDataTypeOutput values.
// You can construct a concrete instance of `AssetModelDataTypeInput` via:
//
//          AssetModelDataTypeArgs{...}
type AssetModelDataTypeInput interface {
	pulumi.Input

	ToAssetModelDataTypeOutput() AssetModelDataTypeOutput
	ToAssetModelDataTypeOutputWithContext(context.Context) AssetModelDataTypeOutput
}

var assetModelDataTypePtrType = reflect.TypeOf((**AssetModelDataType)(nil)).Elem()

type AssetModelDataTypePtrInput interface {
	pulumi.Input

	ToAssetModelDataTypePtrOutput() AssetModelDataTypePtrOutput
	ToAssetModelDataTypePtrOutputWithContext(context.Context) AssetModelDataTypePtrOutput
}

type assetModelDataTypePtr string

func AssetModelDataTypePtr(v string) AssetModelDataTypePtrInput {
	return (*assetModelDataTypePtr)(&v)
}

func (*assetModelDataTypePtr) ElementType() reflect.Type {
	return assetModelDataTypePtrType
}

func (in *assetModelDataTypePtr) ToAssetModelDataTypePtrOutput() AssetModelDataTypePtrOutput {
	return pulumi.ToOutput(in).(AssetModelDataTypePtrOutput)
}

func (in *assetModelDataTypePtr) ToAssetModelDataTypePtrOutputWithContext(ctx context.Context) AssetModelDataTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AssetModelDataTypePtrOutput)
}

type AssetModelDataTypeSpec string

const (
	AssetModelDataTypeSpecAwsalarmState = AssetModelDataTypeSpec("AWS/ALARM_STATE")
)

func (AssetModelDataTypeSpec) ElementType() reflect.Type {
	return reflect.TypeOf((*AssetModelDataTypeSpec)(nil)).Elem()
}

func (e AssetModelDataTypeSpec) ToAssetModelDataTypeSpecOutput() AssetModelDataTypeSpecOutput {
	return pulumi.ToOutput(e).(AssetModelDataTypeSpecOutput)
}

func (e AssetModelDataTypeSpec) ToAssetModelDataTypeSpecOutputWithContext(ctx context.Context) AssetModelDataTypeSpecOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AssetModelDataTypeSpecOutput)
}

func (e AssetModelDataTypeSpec) ToAssetModelDataTypeSpecPtrOutput() AssetModelDataTypeSpecPtrOutput {
	return e.ToAssetModelDataTypeSpecPtrOutputWithContext(context.Background())
}

func (e AssetModelDataTypeSpec) ToAssetModelDataTypeSpecPtrOutputWithContext(ctx context.Context) AssetModelDataTypeSpecPtrOutput {
	return AssetModelDataTypeSpec(e).ToAssetModelDataTypeSpecOutputWithContext(ctx).ToAssetModelDataTypeSpecPtrOutputWithContext(ctx)
}

func (e AssetModelDataTypeSpec) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssetModelDataTypeSpec) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssetModelDataTypeSpec) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AssetModelDataTypeSpec) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AssetModelDataTypeSpecOutput struct{ *pulumi.OutputState }

func (AssetModelDataTypeSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssetModelDataTypeSpec)(nil)).Elem()
}

func (o AssetModelDataTypeSpecOutput) ToAssetModelDataTypeSpecOutput() AssetModelDataTypeSpecOutput {
	return o
}

func (o AssetModelDataTypeSpecOutput) ToAssetModelDataTypeSpecOutputWithContext(ctx context.Context) AssetModelDataTypeSpecOutput {
	return o
}

func (o AssetModelDataTypeSpecOutput) ToAssetModelDataTypeSpecPtrOutput() AssetModelDataTypeSpecPtrOutput {
	return o.ToAssetModelDataTypeSpecPtrOutputWithContext(context.Background())
}

func (o AssetModelDataTypeSpecOutput) ToAssetModelDataTypeSpecPtrOutputWithContext(ctx context.Context) AssetModelDataTypeSpecPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssetModelDataTypeSpec) *AssetModelDataTypeSpec {
		return &v
	}).(AssetModelDataTypeSpecPtrOutput)
}

func (o AssetModelDataTypeSpecOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AssetModelDataTypeSpecOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssetModelDataTypeSpec) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AssetModelDataTypeSpecOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssetModelDataTypeSpecOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssetModelDataTypeSpec) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AssetModelDataTypeSpecPtrOutput struct{ *pulumi.OutputState }

func (AssetModelDataTypeSpecPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssetModelDataTypeSpec)(nil)).Elem()
}

func (o AssetModelDataTypeSpecPtrOutput) ToAssetModelDataTypeSpecPtrOutput() AssetModelDataTypeSpecPtrOutput {
	return o
}

func (o AssetModelDataTypeSpecPtrOutput) ToAssetModelDataTypeSpecPtrOutputWithContext(ctx context.Context) AssetModelDataTypeSpecPtrOutput {
	return o
}

func (o AssetModelDataTypeSpecPtrOutput) Elem() AssetModelDataTypeSpecOutput {
	return o.ApplyT(func(v *AssetModelDataTypeSpec) AssetModelDataTypeSpec {
		if v != nil {
			return *v
		}
		var ret AssetModelDataTypeSpec
		return ret
	}).(AssetModelDataTypeSpecOutput)
}

func (o AssetModelDataTypeSpecPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssetModelDataTypeSpecPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AssetModelDataTypeSpec) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AssetModelDataTypeSpecInput is an input type that accepts AssetModelDataTypeSpecArgs and AssetModelDataTypeSpecOutput values.
// You can construct a concrete instance of `AssetModelDataTypeSpecInput` via:
//
//          AssetModelDataTypeSpecArgs{...}
type AssetModelDataTypeSpecInput interface {
	pulumi.Input

	ToAssetModelDataTypeSpecOutput() AssetModelDataTypeSpecOutput
	ToAssetModelDataTypeSpecOutputWithContext(context.Context) AssetModelDataTypeSpecOutput
}

var assetModelDataTypeSpecPtrType = reflect.TypeOf((**AssetModelDataTypeSpec)(nil)).Elem()

type AssetModelDataTypeSpecPtrInput interface {
	pulumi.Input

	ToAssetModelDataTypeSpecPtrOutput() AssetModelDataTypeSpecPtrOutput
	ToAssetModelDataTypeSpecPtrOutputWithContext(context.Context) AssetModelDataTypeSpecPtrOutput
}

type assetModelDataTypeSpecPtr string

func AssetModelDataTypeSpecPtr(v string) AssetModelDataTypeSpecPtrInput {
	return (*assetModelDataTypeSpecPtr)(&v)
}

func (*assetModelDataTypeSpecPtr) ElementType() reflect.Type {
	return assetModelDataTypeSpecPtrType
}

func (in *assetModelDataTypeSpecPtr) ToAssetModelDataTypeSpecPtrOutput() AssetModelDataTypeSpecPtrOutput {
	return pulumi.ToOutput(in).(AssetModelDataTypeSpecPtrOutput)
}

func (in *assetModelDataTypeSpecPtr) ToAssetModelDataTypeSpecPtrOutputWithContext(ctx context.Context) AssetModelDataTypeSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AssetModelDataTypeSpecPtrOutput)
}

type AssetModelTypeName string

const (
	AssetModelTypeNameMeasurement = AssetModelTypeName("Measurement")
	AssetModelTypeNameAttribute   = AssetModelTypeName("Attribute")
	AssetModelTypeNameTransform   = AssetModelTypeName("Transform")
	AssetModelTypeNameMetric      = AssetModelTypeName("Metric")
)

func (AssetModelTypeName) ElementType() reflect.Type {
	return reflect.TypeOf((*AssetModelTypeName)(nil)).Elem()
}

func (e AssetModelTypeName) ToAssetModelTypeNameOutput() AssetModelTypeNameOutput {
	return pulumi.ToOutput(e).(AssetModelTypeNameOutput)
}

func (e AssetModelTypeName) ToAssetModelTypeNameOutputWithContext(ctx context.Context) AssetModelTypeNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AssetModelTypeNameOutput)
}

func (e AssetModelTypeName) ToAssetModelTypeNamePtrOutput() AssetModelTypeNamePtrOutput {
	return e.ToAssetModelTypeNamePtrOutputWithContext(context.Background())
}

func (e AssetModelTypeName) ToAssetModelTypeNamePtrOutputWithContext(ctx context.Context) AssetModelTypeNamePtrOutput {
	return AssetModelTypeName(e).ToAssetModelTypeNameOutputWithContext(ctx).ToAssetModelTypeNamePtrOutputWithContext(ctx)
}

func (e AssetModelTypeName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssetModelTypeName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssetModelTypeName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AssetModelTypeName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AssetModelTypeNameOutput struct{ *pulumi.OutputState }

func (AssetModelTypeNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssetModelTypeName)(nil)).Elem()
}

func (o AssetModelTypeNameOutput) ToAssetModelTypeNameOutput() AssetModelTypeNameOutput {
	return o
}

func (o AssetModelTypeNameOutput) ToAssetModelTypeNameOutputWithContext(ctx context.Context) AssetModelTypeNameOutput {
	return o
}

func (o AssetModelTypeNameOutput) ToAssetModelTypeNamePtrOutput() AssetModelTypeNamePtrOutput {
	return o.ToAssetModelTypeNamePtrOutputWithContext(context.Background())
}

func (o AssetModelTypeNameOutput) ToAssetModelTypeNamePtrOutputWithContext(ctx context.Context) AssetModelTypeNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssetModelTypeName) *AssetModelTypeName {
		return &v
	}).(AssetModelTypeNamePtrOutput)
}

func (o AssetModelTypeNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AssetModelTypeNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssetModelTypeName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AssetModelTypeNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssetModelTypeNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssetModelTypeName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AssetModelTypeNamePtrOutput struct{ *pulumi.OutputState }

func (AssetModelTypeNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssetModelTypeName)(nil)).Elem()
}

func (o AssetModelTypeNamePtrOutput) ToAssetModelTypeNamePtrOutput() AssetModelTypeNamePtrOutput {
	return o
}

func (o AssetModelTypeNamePtrOutput) ToAssetModelTypeNamePtrOutputWithContext(ctx context.Context) AssetModelTypeNamePtrOutput {
	return o
}

func (o AssetModelTypeNamePtrOutput) Elem() AssetModelTypeNameOutput {
	return o.ApplyT(func(v *AssetModelTypeName) AssetModelTypeName {
		if v != nil {
			return *v
		}
		var ret AssetModelTypeName
		return ret
	}).(AssetModelTypeNameOutput)
}

func (o AssetModelTypeNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssetModelTypeNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AssetModelTypeName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AssetModelTypeNameInput is an input type that accepts AssetModelTypeNameArgs and AssetModelTypeNameOutput values.
// You can construct a concrete instance of `AssetModelTypeNameInput` via:
//
//          AssetModelTypeNameArgs{...}
type AssetModelTypeNameInput interface {
	pulumi.Input

	ToAssetModelTypeNameOutput() AssetModelTypeNameOutput
	ToAssetModelTypeNameOutputWithContext(context.Context) AssetModelTypeNameOutput
}

var assetModelTypeNamePtrType = reflect.TypeOf((**AssetModelTypeName)(nil)).Elem()

type AssetModelTypeNamePtrInput interface {
	pulumi.Input

	ToAssetModelTypeNamePtrOutput() AssetModelTypeNamePtrOutput
	ToAssetModelTypeNamePtrOutputWithContext(context.Context) AssetModelTypeNamePtrOutput
}

type assetModelTypeNamePtr string

func AssetModelTypeNamePtr(v string) AssetModelTypeNamePtrInput {
	return (*assetModelTypeNamePtr)(&v)
}

func (*assetModelTypeNamePtr) ElementType() reflect.Type {
	return assetModelTypeNamePtrType
}

func (in *assetModelTypeNamePtr) ToAssetModelTypeNamePtrOutput() AssetModelTypeNamePtrOutput {
	return pulumi.ToOutput(in).(AssetModelTypeNamePtrOutput)
}

func (in *assetModelTypeNamePtr) ToAssetModelTypeNamePtrOutputWithContext(ctx context.Context) AssetModelTypeNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AssetModelTypeNamePtrOutput)
}

// The MQTT notification state (ENABLED or DISABLED) for this asset property.
type AssetPropertyNotificationState string

const (
	AssetPropertyNotificationStateEnabled  = AssetPropertyNotificationState("ENABLED")
	AssetPropertyNotificationStateDisabled = AssetPropertyNotificationState("DISABLED")
)

func (AssetPropertyNotificationState) ElementType() reflect.Type {
	return reflect.TypeOf((*AssetPropertyNotificationState)(nil)).Elem()
}

func (e AssetPropertyNotificationState) ToAssetPropertyNotificationStateOutput() AssetPropertyNotificationStateOutput {
	return pulumi.ToOutput(e).(AssetPropertyNotificationStateOutput)
}

func (e AssetPropertyNotificationState) ToAssetPropertyNotificationStateOutputWithContext(ctx context.Context) AssetPropertyNotificationStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AssetPropertyNotificationStateOutput)
}

func (e AssetPropertyNotificationState) ToAssetPropertyNotificationStatePtrOutput() AssetPropertyNotificationStatePtrOutput {
	return e.ToAssetPropertyNotificationStatePtrOutputWithContext(context.Background())
}

func (e AssetPropertyNotificationState) ToAssetPropertyNotificationStatePtrOutputWithContext(ctx context.Context) AssetPropertyNotificationStatePtrOutput {
	return AssetPropertyNotificationState(e).ToAssetPropertyNotificationStateOutputWithContext(ctx).ToAssetPropertyNotificationStatePtrOutputWithContext(ctx)
}

func (e AssetPropertyNotificationState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssetPropertyNotificationState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssetPropertyNotificationState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AssetPropertyNotificationState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AssetPropertyNotificationStateOutput struct{ *pulumi.OutputState }

func (AssetPropertyNotificationStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssetPropertyNotificationState)(nil)).Elem()
}

func (o AssetPropertyNotificationStateOutput) ToAssetPropertyNotificationStateOutput() AssetPropertyNotificationStateOutput {
	return o
}

func (o AssetPropertyNotificationStateOutput) ToAssetPropertyNotificationStateOutputWithContext(ctx context.Context) AssetPropertyNotificationStateOutput {
	return o
}

func (o AssetPropertyNotificationStateOutput) ToAssetPropertyNotificationStatePtrOutput() AssetPropertyNotificationStatePtrOutput {
	return o.ToAssetPropertyNotificationStatePtrOutputWithContext(context.Background())
}

func (o AssetPropertyNotificationStateOutput) ToAssetPropertyNotificationStatePtrOutputWithContext(ctx context.Context) AssetPropertyNotificationStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssetPropertyNotificationState) *AssetPropertyNotificationState {
		return &v
	}).(AssetPropertyNotificationStatePtrOutput)
}

func (o AssetPropertyNotificationStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AssetPropertyNotificationStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssetPropertyNotificationState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AssetPropertyNotificationStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssetPropertyNotificationStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssetPropertyNotificationState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AssetPropertyNotificationStatePtrOutput struct{ *pulumi.OutputState }

func (AssetPropertyNotificationStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssetPropertyNotificationState)(nil)).Elem()
}

func (o AssetPropertyNotificationStatePtrOutput) ToAssetPropertyNotificationStatePtrOutput() AssetPropertyNotificationStatePtrOutput {
	return o
}

func (o AssetPropertyNotificationStatePtrOutput) ToAssetPropertyNotificationStatePtrOutputWithContext(ctx context.Context) AssetPropertyNotificationStatePtrOutput {
	return o
}

func (o AssetPropertyNotificationStatePtrOutput) Elem() AssetPropertyNotificationStateOutput {
	return o.ApplyT(func(v *AssetPropertyNotificationState) AssetPropertyNotificationState {
		if v != nil {
			return *v
		}
		var ret AssetPropertyNotificationState
		return ret
	}).(AssetPropertyNotificationStateOutput)
}

func (o AssetPropertyNotificationStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssetPropertyNotificationStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AssetPropertyNotificationState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AssetPropertyNotificationStateInput is an input type that accepts AssetPropertyNotificationStateArgs and AssetPropertyNotificationStateOutput values.
// You can construct a concrete instance of `AssetPropertyNotificationStateInput` via:
//
//          AssetPropertyNotificationStateArgs{...}
type AssetPropertyNotificationStateInput interface {
	pulumi.Input

	ToAssetPropertyNotificationStateOutput() AssetPropertyNotificationStateOutput
	ToAssetPropertyNotificationStateOutputWithContext(context.Context) AssetPropertyNotificationStateOutput
}

var assetPropertyNotificationStatePtrType = reflect.TypeOf((**AssetPropertyNotificationState)(nil)).Elem()

type AssetPropertyNotificationStatePtrInput interface {
	pulumi.Input

	ToAssetPropertyNotificationStatePtrOutput() AssetPropertyNotificationStatePtrOutput
	ToAssetPropertyNotificationStatePtrOutputWithContext(context.Context) AssetPropertyNotificationStatePtrOutput
}

type assetPropertyNotificationStatePtr string

func AssetPropertyNotificationStatePtr(v string) AssetPropertyNotificationStatePtrInput {
	return (*assetPropertyNotificationStatePtr)(&v)
}

func (*assetPropertyNotificationStatePtr) ElementType() reflect.Type {
	return assetPropertyNotificationStatePtrType
}

func (in *assetPropertyNotificationStatePtr) ToAssetPropertyNotificationStatePtrOutput() AssetPropertyNotificationStatePtrOutput {
	return pulumi.ToOutput(in).(AssetPropertyNotificationStatePtrOutput)
}

func (in *assetPropertyNotificationStatePtr) ToAssetPropertyNotificationStatePtrOutputWithContext(ctx context.Context) AssetPropertyNotificationStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AssetPropertyNotificationStatePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AssetModelDataTypeInput)(nil)).Elem(), AssetModelDataType("STRING"))
	pulumi.RegisterInputType(reflect.TypeOf((*AssetModelDataTypePtrInput)(nil)).Elem(), AssetModelDataType("STRING"))
	pulumi.RegisterInputType(reflect.TypeOf((*AssetModelDataTypeSpecInput)(nil)).Elem(), AssetModelDataTypeSpec("AWS/ALARM_STATE"))
	pulumi.RegisterInputType(reflect.TypeOf((*AssetModelDataTypeSpecPtrInput)(nil)).Elem(), AssetModelDataTypeSpec("AWS/ALARM_STATE"))
	pulumi.RegisterInputType(reflect.TypeOf((*AssetModelTypeNameInput)(nil)).Elem(), AssetModelTypeName("Measurement"))
	pulumi.RegisterInputType(reflect.TypeOf((*AssetModelTypeNamePtrInput)(nil)).Elem(), AssetModelTypeName("Measurement"))
	pulumi.RegisterInputType(reflect.TypeOf((*AssetPropertyNotificationStateInput)(nil)).Elem(), AssetPropertyNotificationState("ENABLED"))
	pulumi.RegisterInputType(reflect.TypeOf((*AssetPropertyNotificationStatePtrInput)(nil)).Elem(), AssetPropertyNotificationState("ENABLED"))
	pulumi.RegisterOutputType(AssetModelDataTypeOutput{})
	pulumi.RegisterOutputType(AssetModelDataTypePtrOutput{})
	pulumi.RegisterOutputType(AssetModelDataTypeSpecOutput{})
	pulumi.RegisterOutputType(AssetModelDataTypeSpecPtrOutput{})
	pulumi.RegisterOutputType(AssetModelTypeNameOutput{})
	pulumi.RegisterOutputType(AssetModelTypeNamePtrOutput{})
	pulumi.RegisterOutputType(AssetPropertyNotificationStateOutput{})
	pulumi.RegisterOutputType(AssetPropertyNotificationStatePtrOutput{})
}
