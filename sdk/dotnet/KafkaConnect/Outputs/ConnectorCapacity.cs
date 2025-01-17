// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KafkaConnect.Outputs
{

    /// <summary>
    /// Information about the capacity allocated to the connector.
    /// </summary>
    [OutputType]
    public sealed class ConnectorCapacity
    {
        public readonly Outputs.ConnectorAutoScaling? AutoScaling;
        public readonly Outputs.ConnectorProvisionedCapacity? ProvisionedCapacity;

        [OutputConstructor]
        private ConnectorCapacity(
            Outputs.ConnectorAutoScaling? autoScaling,

            Outputs.ConnectorProvisionedCapacity? provisionedCapacity)
        {
            AutoScaling = autoScaling;
            ProvisionedCapacity = provisionedCapacity;
        }
    }
}
