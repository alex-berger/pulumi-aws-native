// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    public static class GetNetworkInsightsPath
    {
        /// <summary>
        /// Resource schema for AWS::EC2::NetworkInsightsPath
        /// </summary>
        public static Task<GetNetworkInsightsPathResult> InvokeAsync(GetNetworkInsightsPathArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetworkInsightsPathResult>("aws-native:ec2:getNetworkInsightsPath", args ?? new GetNetworkInsightsPathArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::EC2::NetworkInsightsPath
        /// </summary>
        public static Output<GetNetworkInsightsPathResult> Invoke(GetNetworkInsightsPathInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNetworkInsightsPathResult>("aws-native:ec2:getNetworkInsightsPath", args ?? new GetNetworkInsightsPathInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkInsightsPathArgs : Pulumi.InvokeArgs
    {
        [Input("networkInsightsPathId", required: true)]
        public string NetworkInsightsPathId { get; set; } = null!;

        public GetNetworkInsightsPathArgs()
        {
        }
    }

    public sealed class GetNetworkInsightsPathInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("networkInsightsPathId", required: true)]
        public Input<string> NetworkInsightsPathId { get; set; } = null!;

        public GetNetworkInsightsPathInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNetworkInsightsPathResult
    {
        public readonly string? CreatedDate;
        public readonly string? NetworkInsightsPathArn;
        public readonly string? NetworkInsightsPathId;
        public readonly ImmutableArray<Outputs.NetworkInsightsPathTag> Tags;

        [OutputConstructor]
        private GetNetworkInsightsPathResult(
            string? createdDate,

            string? networkInsightsPathArn,

            string? networkInsightsPathId,

            ImmutableArray<Outputs.NetworkInsightsPathTag> tags)
        {
            CreatedDate = createdDate;
            NetworkInsightsPathArn = networkInsightsPathArn;
            NetworkInsightsPathId = networkInsightsPathId;
            Tags = tags;
        }
    }
}
