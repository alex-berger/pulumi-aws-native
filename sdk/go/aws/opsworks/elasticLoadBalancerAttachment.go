// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::OpsWorks::ElasticLoadBalancerAttachment
//
// Deprecated: ElasticLoadBalancerAttachment is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ElasticLoadBalancerAttachment struct {
	pulumi.CustomResourceState

	ElasticLoadBalancerName pulumi.StringOutput `pulumi:"elasticLoadBalancerName"`
	LayerId                 pulumi.StringOutput `pulumi:"layerId"`
}

// NewElasticLoadBalancerAttachment registers a new resource with the given unique name, arguments, and options.
func NewElasticLoadBalancerAttachment(ctx *pulumi.Context,
	name string, args *ElasticLoadBalancerAttachmentArgs, opts ...pulumi.ResourceOption) (*ElasticLoadBalancerAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ElasticLoadBalancerName == nil {
		return nil, errors.New("invalid value for required argument 'ElasticLoadBalancerName'")
	}
	if args.LayerId == nil {
		return nil, errors.New("invalid value for required argument 'LayerId'")
	}
	var resource ElasticLoadBalancerAttachment
	err := ctx.RegisterResource("aws-native:opsworks:ElasticLoadBalancerAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetElasticLoadBalancerAttachment gets an existing ElasticLoadBalancerAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetElasticLoadBalancerAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ElasticLoadBalancerAttachmentState, opts ...pulumi.ResourceOption) (*ElasticLoadBalancerAttachment, error) {
	var resource ElasticLoadBalancerAttachment
	err := ctx.ReadResource("aws-native:opsworks:ElasticLoadBalancerAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ElasticLoadBalancerAttachment resources.
type elasticLoadBalancerAttachmentState struct {
}

type ElasticLoadBalancerAttachmentState struct {
}

func (ElasticLoadBalancerAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticLoadBalancerAttachmentState)(nil)).Elem()
}

type elasticLoadBalancerAttachmentArgs struct {
	ElasticLoadBalancerName string `pulumi:"elasticLoadBalancerName"`
	LayerId                 string `pulumi:"layerId"`
}

// The set of arguments for constructing a ElasticLoadBalancerAttachment resource.
type ElasticLoadBalancerAttachmentArgs struct {
	ElasticLoadBalancerName pulumi.StringInput
	LayerId                 pulumi.StringInput
}

func (ElasticLoadBalancerAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticLoadBalancerAttachmentArgs)(nil)).Elem()
}

type ElasticLoadBalancerAttachmentInput interface {
	pulumi.Input

	ToElasticLoadBalancerAttachmentOutput() ElasticLoadBalancerAttachmentOutput
	ToElasticLoadBalancerAttachmentOutputWithContext(ctx context.Context) ElasticLoadBalancerAttachmentOutput
}

func (*ElasticLoadBalancerAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticLoadBalancerAttachment)(nil)).Elem()
}

func (i *ElasticLoadBalancerAttachment) ToElasticLoadBalancerAttachmentOutput() ElasticLoadBalancerAttachmentOutput {
	return i.ToElasticLoadBalancerAttachmentOutputWithContext(context.Background())
}

func (i *ElasticLoadBalancerAttachment) ToElasticLoadBalancerAttachmentOutputWithContext(ctx context.Context) ElasticLoadBalancerAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticLoadBalancerAttachmentOutput)
}

type ElasticLoadBalancerAttachmentOutput struct{ *pulumi.OutputState }

func (ElasticLoadBalancerAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticLoadBalancerAttachment)(nil)).Elem()
}

func (o ElasticLoadBalancerAttachmentOutput) ToElasticLoadBalancerAttachmentOutput() ElasticLoadBalancerAttachmentOutput {
	return o
}

func (o ElasticLoadBalancerAttachmentOutput) ToElasticLoadBalancerAttachmentOutputWithContext(ctx context.Context) ElasticLoadBalancerAttachmentOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticLoadBalancerAttachmentInput)(nil)).Elem(), &ElasticLoadBalancerAttachment{})
	pulumi.RegisterOutputType(ElasticLoadBalancerAttachmentOutput{})
}
