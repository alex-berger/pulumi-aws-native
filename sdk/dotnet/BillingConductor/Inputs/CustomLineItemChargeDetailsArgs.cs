// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.BillingConductor.Inputs
{

    public sealed class CustomLineItemChargeDetailsArgs : Pulumi.ResourceArgs
    {
        [Input("flat")]
        public Input<Inputs.CustomLineItemFlatChargeDetailsArgs>? Flat { get; set; }

        [Input("percentage")]
        public Input<Inputs.CustomLineItemPercentageChargeDetailsArgs>? Percentage { get; set; }

        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.BillingConductor.CustomLineItemType> Type { get; set; } = null!;

        public CustomLineItemChargeDetailsArgs()
        {
        }
    }
}
