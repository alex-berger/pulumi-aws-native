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
    public sealed class ScalingPolicyPredictiveScalingCustomizedCapacityMetric
    {
        public readonly ImmutableArray<Outputs.ScalingPolicyMetricDataQuery> MetricDataQueries;

        [OutputConstructor]
        private ScalingPolicyPredictiveScalingCustomizedCapacityMetric(ImmutableArray<Outputs.ScalingPolicyMetricDataQuery> metricDataQueries)
        {
            MetricDataQueries = metricDataQueries;
        }
    }
}
