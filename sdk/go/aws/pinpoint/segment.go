// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-segment.html
type Segment struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-segment.html#cfn-pinpoint-segment-applicationid
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	Arn           pulumi.StringOutput `pulumi:"arn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-segment.html#cfn-pinpoint-segment-dimensions
	Dimensions SegmentSegmentDimensionsPtrOutput `pulumi:"dimensions"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-segment.html#cfn-pinpoint-segment-name
	Name pulumi.StringOutput `pulumi:"name"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-segment.html#cfn-pinpoint-segment-segmentgroups
	SegmentGroups SegmentSegmentGroupsPtrOutput `pulumi:"segmentGroups"`
	SegmentId     pulumi.StringOutput           `pulumi:"segmentId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-segment.html#cfn-pinpoint-segment-tags
	Tags pulumi.AnyOutput `pulumi:"tags"`
}

// NewSegment registers a new resource with the given unique name, arguments, and options.
func NewSegment(ctx *pulumi.Context,
	name string, args *SegmentArgs, opts ...pulumi.ResourceOption) (*Segment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	var resource Segment
	err := ctx.RegisterResource("aws-native:pinpoint:Segment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSegment gets an existing Segment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSegment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SegmentState, opts ...pulumi.ResourceOption) (*Segment, error) {
	var resource Segment
	err := ctx.ReadResource("aws-native:pinpoint:Segment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Segment resources.
type segmentState struct {
}

type SegmentState struct {
}

func (SegmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*segmentState)(nil)).Elem()
}

type segmentArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-segment.html#cfn-pinpoint-segment-applicationid
	ApplicationId string `pulumi:"applicationId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-segment.html#cfn-pinpoint-segment-dimensions
	Dimensions *SegmentSegmentDimensions `pulumi:"dimensions"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-segment.html#cfn-pinpoint-segment-name
	Name string `pulumi:"name"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-segment.html#cfn-pinpoint-segment-segmentgroups
	SegmentGroups *SegmentSegmentGroups `pulumi:"segmentGroups"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-segment.html#cfn-pinpoint-segment-tags
	Tags interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Segment resource.
type SegmentArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-segment.html#cfn-pinpoint-segment-applicationid
	ApplicationId pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-segment.html#cfn-pinpoint-segment-dimensions
	Dimensions SegmentSegmentDimensionsPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-segment.html#cfn-pinpoint-segment-name
	Name pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-segment.html#cfn-pinpoint-segment-segmentgroups
	SegmentGroups SegmentSegmentGroupsPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-segment.html#cfn-pinpoint-segment-tags
	Tags pulumi.Input
}

func (SegmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*segmentArgs)(nil)).Elem()
}

type SegmentInput interface {
	pulumi.Input

	ToSegmentOutput() SegmentOutput
	ToSegmentOutputWithContext(ctx context.Context) SegmentOutput
}

func (*Segment) ElementType() reflect.Type {
	return reflect.TypeOf((*Segment)(nil))
}

func (i *Segment) ToSegmentOutput() SegmentOutput {
	return i.ToSegmentOutputWithContext(context.Background())
}

func (i *Segment) ToSegmentOutputWithContext(ctx context.Context) SegmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SegmentOutput)
}

type SegmentOutput struct{ *pulumi.OutputState }

func (SegmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Segment)(nil))
}

func (o SegmentOutput) ToSegmentOutput() SegmentOutput {
	return o
}

func (o SegmentOutput) ToSegmentOutputWithContext(ctx context.Context) SegmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SegmentOutput{})
}