// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS::S3::MultiRegionAccessPoint is an Amazon S3 resource type that dynamically routes S3 requests to easily satisfy geographic compliance requirements based on customer-defined routing policies.
type MultiRegionAccessPoint struct {
	pulumi.CustomResourceState

	// The alias is a unique identifier to, and is part of the public DNS name for this Multi Region Access Point
	Alias pulumi.StringOutput `pulumi:"alias"`
	// The timestamp of the when the Multi Region Access Point is created
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The name you want to assign to this Multi Region Access Point.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The PublicAccessBlock configuration that you want to apply to this Multi Region Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.
	PublicAccessBlockConfiguration MultiRegionAccessPointPublicAccessBlockConfigurationPtrOutput `pulumi:"publicAccessBlockConfiguration"`
	// The list of buckets that you want to associate this Multi Region Access Point with.
	Regions MultiRegionAccessPointRegionArrayOutput `pulumi:"regions"`
}

// NewMultiRegionAccessPoint registers a new resource with the given unique name, arguments, and options.
func NewMultiRegionAccessPoint(ctx *pulumi.Context,
	name string, args *MultiRegionAccessPointArgs, opts ...pulumi.ResourceOption) (*MultiRegionAccessPoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Regions == nil {
		return nil, errors.New("invalid value for required argument 'Regions'")
	}
	var resource MultiRegionAccessPoint
	err := ctx.RegisterResource("aws-native:s3:MultiRegionAccessPoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMultiRegionAccessPoint gets an existing MultiRegionAccessPoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMultiRegionAccessPoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MultiRegionAccessPointState, opts ...pulumi.ResourceOption) (*MultiRegionAccessPoint, error) {
	var resource MultiRegionAccessPoint
	err := ctx.ReadResource("aws-native:s3:MultiRegionAccessPoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MultiRegionAccessPoint resources.
type multiRegionAccessPointState struct {
}

type MultiRegionAccessPointState struct {
}

func (MultiRegionAccessPointState) ElementType() reflect.Type {
	return reflect.TypeOf((*multiRegionAccessPointState)(nil)).Elem()
}

type multiRegionAccessPointArgs struct {
	// The name you want to assign to this Multi Region Access Point.
	Name *string `pulumi:"name"`
	// The PublicAccessBlock configuration that you want to apply to this Multi Region Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.
	PublicAccessBlockConfiguration *MultiRegionAccessPointPublicAccessBlockConfiguration `pulumi:"publicAccessBlockConfiguration"`
	// The list of buckets that you want to associate this Multi Region Access Point with.
	Regions []MultiRegionAccessPointRegion `pulumi:"regions"`
}

// The set of arguments for constructing a MultiRegionAccessPoint resource.
type MultiRegionAccessPointArgs struct {
	// The name you want to assign to this Multi Region Access Point.
	Name pulumi.StringPtrInput
	// The PublicAccessBlock configuration that you want to apply to this Multi Region Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.
	PublicAccessBlockConfiguration MultiRegionAccessPointPublicAccessBlockConfigurationPtrInput
	// The list of buckets that you want to associate this Multi Region Access Point with.
	Regions MultiRegionAccessPointRegionArrayInput
}

func (MultiRegionAccessPointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*multiRegionAccessPointArgs)(nil)).Elem()
}

type MultiRegionAccessPointInput interface {
	pulumi.Input

	ToMultiRegionAccessPointOutput() MultiRegionAccessPointOutput
	ToMultiRegionAccessPointOutputWithContext(ctx context.Context) MultiRegionAccessPointOutput
}

func (*MultiRegionAccessPoint) ElementType() reflect.Type {
	return reflect.TypeOf((**MultiRegionAccessPoint)(nil)).Elem()
}

func (i *MultiRegionAccessPoint) ToMultiRegionAccessPointOutput() MultiRegionAccessPointOutput {
	return i.ToMultiRegionAccessPointOutputWithContext(context.Background())
}

func (i *MultiRegionAccessPoint) ToMultiRegionAccessPointOutputWithContext(ctx context.Context) MultiRegionAccessPointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultiRegionAccessPointOutput)
}

type MultiRegionAccessPointOutput struct{ *pulumi.OutputState }

func (MultiRegionAccessPointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MultiRegionAccessPoint)(nil)).Elem()
}

func (o MultiRegionAccessPointOutput) ToMultiRegionAccessPointOutput() MultiRegionAccessPointOutput {
	return o
}

func (o MultiRegionAccessPointOutput) ToMultiRegionAccessPointOutputWithContext(ctx context.Context) MultiRegionAccessPointOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MultiRegionAccessPointInput)(nil)).Elem(), &MultiRegionAccessPoint{})
	pulumi.RegisterOutputType(MultiRegionAccessPointOutput{})
}
