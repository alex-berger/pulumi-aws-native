// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KafkaConnect.Inputs
{

    /// <summary>
    /// Details about a fixed capacity allocated to a connector.
    /// </summary>
    public sealed class ConnectorProvisionedCapacityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies how many MSK Connect Units (MCU) are allocated to the connector.
        /// </summary>
        [Input("mcuCount")]
        public Input<int>? McuCount { get; set; }

        /// <summary>
        /// Number of workers for a connector.
        /// </summary>
        [Input("workerCount", required: true)]
        public Input<int> WorkerCount { get; set; } = null!;

        public ConnectorProvisionedCapacityArgs()
        {
        }
    }
}
