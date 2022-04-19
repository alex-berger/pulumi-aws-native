// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package neptune

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Neptune::DBParameterGroup
//
// Deprecated: DBParameterGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type DBParameterGroup struct {
	pulumi.CustomResourceState

	Description pulumi.StringOutput            `pulumi:"description"`
	Family      pulumi.StringOutput            `pulumi:"family"`
	Name        pulumi.StringPtrOutput         `pulumi:"name"`
	Parameters  pulumi.AnyOutput               `pulumi:"parameters"`
	Tags        DBParameterGroupTagArrayOutput `pulumi:"tags"`
}

// NewDBParameterGroup registers a new resource with the given unique name, arguments, and options.
func NewDBParameterGroup(ctx *pulumi.Context,
	name string, args *DBParameterGroupArgs, opts ...pulumi.ResourceOption) (*DBParameterGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Family == nil {
		return nil, errors.New("invalid value for required argument 'Family'")
	}
	if args.Parameters == nil {
		return nil, errors.New("invalid value for required argument 'Parameters'")
	}
	var resource DBParameterGroup
	err := ctx.RegisterResource("aws-native:neptune:DBParameterGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDBParameterGroup gets an existing DBParameterGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDBParameterGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DBParameterGroupState, opts ...pulumi.ResourceOption) (*DBParameterGroup, error) {
	var resource DBParameterGroup
	err := ctx.ReadResource("aws-native:neptune:DBParameterGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DBParameterGroup resources.
type dbparameterGroupState struct {
}

type DBParameterGroupState struct {
}

func (DBParameterGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbparameterGroupState)(nil)).Elem()
}

type dbparameterGroupArgs struct {
	Description string                `pulumi:"description"`
	Family      string                `pulumi:"family"`
	Name        *string               `pulumi:"name"`
	Parameters  interface{}           `pulumi:"parameters"`
	Tags        []DBParameterGroupTag `pulumi:"tags"`
}

// The set of arguments for constructing a DBParameterGroup resource.
type DBParameterGroupArgs struct {
	Description pulumi.StringInput
	Family      pulumi.StringInput
	Name        pulumi.StringPtrInput
	Parameters  pulumi.Input
	Tags        DBParameterGroupTagArrayInput
}

func (DBParameterGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbparameterGroupArgs)(nil)).Elem()
}

type DBParameterGroupInput interface {
	pulumi.Input

	ToDBParameterGroupOutput() DBParameterGroupOutput
	ToDBParameterGroupOutputWithContext(ctx context.Context) DBParameterGroupOutput
}

func (*DBParameterGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**DBParameterGroup)(nil)).Elem()
}

func (i *DBParameterGroup) ToDBParameterGroupOutput() DBParameterGroupOutput {
	return i.ToDBParameterGroupOutputWithContext(context.Background())
}

func (i *DBParameterGroup) ToDBParameterGroupOutputWithContext(ctx context.Context) DBParameterGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBParameterGroupOutput)
}

type DBParameterGroupOutput struct{ *pulumi.OutputState }

func (DBParameterGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DBParameterGroup)(nil)).Elem()
}

func (o DBParameterGroupOutput) ToDBParameterGroupOutput() DBParameterGroupOutput {
	return o
}

func (o DBParameterGroupOutput) ToDBParameterGroupOutputWithContext(ctx context.Context) DBParameterGroupOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DBParameterGroupInput)(nil)).Elem(), &DBParameterGroup{})
	pulumi.RegisterOutputType(DBParameterGroupOutput{})
}
