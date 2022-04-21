// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package neptune

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Neptune::DBClusterParameterGroup
//
// Deprecated: DBClusterParameterGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type DBClusterParameterGroup struct {
	pulumi.CustomResourceState

	Description pulumi.StringOutput                   `pulumi:"description"`
	Family      pulumi.StringOutput                   `pulumi:"family"`
	Name        pulumi.StringPtrOutput                `pulumi:"name"`
	Parameters  pulumi.AnyOutput                      `pulumi:"parameters"`
	Tags        DBClusterParameterGroupTagArrayOutput `pulumi:"tags"`
}

// NewDBClusterParameterGroup registers a new resource with the given unique name, arguments, and options.
func NewDBClusterParameterGroup(ctx *pulumi.Context,
	name string, args *DBClusterParameterGroupArgs, opts ...pulumi.ResourceOption) (*DBClusterParameterGroup, error) {
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
	var resource DBClusterParameterGroup
	err := ctx.RegisterResource("aws-native:neptune:DBClusterParameterGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDBClusterParameterGroup gets an existing DBClusterParameterGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDBClusterParameterGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DBClusterParameterGroupState, opts ...pulumi.ResourceOption) (*DBClusterParameterGroup, error) {
	var resource DBClusterParameterGroup
	err := ctx.ReadResource("aws-native:neptune:DBClusterParameterGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DBClusterParameterGroup resources.
type dbclusterParameterGroupState struct {
}

type DBClusterParameterGroupState struct {
}

func (DBClusterParameterGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbclusterParameterGroupState)(nil)).Elem()
}

type dbclusterParameterGroupArgs struct {
	Description string                       `pulumi:"description"`
	Family      string                       `pulumi:"family"`
	Name        *string                      `pulumi:"name"`
	Parameters  interface{}                  `pulumi:"parameters"`
	Tags        []DBClusterParameterGroupTag `pulumi:"tags"`
}

// The set of arguments for constructing a DBClusterParameterGroup resource.
type DBClusterParameterGroupArgs struct {
	Description pulumi.StringInput
	Family      pulumi.StringInput
	Name        pulumi.StringPtrInput
	Parameters  pulumi.Input
	Tags        DBClusterParameterGroupTagArrayInput
}

func (DBClusterParameterGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbclusterParameterGroupArgs)(nil)).Elem()
}

type DBClusterParameterGroupInput interface {
	pulumi.Input

	ToDBClusterParameterGroupOutput() DBClusterParameterGroupOutput
	ToDBClusterParameterGroupOutputWithContext(ctx context.Context) DBClusterParameterGroupOutput
}

func (*DBClusterParameterGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**DBClusterParameterGroup)(nil)).Elem()
}

func (i *DBClusterParameterGroup) ToDBClusterParameterGroupOutput() DBClusterParameterGroupOutput {
	return i.ToDBClusterParameterGroupOutputWithContext(context.Background())
}

func (i *DBClusterParameterGroup) ToDBClusterParameterGroupOutputWithContext(ctx context.Context) DBClusterParameterGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBClusterParameterGroupOutput)
}

type DBClusterParameterGroupOutput struct{ *pulumi.OutputState }

func (DBClusterParameterGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DBClusterParameterGroup)(nil)).Elem()
}

func (o DBClusterParameterGroupOutput) ToDBClusterParameterGroupOutput() DBClusterParameterGroupOutput {
	return o
}

func (o DBClusterParameterGroupOutput) ToDBClusterParameterGroupOutputWithContext(ctx context.Context) DBClusterParameterGroupOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DBClusterParameterGroupInput)(nil)).Elem(), &DBClusterParameterGroup{})
	pulumi.RegisterOutputType(DBClusterParameterGroupOutput{})
}
