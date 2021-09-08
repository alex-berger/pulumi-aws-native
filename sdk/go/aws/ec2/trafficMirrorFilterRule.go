// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html
type TrafficMirrorFilterRule struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-destinationcidrblock
	DestinationCidrBlock pulumi.StringOutput `pulumi:"destinationCidrBlock"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-destinationportrange
	DestinationPortRange TrafficMirrorFilterRuleTrafficMirrorPortRangePtrOutput `pulumi:"destinationPortRange"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-protocol
	Protocol pulumi.IntPtrOutput `pulumi:"protocol"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-ruleaction
	RuleAction pulumi.StringOutput `pulumi:"ruleAction"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-rulenumber
	RuleNumber pulumi.IntOutput `pulumi:"ruleNumber"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-sourcecidrblock
	SourceCidrBlock pulumi.StringOutput `pulumi:"sourceCidrBlock"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-sourceportrange
	SourcePortRange TrafficMirrorFilterRuleTrafficMirrorPortRangePtrOutput `pulumi:"sourcePortRange"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-trafficdirection
	TrafficDirection pulumi.StringOutput `pulumi:"trafficDirection"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-trafficmirrorfilterid
	TrafficMirrorFilterId pulumi.StringOutput `pulumi:"trafficMirrorFilterId"`
}

// NewTrafficMirrorFilterRule registers a new resource with the given unique name, arguments, and options.
func NewTrafficMirrorFilterRule(ctx *pulumi.Context,
	name string, args *TrafficMirrorFilterRuleArgs, opts ...pulumi.ResourceOption) (*TrafficMirrorFilterRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationCidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'DestinationCidrBlock'")
	}
	if args.RuleAction == nil {
		return nil, errors.New("invalid value for required argument 'RuleAction'")
	}
	if args.RuleNumber == nil {
		return nil, errors.New("invalid value for required argument 'RuleNumber'")
	}
	if args.SourceCidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'SourceCidrBlock'")
	}
	if args.TrafficDirection == nil {
		return nil, errors.New("invalid value for required argument 'TrafficDirection'")
	}
	if args.TrafficMirrorFilterId == nil {
		return nil, errors.New("invalid value for required argument 'TrafficMirrorFilterId'")
	}
	var resource TrafficMirrorFilterRule
	err := ctx.RegisterResource("aws-native:ec2:TrafficMirrorFilterRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficMirrorFilterRule gets an existing TrafficMirrorFilterRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficMirrorFilterRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficMirrorFilterRuleState, opts ...pulumi.ResourceOption) (*TrafficMirrorFilterRule, error) {
	var resource TrafficMirrorFilterRule
	err := ctx.ReadResource("aws-native:ec2:TrafficMirrorFilterRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrafficMirrorFilterRule resources.
type trafficMirrorFilterRuleState struct {
}

type TrafficMirrorFilterRuleState struct {
}

func (TrafficMirrorFilterRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMirrorFilterRuleState)(nil)).Elem()
}

type trafficMirrorFilterRuleArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-description
	Description *string `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-destinationcidrblock
	DestinationCidrBlock string `pulumi:"destinationCidrBlock"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-destinationportrange
	DestinationPortRange *TrafficMirrorFilterRuleTrafficMirrorPortRange `pulumi:"destinationPortRange"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-protocol
	Protocol *int `pulumi:"protocol"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-ruleaction
	RuleAction string `pulumi:"ruleAction"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-rulenumber
	RuleNumber int `pulumi:"ruleNumber"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-sourcecidrblock
	SourceCidrBlock string `pulumi:"sourceCidrBlock"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-sourceportrange
	SourcePortRange *TrafficMirrorFilterRuleTrafficMirrorPortRange `pulumi:"sourcePortRange"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-trafficdirection
	TrafficDirection string `pulumi:"trafficDirection"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-trafficmirrorfilterid
	TrafficMirrorFilterId string `pulumi:"trafficMirrorFilterId"`
}

// The set of arguments for constructing a TrafficMirrorFilterRule resource.
type TrafficMirrorFilterRuleArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-description
	Description pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-destinationcidrblock
	DestinationCidrBlock pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-destinationportrange
	DestinationPortRange TrafficMirrorFilterRuleTrafficMirrorPortRangePtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-protocol
	Protocol pulumi.IntPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-ruleaction
	RuleAction pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-rulenumber
	RuleNumber pulumi.IntInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-sourcecidrblock
	SourceCidrBlock pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-sourceportrange
	SourcePortRange TrafficMirrorFilterRuleTrafficMirrorPortRangePtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-trafficdirection
	TrafficDirection pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-trafficmirrorfilterid
	TrafficMirrorFilterId pulumi.StringInput
}

func (TrafficMirrorFilterRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMirrorFilterRuleArgs)(nil)).Elem()
}

type TrafficMirrorFilterRuleInput interface {
	pulumi.Input

	ToTrafficMirrorFilterRuleOutput() TrafficMirrorFilterRuleOutput
	ToTrafficMirrorFilterRuleOutputWithContext(ctx context.Context) TrafficMirrorFilterRuleOutput
}

func (*TrafficMirrorFilterRule) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficMirrorFilterRule)(nil))
}

func (i *TrafficMirrorFilterRule) ToTrafficMirrorFilterRuleOutput() TrafficMirrorFilterRuleOutput {
	return i.ToTrafficMirrorFilterRuleOutputWithContext(context.Background())
}

func (i *TrafficMirrorFilterRule) ToTrafficMirrorFilterRuleOutputWithContext(ctx context.Context) TrafficMirrorFilterRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficMirrorFilterRuleOutput)
}

type TrafficMirrorFilterRuleOutput struct{ *pulumi.OutputState }

func (TrafficMirrorFilterRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficMirrorFilterRule)(nil))
}

func (o TrafficMirrorFilterRuleOutput) ToTrafficMirrorFilterRuleOutput() TrafficMirrorFilterRuleOutput {
	return o
}

func (o TrafficMirrorFilterRuleOutput) ToTrafficMirrorFilterRuleOutputWithContext(ctx context.Context) TrafficMirrorFilterRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TrafficMirrorFilterRuleOutput{})
}