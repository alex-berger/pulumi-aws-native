// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::DataSync::LocationObjectStorage.
func LookupLocationObjectStorage(ctx *pulumi.Context, args *LookupLocationObjectStorageArgs, opts ...pulumi.InvokeOption) (*LookupLocationObjectStorageResult, error) {
	var rv LookupLocationObjectStorageResult
	err := ctx.Invoke("aws-native:datasync:getLocationObjectStorage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLocationObjectStorageArgs struct {
	// The Amazon Resource Name (ARN) of the location that is created.
	LocationArn string `pulumi:"locationArn"`
}

type LookupLocationObjectStorageResult struct {
	// Optional. The access key is used if credentials are required to access the self-managed object storage server.
	AccessKey *string `pulumi:"accessKey"`
	// The Amazon Resource Name (ARN) of the agents associated with the self-managed object storage server location.
	AgentArns []string `pulumi:"agentArns"`
	// The Amazon Resource Name (ARN) of the location that is created.
	LocationArn *string `pulumi:"locationArn"`
	// The URL of the object storage location that was described.
	LocationUri *string `pulumi:"locationUri"`
	// The port that your self-managed server accepts inbound network traffic on.
	ServerPort *int `pulumi:"serverPort"`
	// The protocol that the object storage server uses to communicate.
	ServerProtocol *LocationObjectStorageServerProtocol `pulumi:"serverProtocol"`
	// An array of key-value pairs to apply to this resource.
	Tags []LocationObjectStorageTag `pulumi:"tags"`
}

func LookupLocationObjectStorageOutput(ctx *pulumi.Context, args LookupLocationObjectStorageOutputArgs, opts ...pulumi.InvokeOption) LookupLocationObjectStorageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLocationObjectStorageResult, error) {
			args := v.(LookupLocationObjectStorageArgs)
			r, err := LookupLocationObjectStorage(ctx, &args, opts...)
			var s LookupLocationObjectStorageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLocationObjectStorageResultOutput)
}

type LookupLocationObjectStorageOutputArgs struct {
	// The Amazon Resource Name (ARN) of the location that is created.
	LocationArn pulumi.StringInput `pulumi:"locationArn"`
}

func (LookupLocationObjectStorageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocationObjectStorageArgs)(nil)).Elem()
}

type LookupLocationObjectStorageResultOutput struct{ *pulumi.OutputState }

func (LookupLocationObjectStorageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocationObjectStorageResult)(nil)).Elem()
}

func (o LookupLocationObjectStorageResultOutput) ToLookupLocationObjectStorageResultOutput() LookupLocationObjectStorageResultOutput {
	return o
}

func (o LookupLocationObjectStorageResultOutput) ToLookupLocationObjectStorageResultOutputWithContext(ctx context.Context) LookupLocationObjectStorageResultOutput {
	return o
}

// Optional. The access key is used if credentials are required to access the self-managed object storage server.
func (o LookupLocationObjectStorageResultOutput) AccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationObjectStorageResult) *string { return v.AccessKey }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the agents associated with the self-managed object storage server location.
func (o LookupLocationObjectStorageResultOutput) AgentArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocationObjectStorageResult) []string { return v.AgentArns }).(pulumi.StringArrayOutput)
}

// The Amazon Resource Name (ARN) of the location that is created.
func (o LookupLocationObjectStorageResultOutput) LocationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationObjectStorageResult) *string { return v.LocationArn }).(pulumi.StringPtrOutput)
}

// The URL of the object storage location that was described.
func (o LookupLocationObjectStorageResultOutput) LocationUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationObjectStorageResult) *string { return v.LocationUri }).(pulumi.StringPtrOutput)
}

// The port that your self-managed server accepts inbound network traffic on.
func (o LookupLocationObjectStorageResultOutput) ServerPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupLocationObjectStorageResult) *int { return v.ServerPort }).(pulumi.IntPtrOutput)
}

// The protocol that the object storage server uses to communicate.
func (o LookupLocationObjectStorageResultOutput) ServerProtocol() LocationObjectStorageServerProtocolPtrOutput {
	return o.ApplyT(func(v LookupLocationObjectStorageResult) *LocationObjectStorageServerProtocol {
		return v.ServerProtocol
	}).(LocationObjectStorageServerProtocolPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupLocationObjectStorageResultOutput) Tags() LocationObjectStorageTagArrayOutput {
	return o.ApplyT(func(v LookupLocationObjectStorageResult) []LocationObjectStorageTag { return v.Tags }).(LocationObjectStorageTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocationObjectStorageResultOutput{})
}
