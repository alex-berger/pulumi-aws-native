// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ivs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IVS::RecordingConfiguration
type RecordingConfiguration struct {
	pulumi.CustomResourceState

	// Recording Configuration ARN is automatically generated on creation and assigned as the unique identifier.
	Arn                      pulumi.StringOutput                                  `pulumi:"arn"`
	DestinationConfiguration RecordingConfigurationDestinationConfigurationOutput `pulumi:"destinationConfiguration"`
	// Recording Configuration Name.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Recording Configuration State.
	State RecordingConfigurationStateEnumOutput `pulumi:"state"`
	// A list of key-value pairs that contain metadata for the asset model.
	Tags                   RecordingConfigurationTagArrayOutput                  `pulumi:"tags"`
	ThumbnailConfiguration RecordingConfigurationThumbnailConfigurationPtrOutput `pulumi:"thumbnailConfiguration"`
}

// NewRecordingConfiguration registers a new resource with the given unique name, arguments, and options.
func NewRecordingConfiguration(ctx *pulumi.Context,
	name string, args *RecordingConfigurationArgs, opts ...pulumi.ResourceOption) (*RecordingConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'DestinationConfiguration'")
	}
	var resource RecordingConfiguration
	err := ctx.RegisterResource("aws-native:ivs:RecordingConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRecordingConfiguration gets an existing RecordingConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRecordingConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RecordingConfigurationState, opts ...pulumi.ResourceOption) (*RecordingConfiguration, error) {
	var resource RecordingConfiguration
	err := ctx.ReadResource("aws-native:ivs:RecordingConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RecordingConfiguration resources.
type recordingConfigurationState struct {
}

type RecordingConfigurationState struct {
}

func (RecordingConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*recordingConfigurationState)(nil)).Elem()
}

type recordingConfigurationArgs struct {
	DestinationConfiguration RecordingConfigurationDestinationConfiguration `pulumi:"destinationConfiguration"`
	// Recording Configuration Name.
	Name *string `pulumi:"name"`
	// A list of key-value pairs that contain metadata for the asset model.
	Tags                   []RecordingConfigurationTag                   `pulumi:"tags"`
	ThumbnailConfiguration *RecordingConfigurationThumbnailConfiguration `pulumi:"thumbnailConfiguration"`
}

// The set of arguments for constructing a RecordingConfiguration resource.
type RecordingConfigurationArgs struct {
	DestinationConfiguration RecordingConfigurationDestinationConfigurationInput
	// Recording Configuration Name.
	Name pulumi.StringPtrInput
	// A list of key-value pairs that contain metadata for the asset model.
	Tags                   RecordingConfigurationTagArrayInput
	ThumbnailConfiguration RecordingConfigurationThumbnailConfigurationPtrInput
}

func (RecordingConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*recordingConfigurationArgs)(nil)).Elem()
}

type RecordingConfigurationInput interface {
	pulumi.Input

	ToRecordingConfigurationOutput() RecordingConfigurationOutput
	ToRecordingConfigurationOutputWithContext(ctx context.Context) RecordingConfigurationOutput
}

func (*RecordingConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**RecordingConfiguration)(nil)).Elem()
}

func (i *RecordingConfiguration) ToRecordingConfigurationOutput() RecordingConfigurationOutput {
	return i.ToRecordingConfigurationOutputWithContext(context.Background())
}

func (i *RecordingConfiguration) ToRecordingConfigurationOutputWithContext(ctx context.Context) RecordingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordingConfigurationOutput)
}

type RecordingConfigurationOutput struct{ *pulumi.OutputState }

func (RecordingConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecordingConfiguration)(nil)).Elem()
}

func (o RecordingConfigurationOutput) ToRecordingConfigurationOutput() RecordingConfigurationOutput {
	return o
}

func (o RecordingConfigurationOutput) ToRecordingConfigurationOutputWithContext(ctx context.Context) RecordingConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RecordingConfigurationInput)(nil)).Elem(), &RecordingConfiguration{})
	pulumi.RegisterOutputType(RecordingConfigurationOutput{})
}
