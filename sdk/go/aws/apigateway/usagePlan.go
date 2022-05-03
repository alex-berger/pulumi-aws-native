// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGateway::UsagePlan
type UsagePlan struct {
	pulumi.CustomResourceState

	// The API stages to associate with this usage plan.
	ApiStages UsagePlanApiStageArrayOutput `pulumi:"apiStages"`
	// A description of the usage plan.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Configures the number of requests that users can make within a given interval.
	Quota UsagePlanQuotaSettingsPtrOutput `pulumi:"quota"`
	// An array of arbitrary tags (key-value pairs) to associate with the usage plan.
	Tags UsagePlanTagArrayOutput `pulumi:"tags"`
	// Configures the overall request rate (average requests per second) and burst capacity.
	Throttle UsagePlanThrottleSettingsPtrOutput `pulumi:"throttle"`
	// A name for the usage plan.
	UsagePlanName pulumi.StringPtrOutput `pulumi:"usagePlanName"`
}

// NewUsagePlan registers a new resource with the given unique name, arguments, and options.
func NewUsagePlan(ctx *pulumi.Context,
	name string, args *UsagePlanArgs, opts ...pulumi.ResourceOption) (*UsagePlan, error) {
	if args == nil {
		args = &UsagePlanArgs{}
	}

	var resource UsagePlan
	err := ctx.RegisterResource("aws-native:apigateway:UsagePlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUsagePlan gets an existing UsagePlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUsagePlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UsagePlanState, opts ...pulumi.ResourceOption) (*UsagePlan, error) {
	var resource UsagePlan
	err := ctx.ReadResource("aws-native:apigateway:UsagePlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UsagePlan resources.
type usagePlanState struct {
}

type UsagePlanState struct {
}

func (UsagePlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*usagePlanState)(nil)).Elem()
}

type usagePlanArgs struct {
	// The API stages to associate with this usage plan.
	ApiStages []UsagePlanApiStage `pulumi:"apiStages"`
	// A description of the usage plan.
	Description *string `pulumi:"description"`
	// Configures the number of requests that users can make within a given interval.
	Quota *UsagePlanQuotaSettings `pulumi:"quota"`
	// An array of arbitrary tags (key-value pairs) to associate with the usage plan.
	Tags []UsagePlanTag `pulumi:"tags"`
	// Configures the overall request rate (average requests per second) and burst capacity.
	Throttle *UsagePlanThrottleSettings `pulumi:"throttle"`
	// A name for the usage plan.
	UsagePlanName *string `pulumi:"usagePlanName"`
}

// The set of arguments for constructing a UsagePlan resource.
type UsagePlanArgs struct {
	// The API stages to associate with this usage plan.
	ApiStages UsagePlanApiStageArrayInput
	// A description of the usage plan.
	Description pulumi.StringPtrInput
	// Configures the number of requests that users can make within a given interval.
	Quota UsagePlanQuotaSettingsPtrInput
	// An array of arbitrary tags (key-value pairs) to associate with the usage plan.
	Tags UsagePlanTagArrayInput
	// Configures the overall request rate (average requests per second) and burst capacity.
	Throttle UsagePlanThrottleSettingsPtrInput
	// A name for the usage plan.
	UsagePlanName pulumi.StringPtrInput
}

func (UsagePlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*usagePlanArgs)(nil)).Elem()
}

type UsagePlanInput interface {
	pulumi.Input

	ToUsagePlanOutput() UsagePlanOutput
	ToUsagePlanOutputWithContext(ctx context.Context) UsagePlanOutput
}

func (*UsagePlan) ElementType() reflect.Type {
	return reflect.TypeOf((**UsagePlan)(nil)).Elem()
}

func (i *UsagePlan) ToUsagePlanOutput() UsagePlanOutput {
	return i.ToUsagePlanOutputWithContext(context.Background())
}

func (i *UsagePlan) ToUsagePlanOutputWithContext(ctx context.Context) UsagePlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsagePlanOutput)
}

type UsagePlanOutput struct{ *pulumi.OutputState }

func (UsagePlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UsagePlan)(nil)).Elem()
}

func (o UsagePlanOutput) ToUsagePlanOutput() UsagePlanOutput {
	return o
}

func (o UsagePlanOutput) ToUsagePlanOutputWithContext(ctx context.Context) UsagePlanOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UsagePlanInput)(nil)).Elem(), &UsagePlan{})
	pulumi.RegisterOutputType(UsagePlanOutput{})
}
