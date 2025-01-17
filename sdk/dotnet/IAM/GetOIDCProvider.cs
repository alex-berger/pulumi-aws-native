// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IAM
{
    public static class GetOIDCProvider
    {
        /// <summary>
        /// Resource Type definition for AWS::IAM::OIDCProvider
        /// </summary>
        public static Task<GetOIDCProviderResult> InvokeAsync(GetOIDCProviderArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOIDCProviderResult>("aws-native:iam:getOIDCProvider", args ?? new GetOIDCProviderArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::IAM::OIDCProvider
        /// </summary>
        public static Output<GetOIDCProviderResult> Invoke(GetOIDCProviderInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetOIDCProviderResult>("aws-native:iam:getOIDCProvider", args ?? new GetOIDCProviderInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOIDCProviderArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the OIDC provider
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetOIDCProviderArgs()
        {
        }
    }

    public sealed class GetOIDCProviderInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the OIDC provider
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetOIDCProviderInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetOIDCProviderResult
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the OIDC provider
        /// </summary>
        public readonly string? Arn;
        public readonly ImmutableArray<string> ClientIdList;
        public readonly ImmutableArray<Outputs.OIDCProviderTag> Tags;
        public readonly ImmutableArray<string> ThumbprintList;

        [OutputConstructor]
        private GetOIDCProviderResult(
            string? arn,

            ImmutableArray<string> clientIdList,

            ImmutableArray<Outputs.OIDCProviderTag> tags,

            ImmutableArray<string> thumbprintList)
        {
            Arn = arn;
            ClientIdList = clientIdList;
            Tags = tags;
            ThumbprintList = thumbprintList;
        }
    }
}
