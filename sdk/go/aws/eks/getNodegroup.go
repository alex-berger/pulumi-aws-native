// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eks

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EKS::Nodegroup
func LookupNodegroup(ctx *pulumi.Context, args *LookupNodegroupArgs, opts ...pulumi.InvokeOption) (*LookupNodegroupResult, error) {
	var rv LookupNodegroupResult
	err := ctx.Invoke("aws-native:eks:getNodegroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNodegroupArgs struct {
	Id string `pulumi:"id"`
}

type LookupNodegroupResult struct {
	Arn                *string                               `pulumi:"arn"`
	ForceUpdateEnabled *bool                                 `pulumi:"forceUpdateEnabled"`
	Id                 *string                               `pulumi:"id"`
	Labels             interface{}                           `pulumi:"labels"`
	LaunchTemplate     *NodegroupLaunchTemplateSpecification `pulumi:"launchTemplate"`
	ReleaseVersion     *string                               `pulumi:"releaseVersion"`
	ScalingConfig      *NodegroupScalingConfig               `pulumi:"scalingConfig"`
	Tags               interface{}                           `pulumi:"tags"`
	Taints             []NodegroupTaint                      `pulumi:"taints"`
	UpdateConfig       *NodegroupUpdateConfig                `pulumi:"updateConfig"`
	Version            *string                               `pulumi:"version"`
}

func LookupNodegroupOutput(ctx *pulumi.Context, args LookupNodegroupOutputArgs, opts ...pulumi.InvokeOption) LookupNodegroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNodegroupResult, error) {
			args := v.(LookupNodegroupArgs)
			r, err := LookupNodegroup(ctx, &args, opts...)
			return *r, err
		}).(LookupNodegroupResultOutput)
}

type LookupNodegroupOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupNodegroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNodegroupArgs)(nil)).Elem()
}

type LookupNodegroupResultOutput struct{ *pulumi.OutputState }

func (LookupNodegroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNodegroupResult)(nil)).Elem()
}

func (o LookupNodegroupResultOutput) ToLookupNodegroupResultOutput() LookupNodegroupResultOutput {
	return o
}

func (o LookupNodegroupResultOutput) ToLookupNodegroupResultOutputWithContext(ctx context.Context) LookupNodegroupResultOutput {
	return o
}

func (o LookupNodegroupResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodegroupResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupNodegroupResultOutput) ForceUpdateEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNodegroupResult) *bool { return v.ForceUpdateEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupNodegroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodegroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupNodegroupResultOutput) Labels() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupNodegroupResult) interface{} { return v.Labels }).(pulumi.AnyOutput)
}

func (o LookupNodegroupResultOutput) LaunchTemplate() NodegroupLaunchTemplateSpecificationPtrOutput {
	return o.ApplyT(func(v LookupNodegroupResult) *NodegroupLaunchTemplateSpecification { return v.LaunchTemplate }).(NodegroupLaunchTemplateSpecificationPtrOutput)
}

func (o LookupNodegroupResultOutput) ReleaseVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodegroupResult) *string { return v.ReleaseVersion }).(pulumi.StringPtrOutput)
}

func (o LookupNodegroupResultOutput) ScalingConfig() NodegroupScalingConfigPtrOutput {
	return o.ApplyT(func(v LookupNodegroupResult) *NodegroupScalingConfig { return v.ScalingConfig }).(NodegroupScalingConfigPtrOutput)
}

func (o LookupNodegroupResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupNodegroupResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func (o LookupNodegroupResultOutput) Taints() NodegroupTaintArrayOutput {
	return o.ApplyT(func(v LookupNodegroupResult) []NodegroupTaint { return v.Taints }).(NodegroupTaintArrayOutput)
}

func (o LookupNodegroupResultOutput) UpdateConfig() NodegroupUpdateConfigPtrOutput {
	return o.ApplyT(func(v LookupNodegroupResult) *NodegroupUpdateConfig { return v.UpdateConfig }).(NodegroupUpdateConfigPtrOutput)
}

func (o LookupNodegroupResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodegroupResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNodegroupResultOutput{})
}