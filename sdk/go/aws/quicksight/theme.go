// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the AWS::QuickSight::Theme Resource Type.
type Theme struct {
	pulumi.CustomResourceState

	// <p>The Amazon Resource Name (ARN) of the theme.</p>
	Arn          pulumi.StringOutput `pulumi:"arn"`
	AwsAccountId pulumi.StringOutput `pulumi:"awsAccountId"`
	// <p>The ID of the theme that a custom theme will inherit from. All themes inherit from one of
	// 			the starting themes defined by Amazon QuickSight. For a list of the starting themes, use
	// 				<code>ListThemes</code> or choose <b>Themes</b> from
	// 			within a QuickSight analysis. </p>
	BaseThemeId   pulumi.StringPtrOutput      `pulumi:"baseThemeId"`
	Configuration ThemeConfigurationPtrOutput `pulumi:"configuration"`
	// <p>The date and time that the theme was created.</p>
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// <p>The date and time that the theme was last updated.</p>
	LastUpdatedTime pulumi.StringOutput `pulumi:"lastUpdatedTime"`
	// <p>A display name for the theme.</p>
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// <p>A valid grouping of resource permissions to apply to the new theme.
	// 			</p>
	Permissions ThemeResourcePermissionArrayOutput `pulumi:"permissions"`
	// <p>A map of the key-value pairs for the resource tag or tags that you want to add to the
	// 			resource.</p>
	Tags    ThemeTagArrayOutput `pulumi:"tags"`
	ThemeId pulumi.StringOutput `pulumi:"themeId"`
	Type    ThemeTypeOutput     `pulumi:"type"`
	Version ThemeVersionOutput  `pulumi:"version"`
	// <p>A description of the first version of the theme that you're creating. Every time
	// 				<code>UpdateTheme</code> is called, a new version is created. Each version of the
	// 			theme has a description of the version in the <code>VersionDescription</code>
	// 			field.</p>
	VersionDescription pulumi.StringPtrOutput `pulumi:"versionDescription"`
}

// NewTheme registers a new resource with the given unique name, arguments, and options.
func NewTheme(ctx *pulumi.Context,
	name string, args *ThemeArgs, opts ...pulumi.ResourceOption) (*Theme, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AwsAccountId == nil {
		return nil, errors.New("invalid value for required argument 'AwsAccountId'")
	}
	if args.ThemeId == nil {
		return nil, errors.New("invalid value for required argument 'ThemeId'")
	}
	var resource Theme
	err := ctx.RegisterResource("aws-native:quicksight:Theme", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws-native:quicksight:Theme", name, id, state, &resource, opts...)
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
	AwsAccountId string `pulumi:"awsAccountId"`
	// <p>The ID of the theme that a custom theme will inherit from. All themes inherit from one of
	// 			the starting themes defined by Amazon QuickSight. For a list of the starting themes, use
	// 				<code>ListThemes</code> or choose <b>Themes</b> from
	// 			within a QuickSight analysis. </p>
	BaseThemeId   *string             `pulumi:"baseThemeId"`
	Configuration *ThemeConfiguration `pulumi:"configuration"`
	// <p>A display name for the theme.</p>
	Name *string `pulumi:"name"`
	// <p>A valid grouping of resource permissions to apply to the new theme.
	// 			</p>
	Permissions []ThemeResourcePermission `pulumi:"permissions"`
	// <p>A map of the key-value pairs for the resource tag or tags that you want to add to the
	// 			resource.</p>
	Tags    []ThemeTag `pulumi:"tags"`
	ThemeId string     `pulumi:"themeId"`
	// <p>A description of the first version of the theme that you're creating. Every time
	// 				<code>UpdateTheme</code> is called, a new version is created. Each version of the
	// 			theme has a description of the version in the <code>VersionDescription</code>
	// 			field.</p>
	VersionDescription *string `pulumi:"versionDescription"`
}

// The set of arguments for constructing a Theme resource.
type ThemeArgs struct {
	AwsAccountId pulumi.StringInput
	// <p>The ID of the theme that a custom theme will inherit from. All themes inherit from one of
	// 			the starting themes defined by Amazon QuickSight. For a list of the starting themes, use
	// 				<code>ListThemes</code> or choose <b>Themes</b> from
	// 			within a QuickSight analysis. </p>
	BaseThemeId   pulumi.StringPtrInput
	Configuration ThemeConfigurationPtrInput
	// <p>A display name for the theme.</p>
	Name pulumi.StringPtrInput
	// <p>A valid grouping of resource permissions to apply to the new theme.
	// 			</p>
	Permissions ThemeResourcePermissionArrayInput
	// <p>A map of the key-value pairs for the resource tag or tags that you want to add to the
	// 			resource.</p>
	Tags    ThemeTagArrayInput
	ThemeId pulumi.StringInput
	// <p>A description of the first version of the theme that you're creating. Every time
	// 				<code>UpdateTheme</code> is called, a new version is created. Each version of the
	// 			theme has a description of the version in the <code>VersionDescription</code>
	// 			field.</p>
	VersionDescription pulumi.StringPtrInput
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
