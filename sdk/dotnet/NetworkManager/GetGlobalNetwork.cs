// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkManager
{
    public static class GetGlobalNetwork
    {
        /// <summary>
        /// The AWS::NetworkManager::GlobalNetwork type specifies a global network of the user's account
        /// </summary>
        public static Task<GetGlobalNetworkResult> InvokeAsync(GetGlobalNetworkArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGlobalNetworkResult>("aws-native:networkmanager:getGlobalNetwork", args ?? new GetGlobalNetworkArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::NetworkManager::GlobalNetwork type specifies a global network of the user's account
        /// </summary>
        public static Output<GetGlobalNetworkResult> Invoke(GetGlobalNetworkInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGlobalNetworkResult>("aws-native:networkmanager:getGlobalNetwork", args ?? new GetGlobalNetworkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGlobalNetworkArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the global network.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetGlobalNetworkArgs()
        {
        }
    }

    public sealed class GetGlobalNetworkInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the global network.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetGlobalNetworkInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGlobalNetworkResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the global network.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The description of the global network.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The ID of the global network.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The tags for the global network.
        /// </summary>
        public readonly ImmutableArray<Outputs.GlobalNetworkTag> Tags;

        [OutputConstructor]
        private GetGlobalNetworkResult(
            string? arn,

            string? description,

            string? id,

            ImmutableArray<Outputs.GlobalNetworkTag> tags)
        {
            Arn = arn;
            Description = description;
            Id = id;
            Tags = tags;
        }
    }
}
