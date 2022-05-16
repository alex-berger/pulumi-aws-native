// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SNS::TopicPolicy
//
// Deprecated: TopicPolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type TopicPolicy struct {
	pulumi.CustomResourceState

	PolicyDocument pulumi.AnyOutput         `pulumi:"policyDocument"`
	Topics         pulumi.StringArrayOutput `pulumi:"topics"`
}

// NewTopicPolicy registers a new resource with the given unique name, arguments, and options.
func NewTopicPolicy(ctx *pulumi.Context,
	name string, args *TopicPolicyArgs, opts ...pulumi.ResourceOption) (*TopicPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyDocument == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDocument'")
	}
	if args.Topics == nil {
		return nil, errors.New("invalid value for required argument 'Topics'")
	}
	var resource TopicPolicy
	err := ctx.RegisterResource("aws-native:sns:TopicPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopicPolicy gets an existing TopicPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopicPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicPolicyState, opts ...pulumi.ResourceOption) (*TopicPolicy, error) {
	var resource TopicPolicy
	err := ctx.ReadResource("aws-native:sns:TopicPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TopicPolicy resources.
type topicPolicyState struct {
}

type TopicPolicyState struct {
}

func (TopicPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicPolicyState)(nil)).Elem()
}

type topicPolicyArgs struct {
	PolicyDocument interface{} `pulumi:"policyDocument"`
	Topics         []string    `pulumi:"topics"`
}

// The set of arguments for constructing a TopicPolicy resource.
type TopicPolicyArgs struct {
	PolicyDocument pulumi.Input
	Topics         pulumi.StringArrayInput
}

func (TopicPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicPolicyArgs)(nil)).Elem()
}

type TopicPolicyInput interface {
	pulumi.Input

	ToTopicPolicyOutput() TopicPolicyOutput
	ToTopicPolicyOutputWithContext(ctx context.Context) TopicPolicyOutput
}

func (*TopicPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicPolicy)(nil)).Elem()
}

func (i *TopicPolicy) ToTopicPolicyOutput() TopicPolicyOutput {
	return i.ToTopicPolicyOutputWithContext(context.Background())
}

func (i *TopicPolicy) ToTopicPolicyOutputWithContext(ctx context.Context) TopicPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicPolicyOutput)
}

type TopicPolicyOutput struct{ *pulumi.OutputState }

func (TopicPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicPolicy)(nil)).Elem()
}

func (o TopicPolicyOutput) ToTopicPolicyOutput() TopicPolicyOutput {
	return o
}

func (o TopicPolicyOutput) ToTopicPolicyOutputWithContext(ctx context.Context) TopicPolicyOutput {
	return o
}

func (o TopicPolicyOutput) PolicyDocument() pulumi.AnyOutput {
	return o.ApplyT(func(v *TopicPolicy) pulumi.AnyOutput { return v.PolicyDocument }).(pulumi.AnyOutput)
}

func (o TopicPolicyOutput) Topics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TopicPolicy) pulumi.StringArrayOutput { return v.Topics }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TopicPolicyInput)(nil)).Elem(), &TopicPolicy{})
	pulumi.RegisterOutputType(TopicPolicyOutput{})
}
