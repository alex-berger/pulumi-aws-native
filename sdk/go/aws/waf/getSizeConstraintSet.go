// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::WAF::SizeConstraintSet
func LookupSizeConstraintSet(ctx *pulumi.Context, args *LookupSizeConstraintSetArgs, opts ...pulumi.InvokeOption) (*LookupSizeConstraintSetResult, error) {
	var rv LookupSizeConstraintSetResult
	err := ctx.Invoke("aws-native:waf:getSizeConstraintSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSizeConstraintSetArgs struct {
	Id string `pulumi:"id"`
}

type LookupSizeConstraintSetResult struct {
	Id              *string                           `pulumi:"id"`
	SizeConstraints []SizeConstraintSetSizeConstraint `pulumi:"sizeConstraints"`
}

func LookupSizeConstraintSetOutput(ctx *pulumi.Context, args LookupSizeConstraintSetOutputArgs, opts ...pulumi.InvokeOption) LookupSizeConstraintSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSizeConstraintSetResult, error) {
			args := v.(LookupSizeConstraintSetArgs)
			r, err := LookupSizeConstraintSet(ctx, &args, opts...)
			var s LookupSizeConstraintSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSizeConstraintSetResultOutput)
}

type LookupSizeConstraintSetOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupSizeConstraintSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSizeConstraintSetArgs)(nil)).Elem()
}

type LookupSizeConstraintSetResultOutput struct{ *pulumi.OutputState }

func (LookupSizeConstraintSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSizeConstraintSetResult)(nil)).Elem()
}

func (o LookupSizeConstraintSetResultOutput) ToLookupSizeConstraintSetResultOutput() LookupSizeConstraintSetResultOutput {
	return o
}

func (o LookupSizeConstraintSetResultOutput) ToLookupSizeConstraintSetResultOutputWithContext(ctx context.Context) LookupSizeConstraintSetResultOutput {
	return o
}

func (o LookupSizeConstraintSetResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSizeConstraintSetResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSizeConstraintSetResultOutput) SizeConstraints() SizeConstraintSetSizeConstraintArrayOutput {
	return o.ApplyT(func(v LookupSizeConstraintSetResult) []SizeConstraintSetSizeConstraint { return v.SizeConstraints }).(SizeConstraintSetSizeConstraintArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSizeConstraintSetResultOutput{})
}
