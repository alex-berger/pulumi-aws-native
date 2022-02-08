// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SNS::TopicPolicy
func LookupTopicPolicy(ctx *pulumi.Context, args *LookupTopicPolicyArgs, opts ...pulumi.InvokeOption) (*LookupTopicPolicyResult, error) {
	var rv LookupTopicPolicyResult
	err := ctx.Invoke("aws-native:sns:getTopicPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTopicPolicyArgs struct {
	Id string `pulumi:"id"`
}

type LookupTopicPolicyResult struct {
	Id             *string     `pulumi:"id"`
	PolicyDocument interface{} `pulumi:"policyDocument"`
	Topics         []string    `pulumi:"topics"`
}

func LookupTopicPolicyOutput(ctx *pulumi.Context, args LookupTopicPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupTopicPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTopicPolicyResult, error) {
			args := v.(LookupTopicPolicyArgs)
			r, err := LookupTopicPolicy(ctx, &args, opts...)
			return *r, err
		}).(LookupTopicPolicyResultOutput)
}

type LookupTopicPolicyOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupTopicPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicPolicyArgs)(nil)).Elem()
}

type LookupTopicPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupTopicPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicPolicyResult)(nil)).Elem()
}

func (o LookupTopicPolicyResultOutput) ToLookupTopicPolicyResultOutput() LookupTopicPolicyResultOutput {
	return o
}

func (o LookupTopicPolicyResultOutput) ToLookupTopicPolicyResultOutputWithContext(ctx context.Context) LookupTopicPolicyResultOutput {
	return o
}

func (o LookupTopicPolicyResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicPolicyResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupTopicPolicyResultOutput) PolicyDocument() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupTopicPolicyResult) interface{} { return v.PolicyDocument }).(pulumi.AnyOutput)
}

func (o LookupTopicPolicyResultOutput) Topics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTopicPolicyResult) []string { return v.Topics }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTopicPolicyResultOutput{})
}