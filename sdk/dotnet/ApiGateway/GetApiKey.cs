// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway
{
    public static class GetApiKey
    {
        /// <summary>
        /// Resource Type definition for AWS::ApiGateway::ApiKey
        /// </summary>
        public static Task<GetApiKeyResult> InvokeAsync(GetApiKeyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApiKeyResult>("aws-native:apigateway:getApiKey", args ?? new GetApiKeyArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ApiGateway::ApiKey
        /// </summary>
        public static Output<GetApiKeyResult> Invoke(GetApiKeyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetApiKeyResult>("aws-native:apigateway:getApiKey", args ?? new GetApiKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApiKeyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A Unique Key ID which identifies the API Key. Generated by the Create API and returned by the Read and List APIs 
        /// </summary>
        [Input("aPIKeyId", required: true)]
        public string APIKeyId { get; set; } = null!;

        public GetApiKeyArgs()
        {
        }
    }

    public sealed class GetApiKeyInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A Unique Key ID which identifies the API Key. Generated by the Create API and returned by the Read and List APIs 
        /// </summary>
        [Input("aPIKeyId", required: true)]
        public Input<string> APIKeyId { get; set; } = null!;

        public GetApiKeyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApiKeyResult
    {
        /// <summary>
        /// A Unique Key ID which identifies the API Key. Generated by the Create API and returned by the Read and List APIs 
        /// </summary>
        public readonly string? APIKeyId;
        /// <summary>
        /// An AWS Marketplace customer identifier to use when integrating with the AWS SaaS Marketplace.
        /// </summary>
        public readonly string? CustomerId;
        /// <summary>
        /// A description of the purpose of the API key.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Indicates whether the API key can be used by clients.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// A list of stages to associate with this API key.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApiKeyStageKey> StageKeys;
        /// <summary>
        /// An array of arbitrary tags (key-value pairs) to associate with the API key.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApiKeyTag> Tags;

        [OutputConstructor]
        private GetApiKeyResult(
            string? aPIKeyId,

            string? customerId,

            string? description,

            bool? enabled,

            ImmutableArray<Outputs.ApiKeyStageKey> stageKeys,

            ImmutableArray<Outputs.ApiKeyTag> tags)
        {
            APIKeyId = aPIKeyId;
            CustomerId = customerId;
            Description = description;
            Enabled = enabled;
            StageKeys = stageKeys;
            Tags = tags;
        }
    }
}
