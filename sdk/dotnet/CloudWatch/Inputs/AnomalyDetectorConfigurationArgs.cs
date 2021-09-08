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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-configuration.html
    /// </summary>
    public sealed class AnomalyDetectorConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("excludedTimeRanges")]
        private InputList<Inputs.AnomalyDetectorRangeArgs>? _excludedTimeRanges;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-configuration.html#cfn-cloudwatch-anomalydetector-configuration-excludedtimeranges
        /// </summary>
        public InputList<Inputs.AnomalyDetectorRangeArgs> ExcludedTimeRanges
        {
            get => _excludedTimeRanges ?? (_excludedTimeRanges = new InputList<Inputs.AnomalyDetectorRangeArgs>());
            set => _excludedTimeRanges = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-configuration.html#cfn-cloudwatch-anomalydetector-configuration-metrictimezone
        /// </summary>
        [Input("metricTimeZone")]
        public Input<string>? MetricTimeZone { get; set; }

        public AnomalyDetectorConfigurationArgs()
        {
        }
    }
}
