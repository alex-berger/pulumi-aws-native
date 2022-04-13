// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databrew

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::DataBrew::Schedule.
func LookupSchedule(ctx *pulumi.Context, args *LookupScheduleArgs, opts ...pulumi.InvokeOption) (*LookupScheduleResult, error) {
	var rv LookupScheduleResult
	err := ctx.Invoke("aws-native:databrew:getSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduleArgs struct {
	// Schedule Name
	Name string `pulumi:"name"`
}

type LookupScheduleResult struct {
	// Schedule cron
	CronExpression *string  `pulumi:"cronExpression"`
	JobNames       []string `pulumi:"jobNames"`
}

func LookupScheduleOutput(ctx *pulumi.Context, args LookupScheduleOutputArgs, opts ...pulumi.InvokeOption) LookupScheduleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScheduleResult, error) {
			args := v.(LookupScheduleArgs)
			r, err := LookupSchedule(ctx, &args, opts...)
			var s LookupScheduleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScheduleResultOutput)
}

type LookupScheduleOutputArgs struct {
	// Schedule Name
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupScheduleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduleArgs)(nil)).Elem()
}

type LookupScheduleResultOutput struct{ *pulumi.OutputState }

func (LookupScheduleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduleResult)(nil)).Elem()
}

func (o LookupScheduleResultOutput) ToLookupScheduleResultOutput() LookupScheduleResultOutput {
	return o
}

func (o LookupScheduleResultOutput) ToLookupScheduleResultOutputWithContext(ctx context.Context) LookupScheduleResultOutput {
	return o
}

// Schedule cron
func (o LookupScheduleResultOutput) CronExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.CronExpression }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResultOutput) JobNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupScheduleResult) []string { return v.JobNames }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScheduleResultOutput{})
}
