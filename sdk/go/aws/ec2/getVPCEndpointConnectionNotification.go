// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::VPCEndpointConnectionNotification
func LookupVPCEndpointConnectionNotification(ctx *pulumi.Context, args *LookupVPCEndpointConnectionNotificationArgs, opts ...pulumi.InvokeOption) (*LookupVPCEndpointConnectionNotificationResult, error) {
	var rv LookupVPCEndpointConnectionNotificationResult
	err := ctx.Invoke("aws-native:ec2:getVPCEndpointConnectionNotification", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVPCEndpointConnectionNotificationArgs struct {
	Id string `pulumi:"id"`
}

type LookupVPCEndpointConnectionNotificationResult struct {
	ConnectionEvents          []string `pulumi:"connectionEvents"`
	ConnectionNotificationArn *string  `pulumi:"connectionNotificationArn"`
	Id                        *string  `pulumi:"id"`
}

func LookupVPCEndpointConnectionNotificationOutput(ctx *pulumi.Context, args LookupVPCEndpointConnectionNotificationOutputArgs, opts ...pulumi.InvokeOption) LookupVPCEndpointConnectionNotificationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVPCEndpointConnectionNotificationResult, error) {
			args := v.(LookupVPCEndpointConnectionNotificationArgs)
			r, err := LookupVPCEndpointConnectionNotification(ctx, &args, opts...)
			return *r, err
		}).(LookupVPCEndpointConnectionNotificationResultOutput)
}

type LookupVPCEndpointConnectionNotificationOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupVPCEndpointConnectionNotificationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVPCEndpointConnectionNotificationArgs)(nil)).Elem()
}

type LookupVPCEndpointConnectionNotificationResultOutput struct{ *pulumi.OutputState }

func (LookupVPCEndpointConnectionNotificationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVPCEndpointConnectionNotificationResult)(nil)).Elem()
}

func (o LookupVPCEndpointConnectionNotificationResultOutput) ToLookupVPCEndpointConnectionNotificationResultOutput() LookupVPCEndpointConnectionNotificationResultOutput {
	return o
}

func (o LookupVPCEndpointConnectionNotificationResultOutput) ToLookupVPCEndpointConnectionNotificationResultOutputWithContext(ctx context.Context) LookupVPCEndpointConnectionNotificationResultOutput {
	return o
}

func (o LookupVPCEndpointConnectionNotificationResultOutput) ConnectionEvents() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVPCEndpointConnectionNotificationResult) []string { return v.ConnectionEvents }).(pulumi.StringArrayOutput)
}

func (o LookupVPCEndpointConnectionNotificationResultOutput) ConnectionNotificationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVPCEndpointConnectionNotificationResult) *string { return v.ConnectionNotificationArn }).(pulumi.StringPtrOutput)
}

func (o LookupVPCEndpointConnectionNotificationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVPCEndpointConnectionNotificationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVPCEndpointConnectionNotificationResultOutput{})
}