// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ServiceCatalog::ResourceUpdateConstraint
func LookupResourceUpdateConstraint(ctx *pulumi.Context, args *LookupResourceUpdateConstraintArgs, opts ...pulumi.InvokeOption) (*LookupResourceUpdateConstraintResult, error) {
	var rv LookupResourceUpdateConstraintResult
	err := ctx.Invoke("aws-native:servicecatalog:getResourceUpdateConstraint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceUpdateConstraintArgs struct {
	Id string `pulumi:"id"`
}

type LookupResourceUpdateConstraintResult struct {
	AcceptLanguage                *string `pulumi:"acceptLanguage"`
	Description                   *string `pulumi:"description"`
	Id                            *string `pulumi:"id"`
	TagUpdateOnProvisionedProduct *string `pulumi:"tagUpdateOnProvisionedProduct"`
}

func LookupResourceUpdateConstraintOutput(ctx *pulumi.Context, args LookupResourceUpdateConstraintOutputArgs, opts ...pulumi.InvokeOption) LookupResourceUpdateConstraintResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourceUpdateConstraintResult, error) {
			args := v.(LookupResourceUpdateConstraintArgs)
			r, err := LookupResourceUpdateConstraint(ctx, &args, opts...)
			var s LookupResourceUpdateConstraintResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourceUpdateConstraintResultOutput)
}

type LookupResourceUpdateConstraintOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupResourceUpdateConstraintOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceUpdateConstraintArgs)(nil)).Elem()
}

type LookupResourceUpdateConstraintResultOutput struct{ *pulumi.OutputState }

func (LookupResourceUpdateConstraintResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceUpdateConstraintResult)(nil)).Elem()
}

func (o LookupResourceUpdateConstraintResultOutput) ToLookupResourceUpdateConstraintResultOutput() LookupResourceUpdateConstraintResultOutput {
	return o
}

func (o LookupResourceUpdateConstraintResultOutput) ToLookupResourceUpdateConstraintResultOutputWithContext(ctx context.Context) LookupResourceUpdateConstraintResultOutput {
	return o
}

func (o LookupResourceUpdateConstraintResultOutput) AcceptLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceUpdateConstraintResult) *string { return v.AcceptLanguage }).(pulumi.StringPtrOutput)
}

func (o LookupResourceUpdateConstraintResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceUpdateConstraintResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupResourceUpdateConstraintResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceUpdateConstraintResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupResourceUpdateConstraintResultOutput) TagUpdateOnProvisionedProduct() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceUpdateConstraintResult) *string { return v.TagUpdateOnProvisionedProduct }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceUpdateConstraintResultOutput{})
}
