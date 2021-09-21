// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Budgets.Outputs
{

    [OutputType]
    public sealed class BudgetSubscriber
    {
        public readonly string Address;
        public readonly string SubscriptionType;

        [OutputConstructor]
        private BudgetSubscriber(
            string address,

            string subscriptionType)
        {
            Address = address;
            SubscriptionType = subscriptionType;
        }
    }
}