// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::DataSync::LocationEFS.
func LookupLocationEFS(ctx *pulumi.Context, args *LookupLocationEFSArgs, opts ...pulumi.InvokeOption) (*LookupLocationEFSResult, error) {
	var rv LookupLocationEFSResult
	err := ctx.Invoke("aws-native:datasync:getLocationEFS", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLocationEFSArgs struct {
	// The Amazon Resource Name (ARN) of the Amazon EFS file system location that is created.
	LocationArn string `pulumi:"locationArn"`
}

type LookupLocationEFSResult struct {
	// The Amazon Resource Name (ARN) of the Amazon EFS file system location that is created.
	LocationArn *string `pulumi:"locationArn"`
	// The URL of the EFS location that was described.
	LocationUri *string `pulumi:"locationUri"`
	// An array of key-value pairs to apply to this resource.
	Tags []LocationEFSTag `pulumi:"tags"`
}

func LookupLocationEFSOutput(ctx *pulumi.Context, args LookupLocationEFSOutputArgs, opts ...pulumi.InvokeOption) LookupLocationEFSResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLocationEFSResult, error) {
			args := v.(LookupLocationEFSArgs)
			r, err := LookupLocationEFS(ctx, &args, opts...)
			return *r, err
		}).(LookupLocationEFSResultOutput)
}

type LookupLocationEFSOutputArgs struct {
	// The Amazon Resource Name (ARN) of the Amazon EFS file system location that is created.
	LocationArn pulumi.StringInput `pulumi:"locationArn"`
}

func (LookupLocationEFSOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocationEFSArgs)(nil)).Elem()
}

type LookupLocationEFSResultOutput struct{ *pulumi.OutputState }

func (LookupLocationEFSResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocationEFSResult)(nil)).Elem()
}

func (o LookupLocationEFSResultOutput) ToLookupLocationEFSResultOutput() LookupLocationEFSResultOutput {
	return o
}

func (o LookupLocationEFSResultOutput) ToLookupLocationEFSResultOutputWithContext(ctx context.Context) LookupLocationEFSResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the Amazon EFS file system location that is created.
func (o LookupLocationEFSResultOutput) LocationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationEFSResult) *string { return v.LocationArn }).(pulumi.StringPtrOutput)
}

// The URL of the EFS location that was described.
func (o LookupLocationEFSResultOutput) LocationUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationEFSResult) *string { return v.LocationUri }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupLocationEFSResultOutput) Tags() LocationEFSTagArrayOutput {
	return o.ApplyT(func(v LookupLocationEFSResult) []LocationEFSTag { return v.Tags }).(LocationEFSTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocationEFSResultOutput{})
}