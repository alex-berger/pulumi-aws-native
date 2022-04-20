// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::VolumeAttachment
func LookupVolumeAttachment(ctx *pulumi.Context, args *LookupVolumeAttachmentArgs, opts ...pulumi.InvokeOption) (*LookupVolumeAttachmentResult, error) {
	var rv LookupVolumeAttachmentResult
	err := ctx.Invoke("aws-native:ec2:getVolumeAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVolumeAttachmentArgs struct {
	Id string `pulumi:"id"`
}

type LookupVolumeAttachmentResult struct {
	Id *string `pulumi:"id"`
}

func LookupVolumeAttachmentOutput(ctx *pulumi.Context, args LookupVolumeAttachmentOutputArgs, opts ...pulumi.InvokeOption) LookupVolumeAttachmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVolumeAttachmentResult, error) {
			args := v.(LookupVolumeAttachmentArgs)
			r, err := LookupVolumeAttachment(ctx, &args, opts...)
			var s LookupVolumeAttachmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVolumeAttachmentResultOutput)
}

type LookupVolumeAttachmentOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupVolumeAttachmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeAttachmentArgs)(nil)).Elem()
}

type LookupVolumeAttachmentResultOutput struct{ *pulumi.OutputState }

func (LookupVolumeAttachmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeAttachmentResult)(nil)).Elem()
}

func (o LookupVolumeAttachmentResultOutput) ToLookupVolumeAttachmentResultOutput() LookupVolumeAttachmentResultOutput {
	return o
}

func (o LookupVolumeAttachmentResultOutput) ToLookupVolumeAttachmentResultOutputWithContext(ctx context.Context) LookupVolumeAttachmentResultOutput {
	return o
}

func (o LookupVolumeAttachmentResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeAttachmentResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeAttachmentResultOutput{})
}
