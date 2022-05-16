// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mediaconnect

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::MediaConnect::FlowEntitlement
func LookupFlowEntitlement(ctx *pulumi.Context, args *LookupFlowEntitlementArgs, opts ...pulumi.InvokeOption) (*LookupFlowEntitlementResult, error) {
	var rv LookupFlowEntitlementResult
	err := ctx.Invoke("aws-native:mediaconnect:getFlowEntitlement", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFlowEntitlementArgs struct {
	// The ARN of the entitlement.
	EntitlementArn string `pulumi:"entitlementArn"`
}

type LookupFlowEntitlementResult struct {
	// A description of the entitlement.
	Description *string `pulumi:"description"`
	// The type of encryption that will be used on the output that is associated with this entitlement.
	Encryption *FlowEntitlementEncryption `pulumi:"encryption"`
	// The ARN of the entitlement.
	EntitlementArn *string `pulumi:"entitlementArn"`
	//  An indication of whether the entitlement is enabled.
	EntitlementStatus *FlowEntitlementEntitlementStatus `pulumi:"entitlementStatus"`
	// The ARN of the flow.
	FlowArn *string `pulumi:"flowArn"`
	// The AWS account IDs that you want to share your content with. The receiving accounts (subscribers) will be allowed to create their own flow using your content as the source.
	Subscribers []string `pulumi:"subscribers"`
}

func LookupFlowEntitlementOutput(ctx *pulumi.Context, args LookupFlowEntitlementOutputArgs, opts ...pulumi.InvokeOption) LookupFlowEntitlementResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFlowEntitlementResult, error) {
			args := v.(LookupFlowEntitlementArgs)
			r, err := LookupFlowEntitlement(ctx, &args, opts...)
			var s LookupFlowEntitlementResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFlowEntitlementResultOutput)
}

type LookupFlowEntitlementOutputArgs struct {
	// The ARN of the entitlement.
	EntitlementArn pulumi.StringInput `pulumi:"entitlementArn"`
}

func (LookupFlowEntitlementOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFlowEntitlementArgs)(nil)).Elem()
}

type LookupFlowEntitlementResultOutput struct{ *pulumi.OutputState }

func (LookupFlowEntitlementResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFlowEntitlementResult)(nil)).Elem()
}

func (o LookupFlowEntitlementResultOutput) ToLookupFlowEntitlementResultOutput() LookupFlowEntitlementResultOutput {
	return o
}

func (o LookupFlowEntitlementResultOutput) ToLookupFlowEntitlementResultOutputWithContext(ctx context.Context) LookupFlowEntitlementResultOutput {
	return o
}

// A description of the entitlement.
func (o LookupFlowEntitlementResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFlowEntitlementResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The type of encryption that will be used on the output that is associated with this entitlement.
func (o LookupFlowEntitlementResultOutput) Encryption() FlowEntitlementEncryptionPtrOutput {
	return o.ApplyT(func(v LookupFlowEntitlementResult) *FlowEntitlementEncryption { return v.Encryption }).(FlowEntitlementEncryptionPtrOutput)
}

// The ARN of the entitlement.
func (o LookupFlowEntitlementResultOutput) EntitlementArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFlowEntitlementResult) *string { return v.EntitlementArn }).(pulumi.StringPtrOutput)
}

//  An indication of whether the entitlement is enabled.
func (o LookupFlowEntitlementResultOutput) EntitlementStatus() FlowEntitlementEntitlementStatusPtrOutput {
	return o.ApplyT(func(v LookupFlowEntitlementResult) *FlowEntitlementEntitlementStatus { return v.EntitlementStatus }).(FlowEntitlementEntitlementStatusPtrOutput)
}

// The ARN of the flow.
func (o LookupFlowEntitlementResultOutput) FlowArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFlowEntitlementResult) *string { return v.FlowArn }).(pulumi.StringPtrOutput)
}

// The AWS account IDs that you want to share your content with. The receiving accounts (subscribers) will be allowed to create their own flow using your content as the source.
func (o LookupFlowEntitlementResultOutput) Subscribers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFlowEntitlementResult) []string { return v.Subscribers }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFlowEntitlementResultOutput{})
}