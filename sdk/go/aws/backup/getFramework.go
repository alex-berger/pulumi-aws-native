// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package backup

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Contains detailed information about a framework. Frameworks contain controls, which evaluate and report on your backup events and resources. Frameworks generate daily compliance results.
func LookupFramework(ctx *pulumi.Context, args *LookupFrameworkArgs, opts ...pulumi.InvokeOption) (*LookupFrameworkResult, error) {
	var rv LookupFrameworkResult
	err := ctx.Invoke("aws-native:backup:getFramework", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFrameworkArgs struct {
	// An Amazon Resource Name (ARN) that uniquely identifies Framework as a resource
	FrameworkArn string `pulumi:"frameworkArn"`
}

type LookupFrameworkResult struct {
	// The date and time that a framework is created, in Unix format and Coordinated Universal Time (UTC). The value of `CreationTime` is accurate to milliseconds. For example, the value 1516925490.087 represents Friday, January 26, 2018 12:11:30.087 AM.
	CreationTime *float64 `pulumi:"creationTime"`
	// The deployment status of a framework. The statuses are: `CREATE_IN_PROGRESS | UPDATE_IN_PROGRESS | DELETE_IN_PROGRESS | COMPLETED | FAILED`
	DeploymentStatus *string `pulumi:"deploymentStatus"`
	// An Amazon Resource Name (ARN) that uniquely identifies Framework as a resource
	FrameworkArn *string `pulumi:"frameworkArn"`
	// Contains detailed information about all of the controls of a framework. Each framework must contain at least one control.
	FrameworkControls []FrameworkControl `pulumi:"frameworkControls"`
	// An optional description of the framework with a maximum 1,024 characters.
	FrameworkDescription *string `pulumi:"frameworkDescription"`
	// A framework consists of one or more controls. Each control governs a resource, such as backup plans, backup selections, backup vaults, or recovery points. You can also turn AWS Config recording on or off for each resource. The statuses are:
	//
	// `ACTIVE` when recording is turned on for all resources governed by the framework.
	//
	// `PARTIALLY_ACTIVE` when recording is turned off for at least one resource governed by the framework.
	//
	// `INACTIVE` when recording is turned off for all resources governed by the framework.
	//
	// `UNAVAILABLE` when AWS Backup is unable to validate recording status at this time.
	FrameworkStatus *string `pulumi:"frameworkStatus"`
	// Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.
	FrameworkTags []FrameworkTag `pulumi:"frameworkTags"`
}

func LookupFrameworkOutput(ctx *pulumi.Context, args LookupFrameworkOutputArgs, opts ...pulumi.InvokeOption) LookupFrameworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFrameworkResult, error) {
			args := v.(LookupFrameworkArgs)
			r, err := LookupFramework(ctx, &args, opts...)
			var s LookupFrameworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFrameworkResultOutput)
}

type LookupFrameworkOutputArgs struct {
	// An Amazon Resource Name (ARN) that uniquely identifies Framework as a resource
	FrameworkArn pulumi.StringInput `pulumi:"frameworkArn"`
}

func (LookupFrameworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFrameworkArgs)(nil)).Elem()
}

type LookupFrameworkResultOutput struct{ *pulumi.OutputState }

func (LookupFrameworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFrameworkResult)(nil)).Elem()
}

func (o LookupFrameworkResultOutput) ToLookupFrameworkResultOutput() LookupFrameworkResultOutput {
	return o
}

func (o LookupFrameworkResultOutput) ToLookupFrameworkResultOutputWithContext(ctx context.Context) LookupFrameworkResultOutput {
	return o
}

// The date and time that a framework is created, in Unix format and Coordinated Universal Time (UTC). The value of `CreationTime` is accurate to milliseconds. For example, the value 1516925490.087 represents Friday, January 26, 2018 12:11:30.087 AM.
func (o LookupFrameworkResultOutput) CreationTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupFrameworkResult) *float64 { return v.CreationTime }).(pulumi.Float64PtrOutput)
}

// The deployment status of a framework. The statuses are: `CREATE_IN_PROGRESS | UPDATE_IN_PROGRESS | DELETE_IN_PROGRESS | COMPLETED | FAILED`
func (o LookupFrameworkResultOutput) DeploymentStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFrameworkResult) *string { return v.DeploymentStatus }).(pulumi.StringPtrOutput)
}

// An Amazon Resource Name (ARN) that uniquely identifies Framework as a resource
func (o LookupFrameworkResultOutput) FrameworkArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFrameworkResult) *string { return v.FrameworkArn }).(pulumi.StringPtrOutput)
}

// Contains detailed information about all of the controls of a framework. Each framework must contain at least one control.
func (o LookupFrameworkResultOutput) FrameworkControls() FrameworkControlArrayOutput {
	return o.ApplyT(func(v LookupFrameworkResult) []FrameworkControl { return v.FrameworkControls }).(FrameworkControlArrayOutput)
}

// An optional description of the framework with a maximum 1,024 characters.
func (o LookupFrameworkResultOutput) FrameworkDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFrameworkResult) *string { return v.FrameworkDescription }).(pulumi.StringPtrOutput)
}

// A framework consists of one or more controls. Each control governs a resource, such as backup plans, backup selections, backup vaults, or recovery points. You can also turn AWS Config recording on or off for each resource. The statuses are:
//
// `ACTIVE` when recording is turned on for all resources governed by the framework.
//
// `PARTIALLY_ACTIVE` when recording is turned off for at least one resource governed by the framework.
//
// `INACTIVE` when recording is turned off for all resources governed by the framework.
//
// `UNAVAILABLE` when AWS Backup is unable to validate recording status at this time.
func (o LookupFrameworkResultOutput) FrameworkStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFrameworkResult) *string { return v.FrameworkStatus }).(pulumi.StringPtrOutput)
}

// Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.
func (o LookupFrameworkResultOutput) FrameworkTags() FrameworkTagArrayOutput {
	return o.ApplyT(func(v LookupFrameworkResult) []FrameworkTag { return v.FrameworkTags }).(FrameworkTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFrameworkResultOutput{})
}
