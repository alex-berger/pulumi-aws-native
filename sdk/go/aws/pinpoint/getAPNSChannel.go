// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Pinpoint::APNSChannel
func LookupAPNSChannel(ctx *pulumi.Context, args *LookupAPNSChannelArgs, opts ...pulumi.InvokeOption) (*LookupAPNSChannelResult, error) {
	var rv LookupAPNSChannelResult
	err := ctx.Invoke("aws-native:pinpoint:getAPNSChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAPNSChannelArgs struct {
	Id string `pulumi:"id"`
}

type LookupAPNSChannelResult struct {
	BundleId                    *string `pulumi:"bundleId"`
	Certificate                 *string `pulumi:"certificate"`
	DefaultAuthenticationMethod *string `pulumi:"defaultAuthenticationMethod"`
	Enabled                     *bool   `pulumi:"enabled"`
	Id                          *string `pulumi:"id"`
	PrivateKey                  *string `pulumi:"privateKey"`
	TeamId                      *string `pulumi:"teamId"`
	TokenKey                    *string `pulumi:"tokenKey"`
	TokenKeyId                  *string `pulumi:"tokenKeyId"`
}

func LookupAPNSChannelOutput(ctx *pulumi.Context, args LookupAPNSChannelOutputArgs, opts ...pulumi.InvokeOption) LookupAPNSChannelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAPNSChannelResult, error) {
			args := v.(LookupAPNSChannelArgs)
			r, err := LookupAPNSChannel(ctx, &args, opts...)
			var s LookupAPNSChannelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAPNSChannelResultOutput)
}

type LookupAPNSChannelOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupAPNSChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAPNSChannelArgs)(nil)).Elem()
}

type LookupAPNSChannelResultOutput struct{ *pulumi.OutputState }

func (LookupAPNSChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAPNSChannelResult)(nil)).Elem()
}

func (o LookupAPNSChannelResultOutput) ToLookupAPNSChannelResultOutput() LookupAPNSChannelResultOutput {
	return o
}

func (o LookupAPNSChannelResultOutput) ToLookupAPNSChannelResultOutputWithContext(ctx context.Context) LookupAPNSChannelResultOutput {
	return o
}

func (o LookupAPNSChannelResultOutput) BundleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAPNSChannelResult) *string { return v.BundleId }).(pulumi.StringPtrOutput)
}

func (o LookupAPNSChannelResultOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAPNSChannelResult) *string { return v.Certificate }).(pulumi.StringPtrOutput)
}

func (o LookupAPNSChannelResultOutput) DefaultAuthenticationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAPNSChannelResult) *string { return v.DefaultAuthenticationMethod }).(pulumi.StringPtrOutput)
}

func (o LookupAPNSChannelResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAPNSChannelResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupAPNSChannelResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAPNSChannelResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupAPNSChannelResultOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAPNSChannelResult) *string { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

func (o LookupAPNSChannelResultOutput) TeamId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAPNSChannelResult) *string { return v.TeamId }).(pulumi.StringPtrOutput)
}

func (o LookupAPNSChannelResultOutput) TokenKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAPNSChannelResult) *string { return v.TokenKey }).(pulumi.StringPtrOutput)
}

func (o LookupAPNSChannelResultOutput) TokenKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAPNSChannelResult) *string { return v.TokenKeyId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAPNSChannelResultOutput{})
}