// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CE
{
    public static class GetAnomalySubscription
    {
        /// <summary>
        /// AWS Cost Anomaly Detection leverages advanced Machine Learning technologies to identify anomalous spend and root causes, so you can quickly take action. Create subscription to be notified
        /// </summary>
        public static Task<GetAnomalySubscriptionResult> InvokeAsync(GetAnomalySubscriptionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAnomalySubscriptionResult>("aws-native:ce:getAnomalySubscription", args ?? new GetAnomalySubscriptionArgs(), options.WithDefaults());

        /// <summary>
        /// AWS Cost Anomaly Detection leverages advanced Machine Learning technologies to identify anomalous spend and root causes, so you can quickly take action. Create subscription to be notified
        /// </summary>
        public static Output<GetAnomalySubscriptionResult> Invoke(GetAnomalySubscriptionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAnomalySubscriptionResult>("aws-native:ce:getAnomalySubscription", args ?? new GetAnomalySubscriptionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAnomalySubscriptionArgs : Pulumi.InvokeArgs
    {
        [Input("subscriptionArn", required: true)]
        public string SubscriptionArn { get; set; } = null!;

        public GetAnomalySubscriptionArgs()
        {
        }
    }

    public sealed class GetAnomalySubscriptionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("subscriptionArn", required: true)]
        public Input<string> SubscriptionArn { get; set; } = null!;

        public GetAnomalySubscriptionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAnomalySubscriptionResult
    {
        /// <summary>
        /// The accountId
        /// </summary>
        public readonly string? AccountId;
        /// <summary>
        /// The frequency at which anomaly reports are sent over email. 
        /// </summary>
        public readonly Pulumi.AwsNative.CE.AnomalySubscriptionFrequency? Frequency;
        /// <summary>
        /// A list of cost anomaly monitors.
        /// </summary>
        public readonly ImmutableArray<string> MonitorArnList;
        /// <summary>
        /// A list of subscriber
        /// </summary>
        public readonly ImmutableArray<Outputs.AnomalySubscriptionSubscriber> Subscribers;
        public readonly string? SubscriptionArn;
        /// <summary>
        /// The name of the subscription.
        /// </summary>
        public readonly string? SubscriptionName;
        /// <summary>
        /// The dollar value that triggers a notification if the threshold is exceeded. 
        /// </summary>
        public readonly double? Threshold;

        [OutputConstructor]
        private GetAnomalySubscriptionResult(
            string? accountId,

            Pulumi.AwsNative.CE.AnomalySubscriptionFrequency? frequency,

            ImmutableArray<string> monitorArnList,

            ImmutableArray<Outputs.AnomalySubscriptionSubscriber> subscribers,

            string? subscriptionArn,

            string? subscriptionName,

            double? threshold)
        {
            AccountId = accountId;
            Frequency = frequency;
            MonitorArnList = monitorArnList;
            Subscribers = subscribers;
            SubscriptionArn = subscriptionArn;
            SubscriptionName = subscriptionName;
            Threshold = threshold;
        }
    }
}
