// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudWatch.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-metricstat.html
    /// </summary>
    public sealed class AlarmMetricStatArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-metricstat.html#cfn-cloudwatch-alarm-metricstat-metric
        /// </summary>
        [Input("metric", required: true)]
        public Input<Inputs.AlarmMetricArgs> Metric { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-metricstat.html#cfn-cloudwatch-alarm-metricstat-period
        /// </summary>
        [Input("period", required: true)]
        public Input<int> Period { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-metricstat.html#cfn-cloudwatch-alarm-metricstat-stat
        /// </summary>
        [Input("stat", required: true)]
        public Input<string> Stat { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-metricstat.html#cfn-cloudwatch-alarm-metricstat-unit
        /// </summary>
        [Input("unit")]
        public Input<string>? Unit { get; set; }

        public AlarmMetricStatArgs()
        {
        }
    }
}