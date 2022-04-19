// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3Outposts
{
    public static class GetBucket
    {
        /// <summary>
        /// Resource Type Definition for AWS::S3Outposts::Bucket
        /// </summary>
        public static Task<GetBucketResult> InvokeAsync(GetBucketArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBucketResult>("aws-native:s3outposts:getBucket", args ?? new GetBucketArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type Definition for AWS::S3Outposts::Bucket
        /// </summary>
        public static Output<GetBucketResult> Invoke(GetBucketInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBucketResult>("aws-native:s3outposts:getBucket", args ?? new GetBucketInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBucketArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the specified bucket.
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetBucketArgs()
        {
        }
    }

    public sealed class GetBucketInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the specified bucket.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetBucketInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBucketResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the specified bucket.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// Rules that define how Amazon S3Outposts manages objects during their lifetime.
        /// </summary>
        public readonly Outputs.BucketLifecycleConfiguration? LifecycleConfiguration;
        /// <summary>
        /// An arbitrary set of tags (key-value pairs) for this S3Outposts bucket.
        /// </summary>
        public readonly ImmutableArray<Outputs.BucketTag> Tags;

        [OutputConstructor]
        private GetBucketResult(
            string? arn,

            Outputs.BucketLifecycleConfiguration? lifecycleConfiguration,

            ImmutableArray<Outputs.BucketTag> tags)
        {
            Arn = arn;
            LifecycleConfiguration = lifecycleConfiguration;
            Tags = tags;
        }
    }
}
