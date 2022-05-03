// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kinesisanalyticsv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::KinesisAnalyticsV2::ApplicationOutput
func LookupApplicationOutputResource(ctx *pulumi.Context, args *LookupApplicationOutputResourceArgs, opts ...pulumi.InvokeOption) (*LookupApplicationOutputResourceResult, error) {
	var rv LookupApplicationOutputResourceResult
	err := ctx.Invoke("aws-native:kinesisanalyticsv2:getApplicationOutputResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationOutputResourceArgs struct {
	Id string `pulumi:"id"`
}

type LookupApplicationOutputResourceResult struct {
	Id     *string                              `pulumi:"id"`
	Output *ApplicationOutputResourceOutputType `pulumi:"output"`
}

func LookupApplicationOutputResourceOutput(ctx *pulumi.Context, args LookupApplicationOutputResourceOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationOutputResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationOutputResourceResult, error) {
			args := v.(LookupApplicationOutputResourceArgs)
			r, err := LookupApplicationOutputResource(ctx, &args, opts...)
			var s LookupApplicationOutputResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationOutputResourceResultOutput)
}

type LookupApplicationOutputResourceOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupApplicationOutputResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationOutputResourceArgs)(nil)).Elem()
}

type LookupApplicationOutputResourceResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationOutputResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationOutputResourceResult)(nil)).Elem()
}

func (o LookupApplicationOutputResourceResultOutput) ToLookupApplicationOutputResourceResultOutput() LookupApplicationOutputResourceResultOutput {
	return o
}

func (o LookupApplicationOutputResourceResultOutput) ToLookupApplicationOutputResourceResultOutputWithContext(ctx context.Context) LookupApplicationOutputResourceResultOutput {
	return o
}

func (o LookupApplicationOutputResourceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationOutputResourceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationOutputResourceResultOutput) Output() ApplicationOutputResourceOutputTypePtrOutput {
	return o.ApplyT(func(v LookupApplicationOutputResourceResult) *ApplicationOutputResourceOutputType { return v.Output }).(ApplicationOutputResourceOutputTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationOutputResourceResultOutput{})
}
