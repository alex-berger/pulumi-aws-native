// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package healthlake

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// HealthLake FHIR Datastore
type FHIRDatastore struct {
	pulumi.CustomResourceState

	CreatedAt            FHIRDatastoreCreatedAtOutput            `pulumi:"createdAt"`
	DatastoreArn         pulumi.StringOutput                     `pulumi:"datastoreArn"`
	DatastoreEndpoint    pulumi.StringOutput                     `pulumi:"datastoreEndpoint"`
	DatastoreId          pulumi.StringOutput                     `pulumi:"datastoreId"`
	DatastoreName        pulumi.StringPtrOutput                  `pulumi:"datastoreName"`
	DatastoreStatus      FHIRDatastoreDatastoreStatusOutput      `pulumi:"datastoreStatus"`
	DatastoreTypeVersion FHIRDatastoreDatastoreTypeVersionOutput `pulumi:"datastoreTypeVersion"`
	PreloadDataConfig    FHIRDatastorePreloadDataConfigPtrOutput `pulumi:"preloadDataConfig"`
	SseConfiguration     FHIRDatastoreSseConfigurationPtrOutput  `pulumi:"sseConfiguration"`
	Tags                 FHIRDatastoreTagArrayOutput             `pulumi:"tags"`
}

// NewFHIRDatastore registers a new resource with the given unique name, arguments, and options.
func NewFHIRDatastore(ctx *pulumi.Context,
	name string, args *FHIRDatastoreArgs, opts ...pulumi.ResourceOption) (*FHIRDatastore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatastoreTypeVersion == nil {
		return nil, errors.New("invalid value for required argument 'DatastoreTypeVersion'")
	}
	var resource FHIRDatastore
	err := ctx.RegisterResource("aws-native:healthlake:FHIRDatastore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFHIRDatastore gets an existing FHIRDatastore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFHIRDatastore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FHIRDatastoreState, opts ...pulumi.ResourceOption) (*FHIRDatastore, error) {
	var resource FHIRDatastore
	err := ctx.ReadResource("aws-native:healthlake:FHIRDatastore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FHIRDatastore resources.
type fhirdatastoreState struct {
}

type FHIRDatastoreState struct {
}

func (FHIRDatastoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*fhirdatastoreState)(nil)).Elem()
}

type fhirdatastoreArgs struct {
	DatastoreName        *string                           `pulumi:"datastoreName"`
	DatastoreTypeVersion FHIRDatastoreDatastoreTypeVersion `pulumi:"datastoreTypeVersion"`
	PreloadDataConfig    *FHIRDatastorePreloadDataConfig   `pulumi:"preloadDataConfig"`
	SseConfiguration     *FHIRDatastoreSseConfiguration    `pulumi:"sseConfiguration"`
	Tags                 []FHIRDatastoreTag                `pulumi:"tags"`
}

// The set of arguments for constructing a FHIRDatastore resource.
type FHIRDatastoreArgs struct {
	DatastoreName        pulumi.StringPtrInput
	DatastoreTypeVersion FHIRDatastoreDatastoreTypeVersionInput
	PreloadDataConfig    FHIRDatastorePreloadDataConfigPtrInput
	SseConfiguration     FHIRDatastoreSseConfigurationPtrInput
	Tags                 FHIRDatastoreTagArrayInput
}

func (FHIRDatastoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fhirdatastoreArgs)(nil)).Elem()
}

type FHIRDatastoreInput interface {
	pulumi.Input

	ToFHIRDatastoreOutput() FHIRDatastoreOutput
	ToFHIRDatastoreOutputWithContext(ctx context.Context) FHIRDatastoreOutput
}

func (*FHIRDatastore) ElementType() reflect.Type {
	return reflect.TypeOf((**FHIRDatastore)(nil)).Elem()
}

func (i *FHIRDatastore) ToFHIRDatastoreOutput() FHIRDatastoreOutput {
	return i.ToFHIRDatastoreOutputWithContext(context.Background())
}

func (i *FHIRDatastore) ToFHIRDatastoreOutputWithContext(ctx context.Context) FHIRDatastoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FHIRDatastoreOutput)
}

type FHIRDatastoreOutput struct{ *pulumi.OutputState }

func (FHIRDatastoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FHIRDatastore)(nil)).Elem()
}

func (o FHIRDatastoreOutput) ToFHIRDatastoreOutput() FHIRDatastoreOutput {
	return o
}

func (o FHIRDatastoreOutput) ToFHIRDatastoreOutputWithContext(ctx context.Context) FHIRDatastoreOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FHIRDatastoreInput)(nil)).Elem(), &FHIRDatastore{})
	pulumi.RegisterOutputType(FHIRDatastoreOutput{})
}
