// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ImageBuilder.Inputs
{

    /// <summary>
    /// The logging configuration of the infrastructure configuration.
    /// </summary>
    public sealed class InfrastructureConfigurationLoggingArgs : Pulumi.ResourceArgs
    {
        [Input("s3Logs")]
        public Input<Inputs.InfrastructureConfigurationS3LogsArgs>? S3Logs { get; set; }

        public InfrastructureConfigurationLoggingArgs()
        {
        }
    }
}
