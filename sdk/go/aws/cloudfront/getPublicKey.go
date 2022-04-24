// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::CloudFront::PublicKey
func LookupPublicKey(ctx *pulumi.Context, args *LookupPublicKeyArgs, opts ...pulumi.InvokeOption) (*LookupPublicKeyResult, error) {
	var rv LookupPublicKeyResult
	err := ctx.Invoke("aws-native:cloudfront:getPublicKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPublicKeyArgs struct {
	Id string `pulumi:"id"`
}

type LookupPublicKeyResult struct {
	CreatedTime     *string          `pulumi:"createdTime"`
	Id              *string          `pulumi:"id"`
	PublicKeyConfig *PublicKeyConfig `pulumi:"publicKeyConfig"`
}

func LookupPublicKeyOutput(ctx *pulumi.Context, args LookupPublicKeyOutputArgs, opts ...pulumi.InvokeOption) LookupPublicKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPublicKeyResult, error) {
			args := v.(LookupPublicKeyArgs)
			r, err := LookupPublicKey(ctx, &args, opts...)
			var s LookupPublicKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPublicKeyResultOutput)
}

type LookupPublicKeyOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupPublicKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicKeyArgs)(nil)).Elem()
}

type LookupPublicKeyResultOutput struct{ *pulumi.OutputState }

func (LookupPublicKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicKeyResult)(nil)).Elem()
}

func (o LookupPublicKeyResultOutput) ToLookupPublicKeyResultOutput() LookupPublicKeyResultOutput {
	return o
}

func (o LookupPublicKeyResultOutput) ToLookupPublicKeyResultOutputWithContext(ctx context.Context) LookupPublicKeyResultOutput {
	return o
}

func (o LookupPublicKeyResultOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicKeyResult) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

func (o LookupPublicKeyResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicKeyResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupPublicKeyResultOutput) PublicKeyConfig() PublicKeyConfigPtrOutput {
	return o.ApplyT(func(v LookupPublicKeyResult) *PublicKeyConfig { return v.PublicKeyConfig }).(PublicKeyConfigPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPublicKeyResultOutput{})
}
