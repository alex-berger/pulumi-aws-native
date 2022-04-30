// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScaling.Inputs
{

    public sealed class ScalingPolicyMetricArgs : Pulumi.ResourceArgs
    {
        [Input("dimensions")]
        private InputList<Inputs.ScalingPolicyMetricDimensionArgs>? _dimensions;
        public InputList<Inputs.ScalingPolicyMetricDimensionArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.ScalingPolicyMetricDimensionArgs>());
            set => _dimensions = value;
        }

        [Input("metricName", required: true)]
        public Input<string> MetricName { get; set; } = null!;

        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        public ScalingPolicyMetricArgs()
        {
        }
    }
}
