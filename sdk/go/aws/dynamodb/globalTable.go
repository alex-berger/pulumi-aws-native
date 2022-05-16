// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dynamodb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Version: None. Resource Type definition for AWS::DynamoDB::GlobalTable
type GlobalTable struct {
	pulumi.CustomResourceState

	Arn                                pulumi.StringOutput                                    `pulumi:"arn"`
	AttributeDefinitions               GlobalTableAttributeDefinitionArrayOutput              `pulumi:"attributeDefinitions"`
	BillingMode                        pulumi.StringPtrOutput                                 `pulumi:"billingMode"`
	GlobalSecondaryIndexes             GlobalTableGlobalSecondaryIndexArrayOutput             `pulumi:"globalSecondaryIndexes"`
	KeySchema                          GlobalTableKeySchemaArrayOutput                        `pulumi:"keySchema"`
	LocalSecondaryIndexes              GlobalTableLocalSecondaryIndexArrayOutput              `pulumi:"localSecondaryIndexes"`
	Replicas                           GlobalTableReplicaSpecificationArrayOutput             `pulumi:"replicas"`
	SSESpecification                   GlobalTableSSESpecificationPtrOutput                   `pulumi:"sSESpecification"`
	StreamArn                          pulumi.StringOutput                                    `pulumi:"streamArn"`
	StreamSpecification                GlobalTableStreamSpecificationPtrOutput                `pulumi:"streamSpecification"`
	TableId                            pulumi.StringOutput                                    `pulumi:"tableId"`
	TableName                          pulumi.StringPtrOutput                                 `pulumi:"tableName"`
	TimeToLiveSpecification            GlobalTableTimeToLiveSpecificationPtrOutput            `pulumi:"timeToLiveSpecification"`
	WriteProvisionedThroughputSettings GlobalTableWriteProvisionedThroughputSettingsPtrOutput `pulumi:"writeProvisionedThroughputSettings"`
}

// NewGlobalTable registers a new resource with the given unique name, arguments, and options.
func NewGlobalTable(ctx *pulumi.Context,
	name string, args *GlobalTableArgs, opts ...pulumi.ResourceOption) (*GlobalTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AttributeDefinitions == nil {
		return nil, errors.New("invalid value for required argument 'AttributeDefinitions'")
	}
	if args.KeySchema == nil {
		return nil, errors.New("invalid value for required argument 'KeySchema'")
	}
	if args.Replicas == nil {
		return nil, errors.New("invalid value for required argument 'Replicas'")
	}
	var resource GlobalTable
	err := ctx.RegisterResource("aws-native:dynamodb:GlobalTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobalTable gets an existing GlobalTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalTableState, opts ...pulumi.ResourceOption) (*GlobalTable, error) {
	var resource GlobalTable
	err := ctx.ReadResource("aws-native:dynamodb:GlobalTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GlobalTable resources.
type globalTableState struct {
}

type GlobalTableState struct {
}

func (GlobalTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalTableState)(nil)).Elem()
}

type globalTableArgs struct {
	AttributeDefinitions               []GlobalTableAttributeDefinition               `pulumi:"attributeDefinitions"`
	BillingMode                        *string                                        `pulumi:"billingMode"`
	GlobalSecondaryIndexes             []GlobalTableGlobalSecondaryIndex              `pulumi:"globalSecondaryIndexes"`
	KeySchema                          []GlobalTableKeySchema                         `pulumi:"keySchema"`
	LocalSecondaryIndexes              []GlobalTableLocalSecondaryIndex               `pulumi:"localSecondaryIndexes"`
	Replicas                           []GlobalTableReplicaSpecification              `pulumi:"replicas"`
	SSESpecification                   *GlobalTableSSESpecification                   `pulumi:"sSESpecification"`
	StreamSpecification                *GlobalTableStreamSpecification                `pulumi:"streamSpecification"`
	TableName                          *string                                        `pulumi:"tableName"`
	TimeToLiveSpecification            *GlobalTableTimeToLiveSpecification            `pulumi:"timeToLiveSpecification"`
	WriteProvisionedThroughputSettings *GlobalTableWriteProvisionedThroughputSettings `pulumi:"writeProvisionedThroughputSettings"`
}

// The set of arguments for constructing a GlobalTable resource.
type GlobalTableArgs struct {
	AttributeDefinitions               GlobalTableAttributeDefinitionArrayInput
	BillingMode                        pulumi.StringPtrInput
	GlobalSecondaryIndexes             GlobalTableGlobalSecondaryIndexArrayInput
	KeySchema                          GlobalTableKeySchemaArrayInput
	LocalSecondaryIndexes              GlobalTableLocalSecondaryIndexArrayInput
	Replicas                           GlobalTableReplicaSpecificationArrayInput
	SSESpecification                   GlobalTableSSESpecificationPtrInput
	StreamSpecification                GlobalTableStreamSpecificationPtrInput
	TableName                          pulumi.StringPtrInput
	TimeToLiveSpecification            GlobalTableTimeToLiveSpecificationPtrInput
	WriteProvisionedThroughputSettings GlobalTableWriteProvisionedThroughputSettingsPtrInput
}

func (GlobalTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalTableArgs)(nil)).Elem()
}

type GlobalTableInput interface {
	pulumi.Input

	ToGlobalTableOutput() GlobalTableOutput
	ToGlobalTableOutputWithContext(ctx context.Context) GlobalTableOutput
}

func (*GlobalTable) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalTable)(nil)).Elem()
}

func (i *GlobalTable) ToGlobalTableOutput() GlobalTableOutput {
	return i.ToGlobalTableOutputWithContext(context.Background())
}

func (i *GlobalTable) ToGlobalTableOutputWithContext(ctx context.Context) GlobalTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalTableOutput)
}

type GlobalTableOutput struct{ *pulumi.OutputState }

func (GlobalTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalTable)(nil)).Elem()
}

func (o GlobalTableOutput) ToGlobalTableOutput() GlobalTableOutput {
	return o
}

func (o GlobalTableOutput) ToGlobalTableOutputWithContext(ctx context.Context) GlobalTableOutput {
	return o
}

func (o GlobalTableOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalTable) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o GlobalTableOutput) AttributeDefinitions() GlobalTableAttributeDefinitionArrayOutput {
	return o.ApplyT(func(v *GlobalTable) GlobalTableAttributeDefinitionArrayOutput { return v.AttributeDefinitions }).(GlobalTableAttributeDefinitionArrayOutput)
}

func (o GlobalTableOutput) BillingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GlobalTable) pulumi.StringPtrOutput { return v.BillingMode }).(pulumi.StringPtrOutput)
}

func (o GlobalTableOutput) GlobalSecondaryIndexes() GlobalTableGlobalSecondaryIndexArrayOutput {
	return o.ApplyT(func(v *GlobalTable) GlobalTableGlobalSecondaryIndexArrayOutput { return v.GlobalSecondaryIndexes }).(GlobalTableGlobalSecondaryIndexArrayOutput)
}

func (o GlobalTableOutput) KeySchema() GlobalTableKeySchemaArrayOutput {
	return o.ApplyT(func(v *GlobalTable) GlobalTableKeySchemaArrayOutput { return v.KeySchema }).(GlobalTableKeySchemaArrayOutput)
}

func (o GlobalTableOutput) LocalSecondaryIndexes() GlobalTableLocalSecondaryIndexArrayOutput {
	return o.ApplyT(func(v *GlobalTable) GlobalTableLocalSecondaryIndexArrayOutput { return v.LocalSecondaryIndexes }).(GlobalTableLocalSecondaryIndexArrayOutput)
}

func (o GlobalTableOutput) Replicas() GlobalTableReplicaSpecificationArrayOutput {
	return o.ApplyT(func(v *GlobalTable) GlobalTableReplicaSpecificationArrayOutput { return v.Replicas }).(GlobalTableReplicaSpecificationArrayOutput)
}

func (o GlobalTableOutput) SSESpecification() GlobalTableSSESpecificationPtrOutput {
	return o.ApplyT(func(v *GlobalTable) GlobalTableSSESpecificationPtrOutput { return v.SSESpecification }).(GlobalTableSSESpecificationPtrOutput)
}

func (o GlobalTableOutput) StreamArn() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalTable) pulumi.StringOutput { return v.StreamArn }).(pulumi.StringOutput)
}

func (o GlobalTableOutput) StreamSpecification() GlobalTableStreamSpecificationPtrOutput {
	return o.ApplyT(func(v *GlobalTable) GlobalTableStreamSpecificationPtrOutput { return v.StreamSpecification }).(GlobalTableStreamSpecificationPtrOutput)
}

func (o GlobalTableOutput) TableId() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalTable) pulumi.StringOutput { return v.TableId }).(pulumi.StringOutput)
}

func (o GlobalTableOutput) TableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GlobalTable) pulumi.StringPtrOutput { return v.TableName }).(pulumi.StringPtrOutput)
}

func (o GlobalTableOutput) TimeToLiveSpecification() GlobalTableTimeToLiveSpecificationPtrOutput {
	return o.ApplyT(func(v *GlobalTable) GlobalTableTimeToLiveSpecificationPtrOutput { return v.TimeToLiveSpecification }).(GlobalTableTimeToLiveSpecificationPtrOutput)
}

func (o GlobalTableOutput) WriteProvisionedThroughputSettings() GlobalTableWriteProvisionedThroughputSettingsPtrOutput {
	return o.ApplyT(func(v *GlobalTable) GlobalTableWriteProvisionedThroughputSettingsPtrOutput {
		return v.WriteProvisionedThroughputSettings
	}).(GlobalTableWriteProvisionedThroughputSettingsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalTableInput)(nil)).Elem(), &GlobalTable{})
	pulumi.RegisterOutputType(GlobalTableOutput{})
}
