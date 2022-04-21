// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IoT::Thing
func LookupThing(ctx *pulumi.Context, args *LookupThingArgs, opts ...pulumi.InvokeOption) (*LookupThingResult, error) {
	var rv LookupThingResult
	err := ctx.Invoke("aws-native:iot:getThing", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupThingArgs struct {
	Id string `pulumi:"id"`
}

type LookupThingResult struct {
	AttributePayload *ThingAttributePayload `pulumi:"attributePayload"`
	Id               *string                `pulumi:"id"`
}

func LookupThingOutput(ctx *pulumi.Context, args LookupThingOutputArgs, opts ...pulumi.InvokeOption) LookupThingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupThingResult, error) {
			args := v.(LookupThingArgs)
			r, err := LookupThing(ctx, &args, opts...)
			var s LookupThingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupThingResultOutput)
}

type LookupThingOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupThingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupThingArgs)(nil)).Elem()
}

type LookupThingResultOutput struct{ *pulumi.OutputState }

func (LookupThingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupThingResult)(nil)).Elem()
}

func (o LookupThingResultOutput) ToLookupThingResultOutput() LookupThingResultOutput {
	return o
}

func (o LookupThingResultOutput) ToLookupThingResultOutputWithContext(ctx context.Context) LookupThingResultOutput {
	return o
}

func (o LookupThingResultOutput) AttributePayload() ThingAttributePayloadPtrOutput {
	return o.ApplyT(func(v LookupThingResult) *ThingAttributePayload { return v.AttributePayload }).(ThingAttributePayloadPtrOutput)
}

func (o LookupThingResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThingResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupThingResultOutput{})
}
