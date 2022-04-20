// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the AWS::QuickSight::Analysis Resource Type.
func LookupAnalysis(ctx *pulumi.Context, args *LookupAnalysisArgs, opts ...pulumi.InvokeOption) (*LookupAnalysisResult, error) {
	var rv LookupAnalysisResult
	err := ctx.Invoke("aws-native:quicksight:getAnalysis", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAnalysisArgs struct {
	AnalysisId   string `pulumi:"analysisId"`
	AwsAccountId string `pulumi:"awsAccountId"`
}

type LookupAnalysisResult struct {
	// <p>The Amazon Resource Name (ARN) of the analysis.</p>
	Arn *string `pulumi:"arn"`
	// <p>The time that the analysis was created.</p>
	CreatedTime *string `pulumi:"createdTime"`
	// <p>The ARNs of the datasets of the analysis.</p>
	DataSetArns []string `pulumi:"dataSetArns"`
	// <p>Errors associated with the analysis.</p>
	Errors []AnalysisError `pulumi:"errors"`
	// <p>The descriptive name of the analysis.</p>
	Name *string `pulumi:"name"`
	// <p>A structure that describes the principals and the resource-level permissions on an
	//             analysis. You can use the <code>Permissions</code> structure to grant permissions by
	//             providing a list of AWS Identity and Access Management (IAM) action information for each
	//             principal listed by Amazon Resource Name (ARN). </p>
	//
	//         <p>To specify no permissions, omit <code>Permissions</code>.</p>
	Permissions []AnalysisResourcePermission `pulumi:"permissions"`
	// <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the
	//             analysis.</p>
	Tags []AnalysisTag `pulumi:"tags"`
	// <p>The ARN of the theme of the analysis.</p>
	ThemeArn *string `pulumi:"themeArn"`
}

func LookupAnalysisOutput(ctx *pulumi.Context, args LookupAnalysisOutputArgs, opts ...pulumi.InvokeOption) LookupAnalysisResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAnalysisResult, error) {
			args := v.(LookupAnalysisArgs)
			r, err := LookupAnalysis(ctx, &args, opts...)
			var s LookupAnalysisResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAnalysisResultOutput)
}

type LookupAnalysisOutputArgs struct {
	AnalysisId   pulumi.StringInput `pulumi:"analysisId"`
	AwsAccountId pulumi.StringInput `pulumi:"awsAccountId"`
}

func (LookupAnalysisOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnalysisArgs)(nil)).Elem()
}

type LookupAnalysisResultOutput struct{ *pulumi.OutputState }

func (LookupAnalysisResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnalysisResult)(nil)).Elem()
}

func (o LookupAnalysisResultOutput) ToLookupAnalysisResultOutput() LookupAnalysisResultOutput {
	return o
}

func (o LookupAnalysisResultOutput) ToLookupAnalysisResultOutputWithContext(ctx context.Context) LookupAnalysisResultOutput {
	return o
}

// <p>The Amazon Resource Name (ARN) of the analysis.</p>
func (o LookupAnalysisResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnalysisResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// <p>The time that the analysis was created.</p>
func (o LookupAnalysisResultOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnalysisResult) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

// <p>The ARNs of the datasets of the analysis.</p>
func (o LookupAnalysisResultOutput) DataSetArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAnalysisResult) []string { return v.DataSetArns }).(pulumi.StringArrayOutput)
}

// <p>Errors associated with the analysis.</p>
func (o LookupAnalysisResultOutput) Errors() AnalysisErrorArrayOutput {
	return o.ApplyT(func(v LookupAnalysisResult) []AnalysisError { return v.Errors }).(AnalysisErrorArrayOutput)
}

// <p>The descriptive name of the analysis.</p>
func (o LookupAnalysisResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnalysisResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// <p>A structure that describes the principals and the resource-level permissions on an
//             analysis. You can use the <code>Permissions</code> structure to grant permissions by
//             providing a list of AWS Identity and Access Management (IAM) action information for each
//             principal listed by Amazon Resource Name (ARN). </p>
//
//         <p>To specify no permissions, omit <code>Permissions</code>.</p>
func (o LookupAnalysisResultOutput) Permissions() AnalysisResourcePermissionArrayOutput {
	return o.ApplyT(func(v LookupAnalysisResult) []AnalysisResourcePermission { return v.Permissions }).(AnalysisResourcePermissionArrayOutput)
}

// <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the
//             analysis.</p>
func (o LookupAnalysisResultOutput) Tags() AnalysisTagArrayOutput {
	return o.ApplyT(func(v LookupAnalysisResult) []AnalysisTag { return v.Tags }).(AnalysisTagArrayOutput)
}

// <p>The ARN of the theme of the analysis.</p>
func (o LookupAnalysisResultOutput) ThemeArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnalysisResult) *string { return v.ThemeArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAnalysisResultOutput{})
}
