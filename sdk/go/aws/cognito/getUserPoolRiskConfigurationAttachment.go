// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Cognito::UserPoolRiskConfigurationAttachment
func LookupUserPoolRiskConfigurationAttachment(ctx *pulumi.Context, args *LookupUserPoolRiskConfigurationAttachmentArgs, opts ...pulumi.InvokeOption) (*LookupUserPoolRiskConfigurationAttachmentResult, error) {
	var rv LookupUserPoolRiskConfigurationAttachmentResult
	err := ctx.Invoke("aws-native:cognito:getUserPoolRiskConfigurationAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserPoolRiskConfigurationAttachmentArgs struct {
	Id string `pulumi:"id"`
}

type LookupUserPoolRiskConfigurationAttachmentResult struct {
	AccountTakeoverRiskConfiguration        *UserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationType        `pulumi:"accountTakeoverRiskConfiguration"`
	CompromisedCredentialsRiskConfiguration *UserPoolRiskConfigurationAttachmentCompromisedCredentialsRiskConfigurationType `pulumi:"compromisedCredentialsRiskConfiguration"`
	Id                                      *string                                                                         `pulumi:"id"`
	RiskExceptionConfiguration              *UserPoolRiskConfigurationAttachmentRiskExceptionConfigurationType              `pulumi:"riskExceptionConfiguration"`
}

func LookupUserPoolRiskConfigurationAttachmentOutput(ctx *pulumi.Context, args LookupUserPoolRiskConfigurationAttachmentOutputArgs, opts ...pulumi.InvokeOption) LookupUserPoolRiskConfigurationAttachmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserPoolRiskConfigurationAttachmentResult, error) {
			args := v.(LookupUserPoolRiskConfigurationAttachmentArgs)
			r, err := LookupUserPoolRiskConfigurationAttachment(ctx, &args, opts...)
			var s LookupUserPoolRiskConfigurationAttachmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserPoolRiskConfigurationAttachmentResultOutput)
}

type LookupUserPoolRiskConfigurationAttachmentOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupUserPoolRiskConfigurationAttachmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserPoolRiskConfigurationAttachmentArgs)(nil)).Elem()
}

type LookupUserPoolRiskConfigurationAttachmentResultOutput struct{ *pulumi.OutputState }

func (LookupUserPoolRiskConfigurationAttachmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserPoolRiskConfigurationAttachmentResult)(nil)).Elem()
}

func (o LookupUserPoolRiskConfigurationAttachmentResultOutput) ToLookupUserPoolRiskConfigurationAttachmentResultOutput() LookupUserPoolRiskConfigurationAttachmentResultOutput {
	return o
}

func (o LookupUserPoolRiskConfigurationAttachmentResultOutput) ToLookupUserPoolRiskConfigurationAttachmentResultOutputWithContext(ctx context.Context) LookupUserPoolRiskConfigurationAttachmentResultOutput {
	return o
}

func (o LookupUserPoolRiskConfigurationAttachmentResultOutput) AccountTakeoverRiskConfiguration() UserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationTypePtrOutput {
	return o.ApplyT(func(v LookupUserPoolRiskConfigurationAttachmentResult) *UserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationType {
		return v.AccountTakeoverRiskConfiguration
	}).(UserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationTypePtrOutput)
}

func (o LookupUserPoolRiskConfigurationAttachmentResultOutput) CompromisedCredentialsRiskConfiguration() UserPoolRiskConfigurationAttachmentCompromisedCredentialsRiskConfigurationTypePtrOutput {
	return o.ApplyT(func(v LookupUserPoolRiskConfigurationAttachmentResult) *UserPoolRiskConfigurationAttachmentCompromisedCredentialsRiskConfigurationType {
		return v.CompromisedCredentialsRiskConfiguration
	}).(UserPoolRiskConfigurationAttachmentCompromisedCredentialsRiskConfigurationTypePtrOutput)
}

func (o LookupUserPoolRiskConfigurationAttachmentResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolRiskConfigurationAttachmentResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupUserPoolRiskConfigurationAttachmentResultOutput) RiskExceptionConfiguration() UserPoolRiskConfigurationAttachmentRiskExceptionConfigurationTypePtrOutput {
	return o.ApplyT(func(v LookupUserPoolRiskConfigurationAttachmentResult) *UserPoolRiskConfigurationAttachmentRiskExceptionConfigurationType {
		return v.RiskExceptionConfiguration
	}).(UserPoolRiskConfigurationAttachmentRiskExceptionConfigurationTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserPoolRiskConfigurationAttachmentResultOutput{})
}