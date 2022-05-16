// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Configuration
{
    public static class GetAggregationAuthorization
    {
        /// <summary>
        /// Resource Type definition for AWS::Config::AggregationAuthorization
        /// </summary>
        public static Task<GetAggregationAuthorizationResult> InvokeAsync(GetAggregationAuthorizationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAggregationAuthorizationResult>("aws-native:configuration:getAggregationAuthorization", args ?? new GetAggregationAuthorizationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Config::AggregationAuthorization
        /// </summary>
        public static Output<GetAggregationAuthorizationResult> Invoke(GetAggregationAuthorizationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAggregationAuthorizationResult>("aws-native:configuration:getAggregationAuthorization", args ?? new GetAggregationAuthorizationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAggregationAuthorizationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the AggregationAuthorization.
        /// </summary>
        [Input("aggregationAuthorizationArn", required: true)]
        public string AggregationAuthorizationArn { get; set; } = null!;

        public GetAggregationAuthorizationArgs()
        {
        }
    }

    public sealed class GetAggregationAuthorizationInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the AggregationAuthorization.
        /// </summary>
        [Input("aggregationAuthorizationArn", required: true)]
        public Input<string> AggregationAuthorizationArn { get; set; } = null!;

        public GetAggregationAuthorizationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAggregationAuthorizationResult
    {
        /// <summary>
        /// The ARN of the AggregationAuthorization.
        /// </summary>
        public readonly string? AggregationAuthorizationArn;
        /// <summary>
        /// The tags for the AggregationAuthorization.
        /// </summary>
        public readonly ImmutableArray<Outputs.AggregationAuthorizationTag> Tags;

        [OutputConstructor]
        private GetAggregationAuthorizationResult(
            string? aggregationAuthorizationArn,

            ImmutableArray<Outputs.AggregationAuthorizationTag> tags)
        {
            AggregationAuthorizationArn = aggregationAuthorizationArn;
            Tags = tags;
        }
    }
}