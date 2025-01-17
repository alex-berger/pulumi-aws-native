// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GuardDuty.Inputs
{

    public sealed class DetectorCFNDataSourceConfigurationsArgs : Pulumi.ResourceArgs
    {
        [Input("kubernetes")]
        public Input<Inputs.DetectorCFNKubernetesConfigurationArgs>? Kubernetes { get; set; }

        [Input("s3Logs")]
        public Input<Inputs.DetectorCFNS3LogsConfigurationArgs>? S3Logs { get; set; }

        public DetectorCFNDataSourceConfigurationsArgs()
        {
        }
    }
}
