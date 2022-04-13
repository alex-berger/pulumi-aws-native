// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource used to control (enable/disable) DNSSEC in a specific hosted zone.
type DNSSEC struct {
	pulumi.CustomResourceState

	// The unique string (ID) used to identify a hosted zone.
	HostedZoneId pulumi.StringOutput `pulumi:"hostedZoneId"`
}

// NewDNSSEC registers a new resource with the given unique name, arguments, and options.
func NewDNSSEC(ctx *pulumi.Context,
	name string, args *DNSSECArgs, opts ...pulumi.ResourceOption) (*DNSSEC, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostedZoneId == nil {
		return nil, errors.New("invalid value for required argument 'HostedZoneId'")
	}
	var resource DNSSEC
	err := ctx.RegisterResource("aws-native:route53:DNSSEC", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDNSSEC gets an existing DNSSEC resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDNSSEC(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DNSSECState, opts ...pulumi.ResourceOption) (*DNSSEC, error) {
	var resource DNSSEC
	err := ctx.ReadResource("aws-native:route53:DNSSEC", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DNSSEC resources.
type dnssecState struct {
}

type DNSSECState struct {
}

func (DNSSECState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnssecState)(nil)).Elem()
}

type dnssecArgs struct {
	// The unique string (ID) used to identify a hosted zone.
	HostedZoneId string `pulumi:"hostedZoneId"`
}

// The set of arguments for constructing a DNSSEC resource.
type DNSSECArgs struct {
	// The unique string (ID) used to identify a hosted zone.
	HostedZoneId pulumi.StringInput
}

func (DNSSECArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnssecArgs)(nil)).Elem()
}

type DNSSECInput interface {
	pulumi.Input

	ToDNSSECOutput() DNSSECOutput
	ToDNSSECOutputWithContext(ctx context.Context) DNSSECOutput
}

func (*DNSSEC) ElementType() reflect.Type {
	return reflect.TypeOf((**DNSSEC)(nil)).Elem()
}

func (i *DNSSEC) ToDNSSECOutput() DNSSECOutput {
	return i.ToDNSSECOutputWithContext(context.Background())
}

func (i *DNSSEC) ToDNSSECOutputWithContext(ctx context.Context) DNSSECOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DNSSECOutput)
}

type DNSSECOutput struct{ *pulumi.OutputState }

func (DNSSECOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DNSSEC)(nil)).Elem()
}

func (o DNSSECOutput) ToDNSSECOutput() DNSSECOutput {
	return o
}

func (o DNSSECOutput) ToDNSSECOutputWithContext(ctx context.Context) DNSSECOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DNSSECInput)(nil)).Elem(), &DNSSEC{})
	pulumi.RegisterOutputType(DNSSECOutput{})
}
