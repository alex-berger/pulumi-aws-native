// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScaling.Inputs
{

    public sealed class ScalingPolicyMetricStatArgs : Pulumi.ResourceArgs
    {
        [Input("metric", required: true)]
        public Input<Inputs.ScalingPolicyMetricArgs> Metric { get; set; } = null!;

        [Input("stat", required: true)]
        public Input<string> Stat { get; set; } = null!;

        [Input("unit")]
        public Input<string>? Unit { get; set; }

        public ScalingPolicyMetricStatArgs()
        {
        }
    }
}
