// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    public sealed class TrafficMirrorFilterRuleTrafficMirrorPortRangeArgs : Pulumi.ResourceArgs
    {
        [Input("fromPort", required: true)]
        public Input<int> FromPort { get; set; } = null!;

        [Input("toPort", required: true)]
        public Input<int> ToPort { get; set; } = null!;

        public TrafficMirrorFilterRuleTrafficMirrorPortRangeArgs()
        {
        }
    }
}
