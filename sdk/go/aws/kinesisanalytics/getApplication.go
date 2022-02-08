// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kinesisanalytics

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::KinesisAnalytics::Application
func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	var rv LookupApplicationResult
	err := ctx.Invoke("aws-native:kinesisanalytics:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationArgs struct {
	Id string `pulumi:"id"`
}

type LookupApplicationResult struct {
	ApplicationCode        *string                `pulumi:"applicationCode"`
	ApplicationDescription *string                `pulumi:"applicationDescription"`
	Id                     *string                `pulumi:"id"`
	Inputs                 []ApplicationInputType `pulumi:"inputs"`
}

func LookupApplicationOutput(ctx *pulumi.Context, args LookupApplicationOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationResult, error) {
			args := v.(LookupApplicationArgs)
			r, err := LookupApplication(ctx, &args, opts...)
			return *r, err
		}).(LookupApplicationResultOutput)
}

type LookupApplicationOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupApplicationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationArgs)(nil)).Elem()
}

type LookupApplicationResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationResult)(nil)).Elem()
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutput() LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutputWithContext(ctx context.Context) LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) ApplicationCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.ApplicationCode }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) ApplicationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.ApplicationDescription }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) Inputs() ApplicationInputTypeArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []ApplicationInputType { return v.Inputs }).(ApplicationInputTypeArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationResultOutput{})
}