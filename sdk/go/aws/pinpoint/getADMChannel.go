// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Pinpoint::ADMChannel
func LookupADMChannel(ctx *pulumi.Context, args *LookupADMChannelArgs, opts ...pulumi.InvokeOption) (*LookupADMChannelResult, error) {
	var rv LookupADMChannelResult
	err := ctx.Invoke("aws-native:pinpoint:getADMChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADMChannelArgs struct {
	Id string `pulumi:"id"`
}

type LookupADMChannelResult struct {
	ClientId     *string `pulumi:"clientId"`
	ClientSecret *string `pulumi:"clientSecret"`
	Enabled      *bool   `pulumi:"enabled"`
	Id           *string `pulumi:"id"`
}

func LookupADMChannelOutput(ctx *pulumi.Context, args LookupADMChannelOutputArgs, opts ...pulumi.InvokeOption) LookupADMChannelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupADMChannelResult, error) {
			args := v.(LookupADMChannelArgs)
			r, err := LookupADMChannel(ctx, &args, opts...)
			var s LookupADMChannelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupADMChannelResultOutput)
}

type LookupADMChannelOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupADMChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADMChannelArgs)(nil)).Elem()
}

type LookupADMChannelResultOutput struct{ *pulumi.OutputState }

func (LookupADMChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADMChannelResult)(nil)).Elem()
}

func (o LookupADMChannelResultOutput) ToLookupADMChannelResultOutput() LookupADMChannelResultOutput {
	return o
}

func (o LookupADMChannelResultOutput) ToLookupADMChannelResultOutputWithContext(ctx context.Context) LookupADMChannelResultOutput {
	return o
}

func (o LookupADMChannelResultOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupADMChannelResult) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o LookupADMChannelResultOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupADMChannelResult) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o LookupADMChannelResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupADMChannelResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupADMChannelResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupADMChannelResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupADMChannelResultOutput{})
}
