// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LookoutMetrics
{
    public static class GetAnomalyDetector
    {
        /// <summary>
        /// An Amazon Lookout for Metrics Detector
        /// </summary>
        public static Task<GetAnomalyDetectorResult> InvokeAsync(GetAnomalyDetectorArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAnomalyDetectorResult>("aws-native:lookoutmetrics:getAnomalyDetector", args ?? new GetAnomalyDetectorArgs(), options.WithDefaults());

        /// <summary>
        /// An Amazon Lookout for Metrics Detector
        /// </summary>
        public static Output<GetAnomalyDetectorResult> Invoke(GetAnomalyDetectorInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAnomalyDetectorResult>("aws-native:lookoutmetrics:getAnomalyDetector", args ?? new GetAnomalyDetectorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAnomalyDetectorArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetAnomalyDetectorArgs()
        {
        }
    }

    public sealed class GetAnomalyDetectorInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetAnomalyDetectorInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAnomalyDetectorResult
    {
        /// <summary>
        /// Configuration options for the AnomalyDetector
        /// </summary>
        public readonly Outputs.AnomalyDetectorConfig? AnomalyDetectorConfig;
        /// <summary>
        /// A description for the AnomalyDetector.
        /// </summary>
        public readonly string? AnomalyDetectorDescription;
        public readonly string? Arn;
        /// <summary>
        /// KMS key used to encrypt the AnomalyDetector data
        /// </summary>
        public readonly string? KmsKeyArn;
        /// <summary>
        /// List of metric sets for anomaly detection
        /// </summary>
        public readonly ImmutableArray<Outputs.AnomalyDetectorMetricSet> MetricSetList;

        [OutputConstructor]
        private GetAnomalyDetectorResult(
            Outputs.AnomalyDetectorConfig? anomalyDetectorConfig,

            string? anomalyDetectorDescription,

            string? arn,

            string? kmsKeyArn,

            ImmutableArray<Outputs.AnomalyDetectorMetricSet> metricSetList)
        {
            AnomalyDetectorConfig = anomalyDetectorConfig;
            AnomalyDetectorDescription = anomalyDetectorDescription;
            Arn = arn;
            KmsKeyArn = kmsKeyArn;
            MetricSetList = metricSetList;
        }
    }
}
