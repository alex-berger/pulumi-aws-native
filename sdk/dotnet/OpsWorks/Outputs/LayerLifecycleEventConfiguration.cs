// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpsWorks.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration.html
    /// </summary>
    [OutputType]
    public sealed class LayerLifecycleEventConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration.html#cfn-opsworks-layer-lifecycleconfiguration-shutdowneventconfiguration
        /// </summary>
        public readonly Outputs.LayerShutdownEventConfiguration? ShutdownEventConfiguration;

        [OutputConstructor]
        private LayerLifecycleEventConfiguration(Outputs.LayerShutdownEventConfiguration? shutdownEventConfiguration)
        {
            ShutdownEventConfiguration = shutdownEventConfiguration;
        }
    }
}
