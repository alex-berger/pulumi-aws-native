// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Budgets.Inputs
{

    public sealed class BudgetsActionSubscriberArgs : Pulumi.ResourceArgs
    {
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.Budgets.BudgetsActionSubscriberType> Type { get; set; } = null!;

        public BudgetsActionSubscriberArgs()
        {
        }
    }
}