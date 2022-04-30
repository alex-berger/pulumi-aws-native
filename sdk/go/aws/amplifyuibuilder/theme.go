// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package amplifyuibuilder

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::AmplifyUIBuilder::Theme Resource Type
type Theme struct {
	pulumi.CustomResourceState

	AppId           pulumi.StringOutput    `pulumi:"appId"`
	CreatedAt       pulumi.StringOutput    `pulumi:"createdAt"`
	EnvironmentName pulumi.StringOutput    `pulumi:"environmentName"`
	ModifiedAt      pulumi.StringOutput    `pulumi:"modifiedAt"`
	Name            pulumi.StringOutput    `pulumi:"name"`
	Overrides       ThemeValuesArrayOutput `pulumi:"overrides"`
	Tags            ThemeTagsPtrOutput     `pulumi:"tags"`
	Values          ThemeValuesArrayOutput `pulumi:"values"`
}

// NewTheme registers a new resource with the given unique name, arguments, and options.
func NewTheme(ctx *pulumi.Context,
	name string, args *ThemeArgs, opts ...pulumi.ResourceOption) (*Theme, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Values == nil {
		return nil, errors.New("invalid value for required argument 'Values'")
	}
	var resource Theme
	err := ctx.RegisterResource("aws-native:amplifyuibuilder:Theme", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTheme gets an existing Theme resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTheme(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ThemeState, opts ...pulumi.ResourceOption) (*Theme, error) {
	var resource Theme
	err := ctx.ReadResource("aws-native:amplifyuibuilder:Theme", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Theme resources.
type themeState struct {
}

type ThemeState struct {
}

func (ThemeState) ElementType() reflect.Type {
	return reflect.TypeOf((*themeState)(nil)).Elem()
}

type themeArgs struct {
	Name      *string       `pulumi:"name"`
	Overrides []ThemeValues `pulumi:"overrides"`
	Tags      *ThemeTags    `pulumi:"tags"`
	Values    []ThemeValues `pulumi:"values"`
}

// The set of arguments for constructing a Theme resource.
type ThemeArgs struct {
	Name      pulumi.StringPtrInput
	Overrides ThemeValuesArrayInput
	Tags      ThemeTagsPtrInput
	Values    ThemeValuesArrayInput
}

func (ThemeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*themeArgs)(nil)).Elem()
}

type ThemeInput interface {
	pulumi.Input

	ToThemeOutput() ThemeOutput
	ToThemeOutputWithContext(ctx context.Context) ThemeOutput
}

func (*Theme) ElementType() reflect.Type {
	return reflect.TypeOf((**Theme)(nil)).Elem()
}

func (i *Theme) ToThemeOutput() ThemeOutput {
	return i.ToThemeOutputWithContext(context.Background())
}

func (i *Theme) ToThemeOutputWithContext(ctx context.Context) ThemeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThemeOutput)
}

type ThemeOutput struct{ *pulumi.OutputState }

func (ThemeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Theme)(nil)).Elem()
}

func (o ThemeOutput) ToThemeOutput() ThemeOutput {
	return o
}

func (o ThemeOutput) ToThemeOutputWithContext(ctx context.Context) ThemeOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ThemeInput)(nil)).Elem(), &Theme{})
	pulumi.RegisterOutputType(ThemeOutput{})
}
