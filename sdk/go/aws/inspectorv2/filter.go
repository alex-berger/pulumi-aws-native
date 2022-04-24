// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package inspectorv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Inspector Filter resource schema
type Filter struct {
	pulumi.CustomResourceState

	// Findings filter ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Findings filter description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Findings filter action.
	FilterAction FilterActionOutput `pulumi:"filterAction"`
	// Findings filter criteria.
	FilterCriteria FilterCriteriaOutput `pulumi:"filterCriteria"`
	// Findings filter name.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewFilter registers a new resource with the given unique name, arguments, and options.
func NewFilter(ctx *pulumi.Context,
	name string, args *FilterArgs, opts ...pulumi.ResourceOption) (*Filter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FilterAction == nil {
		return nil, errors.New("invalid value for required argument 'FilterAction'")
	}
	if args.FilterCriteria == nil {
		return nil, errors.New("invalid value for required argument 'FilterCriteria'")
	}
	var resource Filter
	err := ctx.RegisterResource("aws-native:inspectorv2:Filter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFilter gets an existing Filter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FilterState, opts ...pulumi.ResourceOption) (*Filter, error) {
	var resource Filter
	err := ctx.ReadResource("aws-native:inspectorv2:Filter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Filter resources.
type filterState struct {
}

type FilterState struct {
}

func (FilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*filterState)(nil)).Elem()
}

type filterArgs struct {
	// Findings filter description.
	Description *string `pulumi:"description"`
	// Findings filter action.
	FilterAction FilterAction `pulumi:"filterAction"`
	// Findings filter criteria.
	FilterCriteria FilterCriteria `pulumi:"filterCriteria"`
	// Findings filter name.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Filter resource.
type FilterArgs struct {
	// Findings filter description.
	Description pulumi.StringPtrInput
	// Findings filter action.
	FilterAction FilterActionInput
	// Findings filter criteria.
	FilterCriteria FilterCriteriaInput
	// Findings filter name.
	Name pulumi.StringPtrInput
}

func (FilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*filterArgs)(nil)).Elem()
}

type FilterInput interface {
	pulumi.Input

	ToFilterOutput() FilterOutput
	ToFilterOutputWithContext(ctx context.Context) FilterOutput
}

func (*Filter) ElementType() reflect.Type {
	return reflect.TypeOf((**Filter)(nil)).Elem()
}

func (i *Filter) ToFilterOutput() FilterOutput {
	return i.ToFilterOutputWithContext(context.Background())
}

func (i *Filter) ToFilterOutputWithContext(ctx context.Context) FilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterOutput)
}

type FilterOutput struct{ *pulumi.OutputState }

func (FilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Filter)(nil)).Elem()
}

func (o FilterOutput) ToFilterOutput() FilterOutput {
	return o
}

func (o FilterOutput) ToFilterOutputWithContext(ctx context.Context) FilterOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FilterInput)(nil)).Elem(), &Filter{})
	pulumi.RegisterOutputType(FilterOutput{})
}
