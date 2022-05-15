// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Route53::RecordSet
//
// Deprecated: RecordSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type RecordSet struct {
	pulumi.CustomResourceState

	AliasTarget      RecordSetAliasTargetPtrOutput `pulumi:"aliasTarget"`
	Comment          pulumi.StringPtrOutput        `pulumi:"comment"`
	Failover         pulumi.StringPtrOutput        `pulumi:"failover"`
	GeoLocation      RecordSetGeoLocationPtrOutput `pulumi:"geoLocation"`
	HealthCheckId    pulumi.StringPtrOutput        `pulumi:"healthCheckId"`
	HostedZoneId     pulumi.StringPtrOutput        `pulumi:"hostedZoneId"`
	HostedZoneName   pulumi.StringPtrOutput        `pulumi:"hostedZoneName"`
	MultiValueAnswer pulumi.BoolPtrOutput          `pulumi:"multiValueAnswer"`
	Name             pulumi.StringOutput           `pulumi:"name"`
	Region           pulumi.StringPtrOutput        `pulumi:"region"`
	ResourceRecords  pulumi.StringArrayOutput      `pulumi:"resourceRecords"`
	SetIdentifier    pulumi.StringPtrOutput        `pulumi:"setIdentifier"`
	TTL              pulumi.StringPtrOutput        `pulumi:"tTL"`
	Type             pulumi.StringOutput           `pulumi:"type"`
	Weight           pulumi.IntPtrOutput           `pulumi:"weight"`
}

// NewRecordSet registers a new resource with the given unique name, arguments, and options.
func NewRecordSet(ctx *pulumi.Context,
	name string, args *RecordSetArgs, opts ...pulumi.ResourceOption) (*RecordSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource RecordSet
	err := ctx.RegisterResource("aws-native:route53:RecordSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRecordSet gets an existing RecordSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRecordSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RecordSetState, opts ...pulumi.ResourceOption) (*RecordSet, error) {
	var resource RecordSet
	err := ctx.ReadResource("aws-native:route53:RecordSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RecordSet resources.
type recordSetState struct {
}

type RecordSetState struct {
}

func (RecordSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*recordSetState)(nil)).Elem()
}

type recordSetArgs struct {
	AliasTarget      *RecordSetAliasTarget `pulumi:"aliasTarget"`
	Comment          *string               `pulumi:"comment"`
	Failover         *string               `pulumi:"failover"`
	GeoLocation      *RecordSetGeoLocation `pulumi:"geoLocation"`
	HealthCheckId    *string               `pulumi:"healthCheckId"`
	HostedZoneId     *string               `pulumi:"hostedZoneId"`
	HostedZoneName   *string               `pulumi:"hostedZoneName"`
	MultiValueAnswer *bool                 `pulumi:"multiValueAnswer"`
	Name             *string               `pulumi:"name"`
	Region           *string               `pulumi:"region"`
	ResourceRecords  []string              `pulumi:"resourceRecords"`
	SetIdentifier    *string               `pulumi:"setIdentifier"`
	TTL              *string               `pulumi:"tTL"`
	Type             string                `pulumi:"type"`
	Weight           *int                  `pulumi:"weight"`
}

// The set of arguments for constructing a RecordSet resource.
type RecordSetArgs struct {
	AliasTarget      RecordSetAliasTargetPtrInput
	Comment          pulumi.StringPtrInput
	Failover         pulumi.StringPtrInput
	GeoLocation      RecordSetGeoLocationPtrInput
	HealthCheckId    pulumi.StringPtrInput
	HostedZoneId     pulumi.StringPtrInput
	HostedZoneName   pulumi.StringPtrInput
	MultiValueAnswer pulumi.BoolPtrInput
	Name             pulumi.StringPtrInput
	Region           pulumi.StringPtrInput
	ResourceRecords  pulumi.StringArrayInput
	SetIdentifier    pulumi.StringPtrInput
	TTL              pulumi.StringPtrInput
	Type             pulumi.StringInput
	Weight           pulumi.IntPtrInput
}

func (RecordSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*recordSetArgs)(nil)).Elem()
}

type RecordSetInput interface {
	pulumi.Input

	ToRecordSetOutput() RecordSetOutput
	ToRecordSetOutputWithContext(ctx context.Context) RecordSetOutput
}

func (*RecordSet) ElementType() reflect.Type {
	return reflect.TypeOf((**RecordSet)(nil)).Elem()
}

func (i *RecordSet) ToRecordSetOutput() RecordSetOutput {
	return i.ToRecordSetOutputWithContext(context.Background())
}

func (i *RecordSet) ToRecordSetOutputWithContext(ctx context.Context) RecordSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordSetOutput)
}

type RecordSetOutput struct{ *pulumi.OutputState }

func (RecordSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecordSet)(nil)).Elem()
}

func (o RecordSetOutput) ToRecordSetOutput() RecordSetOutput {
	return o
}

func (o RecordSetOutput) ToRecordSetOutputWithContext(ctx context.Context) RecordSetOutput {
	return o
}

func (o RecordSetOutput) AliasTarget() RecordSetAliasTargetPtrOutput {
	return o.ApplyT(func(v *RecordSet) RecordSetAliasTargetPtrOutput { return v.AliasTarget }).(RecordSetAliasTargetPtrOutput)
}

func (o RecordSetOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o RecordSetOutput) Failover() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.StringPtrOutput { return v.Failover }).(pulumi.StringPtrOutput)
}

func (o RecordSetOutput) GeoLocation() RecordSetGeoLocationPtrOutput {
	return o.ApplyT(func(v *RecordSet) RecordSetGeoLocationPtrOutput { return v.GeoLocation }).(RecordSetGeoLocationPtrOutput)
}

func (o RecordSetOutput) HealthCheckId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.StringPtrOutput { return v.HealthCheckId }).(pulumi.StringPtrOutput)
}

func (o RecordSetOutput) HostedZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.StringPtrOutput { return v.HostedZoneId }).(pulumi.StringPtrOutput)
}

func (o RecordSetOutput) HostedZoneName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.StringPtrOutput { return v.HostedZoneName }).(pulumi.StringPtrOutput)
}

func (o RecordSetOutput) MultiValueAnswer() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.BoolPtrOutput { return v.MultiValueAnswer }).(pulumi.BoolPtrOutput)
}

func (o RecordSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RecordSetOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

func (o RecordSetOutput) ResourceRecords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.StringArrayOutput { return v.ResourceRecords }).(pulumi.StringArrayOutput)
}

func (o RecordSetOutput) SetIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.StringPtrOutput { return v.SetIdentifier }).(pulumi.StringPtrOutput)
}

func (o RecordSetOutput) TTL() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.StringPtrOutput { return v.TTL }).(pulumi.StringPtrOutput)
}

func (o RecordSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o RecordSetOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.IntPtrOutput { return v.Weight }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RecordSetInput)(nil)).Elem(), &RecordSet{})
	pulumi.RegisterOutputType(RecordSetOutput{})
}
