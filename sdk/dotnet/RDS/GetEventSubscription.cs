// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RDS
{
    public static class GetEventSubscription
    {
        /// <summary>
        /// Resource Type definition for AWS::RDS::EventSubscription
        /// </summary>
        public static Task<GetEventSubscriptionResult> InvokeAsync(GetEventSubscriptionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEventSubscriptionResult>("aws-native:rds:getEventSubscription", args ?? new GetEventSubscriptionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::RDS::EventSubscription
        /// </summary>
        public static Output<GetEventSubscriptionResult> Invoke(GetEventSubscriptionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEventSubscriptionResult>("aws-native:rds:getEventSubscription", args ?? new GetEventSubscriptionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEventSubscriptionArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetEventSubscriptionArgs()
        {
        }
    }

    public sealed class GetEventSubscriptionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetEventSubscriptionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEventSubscriptionResult
    {
        public readonly bool? Enabled;
        public readonly ImmutableArray<string> EventCategories;
        public readonly string? Id;
        public readonly ImmutableArray<string> SourceIds;
        public readonly string? SourceType;

        [OutputConstructor]
        private GetEventSubscriptionResult(
            bool? enabled,

            ImmutableArray<string> eventCategories,

            string? id,

            ImmutableArray<string> sourceIds,

            string? sourceType)
        {
            Enabled = enabled;
            EventCategories = eventCategories;
            Id = id;
            SourceIds = sourceIds;
            SourceType = sourceType;
        }
    }
}
