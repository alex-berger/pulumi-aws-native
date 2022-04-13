// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// The output configuration for monitoring jobs.
    /// </summary>
    [OutputType]
    public sealed class MonitoringScheduleMonitoringOutputConfig
    {
        /// <summary>
        /// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
        /// </summary>
        public readonly string? KmsKeyId;
        /// <summary>
        /// Monitoring outputs for monitoring jobs. This is where the output of the periodic monitoring jobs is uploaded.
        /// </summary>
        public readonly ImmutableArray<Outputs.MonitoringScheduleMonitoringOutput> MonitoringOutputs;

        [OutputConstructor]
        private MonitoringScheduleMonitoringOutputConfig(
            string? kmsKeyId,

            ImmutableArray<Outputs.MonitoringScheduleMonitoringOutput> monitoringOutputs)
        {
            KmsKeyId = kmsKeyId;
            MonitoringOutputs = monitoringOutputs;
        }
    }
}
