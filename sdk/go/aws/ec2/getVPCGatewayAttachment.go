// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::VPCGatewayAttachment
func LookupVPCGatewayAttachment(ctx *pulumi.Context, args *LookupVPCGatewayAttachmentArgs, opts ...pulumi.InvokeOption) (*LookupVPCGatewayAttachmentResult, error) {
	var rv LookupVPCGatewayAttachmentResult
	err := ctx.Invoke("aws-native:ec2:getVPCGatewayAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVPCGatewayAttachmentArgs struct {
	Id string `pulumi:"id"`
}

type LookupVPCGatewayAttachmentResult struct {
	Id                *string `pulumi:"id"`
	InternetGatewayId *string `pulumi:"internetGatewayId"`
	VpcId             *string `pulumi:"vpcId"`
	VpnGatewayId      *string `pulumi:"vpnGatewayId"`
}

func LookupVPCGatewayAttachmentOutput(ctx *pulumi.Context, args LookupVPCGatewayAttachmentOutputArgs, opts ...pulumi.InvokeOption) LookupVPCGatewayAttachmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVPCGatewayAttachmentResult, error) {
			args := v.(LookupVPCGatewayAttachmentArgs)
			r, err := LookupVPCGatewayAttachment(ctx, &args, opts...)
			return *r, err
		}).(LookupVPCGatewayAttachmentResultOutput)
}

type LookupVPCGatewayAttachmentOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupVPCGatewayAttachmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVPCGatewayAttachmentArgs)(nil)).Elem()
}

type LookupVPCGatewayAttachmentResultOutput struct{ *pulumi.OutputState }

func (LookupVPCGatewayAttachmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVPCGatewayAttachmentResult)(nil)).Elem()
}

func (o LookupVPCGatewayAttachmentResultOutput) ToLookupVPCGatewayAttachmentResultOutput() LookupVPCGatewayAttachmentResultOutput {
	return o
}

func (o LookupVPCGatewayAttachmentResultOutput) ToLookupVPCGatewayAttachmentResultOutputWithContext(ctx context.Context) LookupVPCGatewayAttachmentResultOutput {
	return o
}

func (o LookupVPCGatewayAttachmentResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVPCGatewayAttachmentResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVPCGatewayAttachmentResultOutput) InternetGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVPCGatewayAttachmentResult) *string { return v.InternetGatewayId }).(pulumi.StringPtrOutput)
}

func (o LookupVPCGatewayAttachmentResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVPCGatewayAttachmentResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func (o LookupVPCGatewayAttachmentResultOutput) VpnGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVPCGatewayAttachmentResult) *string { return v.VpnGatewayId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVPCGatewayAttachmentResultOutput{})
}