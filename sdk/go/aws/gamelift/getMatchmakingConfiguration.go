// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gamelift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::GameLift::MatchmakingConfiguration
func LookupMatchmakingConfiguration(ctx *pulumi.Context, args *LookupMatchmakingConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupMatchmakingConfigurationResult, error) {
	var rv LookupMatchmakingConfigurationResult
	err := ctx.Invoke("aws-native:gamelift:getMatchmakingConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMatchmakingConfigurationArgs struct {
	Id string `pulumi:"id"`
}

type LookupMatchmakingConfigurationResult struct {
	AcceptanceRequired       *bool                                  `pulumi:"acceptanceRequired"`
	AcceptanceTimeoutSeconds *int                                   `pulumi:"acceptanceTimeoutSeconds"`
	AdditionalPlayerCount    *int                                   `pulumi:"additionalPlayerCount"`
	Arn                      *string                                `pulumi:"arn"`
	BackfillMode             *string                                `pulumi:"backfillMode"`
	CustomEventData          *string                                `pulumi:"customEventData"`
	Description              *string                                `pulumi:"description"`
	FlexMatchMode            *string                                `pulumi:"flexMatchMode"`
	GameProperties           []MatchmakingConfigurationGameProperty `pulumi:"gameProperties"`
	GameSessionData          *string                                `pulumi:"gameSessionData"`
	GameSessionQueueArns     []string                               `pulumi:"gameSessionQueueArns"`
	Id                       *string                                `pulumi:"id"`
	NotificationTarget       *string                                `pulumi:"notificationTarget"`
	RequestTimeoutSeconds    *int                                   `pulumi:"requestTimeoutSeconds"`
	RuleSetName              *string                                `pulumi:"ruleSetName"`
	Tags                     []MatchmakingConfigurationTag          `pulumi:"tags"`
}

func LookupMatchmakingConfigurationOutput(ctx *pulumi.Context, args LookupMatchmakingConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupMatchmakingConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMatchmakingConfigurationResult, error) {
			args := v.(LookupMatchmakingConfigurationArgs)
			r, err := LookupMatchmakingConfiguration(ctx, &args, opts...)
			var s LookupMatchmakingConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMatchmakingConfigurationResultOutput)
}

type LookupMatchmakingConfigurationOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupMatchmakingConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMatchmakingConfigurationArgs)(nil)).Elem()
}

type LookupMatchmakingConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupMatchmakingConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMatchmakingConfigurationResult)(nil)).Elem()
}

func (o LookupMatchmakingConfigurationResultOutput) ToLookupMatchmakingConfigurationResultOutput() LookupMatchmakingConfigurationResultOutput {
	return o
}

func (o LookupMatchmakingConfigurationResultOutput) ToLookupMatchmakingConfigurationResultOutputWithContext(ctx context.Context) LookupMatchmakingConfigurationResultOutput {
	return o
}

func (o LookupMatchmakingConfigurationResultOutput) AcceptanceRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupMatchmakingConfigurationResult) *bool { return v.AcceptanceRequired }).(pulumi.BoolPtrOutput)
}

func (o LookupMatchmakingConfigurationResultOutput) AcceptanceTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupMatchmakingConfigurationResult) *int { return v.AcceptanceTimeoutSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupMatchmakingConfigurationResultOutput) AdditionalPlayerCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupMatchmakingConfigurationResult) *int { return v.AdditionalPlayerCount }).(pulumi.IntPtrOutput)
}

func (o LookupMatchmakingConfigurationResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMatchmakingConfigurationResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupMatchmakingConfigurationResultOutput) BackfillMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMatchmakingConfigurationResult) *string { return v.BackfillMode }).(pulumi.StringPtrOutput)
}

func (o LookupMatchmakingConfigurationResultOutput) CustomEventData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMatchmakingConfigurationResult) *string { return v.CustomEventData }).(pulumi.StringPtrOutput)
}

func (o LookupMatchmakingConfigurationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMatchmakingConfigurationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupMatchmakingConfigurationResultOutput) FlexMatchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMatchmakingConfigurationResult) *string { return v.FlexMatchMode }).(pulumi.StringPtrOutput)
}

func (o LookupMatchmakingConfigurationResultOutput) GameProperties() MatchmakingConfigurationGamePropertyArrayOutput {
	return o.ApplyT(func(v LookupMatchmakingConfigurationResult) []MatchmakingConfigurationGameProperty {
		return v.GameProperties
	}).(MatchmakingConfigurationGamePropertyArrayOutput)
}

func (o LookupMatchmakingConfigurationResultOutput) GameSessionData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMatchmakingConfigurationResult) *string { return v.GameSessionData }).(pulumi.StringPtrOutput)
}

func (o LookupMatchmakingConfigurationResultOutput) GameSessionQueueArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMatchmakingConfigurationResult) []string { return v.GameSessionQueueArns }).(pulumi.StringArrayOutput)
}

func (o LookupMatchmakingConfigurationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMatchmakingConfigurationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupMatchmakingConfigurationResultOutput) NotificationTarget() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMatchmakingConfigurationResult) *string { return v.NotificationTarget }).(pulumi.StringPtrOutput)
}

func (o LookupMatchmakingConfigurationResultOutput) RequestTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupMatchmakingConfigurationResult) *int { return v.RequestTimeoutSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupMatchmakingConfigurationResultOutput) RuleSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMatchmakingConfigurationResult) *string { return v.RuleSetName }).(pulumi.StringPtrOutput)
}

func (o LookupMatchmakingConfigurationResultOutput) Tags() MatchmakingConfigurationTagArrayOutput {
	return o.ApplyT(func(v LookupMatchmakingConfigurationResult) []MatchmakingConfigurationTag { return v.Tags }).(MatchmakingConfigurationTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMatchmakingConfigurationResultOutput{})
}
