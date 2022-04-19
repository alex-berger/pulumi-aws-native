// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancingv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ElasticLoadBalancingV2::ListenerRule
func LookupListenerRule(ctx *pulumi.Context, args *LookupListenerRuleArgs, opts ...pulumi.InvokeOption) (*LookupListenerRuleResult, error) {
	var rv LookupListenerRuleResult
	err := ctx.Invoke("aws-native:elasticloadbalancingv2:getListenerRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupListenerRuleArgs struct {
	RuleArn string `pulumi:"ruleArn"`
}

type LookupListenerRuleResult struct {
	Actions    []ListenerRuleAction        `pulumi:"actions"`
	Conditions []ListenerRuleRuleCondition `pulumi:"conditions"`
	IsDefault  *bool                       `pulumi:"isDefault"`
	Priority   *int                        `pulumi:"priority"`
	RuleArn    *string                     `pulumi:"ruleArn"`
}

func LookupListenerRuleOutput(ctx *pulumi.Context, args LookupListenerRuleOutputArgs, opts ...pulumi.InvokeOption) LookupListenerRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupListenerRuleResult, error) {
			args := v.(LookupListenerRuleArgs)
			r, err := LookupListenerRule(ctx, &args, opts...)
			var s LookupListenerRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupListenerRuleResultOutput)
}

type LookupListenerRuleOutputArgs struct {
	RuleArn pulumi.StringInput `pulumi:"ruleArn"`
}

func (LookupListenerRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupListenerRuleArgs)(nil)).Elem()
}

type LookupListenerRuleResultOutput struct{ *pulumi.OutputState }

func (LookupListenerRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupListenerRuleResult)(nil)).Elem()
}

func (o LookupListenerRuleResultOutput) ToLookupListenerRuleResultOutput() LookupListenerRuleResultOutput {
	return o
}

func (o LookupListenerRuleResultOutput) ToLookupListenerRuleResultOutputWithContext(ctx context.Context) LookupListenerRuleResultOutput {
	return o
}

func (o LookupListenerRuleResultOutput) Actions() ListenerRuleActionArrayOutput {
	return o.ApplyT(func(v LookupListenerRuleResult) []ListenerRuleAction { return v.Actions }).(ListenerRuleActionArrayOutput)
}

func (o LookupListenerRuleResultOutput) Conditions() ListenerRuleRuleConditionArrayOutput {
	return o.ApplyT(func(v LookupListenerRuleResult) []ListenerRuleRuleCondition { return v.Conditions }).(ListenerRuleRuleConditionArrayOutput)
}

func (o LookupListenerRuleResultOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupListenerRuleResult) *bool { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

func (o LookupListenerRuleResultOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupListenerRuleResult) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o LookupListenerRuleResultOutput) RuleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupListenerRuleResult) *string { return v.RuleArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupListenerRuleResultOutput{})
}
