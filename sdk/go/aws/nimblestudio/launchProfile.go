// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package nimblestudio

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a launch profile which delegates access to a collection of studio components to studio users
type LaunchProfile struct {
	pulumi.CustomResourceState

	// <p>The description.</p>
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// <p>Specifies the IDs of the EC2 subnets where streaming sessions will be accessible from.
	//             These subnets must support the specified instance types. </p>
	Ec2SubnetIds    pulumi.StringArrayOutput `pulumi:"ec2SubnetIds"`
	LaunchProfileId pulumi.StringOutput      `pulumi:"launchProfileId"`
	// <p>The version number of the protocol that is used by the launch profile. The only valid
	//             version is "2021-03-31".</p>
	LaunchProfileProtocolVersions pulumi.StringArrayOutput `pulumi:"launchProfileProtocolVersions"`
	// <p>The name for the launch profile.</p>
	Name                pulumi.StringOutput                    `pulumi:"name"`
	StreamConfiguration LaunchProfileStreamConfigurationOutput `pulumi:"streamConfiguration"`
	// <p>Unique identifiers for a collection of studio components that can be used with this
	//             launch profile.</p>
	StudioComponentIds pulumi.StringArrayOutput `pulumi:"studioComponentIds"`
	// <p>The studio ID. </p>
	StudioId pulumi.StringOutput        `pulumi:"studioId"`
	Tags     LaunchProfileTagsPtrOutput `pulumi:"tags"`
}

// NewLaunchProfile registers a new resource with the given unique name, arguments, and options.
func NewLaunchProfile(ctx *pulumi.Context,
	name string, args *LaunchProfileArgs, opts ...pulumi.ResourceOption) (*LaunchProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ec2SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'Ec2SubnetIds'")
	}
	if args.LaunchProfileProtocolVersions == nil {
		return nil, errors.New("invalid value for required argument 'LaunchProfileProtocolVersions'")
	}
	if args.StreamConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'StreamConfiguration'")
	}
	if args.StudioComponentIds == nil {
		return nil, errors.New("invalid value for required argument 'StudioComponentIds'")
	}
	if args.StudioId == nil {
		return nil, errors.New("invalid value for required argument 'StudioId'")
	}
	var resource LaunchProfile
	err := ctx.RegisterResource("aws-native:nimblestudio:LaunchProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLaunchProfile gets an existing LaunchProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLaunchProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LaunchProfileState, opts ...pulumi.ResourceOption) (*LaunchProfile, error) {
	var resource LaunchProfile
	err := ctx.ReadResource("aws-native:nimblestudio:LaunchProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LaunchProfile resources.
type launchProfileState struct {
}

type LaunchProfileState struct {
}

func (LaunchProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*launchProfileState)(nil)).Elem()
}

type launchProfileArgs struct {
	// <p>The description.</p>
	Description *string `pulumi:"description"`
	// <p>Specifies the IDs of the EC2 subnets where streaming sessions will be accessible from.
	//             These subnets must support the specified instance types. </p>
	Ec2SubnetIds []string `pulumi:"ec2SubnetIds"`
	// <p>The version number of the protocol that is used by the launch profile. The only valid
	//             version is "2021-03-31".</p>
	LaunchProfileProtocolVersions []string `pulumi:"launchProfileProtocolVersions"`
	// <p>The name for the launch profile.</p>
	Name                *string                          `pulumi:"name"`
	StreamConfiguration LaunchProfileStreamConfiguration `pulumi:"streamConfiguration"`
	// <p>Unique identifiers for a collection of studio components that can be used with this
	//             launch profile.</p>
	StudioComponentIds []string `pulumi:"studioComponentIds"`
	// <p>The studio ID. </p>
	StudioId string             `pulumi:"studioId"`
	Tags     *LaunchProfileTags `pulumi:"tags"`
}

// The set of arguments for constructing a LaunchProfile resource.
type LaunchProfileArgs struct {
	// <p>The description.</p>
	Description pulumi.StringPtrInput
	// <p>Specifies the IDs of the EC2 subnets where streaming sessions will be accessible from.
	//             These subnets must support the specified instance types. </p>
	Ec2SubnetIds pulumi.StringArrayInput
	// <p>The version number of the protocol that is used by the launch profile. The only valid
	//             version is "2021-03-31".</p>
	LaunchProfileProtocolVersions pulumi.StringArrayInput
	// <p>The name for the launch profile.</p>
	Name                pulumi.StringPtrInput
	StreamConfiguration LaunchProfileStreamConfigurationInput
	// <p>Unique identifiers for a collection of studio components that can be used with this
	//             launch profile.</p>
	StudioComponentIds pulumi.StringArrayInput
	// <p>The studio ID. </p>
	StudioId pulumi.StringInput
	Tags     LaunchProfileTagsPtrInput
}

func (LaunchProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*launchProfileArgs)(nil)).Elem()
}

type LaunchProfileInput interface {
	pulumi.Input

	ToLaunchProfileOutput() LaunchProfileOutput
	ToLaunchProfileOutputWithContext(ctx context.Context) LaunchProfileOutput
}

func (*LaunchProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**LaunchProfile)(nil)).Elem()
}

func (i *LaunchProfile) ToLaunchProfileOutput() LaunchProfileOutput {
	return i.ToLaunchProfileOutputWithContext(context.Background())
}

func (i *LaunchProfile) ToLaunchProfileOutputWithContext(ctx context.Context) LaunchProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchProfileOutput)
}

type LaunchProfileOutput struct{ *pulumi.OutputState }

func (LaunchProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LaunchProfile)(nil)).Elem()
}

func (o LaunchProfileOutput) ToLaunchProfileOutput() LaunchProfileOutput {
	return o
}

func (o LaunchProfileOutput) ToLaunchProfileOutputWithContext(ctx context.Context) LaunchProfileOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LaunchProfileInput)(nil)).Elem(), &LaunchProfile{})
	pulumi.RegisterOutputType(LaunchProfileOutput{})
}
