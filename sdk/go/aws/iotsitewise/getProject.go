// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotsitewise

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::IoTSiteWise::Project
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	var rv LookupProjectResult
	err := ctx.Invoke("aws-native:iotsitewise:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProjectArgs struct {
	// The ID of the project.
	ProjectId string `pulumi:"projectId"`
}

type LookupProjectResult struct {
	// The IDs of the assets to be associated to the project.
	AssetIds []string `pulumi:"assetIds"`
	// The ARN of the project.
	ProjectArn *string `pulumi:"projectArn"`
	// A description for the project.
	ProjectDescription *string `pulumi:"projectDescription"`
	// The ID of the project.
	ProjectId *string `pulumi:"projectId"`
	// A friendly name for the project.
	ProjectName *string `pulumi:"projectName"`
	// A list of key-value pairs that contain metadata for the project.
	Tags []ProjectTag `pulumi:"tags"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectResult, error) {
			args := v.(LookupProjectArgs)
			r, err := LookupProject(ctx, &args, opts...)
			var s LookupProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectResultOutput)
}

type LookupProjectOutputArgs struct {
	// The ID of the project.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

// The IDs of the assets to be associated to the project.
func (o LookupProjectResultOutput) AssetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []string { return v.AssetIds }).(pulumi.StringArrayOutput)
}

// The ARN of the project.
func (o LookupProjectResultOutput) ProjectArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.ProjectArn }).(pulumi.StringPtrOutput)
}

// A description for the project.
func (o LookupProjectResultOutput) ProjectDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.ProjectDescription }).(pulumi.StringPtrOutput)
}

// The ID of the project.
func (o LookupProjectResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// A friendly name for the project.
func (o LookupProjectResultOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// A list of key-value pairs that contain metadata for the project.
func (o LookupProjectResultOutput) Tags() ProjectTagArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []ProjectTag { return v.Tags }).(ProjectTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
