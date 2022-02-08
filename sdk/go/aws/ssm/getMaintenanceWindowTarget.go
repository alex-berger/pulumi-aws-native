// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SSM::MaintenanceWindowTarget
func LookupMaintenanceWindowTarget(ctx *pulumi.Context, args *LookupMaintenanceWindowTargetArgs, opts ...pulumi.InvokeOption) (*LookupMaintenanceWindowTargetResult, error) {
	var rv LookupMaintenanceWindowTargetResult
	err := ctx.Invoke("aws-native:ssm:getMaintenanceWindowTarget", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMaintenanceWindowTargetArgs struct {
	Id string `pulumi:"id"`
}

type LookupMaintenanceWindowTargetResult struct {
	Description      *string                          `pulumi:"description"`
	Id               *string                          `pulumi:"id"`
	Name             *string                          `pulumi:"name"`
	OwnerInformation *string                          `pulumi:"ownerInformation"`
	ResourceType     *string                          `pulumi:"resourceType"`
	Targets          []MaintenanceWindowTargetTargets `pulumi:"targets"`
}

func LookupMaintenanceWindowTargetOutput(ctx *pulumi.Context, args LookupMaintenanceWindowTargetOutputArgs, opts ...pulumi.InvokeOption) LookupMaintenanceWindowTargetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMaintenanceWindowTargetResult, error) {
			args := v.(LookupMaintenanceWindowTargetArgs)
			r, err := LookupMaintenanceWindowTarget(ctx, &args, opts...)
			return *r, err
		}).(LookupMaintenanceWindowTargetResultOutput)
}

type LookupMaintenanceWindowTargetOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupMaintenanceWindowTargetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMaintenanceWindowTargetArgs)(nil)).Elem()
}

type LookupMaintenanceWindowTargetResultOutput struct{ *pulumi.OutputState }

func (LookupMaintenanceWindowTargetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMaintenanceWindowTargetResult)(nil)).Elem()
}

func (o LookupMaintenanceWindowTargetResultOutput) ToLookupMaintenanceWindowTargetResultOutput() LookupMaintenanceWindowTargetResultOutput {
	return o
}

func (o LookupMaintenanceWindowTargetResultOutput) ToLookupMaintenanceWindowTargetResultOutputWithContext(ctx context.Context) LookupMaintenanceWindowTargetResultOutput {
	return o
}

func (o LookupMaintenanceWindowTargetResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowTargetResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceWindowTargetResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowTargetResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceWindowTargetResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowTargetResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceWindowTargetResultOutput) OwnerInformation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowTargetResult) *string { return v.OwnerInformation }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceWindowTargetResultOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowTargetResult) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceWindowTargetResultOutput) Targets() MaintenanceWindowTargetTargetsArrayOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowTargetResult) []MaintenanceWindowTargetTargets { return v.Targets }).(MaintenanceWindowTargetTargetsArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMaintenanceWindowTargetResultOutput{})
}