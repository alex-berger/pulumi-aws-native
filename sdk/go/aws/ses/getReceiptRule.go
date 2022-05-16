// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SES::ReceiptRule
func LookupReceiptRule(ctx *pulumi.Context, args *LookupReceiptRuleArgs, opts ...pulumi.InvokeOption) (*LookupReceiptRuleResult, error) {
	var rv LookupReceiptRuleResult
	err := ctx.Invoke("aws-native:ses:getReceiptRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReceiptRuleArgs struct {
	Id string `pulumi:"id"`
}

type LookupReceiptRuleResult struct {
	After *string          `pulumi:"after"`
	Id    *string          `pulumi:"id"`
	Rule  *ReceiptRuleRule `pulumi:"rule"`
}

func LookupReceiptRuleOutput(ctx *pulumi.Context, args LookupReceiptRuleOutputArgs, opts ...pulumi.InvokeOption) LookupReceiptRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReceiptRuleResult, error) {
			args := v.(LookupReceiptRuleArgs)
			r, err := LookupReceiptRule(ctx, &args, opts...)
			var s LookupReceiptRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReceiptRuleResultOutput)
}

type LookupReceiptRuleOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupReceiptRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReceiptRuleArgs)(nil)).Elem()
}

type LookupReceiptRuleResultOutput struct{ *pulumi.OutputState }

func (LookupReceiptRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReceiptRuleResult)(nil)).Elem()
}

func (o LookupReceiptRuleResultOutput) ToLookupReceiptRuleResultOutput() LookupReceiptRuleResultOutput {
	return o
}

func (o LookupReceiptRuleResultOutput) ToLookupReceiptRuleResultOutputWithContext(ctx context.Context) LookupReceiptRuleResultOutput {
	return o
}

func (o LookupReceiptRuleResultOutput) After() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReceiptRuleResult) *string { return v.After }).(pulumi.StringPtrOutput)
}

func (o LookupReceiptRuleResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReceiptRuleResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupReceiptRuleResultOutput) Rule() ReceiptRuleRulePtrOutput {
	return o.ApplyT(func(v LookupReceiptRuleResult) *ReceiptRuleRule { return v.Rule }).(ReceiptRuleRulePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReceiptRuleResultOutput{})
}
