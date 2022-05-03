// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CE
{
    /// <summary>
    /// AWS Cost Anomaly Detection leverages advanced Machine Learning technologies to identify anomalous spend and root causes, so you can quickly take action. Create subscription to be notified
    /// </summary>
    [Obsolete(@"AnomalySubscription is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:ce:AnomalySubscription")]
    public partial class AnomalySubscription : Pulumi.CustomResource
    {
        /// <summary>
        /// The accountId
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// The frequency at which anomaly reports are sent over email. 
        /// </summary>
        [Output("frequency")]
        public Output<Pulumi.AwsNative.CE.AnomalySubscriptionFrequency> Frequency { get; private set; } = null!;

        /// <summary>
        /// A list of cost anomaly monitors.
        /// </summary>
        [Output("monitorArnList")]
        public Output<ImmutableArray<string>> MonitorArnList { get; private set; } = null!;

        /// <summary>
        /// Tags to assign to subscription.
        /// </summary>
        [Output("resourceTags")]
        public Output<ImmutableArray<Outputs.AnomalySubscriptionResourceTag>> ResourceTags { get; private set; } = null!;

        /// <summary>
        /// A list of subscriber
        /// </summary>
        [Output("subscribers")]
        public Output<ImmutableArray<Outputs.AnomalySubscriptionSubscriber>> Subscribers { get; private set; } = null!;

        [Output("subscriptionArn")]
        public Output<string> SubscriptionArn { get; private set; } = null!;

        /// <summary>
        /// The name of the subscription.
        /// </summary>
        [Output("subscriptionName")]
        public Output<string> SubscriptionName { get; private set; } = null!;

        /// <summary>
        /// The dollar value that triggers a notification if the threshold is exceeded. 
        /// </summary>
        [Output("threshold")]
        public Output<double> Threshold { get; private set; } = null!;


        /// <summary>
        /// Create a AnomalySubscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AnomalySubscription(string name, AnomalySubscriptionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ce:AnomalySubscription", name, args ?? new AnomalySubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AnomalySubscription(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ce:AnomalySubscription", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AnomalySubscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AnomalySubscription Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AnomalySubscription(name, id, options);
        }
    }

    public sealed class AnomalySubscriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The frequency at which anomaly reports are sent over email. 
        /// </summary>
        [Input("frequency", required: true)]
        public Input<Pulumi.AwsNative.CE.AnomalySubscriptionFrequency> Frequency { get; set; } = null!;

        [Input("monitorArnList", required: true)]
        private InputList<string>? _monitorArnList;

        /// <summary>
        /// A list of cost anomaly monitors.
        /// </summary>
        public InputList<string> MonitorArnList
        {
            get => _monitorArnList ?? (_monitorArnList = new InputList<string>());
            set => _monitorArnList = value;
        }

        [Input("resourceTags")]
        private InputList<Inputs.AnomalySubscriptionResourceTagArgs>? _resourceTags;

        /// <summary>
        /// Tags to assign to subscription.
        /// </summary>
        public InputList<Inputs.AnomalySubscriptionResourceTagArgs> ResourceTags
        {
            get => _resourceTags ?? (_resourceTags = new InputList<Inputs.AnomalySubscriptionResourceTagArgs>());
            set => _resourceTags = value;
        }

        [Input("subscribers", required: true)]
        private InputList<Inputs.AnomalySubscriptionSubscriberArgs>? _subscribers;

        /// <summary>
        /// A list of subscriber
        /// </summary>
        public InputList<Inputs.AnomalySubscriptionSubscriberArgs> Subscribers
        {
            get => _subscribers ?? (_subscribers = new InputList<Inputs.AnomalySubscriptionSubscriberArgs>());
            set => _subscribers = value;
        }

        /// <summary>
        /// The name of the subscription.
        /// </summary>
        [Input("subscriptionName", required: true)]
        public Input<string> SubscriptionName { get; set; } = null!;

        /// <summary>
        /// The dollar value that triggers a notification if the threshold is exceeded. 
        /// </summary>
        [Input("threshold", required: true)]
        public Input<double> Threshold { get; set; } = null!;

        public AnomalySubscriptionArgs()
        {
        }
    }
}
