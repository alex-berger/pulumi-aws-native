// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kendra

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Kendra DataSource
type DataSource struct {
	pulumi.CustomResourceState

	Arn                                   pulumi.StringOutput                                      `pulumi:"arn"`
	CustomDocumentEnrichmentConfiguration DataSourceCustomDocumentEnrichmentConfigurationPtrOutput `pulumi:"customDocumentEnrichmentConfiguration"`
	DataSourceConfiguration               DataSourceConfigurationPtrOutput                         `pulumi:"dataSourceConfiguration"`
	Description                           pulumi.StringPtrOutput                                   `pulumi:"description"`
	IndexId                               pulumi.StringOutput                                      `pulumi:"indexId"`
	Name                                  pulumi.StringOutput                                      `pulumi:"name"`
	RoleArn                               pulumi.StringPtrOutput                                   `pulumi:"roleArn"`
	Schedule                              pulumi.StringPtrOutput                                   `pulumi:"schedule"`
	// Tags for labeling the data source
	Tags DataSourceTagArrayOutput `pulumi:"tags"`
	Type DataSourceTypeOutput     `pulumi:"type"`
}

// NewDataSource registers a new resource with the given unique name, arguments, and options.
func NewDataSource(ctx *pulumi.Context,
	name string, args *DataSourceArgs, opts ...pulumi.ResourceOption) (*DataSource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IndexId == nil {
		return nil, errors.New("invalid value for required argument 'IndexId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource DataSource
	err := ctx.RegisterResource("aws-native:kendra:DataSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataSource gets an existing DataSource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataSourceState, opts ...pulumi.ResourceOption) (*DataSource, error) {
	var resource DataSource
	err := ctx.ReadResource("aws-native:kendra:DataSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataSource resources.
type dataSourceState struct {
}

type DataSourceState struct {
}

func (DataSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceState)(nil)).Elem()
}

type dataSourceArgs struct {
	CustomDocumentEnrichmentConfiguration *DataSourceCustomDocumentEnrichmentConfiguration `pulumi:"customDocumentEnrichmentConfiguration"`
	DataSourceConfiguration               *DataSourceConfiguration                         `pulumi:"dataSourceConfiguration"`
	Description                           *string                                          `pulumi:"description"`
	IndexId                               string                                           `pulumi:"indexId"`
	Name                                  *string                                          `pulumi:"name"`
	RoleArn                               *string                                          `pulumi:"roleArn"`
	Schedule                              *string                                          `pulumi:"schedule"`
	// Tags for labeling the data source
	Tags []DataSourceTag `pulumi:"tags"`
	Type DataSourceType  `pulumi:"type"`
}

// The set of arguments for constructing a DataSource resource.
type DataSourceArgs struct {
	CustomDocumentEnrichmentConfiguration DataSourceCustomDocumentEnrichmentConfigurationPtrInput
	DataSourceConfiguration               DataSourceConfigurationPtrInput
	Description                           pulumi.StringPtrInput
	IndexId                               pulumi.StringInput
	Name                                  pulumi.StringPtrInput
	RoleArn                               pulumi.StringPtrInput
	Schedule                              pulumi.StringPtrInput
	// Tags for labeling the data source
	Tags DataSourceTagArrayInput
	Type DataSourceTypeInput
}

func (DataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceArgs)(nil)).Elem()
}

type DataSourceInput interface {
	pulumi.Input

	ToDataSourceOutput() DataSourceOutput
	ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput
}

func (*DataSource) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSource)(nil)).Elem()
}

func (i *DataSource) ToDataSourceOutput() DataSourceOutput {
	return i.ToDataSourceOutputWithContext(context.Background())
}

func (i *DataSource) ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceOutput)
}

type DataSourceOutput struct{ *pulumi.OutputState }

func (DataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSource)(nil)).Elem()
}

func (o DataSourceOutput) ToDataSourceOutput() DataSourceOutput {
	return o
}

func (o DataSourceOutput) ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataSourceInput)(nil)).Elem(), &DataSource{})
	pulumi.RegisterOutputType(DataSourceOutput{})
}
