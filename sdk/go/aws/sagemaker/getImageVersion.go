// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::ImageVersion
func LookupImageVersion(ctx *pulumi.Context, args *LookupImageVersionArgs, opts ...pulumi.InvokeOption) (*LookupImageVersionResult, error) {
	var rv LookupImageVersionResult
	err := ctx.Invoke("aws-native:sagemaker:getImageVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupImageVersionArgs struct {
	ImageVersionArn string `pulumi:"imageVersionArn"`
}

type LookupImageVersionResult struct {
	ContainerImage  *string `pulumi:"containerImage"`
	ImageArn        *string `pulumi:"imageArn"`
	ImageVersionArn *string `pulumi:"imageVersionArn"`
	Version         *int    `pulumi:"version"`
}

func LookupImageVersionOutput(ctx *pulumi.Context, args LookupImageVersionOutputArgs, opts ...pulumi.InvokeOption) LookupImageVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupImageVersionResult, error) {
			args := v.(LookupImageVersionArgs)
			r, err := LookupImageVersion(ctx, &args, opts...)
			var s LookupImageVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupImageVersionResultOutput)
}

type LookupImageVersionOutputArgs struct {
	ImageVersionArn pulumi.StringInput `pulumi:"imageVersionArn"`
}

func (LookupImageVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageVersionArgs)(nil)).Elem()
}

type LookupImageVersionResultOutput struct{ *pulumi.OutputState }

func (LookupImageVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageVersionResult)(nil)).Elem()
}

func (o LookupImageVersionResultOutput) ToLookupImageVersionResultOutput() LookupImageVersionResultOutput {
	return o
}

func (o LookupImageVersionResultOutput) ToLookupImageVersionResultOutputWithContext(ctx context.Context) LookupImageVersionResultOutput {
	return o
}

func (o LookupImageVersionResultOutput) ContainerImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageVersionResult) *string { return v.ContainerImage }).(pulumi.StringPtrOutput)
}

func (o LookupImageVersionResultOutput) ImageArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageVersionResult) *string { return v.ImageArn }).(pulumi.StringPtrOutput)
}

func (o LookupImageVersionResultOutput) ImageVersionArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageVersionResult) *string { return v.ImageVersionArn }).(pulumi.StringPtrOutput)
}

func (o LookupImageVersionResultOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupImageVersionResult) *int { return v.Version }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupImageVersionResultOutput{})
}
