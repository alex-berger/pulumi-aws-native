// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package imagebuilder

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::ImageBuilder::Component
type Component struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the component.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The change description of the component.
	ChangeDescription pulumi.StringPtrOutput `pulumi:"changeDescription"`
	// The data of the component.
	Data pulumi.StringPtrOutput `pulumi:"data"`
	// The description of the component.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The encryption status of the component.
	Encrypted pulumi.BoolOutput `pulumi:"encrypted"`
	// The KMS key identifier used to encrypt the component.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// The name of the component.
	Name pulumi.StringOutput `pulumi:"name"`
	// The platform of the component.
	Platform ComponentPlatformOutput `pulumi:"platform"`
	// The operating system (OS) version supported by the component.
	SupportedOsVersions pulumi.StringArrayOutput `pulumi:"supportedOsVersions"`
	// The tags associated with the component.
	Tags pulumi.AnyOutput `pulumi:"tags"`
	// The type of the component denotes whether the component is used to build the image or only to test it.
	Type ComponentTypeOutput `pulumi:"type"`
	// The uri of the component.
	Uri pulumi.StringPtrOutput `pulumi:"uri"`
	// The version of the component.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewComponent registers a new resource with the given unique name, arguments, and options.
func NewComponent(ctx *pulumi.Context,
	name string, args *ComponentArgs, opts ...pulumi.ResourceOption) (*Component, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Platform == nil {
		return nil, errors.New("invalid value for required argument 'Platform'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	var resource Component
	err := ctx.RegisterResource("aws-native:imagebuilder:Component", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComponent gets an existing Component resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComponent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComponentState, opts ...pulumi.ResourceOption) (*Component, error) {
	var resource Component
	err := ctx.ReadResource("aws-native:imagebuilder:Component", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Component resources.
type componentState struct {
}

type ComponentState struct {
}

func (ComponentState) ElementType() reflect.Type {
	return reflect.TypeOf((*componentState)(nil)).Elem()
}

type componentArgs struct {
	// The change description of the component.
	ChangeDescription *string `pulumi:"changeDescription"`
	// The data of the component.
	Data *string `pulumi:"data"`
	// The description of the component.
	Description *string `pulumi:"description"`
	// The KMS key identifier used to encrypt the component.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The name of the component.
	Name *string `pulumi:"name"`
	// The platform of the component.
	Platform ComponentPlatform `pulumi:"platform"`
	// The operating system (OS) version supported by the component.
	SupportedOsVersions []string `pulumi:"supportedOsVersions"`
	// The tags associated with the component.
	Tags interface{} `pulumi:"tags"`
	// The uri of the component.
	Uri *string `pulumi:"uri"`
	// The version of the component.
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a Component resource.
type ComponentArgs struct {
	// The change description of the component.
	ChangeDescription pulumi.StringPtrInput
	// The data of the component.
	Data pulumi.StringPtrInput
	// The description of the component.
	Description pulumi.StringPtrInput
	// The KMS key identifier used to encrypt the component.
	KmsKeyId pulumi.StringPtrInput
	// The name of the component.
	Name pulumi.StringPtrInput
	// The platform of the component.
	Platform ComponentPlatformInput
	// The operating system (OS) version supported by the component.
	SupportedOsVersions pulumi.StringArrayInput
	// The tags associated with the component.
	Tags pulumi.Input
	// The uri of the component.
	Uri pulumi.StringPtrInput
	// The version of the component.
	Version pulumi.StringInput
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type ComponentInput interface {
	pulumi.Input

	ToComponentOutput() ComponentOutput
	ToComponentOutputWithContext(ctx context.Context) ComponentOutput
}

func (*Component) ElementType() reflect.Type {
	return reflect.TypeOf((**Component)(nil)).Elem()
}

func (i *Component) ToComponentOutput() ComponentOutput {
	return i.ToComponentOutputWithContext(context.Background())
}

func (i *Component) ToComponentOutputWithContext(ctx context.Context) ComponentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentOutput)
}

type ComponentOutput struct{ *pulumi.OutputState }

func (ComponentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Component)(nil)).Elem()
}

func (o ComponentOutput) ToComponentOutput() ComponentOutput {
	return o
}

func (o ComponentOutput) ToComponentOutputWithContext(ctx context.Context) ComponentOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComponentInput)(nil)).Elem(), &Component{})
	pulumi.RegisterOutputType(ComponentOutput{})
}
