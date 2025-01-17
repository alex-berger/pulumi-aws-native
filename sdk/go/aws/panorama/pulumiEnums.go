// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package panorama

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationInstanceHealthStatus string

const (
	ApplicationInstanceHealthStatusRunning      = ApplicationInstanceHealthStatus("RUNNING")
	ApplicationInstanceHealthStatusError        = ApplicationInstanceHealthStatus("ERROR")
	ApplicationInstanceHealthStatusNotAvailable = ApplicationInstanceHealthStatus("NOT_AVAILABLE")
)

type ApplicationInstanceHealthStatusOutput struct{ *pulumi.OutputState }

func (ApplicationInstanceHealthStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInstanceHealthStatus)(nil)).Elem()
}

func (o ApplicationInstanceHealthStatusOutput) ToApplicationInstanceHealthStatusOutput() ApplicationInstanceHealthStatusOutput {
	return o
}

func (o ApplicationInstanceHealthStatusOutput) ToApplicationInstanceHealthStatusOutputWithContext(ctx context.Context) ApplicationInstanceHealthStatusOutput {
	return o
}

func (o ApplicationInstanceHealthStatusOutput) ToApplicationInstanceHealthStatusPtrOutput() ApplicationInstanceHealthStatusPtrOutput {
	return o.ToApplicationInstanceHealthStatusPtrOutputWithContext(context.Background())
}

func (o ApplicationInstanceHealthStatusOutput) ToApplicationInstanceHealthStatusPtrOutputWithContext(ctx context.Context) ApplicationInstanceHealthStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationInstanceHealthStatus) *ApplicationInstanceHealthStatus {
		return &v
	}).(ApplicationInstanceHealthStatusPtrOutput)
}

func (o ApplicationInstanceHealthStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApplicationInstanceHealthStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationInstanceHealthStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApplicationInstanceHealthStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationInstanceHealthStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationInstanceHealthStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationInstanceHealthStatusPtrOutput struct{ *pulumi.OutputState }

func (ApplicationInstanceHealthStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInstanceHealthStatus)(nil)).Elem()
}

func (o ApplicationInstanceHealthStatusPtrOutput) ToApplicationInstanceHealthStatusPtrOutput() ApplicationInstanceHealthStatusPtrOutput {
	return o
}

func (o ApplicationInstanceHealthStatusPtrOutput) ToApplicationInstanceHealthStatusPtrOutputWithContext(ctx context.Context) ApplicationInstanceHealthStatusPtrOutput {
	return o
}

func (o ApplicationInstanceHealthStatusPtrOutput) Elem() ApplicationInstanceHealthStatusOutput {
	return o.ApplyT(func(v *ApplicationInstanceHealthStatus) ApplicationInstanceHealthStatus {
		if v != nil {
			return *v
		}
		var ret ApplicationInstanceHealthStatus
		return ret
	}).(ApplicationInstanceHealthStatusOutput)
}

func (o ApplicationInstanceHealthStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationInstanceHealthStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApplicationInstanceHealthStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationInstanceStatus string

const (
	ApplicationInstanceStatusDeploymentPending    = ApplicationInstanceStatus("DEPLOYMENT_PENDING")
	ApplicationInstanceStatusDeploymentRequested  = ApplicationInstanceStatus("DEPLOYMENT_REQUESTED")
	ApplicationInstanceStatusDeploymentInProgress = ApplicationInstanceStatus("DEPLOYMENT_IN_PROGRESS")
	ApplicationInstanceStatusDeploymentError      = ApplicationInstanceStatus("DEPLOYMENT_ERROR")
	ApplicationInstanceStatusDeploymentSucceeded  = ApplicationInstanceStatus("DEPLOYMENT_SUCCEEDED")
	ApplicationInstanceStatusRemovalPending       = ApplicationInstanceStatus("REMOVAL_PENDING")
	ApplicationInstanceStatusRemovalRequested     = ApplicationInstanceStatus("REMOVAL_REQUESTED")
	ApplicationInstanceStatusRemovalInProgress    = ApplicationInstanceStatus("REMOVAL_IN_PROGRESS")
	ApplicationInstanceStatusRemovalFailed        = ApplicationInstanceStatus("REMOVAL_FAILED")
	ApplicationInstanceStatusRemovalSucceeded     = ApplicationInstanceStatus("REMOVAL_SUCCEEDED")
)

type ApplicationInstanceStatusOutput struct{ *pulumi.OutputState }

func (ApplicationInstanceStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInstanceStatus)(nil)).Elem()
}

func (o ApplicationInstanceStatusOutput) ToApplicationInstanceStatusOutput() ApplicationInstanceStatusOutput {
	return o
}

func (o ApplicationInstanceStatusOutput) ToApplicationInstanceStatusOutputWithContext(ctx context.Context) ApplicationInstanceStatusOutput {
	return o
}

func (o ApplicationInstanceStatusOutput) ToApplicationInstanceStatusPtrOutput() ApplicationInstanceStatusPtrOutput {
	return o.ToApplicationInstanceStatusPtrOutputWithContext(context.Background())
}

func (o ApplicationInstanceStatusOutput) ToApplicationInstanceStatusPtrOutputWithContext(ctx context.Context) ApplicationInstanceStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationInstanceStatus) *ApplicationInstanceStatus {
		return &v
	}).(ApplicationInstanceStatusPtrOutput)
}

func (o ApplicationInstanceStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApplicationInstanceStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationInstanceStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApplicationInstanceStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationInstanceStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationInstanceStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationInstanceStatusPtrOutput struct{ *pulumi.OutputState }

func (ApplicationInstanceStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInstanceStatus)(nil)).Elem()
}

func (o ApplicationInstanceStatusPtrOutput) ToApplicationInstanceStatusPtrOutput() ApplicationInstanceStatusPtrOutput {
	return o
}

func (o ApplicationInstanceStatusPtrOutput) ToApplicationInstanceStatusPtrOutputWithContext(ctx context.Context) ApplicationInstanceStatusPtrOutput {
	return o
}

func (o ApplicationInstanceStatusPtrOutput) Elem() ApplicationInstanceStatusOutput {
	return o.ApplyT(func(v *ApplicationInstanceStatus) ApplicationInstanceStatus {
		if v != nil {
			return *v
		}
		var ret ApplicationInstanceStatus
		return ret
	}).(ApplicationInstanceStatusOutput)
}

func (o ApplicationInstanceStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationInstanceStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApplicationInstanceStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationInstanceStatusFilter string

const (
	ApplicationInstanceStatusFilterDeploymentSucceeded  = ApplicationInstanceStatusFilter("DEPLOYMENT_SUCCEEDED")
	ApplicationInstanceStatusFilterDeploymentError      = ApplicationInstanceStatusFilter("DEPLOYMENT_ERROR")
	ApplicationInstanceStatusFilterRemovalSucceeded     = ApplicationInstanceStatusFilter("REMOVAL_SUCCEEDED")
	ApplicationInstanceStatusFilterRemovalFailed        = ApplicationInstanceStatusFilter("REMOVAL_FAILED")
	ApplicationInstanceStatusFilterProcessingDeployment = ApplicationInstanceStatusFilter("PROCESSING_DEPLOYMENT")
	ApplicationInstanceStatusFilterProcessingRemoval    = ApplicationInstanceStatusFilter("PROCESSING_REMOVAL")
)

func (ApplicationInstanceStatusFilter) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInstanceStatusFilter)(nil)).Elem()
}

func (e ApplicationInstanceStatusFilter) ToApplicationInstanceStatusFilterOutput() ApplicationInstanceStatusFilterOutput {
	return pulumi.ToOutput(e).(ApplicationInstanceStatusFilterOutput)
}

func (e ApplicationInstanceStatusFilter) ToApplicationInstanceStatusFilterOutputWithContext(ctx context.Context) ApplicationInstanceStatusFilterOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApplicationInstanceStatusFilterOutput)
}

func (e ApplicationInstanceStatusFilter) ToApplicationInstanceStatusFilterPtrOutput() ApplicationInstanceStatusFilterPtrOutput {
	return e.ToApplicationInstanceStatusFilterPtrOutputWithContext(context.Background())
}

func (e ApplicationInstanceStatusFilter) ToApplicationInstanceStatusFilterPtrOutputWithContext(ctx context.Context) ApplicationInstanceStatusFilterPtrOutput {
	return ApplicationInstanceStatusFilter(e).ToApplicationInstanceStatusFilterOutputWithContext(ctx).ToApplicationInstanceStatusFilterPtrOutputWithContext(ctx)
}

func (e ApplicationInstanceStatusFilter) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationInstanceStatusFilter) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationInstanceStatusFilter) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApplicationInstanceStatusFilter) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ApplicationInstanceStatusFilterOutput struct{ *pulumi.OutputState }

func (ApplicationInstanceStatusFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInstanceStatusFilter)(nil)).Elem()
}

func (o ApplicationInstanceStatusFilterOutput) ToApplicationInstanceStatusFilterOutput() ApplicationInstanceStatusFilterOutput {
	return o
}

func (o ApplicationInstanceStatusFilterOutput) ToApplicationInstanceStatusFilterOutputWithContext(ctx context.Context) ApplicationInstanceStatusFilterOutput {
	return o
}

func (o ApplicationInstanceStatusFilterOutput) ToApplicationInstanceStatusFilterPtrOutput() ApplicationInstanceStatusFilterPtrOutput {
	return o.ToApplicationInstanceStatusFilterPtrOutputWithContext(context.Background())
}

func (o ApplicationInstanceStatusFilterOutput) ToApplicationInstanceStatusFilterPtrOutputWithContext(ctx context.Context) ApplicationInstanceStatusFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationInstanceStatusFilter) *ApplicationInstanceStatusFilter {
		return &v
	}).(ApplicationInstanceStatusFilterPtrOutput)
}

func (o ApplicationInstanceStatusFilterOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApplicationInstanceStatusFilterOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationInstanceStatusFilter) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApplicationInstanceStatusFilterOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationInstanceStatusFilterOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationInstanceStatusFilter) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationInstanceStatusFilterPtrOutput struct{ *pulumi.OutputState }

func (ApplicationInstanceStatusFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInstanceStatusFilter)(nil)).Elem()
}

func (o ApplicationInstanceStatusFilterPtrOutput) ToApplicationInstanceStatusFilterPtrOutput() ApplicationInstanceStatusFilterPtrOutput {
	return o
}

func (o ApplicationInstanceStatusFilterPtrOutput) ToApplicationInstanceStatusFilterPtrOutputWithContext(ctx context.Context) ApplicationInstanceStatusFilterPtrOutput {
	return o
}

func (o ApplicationInstanceStatusFilterPtrOutput) Elem() ApplicationInstanceStatusFilterOutput {
	return o.ApplyT(func(v *ApplicationInstanceStatusFilter) ApplicationInstanceStatusFilter {
		if v != nil {
			return *v
		}
		var ret ApplicationInstanceStatusFilter
		return ret
	}).(ApplicationInstanceStatusFilterOutput)
}

func (o ApplicationInstanceStatusFilterPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationInstanceStatusFilterPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApplicationInstanceStatusFilter) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ApplicationInstanceStatusFilterInput is an input type that accepts ApplicationInstanceStatusFilterArgs and ApplicationInstanceStatusFilterOutput values.
// You can construct a concrete instance of `ApplicationInstanceStatusFilterInput` via:
//
//          ApplicationInstanceStatusFilterArgs{...}
type ApplicationInstanceStatusFilterInput interface {
	pulumi.Input

	ToApplicationInstanceStatusFilterOutput() ApplicationInstanceStatusFilterOutput
	ToApplicationInstanceStatusFilterOutputWithContext(context.Context) ApplicationInstanceStatusFilterOutput
}

var applicationInstanceStatusFilterPtrType = reflect.TypeOf((**ApplicationInstanceStatusFilter)(nil)).Elem()

type ApplicationInstanceStatusFilterPtrInput interface {
	pulumi.Input

	ToApplicationInstanceStatusFilterPtrOutput() ApplicationInstanceStatusFilterPtrOutput
	ToApplicationInstanceStatusFilterPtrOutputWithContext(context.Context) ApplicationInstanceStatusFilterPtrOutput
}

type applicationInstanceStatusFilterPtr string

func ApplicationInstanceStatusFilterPtr(v string) ApplicationInstanceStatusFilterPtrInput {
	return (*applicationInstanceStatusFilterPtr)(&v)
}

func (*applicationInstanceStatusFilterPtr) ElementType() reflect.Type {
	return applicationInstanceStatusFilterPtrType
}

func (in *applicationInstanceStatusFilterPtr) ToApplicationInstanceStatusFilterPtrOutput() ApplicationInstanceStatusFilterPtrOutput {
	return pulumi.ToOutput(in).(ApplicationInstanceStatusFilterPtrOutput)
}

func (in *applicationInstanceStatusFilterPtr) ToApplicationInstanceStatusFilterPtrOutputWithContext(ctx context.Context) ApplicationInstanceStatusFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApplicationInstanceStatusFilterPtrOutput)
}

type PackageVersionStatus string

const (
	PackageVersionStatusRegisterPending   = PackageVersionStatus("REGISTER_PENDING")
	PackageVersionStatusRegisterCompleted = PackageVersionStatus("REGISTER_COMPLETED")
	PackageVersionStatusFailed            = PackageVersionStatus("FAILED")
	PackageVersionStatusDeleting          = PackageVersionStatus("DELETING")
)

type PackageVersionStatusOutput struct{ *pulumi.OutputState }

func (PackageVersionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PackageVersionStatus)(nil)).Elem()
}

func (o PackageVersionStatusOutput) ToPackageVersionStatusOutput() PackageVersionStatusOutput {
	return o
}

func (o PackageVersionStatusOutput) ToPackageVersionStatusOutputWithContext(ctx context.Context) PackageVersionStatusOutput {
	return o
}

func (o PackageVersionStatusOutput) ToPackageVersionStatusPtrOutput() PackageVersionStatusPtrOutput {
	return o.ToPackageVersionStatusPtrOutputWithContext(context.Background())
}

func (o PackageVersionStatusOutput) ToPackageVersionStatusPtrOutputWithContext(ctx context.Context) PackageVersionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PackageVersionStatus) *PackageVersionStatus {
		return &v
	}).(PackageVersionStatusPtrOutput)
}

func (o PackageVersionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PackageVersionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PackageVersionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PackageVersionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PackageVersionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PackageVersionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PackageVersionStatusPtrOutput struct{ *pulumi.OutputState }

func (PackageVersionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PackageVersionStatus)(nil)).Elem()
}

func (o PackageVersionStatusPtrOutput) ToPackageVersionStatusPtrOutput() PackageVersionStatusPtrOutput {
	return o
}

func (o PackageVersionStatusPtrOutput) ToPackageVersionStatusPtrOutputWithContext(ctx context.Context) PackageVersionStatusPtrOutput {
	return o
}

func (o PackageVersionStatusPtrOutput) Elem() PackageVersionStatusOutput {
	return o.ApplyT(func(v *PackageVersionStatus) PackageVersionStatus {
		if v != nil {
			return *v
		}
		var ret PackageVersionStatus
		return ret
	}).(PackageVersionStatusOutput)
}

func (o PackageVersionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PackageVersionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PackageVersionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationInstanceStatusFilterInput)(nil)).Elem(), ApplicationInstanceStatusFilter("DEPLOYMENT_SUCCEEDED"))
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationInstanceStatusFilterPtrInput)(nil)).Elem(), ApplicationInstanceStatusFilter("DEPLOYMENT_SUCCEEDED"))
	pulumi.RegisterOutputType(ApplicationInstanceHealthStatusOutput{})
	pulumi.RegisterOutputType(ApplicationInstanceHealthStatusPtrOutput{})
	pulumi.RegisterOutputType(ApplicationInstanceStatusOutput{})
	pulumi.RegisterOutputType(ApplicationInstanceStatusPtrOutput{})
	pulumi.RegisterOutputType(ApplicationInstanceStatusFilterOutput{})
	pulumi.RegisterOutputType(ApplicationInstanceStatusFilterPtrOutput{})
	pulumi.RegisterOutputType(PackageVersionStatusOutput{})
	pulumi.RegisterOutputType(PackageVersionStatusPtrOutput{})
}
