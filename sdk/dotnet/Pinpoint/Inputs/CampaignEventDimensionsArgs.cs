// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Inputs
{

    public sealed class CampaignEventDimensionsArgs : Pulumi.ResourceArgs
    {
        [Input("attributes")]
        public Input<object>? Attributes { get; set; }

        [Input("eventType")]
        public Input<Inputs.CampaignSetDimensionArgs>? EventType { get; set; }

        [Input("metrics")]
        public Input<object>? Metrics { get; set; }

        public CampaignEventDimensionsArgs()
        {
        }
    }
}
