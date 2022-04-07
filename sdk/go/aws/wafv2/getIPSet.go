// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Contains a list of IP addresses. This can be either IPV4 or IPV6. The list will be mutually
func LookupIPSet(ctx *pulumi.Context, args *LookupIPSetArgs, opts ...pulumi.InvokeOption) (*LookupIPSetResult, error) {
	var rv LookupIPSetResult
	err := ctx.Invoke("aws-native:wafv2:getIPSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIPSetArgs struct {
	Id    string     `pulumi:"id"`
	Name  string     `pulumi:"name"`
	Scope IPSetScope `pulumi:"scope"`
}

type LookupIPSetResult struct {
	// List of IPAddresses.
	Addresses        []string               `pulumi:"addresses"`
	Arn              *string                `pulumi:"arn"`
	Description      *string                `pulumi:"description"`
	IPAddressVersion *IPSetIPAddressVersion `pulumi:"iPAddressVersion"`
	Id               *string                `pulumi:"id"`
	Tags             []IPSetTag             `pulumi:"tags"`
}

func LookupIPSetOutput(ctx *pulumi.Context, args LookupIPSetOutputArgs, opts ...pulumi.InvokeOption) LookupIPSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIPSetResult, error) {
			args := v.(LookupIPSetArgs)
			r, err := LookupIPSet(ctx, &args, opts...)
			var s LookupIPSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIPSetResultOutput)
}

type LookupIPSetOutputArgs struct {
	Id    pulumi.StringInput `pulumi:"id"`
	Name  pulumi.StringInput `pulumi:"name"`
	Scope IPSetScopeInput    `pulumi:"scope"`
}

func (LookupIPSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIPSetArgs)(nil)).Elem()
}

type LookupIPSetResultOutput struct{ *pulumi.OutputState }

func (LookupIPSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIPSetResult)(nil)).Elem()
}

func (o LookupIPSetResultOutput) ToLookupIPSetResultOutput() LookupIPSetResultOutput {
	return o
}

func (o LookupIPSetResultOutput) ToLookupIPSetResultOutputWithContext(ctx context.Context) LookupIPSetResultOutput {
	return o
}

// List of IPAddresses.
func (o LookupIPSetResultOutput) Addresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIPSetResult) []string { return v.Addresses }).(pulumi.StringArrayOutput)
}

func (o LookupIPSetResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPSetResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupIPSetResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPSetResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupIPSetResultOutput) IPAddressVersion() IPSetIPAddressVersionPtrOutput {
	return o.ApplyT(func(v LookupIPSetResult) *IPSetIPAddressVersion { return v.IPAddressVersion }).(IPSetIPAddressVersionPtrOutput)
}

func (o LookupIPSetResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPSetResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupIPSetResultOutput) Tags() IPSetTagArrayOutput {
	return o.ApplyT(func(v LookupIPSetResult) []IPSetTag { return v.Tags }).(IPSetTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIPSetResultOutput{})
}
