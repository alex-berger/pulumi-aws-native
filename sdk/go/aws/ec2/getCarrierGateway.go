// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An example resource schema demonstrating some basic constructs and validation rules.
func LookupCarrierGateway(ctx *pulumi.Context, args *LookupCarrierGatewayArgs, opts ...pulumi.InvokeOption) (*LookupCarrierGatewayResult, error) {
	var rv LookupCarrierGatewayResult
	err := ctx.Invoke("aws-native:ec2:getCarrierGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCarrierGatewayArgs struct {
	// The ID of the carrier gateway.
	CarrierGatewayId string `pulumi:"carrierGatewayId"`
}

type LookupCarrierGatewayResult struct {
	// The ID of the carrier gateway.
	CarrierGatewayId *string `pulumi:"carrierGatewayId"`
	// The ID of the owner.
	OwnerId *string `pulumi:"ownerId"`
	// The state of the carrier gateway.
	State *string `pulumi:"state"`
	// The tags for the carrier gateway.
	Tags []CarrierGatewayTag `pulumi:"tags"`
}

func LookupCarrierGatewayOutput(ctx *pulumi.Context, args LookupCarrierGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupCarrierGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCarrierGatewayResult, error) {
			args := v.(LookupCarrierGatewayArgs)
			r, err := LookupCarrierGateway(ctx, &args, opts...)
			var s LookupCarrierGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCarrierGatewayResultOutput)
}

type LookupCarrierGatewayOutputArgs struct {
	// The ID of the carrier gateway.
	CarrierGatewayId pulumi.StringInput `pulumi:"carrierGatewayId"`
}

func (LookupCarrierGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCarrierGatewayArgs)(nil)).Elem()
}

type LookupCarrierGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupCarrierGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCarrierGatewayResult)(nil)).Elem()
}

func (o LookupCarrierGatewayResultOutput) ToLookupCarrierGatewayResultOutput() LookupCarrierGatewayResultOutput {
	return o
}

func (o LookupCarrierGatewayResultOutput) ToLookupCarrierGatewayResultOutputWithContext(ctx context.Context) LookupCarrierGatewayResultOutput {
	return o
}

// The ID of the carrier gateway.
func (o LookupCarrierGatewayResultOutput) CarrierGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCarrierGatewayResult) *string { return v.CarrierGatewayId }).(pulumi.StringPtrOutput)
}

// The ID of the owner.
func (o LookupCarrierGatewayResultOutput) OwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCarrierGatewayResult) *string { return v.OwnerId }).(pulumi.StringPtrOutput)
}

// The state of the carrier gateway.
func (o LookupCarrierGatewayResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCarrierGatewayResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// The tags for the carrier gateway.
func (o LookupCarrierGatewayResultOutput) Tags() CarrierGatewayTagArrayOutput {
	return o.ApplyT(func(v LookupCarrierGatewayResult) []CarrierGatewayTag { return v.Tags }).(CarrierGatewayTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCarrierGatewayResultOutput{})
}
