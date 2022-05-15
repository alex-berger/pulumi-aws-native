// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Glue::Partition
//
// Deprecated: Partition is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Partition struct {
	pulumi.CustomResourceState

	CatalogId      pulumi.StringOutput      `pulumi:"catalogId"`
	DatabaseName   pulumi.StringOutput      `pulumi:"databaseName"`
	PartitionInput PartitionInputTypeOutput `pulumi:"partitionInput"`
	TableName      pulumi.StringOutput      `pulumi:"tableName"`
}

// NewPartition registers a new resource with the given unique name, arguments, and options.
func NewPartition(ctx *pulumi.Context,
	name string, args *PartitionArgs, opts ...pulumi.ResourceOption) (*Partition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CatalogId == nil {
		return nil, errors.New("invalid value for required argument 'CatalogId'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.PartitionInput == nil {
		return nil, errors.New("invalid value for required argument 'PartitionInput'")
	}
	if args.TableName == nil {
		return nil, errors.New("invalid value for required argument 'TableName'")
	}
	var resource Partition
	err := ctx.RegisterResource("aws-native:glue:Partition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPartition gets an existing Partition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPartition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PartitionState, opts ...pulumi.ResourceOption) (*Partition, error) {
	var resource Partition
	err := ctx.ReadResource("aws-native:glue:Partition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Partition resources.
type partitionState struct {
}

type PartitionState struct {
}

func (PartitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*partitionState)(nil)).Elem()
}

type partitionArgs struct {
	CatalogId      string             `pulumi:"catalogId"`
	DatabaseName   string             `pulumi:"databaseName"`
	PartitionInput PartitionInputType `pulumi:"partitionInput"`
	TableName      string             `pulumi:"tableName"`
}

// The set of arguments for constructing a Partition resource.
type PartitionArgs struct {
	CatalogId      pulumi.StringInput
	DatabaseName   pulumi.StringInput
	PartitionInput PartitionInputTypeInput
	TableName      pulumi.StringInput
}

func (PartitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*partitionArgs)(nil)).Elem()
}

type PartitionInput interface {
	pulumi.Input

	ToPartitionOutput() PartitionOutput
	ToPartitionOutputWithContext(ctx context.Context) PartitionOutput
}

func (*Partition) ElementType() reflect.Type {
	return reflect.TypeOf((**Partition)(nil)).Elem()
}

func (i *Partition) ToPartitionOutput() PartitionOutput {
	return i.ToPartitionOutputWithContext(context.Background())
}

func (i *Partition) ToPartitionOutputWithContext(ctx context.Context) PartitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartitionOutput)
}

type PartitionOutput struct{ *pulumi.OutputState }

func (PartitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Partition)(nil)).Elem()
}

func (o PartitionOutput) ToPartitionOutput() PartitionOutput {
	return o
}

func (o PartitionOutput) ToPartitionOutputWithContext(ctx context.Context) PartitionOutput {
	return o
}

func (o PartitionOutput) CatalogId() pulumi.StringOutput {
	return o.ApplyT(func(v *Partition) pulumi.StringOutput { return v.CatalogId }).(pulumi.StringOutput)
}

func (o PartitionOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *Partition) pulumi.StringOutput { return v.DatabaseName }).(pulumi.StringOutput)
}

func (o PartitionOutput) PartitionInput() PartitionInputTypeOutput {
	return o.ApplyT(func(v *Partition) PartitionInputTypeOutput { return v.PartitionInput }).(PartitionInputTypeOutput)
}

func (o PartitionOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v *Partition) pulumi.StringOutput { return v.TableName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PartitionInput)(nil)).Elem(), &Partition{})
	pulumi.RegisterOutputType(PartitionOutput{})
}
