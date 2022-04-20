// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Lambda::EventSourceMapping
func LookupEventSourceMapping(ctx *pulumi.Context, args *LookupEventSourceMappingArgs, opts ...pulumi.InvokeOption) (*LookupEventSourceMappingResult, error) {
	var rv LookupEventSourceMappingResult
	err := ctx.Invoke("aws-native:lambda:getEventSourceMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventSourceMappingArgs struct {
	// Event Source Mapping Identifier UUID.
	Id string `pulumi:"id"`
}

type LookupEventSourceMappingResult struct {
	// The maximum number of items to retrieve in a single batch.
	BatchSize *int `pulumi:"batchSize"`
	// (Streams) If the function returns an error, split the batch in two and retry.
	BisectBatchOnFunctionError *bool `pulumi:"bisectBatchOnFunctionError"`
	// (Streams) An Amazon SQS queue or Amazon SNS topic destination for discarded records.
	DestinationConfig *EventSourceMappingDestinationConfig `pulumi:"destinationConfig"`
	// Disables the event source mapping to pause polling and invocation.
	Enabled *bool `pulumi:"enabled"`
	// The filter criteria to control event filtering.
	FilterCriteria *EventSourceMappingFilterCriteria `pulumi:"filterCriteria"`
	// The name of the Lambda function.
	FunctionName *string `pulumi:"functionName"`
	// (Streams) A list of response types supported by the function.
	FunctionResponseTypes []EventSourceMappingFunctionResponseTypesItem `pulumi:"functionResponseTypes"`
	// Event Source Mapping Identifier UUID.
	Id *string `pulumi:"id"`
	// (Streams) The maximum amount of time to gather records before invoking the function, in seconds.
	MaximumBatchingWindowInSeconds *int `pulumi:"maximumBatchingWindowInSeconds"`
	// (Streams) The maximum age of a record that Lambda sends to a function for processing.
	MaximumRecordAgeInSeconds *int `pulumi:"maximumRecordAgeInSeconds"`
	// (Streams) The maximum number of times to retry when the function returns an error.
	MaximumRetryAttempts *int `pulumi:"maximumRetryAttempts"`
	// (Streams) The number of batches to process from each shard concurrently.
	ParallelizationFactor *int `pulumi:"parallelizationFactor"`
	// (ActiveMQ) A list of ActiveMQ queues.
	Queues []string `pulumi:"queues"`
	// A list of SourceAccessConfiguration.
	SourceAccessConfigurations []EventSourceMappingSourceAccessConfiguration `pulumi:"sourceAccessConfigurations"`
	// (Kafka) A list of Kafka topics.
	Topics []string `pulumi:"topics"`
	// (Streams) Tumbling window (non-overlapping time window) duration to perform aggregations.
	TumblingWindowInSeconds *int `pulumi:"tumblingWindowInSeconds"`
}

func LookupEventSourceMappingOutput(ctx *pulumi.Context, args LookupEventSourceMappingOutputArgs, opts ...pulumi.InvokeOption) LookupEventSourceMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventSourceMappingResult, error) {
			args := v.(LookupEventSourceMappingArgs)
			r, err := LookupEventSourceMapping(ctx, &args, opts...)
			var s LookupEventSourceMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEventSourceMappingResultOutput)
}

type LookupEventSourceMappingOutputArgs struct {
	// Event Source Mapping Identifier UUID.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupEventSourceMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventSourceMappingArgs)(nil)).Elem()
}

type LookupEventSourceMappingResultOutput struct{ *pulumi.OutputState }

func (LookupEventSourceMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventSourceMappingResult)(nil)).Elem()
}

func (o LookupEventSourceMappingResultOutput) ToLookupEventSourceMappingResultOutput() LookupEventSourceMappingResultOutput {
	return o
}

func (o LookupEventSourceMappingResultOutput) ToLookupEventSourceMappingResultOutputWithContext(ctx context.Context) LookupEventSourceMappingResultOutput {
	return o
}

// The maximum number of items to retrieve in a single batch.
func (o LookupEventSourceMappingResultOutput) BatchSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *int { return v.BatchSize }).(pulumi.IntPtrOutput)
}

// (Streams) If the function returns an error, split the batch in two and retry.
func (o LookupEventSourceMappingResultOutput) BisectBatchOnFunctionError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *bool { return v.BisectBatchOnFunctionError }).(pulumi.BoolPtrOutput)
}

// (Streams) An Amazon SQS queue or Amazon SNS topic destination for discarded records.
func (o LookupEventSourceMappingResultOutput) DestinationConfig() EventSourceMappingDestinationConfigPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *EventSourceMappingDestinationConfig {
		return v.DestinationConfig
	}).(EventSourceMappingDestinationConfigPtrOutput)
}

// Disables the event source mapping to pause polling and invocation.
func (o LookupEventSourceMappingResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The filter criteria to control event filtering.
func (o LookupEventSourceMappingResultOutput) FilterCriteria() EventSourceMappingFilterCriteriaPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *EventSourceMappingFilterCriteria { return v.FilterCriteria }).(EventSourceMappingFilterCriteriaPtrOutput)
}

// The name of the Lambda function.
func (o LookupEventSourceMappingResultOutput) FunctionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *string { return v.FunctionName }).(pulumi.StringPtrOutput)
}

// (Streams) A list of response types supported by the function.
func (o LookupEventSourceMappingResultOutput) FunctionResponseTypes() EventSourceMappingFunctionResponseTypesItemArrayOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) []EventSourceMappingFunctionResponseTypesItem {
		return v.FunctionResponseTypes
	}).(EventSourceMappingFunctionResponseTypesItemArrayOutput)
}

// Event Source Mapping Identifier UUID.
func (o LookupEventSourceMappingResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// (Streams) The maximum amount of time to gather records before invoking the function, in seconds.
func (o LookupEventSourceMappingResultOutput) MaximumBatchingWindowInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *int { return v.MaximumBatchingWindowInSeconds }).(pulumi.IntPtrOutput)
}

// (Streams) The maximum age of a record that Lambda sends to a function for processing.
func (o LookupEventSourceMappingResultOutput) MaximumRecordAgeInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *int { return v.MaximumRecordAgeInSeconds }).(pulumi.IntPtrOutput)
}

// (Streams) The maximum number of times to retry when the function returns an error.
func (o LookupEventSourceMappingResultOutput) MaximumRetryAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *int { return v.MaximumRetryAttempts }).(pulumi.IntPtrOutput)
}

// (Streams) The number of batches to process from each shard concurrently.
func (o LookupEventSourceMappingResultOutput) ParallelizationFactor() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *int { return v.ParallelizationFactor }).(pulumi.IntPtrOutput)
}

// (ActiveMQ) A list of ActiveMQ queues.
func (o LookupEventSourceMappingResultOutput) Queues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) []string { return v.Queues }).(pulumi.StringArrayOutput)
}

// A list of SourceAccessConfiguration.
func (o LookupEventSourceMappingResultOutput) SourceAccessConfigurations() EventSourceMappingSourceAccessConfigurationArrayOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) []EventSourceMappingSourceAccessConfiguration {
		return v.SourceAccessConfigurations
	}).(EventSourceMappingSourceAccessConfigurationArrayOutput)
}

// (Kafka) A list of Kafka topics.
func (o LookupEventSourceMappingResultOutput) Topics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) []string { return v.Topics }).(pulumi.StringArrayOutput)
}

// (Streams) Tumbling window (non-overlapping time window) duration to perform aggregations.
func (o LookupEventSourceMappingResultOutput) TumblingWindowInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *int { return v.TumblingWindowInSeconds }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventSourceMappingResultOutput{})
}
