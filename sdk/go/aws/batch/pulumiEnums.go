// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package batch

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type JobQueueStateEnum string

const (
	JobQueueStateEnumDisabled = JobQueueStateEnum("DISABLED")
	JobQueueStateEnumEnabled  = JobQueueStateEnum("ENABLED")
)

func (JobQueueStateEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*JobQueueStateEnum)(nil)).Elem()
}

func (e JobQueueStateEnum) ToJobQueueStateEnumOutput() JobQueueStateEnumOutput {
	return pulumi.ToOutput(e).(JobQueueStateEnumOutput)
}

func (e JobQueueStateEnum) ToJobQueueStateEnumOutputWithContext(ctx context.Context) JobQueueStateEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(JobQueueStateEnumOutput)
}

func (e JobQueueStateEnum) ToJobQueueStateEnumPtrOutput() JobQueueStateEnumPtrOutput {
	return e.ToJobQueueStateEnumPtrOutputWithContext(context.Background())
}

func (e JobQueueStateEnum) ToJobQueueStateEnumPtrOutputWithContext(ctx context.Context) JobQueueStateEnumPtrOutput {
	return JobQueueStateEnum(e).ToJobQueueStateEnumOutputWithContext(ctx).ToJobQueueStateEnumPtrOutputWithContext(ctx)
}

func (e JobQueueStateEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobQueueStateEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobQueueStateEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e JobQueueStateEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type JobQueueStateEnumOutput struct{ *pulumi.OutputState }

func (JobQueueStateEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobQueueStateEnum)(nil)).Elem()
}

func (o JobQueueStateEnumOutput) ToJobQueueStateEnumOutput() JobQueueStateEnumOutput {
	return o
}

func (o JobQueueStateEnumOutput) ToJobQueueStateEnumOutputWithContext(ctx context.Context) JobQueueStateEnumOutput {
	return o
}

func (o JobQueueStateEnumOutput) ToJobQueueStateEnumPtrOutput() JobQueueStateEnumPtrOutput {
	return o.ToJobQueueStateEnumPtrOutputWithContext(context.Background())
}

func (o JobQueueStateEnumOutput) ToJobQueueStateEnumPtrOutputWithContext(ctx context.Context) JobQueueStateEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobQueueStateEnum) *JobQueueStateEnum {
		return &v
	}).(JobQueueStateEnumPtrOutput)
}

func (o JobQueueStateEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o JobQueueStateEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e JobQueueStateEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o JobQueueStateEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o JobQueueStateEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e JobQueueStateEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type JobQueueStateEnumPtrOutput struct{ *pulumi.OutputState }

func (JobQueueStateEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobQueueStateEnum)(nil)).Elem()
}

func (o JobQueueStateEnumPtrOutput) ToJobQueueStateEnumPtrOutput() JobQueueStateEnumPtrOutput {
	return o
}

func (o JobQueueStateEnumPtrOutput) ToJobQueueStateEnumPtrOutputWithContext(ctx context.Context) JobQueueStateEnumPtrOutput {
	return o
}

func (o JobQueueStateEnumPtrOutput) Elem() JobQueueStateEnumOutput {
	return o.ApplyT(func(v *JobQueueStateEnum) JobQueueStateEnum {
		if v != nil {
			return *v
		}
		var ret JobQueueStateEnum
		return ret
	}).(JobQueueStateEnumOutput)
}

func (o JobQueueStateEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o JobQueueStateEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *JobQueueStateEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// JobQueueStateEnumInput is an input type that accepts JobQueueStateEnumArgs and JobQueueStateEnumOutput values.
// You can construct a concrete instance of `JobQueueStateEnumInput` via:
//
//          JobQueueStateEnumArgs{...}
type JobQueueStateEnumInput interface {
	pulumi.Input

	ToJobQueueStateEnumOutput() JobQueueStateEnumOutput
	ToJobQueueStateEnumOutputWithContext(context.Context) JobQueueStateEnumOutput
}

var jobQueueStateEnumPtrType = reflect.TypeOf((**JobQueueStateEnum)(nil)).Elem()

type JobQueueStateEnumPtrInput interface {
	pulumi.Input

	ToJobQueueStateEnumPtrOutput() JobQueueStateEnumPtrOutput
	ToJobQueueStateEnumPtrOutputWithContext(context.Context) JobQueueStateEnumPtrOutput
}

type jobQueueStateEnumPtr string

func JobQueueStateEnumPtr(v string) JobQueueStateEnumPtrInput {
	return (*jobQueueStateEnumPtr)(&v)
}

func (*jobQueueStateEnumPtr) ElementType() reflect.Type {
	return jobQueueStateEnumPtrType
}

func (in *jobQueueStateEnumPtr) ToJobQueueStateEnumPtrOutput() JobQueueStateEnumPtrOutput {
	return pulumi.ToOutput(in).(JobQueueStateEnumPtrOutput)
}

func (in *jobQueueStateEnumPtr) ToJobQueueStateEnumPtrOutputWithContext(ctx context.Context) JobQueueStateEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(JobQueueStateEnumPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JobQueueStateEnumInput)(nil)).Elem(), JobQueueStateEnum("DISABLED"))
	pulumi.RegisterInputType(reflect.TypeOf((*JobQueueStateEnumPtrInput)(nil)).Elem(), JobQueueStateEnum("DISABLED"))
	pulumi.RegisterOutputType(JobQueueStateEnumOutput{})
	pulumi.RegisterOutputType(JobQueueStateEnumPtrOutput{})
}
