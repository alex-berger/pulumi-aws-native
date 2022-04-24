// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::SubnetNetworkAclAssociation
func LookupSubnetNetworkAclAssociation(ctx *pulumi.Context, args *LookupSubnetNetworkAclAssociationArgs, opts ...pulumi.InvokeOption) (*LookupSubnetNetworkAclAssociationResult, error) {
	var rv LookupSubnetNetworkAclAssociationResult
	err := ctx.Invoke("aws-native:ec2:getSubnetNetworkAclAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubnetNetworkAclAssociationArgs struct {
	AssociationId string `pulumi:"associationId"`
}

type LookupSubnetNetworkAclAssociationResult struct {
	AssociationId *string `pulumi:"associationId"`
}

func LookupSubnetNetworkAclAssociationOutput(ctx *pulumi.Context, args LookupSubnetNetworkAclAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupSubnetNetworkAclAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubnetNetworkAclAssociationResult, error) {
			args := v.(LookupSubnetNetworkAclAssociationArgs)
			r, err := LookupSubnetNetworkAclAssociation(ctx, &args, opts...)
			var s LookupSubnetNetworkAclAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubnetNetworkAclAssociationResultOutput)
}

type LookupSubnetNetworkAclAssociationOutputArgs struct {
	AssociationId pulumi.StringInput `pulumi:"associationId"`
}

func (LookupSubnetNetworkAclAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubnetNetworkAclAssociationArgs)(nil)).Elem()
}

type LookupSubnetNetworkAclAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupSubnetNetworkAclAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubnetNetworkAclAssociationResult)(nil)).Elem()
}

func (o LookupSubnetNetworkAclAssociationResultOutput) ToLookupSubnetNetworkAclAssociationResultOutput() LookupSubnetNetworkAclAssociationResultOutput {
	return o
}

func (o LookupSubnetNetworkAclAssociationResultOutput) ToLookupSubnetNetworkAclAssociationResultOutputWithContext(ctx context.Context) LookupSubnetNetworkAclAssociationResultOutput {
	return o
}

func (o LookupSubnetNetworkAclAssociationResultOutput) AssociationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubnetNetworkAclAssociationResult) *string { return v.AssociationId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubnetNetworkAclAssociationResultOutput{})
}
