// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dax

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClusterSSESpecification struct {
	SSEEnabled *bool `pulumi:"sSEEnabled"`
}

// ClusterSSESpecificationInput is an input type that accepts ClusterSSESpecificationArgs and ClusterSSESpecificationOutput values.
// You can construct a concrete instance of `ClusterSSESpecificationInput` via:
//
//          ClusterSSESpecificationArgs{...}
type ClusterSSESpecificationInput interface {
	pulumi.Input

	ToClusterSSESpecificationOutput() ClusterSSESpecificationOutput
	ToClusterSSESpecificationOutputWithContext(context.Context) ClusterSSESpecificationOutput
}

type ClusterSSESpecificationArgs struct {
	SSEEnabled pulumi.BoolPtrInput `pulumi:"sSEEnabled"`
}

func (ClusterSSESpecificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSSESpecification)(nil)).Elem()
}

func (i ClusterSSESpecificationArgs) ToClusterSSESpecificationOutput() ClusterSSESpecificationOutput {
	return i.ToClusterSSESpecificationOutputWithContext(context.Background())
}

func (i ClusterSSESpecificationArgs) ToClusterSSESpecificationOutputWithContext(ctx context.Context) ClusterSSESpecificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSSESpecificationOutput)
}

func (i ClusterSSESpecificationArgs) ToClusterSSESpecificationPtrOutput() ClusterSSESpecificationPtrOutput {
	return i.ToClusterSSESpecificationPtrOutputWithContext(context.Background())
}

func (i ClusterSSESpecificationArgs) ToClusterSSESpecificationPtrOutputWithContext(ctx context.Context) ClusterSSESpecificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSSESpecificationOutput).ToClusterSSESpecificationPtrOutputWithContext(ctx)
}

// ClusterSSESpecificationPtrInput is an input type that accepts ClusterSSESpecificationArgs, ClusterSSESpecificationPtr and ClusterSSESpecificationPtrOutput values.
// You can construct a concrete instance of `ClusterSSESpecificationPtrInput` via:
//
//          ClusterSSESpecificationArgs{...}
//
//  or:
//
//          nil
type ClusterSSESpecificationPtrInput interface {
	pulumi.Input

	ToClusterSSESpecificationPtrOutput() ClusterSSESpecificationPtrOutput
	ToClusterSSESpecificationPtrOutputWithContext(context.Context) ClusterSSESpecificationPtrOutput
}

type clusterSSESpecificationPtrType ClusterSSESpecificationArgs

func ClusterSSESpecificationPtr(v *ClusterSSESpecificationArgs) ClusterSSESpecificationPtrInput {
	return (*clusterSSESpecificationPtrType)(v)
}

func (*clusterSSESpecificationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSSESpecification)(nil)).Elem()
}

func (i *clusterSSESpecificationPtrType) ToClusterSSESpecificationPtrOutput() ClusterSSESpecificationPtrOutput {
	return i.ToClusterSSESpecificationPtrOutputWithContext(context.Background())
}

func (i *clusterSSESpecificationPtrType) ToClusterSSESpecificationPtrOutputWithContext(ctx context.Context) ClusterSSESpecificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSSESpecificationPtrOutput)
}

type ClusterSSESpecificationOutput struct{ *pulumi.OutputState }

func (ClusterSSESpecificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSSESpecification)(nil)).Elem()
}

func (o ClusterSSESpecificationOutput) ToClusterSSESpecificationOutput() ClusterSSESpecificationOutput {
	return o
}

func (o ClusterSSESpecificationOutput) ToClusterSSESpecificationOutputWithContext(ctx context.Context) ClusterSSESpecificationOutput {
	return o
}

func (o ClusterSSESpecificationOutput) ToClusterSSESpecificationPtrOutput() ClusterSSESpecificationPtrOutput {
	return o.ToClusterSSESpecificationPtrOutputWithContext(context.Background())
}

func (o ClusterSSESpecificationOutput) ToClusterSSESpecificationPtrOutputWithContext(ctx context.Context) ClusterSSESpecificationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterSSESpecification) *ClusterSSESpecification {
		return &v
	}).(ClusterSSESpecificationPtrOutput)
}

func (o ClusterSSESpecificationOutput) SSEEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ClusterSSESpecification) *bool { return v.SSEEnabled }).(pulumi.BoolPtrOutput)
}

type ClusterSSESpecificationPtrOutput struct{ *pulumi.OutputState }

func (ClusterSSESpecificationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSSESpecification)(nil)).Elem()
}

func (o ClusterSSESpecificationPtrOutput) ToClusterSSESpecificationPtrOutput() ClusterSSESpecificationPtrOutput {
	return o
}

func (o ClusterSSESpecificationPtrOutput) ToClusterSSESpecificationPtrOutputWithContext(ctx context.Context) ClusterSSESpecificationPtrOutput {
	return o
}

func (o ClusterSSESpecificationPtrOutput) Elem() ClusterSSESpecificationOutput {
	return o.ApplyT(func(v *ClusterSSESpecification) ClusterSSESpecification {
		if v != nil {
			return *v
		}
		var ret ClusterSSESpecification
		return ret
	}).(ClusterSSESpecificationOutput)
}

func (o ClusterSSESpecificationPtrOutput) SSEEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClusterSSESpecification) *bool {
		if v == nil {
			return nil
		}
		return v.SSEEnabled
	}).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterSSESpecificationInput)(nil)).Elem(), ClusterSSESpecificationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterSSESpecificationPtrInput)(nil)).Elem(), ClusterSSESpecificationArgs{})
	pulumi.RegisterOutputType(ClusterSSESpecificationOutput{})
	pulumi.RegisterOutputType(ClusterSSESpecificationPtrOutput{})
}