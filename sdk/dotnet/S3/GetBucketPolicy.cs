// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3
{
    public static class GetBucketPolicy
    {
        /// <summary>
        /// Resource Type definition for AWS::S3::BucketPolicy
        /// </summary>
        public static Task<GetBucketPolicyResult> InvokeAsync(GetBucketPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBucketPolicyResult>("aws-native:s3:getBucketPolicy", args ?? new GetBucketPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::S3::BucketPolicy
        /// </summary>
        public static Output<GetBucketPolicyResult> Invoke(GetBucketPolicyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBucketPolicyResult>("aws-native:s3:getBucketPolicy", args ?? new GetBucketPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBucketPolicyArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetBucketPolicyArgs()
        {
        }
    }

    public sealed class GetBucketPolicyInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetBucketPolicyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBucketPolicyResult
    {
        public readonly string? Id;
        public readonly object? PolicyDocument;

        [OutputConstructor]
        private GetBucketPolicyResult(
            string? id,

            object? policyDocument)
        {
            Id = id;
            PolicyDocument = policyDocument;
        }
    }
}
