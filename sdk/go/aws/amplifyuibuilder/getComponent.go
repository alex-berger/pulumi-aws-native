// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package amplifyuibuilder

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::AmplifyUIBuilder::Component Resource Type
func LookupComponent(ctx *pulumi.Context, args *LookupComponentArgs, opts ...pulumi.InvokeOption) (*LookupComponentResult, error) {
	var rv LookupComponentResult
	err := ctx.Invoke("aws-native:amplifyuibuilder:getComponent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupComponentArgs struct {
	AppId           string `pulumi:"appId"`
	EnvironmentName string `pulumi:"environmentName"`
	Id              string `pulumi:"id"`
}

type LookupComponentResult struct {
	AppId                *string                        `pulumi:"appId"`
	BindingProperties    *ComponentBindingProperties    `pulumi:"bindingProperties"`
	Children             []ComponentChild               `pulumi:"children"`
	CollectionProperties *ComponentCollectionProperties `pulumi:"collectionProperties"`
	ComponentType        *string                        `pulumi:"componentType"`
	EnvironmentName      *string                        `pulumi:"environmentName"`
	Events               *ComponentEvents               `pulumi:"events"`
	Id                   *string                        `pulumi:"id"`
	Name                 *string                        `pulumi:"name"`
	Overrides            *ComponentOverrides            `pulumi:"overrides"`
	Properties           *ComponentProperties           `pulumi:"properties"`
	SchemaVersion        *string                        `pulumi:"schemaVersion"`
	SourceId             *string                        `pulumi:"sourceId"`
	Variants             []ComponentVariant             `pulumi:"variants"`
}

func LookupComponentOutput(ctx *pulumi.Context, args LookupComponentOutputArgs, opts ...pulumi.InvokeOption) LookupComponentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupComponentResult, error) {
			args := v.(LookupComponentArgs)
			r, err := LookupComponent(ctx, &args, opts...)
			var s LookupComponentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupComponentResultOutput)
}

type LookupComponentOutputArgs struct {
	AppId           pulumi.StringInput `pulumi:"appId"`
	EnvironmentName pulumi.StringInput `pulumi:"environmentName"`
	Id              pulumi.StringInput `pulumi:"id"`
}

func (LookupComponentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComponentArgs)(nil)).Elem()
}

type LookupComponentResultOutput struct{ *pulumi.OutputState }

func (LookupComponentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComponentResult)(nil)).Elem()
}

func (o LookupComponentResultOutput) ToLookupComponentResultOutput() LookupComponentResultOutput {
	return o
}

func (o LookupComponentResultOutput) ToLookupComponentResultOutputWithContext(ctx context.Context) LookupComponentResultOutput {
	return o
}

func (o LookupComponentResultOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o LookupComponentResultOutput) BindingProperties() ComponentBindingPropertiesPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *ComponentBindingProperties { return v.BindingProperties }).(ComponentBindingPropertiesPtrOutput)
}

func (o LookupComponentResultOutput) Children() ComponentChildArrayOutput {
	return o.ApplyT(func(v LookupComponentResult) []ComponentChild { return v.Children }).(ComponentChildArrayOutput)
}

func (o LookupComponentResultOutput) CollectionProperties() ComponentCollectionPropertiesPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *ComponentCollectionProperties { return v.CollectionProperties }).(ComponentCollectionPropertiesPtrOutput)
}

func (o LookupComponentResultOutput) ComponentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *string { return v.ComponentType }).(pulumi.StringPtrOutput)
}

func (o LookupComponentResultOutput) EnvironmentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *string { return v.EnvironmentName }).(pulumi.StringPtrOutput)
}

func (o LookupComponentResultOutput) Events() ComponentEventsPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *ComponentEvents { return v.Events }).(ComponentEventsPtrOutput)
}

func (o LookupComponentResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupComponentResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupComponentResultOutput) Overrides() ComponentOverridesPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *ComponentOverrides { return v.Overrides }).(ComponentOverridesPtrOutput)
}

func (o LookupComponentResultOutput) Properties() ComponentPropertiesPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *ComponentProperties { return v.Properties }).(ComponentPropertiesPtrOutput)
}

func (o LookupComponentResultOutput) SchemaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *string { return v.SchemaVersion }).(pulumi.StringPtrOutput)
}

func (o LookupComponentResultOutput) SourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *string { return v.SourceId }).(pulumi.StringPtrOutput)
}

func (o LookupComponentResultOutput) Variants() ComponentVariantArrayOutput {
	return o.ApplyT(func(v LookupComponentResult) []ComponentVariant { return v.Variants }).(ComponentVariantArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComponentResultOutput{})
}
