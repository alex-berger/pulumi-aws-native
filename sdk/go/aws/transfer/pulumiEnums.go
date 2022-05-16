// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package transfer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A flag that indicates whether or not to overwrite an existing file of the same name. The default is FALSE.
type WorkflowStepCopyStepDetailsPropertiesOverwriteExisting string

const (
	WorkflowStepCopyStepDetailsPropertiesOverwriteExistingTrue  = WorkflowStepCopyStepDetailsPropertiesOverwriteExisting("TRUE")
	WorkflowStepCopyStepDetailsPropertiesOverwriteExistingFalse = WorkflowStepCopyStepDetailsPropertiesOverwriteExisting("FALSE")
)

func (WorkflowStepCopyStepDetailsPropertiesOverwriteExisting) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowStepCopyStepDetailsPropertiesOverwriteExisting)(nil)).Elem()
}

func (e WorkflowStepCopyStepDetailsPropertiesOverwriteExisting) ToWorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput() WorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput {
	return pulumi.ToOutput(e).(WorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput)
}

func (e WorkflowStepCopyStepDetailsPropertiesOverwriteExisting) ToWorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutputWithContext(ctx context.Context) WorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput)
}

func (e WorkflowStepCopyStepDetailsPropertiesOverwriteExisting) ToWorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput() WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput {
	return e.ToWorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutputWithContext(context.Background())
}

func (e WorkflowStepCopyStepDetailsPropertiesOverwriteExisting) ToWorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutputWithContext(ctx context.Context) WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput {
	return WorkflowStepCopyStepDetailsPropertiesOverwriteExisting(e).ToWorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutputWithContext(ctx).ToWorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutputWithContext(ctx)
}

func (e WorkflowStepCopyStepDetailsPropertiesOverwriteExisting) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkflowStepCopyStepDetailsPropertiesOverwriteExisting) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkflowStepCopyStepDetailsPropertiesOverwriteExisting) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WorkflowStepCopyStepDetailsPropertiesOverwriteExisting) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput struct{ *pulumi.OutputState }

func (WorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowStepCopyStepDetailsPropertiesOverwriteExisting)(nil)).Elem()
}

func (o WorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput) ToWorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput() WorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput {
	return o
}

func (o WorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput) ToWorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutputWithContext(ctx context.Context) WorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput {
	return o
}

func (o WorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput) ToWorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput() WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput {
	return o.ToWorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutputWithContext(context.Background())
}

func (o WorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput) ToWorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutputWithContext(ctx context.Context) WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkflowStepCopyStepDetailsPropertiesOverwriteExisting) *WorkflowStepCopyStepDetailsPropertiesOverwriteExisting {
		return &v
	}).(WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput)
}

func (o WorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkflowStepCopyStepDetailsPropertiesOverwriteExisting) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkflowStepCopyStepDetailsPropertiesOverwriteExisting) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput struct{ *pulumi.OutputState }

func (WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowStepCopyStepDetailsPropertiesOverwriteExisting)(nil)).Elem()
}

func (o WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput) ToWorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput() WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput {
	return o
}

func (o WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput) ToWorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutputWithContext(ctx context.Context) WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput {
	return o
}

func (o WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput) Elem() WorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput {
	return o.ApplyT(func(v *WorkflowStepCopyStepDetailsPropertiesOverwriteExisting) WorkflowStepCopyStepDetailsPropertiesOverwriteExisting {
		if v != nil {
			return *v
		}
		var ret WorkflowStepCopyStepDetailsPropertiesOverwriteExisting
		return ret
	}).(WorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput)
}

func (o WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WorkflowStepCopyStepDetailsPropertiesOverwriteExisting) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// WorkflowStepCopyStepDetailsPropertiesOverwriteExistingInput is an input type that accepts WorkflowStepCopyStepDetailsPropertiesOverwriteExistingArgs and WorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput values.
// You can construct a concrete instance of `WorkflowStepCopyStepDetailsPropertiesOverwriteExistingInput` via:
//
//          WorkflowStepCopyStepDetailsPropertiesOverwriteExistingArgs{...}
type WorkflowStepCopyStepDetailsPropertiesOverwriteExistingInput interface {
	pulumi.Input

	ToWorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput() WorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput
	ToWorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutputWithContext(context.Context) WorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput
}

var workflowStepCopyStepDetailsPropertiesOverwriteExistingPtrType = reflect.TypeOf((**WorkflowStepCopyStepDetailsPropertiesOverwriteExisting)(nil)).Elem()

type WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrInput interface {
	pulumi.Input

	ToWorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput() WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput
	ToWorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutputWithContext(context.Context) WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput
}

type workflowStepCopyStepDetailsPropertiesOverwriteExistingPtr string

func WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtr(v string) WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrInput {
	return (*workflowStepCopyStepDetailsPropertiesOverwriteExistingPtr)(&v)
}

func (*workflowStepCopyStepDetailsPropertiesOverwriteExistingPtr) ElementType() reflect.Type {
	return workflowStepCopyStepDetailsPropertiesOverwriteExistingPtrType
}

func (in *workflowStepCopyStepDetailsPropertiesOverwriteExistingPtr) ToWorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput() WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput {
	return pulumi.ToOutput(in).(WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput)
}

func (in *workflowStepCopyStepDetailsPropertiesOverwriteExistingPtr) ToWorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutputWithContext(ctx context.Context) WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput)
}

type WorkflowStepType string

const (
	WorkflowStepTypeCopy   = WorkflowStepType("COPY")
	WorkflowStepTypeCustom = WorkflowStepType("CUSTOM")
	WorkflowStepTypeDelete = WorkflowStepType("DELETE")
	WorkflowStepTypeTag    = WorkflowStepType("TAG")
)

func (WorkflowStepType) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowStepType)(nil)).Elem()
}

func (e WorkflowStepType) ToWorkflowStepTypeOutput() WorkflowStepTypeOutput {
	return pulumi.ToOutput(e).(WorkflowStepTypeOutput)
}

func (e WorkflowStepType) ToWorkflowStepTypeOutputWithContext(ctx context.Context) WorkflowStepTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WorkflowStepTypeOutput)
}

func (e WorkflowStepType) ToWorkflowStepTypePtrOutput() WorkflowStepTypePtrOutput {
	return e.ToWorkflowStepTypePtrOutputWithContext(context.Background())
}

func (e WorkflowStepType) ToWorkflowStepTypePtrOutputWithContext(ctx context.Context) WorkflowStepTypePtrOutput {
	return WorkflowStepType(e).ToWorkflowStepTypeOutputWithContext(ctx).ToWorkflowStepTypePtrOutputWithContext(ctx)
}

func (e WorkflowStepType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkflowStepType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkflowStepType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WorkflowStepType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WorkflowStepTypeOutput struct{ *pulumi.OutputState }

func (WorkflowStepTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowStepType)(nil)).Elem()
}

func (o WorkflowStepTypeOutput) ToWorkflowStepTypeOutput() WorkflowStepTypeOutput {
	return o
}

func (o WorkflowStepTypeOutput) ToWorkflowStepTypeOutputWithContext(ctx context.Context) WorkflowStepTypeOutput {
	return o
}

func (o WorkflowStepTypeOutput) ToWorkflowStepTypePtrOutput() WorkflowStepTypePtrOutput {
	return o.ToWorkflowStepTypePtrOutputWithContext(context.Background())
}

func (o WorkflowStepTypeOutput) ToWorkflowStepTypePtrOutputWithContext(ctx context.Context) WorkflowStepTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkflowStepType) *WorkflowStepType {
		return &v
	}).(WorkflowStepTypePtrOutput)
}

func (o WorkflowStepTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WorkflowStepTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkflowStepType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WorkflowStepTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkflowStepTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkflowStepType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WorkflowStepTypePtrOutput struct{ *pulumi.OutputState }

func (WorkflowStepTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowStepType)(nil)).Elem()
}

func (o WorkflowStepTypePtrOutput) ToWorkflowStepTypePtrOutput() WorkflowStepTypePtrOutput {
	return o
}

func (o WorkflowStepTypePtrOutput) ToWorkflowStepTypePtrOutputWithContext(ctx context.Context) WorkflowStepTypePtrOutput {
	return o
}

func (o WorkflowStepTypePtrOutput) Elem() WorkflowStepTypeOutput {
	return o.ApplyT(func(v *WorkflowStepType) WorkflowStepType {
		if v != nil {
			return *v
		}
		var ret WorkflowStepType
		return ret
	}).(WorkflowStepTypeOutput)
}

func (o WorkflowStepTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkflowStepTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WorkflowStepType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// WorkflowStepTypeInput is an input type that accepts WorkflowStepTypeArgs and WorkflowStepTypeOutput values.
// You can construct a concrete instance of `WorkflowStepTypeInput` via:
//
//          WorkflowStepTypeArgs{...}
type WorkflowStepTypeInput interface {
	pulumi.Input

	ToWorkflowStepTypeOutput() WorkflowStepTypeOutput
	ToWorkflowStepTypeOutputWithContext(context.Context) WorkflowStepTypeOutput
}

var workflowStepTypePtrType = reflect.TypeOf((**WorkflowStepType)(nil)).Elem()

type WorkflowStepTypePtrInput interface {
	pulumi.Input

	ToWorkflowStepTypePtrOutput() WorkflowStepTypePtrOutput
	ToWorkflowStepTypePtrOutputWithContext(context.Context) WorkflowStepTypePtrOutput
}

type workflowStepTypePtr string

func WorkflowStepTypePtr(v string) WorkflowStepTypePtrInput {
	return (*workflowStepTypePtr)(&v)
}

func (*workflowStepTypePtr) ElementType() reflect.Type {
	return workflowStepTypePtrType
}

func (in *workflowStepTypePtr) ToWorkflowStepTypePtrOutput() WorkflowStepTypePtrOutput {
	return pulumi.ToOutput(in).(WorkflowStepTypePtrOutput)
}

func (in *workflowStepTypePtr) ToWorkflowStepTypePtrOutputWithContext(ctx context.Context) WorkflowStepTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WorkflowStepTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowStepCopyStepDetailsPropertiesOverwriteExistingInput)(nil)).Elem(), WorkflowStepCopyStepDetailsPropertiesOverwriteExisting("TRUE"))
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrInput)(nil)).Elem(), WorkflowStepCopyStepDetailsPropertiesOverwriteExisting("TRUE"))
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowStepTypeInput)(nil)).Elem(), WorkflowStepType("COPY"))
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowStepTypePtrInput)(nil)).Elem(), WorkflowStepType("COPY"))
	pulumi.RegisterOutputType(WorkflowStepCopyStepDetailsPropertiesOverwriteExistingOutput{})
	pulumi.RegisterOutputType(WorkflowStepCopyStepDetailsPropertiesOverwriteExistingPtrOutput{})
	pulumi.RegisterOutputType(WorkflowStepTypeOutput{})
	pulumi.RegisterOutputType(WorkflowStepTypePtrOutput{})
}