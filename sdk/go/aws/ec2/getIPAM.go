// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Schema of AWS::EC2::IPAM Type
func LookupIPAM(ctx *pulumi.Context, args *LookupIPAMArgs, opts ...pulumi.InvokeOption) (*LookupIPAMResult, error) {
	var rv LookupIPAMResult
	err := ctx.Invoke("aws-native:ec2:getIPAM", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIPAMArgs struct {
	// Id of the IPAM.
	IpamId string `pulumi:"ipamId"`
}

type LookupIPAMResult struct {
	// The Amazon Resource Name (ARN) of the IPAM.
	Arn         *string `pulumi:"arn"`
	Description *string `pulumi:"description"`
	// Id of the IPAM.
	IpamId *string `pulumi:"ipamId"`
	// The regions IPAM is enabled for. Allows pools to be created in these regions, as well as enabling monitoring
	OperatingRegions []IPAMIpamOperatingRegion `pulumi:"operatingRegions"`
	// The Id of the default scope for publicly routable IP space, created with this IPAM.
	PrivateDefaultScopeId *string `pulumi:"privateDefaultScopeId"`
	// The Id of the default scope for publicly routable IP space, created with this IPAM.
	PublicDefaultScopeId *string `pulumi:"publicDefaultScopeId"`
	// The number of scopes that currently exist in this IPAM.
	ScopeCount *int `pulumi:"scopeCount"`
	// An array of key-value pairs to apply to this resource.
	Tags []IPAMTag `pulumi:"tags"`
}

func LookupIPAMOutput(ctx *pulumi.Context, args LookupIPAMOutputArgs, opts ...pulumi.InvokeOption) LookupIPAMResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIPAMResult, error) {
			args := v.(LookupIPAMArgs)
			r, err := LookupIPAM(ctx, &args, opts...)
			var s LookupIPAMResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIPAMResultOutput)
}

type LookupIPAMOutputArgs struct {
	// Id of the IPAM.
	IpamId pulumi.StringInput `pulumi:"ipamId"`
}

func (LookupIPAMOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIPAMArgs)(nil)).Elem()
}

type LookupIPAMResultOutput struct{ *pulumi.OutputState }

func (LookupIPAMResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIPAMResult)(nil)).Elem()
}

func (o LookupIPAMResultOutput) ToLookupIPAMResultOutput() LookupIPAMResultOutput {
	return o
}

func (o LookupIPAMResultOutput) ToLookupIPAMResultOutputWithContext(ctx context.Context) LookupIPAMResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the IPAM.
func (o LookupIPAMResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPAMResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupIPAMResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPAMResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Id of the IPAM.
func (o LookupIPAMResultOutput) IpamId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPAMResult) *string { return v.IpamId }).(pulumi.StringPtrOutput)
}

// The regions IPAM is enabled for. Allows pools to be created in these regions, as well as enabling monitoring
func (o LookupIPAMResultOutput) OperatingRegions() IPAMIpamOperatingRegionArrayOutput {
	return o.ApplyT(func(v LookupIPAMResult) []IPAMIpamOperatingRegion { return v.OperatingRegions }).(IPAMIpamOperatingRegionArrayOutput)
}

// The Id of the default scope for publicly routable IP space, created with this IPAM.
func (o LookupIPAMResultOutput) PrivateDefaultScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPAMResult) *string { return v.PrivateDefaultScopeId }).(pulumi.StringPtrOutput)
}

// The Id of the default scope for publicly routable IP space, created with this IPAM.
func (o LookupIPAMResultOutput) PublicDefaultScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPAMResult) *string { return v.PublicDefaultScopeId }).(pulumi.StringPtrOutput)
}

// The number of scopes that currently exist in this IPAM.
func (o LookupIPAMResultOutput) ScopeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupIPAMResult) *int { return v.ScopeCount }).(pulumi.IntPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupIPAMResultOutput) Tags() IPAMTagArrayOutput {
	return o.ApplyT(func(v LookupIPAMResult) []IPAMTag { return v.Tags }).(IPAMTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIPAMResultOutput{})
}
