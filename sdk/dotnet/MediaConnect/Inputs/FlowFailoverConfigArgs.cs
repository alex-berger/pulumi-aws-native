// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaConnect.Inputs
{

    /// <summary>
    /// The settings for source failover
    /// </summary>
    public sealed class FlowFailoverConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Search window time to look for dash-7 packets
        /// </summary>
        [Input("recoveryWindow")]
        public Input<int>? RecoveryWindow { get; set; }

        [Input("state")]
        public Input<Pulumi.AwsNative.MediaConnect.FlowFailoverConfigState>? State { get; set; }

        public FlowFailoverConfigArgs()
        {
        }
    }
}
