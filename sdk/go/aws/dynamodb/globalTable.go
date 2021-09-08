// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dynamodb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html
type GlobalTable struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-attributedefinitions
	AttributeDefinitions GlobalTableAttributeDefinitionArrayOutput `pulumi:"attributeDefinitions"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-billingmode
	BillingMode pulumi.StringPtrOutput `pulumi:"billingMode"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-globalsecondaryindexes
	GlobalSecondaryIndexes GlobalTableGlobalSecondaryIndexArrayOutput `pulumi:"globalSecondaryIndexes"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-keyschema
	KeySchema GlobalTableKeySchemaArrayOutput `pulumi:"keySchema"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-localsecondaryindexes
	LocalSecondaryIndexes GlobalTableLocalSecondaryIndexArrayOutput `pulumi:"localSecondaryIndexes"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-replicas
	Replicas GlobalTableReplicaSpecificationArrayOutput `pulumi:"replicas"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-ssespecification
	SSESpecification GlobalTableSSESpecificationPtrOutput `pulumi:"sSESpecification"`
	StreamArn        pulumi.StringOutput                  `pulumi:"streamArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-streamspecification
	StreamSpecification GlobalTableStreamSpecificationPtrOutput `pulumi:"streamSpecification"`
	TableId             pulumi.StringOutput                     `pulumi:"tableId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-tablename
	TableName pulumi.StringPtrOutput `pulumi:"tableName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-timetolivespecification
	TimeToLiveSpecification GlobalTableTimeToLiveSpecificationPtrOutput `pulumi:"timeToLiveSpecification"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-writeprovisionedthroughputsettings
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
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-attributedefinitions
	AttributeDefinitions []GlobalTableAttributeDefinition `pulumi:"attributeDefinitions"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-billingmode
	BillingMode *string `pulumi:"billingMode"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-globalsecondaryindexes
	GlobalSecondaryIndexes []GlobalTableGlobalSecondaryIndex `pulumi:"globalSecondaryIndexes"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-keyschema
	KeySchema []GlobalTableKeySchema `pulumi:"keySchema"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-localsecondaryindexes
	LocalSecondaryIndexes []GlobalTableLocalSecondaryIndex `pulumi:"localSecondaryIndexes"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-replicas
	Replicas []GlobalTableReplicaSpecification `pulumi:"replicas"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-ssespecification
	SSESpecification *GlobalTableSSESpecification `pulumi:"sSESpecification"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-streamspecification
	StreamSpecification *GlobalTableStreamSpecification `pulumi:"streamSpecification"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-tablename
	TableName *string `pulumi:"tableName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-timetolivespecification
	TimeToLiveSpecification *GlobalTableTimeToLiveSpecification `pulumi:"timeToLiveSpecification"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-writeprovisionedthroughputsettings
	WriteProvisionedThroughputSettings *GlobalTableWriteProvisionedThroughputSettings `pulumi:"writeProvisionedThroughputSettings"`
}

// The set of arguments for constructing a GlobalTable resource.
type GlobalTableArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-attributedefinitions
	AttributeDefinitions GlobalTableAttributeDefinitionArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-billingmode
	BillingMode pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-globalsecondaryindexes
	GlobalSecondaryIndexes GlobalTableGlobalSecondaryIndexArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-keyschema
	KeySchema GlobalTableKeySchemaArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-localsecondaryindexes
	LocalSecondaryIndexes GlobalTableLocalSecondaryIndexArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-replicas
	Replicas GlobalTableReplicaSpecificationArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-ssespecification
	SSESpecification GlobalTableSSESpecificationPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-streamspecification
	StreamSpecification GlobalTableStreamSpecificationPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-tablename
	TableName pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-timetolivespecification
	TimeToLiveSpecification GlobalTableTimeToLiveSpecificationPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-writeprovisionedthroughputsettings
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
	return reflect.TypeOf((*GlobalTable)(nil))
}

func (i *GlobalTable) ToGlobalTableOutput() GlobalTableOutput {
	return i.ToGlobalTableOutputWithContext(context.Background())
}

func (i *GlobalTable) ToGlobalTableOutputWithContext(ctx context.Context) GlobalTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalTableOutput)
}

type GlobalTableOutput struct{ *pulumi.OutputState }

func (GlobalTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalTable)(nil))
}

func (o GlobalTableOutput) ToGlobalTableOutput() GlobalTableOutput {
	return o
}

func (o GlobalTableOutput) ToGlobalTableOutputWithContext(ctx context.Context) GlobalTableOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GlobalTableOutput{})
}