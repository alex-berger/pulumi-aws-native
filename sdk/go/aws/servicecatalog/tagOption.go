// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ServiceCatalog::TagOption
//
// Deprecated: TagOption is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.
type TagOption struct {
	pulumi.CustomResourceState

	Active pulumi.BoolPtrOutput `pulumi:"active"`
	Key    pulumi.StringOutput  `pulumi:"key"`
	Value  pulumi.StringOutput  `pulumi:"value"`
}

// NewTagOption registers a new resource with the given unique name, arguments, and options.
func NewTagOption(ctx *pulumi.Context,
	name string, args *TagOptionArgs, opts ...pulumi.ResourceOption) (*TagOption, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	var resource TagOption
	err := ctx.RegisterResource("aws-native:servicecatalog:TagOption", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagOption gets an existing TagOption resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagOption(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagOptionState, opts ...pulumi.ResourceOption) (*TagOption, error) {
	var resource TagOption
	err := ctx.ReadResource("aws-native:servicecatalog:TagOption", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagOption resources.
type tagOptionState struct {
}

type TagOptionState struct {
}

func (TagOptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagOptionState)(nil)).Elem()
}

type tagOptionArgs struct {
	Active *bool  `pulumi:"active"`
	Key    string `pulumi:"key"`
	Value  string `pulumi:"value"`
}

// The set of arguments for constructing a TagOption resource.
type TagOptionArgs struct {
	Active pulumi.BoolPtrInput
	Key    pulumi.StringInput
	Value  pulumi.StringInput
}

func (TagOptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagOptionArgs)(nil)).Elem()
}

type TagOptionInput interface {
	pulumi.Input

	ToTagOptionOutput() TagOptionOutput
	ToTagOptionOutputWithContext(ctx context.Context) TagOptionOutput
}

func (*TagOption) ElementType() reflect.Type {
	return reflect.TypeOf((*TagOption)(nil))
}

func (i *TagOption) ToTagOptionOutput() TagOptionOutput {
	return i.ToTagOptionOutputWithContext(context.Background())
}

func (i *TagOption) ToTagOptionOutputWithContext(ctx context.Context) TagOptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagOptionOutput)
}

type TagOptionOutput struct{ *pulumi.OutputState }

func (TagOptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagOption)(nil))
}

func (o TagOptionOutput) ToTagOptionOutput() TagOptionOutput {
	return o
}

func (o TagOptionOutput) ToTagOptionOutputWithContext(ctx context.Context) TagOptionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TagOptionOutput{})
}