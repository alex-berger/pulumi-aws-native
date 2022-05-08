// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::MonitoringSchedule
func LookupMonitoringSchedule(ctx *pulumi.Context, args *LookupMonitoringScheduleArgs, opts ...pulumi.InvokeOption) (*LookupMonitoringScheduleResult, error) {
	var rv LookupMonitoringScheduleResult
	err := ctx.Invoke("aws-native:sagemaker:getMonitoringSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMonitoringScheduleArgs struct {
	// The Amazon Resource Name (ARN) of the monitoring schedule.
	MonitoringScheduleArn string `pulumi:"monitoringScheduleArn"`
}

type LookupMonitoringScheduleResult struct {
	// The time at which the schedule was created.
	CreationTime *string `pulumi:"creationTime"`
	EndpointName *string `pulumi:"endpointName"`
	// Contains the reason a monitoring job failed, if it failed.
	FailureReason *string `pulumi:"failureReason"`
	// A timestamp that indicates the last time the monitoring job was modified.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// Describes metadata on the last execution to run, if there was one.
	LastMonitoringExecutionSummary *MonitoringScheduleMonitoringExecutionSummary `pulumi:"lastMonitoringExecutionSummary"`
	// The Amazon Resource Name (ARN) of the monitoring schedule.
	MonitoringScheduleArn    *string                   `pulumi:"monitoringScheduleArn"`
	MonitoringScheduleConfig *MonitoringScheduleConfig `pulumi:"monitoringScheduleConfig"`
	// The status of a schedule job.
	MonitoringScheduleStatus *MonitoringScheduleStatus `pulumi:"monitoringScheduleStatus"`
	// An array of key-value pairs to apply to this resource.
	Tags []MonitoringScheduleTag `pulumi:"tags"`
}

func LookupMonitoringScheduleOutput(ctx *pulumi.Context, args LookupMonitoringScheduleOutputArgs, opts ...pulumi.InvokeOption) LookupMonitoringScheduleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMonitoringScheduleResult, error) {
			args := v.(LookupMonitoringScheduleArgs)
			r, err := LookupMonitoringSchedule(ctx, &args, opts...)
			var s LookupMonitoringScheduleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMonitoringScheduleResultOutput)
}

type LookupMonitoringScheduleOutputArgs struct {
	// The Amazon Resource Name (ARN) of the monitoring schedule.
	MonitoringScheduleArn pulumi.StringInput `pulumi:"monitoringScheduleArn"`
}

func (LookupMonitoringScheduleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMonitoringScheduleArgs)(nil)).Elem()
}

type LookupMonitoringScheduleResultOutput struct{ *pulumi.OutputState }

func (LookupMonitoringScheduleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMonitoringScheduleResult)(nil)).Elem()
}

func (o LookupMonitoringScheduleResultOutput) ToLookupMonitoringScheduleResultOutput() LookupMonitoringScheduleResultOutput {
	return o
}

func (o LookupMonitoringScheduleResultOutput) ToLookupMonitoringScheduleResultOutputWithContext(ctx context.Context) LookupMonitoringScheduleResultOutput {
	return o
}

// The time at which the schedule was created.
func (o LookupMonitoringScheduleResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMonitoringScheduleResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o LookupMonitoringScheduleResultOutput) EndpointName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMonitoringScheduleResult) *string { return v.EndpointName }).(pulumi.StringPtrOutput)
}

// Contains the reason a monitoring job failed, if it failed.
func (o LookupMonitoringScheduleResultOutput) FailureReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMonitoringScheduleResult) *string { return v.FailureReason }).(pulumi.StringPtrOutput)
}

// A timestamp that indicates the last time the monitoring job was modified.
func (o LookupMonitoringScheduleResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMonitoringScheduleResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

// Describes metadata on the last execution to run, if there was one.
func (o LookupMonitoringScheduleResultOutput) LastMonitoringExecutionSummary() MonitoringScheduleMonitoringExecutionSummaryPtrOutput {
	return o.ApplyT(func(v LookupMonitoringScheduleResult) *MonitoringScheduleMonitoringExecutionSummary {
		return v.LastMonitoringExecutionSummary
	}).(MonitoringScheduleMonitoringExecutionSummaryPtrOutput)
}

// The Amazon Resource Name (ARN) of the monitoring schedule.
func (o LookupMonitoringScheduleResultOutput) MonitoringScheduleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMonitoringScheduleResult) *string { return v.MonitoringScheduleArn }).(pulumi.StringPtrOutput)
}

func (o LookupMonitoringScheduleResultOutput) MonitoringScheduleConfig() MonitoringScheduleConfigPtrOutput {
	return o.ApplyT(func(v LookupMonitoringScheduleResult) *MonitoringScheduleConfig { return v.MonitoringScheduleConfig }).(MonitoringScheduleConfigPtrOutput)
}

// The status of a schedule job.
func (o LookupMonitoringScheduleResultOutput) MonitoringScheduleStatus() MonitoringScheduleStatusPtrOutput {
	return o.ApplyT(func(v LookupMonitoringScheduleResult) *MonitoringScheduleStatus { return v.MonitoringScheduleStatus }).(MonitoringScheduleStatusPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupMonitoringScheduleResultOutput) Tags() MonitoringScheduleTagArrayOutput {
	return o.ApplyT(func(v LookupMonitoringScheduleResult) []MonitoringScheduleTag { return v.Tags }).(MonitoringScheduleTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMonitoringScheduleResultOutput{})
}
