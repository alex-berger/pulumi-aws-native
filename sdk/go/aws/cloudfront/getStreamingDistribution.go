// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::CloudFront::StreamingDistribution
func LookupStreamingDistribution(ctx *pulumi.Context, args *LookupStreamingDistributionArgs, opts ...pulumi.InvokeOption) (*LookupStreamingDistributionResult, error) {
	var rv LookupStreamingDistributionResult
	err := ctx.Invoke("aws-native:cloudfront:getStreamingDistribution", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStreamingDistributionArgs struct {
	Id string `pulumi:"id"`
}

type LookupStreamingDistributionResult struct {
	DomainName                  *string                      `pulumi:"domainName"`
	Id                          *string                      `pulumi:"id"`
	StreamingDistributionConfig *StreamingDistributionConfig `pulumi:"streamingDistributionConfig"`
	Tags                        []StreamingDistributionTag   `pulumi:"tags"`
}

func LookupStreamingDistributionOutput(ctx *pulumi.Context, args LookupStreamingDistributionOutputArgs, opts ...pulumi.InvokeOption) LookupStreamingDistributionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStreamingDistributionResult, error) {
			args := v.(LookupStreamingDistributionArgs)
			r, err := LookupStreamingDistribution(ctx, &args, opts...)
			var s LookupStreamingDistributionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStreamingDistributionResultOutput)
}

type LookupStreamingDistributionOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupStreamingDistributionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamingDistributionArgs)(nil)).Elem()
}

type LookupStreamingDistributionResultOutput struct{ *pulumi.OutputState }

func (LookupStreamingDistributionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamingDistributionResult)(nil)).Elem()
}

func (o LookupStreamingDistributionResultOutput) ToLookupStreamingDistributionResultOutput() LookupStreamingDistributionResultOutput {
	return o
}

func (o LookupStreamingDistributionResultOutput) ToLookupStreamingDistributionResultOutputWithContext(ctx context.Context) LookupStreamingDistributionResultOutput {
	return o
}

func (o LookupStreamingDistributionResultOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingDistributionResult) *string { return v.DomainName }).(pulumi.StringPtrOutput)
}

func (o LookupStreamingDistributionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingDistributionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupStreamingDistributionResultOutput) StreamingDistributionConfig() StreamingDistributionConfigPtrOutput {
	return o.ApplyT(func(v LookupStreamingDistributionResult) *StreamingDistributionConfig {
		return v.StreamingDistributionConfig
	}).(StreamingDistributionConfigPtrOutput)
}

func (o LookupStreamingDistributionResultOutput) Tags() StreamingDistributionTagArrayOutput {
	return o.ApplyT(func(v LookupStreamingDistributionResult) []StreamingDistributionTag { return v.Tags }).(StreamingDistributionTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStreamingDistributionResultOutput{})
}
