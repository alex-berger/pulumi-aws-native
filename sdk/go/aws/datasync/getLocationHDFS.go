// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::DataSync::LocationHDFS.
func LookupLocationHDFS(ctx *pulumi.Context, args *LookupLocationHDFSArgs, opts ...pulumi.InvokeOption) (*LookupLocationHDFSResult, error) {
	var rv LookupLocationHDFSResult
	err := ctx.Invoke("aws-native:datasync:getLocationHDFS", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLocationHDFSArgs struct {
	// The Amazon Resource Name (ARN) of the HDFS location.
	LocationArn string `pulumi:"locationArn"`
}

type LookupLocationHDFSResult struct {
	// ARN(s) of the agent(s) to use for an HDFS location.
	AgentArns []string `pulumi:"agentArns"`
	// The authentication mode used to determine identity of user.
	AuthenticationType *LocationHDFSAuthenticationType `pulumi:"authenticationType"`
	// Size of chunks (blocks) in bytes that the data is divided into when stored in the HDFS cluster.
	BlockSize *int `pulumi:"blockSize"`
	// The unique identity, or principal, to which Kerberos can assign tickets.
	KerberosPrincipal *string `pulumi:"kerberosPrincipal"`
	// The identifier for the Key Management Server where the encryption keys that encrypt data inside HDFS clusters are stored.
	KmsKeyProviderUri *string `pulumi:"kmsKeyProviderUri"`
	// The Amazon Resource Name (ARN) of the HDFS location.
	LocationArn *string `pulumi:"locationArn"`
	// The URL of the HDFS location that was described.
	LocationUri *string `pulumi:"locationUri"`
	// An array of Name Node(s) of the HDFS location.
	NameNodes        []LocationHDFSNameNode        `pulumi:"nameNodes"`
	QopConfiguration *LocationHDFSQopConfiguration `pulumi:"qopConfiguration"`
	// Number of copies of each block that exists inside the HDFS cluster.
	ReplicationFactor *int `pulumi:"replicationFactor"`
	// The user name that has read and write permissions on the specified HDFS cluster.
	SimpleUser *string `pulumi:"simpleUser"`
	// An array of key-value pairs to apply to this resource.
	Tags []LocationHDFSTag `pulumi:"tags"`
}

func LookupLocationHDFSOutput(ctx *pulumi.Context, args LookupLocationHDFSOutputArgs, opts ...pulumi.InvokeOption) LookupLocationHDFSResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLocationHDFSResult, error) {
			args := v.(LookupLocationHDFSArgs)
			r, err := LookupLocationHDFS(ctx, &args, opts...)
			var s LookupLocationHDFSResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLocationHDFSResultOutput)
}

type LookupLocationHDFSOutputArgs struct {
	// The Amazon Resource Name (ARN) of the HDFS location.
	LocationArn pulumi.StringInput `pulumi:"locationArn"`
}

func (LookupLocationHDFSOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocationHDFSArgs)(nil)).Elem()
}

type LookupLocationHDFSResultOutput struct{ *pulumi.OutputState }

func (LookupLocationHDFSResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocationHDFSResult)(nil)).Elem()
}

func (o LookupLocationHDFSResultOutput) ToLookupLocationHDFSResultOutput() LookupLocationHDFSResultOutput {
	return o
}

func (o LookupLocationHDFSResultOutput) ToLookupLocationHDFSResultOutputWithContext(ctx context.Context) LookupLocationHDFSResultOutput {
	return o
}

// ARN(s) of the agent(s) to use for an HDFS location.
func (o LookupLocationHDFSResultOutput) AgentArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocationHDFSResult) []string { return v.AgentArns }).(pulumi.StringArrayOutput)
}

// The authentication mode used to determine identity of user.
func (o LookupLocationHDFSResultOutput) AuthenticationType() LocationHDFSAuthenticationTypePtrOutput {
	return o.ApplyT(func(v LookupLocationHDFSResult) *LocationHDFSAuthenticationType { return v.AuthenticationType }).(LocationHDFSAuthenticationTypePtrOutput)
}

// Size of chunks (blocks) in bytes that the data is divided into when stored in the HDFS cluster.
func (o LookupLocationHDFSResultOutput) BlockSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupLocationHDFSResult) *int { return v.BlockSize }).(pulumi.IntPtrOutput)
}

// The unique identity, or principal, to which Kerberos can assign tickets.
func (o LookupLocationHDFSResultOutput) KerberosPrincipal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationHDFSResult) *string { return v.KerberosPrincipal }).(pulumi.StringPtrOutput)
}

// The identifier for the Key Management Server where the encryption keys that encrypt data inside HDFS clusters are stored.
func (o LookupLocationHDFSResultOutput) KmsKeyProviderUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationHDFSResult) *string { return v.KmsKeyProviderUri }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the HDFS location.
func (o LookupLocationHDFSResultOutput) LocationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationHDFSResult) *string { return v.LocationArn }).(pulumi.StringPtrOutput)
}

// The URL of the HDFS location that was described.
func (o LookupLocationHDFSResultOutput) LocationUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationHDFSResult) *string { return v.LocationUri }).(pulumi.StringPtrOutput)
}

// An array of Name Node(s) of the HDFS location.
func (o LookupLocationHDFSResultOutput) NameNodes() LocationHDFSNameNodeArrayOutput {
	return o.ApplyT(func(v LookupLocationHDFSResult) []LocationHDFSNameNode { return v.NameNodes }).(LocationHDFSNameNodeArrayOutput)
}

func (o LookupLocationHDFSResultOutput) QopConfiguration() LocationHDFSQopConfigurationPtrOutput {
	return o.ApplyT(func(v LookupLocationHDFSResult) *LocationHDFSQopConfiguration { return v.QopConfiguration }).(LocationHDFSQopConfigurationPtrOutput)
}

// Number of copies of each block that exists inside the HDFS cluster.
func (o LookupLocationHDFSResultOutput) ReplicationFactor() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupLocationHDFSResult) *int { return v.ReplicationFactor }).(pulumi.IntPtrOutput)
}

// The user name that has read and write permissions on the specified HDFS cluster.
func (o LookupLocationHDFSResultOutput) SimpleUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationHDFSResult) *string { return v.SimpleUser }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupLocationHDFSResultOutput) Tags() LocationHDFSTagArrayOutput {
	return o.ApplyT(func(v LookupLocationHDFSResult) []LocationHDFSTag { return v.Tags }).(LocationHDFSTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocationHDFSResultOutput{})
}
