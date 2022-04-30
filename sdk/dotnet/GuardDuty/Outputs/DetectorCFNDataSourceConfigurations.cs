// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GuardDuty.Outputs
{

    [OutputType]
    public sealed class DetectorCFNDataSourceConfigurations
    {
        public readonly Outputs.DetectorCFNKubernetesConfiguration? Kubernetes;
        public readonly Outputs.DetectorCFNS3LogsConfiguration? S3Logs;

        [OutputConstructor]
        private DetectorCFNDataSourceConfigurations(
            Outputs.DetectorCFNKubernetesConfiguration? kubernetes,

            Outputs.DetectorCFNS3LogsConfiguration? s3Logs)
        {
            Kubernetes = kubernetes;
            S3Logs = s3Logs;
        }
    }
}
