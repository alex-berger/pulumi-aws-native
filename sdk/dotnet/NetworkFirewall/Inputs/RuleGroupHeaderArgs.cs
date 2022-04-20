// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Inputs
{

    public sealed class RuleGroupHeaderArgs : Pulumi.ResourceArgs
    {
        [Input("destination", required: true)]
        public Input<string> Destination { get; set; } = null!;

        [Input("destinationPort", required: true)]
        public Input<string> DestinationPort { get; set; } = null!;

        [Input("direction", required: true)]
        public Input<Pulumi.AwsNative.NetworkFirewall.RuleGroupHeaderDirection> Direction { get; set; } = null!;

        [Input("protocol", required: true)]
        public Input<Pulumi.AwsNative.NetworkFirewall.RuleGroupHeaderProtocol> Protocol { get; set; } = null!;

        [Input("source", required: true)]
        public Input<string> Source { get; set; } = null!;

        [Input("sourcePort", required: true)]
        public Input<string> SourcePort { get; set; } = null!;

        public RuleGroupHeaderArgs()
        {
        }
    }
}
