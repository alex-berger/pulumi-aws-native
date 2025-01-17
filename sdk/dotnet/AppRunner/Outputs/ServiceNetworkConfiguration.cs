// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppRunner.Outputs
{

    /// <summary>
    /// Network configuration
    /// </summary>
    [OutputType]
    public sealed class ServiceNetworkConfiguration
    {
        public readonly Outputs.ServiceEgressConfiguration EgressConfiguration;

        [OutputConstructor]
        private ServiceNetworkConfiguration(Outputs.ServiceEgressConfiguration egressConfiguration)
        {
            EgressConfiguration = egressConfiguration;
        }
    }
}
