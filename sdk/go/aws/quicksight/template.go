// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the AWS::QuickSight::Template Resource Type.
type Template struct {
	pulumi.CustomResourceState

	// <p>The Amazon Resource Name (ARN) of the template.</p>
	Arn          pulumi.StringOutput `pulumi:"arn"`
	AwsAccountId pulumi.StringOutput `pulumi:"awsAccountId"`
	// <p>Time when this was created.</p>
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// <p>Time when this was last updated.</p>
	LastUpdatedTime pulumi.StringOutput `pulumi:"lastUpdatedTime"`
	// <p>A display name for the template.</p>
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// <p>A list of resource permissions to be set on the template. </p>
	Permissions  TemplateResourcePermissionArrayOutput `pulumi:"permissions"`
	SourceEntity TemplateSourceEntityOutput            `pulumi:"sourceEntity"`
	// <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the resource.</p>
	Tags       TemplateTagArrayOutput `pulumi:"tags"`
	TemplateId pulumi.StringOutput    `pulumi:"templateId"`
	Version    TemplateVersionOutput  `pulumi:"version"`
	// <p>A description of the current template version being created. This API operation creates the
	// 			first version of the template. Every time <code>UpdateTemplate</code> is called, a new
	// 			version is created. Each version of the template maintains a description of the version
	// 			in the <code>VersionDescription</code> field.</p>
	VersionDescription pulumi.StringPtrOutput `pulumi:"versionDescription"`
}

// NewTemplate registers a new resource with the given unique name, arguments, and options.
func NewTemplate(ctx *pulumi.Context,
	name string, args *TemplateArgs, opts ...pulumi.ResourceOption) (*Template, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AwsAccountId == nil {
		return nil, errors.New("invalid value for required argument 'AwsAccountId'")
	}
	if args.SourceEntity == nil {
		return nil, errors.New("invalid value for required argument 'SourceEntity'")
	}
	if args.TemplateId == nil {
		return nil, errors.New("invalid value for required argument 'TemplateId'")
	}
	var resource Template
	err := ctx.RegisterResource("aws-native:quicksight:Template", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTemplate gets an existing Template resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateState, opts ...pulumi.ResourceOption) (*Template, error) {
	var resource Template
	err := ctx.ReadResource("aws-native:quicksight:Template", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Template resources.
type templateState struct {
}

type TemplateState struct {
}

func (TemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateState)(nil)).Elem()
}

type templateArgs struct {
	AwsAccountId string `pulumi:"awsAccountId"`
	// <p>A display name for the template.</p>
	Name *string `pulumi:"name"`
	// <p>A list of resource permissions to be set on the template. </p>
	Permissions  []TemplateResourcePermission `pulumi:"permissions"`
	SourceEntity TemplateSourceEntity         `pulumi:"sourceEntity"`
	// <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the resource.</p>
	Tags       []TemplateTag `pulumi:"tags"`
	TemplateId string        `pulumi:"templateId"`
	// <p>A description of the current template version being created. This API operation creates the
	// 			first version of the template. Every time <code>UpdateTemplate</code> is called, a new
	// 			version is created. Each version of the template maintains a description of the version
	// 			in the <code>VersionDescription</code> field.</p>
	VersionDescription *string `pulumi:"versionDescription"`
}

// The set of arguments for constructing a Template resource.
type TemplateArgs struct {
	AwsAccountId pulumi.StringInput
	// <p>A display name for the template.</p>
	Name pulumi.StringPtrInput
	// <p>A list of resource permissions to be set on the template. </p>
	Permissions  TemplateResourcePermissionArrayInput
	SourceEntity TemplateSourceEntityInput
	// <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the resource.</p>
	Tags       TemplateTagArrayInput
	TemplateId pulumi.StringInput
	// <p>A description of the current template version being created. This API operation creates the
	// 			first version of the template. Every time <code>UpdateTemplate</code> is called, a new
	// 			version is created. Each version of the template maintains a description of the version
	// 			in the <code>VersionDescription</code> field.</p>
	VersionDescription pulumi.StringPtrInput
}

func (TemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateArgs)(nil)).Elem()
}

type TemplateInput interface {
	pulumi.Input

	ToTemplateOutput() TemplateOutput
	ToTemplateOutputWithContext(ctx context.Context) TemplateOutput
}

func (*Template) ElementType() reflect.Type {
	return reflect.TypeOf((**Template)(nil)).Elem()
}

func (i *Template) ToTemplateOutput() TemplateOutput {
	return i.ToTemplateOutputWithContext(context.Background())
}

func (i *Template) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateOutput)
}

type TemplateOutput struct{ *pulumi.OutputState }

func (TemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Template)(nil)).Elem()
}

func (o TemplateOutput) ToTemplateOutput() TemplateOutput {
	return o
}

func (o TemplateOutput) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return o
}

// <p>The Amazon Resource Name (ARN) of the template.</p>
func (o TemplateOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o TemplateOutput) AwsAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.AwsAccountId }).(pulumi.StringOutput)
}

// <p>Time when this was created.</p>
func (o TemplateOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// <p>Time when this was last updated.</p>
func (o TemplateOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

// <p>A display name for the template.</p>
func (o TemplateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Template) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// <p>A list of resource permissions to be set on the template. </p>
func (o TemplateOutput) Permissions() TemplateResourcePermissionArrayOutput {
	return o.ApplyT(func(v *Template) TemplateResourcePermissionArrayOutput { return v.Permissions }).(TemplateResourcePermissionArrayOutput)
}

func (o TemplateOutput) SourceEntity() TemplateSourceEntityOutput {
	return o.ApplyT(func(v *Template) TemplateSourceEntityOutput { return v.SourceEntity }).(TemplateSourceEntityOutput)
}

// <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the resource.</p>
func (o TemplateOutput) Tags() TemplateTagArrayOutput {
	return o.ApplyT(func(v *Template) TemplateTagArrayOutput { return v.Tags }).(TemplateTagArrayOutput)
}

func (o TemplateOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.TemplateId }).(pulumi.StringOutput)
}

func (o TemplateOutput) Version() TemplateVersionOutput {
	return o.ApplyT(func(v *Template) TemplateVersionOutput { return v.Version }).(TemplateVersionOutput)
}

// <p>A description of the current template version being created. This API operation creates the
// 			first version of the template. Every time <code>UpdateTemplate</code> is called, a new
// 			version is created. Each version of the template maintains a description of the version
// 			in the <code>VersionDescription</code> field.</p>
func (o TemplateOutput) VersionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Template) pulumi.StringPtrOutput { return v.VersionDescription }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateInput)(nil)).Elem(), &Template{})
	pulumi.RegisterOutputType(TemplateOutput{})
}
