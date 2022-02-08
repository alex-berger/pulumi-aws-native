// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3ObjectLambda
{
    public static class GetAccessPointPolicy
    {
        /// <summary>
        /// AWS::S3ObjectLambda::AccessPointPolicy resource is an Amazon S3ObjectLambda policy type that you can use to control permissions for your S3ObjectLambda
        /// </summary>
        public static Task<GetAccessPointPolicyResult> InvokeAsync(GetAccessPointPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccessPointPolicyResult>("aws-native:s3objectlambda:getAccessPointPolicy", args ?? new GetAccessPointPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// AWS::S3ObjectLambda::AccessPointPolicy resource is an Amazon S3ObjectLambda policy type that you can use to control permissions for your S3ObjectLambda
        /// </summary>
        public static Output<GetAccessPointPolicyResult> Invoke(GetAccessPointPolicyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAccessPointPolicyResult>("aws-native:s3objectlambda:getAccessPointPolicy", args ?? new GetAccessPointPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccessPointPolicyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Amazon S3 ObjectLambdaAccessPoint to which the policy applies.
        /// </summary>
        [Input("objectLambdaAccessPoint", required: true)]
        public string ObjectLambdaAccessPoint { get; set; } = null!;

        public GetAccessPointPolicyArgs()
        {
        }
    }

    public sealed class GetAccessPointPolicyInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Amazon S3 ObjectLambdaAccessPoint to which the policy applies.
        /// </summary>
        [Input("objectLambdaAccessPoint", required: true)]
        public Input<string> ObjectLambdaAccessPoint { get; set; } = null!;

        public GetAccessPointPolicyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAccessPointPolicyResult
    {
        /// <summary>
        /// A policy document containing permissions to add to the specified ObjectLambdaAccessPoint. For more information, see Access Policy Language Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the Amazon Simple Storage Service Developer Guide. 
        /// </summary>
        public readonly object? PolicyDocument;

        [OutputConstructor]
        private GetAccessPointPolicyResult(object? policyDocument)
        {
            PolicyDocument = policyDocument;
        }
    }
}