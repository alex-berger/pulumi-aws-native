// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::WAFRegional::SqlInjectionMatchSet
//
// Deprecated: SqlInjectionMatchSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type SqlInjectionMatchSet struct {
	pulumi.CustomResourceState

	Name                    pulumi.StringOutput                                   `pulumi:"name"`
	SqlInjectionMatchTuples SqlInjectionMatchSetSqlInjectionMatchTupleArrayOutput `pulumi:"sqlInjectionMatchTuples"`
}

// NewSqlInjectionMatchSet registers a new resource with the given unique name, arguments, and options.
func NewSqlInjectionMatchSet(ctx *pulumi.Context,
	name string, args *SqlInjectionMatchSetArgs, opts ...pulumi.ResourceOption) (*SqlInjectionMatchSet, error) {
	if args == nil {
		args = &SqlInjectionMatchSetArgs{}
	}

	var resource SqlInjectionMatchSet
	err := ctx.RegisterResource("aws-native:wafregional:SqlInjectionMatchSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlInjectionMatchSet gets an existing SqlInjectionMatchSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlInjectionMatchSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlInjectionMatchSetState, opts ...pulumi.ResourceOption) (*SqlInjectionMatchSet, error) {
	var resource SqlInjectionMatchSet
	err := ctx.ReadResource("aws-native:wafregional:SqlInjectionMatchSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlInjectionMatchSet resources.
type sqlInjectionMatchSetState struct {
}

type SqlInjectionMatchSetState struct {
}

func (SqlInjectionMatchSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlInjectionMatchSetState)(nil)).Elem()
}

type sqlInjectionMatchSetArgs struct {
	Name                    *string                                      `pulumi:"name"`
	SqlInjectionMatchTuples []SqlInjectionMatchSetSqlInjectionMatchTuple `pulumi:"sqlInjectionMatchTuples"`
}

// The set of arguments for constructing a SqlInjectionMatchSet resource.
type SqlInjectionMatchSetArgs struct {
	Name                    pulumi.StringPtrInput
	SqlInjectionMatchTuples SqlInjectionMatchSetSqlInjectionMatchTupleArrayInput
}

func (SqlInjectionMatchSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlInjectionMatchSetArgs)(nil)).Elem()
}

type SqlInjectionMatchSetInput interface {
	pulumi.Input

	ToSqlInjectionMatchSetOutput() SqlInjectionMatchSetOutput
	ToSqlInjectionMatchSetOutputWithContext(ctx context.Context) SqlInjectionMatchSetOutput
}

func (*SqlInjectionMatchSet) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlInjectionMatchSet)(nil)).Elem()
}

func (i *SqlInjectionMatchSet) ToSqlInjectionMatchSetOutput() SqlInjectionMatchSetOutput {
	return i.ToSqlInjectionMatchSetOutputWithContext(context.Background())
}

func (i *SqlInjectionMatchSet) ToSqlInjectionMatchSetOutputWithContext(ctx context.Context) SqlInjectionMatchSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlInjectionMatchSetOutput)
}

type SqlInjectionMatchSetOutput struct{ *pulumi.OutputState }

func (SqlInjectionMatchSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlInjectionMatchSet)(nil)).Elem()
}

func (o SqlInjectionMatchSetOutput) ToSqlInjectionMatchSetOutput() SqlInjectionMatchSetOutput {
	return o
}

func (o SqlInjectionMatchSetOutput) ToSqlInjectionMatchSetOutputWithContext(ctx context.Context) SqlInjectionMatchSetOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SqlInjectionMatchSetInput)(nil)).Elem(), &SqlInjectionMatchSet{})
	pulumi.RegisterOutputType(SqlInjectionMatchSetOutput{})
}
