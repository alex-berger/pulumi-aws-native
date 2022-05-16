// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScaling.Outputs
{

    [OutputType]
    public sealed class ScalingPolicyMetricStat
    {
        public readonly Outputs.ScalingPolicyMetric Metric;
        public readonly string Stat;
        public readonly string? Unit;

        [OutputConstructor]
        private ScalingPolicyMetricStat(
            Outputs.ScalingPolicyMetric metric,

            string stat,

            string? unit)
        {
            Metric = metric;
            Stat = stat;
            Unit = unit;
        }
    }
}