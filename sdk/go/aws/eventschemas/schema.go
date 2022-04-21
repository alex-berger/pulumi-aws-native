// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventschemas

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EventSchemas::Schema
//
// Deprecated: Schema is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Schema struct {
	pulumi.CustomResourceState

	Content       pulumi.StringOutput        `pulumi:"content"`
	Description   pulumi.StringPtrOutput     `pulumi:"description"`
	RegistryName  pulumi.StringOutput        `pulumi:"registryName"`
	SchemaArn     pulumi.StringOutput        `pulumi:"schemaArn"`
	SchemaName    pulumi.StringPtrOutput     `pulumi:"schemaName"`
	SchemaVersion pulumi.StringOutput        `pulumi:"schemaVersion"`
	Tags          SchemaTagsEntryArrayOutput `pulumi:"tags"`
	Type          pulumi.StringOutput        `pulumi:"type"`
}

// NewSchema registers a new resource with the given unique name, arguments, and options.
func NewSchema(ctx *pulumi.Context,
	name string, args *SchemaArgs, opts ...pulumi.ResourceOption) (*Schema, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource Schema
	err := ctx.RegisterResource("aws-native:eventschemas:Schema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSchema gets an existing Schema resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SchemaState, opts ...pulumi.ResourceOption) (*Schema, error) {
	var resource Schema
	err := ctx.ReadResource("aws-native:eventschemas:Schema", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Schema resources.
type schemaState struct {
}

type SchemaState struct {
}

func (SchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaState)(nil)).Elem()
}

type schemaArgs struct {
	Content      string            `pulumi:"content"`
	Description  *string           `pulumi:"description"`
	RegistryName string            `pulumi:"registryName"`
	SchemaName   *string           `pulumi:"schemaName"`
	Tags         []SchemaTagsEntry `pulumi:"tags"`
	Type         string            `pulumi:"type"`
}

// The set of arguments for constructing a Schema resource.
type SchemaArgs struct {
	Content      pulumi.StringInput
	Description  pulumi.StringPtrInput
	RegistryName pulumi.StringInput
	SchemaName   pulumi.StringPtrInput
	Tags         SchemaTagsEntryArrayInput
	Type         pulumi.StringInput
}

func (SchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaArgs)(nil)).Elem()
}

type SchemaInput interface {
	pulumi.Input

	ToSchemaOutput() SchemaOutput
	ToSchemaOutputWithContext(ctx context.Context) SchemaOutput
}

func (*Schema) ElementType() reflect.Type {
	return reflect.TypeOf((**Schema)(nil)).Elem()
}

func (i *Schema) ToSchemaOutput() SchemaOutput {
	return i.ToSchemaOutputWithContext(context.Background())
}

func (i *Schema) ToSchemaOutputWithContext(ctx context.Context) SchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaOutput)
}

type SchemaOutput struct{ *pulumi.OutputState }

func (SchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Schema)(nil)).Elem()
}

func (o SchemaOutput) ToSchemaOutput() SchemaOutput {
	return o
}

func (o SchemaOutput) ToSchemaOutputWithContext(ctx context.Context) SchemaOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaInput)(nil)).Elem(), &Schema{})
	pulumi.RegisterOutputType(SchemaOutput{})
}
