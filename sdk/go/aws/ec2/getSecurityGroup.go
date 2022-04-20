// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::SecurityGroup
func LookupSecurityGroup(ctx *pulumi.Context, args *LookupSecurityGroupArgs, opts ...pulumi.InvokeOption) (*LookupSecurityGroupResult, error) {
	var rv LookupSecurityGroupResult
	err := ctx.Invoke("aws-native:ec2:getSecurityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityGroupArgs struct {
	Id string `pulumi:"id"`
}

type LookupSecurityGroupResult struct {
	GroupId              *string                    `pulumi:"groupId"`
	Id                   *string                    `pulumi:"id"`
	SecurityGroupEgress  []SecurityGroupEgressType  `pulumi:"securityGroupEgress"`
	SecurityGroupIngress []SecurityGroupIngressType `pulumi:"securityGroupIngress"`
	Tags                 []SecurityGroupTag         `pulumi:"tags"`
}

func LookupSecurityGroupOutput(ctx *pulumi.Context, args LookupSecurityGroupOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityGroupResult, error) {
			args := v.(LookupSecurityGroupArgs)
			r, err := LookupSecurityGroup(ctx, &args, opts...)
			var s LookupSecurityGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityGroupResultOutput)
}

type LookupSecurityGroupOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupSecurityGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityGroupArgs)(nil)).Elem()
}

type LookupSecurityGroupResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityGroupResult)(nil)).Elem()
}

func (o LookupSecurityGroupResultOutput) ToLookupSecurityGroupResultOutput() LookupSecurityGroupResultOutput {
	return o
}

func (o LookupSecurityGroupResultOutput) ToLookupSecurityGroupResultOutputWithContext(ctx context.Context) LookupSecurityGroupResultOutput {
	return o
}

func (o LookupSecurityGroupResultOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityGroupResult) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityGroupResultOutput) SecurityGroupEgress() SecurityGroupEgressTypeArrayOutput {
	return o.ApplyT(func(v LookupSecurityGroupResult) []SecurityGroupEgressType { return v.SecurityGroupEgress }).(SecurityGroupEgressTypeArrayOutput)
}

func (o LookupSecurityGroupResultOutput) SecurityGroupIngress() SecurityGroupIngressTypeArrayOutput {
	return o.ApplyT(func(v LookupSecurityGroupResult) []SecurityGroupIngressType { return v.SecurityGroupIngress }).(SecurityGroupIngressTypeArrayOutput)
}

func (o LookupSecurityGroupResultOutput) Tags() SecurityGroupTagArrayOutput {
	return o.ApplyT(func(v LookupSecurityGroupResult) []SecurityGroupTag { return v.Tags }).(SecurityGroupTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityGroupResultOutput{})
}
