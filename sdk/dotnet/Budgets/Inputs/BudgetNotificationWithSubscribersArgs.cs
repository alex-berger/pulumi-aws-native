// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Budgets.Inputs
{

    public sealed class BudgetNotificationWithSubscribersArgs : Pulumi.ResourceArgs
    {
        [Input("notification", required: true)]
        public Input<Inputs.BudgetNotificationArgs> Notification { get; set; } = null!;

        [Input("subscribers", required: true)]
        private InputList<Inputs.BudgetSubscriberArgs>? _subscribers;
        public InputList<Inputs.BudgetSubscriberArgs> Subscribers
        {
            get => _subscribers ?? (_subscribers = new InputList<Inputs.BudgetSubscriberArgs>());
            set => _subscribers = value;
        }

        public BudgetNotificationWithSubscribersArgs()
        {
        }
    }
}
