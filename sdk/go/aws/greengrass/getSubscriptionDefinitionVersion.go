// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package greengrass

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Greengrass::SubscriptionDefinitionVersion
func LookupSubscriptionDefinitionVersion(ctx *pulumi.Context, args *LookupSubscriptionDefinitionVersionArgs, opts ...pulumi.InvokeOption) (*LookupSubscriptionDefinitionVersionResult, error) {
	var rv LookupSubscriptionDefinitionVersionResult
	err := ctx.Invoke("aws-native:greengrass:getSubscriptionDefinitionVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubscriptionDefinitionVersionArgs struct {
	Id string `pulumi:"id"`
}

type LookupSubscriptionDefinitionVersionResult struct {
	Id *string `pulumi:"id"`
}

func LookupSubscriptionDefinitionVersionOutput(ctx *pulumi.Context, args LookupSubscriptionDefinitionVersionOutputArgs, opts ...pulumi.InvokeOption) LookupSubscriptionDefinitionVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubscriptionDefinitionVersionResult, error) {
			args := v.(LookupSubscriptionDefinitionVersionArgs)
			r, err := LookupSubscriptionDefinitionVersion(ctx, &args, opts...)
			var s LookupSubscriptionDefinitionVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubscriptionDefinitionVersionResultOutput)
}

type LookupSubscriptionDefinitionVersionOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupSubscriptionDefinitionVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionDefinitionVersionArgs)(nil)).Elem()
}

type LookupSubscriptionDefinitionVersionResultOutput struct{ *pulumi.OutputState }

func (LookupSubscriptionDefinitionVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionDefinitionVersionResult)(nil)).Elem()
}

func (o LookupSubscriptionDefinitionVersionResultOutput) ToLookupSubscriptionDefinitionVersionResultOutput() LookupSubscriptionDefinitionVersionResultOutput {
	return o
}

func (o LookupSubscriptionDefinitionVersionResultOutput) ToLookupSubscriptionDefinitionVersionResultOutputWithContext(ctx context.Context) LookupSubscriptionDefinitionVersionResultOutput {
	return o
}

func (o LookupSubscriptionDefinitionVersionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionDefinitionVersionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubscriptionDefinitionVersionResultOutput{})
}
