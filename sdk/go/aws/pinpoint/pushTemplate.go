// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html
type PushTemplate struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-adm
	ADM PushTemplateAndroidPushNotificationTemplatePtrOutput `pulumi:"aDM"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-apns
	APNS PushTemplateAPNSPushNotificationTemplatePtrOutput `pulumi:"aPNS"`
	Arn  pulumi.StringOutput                               `pulumi:"arn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-baidu
	Baidu PushTemplateAndroidPushNotificationTemplatePtrOutput `pulumi:"baidu"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-default
	Default PushTemplateDefaultPushNotificationTemplatePtrOutput `pulumi:"default"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-defaultsubstitutions
	DefaultSubstitutions pulumi.StringPtrOutput `pulumi:"defaultSubstitutions"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-gcm
	GCM PushTemplateAndroidPushNotificationTemplatePtrOutput `pulumi:"gCM"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-tags
	Tags pulumi.AnyOutput `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-templatedescription
	TemplateDescription pulumi.StringPtrOutput `pulumi:"templateDescription"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-templatename
	TemplateName pulumi.StringOutput `pulumi:"templateName"`
}

// NewPushTemplate registers a new resource with the given unique name, arguments, and options.
func NewPushTemplate(ctx *pulumi.Context,
	name string, args *PushTemplateArgs, opts ...pulumi.ResourceOption) (*PushTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TemplateName == nil {
		return nil, errors.New("invalid value for required argument 'TemplateName'")
	}
	var resource PushTemplate
	err := ctx.RegisterResource("aws-native:pinpoint:PushTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPushTemplate gets an existing PushTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPushTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PushTemplateState, opts ...pulumi.ResourceOption) (*PushTemplate, error) {
	var resource PushTemplate
	err := ctx.ReadResource("aws-native:pinpoint:PushTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PushTemplate resources.
type pushTemplateState struct {
}

type PushTemplateState struct {
}

func (PushTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*pushTemplateState)(nil)).Elem()
}

type pushTemplateArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-adm
	ADM *PushTemplateAndroidPushNotificationTemplate `pulumi:"aDM"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-apns
	APNS *PushTemplateAPNSPushNotificationTemplate `pulumi:"aPNS"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-baidu
	Baidu *PushTemplateAndroidPushNotificationTemplate `pulumi:"baidu"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-default
	Default *PushTemplateDefaultPushNotificationTemplate `pulumi:"default"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-defaultsubstitutions
	DefaultSubstitutions *string `pulumi:"defaultSubstitutions"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-gcm
	GCM *PushTemplateAndroidPushNotificationTemplate `pulumi:"gCM"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-tags
	Tags interface{} `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-templatedescription
	TemplateDescription *string `pulumi:"templateDescription"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-templatename
	TemplateName string `pulumi:"templateName"`
}

// The set of arguments for constructing a PushTemplate resource.
type PushTemplateArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-adm
	ADM PushTemplateAndroidPushNotificationTemplatePtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-apns
	APNS PushTemplateAPNSPushNotificationTemplatePtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-baidu
	Baidu PushTemplateAndroidPushNotificationTemplatePtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-default
	Default PushTemplateDefaultPushNotificationTemplatePtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-defaultsubstitutions
	DefaultSubstitutions pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-gcm
	GCM PushTemplateAndroidPushNotificationTemplatePtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-tags
	Tags pulumi.Input
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-templatedescription
	TemplateDescription pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-templatename
	TemplateName pulumi.StringInput
}

func (PushTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pushTemplateArgs)(nil)).Elem()
}

type PushTemplateInput interface {
	pulumi.Input

	ToPushTemplateOutput() PushTemplateOutput
	ToPushTemplateOutputWithContext(ctx context.Context) PushTemplateOutput
}

func (*PushTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((*PushTemplate)(nil))
}

func (i *PushTemplate) ToPushTemplateOutput() PushTemplateOutput {
	return i.ToPushTemplateOutputWithContext(context.Background())
}

func (i *PushTemplate) ToPushTemplateOutputWithContext(ctx context.Context) PushTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PushTemplateOutput)
}

type PushTemplateOutput struct{ *pulumi.OutputState }

func (PushTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PushTemplate)(nil))
}

func (o PushTemplateOutput) ToPushTemplateOutput() PushTemplateOutput {
	return o
}

func (o PushTemplateOutput) ToPushTemplateOutputWithContext(ctx context.Context) PushTemplateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PushTemplateOutput{})
}