// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScaling.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-metricscollection.html
    /// </summary>
    public sealed class AutoScalingGroupMetricsCollectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-metricscollection.html#cfn-as-metricscollection-granularity
        /// </summary>
        [Input("granularity", required: true)]
        public Input<string> Granularity { get; set; } = null!;

        [Input("metrics")]
        private InputList<string>? _metrics;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-metricscollection.html#cfn-as-metricscollection-metrics
        /// </summary>
        public InputList<string> Metrics
        {
            get => _metrics ?? (_metrics = new InputList<string>());
            set => _metrics = value;
        }

        public AutoScalingGroupMetricsCollectionArgs()
        {
        }
    }
}