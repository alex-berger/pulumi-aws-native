// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html
type CapacityReservation struct {
	pulumi.CustomResourceState

	AvailabilityZone       pulumi.StringOutput `pulumi:"availabilityZone"`
	AvailableInstanceCount pulumi.IntOutput    `pulumi:"availableInstanceCount"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-ebsoptimized
	EbsOptimized pulumi.BoolPtrOutput `pulumi:"ebsOptimized"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-enddate
	EndDate pulumi.StringPtrOutput `pulumi:"endDate"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-enddatetype
	EndDateType pulumi.StringPtrOutput `pulumi:"endDateType"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-ephemeralstorage
	EphemeralStorage pulumi.BoolPtrOutput `pulumi:"ephemeralStorage"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-instancecount
	InstanceCount pulumi.IntOutput `pulumi:"instanceCount"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-instancematchcriteria
	InstanceMatchCriteria pulumi.StringPtrOutput `pulumi:"instanceMatchCriteria"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-instanceplatform
	InstancePlatform pulumi.StringOutput `pulumi:"instancePlatform"`
	InstanceType     pulumi.StringOutput `pulumi:"instanceType"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-tagspecifications
	TagSpecifications  CapacityReservationTagSpecificationArrayOutput `pulumi:"tagSpecifications"`
	Tenancy            pulumi.StringOutput                            `pulumi:"tenancy"`
	TotalInstanceCount pulumi.IntOutput                               `pulumi:"totalInstanceCount"`
}

// NewCapacityReservation registers a new resource with the given unique name, arguments, and options.
func NewCapacityReservation(ctx *pulumi.Context,
	name string, args *CapacityReservationArgs, opts ...pulumi.ResourceOption) (*CapacityReservation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZone == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZone'")
	}
	if args.InstanceCount == nil {
		return nil, errors.New("invalid value for required argument 'InstanceCount'")
	}
	if args.InstancePlatform == nil {
		return nil, errors.New("invalid value for required argument 'InstancePlatform'")
	}
	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	var resource CapacityReservation
	err := ctx.RegisterResource("aws-native:ec2:CapacityReservation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCapacityReservation gets an existing CapacityReservation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCapacityReservation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CapacityReservationState, opts ...pulumi.ResourceOption) (*CapacityReservation, error) {
	var resource CapacityReservation
	err := ctx.ReadResource("aws-native:ec2:CapacityReservation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CapacityReservation resources.
type capacityReservationState struct {
}

type CapacityReservationState struct {
}

func (CapacityReservationState) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityReservationState)(nil)).Elem()
}

type capacityReservationArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-availabilityzone
	AvailabilityZone string `pulumi:"availabilityZone"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-ebsoptimized
	EbsOptimized *bool `pulumi:"ebsOptimized"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-enddate
	EndDate *string `pulumi:"endDate"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-enddatetype
	EndDateType *string `pulumi:"endDateType"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-ephemeralstorage
	EphemeralStorage *bool `pulumi:"ephemeralStorage"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-instancecount
	InstanceCount int `pulumi:"instanceCount"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-instancematchcriteria
	InstanceMatchCriteria *string `pulumi:"instanceMatchCriteria"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-instanceplatform
	InstancePlatform string `pulumi:"instancePlatform"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-instancetype
	InstanceType string `pulumi:"instanceType"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-tagspecifications
	TagSpecifications []CapacityReservationTagSpecification `pulumi:"tagSpecifications"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-tenancy
	Tenancy *string `pulumi:"tenancy"`
}

// The set of arguments for constructing a CapacityReservation resource.
type CapacityReservationArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-availabilityzone
	AvailabilityZone pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-ebsoptimized
	EbsOptimized pulumi.BoolPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-enddate
	EndDate pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-enddatetype
	EndDateType pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-ephemeralstorage
	EphemeralStorage pulumi.BoolPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-instancecount
	InstanceCount pulumi.IntInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-instancematchcriteria
	InstanceMatchCriteria pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-instanceplatform
	InstancePlatform pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-instancetype
	InstanceType pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-tagspecifications
	TagSpecifications CapacityReservationTagSpecificationArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-tenancy
	Tenancy pulumi.StringPtrInput
}

func (CapacityReservationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityReservationArgs)(nil)).Elem()
}

type CapacityReservationInput interface {
	pulumi.Input

	ToCapacityReservationOutput() CapacityReservationOutput
	ToCapacityReservationOutputWithContext(ctx context.Context) CapacityReservationOutput
}

func (*CapacityReservation) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacityReservation)(nil))
}

func (i *CapacityReservation) ToCapacityReservationOutput() CapacityReservationOutput {
	return i.ToCapacityReservationOutputWithContext(context.Background())
}

func (i *CapacityReservation) ToCapacityReservationOutputWithContext(ctx context.Context) CapacityReservationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityReservationOutput)
}

type CapacityReservationOutput struct{ *pulumi.OutputState }

func (CapacityReservationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacityReservation)(nil))
}

func (o CapacityReservationOutput) ToCapacityReservationOutput() CapacityReservationOutput {
	return o
}

func (o CapacityReservationOutput) ToCapacityReservationOutputWithContext(ctx context.Context) CapacityReservationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CapacityReservationOutput{})
}