// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::WAFRegional::GeoMatchSet
func LookupGeoMatchSet(ctx *pulumi.Context, args *LookupGeoMatchSetArgs, opts ...pulumi.InvokeOption) (*LookupGeoMatchSetResult, error) {
	var rv LookupGeoMatchSetResult
	err := ctx.Invoke("aws-native:wafregional:getGeoMatchSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGeoMatchSetArgs struct {
	Id string `pulumi:"id"`
}

type LookupGeoMatchSetResult struct {
	GeoMatchConstraints []GeoMatchSetGeoMatchConstraint `pulumi:"geoMatchConstraints"`
	Id                  *string                         `pulumi:"id"`
}

func LookupGeoMatchSetOutput(ctx *pulumi.Context, args LookupGeoMatchSetOutputArgs, opts ...pulumi.InvokeOption) LookupGeoMatchSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGeoMatchSetResult, error) {
			args := v.(LookupGeoMatchSetArgs)
			r, err := LookupGeoMatchSet(ctx, &args, opts...)
			var s LookupGeoMatchSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGeoMatchSetResultOutput)
}

type LookupGeoMatchSetOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupGeoMatchSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGeoMatchSetArgs)(nil)).Elem()
}

type LookupGeoMatchSetResultOutput struct{ *pulumi.OutputState }

func (LookupGeoMatchSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGeoMatchSetResult)(nil)).Elem()
}

func (o LookupGeoMatchSetResultOutput) ToLookupGeoMatchSetResultOutput() LookupGeoMatchSetResultOutput {
	return o
}

func (o LookupGeoMatchSetResultOutput) ToLookupGeoMatchSetResultOutputWithContext(ctx context.Context) LookupGeoMatchSetResultOutput {
	return o
}

func (o LookupGeoMatchSetResultOutput) GeoMatchConstraints() GeoMatchSetGeoMatchConstraintArrayOutput {
	return o.ApplyT(func(v LookupGeoMatchSetResult) []GeoMatchSetGeoMatchConstraint { return v.GeoMatchConstraints }).(GeoMatchSetGeoMatchConstraintArrayOutput)
}

func (o LookupGeoMatchSetResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGeoMatchSetResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGeoMatchSetResultOutput{})
}
