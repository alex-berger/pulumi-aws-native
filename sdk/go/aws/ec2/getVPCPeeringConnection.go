// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::VPCPeeringConnection
func LookupVPCPeeringConnection(ctx *pulumi.Context, args *LookupVPCPeeringConnectionArgs, opts ...pulumi.InvokeOption) (*LookupVPCPeeringConnectionResult, error) {
	var rv LookupVPCPeeringConnectionResult
	err := ctx.Invoke("aws-native:ec2:getVPCPeeringConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVPCPeeringConnectionArgs struct {
	Id string `pulumi:"id"`
}

type LookupVPCPeeringConnectionResult struct {
	Id   *string                   `pulumi:"id"`
	Tags []VPCPeeringConnectionTag `pulumi:"tags"`
}

func LookupVPCPeeringConnectionOutput(ctx *pulumi.Context, args LookupVPCPeeringConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupVPCPeeringConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVPCPeeringConnectionResult, error) {
			args := v.(LookupVPCPeeringConnectionArgs)
			r, err := LookupVPCPeeringConnection(ctx, &args, opts...)
			var s LookupVPCPeeringConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVPCPeeringConnectionResultOutput)
}

type LookupVPCPeeringConnectionOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupVPCPeeringConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVPCPeeringConnectionArgs)(nil)).Elem()
}

type LookupVPCPeeringConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupVPCPeeringConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVPCPeeringConnectionResult)(nil)).Elem()
}

func (o LookupVPCPeeringConnectionResultOutput) ToLookupVPCPeeringConnectionResultOutput() LookupVPCPeeringConnectionResultOutput {
	return o
}

func (o LookupVPCPeeringConnectionResultOutput) ToLookupVPCPeeringConnectionResultOutputWithContext(ctx context.Context) LookupVPCPeeringConnectionResultOutput {
	return o
}

func (o LookupVPCPeeringConnectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVPCPeeringConnectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVPCPeeringConnectionResultOutput) Tags() VPCPeeringConnectionTagArrayOutput {
	return o.ApplyT(func(v LookupVPCPeeringConnectionResult) []VPCPeeringConnectionTag { return v.Tags }).(VPCPeeringConnectionTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVPCPeeringConnectionResultOutput{})
}
