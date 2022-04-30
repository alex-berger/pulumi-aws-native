// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotanalytics

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IoTAnalytics::Dataset
type Dataset struct {
	pulumi.CustomResourceState

	Actions                 DatasetActionArrayOutput                `pulumi:"actions"`
	ContentDeliveryRules    DatasetContentDeliveryRuleArrayOutput   `pulumi:"contentDeliveryRules"`
	DatasetName             pulumi.StringPtrOutput                  `pulumi:"datasetName"`
	LateDataRules           DatasetLateDataRuleArrayOutput          `pulumi:"lateDataRules"`
	RetentionPeriod         DatasetRetentionPeriodPtrOutput         `pulumi:"retentionPeriod"`
	Tags                    DatasetTagArrayOutput                   `pulumi:"tags"`
	Triggers                DatasetTriggerArrayOutput               `pulumi:"triggers"`
	VersioningConfiguration DatasetVersioningConfigurationPtrOutput `pulumi:"versioningConfiguration"`
}

// NewDataset registers a new resource with the given unique name, arguments, and options.
func NewDataset(ctx *pulumi.Context,
	name string, args *DatasetArgs, opts ...pulumi.ResourceOption) (*Dataset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	var resource Dataset
	err := ctx.RegisterResource("aws-native:iotanalytics:Dataset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataset gets an existing Dataset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetState, opts ...pulumi.ResourceOption) (*Dataset, error) {
	var resource Dataset
	err := ctx.ReadResource("aws-native:iotanalytics:Dataset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dataset resources.
type datasetState struct {
}

type DatasetState struct {
}

func (DatasetState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetState)(nil)).Elem()
}

type datasetArgs struct {
	Actions                 []DatasetAction                 `pulumi:"actions"`
	ContentDeliveryRules    []DatasetContentDeliveryRule    `pulumi:"contentDeliveryRules"`
	DatasetName             *string                         `pulumi:"datasetName"`
	LateDataRules           []DatasetLateDataRule           `pulumi:"lateDataRules"`
	RetentionPeriod         *DatasetRetentionPeriod         `pulumi:"retentionPeriod"`
	Tags                    []DatasetTag                    `pulumi:"tags"`
	Triggers                []DatasetTrigger                `pulumi:"triggers"`
	VersioningConfiguration *DatasetVersioningConfiguration `pulumi:"versioningConfiguration"`
}

// The set of arguments for constructing a Dataset resource.
type DatasetArgs struct {
	Actions                 DatasetActionArrayInput
	ContentDeliveryRules    DatasetContentDeliveryRuleArrayInput
	DatasetName             pulumi.StringPtrInput
	LateDataRules           DatasetLateDataRuleArrayInput
	RetentionPeriod         DatasetRetentionPeriodPtrInput
	Tags                    DatasetTagArrayInput
	Triggers                DatasetTriggerArrayInput
	VersioningConfiguration DatasetVersioningConfigurationPtrInput
}

func (DatasetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetArgs)(nil)).Elem()
}

type DatasetInput interface {
	pulumi.Input

	ToDatasetOutput() DatasetOutput
	ToDatasetOutputWithContext(ctx context.Context) DatasetOutput
}

func (*Dataset) ElementType() reflect.Type {
	return reflect.TypeOf((**Dataset)(nil)).Elem()
}

func (i *Dataset) ToDatasetOutput() DatasetOutput {
	return i.ToDatasetOutputWithContext(context.Background())
}

func (i *Dataset) ToDatasetOutputWithContext(ctx context.Context) DatasetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetOutput)
}

type DatasetOutput struct{ *pulumi.OutputState }

func (DatasetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dataset)(nil)).Elem()
}

func (o DatasetOutput) ToDatasetOutput() DatasetOutput {
	return o
}

func (o DatasetOutput) ToDatasetOutputWithContext(ctx context.Context) DatasetOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetInput)(nil)).Elem(), &Dataset{})
	pulumi.RegisterOutputType(DatasetOutput{})
}
