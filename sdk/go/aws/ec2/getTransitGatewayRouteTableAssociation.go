// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::TransitGatewayRouteTableAssociation
func LookupTransitGatewayRouteTableAssociation(ctx *pulumi.Context, args *LookupTransitGatewayRouteTableAssociationArgs, opts ...pulumi.InvokeOption) (*LookupTransitGatewayRouteTableAssociationResult, error) {
	var rv LookupTransitGatewayRouteTableAssociationResult
	err := ctx.Invoke("aws-native:ec2:getTransitGatewayRouteTableAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTransitGatewayRouteTableAssociationArgs struct {
	Id string `pulumi:"id"`
}

type LookupTransitGatewayRouteTableAssociationResult struct {
	Id *string `pulumi:"id"`
}

func LookupTransitGatewayRouteTableAssociationOutput(ctx *pulumi.Context, args LookupTransitGatewayRouteTableAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupTransitGatewayRouteTableAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTransitGatewayRouteTableAssociationResult, error) {
			args := v.(LookupTransitGatewayRouteTableAssociationArgs)
			r, err := LookupTransitGatewayRouteTableAssociation(ctx, &args, opts...)
			var s LookupTransitGatewayRouteTableAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTransitGatewayRouteTableAssociationResultOutput)
}

type LookupTransitGatewayRouteTableAssociationOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupTransitGatewayRouteTableAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTransitGatewayRouteTableAssociationArgs)(nil)).Elem()
}

type LookupTransitGatewayRouteTableAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupTransitGatewayRouteTableAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTransitGatewayRouteTableAssociationResult)(nil)).Elem()
}

func (o LookupTransitGatewayRouteTableAssociationResultOutput) ToLookupTransitGatewayRouteTableAssociationResultOutput() LookupTransitGatewayRouteTableAssociationResultOutput {
	return o
}

func (o LookupTransitGatewayRouteTableAssociationResultOutput) ToLookupTransitGatewayRouteTableAssociationResultOutputWithContext(ctx context.Context) LookupTransitGatewayRouteTableAssociationResultOutput {
	return o
}

func (o LookupTransitGatewayRouteTableAssociationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTransitGatewayRouteTableAssociationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTransitGatewayRouteTableAssociationResultOutput{})
}
