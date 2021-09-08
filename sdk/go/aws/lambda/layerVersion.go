// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html
type LayerVersion struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-compatibleruntimes
	CompatibleRuntimes pulumi.StringArrayOutput `pulumi:"compatibleRuntimes"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-content
	Content LayerVersionContentOutput `pulumi:"content"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-layername
	LayerName pulumi.StringPtrOutput `pulumi:"layerName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-licenseinfo
	LicenseInfo pulumi.StringPtrOutput `pulumi:"licenseInfo"`
}

// NewLayerVersion registers a new resource with the given unique name, arguments, and options.
func NewLayerVersion(ctx *pulumi.Context,
	name string, args *LayerVersionArgs, opts ...pulumi.ResourceOption) (*LayerVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	var resource LayerVersion
	err := ctx.RegisterResource("aws-native:lambda:LayerVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLayerVersion gets an existing LayerVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLayerVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LayerVersionState, opts ...pulumi.ResourceOption) (*LayerVersion, error) {
	var resource LayerVersion
	err := ctx.ReadResource("aws-native:lambda:LayerVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LayerVersion resources.
type layerVersionState struct {
}

type LayerVersionState struct {
}

func (LayerVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*layerVersionState)(nil)).Elem()
}

type layerVersionArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-compatibleruntimes
	CompatibleRuntimes []string `pulumi:"compatibleRuntimes"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-content
	Content LayerVersionContent `pulumi:"content"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-description
	Description *string `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-layername
	LayerName *string `pulumi:"layerName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-licenseinfo
	LicenseInfo *string `pulumi:"licenseInfo"`
}

// The set of arguments for constructing a LayerVersion resource.
type LayerVersionArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-compatibleruntimes
	CompatibleRuntimes pulumi.StringArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-content
	Content LayerVersionContentInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-description
	Description pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-layername
	LayerName pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-licenseinfo
	LicenseInfo pulumi.StringPtrInput
}

func (LayerVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*layerVersionArgs)(nil)).Elem()
}

type LayerVersionInput interface {
	pulumi.Input

	ToLayerVersionOutput() LayerVersionOutput
	ToLayerVersionOutputWithContext(ctx context.Context) LayerVersionOutput
}

func (*LayerVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*LayerVersion)(nil))
}

func (i *LayerVersion) ToLayerVersionOutput() LayerVersionOutput {
	return i.ToLayerVersionOutputWithContext(context.Background())
}

func (i *LayerVersion) ToLayerVersionOutputWithContext(ctx context.Context) LayerVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LayerVersionOutput)
}

type LayerVersionOutput struct{ *pulumi.OutputState }

func (LayerVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LayerVersion)(nil))
}

func (o LayerVersionOutput) ToLayerVersionOutput() LayerVersionOutput {
	return o
}

func (o LayerVersionOutput) ToLayerVersionOutputWithContext(ctx context.Context) LayerVersionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LayerVersionOutput{})
}