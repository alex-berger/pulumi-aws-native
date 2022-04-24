// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kendra

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Kendra index
type Index struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// Capacity units
	CapacityUnits IndexCapacityUnitsConfigurationPtrOutput `pulumi:"capacityUnits"`
	// A description for the index
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Document metadata configurations
	DocumentMetadataConfigurations IndexDocumentMetadataConfigurationArrayOutput `pulumi:"documentMetadataConfigurations"`
	Edition                        IndexEditionOutput                            `pulumi:"edition"`
	Name                           pulumi.StringOutput                           `pulumi:"name"`
	RoleArn                        pulumi.StringOutput                           `pulumi:"roleArn"`
	// Server side encryption configuration
	ServerSideEncryptionConfiguration IndexServerSideEncryptionConfigurationPtrOutput `pulumi:"serverSideEncryptionConfiguration"`
	// Tags for labeling the index
	Tags                    IndexTagArrayOutput                    `pulumi:"tags"`
	UserContextPolicy       IndexUserContextPolicyPtrOutput        `pulumi:"userContextPolicy"`
	UserTokenConfigurations IndexUserTokenConfigurationArrayOutput `pulumi:"userTokenConfigurations"`
}

// NewIndex registers a new resource with the given unique name, arguments, and options.
func NewIndex(ctx *pulumi.Context,
	name string, args *IndexArgs, opts ...pulumi.ResourceOption) (*Index, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Edition == nil {
		return nil, errors.New("invalid value for required argument 'Edition'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource Index
	err := ctx.RegisterResource("aws-native:kendra:Index", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIndex gets an existing Index resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIndex(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IndexState, opts ...pulumi.ResourceOption) (*Index, error) {
	var resource Index
	err := ctx.ReadResource("aws-native:kendra:Index", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Index resources.
type indexState struct {
}

type IndexState struct {
}

func (IndexState) ElementType() reflect.Type {
	return reflect.TypeOf((*indexState)(nil)).Elem()
}

type indexArgs struct {
	// Capacity units
	CapacityUnits *IndexCapacityUnitsConfiguration `pulumi:"capacityUnits"`
	// A description for the index
	Description *string `pulumi:"description"`
	// Document metadata configurations
	DocumentMetadataConfigurations []IndexDocumentMetadataConfiguration `pulumi:"documentMetadataConfigurations"`
	Edition                        IndexEdition                         `pulumi:"edition"`
	Name                           *string                              `pulumi:"name"`
	RoleArn                        string                               `pulumi:"roleArn"`
	// Server side encryption configuration
	ServerSideEncryptionConfiguration *IndexServerSideEncryptionConfiguration `pulumi:"serverSideEncryptionConfiguration"`
	// Tags for labeling the index
	Tags                    []IndexTag                    `pulumi:"tags"`
	UserContextPolicy       *IndexUserContextPolicy       `pulumi:"userContextPolicy"`
	UserTokenConfigurations []IndexUserTokenConfiguration `pulumi:"userTokenConfigurations"`
}

// The set of arguments for constructing a Index resource.
type IndexArgs struct {
	// Capacity units
	CapacityUnits IndexCapacityUnitsConfigurationPtrInput
	// A description for the index
	Description pulumi.StringPtrInput
	// Document metadata configurations
	DocumentMetadataConfigurations IndexDocumentMetadataConfigurationArrayInput
	Edition                        IndexEditionInput
	Name                           pulumi.StringPtrInput
	RoleArn                        pulumi.StringInput
	// Server side encryption configuration
	ServerSideEncryptionConfiguration IndexServerSideEncryptionConfigurationPtrInput
	// Tags for labeling the index
	Tags                    IndexTagArrayInput
	UserContextPolicy       IndexUserContextPolicyPtrInput
	UserTokenConfigurations IndexUserTokenConfigurationArrayInput
}

func (IndexArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*indexArgs)(nil)).Elem()
}

type IndexInput interface {
	pulumi.Input

	ToIndexOutput() IndexOutput
	ToIndexOutputWithContext(ctx context.Context) IndexOutput
}

func (*Index) ElementType() reflect.Type {
	return reflect.TypeOf((**Index)(nil)).Elem()
}

func (i *Index) ToIndexOutput() IndexOutput {
	return i.ToIndexOutputWithContext(context.Background())
}

func (i *Index) ToIndexOutputWithContext(ctx context.Context) IndexOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexOutput)
}

type IndexOutput struct{ *pulumi.OutputState }

func (IndexOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Index)(nil)).Elem()
}

func (o IndexOutput) ToIndexOutput() IndexOutput {
	return o
}

func (o IndexOutput) ToIndexOutputWithContext(ctx context.Context) IndexOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IndexInput)(nil)).Elem(), &Index{})
	pulumi.RegisterOutputType(IndexOutput{})
}
