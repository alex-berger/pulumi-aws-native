// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::NetworkAcl
func LookupNetworkAcl(ctx *pulumi.Context, args *LookupNetworkAclArgs, opts ...pulumi.InvokeOption) (*LookupNetworkAclResult, error) {
	var rv LookupNetworkAclResult
	err := ctx.Invoke("aws-native:ec2:getNetworkAcl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkAclArgs struct {
	Id string `pulumi:"id"`
}

type LookupNetworkAclResult struct {
	Id *string `pulumi:"id"`
	// The tags to assign to the network ACL.
	Tags []NetworkAclTag `pulumi:"tags"`
}

func LookupNetworkAclOutput(ctx *pulumi.Context, args LookupNetworkAclOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkAclResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkAclResult, error) {
			args := v.(LookupNetworkAclArgs)
			r, err := LookupNetworkAcl(ctx, &args, opts...)
			var s LookupNetworkAclResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkAclResultOutput)
}

type LookupNetworkAclOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupNetworkAclOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkAclArgs)(nil)).Elem()
}

type LookupNetworkAclResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkAclResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkAclResult)(nil)).Elem()
}

func (o LookupNetworkAclResultOutput) ToLookupNetworkAclResultOutput() LookupNetworkAclResultOutput {
	return o
}

func (o LookupNetworkAclResultOutput) ToLookupNetworkAclResultOutputWithContext(ctx context.Context) LookupNetworkAclResultOutput {
	return o
}

func (o LookupNetworkAclResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkAclResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The tags to assign to the network ACL.
func (o LookupNetworkAclResultOutput) Tags() NetworkAclTagArrayOutput {
	return o.ApplyT(func(v LookupNetworkAclResult) []NetworkAclTag { return v.Tags }).(NetworkAclTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkAclResultOutput{})
}
