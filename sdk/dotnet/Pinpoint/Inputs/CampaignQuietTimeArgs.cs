// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Inputs
{

    public sealed class CampaignQuietTimeArgs : Pulumi.ResourceArgs
    {
        [Input("end", required: true)]
        public Input<string> End { get; set; } = null!;

        [Input("start", required: true)]
        public Input<string> Start { get; set; } = null!;

        public CampaignQuietTimeArgs()
        {
        }
    }
}
