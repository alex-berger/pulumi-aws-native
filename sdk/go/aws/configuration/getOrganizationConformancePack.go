// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package configuration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::Config::OrganizationConformancePack.
func LookupOrganizationConformancePack(ctx *pulumi.Context, args *LookupOrganizationConformancePackArgs, opts ...pulumi.InvokeOption) (*LookupOrganizationConformancePackResult, error) {
	var rv LookupOrganizationConformancePackResult
	err := ctx.Invoke("aws-native:configuration:getOrganizationConformancePack", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOrganizationConformancePackArgs struct {
	// The name of the organization conformance pack.
	OrganizationConformancePackName string `pulumi:"organizationConformancePackName"`
}

type LookupOrganizationConformancePackResult struct {
	// A list of ConformancePackInputParameter objects.
	ConformancePackInputParameters []OrganizationConformancePackConformancePackInputParameter `pulumi:"conformancePackInputParameters"`
	// AWS Config stores intermediate files while processing conformance pack template.
	DeliveryS3Bucket *string `pulumi:"deliveryS3Bucket"`
	// The prefix for the delivery S3 bucket.
	DeliveryS3KeyPrefix *string `pulumi:"deliveryS3KeyPrefix"`
	// A list of AWS accounts to be excluded from an organization conformance pack while deploying a conformance pack.
	ExcludedAccounts []string `pulumi:"excludedAccounts"`
}

func LookupOrganizationConformancePackOutput(ctx *pulumi.Context, args LookupOrganizationConformancePackOutputArgs, opts ...pulumi.InvokeOption) LookupOrganizationConformancePackResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrganizationConformancePackResult, error) {
			args := v.(LookupOrganizationConformancePackArgs)
			r, err := LookupOrganizationConformancePack(ctx, &args, opts...)
			var s LookupOrganizationConformancePackResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrganizationConformancePackResultOutput)
}

type LookupOrganizationConformancePackOutputArgs struct {
	// The name of the organization conformance pack.
	OrganizationConformancePackName pulumi.StringInput `pulumi:"organizationConformancePackName"`
}

func (LookupOrganizationConformancePackOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationConformancePackArgs)(nil)).Elem()
}

type LookupOrganizationConformancePackResultOutput struct{ *pulumi.OutputState }

func (LookupOrganizationConformancePackResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationConformancePackResult)(nil)).Elem()
}

func (o LookupOrganizationConformancePackResultOutput) ToLookupOrganizationConformancePackResultOutput() LookupOrganizationConformancePackResultOutput {
	return o
}

func (o LookupOrganizationConformancePackResultOutput) ToLookupOrganizationConformancePackResultOutputWithContext(ctx context.Context) LookupOrganizationConformancePackResultOutput {
	return o
}

// A list of ConformancePackInputParameter objects.
func (o LookupOrganizationConformancePackResultOutput) ConformancePackInputParameters() OrganizationConformancePackConformancePackInputParameterArrayOutput {
	return o.ApplyT(func(v LookupOrganizationConformancePackResult) []OrganizationConformancePackConformancePackInputParameter {
		return v.ConformancePackInputParameters
	}).(OrganizationConformancePackConformancePackInputParameterArrayOutput)
}

// AWS Config stores intermediate files while processing conformance pack template.
func (o LookupOrganizationConformancePackResultOutput) DeliveryS3Bucket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrganizationConformancePackResult) *string { return v.DeliveryS3Bucket }).(pulumi.StringPtrOutput)
}

// The prefix for the delivery S3 bucket.
func (o LookupOrganizationConformancePackResultOutput) DeliveryS3KeyPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrganizationConformancePackResult) *string { return v.DeliveryS3KeyPrefix }).(pulumi.StringPtrOutput)
}

// A list of AWS accounts to be excluded from an organization conformance pack while deploying a conformance pack.
func (o LookupOrganizationConformancePackResultOutput) ExcludedAccounts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupOrganizationConformancePackResult) []string { return v.ExcludedAccounts }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrganizationConformancePackResultOutput{})
}
