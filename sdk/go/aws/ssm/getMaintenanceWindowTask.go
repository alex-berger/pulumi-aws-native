// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SSM::MaintenanceWindowTask
func LookupMaintenanceWindowTask(ctx *pulumi.Context, args *LookupMaintenanceWindowTaskArgs, opts ...pulumi.InvokeOption) (*LookupMaintenanceWindowTaskResult, error) {
	var rv LookupMaintenanceWindowTaskResult
	err := ctx.Invoke("aws-native:ssm:getMaintenanceWindowTask", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMaintenanceWindowTaskArgs struct {
	Id string `pulumi:"id"`
}

type LookupMaintenanceWindowTaskResult struct {
	CutoffBehavior           *string                                        `pulumi:"cutoffBehavior"`
	Description              *string                                        `pulumi:"description"`
	Id                       *string                                        `pulumi:"id"`
	LoggingInfo              *MaintenanceWindowTaskLoggingInfo              `pulumi:"loggingInfo"`
	MaxConcurrency           *string                                        `pulumi:"maxConcurrency"`
	MaxErrors                *string                                        `pulumi:"maxErrors"`
	Name                     *string                                        `pulumi:"name"`
	Priority                 *int                                           `pulumi:"priority"`
	ServiceRoleArn           *string                                        `pulumi:"serviceRoleArn"`
	Targets                  []MaintenanceWindowTaskTarget                  `pulumi:"targets"`
	TaskArn                  *string                                        `pulumi:"taskArn"`
	TaskInvocationParameters *MaintenanceWindowTaskTaskInvocationParameters `pulumi:"taskInvocationParameters"`
	TaskParameters           interface{}                                    `pulumi:"taskParameters"`
}

func LookupMaintenanceWindowTaskOutput(ctx *pulumi.Context, args LookupMaintenanceWindowTaskOutputArgs, opts ...pulumi.InvokeOption) LookupMaintenanceWindowTaskResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMaintenanceWindowTaskResult, error) {
			args := v.(LookupMaintenanceWindowTaskArgs)
			r, err := LookupMaintenanceWindowTask(ctx, &args, opts...)
			var s LookupMaintenanceWindowTaskResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMaintenanceWindowTaskResultOutput)
}

type LookupMaintenanceWindowTaskOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupMaintenanceWindowTaskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMaintenanceWindowTaskArgs)(nil)).Elem()
}

type LookupMaintenanceWindowTaskResultOutput struct{ *pulumi.OutputState }

func (LookupMaintenanceWindowTaskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMaintenanceWindowTaskResult)(nil)).Elem()
}

func (o LookupMaintenanceWindowTaskResultOutput) ToLookupMaintenanceWindowTaskResultOutput() LookupMaintenanceWindowTaskResultOutput {
	return o
}

func (o LookupMaintenanceWindowTaskResultOutput) ToLookupMaintenanceWindowTaskResultOutputWithContext(ctx context.Context) LookupMaintenanceWindowTaskResultOutput {
	return o
}

func (o LookupMaintenanceWindowTaskResultOutput) CutoffBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowTaskResult) *string { return v.CutoffBehavior }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceWindowTaskResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowTaskResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceWindowTaskResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowTaskResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceWindowTaskResultOutput) LoggingInfo() MaintenanceWindowTaskLoggingInfoPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowTaskResult) *MaintenanceWindowTaskLoggingInfo { return v.LoggingInfo }).(MaintenanceWindowTaskLoggingInfoPtrOutput)
}

func (o LookupMaintenanceWindowTaskResultOutput) MaxConcurrency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowTaskResult) *string { return v.MaxConcurrency }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceWindowTaskResultOutput) MaxErrors() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowTaskResult) *string { return v.MaxErrors }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceWindowTaskResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowTaskResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceWindowTaskResultOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowTaskResult) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o LookupMaintenanceWindowTaskResultOutput) ServiceRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowTaskResult) *string { return v.ServiceRoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceWindowTaskResultOutput) Targets() MaintenanceWindowTaskTargetArrayOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowTaskResult) []MaintenanceWindowTaskTarget { return v.Targets }).(MaintenanceWindowTaskTargetArrayOutput)
}

func (o LookupMaintenanceWindowTaskResultOutput) TaskArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowTaskResult) *string { return v.TaskArn }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceWindowTaskResultOutput) TaskInvocationParameters() MaintenanceWindowTaskTaskInvocationParametersPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowTaskResult) *MaintenanceWindowTaskTaskInvocationParameters {
		return v.TaskInvocationParameters
	}).(MaintenanceWindowTaskTaskInvocationParametersPtrOutput)
}

func (o LookupMaintenanceWindowTaskResultOutput) TaskParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowTaskResult) interface{} { return v.TaskParameters }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMaintenanceWindowTaskResultOutput{})
}
