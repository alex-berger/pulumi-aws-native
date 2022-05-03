// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Pinpoint::Campaign
func LookupCampaign(ctx *pulumi.Context, args *LookupCampaignArgs, opts ...pulumi.InvokeOption) (*LookupCampaignResult, error) {
	var rv LookupCampaignResult
	err := ctx.Invoke("aws-native:pinpoint:getCampaign", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCampaignArgs struct {
	CampaignId string `pulumi:"campaignId"`
}

type LookupCampaignResult struct {
	AdditionalTreatments []CampaignWriteTreatmentResource `pulumi:"additionalTreatments"`
	Arn                  *string                          `pulumi:"arn"`
	CampaignHook         *CampaignHook                    `pulumi:"campaignHook"`
	CampaignId           *string                          `pulumi:"campaignId"`
	Description          *string                          `pulumi:"description"`
	HoldoutPercent       *int                             `pulumi:"holdoutPercent"`
	IsPaused             *bool                            `pulumi:"isPaused"`
	Limits               *CampaignLimits                  `pulumi:"limits"`
	MessageConfiguration *CampaignMessageConfiguration    `pulumi:"messageConfiguration"`
	Name                 *string                          `pulumi:"name"`
	Priority             *int                             `pulumi:"priority"`
	Schedule             *CampaignSchedule                `pulumi:"schedule"`
	SegmentId            *string                          `pulumi:"segmentId"`
	SegmentVersion       *int                             `pulumi:"segmentVersion"`
	Tags                 interface{}                      `pulumi:"tags"`
	TreatmentDescription *string                          `pulumi:"treatmentDescription"`
	TreatmentName        *string                          `pulumi:"treatmentName"`
}

func LookupCampaignOutput(ctx *pulumi.Context, args LookupCampaignOutputArgs, opts ...pulumi.InvokeOption) LookupCampaignResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCampaignResult, error) {
			args := v.(LookupCampaignArgs)
			r, err := LookupCampaign(ctx, &args, opts...)
			var s LookupCampaignResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCampaignResultOutput)
}

type LookupCampaignOutputArgs struct {
	CampaignId pulumi.StringInput `pulumi:"campaignId"`
}

func (LookupCampaignOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCampaignArgs)(nil)).Elem()
}

type LookupCampaignResultOutput struct{ *pulumi.OutputState }

func (LookupCampaignResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCampaignResult)(nil)).Elem()
}

func (o LookupCampaignResultOutput) ToLookupCampaignResultOutput() LookupCampaignResultOutput {
	return o
}

func (o LookupCampaignResultOutput) ToLookupCampaignResultOutputWithContext(ctx context.Context) LookupCampaignResultOutput {
	return o
}

func (o LookupCampaignResultOutput) AdditionalTreatments() CampaignWriteTreatmentResourceArrayOutput {
	return o.ApplyT(func(v LookupCampaignResult) []CampaignWriteTreatmentResource { return v.AdditionalTreatments }).(CampaignWriteTreatmentResourceArrayOutput)
}

func (o LookupCampaignResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCampaignResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupCampaignResultOutput) CampaignHook() CampaignHookPtrOutput {
	return o.ApplyT(func(v LookupCampaignResult) *CampaignHook { return v.CampaignHook }).(CampaignHookPtrOutput)
}

func (o LookupCampaignResultOutput) CampaignId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCampaignResult) *string { return v.CampaignId }).(pulumi.StringPtrOutput)
}

func (o LookupCampaignResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCampaignResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupCampaignResultOutput) HoldoutPercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCampaignResult) *int { return v.HoldoutPercent }).(pulumi.IntPtrOutput)
}

func (o LookupCampaignResultOutput) IsPaused() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupCampaignResult) *bool { return v.IsPaused }).(pulumi.BoolPtrOutput)
}

func (o LookupCampaignResultOutput) Limits() CampaignLimitsPtrOutput {
	return o.ApplyT(func(v LookupCampaignResult) *CampaignLimits { return v.Limits }).(CampaignLimitsPtrOutput)
}

func (o LookupCampaignResultOutput) MessageConfiguration() CampaignMessageConfigurationPtrOutput {
	return o.ApplyT(func(v LookupCampaignResult) *CampaignMessageConfiguration { return v.MessageConfiguration }).(CampaignMessageConfigurationPtrOutput)
}

func (o LookupCampaignResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCampaignResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupCampaignResultOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCampaignResult) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o LookupCampaignResultOutput) Schedule() CampaignSchedulePtrOutput {
	return o.ApplyT(func(v LookupCampaignResult) *CampaignSchedule { return v.Schedule }).(CampaignSchedulePtrOutput)
}

func (o LookupCampaignResultOutput) SegmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCampaignResult) *string { return v.SegmentId }).(pulumi.StringPtrOutput)
}

func (o LookupCampaignResultOutput) SegmentVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCampaignResult) *int { return v.SegmentVersion }).(pulumi.IntPtrOutput)
}

func (o LookupCampaignResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupCampaignResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func (o LookupCampaignResultOutput) TreatmentDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCampaignResult) *string { return v.TreatmentDescription }).(pulumi.StringPtrOutput)
}

func (o LookupCampaignResultOutput) TreatmentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCampaignResult) *string { return v.TreatmentName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCampaignResultOutput{})
}
