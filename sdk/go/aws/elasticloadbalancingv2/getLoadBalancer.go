// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancingv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ElasticLoadBalancingV2::LoadBalancer
func LookupLoadBalancer(ctx *pulumi.Context, args *LookupLoadBalancerArgs, opts ...pulumi.InvokeOption) (*LookupLoadBalancerResult, error) {
	var rv LookupLoadBalancerResult
	err := ctx.Invoke("aws-native:elasticloadbalancingv2:getLoadBalancer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLoadBalancerArgs struct {
	Id string `pulumi:"id"`
}

type LookupLoadBalancerResult struct {
	CanonicalHostedZoneID  *string                     `pulumi:"canonicalHostedZoneID"`
	DNSName                *string                     `pulumi:"dNSName"`
	Id                     *string                     `pulumi:"id"`
	IpAddressType          *string                     `pulumi:"ipAddressType"`
	LoadBalancerAttributes []LoadBalancerAttribute     `pulumi:"loadBalancerAttributes"`
	LoadBalancerFullName   *string                     `pulumi:"loadBalancerFullName"`
	LoadBalancerName       *string                     `pulumi:"loadBalancerName"`
	SecurityGroups         []string                    `pulumi:"securityGroups"`
	SubnetMappings         []LoadBalancerSubnetMapping `pulumi:"subnetMappings"`
	Subnets                []string                    `pulumi:"subnets"`
	Tags                   []LoadBalancerTag           `pulumi:"tags"`
}

func LookupLoadBalancerOutput(ctx *pulumi.Context, args LookupLoadBalancerOutputArgs, opts ...pulumi.InvokeOption) LookupLoadBalancerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLoadBalancerResult, error) {
			args := v.(LookupLoadBalancerArgs)
			r, err := LookupLoadBalancer(ctx, &args, opts...)
			var s LookupLoadBalancerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLoadBalancerResultOutput)
}

type LookupLoadBalancerOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupLoadBalancerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancerArgs)(nil)).Elem()
}

type LookupLoadBalancerResultOutput struct{ *pulumi.OutputState }

func (LookupLoadBalancerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancerResult)(nil)).Elem()
}

func (o LookupLoadBalancerResultOutput) ToLookupLoadBalancerResultOutput() LookupLoadBalancerResultOutput {
	return o
}

func (o LookupLoadBalancerResultOutput) ToLookupLoadBalancerResultOutputWithContext(ctx context.Context) LookupLoadBalancerResultOutput {
	return o
}

func (o LookupLoadBalancerResultOutput) CanonicalHostedZoneID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.CanonicalHostedZoneID }).(pulumi.StringPtrOutput)
}

func (o LookupLoadBalancerResultOutput) DNSName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.DNSName }).(pulumi.StringPtrOutput)
}

func (o LookupLoadBalancerResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupLoadBalancerResultOutput) IpAddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.IpAddressType }).(pulumi.StringPtrOutput)
}

func (o LookupLoadBalancerResultOutput) LoadBalancerAttributes() LoadBalancerAttributeArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []LoadBalancerAttribute { return v.LoadBalancerAttributes }).(LoadBalancerAttributeArrayOutput)
}

func (o LookupLoadBalancerResultOutput) LoadBalancerFullName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.LoadBalancerFullName }).(pulumi.StringPtrOutput)
}

func (o LookupLoadBalancerResultOutput) LoadBalancerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.LoadBalancerName }).(pulumi.StringPtrOutput)
}

func (o LookupLoadBalancerResultOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []string { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

func (o LookupLoadBalancerResultOutput) SubnetMappings() LoadBalancerSubnetMappingArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []LoadBalancerSubnetMapping { return v.SubnetMappings }).(LoadBalancerSubnetMappingArrayOutput)
}

func (o LookupLoadBalancerResultOutput) Subnets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []string { return v.Subnets }).(pulumi.StringArrayOutput)
}

func (o LookupLoadBalancerResultOutput) Tags() LoadBalancerTagArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []LoadBalancerTag { return v.Tags }).(LoadBalancerTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLoadBalancerResultOutput{})
}
