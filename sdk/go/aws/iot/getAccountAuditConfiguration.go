// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configures the Device Defender audit settings for this account. Settings include how audit notifications are sent and which audit checks are enabled or disabled.
func LookupAccountAuditConfiguration(ctx *pulumi.Context, args *LookupAccountAuditConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupAccountAuditConfigurationResult, error) {
	var rv LookupAccountAuditConfigurationResult
	err := ctx.Invoke("aws-native:iot:getAccountAuditConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccountAuditConfigurationArgs struct {
	// Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).
	AccountId string `pulumi:"accountId"`
}

type LookupAccountAuditConfigurationResult struct {
	AuditCheckConfigurations              *AccountAuditConfigurationAuditCheckConfigurations              `pulumi:"auditCheckConfigurations"`
	AuditNotificationTargetConfigurations *AccountAuditConfigurationAuditNotificationTargetConfigurations `pulumi:"auditNotificationTargetConfigurations"`
	// The ARN of the role that grants permission to AWS IoT to access information about your devices, policies, certificates and other items as required when performing an audit.
	RoleArn *string `pulumi:"roleArn"`
}

func LookupAccountAuditConfigurationOutput(ctx *pulumi.Context, args LookupAccountAuditConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupAccountAuditConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccountAuditConfigurationResult, error) {
			args := v.(LookupAccountAuditConfigurationArgs)
			r, err := LookupAccountAuditConfiguration(ctx, &args, opts...)
			var s LookupAccountAuditConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccountAuditConfigurationResultOutput)
}

type LookupAccountAuditConfigurationOutputArgs struct {
	// Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).
	AccountId pulumi.StringInput `pulumi:"accountId"`
}

func (LookupAccountAuditConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountAuditConfigurationArgs)(nil)).Elem()
}

type LookupAccountAuditConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupAccountAuditConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountAuditConfigurationResult)(nil)).Elem()
}

func (o LookupAccountAuditConfigurationResultOutput) ToLookupAccountAuditConfigurationResultOutput() LookupAccountAuditConfigurationResultOutput {
	return o
}

func (o LookupAccountAuditConfigurationResultOutput) ToLookupAccountAuditConfigurationResultOutputWithContext(ctx context.Context) LookupAccountAuditConfigurationResultOutput {
	return o
}

func (o LookupAccountAuditConfigurationResultOutput) AuditCheckConfigurations() AccountAuditConfigurationAuditCheckConfigurationsPtrOutput {
	return o.ApplyT(func(v LookupAccountAuditConfigurationResult) *AccountAuditConfigurationAuditCheckConfigurations {
		return v.AuditCheckConfigurations
	}).(AccountAuditConfigurationAuditCheckConfigurationsPtrOutput)
}

func (o LookupAccountAuditConfigurationResultOutput) AuditNotificationTargetConfigurations() AccountAuditConfigurationAuditNotificationTargetConfigurationsPtrOutput {
	return o.ApplyT(func(v LookupAccountAuditConfigurationResult) *AccountAuditConfigurationAuditNotificationTargetConfigurations {
		return v.AuditNotificationTargetConfigurations
	}).(AccountAuditConfigurationAuditNotificationTargetConfigurationsPtrOutput)
}

// The ARN of the role that grants permission to AWS IoT to access information about your devices, policies, certificates and other items as required when performing an audit.
func (o LookupAccountAuditConfigurationResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountAuditConfigurationResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccountAuditConfigurationResultOutput{})
}
