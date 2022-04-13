// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ask

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for Alexa::ASK::Skill
//
// Deprecated: Skill is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Skill struct {
	pulumi.CustomResourceState

	AuthenticationConfiguration SkillAuthenticationConfigurationOutput `pulumi:"authenticationConfiguration"`
	SkillPackage                SkillPackageOutput                     `pulumi:"skillPackage"`
	VendorId                    pulumi.StringOutput                    `pulumi:"vendorId"`
}

// NewSkill registers a new resource with the given unique name, arguments, and options.
func NewSkill(ctx *pulumi.Context,
	name string, args *SkillArgs, opts ...pulumi.ResourceOption) (*Skill, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthenticationConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'AuthenticationConfiguration'")
	}
	if args.SkillPackage == nil {
		return nil, errors.New("invalid value for required argument 'SkillPackage'")
	}
	if args.VendorId == nil {
		return nil, errors.New("invalid value for required argument 'VendorId'")
	}
	var resource Skill
	err := ctx.RegisterResource("aws-native:ask:Skill", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSkill gets an existing Skill resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSkill(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SkillState, opts ...pulumi.ResourceOption) (*Skill, error) {
	var resource Skill
	err := ctx.ReadResource("aws-native:ask:Skill", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Skill resources.
type skillState struct {
}

type SkillState struct {
}

func (SkillState) ElementType() reflect.Type {
	return reflect.TypeOf((*skillState)(nil)).Elem()
}

type skillArgs struct {
	AuthenticationConfiguration SkillAuthenticationConfiguration `pulumi:"authenticationConfiguration"`
	SkillPackage                SkillPackage                     `pulumi:"skillPackage"`
	VendorId                    string                           `pulumi:"vendorId"`
}

// The set of arguments for constructing a Skill resource.
type SkillArgs struct {
	AuthenticationConfiguration SkillAuthenticationConfigurationInput
	SkillPackage                SkillPackageInput
	VendorId                    pulumi.StringInput
}

func (SkillArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*skillArgs)(nil)).Elem()
}

type SkillInput interface {
	pulumi.Input

	ToSkillOutput() SkillOutput
	ToSkillOutputWithContext(ctx context.Context) SkillOutput
}

func (*Skill) ElementType() reflect.Type {
	return reflect.TypeOf((**Skill)(nil)).Elem()
}

func (i *Skill) ToSkillOutput() SkillOutput {
	return i.ToSkillOutputWithContext(context.Background())
}

func (i *Skill) ToSkillOutputWithContext(ctx context.Context) SkillOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkillOutput)
}

type SkillOutput struct{ *pulumi.OutputState }

func (SkillOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Skill)(nil)).Elem()
}

func (o SkillOutput) ToSkillOutput() SkillOutput {
	return o
}

func (o SkillOutput) ToSkillOutputWithContext(ctx context.Context) SkillOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SkillInput)(nil)).Elem(), &Skill{})
	pulumi.RegisterOutputType(SkillOutput{})
}
