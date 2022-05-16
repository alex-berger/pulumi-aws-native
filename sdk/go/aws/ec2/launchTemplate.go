// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::LaunchTemplate
//
// Deprecated: LaunchTemplate is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type LaunchTemplate struct {
	pulumi.CustomResourceState

	DefaultVersionNumber pulumi.StringOutput                       `pulumi:"defaultVersionNumber"`
	LatestVersionNumber  pulumi.StringOutput                       `pulumi:"latestVersionNumber"`
	LaunchTemplateData   LaunchTemplateDataPtrOutput               `pulumi:"launchTemplateData"`
	LaunchTemplateName   pulumi.StringPtrOutput                    `pulumi:"launchTemplateName"`
	TagSpecifications    LaunchTemplateTagSpecificationArrayOutput `pulumi:"tagSpecifications"`
}

// NewLaunchTemplate registers a new resource with the given unique name, arguments, and options.
func NewLaunchTemplate(ctx *pulumi.Context,
	name string, args *LaunchTemplateArgs, opts ...pulumi.ResourceOption) (*LaunchTemplate, error) {
	if args == nil {
		args = &LaunchTemplateArgs{}
	}

	var resource LaunchTemplate
	err := ctx.RegisterResource("aws-native:ec2:LaunchTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLaunchTemplate gets an existing LaunchTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLaunchTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LaunchTemplateState, opts ...pulumi.ResourceOption) (*LaunchTemplate, error) {
	var resource LaunchTemplate
	err := ctx.ReadResource("aws-native:ec2:LaunchTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LaunchTemplate resources.
type launchTemplateState struct {
}

type LaunchTemplateState struct {
}

func (LaunchTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*launchTemplateState)(nil)).Elem()
}

type launchTemplateArgs struct {
	LaunchTemplateData *LaunchTemplateData              `pulumi:"launchTemplateData"`
	LaunchTemplateName *string                          `pulumi:"launchTemplateName"`
	TagSpecifications  []LaunchTemplateTagSpecification `pulumi:"tagSpecifications"`
}

// The set of arguments for constructing a LaunchTemplate resource.
type LaunchTemplateArgs struct {
	LaunchTemplateData LaunchTemplateDataPtrInput
	LaunchTemplateName pulumi.StringPtrInput
	TagSpecifications  LaunchTemplateTagSpecificationArrayInput
}

func (LaunchTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*launchTemplateArgs)(nil)).Elem()
}

type LaunchTemplateInput interface {
	pulumi.Input

	ToLaunchTemplateOutput() LaunchTemplateOutput
	ToLaunchTemplateOutputWithContext(ctx context.Context) LaunchTemplateOutput
}

func (*LaunchTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**LaunchTemplate)(nil)).Elem()
}

func (i *LaunchTemplate) ToLaunchTemplateOutput() LaunchTemplateOutput {
	return i.ToLaunchTemplateOutputWithContext(context.Background())
}

func (i *LaunchTemplate) ToLaunchTemplateOutputWithContext(ctx context.Context) LaunchTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchTemplateOutput)
}

type LaunchTemplateOutput struct{ *pulumi.OutputState }

func (LaunchTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LaunchTemplate)(nil)).Elem()
}

func (o LaunchTemplateOutput) ToLaunchTemplateOutput() LaunchTemplateOutput {
	return o
}

func (o LaunchTemplateOutput) ToLaunchTemplateOutputWithContext(ctx context.Context) LaunchTemplateOutput {
	return o
}

func (o LaunchTemplateOutput) DefaultVersionNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *LaunchTemplate) pulumi.StringOutput { return v.DefaultVersionNumber }).(pulumi.StringOutput)
}

func (o LaunchTemplateOutput) LatestVersionNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *LaunchTemplate) pulumi.StringOutput { return v.LatestVersionNumber }).(pulumi.StringOutput)
}

func (o LaunchTemplateOutput) LaunchTemplateData() LaunchTemplateDataPtrOutput {
	return o.ApplyT(func(v *LaunchTemplate) LaunchTemplateDataPtrOutput { return v.LaunchTemplateData }).(LaunchTemplateDataPtrOutput)
}

func (o LaunchTemplateOutput) LaunchTemplateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LaunchTemplate) pulumi.StringPtrOutput { return v.LaunchTemplateName }).(pulumi.StringPtrOutput)
}

func (o LaunchTemplateOutput) TagSpecifications() LaunchTemplateTagSpecificationArrayOutput {
	return o.ApplyT(func(v *LaunchTemplate) LaunchTemplateTagSpecificationArrayOutput { return v.TagSpecifications }).(LaunchTemplateTagSpecificationArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LaunchTemplateInput)(nil)).Elem(), &LaunchTemplate{})
	pulumi.RegisterOutputType(LaunchTemplateOutput{})
}
