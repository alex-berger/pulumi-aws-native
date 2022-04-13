// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront
{
    public static class GetCloudFrontOriginAccessIdentity
    {
        /// <summary>
        /// Resource Type definition for AWS::CloudFront::CloudFrontOriginAccessIdentity
        /// </summary>
        public static Task<GetCloudFrontOriginAccessIdentityResult> InvokeAsync(GetCloudFrontOriginAccessIdentityArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCloudFrontOriginAccessIdentityResult>("aws-native:cloudfront:getCloudFrontOriginAccessIdentity", args ?? new GetCloudFrontOriginAccessIdentityArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::CloudFront::CloudFrontOriginAccessIdentity
        /// </summary>
        public static Output<GetCloudFrontOriginAccessIdentityResult> Invoke(GetCloudFrontOriginAccessIdentityInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCloudFrontOriginAccessIdentityResult>("aws-native:cloudfront:getCloudFrontOriginAccessIdentity", args ?? new GetCloudFrontOriginAccessIdentityInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudFrontOriginAccessIdentityArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetCloudFrontOriginAccessIdentityArgs()
        {
        }
    }

    public sealed class GetCloudFrontOriginAccessIdentityInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetCloudFrontOriginAccessIdentityInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCloudFrontOriginAccessIdentityResult
    {
        public readonly Outputs.CloudFrontOriginAccessIdentityConfig? CloudFrontOriginAccessIdentityConfig;
        public readonly string? Id;
        public readonly string? S3CanonicalUserId;

        [OutputConstructor]
        private GetCloudFrontOriginAccessIdentityResult(
            Outputs.CloudFrontOriginAccessIdentityConfig? cloudFrontOriginAccessIdentityConfig,

            string? id,

            string? s3CanonicalUserId)
        {
            CloudFrontOriginAccessIdentityConfig = cloudFrontOriginAccessIdentityConfig;
            Id = id;
            S3CanonicalUserId = s3CanonicalUserId;
        }
    }
}
