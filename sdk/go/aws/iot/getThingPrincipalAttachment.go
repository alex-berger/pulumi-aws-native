// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IoT::ThingPrincipalAttachment
func LookupThingPrincipalAttachment(ctx *pulumi.Context, args *LookupThingPrincipalAttachmentArgs, opts ...pulumi.InvokeOption) (*LookupThingPrincipalAttachmentResult, error) {
	var rv LookupThingPrincipalAttachmentResult
	err := ctx.Invoke("aws-native:iot:getThingPrincipalAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupThingPrincipalAttachmentArgs struct {
	Id string `pulumi:"id"`
}

type LookupThingPrincipalAttachmentResult struct {
	Id *string `pulumi:"id"`
}

func LookupThingPrincipalAttachmentOutput(ctx *pulumi.Context, args LookupThingPrincipalAttachmentOutputArgs, opts ...pulumi.InvokeOption) LookupThingPrincipalAttachmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupThingPrincipalAttachmentResult, error) {
			args := v.(LookupThingPrincipalAttachmentArgs)
			r, err := LookupThingPrincipalAttachment(ctx, &args, opts...)
			var s LookupThingPrincipalAttachmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupThingPrincipalAttachmentResultOutput)
}

type LookupThingPrincipalAttachmentOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupThingPrincipalAttachmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupThingPrincipalAttachmentArgs)(nil)).Elem()
}

type LookupThingPrincipalAttachmentResultOutput struct{ *pulumi.OutputState }

func (LookupThingPrincipalAttachmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupThingPrincipalAttachmentResult)(nil)).Elem()
}

func (o LookupThingPrincipalAttachmentResultOutput) ToLookupThingPrincipalAttachmentResultOutput() LookupThingPrincipalAttachmentResultOutput {
	return o
}

func (o LookupThingPrincipalAttachmentResultOutput) ToLookupThingPrincipalAttachmentResultOutputWithContext(ctx context.Context) LookupThingPrincipalAttachmentResultOutput {
	return o
}

func (o LookupThingPrincipalAttachmentResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThingPrincipalAttachmentResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupThingPrincipalAttachmentResultOutput{})
}
