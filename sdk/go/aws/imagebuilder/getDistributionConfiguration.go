// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package imagebuilder

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::ImageBuilder::DistributionConfiguration
func LookupDistributionConfiguration(ctx *pulumi.Context, args *LookupDistributionConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupDistributionConfigurationResult, error) {
	var rv LookupDistributionConfigurationResult
	err := ctx.Invoke("aws-native:imagebuilder:getDistributionConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDistributionConfigurationArgs struct {
	// The Amazon Resource Name (ARN) of the distribution configuration.
	Arn string `pulumi:"arn"`
}

type LookupDistributionConfigurationResult struct {
	// The Amazon Resource Name (ARN) of the distribution configuration.
	Arn *string `pulumi:"arn"`
	// The description of the distribution configuration.
	Description *string `pulumi:"description"`
	// The distributions of the distribution configuration.
	Distributions []DistributionConfigurationDistribution `pulumi:"distributions"`
	// The tags associated with the component.
	Tags interface{} `pulumi:"tags"`
}

func LookupDistributionConfigurationOutput(ctx *pulumi.Context, args LookupDistributionConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupDistributionConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDistributionConfigurationResult, error) {
			args := v.(LookupDistributionConfigurationArgs)
			r, err := LookupDistributionConfiguration(ctx, &args, opts...)
			var s LookupDistributionConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDistributionConfigurationResultOutput)
}

type LookupDistributionConfigurationOutputArgs struct {
	// The Amazon Resource Name (ARN) of the distribution configuration.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupDistributionConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDistributionConfigurationArgs)(nil)).Elem()
}

type LookupDistributionConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupDistributionConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDistributionConfigurationResult)(nil)).Elem()
}

func (o LookupDistributionConfigurationResultOutput) ToLookupDistributionConfigurationResultOutput() LookupDistributionConfigurationResultOutput {
	return o
}

func (o LookupDistributionConfigurationResultOutput) ToLookupDistributionConfigurationResultOutputWithContext(ctx context.Context) LookupDistributionConfigurationResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the distribution configuration.
func (o LookupDistributionConfigurationResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDistributionConfigurationResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The description of the distribution configuration.
func (o LookupDistributionConfigurationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDistributionConfigurationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The distributions of the distribution configuration.
func (o LookupDistributionConfigurationResultOutput) Distributions() DistributionConfigurationDistributionArrayOutput {
	return o.ApplyT(func(v LookupDistributionConfigurationResult) []DistributionConfigurationDistribution {
		return v.Distributions
	}).(DistributionConfigurationDistributionArrayOutput)
}

// The tags associated with the component.
func (o LookupDistributionConfigurationResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDistributionConfigurationResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDistributionConfigurationResultOutput{})
}
