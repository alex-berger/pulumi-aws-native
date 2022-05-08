// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devopsguru

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The type of ResourceCollection
type ResourceCollectionType string

const (
	ResourceCollectionTypeAwsCloudFormation = ResourceCollectionType("AWS_CLOUD_FORMATION")
	ResourceCollectionTypeAwsTags           = ResourceCollectionType("AWS_TAGS")
)

func (ResourceCollectionType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceCollectionType)(nil)).Elem()
}

func (e ResourceCollectionType) ToResourceCollectionTypeOutput() ResourceCollectionTypeOutput {
	return pulumi.ToOutput(e).(ResourceCollectionTypeOutput)
}

func (e ResourceCollectionType) ToResourceCollectionTypeOutputWithContext(ctx context.Context) ResourceCollectionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceCollectionTypeOutput)
}

func (e ResourceCollectionType) ToResourceCollectionTypePtrOutput() ResourceCollectionTypePtrOutput {
	return e.ToResourceCollectionTypePtrOutputWithContext(context.Background())
}

func (e ResourceCollectionType) ToResourceCollectionTypePtrOutputWithContext(ctx context.Context) ResourceCollectionTypePtrOutput {
	return ResourceCollectionType(e).ToResourceCollectionTypeOutputWithContext(ctx).ToResourceCollectionTypePtrOutputWithContext(ctx)
}

func (e ResourceCollectionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceCollectionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceCollectionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceCollectionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceCollectionTypeOutput struct{ *pulumi.OutputState }

func (ResourceCollectionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceCollectionType)(nil)).Elem()
}

func (o ResourceCollectionTypeOutput) ToResourceCollectionTypeOutput() ResourceCollectionTypeOutput {
	return o
}

func (o ResourceCollectionTypeOutput) ToResourceCollectionTypeOutputWithContext(ctx context.Context) ResourceCollectionTypeOutput {
	return o
}

func (o ResourceCollectionTypeOutput) ToResourceCollectionTypePtrOutput() ResourceCollectionTypePtrOutput {
	return o.ToResourceCollectionTypePtrOutputWithContext(context.Background())
}

func (o ResourceCollectionTypeOutput) ToResourceCollectionTypePtrOutputWithContext(ctx context.Context) ResourceCollectionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceCollectionType) *ResourceCollectionType {
		return &v
	}).(ResourceCollectionTypePtrOutput)
}

func (o ResourceCollectionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceCollectionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceCollectionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceCollectionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceCollectionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceCollectionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceCollectionTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceCollectionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceCollectionType)(nil)).Elem()
}

func (o ResourceCollectionTypePtrOutput) ToResourceCollectionTypePtrOutput() ResourceCollectionTypePtrOutput {
	return o
}

func (o ResourceCollectionTypePtrOutput) ToResourceCollectionTypePtrOutputWithContext(ctx context.Context) ResourceCollectionTypePtrOutput {
	return o
}

func (o ResourceCollectionTypePtrOutput) Elem() ResourceCollectionTypeOutput {
	return o.ApplyT(func(v *ResourceCollectionType) ResourceCollectionType {
		if v != nil {
			return *v
		}
		var ret ResourceCollectionType
		return ret
	}).(ResourceCollectionTypeOutput)
}

func (o ResourceCollectionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceCollectionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceCollectionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ResourceCollectionTypeInput is an input type that accepts ResourceCollectionTypeArgs and ResourceCollectionTypeOutput values.
// You can construct a concrete instance of `ResourceCollectionTypeInput` via:
//
//          ResourceCollectionTypeArgs{...}
type ResourceCollectionTypeInput interface {
	pulumi.Input

	ToResourceCollectionTypeOutput() ResourceCollectionTypeOutput
	ToResourceCollectionTypeOutputWithContext(context.Context) ResourceCollectionTypeOutput
}

var resourceCollectionTypePtrType = reflect.TypeOf((**ResourceCollectionType)(nil)).Elem()

type ResourceCollectionTypePtrInput interface {
	pulumi.Input

	ToResourceCollectionTypePtrOutput() ResourceCollectionTypePtrOutput
	ToResourceCollectionTypePtrOutputWithContext(context.Context) ResourceCollectionTypePtrOutput
}

type resourceCollectionTypePtr string

func ResourceCollectionTypePtr(v string) ResourceCollectionTypePtrInput {
	return (*resourceCollectionTypePtr)(&v)
}

func (*resourceCollectionTypePtr) ElementType() reflect.Type {
	return resourceCollectionTypePtrType
}

func (in *resourceCollectionTypePtr) ToResourceCollectionTypePtrOutput() ResourceCollectionTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceCollectionTypePtrOutput)
}

func (in *resourceCollectionTypePtr) ToResourceCollectionTypePtrOutputWithContext(ctx context.Context) ResourceCollectionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceCollectionTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceCollectionTypeInput)(nil)).Elem(), ResourceCollectionType("AWS_CLOUD_FORMATION"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceCollectionTypePtrInput)(nil)).Elem(), ResourceCollectionType("AWS_CLOUD_FORMATION"))
	pulumi.RegisterOutputType(ResourceCollectionTypeOutput{})
	pulumi.RegisterOutputType(ResourceCollectionTypePtrOutput{})
}
