// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::VPNGateway
func LookupVPNGateway(ctx *pulumi.Context, args *LookupVPNGatewayArgs, opts ...pulumi.InvokeOption) (*LookupVPNGatewayResult, error) {
	var rv LookupVPNGatewayResult
	err := ctx.Invoke("aws-native:ec2:getVPNGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVPNGatewayArgs struct {
	Id string `pulumi:"id"`
}

type LookupVPNGatewayResult struct {
	Id   *string         `pulumi:"id"`
	Tags []VPNGatewayTag `pulumi:"tags"`
}

func LookupVPNGatewayOutput(ctx *pulumi.Context, args LookupVPNGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupVPNGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVPNGatewayResult, error) {
			args := v.(LookupVPNGatewayArgs)
			r, err := LookupVPNGateway(ctx, &args, opts...)
			var s LookupVPNGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVPNGatewayResultOutput)
}

type LookupVPNGatewayOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupVPNGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVPNGatewayArgs)(nil)).Elem()
}

type LookupVPNGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupVPNGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVPNGatewayResult)(nil)).Elem()
}

func (o LookupVPNGatewayResultOutput) ToLookupVPNGatewayResultOutput() LookupVPNGatewayResultOutput {
	return o
}

func (o LookupVPNGatewayResultOutput) ToLookupVPNGatewayResultOutputWithContext(ctx context.Context) LookupVPNGatewayResultOutput {
	return o
}

func (o LookupVPNGatewayResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVPNGatewayResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVPNGatewayResultOutput) Tags() VPNGatewayTagArrayOutput {
	return o.ApplyT(func(v LookupVPNGatewayResult) []VPNGatewayTag { return v.Tags }).(VPNGatewayTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVPNGatewayResultOutput{})
}
