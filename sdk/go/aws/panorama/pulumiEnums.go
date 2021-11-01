// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

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

func (ApplicationInstanceHealthStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInstanceHealthStatus)(nil)).Elem()
}

func (e ApplicationInstanceHealthStatus) ToApplicationInstanceHealthStatusOutput() ApplicationInstanceHealthStatusOutput {
	return pulumi.ToOutput(e).(ApplicationInstanceHealthStatusOutput)
}

func (e ApplicationInstanceHealthStatus) ToApplicationInstanceHealthStatusOutputWithContext(ctx context.Context) ApplicationInstanceHealthStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApplicationInstanceHealthStatusOutput)
}

func (e ApplicationInstanceHealthStatus) ToApplicationInstanceHealthStatusPtrOutput() ApplicationInstanceHealthStatusPtrOutput {
	return e.ToApplicationInstanceHealthStatusPtrOutputWithContext(context.Background())
}

func (e ApplicationInstanceHealthStatus) ToApplicationInstanceHealthStatusPtrOutputWithContext(ctx context.Context) ApplicationInstanceHealthStatusPtrOutput {
	return ApplicationInstanceHealthStatus(e).ToApplicationInstanceHealthStatusOutputWithContext(ctx).ToApplicationInstanceHealthStatusPtrOutputWithContext(ctx)
}

func (e ApplicationInstanceHealthStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationInstanceHealthStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationInstanceHealthStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApplicationInstanceHealthStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

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

// ApplicationInstanceHealthStatusInput is an input type that accepts ApplicationInstanceHealthStatusArgs and ApplicationInstanceHealthStatusOutput values.
// You can construct a concrete instance of `ApplicationInstanceHealthStatusInput` via:
//
//          ApplicationInstanceHealthStatusArgs{...}
type ApplicationInstanceHealthStatusInput interface {
	pulumi.Input

	ToApplicationInstanceHealthStatusOutput() ApplicationInstanceHealthStatusOutput
	ToApplicationInstanceHealthStatusOutputWithContext(context.Context) ApplicationInstanceHealthStatusOutput
}

var applicationInstanceHealthStatusPtrType = reflect.TypeOf((**ApplicationInstanceHealthStatus)(nil)).Elem()

type ApplicationInstanceHealthStatusPtrInput interface {
	pulumi.Input

	ToApplicationInstanceHealthStatusPtrOutput() ApplicationInstanceHealthStatusPtrOutput
	ToApplicationInstanceHealthStatusPtrOutputWithContext(context.Context) ApplicationInstanceHealthStatusPtrOutput
}

type applicationInstanceHealthStatusPtr string

func ApplicationInstanceHealthStatusPtr(v string) ApplicationInstanceHealthStatusPtrInput {
	return (*applicationInstanceHealthStatusPtr)(&v)
}

func (*applicationInstanceHealthStatusPtr) ElementType() reflect.Type {
	return applicationInstanceHealthStatusPtrType
}

func (in *applicationInstanceHealthStatusPtr) ToApplicationInstanceHealthStatusPtrOutput() ApplicationInstanceHealthStatusPtrOutput {
	return pulumi.ToOutput(in).(ApplicationInstanceHealthStatusPtrOutput)
}

func (in *applicationInstanceHealthStatusPtr) ToApplicationInstanceHealthStatusPtrOutputWithContext(ctx context.Context) ApplicationInstanceHealthStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApplicationInstanceHealthStatusPtrOutput)
}

type ApplicationInstanceStatus string

const (
	ApplicationInstanceStatusDeploymentPending    = ApplicationInstanceStatus("DEPLOYMENT_PENDING")
	ApplicationInstanceStatusDeploymentRequested  = ApplicationInstanceStatus("DEPLOYMENT_REQUESTED")
	ApplicationInstanceStatusDeploymentInProgress = ApplicationInstanceStatus("DEPLOYMENT_IN_PROGRESS")
	ApplicationInstanceStatusDeploymentFailed     = ApplicationInstanceStatus("DEPLOYMENT_FAILED")
	ApplicationInstanceStatusDeploymentSucceeded  = ApplicationInstanceStatus("DEPLOYMENT_SUCCEEDED")
	ApplicationInstanceStatusRemovalPending       = ApplicationInstanceStatus("REMOVAL_PENDING")
	ApplicationInstanceStatusRemovalRequested     = ApplicationInstanceStatus("REMOVAL_REQUESTED")
	ApplicationInstanceStatusRemovalInProgress    = ApplicationInstanceStatus("REMOVAL_IN_PROGRESS")
	ApplicationInstanceStatusRemovalFailed        = ApplicationInstanceStatus("REMOVAL_FAILED")
	ApplicationInstanceStatusRemovalSucceeded     = ApplicationInstanceStatus("REMOVAL_SUCCEEDED")
)

func (ApplicationInstanceStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInstanceStatus)(nil)).Elem()
}

func (e ApplicationInstanceStatus) ToApplicationInstanceStatusOutput() ApplicationInstanceStatusOutput {
	return pulumi.ToOutput(e).(ApplicationInstanceStatusOutput)
}

func (e ApplicationInstanceStatus) ToApplicationInstanceStatusOutputWithContext(ctx context.Context) ApplicationInstanceStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApplicationInstanceStatusOutput)
}

func (e ApplicationInstanceStatus) ToApplicationInstanceStatusPtrOutput() ApplicationInstanceStatusPtrOutput {
	return e.ToApplicationInstanceStatusPtrOutputWithContext(context.Background())
}

func (e ApplicationInstanceStatus) ToApplicationInstanceStatusPtrOutputWithContext(ctx context.Context) ApplicationInstanceStatusPtrOutput {
	return ApplicationInstanceStatus(e).ToApplicationInstanceStatusOutputWithContext(ctx).ToApplicationInstanceStatusPtrOutputWithContext(ctx)
}

func (e ApplicationInstanceStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationInstanceStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationInstanceStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApplicationInstanceStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

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

// ApplicationInstanceStatusInput is an input type that accepts ApplicationInstanceStatusArgs and ApplicationInstanceStatusOutput values.
// You can construct a concrete instance of `ApplicationInstanceStatusInput` via:
//
//          ApplicationInstanceStatusArgs{...}
type ApplicationInstanceStatusInput interface {
	pulumi.Input

	ToApplicationInstanceStatusOutput() ApplicationInstanceStatusOutput
	ToApplicationInstanceStatusOutputWithContext(context.Context) ApplicationInstanceStatusOutput
}

var applicationInstanceStatusPtrType = reflect.TypeOf((**ApplicationInstanceStatus)(nil)).Elem()

type ApplicationInstanceStatusPtrInput interface {
	pulumi.Input

	ToApplicationInstanceStatusPtrOutput() ApplicationInstanceStatusPtrOutput
	ToApplicationInstanceStatusPtrOutputWithContext(context.Context) ApplicationInstanceStatusPtrOutput
}

type applicationInstanceStatusPtr string

func ApplicationInstanceStatusPtr(v string) ApplicationInstanceStatusPtrInput {
	return (*applicationInstanceStatusPtr)(&v)
}

func (*applicationInstanceStatusPtr) ElementType() reflect.Type {
	return applicationInstanceStatusPtrType
}

func (in *applicationInstanceStatusPtr) ToApplicationInstanceStatusPtrOutput() ApplicationInstanceStatusPtrOutput {
	return pulumi.ToOutput(in).(ApplicationInstanceStatusPtrOutput)
}

func (in *applicationInstanceStatusPtr) ToApplicationInstanceStatusPtrOutputWithContext(ctx context.Context) ApplicationInstanceStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApplicationInstanceStatusPtrOutput)
}

type ApplicationInstanceStatusFilter string

const (
	ApplicationInstanceStatusFilterDeploymentSucceeded  = ApplicationInstanceStatusFilter("DEPLOYMENT_SUCCEEDED")
	ApplicationInstanceStatusFilterDeploymentFailed     = ApplicationInstanceStatusFilter("DEPLOYMENT_FAILED")
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

func (PackageVersionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*PackageVersionStatus)(nil)).Elem()
}

func (e PackageVersionStatus) ToPackageVersionStatusOutput() PackageVersionStatusOutput {
	return pulumi.ToOutput(e).(PackageVersionStatusOutput)
}

func (e PackageVersionStatus) ToPackageVersionStatusOutputWithContext(ctx context.Context) PackageVersionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PackageVersionStatusOutput)
}

func (e PackageVersionStatus) ToPackageVersionStatusPtrOutput() PackageVersionStatusPtrOutput {
	return e.ToPackageVersionStatusPtrOutputWithContext(context.Background())
}

func (e PackageVersionStatus) ToPackageVersionStatusPtrOutputWithContext(ctx context.Context) PackageVersionStatusPtrOutput {
	return PackageVersionStatus(e).ToPackageVersionStatusOutputWithContext(ctx).ToPackageVersionStatusPtrOutputWithContext(ctx)
}

func (e PackageVersionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PackageVersionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PackageVersionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PackageVersionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

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

// PackageVersionStatusInput is an input type that accepts PackageVersionStatusArgs and PackageVersionStatusOutput values.
// You can construct a concrete instance of `PackageVersionStatusInput` via:
//
//          PackageVersionStatusArgs{...}
type PackageVersionStatusInput interface {
	pulumi.Input

	ToPackageVersionStatusOutput() PackageVersionStatusOutput
	ToPackageVersionStatusOutputWithContext(context.Context) PackageVersionStatusOutput
}

var packageVersionStatusPtrType = reflect.TypeOf((**PackageVersionStatus)(nil)).Elem()

type PackageVersionStatusPtrInput interface {
	pulumi.Input

	ToPackageVersionStatusPtrOutput() PackageVersionStatusPtrOutput
	ToPackageVersionStatusPtrOutputWithContext(context.Context) PackageVersionStatusPtrOutput
}

type packageVersionStatusPtr string

func PackageVersionStatusPtr(v string) PackageVersionStatusPtrInput {
	return (*packageVersionStatusPtr)(&v)
}

func (*packageVersionStatusPtr) ElementType() reflect.Type {
	return packageVersionStatusPtrType
}

func (in *packageVersionStatusPtr) ToPackageVersionStatusPtrOutput() PackageVersionStatusPtrOutput {
	return pulumi.ToOutput(in).(PackageVersionStatusPtrOutput)
}

func (in *packageVersionStatusPtr) ToPackageVersionStatusPtrOutputWithContext(ctx context.Context) PackageVersionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PackageVersionStatusPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationInstanceHealthStatusOutput{})
	pulumi.RegisterOutputType(ApplicationInstanceHealthStatusPtrOutput{})
	pulumi.RegisterOutputType(ApplicationInstanceStatusOutput{})
	pulumi.RegisterOutputType(ApplicationInstanceStatusPtrOutput{})
	pulumi.RegisterOutputType(ApplicationInstanceStatusFilterOutput{})
	pulumi.RegisterOutputType(ApplicationInstanceStatusFilterPtrOutput{})
	pulumi.RegisterOutputType(PackageVersionStatusOutput{})
	pulumi.RegisterOutputType(PackageVersionStatusPtrOutput{})
}