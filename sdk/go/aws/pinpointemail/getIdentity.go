// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pinpointemail

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::PinpointEmail::Identity
func LookupIdentity(ctx *pulumi.Context, args *LookupIdentityArgs, opts ...pulumi.InvokeOption) (*LookupIdentityResult, error) {
	var rv LookupIdentityResult
	err := ctx.Invoke("aws-native:pinpointemail:getIdentity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIdentityArgs struct {
	Id string `pulumi:"id"`
}

type LookupIdentityResult struct {
	DkimSigningEnabled        *bool                       `pulumi:"dkimSigningEnabled"`
	FeedbackForwardingEnabled *bool                       `pulumi:"feedbackForwardingEnabled"`
	Id                        *string                     `pulumi:"id"`
	IdentityDNSRecordName1    *string                     `pulumi:"identityDNSRecordName1"`
	IdentityDNSRecordName2    *string                     `pulumi:"identityDNSRecordName2"`
	IdentityDNSRecordName3    *string                     `pulumi:"identityDNSRecordName3"`
	IdentityDNSRecordValue1   *string                     `pulumi:"identityDNSRecordValue1"`
	IdentityDNSRecordValue2   *string                     `pulumi:"identityDNSRecordValue2"`
	IdentityDNSRecordValue3   *string                     `pulumi:"identityDNSRecordValue3"`
	MailFromAttributes        *IdentityMailFromAttributes `pulumi:"mailFromAttributes"`
	Tags                      []IdentityTags              `pulumi:"tags"`
}

func LookupIdentityOutput(ctx *pulumi.Context, args LookupIdentityOutputArgs, opts ...pulumi.InvokeOption) LookupIdentityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIdentityResult, error) {
			args := v.(LookupIdentityArgs)
			r, err := LookupIdentity(ctx, &args, opts...)
			var s LookupIdentityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIdentityResultOutput)
}

type LookupIdentityOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupIdentityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdentityArgs)(nil)).Elem()
}

type LookupIdentityResultOutput struct{ *pulumi.OutputState }

func (LookupIdentityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdentityResult)(nil)).Elem()
}

func (o LookupIdentityResultOutput) ToLookupIdentityResultOutput() LookupIdentityResultOutput {
	return o
}

func (o LookupIdentityResultOutput) ToLookupIdentityResultOutputWithContext(ctx context.Context) LookupIdentityResultOutput {
	return o
}

func (o LookupIdentityResultOutput) DkimSigningEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupIdentityResult) *bool { return v.DkimSigningEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupIdentityResultOutput) FeedbackForwardingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupIdentityResult) *bool { return v.FeedbackForwardingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupIdentityResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdentityResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupIdentityResultOutput) IdentityDNSRecordName1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdentityResult) *string { return v.IdentityDNSRecordName1 }).(pulumi.StringPtrOutput)
}

func (o LookupIdentityResultOutput) IdentityDNSRecordName2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdentityResult) *string { return v.IdentityDNSRecordName2 }).(pulumi.StringPtrOutput)
}

func (o LookupIdentityResultOutput) IdentityDNSRecordName3() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdentityResult) *string { return v.IdentityDNSRecordName3 }).(pulumi.StringPtrOutput)
}

func (o LookupIdentityResultOutput) IdentityDNSRecordValue1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdentityResult) *string { return v.IdentityDNSRecordValue1 }).(pulumi.StringPtrOutput)
}

func (o LookupIdentityResultOutput) IdentityDNSRecordValue2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdentityResult) *string { return v.IdentityDNSRecordValue2 }).(pulumi.StringPtrOutput)
}

func (o LookupIdentityResultOutput) IdentityDNSRecordValue3() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdentityResult) *string { return v.IdentityDNSRecordValue3 }).(pulumi.StringPtrOutput)
}

func (o LookupIdentityResultOutput) MailFromAttributes() IdentityMailFromAttributesPtrOutput {
	return o.ApplyT(func(v LookupIdentityResult) *IdentityMailFromAttributes { return v.MailFromAttributes }).(IdentityMailFromAttributesPtrOutput)
}

func (o LookupIdentityResultOutput) Tags() IdentityTagsArrayOutput {
	return o.ApplyT(func(v LookupIdentityResult) []IdentityTags { return v.Tags }).(IdentityTagsArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIdentityResultOutput{})
}
