// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot1click

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IoT1Click::Placement
func LookupPlacement(ctx *pulumi.Context, args *LookupPlacementArgs, opts ...pulumi.InvokeOption) (*LookupPlacementResult, error) {
	var rv LookupPlacementResult
	err := ctx.Invoke("aws-native:iot1click:getPlacement", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPlacementArgs struct {
	Id string `pulumi:"id"`
}

type LookupPlacementResult struct {
	Attributes interface{} `pulumi:"attributes"`
	Id         *string     `pulumi:"id"`
}

func LookupPlacementOutput(ctx *pulumi.Context, args LookupPlacementOutputArgs, opts ...pulumi.InvokeOption) LookupPlacementResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPlacementResult, error) {
			args := v.(LookupPlacementArgs)
			r, err := LookupPlacement(ctx, &args, opts...)
			var s LookupPlacementResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPlacementResultOutput)
}

type LookupPlacementOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupPlacementOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlacementArgs)(nil)).Elem()
}

type LookupPlacementResultOutput struct{ *pulumi.OutputState }

func (LookupPlacementResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlacementResult)(nil)).Elem()
}

func (o LookupPlacementResultOutput) ToLookupPlacementResultOutput() LookupPlacementResultOutput {
	return o
}

func (o LookupPlacementResultOutput) ToLookupPlacementResultOutputWithContext(ctx context.Context) LookupPlacementResultOutput {
	return o
}

func (o LookupPlacementResultOutput) Attributes() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPlacementResult) interface{} { return v.Attributes }).(pulumi.AnyOutput)
}

func (o LookupPlacementResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPlacementResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPlacementResultOutput{})
}
