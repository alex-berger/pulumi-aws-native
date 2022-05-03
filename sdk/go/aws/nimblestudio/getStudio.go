// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nimblestudio

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a studio that contains other Nimble Studio resources
func LookupStudio(ctx *pulumi.Context, args *LookupStudioArgs, opts ...pulumi.InvokeOption) (*LookupStudioResult, error) {
	var rv LookupStudioResult
	err := ctx.Invoke("aws-native:nimblestudio:getStudio", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStudioArgs struct {
	StudioId string `pulumi:"studioId"`
}

type LookupStudioResult struct {
	// <p>The IAM role that Studio Admins will assume when logging in to the Nimble Studio portal.</p>
	AdminRoleArn *string `pulumi:"adminRoleArn"`
	// <p>A friendly name for the studio.</p>
	DisplayName *string `pulumi:"displayName"`
	// <p>The Amazon Web Services Region where the studio resource is located.</p>
	HomeRegion *string `pulumi:"homeRegion"`
	// <p>The Amazon Web Services SSO application client ID used to integrate with Amazon Web Services SSO to enable Amazon Web Services SSO users to log in to Nimble Studio portal.</p>
	SsoClientId                   *string                        `pulumi:"ssoClientId"`
	StudioEncryptionConfiguration *StudioEncryptionConfiguration `pulumi:"studioEncryptionConfiguration"`
	StudioId                      *string                        `pulumi:"studioId"`
	// <p>The address of the web page for the studio.</p>
	StudioUrl *string `pulumi:"studioUrl"`
	// <p>The IAM role that Studio Users will assume when logging in to the Nimble Studio portal.</p>
	UserRoleArn *string `pulumi:"userRoleArn"`
}

func LookupStudioOutput(ctx *pulumi.Context, args LookupStudioOutputArgs, opts ...pulumi.InvokeOption) LookupStudioResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStudioResult, error) {
			args := v.(LookupStudioArgs)
			r, err := LookupStudio(ctx, &args, opts...)
			var s LookupStudioResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStudioResultOutput)
}

type LookupStudioOutputArgs struct {
	StudioId pulumi.StringInput `pulumi:"studioId"`
}

func (LookupStudioOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStudioArgs)(nil)).Elem()
}

type LookupStudioResultOutput struct{ *pulumi.OutputState }

func (LookupStudioResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStudioResult)(nil)).Elem()
}

func (o LookupStudioResultOutput) ToLookupStudioResultOutput() LookupStudioResultOutput {
	return o
}

func (o LookupStudioResultOutput) ToLookupStudioResultOutputWithContext(ctx context.Context) LookupStudioResultOutput {
	return o
}

// <p>The IAM role that Studio Admins will assume when logging in to the Nimble Studio portal.</p>
func (o LookupStudioResultOutput) AdminRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioResult) *string { return v.AdminRoleArn }).(pulumi.StringPtrOutput)
}

// <p>A friendly name for the studio.</p>
func (o LookupStudioResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// <p>The Amazon Web Services Region where the studio resource is located.</p>
func (o LookupStudioResultOutput) HomeRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioResult) *string { return v.HomeRegion }).(pulumi.StringPtrOutput)
}

// <p>The Amazon Web Services SSO application client ID used to integrate with Amazon Web Services SSO to enable Amazon Web Services SSO users to log in to Nimble Studio portal.</p>
func (o LookupStudioResultOutput) SsoClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioResult) *string { return v.SsoClientId }).(pulumi.StringPtrOutput)
}

func (o LookupStudioResultOutput) StudioEncryptionConfiguration() StudioEncryptionConfigurationPtrOutput {
	return o.ApplyT(func(v LookupStudioResult) *StudioEncryptionConfiguration { return v.StudioEncryptionConfiguration }).(StudioEncryptionConfigurationPtrOutput)
}

func (o LookupStudioResultOutput) StudioId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioResult) *string { return v.StudioId }).(pulumi.StringPtrOutput)
}

// <p>The address of the web page for the studio.</p>
func (o LookupStudioResultOutput) StudioUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioResult) *string { return v.StudioUrl }).(pulumi.StringPtrOutput)
}

// <p>The IAM role that Studio Users will assume when logging in to the Nimble Studio portal.</p>
func (o LookupStudioResultOutput) UserRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioResult) *string { return v.UserRoleArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStudioResultOutput{})
}
