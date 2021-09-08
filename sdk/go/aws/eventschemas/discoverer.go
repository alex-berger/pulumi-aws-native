// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventschemas

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-discoverer.html
type Discoverer struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-discoverer.html#cfn-eventschemas-discoverer-description
	Description   pulumi.StringPtrOutput `pulumi:"description"`
	DiscovererArn pulumi.StringOutput    `pulumi:"discovererArn"`
	DiscovererId  pulumi.StringOutput    `pulumi:"discovererId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-discoverer.html#cfn-eventschemas-discoverer-sourcearn
	SourceArn pulumi.StringOutput `pulumi:"sourceArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-discoverer.html#cfn-eventschemas-discoverer-tags
	Tags DiscovererTagsEntryArrayOutput `pulumi:"tags"`
}

// NewDiscoverer registers a new resource with the given unique name, arguments, and options.
func NewDiscoverer(ctx *pulumi.Context,
	name string, args *DiscovererArgs, opts ...pulumi.ResourceOption) (*Discoverer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SourceArn == nil {
		return nil, errors.New("invalid value for required argument 'SourceArn'")
	}
	var resource Discoverer
	err := ctx.RegisterResource("aws-native:eventschemas:Discoverer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiscoverer gets an existing Discoverer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiscoverer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiscovererState, opts ...pulumi.ResourceOption) (*Discoverer, error) {
	var resource Discoverer
	err := ctx.ReadResource("aws-native:eventschemas:Discoverer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Discoverer resources.
type discovererState struct {
}

type DiscovererState struct {
}

func (DiscovererState) ElementType() reflect.Type {
	return reflect.TypeOf((*discovererState)(nil)).Elem()
}

type discovererArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-discoverer.html#cfn-eventschemas-discoverer-description
	Description *string `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-discoverer.html#cfn-eventschemas-discoverer-sourcearn
	SourceArn string `pulumi:"sourceArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-discoverer.html#cfn-eventschemas-discoverer-tags
	Tags []DiscovererTagsEntry `pulumi:"tags"`
}

// The set of arguments for constructing a Discoverer resource.
type DiscovererArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-discoverer.html#cfn-eventschemas-discoverer-description
	Description pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-discoverer.html#cfn-eventschemas-discoverer-sourcearn
	SourceArn pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-discoverer.html#cfn-eventschemas-discoverer-tags
	Tags DiscovererTagsEntryArrayInput
}

func (DiscovererArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*discovererArgs)(nil)).Elem()
}

type DiscovererInput interface {
	pulumi.Input

	ToDiscovererOutput() DiscovererOutput
	ToDiscovererOutputWithContext(ctx context.Context) DiscovererOutput
}

func (*Discoverer) ElementType() reflect.Type {
	return reflect.TypeOf((*Discoverer)(nil))
}

func (i *Discoverer) ToDiscovererOutput() DiscovererOutput {
	return i.ToDiscovererOutputWithContext(context.Background())
}

func (i *Discoverer) ToDiscovererOutputWithContext(ctx context.Context) DiscovererOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiscovererOutput)
}

type DiscovererOutput struct{ *pulumi.OutputState }

func (DiscovererOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Discoverer)(nil))
}

func (o DiscovererOutput) ToDiscovererOutput() DiscovererOutput {
	return o
}

func (o DiscovererOutput) ToDiscovererOutputWithContext(ctx context.Context) DiscovererOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DiscovererOutput{})
}