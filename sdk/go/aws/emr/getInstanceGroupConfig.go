// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EMR::InstanceGroupConfig
func LookupInstanceGroupConfig(ctx *pulumi.Context, args *LookupInstanceGroupConfigArgs, opts ...pulumi.InvokeOption) (*LookupInstanceGroupConfigResult, error) {
	var rv LookupInstanceGroupConfigResult
	err := ctx.Invoke("aws-native:emr:getInstanceGroupConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceGroupConfigArgs struct {
	Id string `pulumi:"id"`
}

type LookupInstanceGroupConfigResult struct {
	AutoScalingPolicy *InstanceGroupConfigAutoScalingPolicy `pulumi:"autoScalingPolicy"`
	Id                *string                               `pulumi:"id"`
	InstanceCount     *int                                  `pulumi:"instanceCount"`
}

func LookupInstanceGroupConfigOutput(ctx *pulumi.Context, args LookupInstanceGroupConfigOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceGroupConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceGroupConfigResult, error) {
			args := v.(LookupInstanceGroupConfigArgs)
			r, err := LookupInstanceGroupConfig(ctx, &args, opts...)
			var s LookupInstanceGroupConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstanceGroupConfigResultOutput)
}

type LookupInstanceGroupConfigOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupInstanceGroupConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceGroupConfigArgs)(nil)).Elem()
}

type LookupInstanceGroupConfigResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceGroupConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceGroupConfigResult)(nil)).Elem()
}

func (o LookupInstanceGroupConfigResultOutput) ToLookupInstanceGroupConfigResultOutput() LookupInstanceGroupConfigResultOutput {
	return o
}

func (o LookupInstanceGroupConfigResultOutput) ToLookupInstanceGroupConfigResultOutputWithContext(ctx context.Context) LookupInstanceGroupConfigResultOutput {
	return o
}

func (o LookupInstanceGroupConfigResultOutput) AutoScalingPolicy() InstanceGroupConfigAutoScalingPolicyPtrOutput {
	return o.ApplyT(func(v LookupInstanceGroupConfigResult) *InstanceGroupConfigAutoScalingPolicy {
		return v.AutoScalingPolicy
	}).(InstanceGroupConfigAutoScalingPolicyPtrOutput)
}

func (o LookupInstanceGroupConfigResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceGroupConfigResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceGroupConfigResultOutput) InstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupInstanceGroupConfigResult) *int { return v.InstanceCount }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceGroupConfigResultOutput{})
}
