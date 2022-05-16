// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    /// <summary>
    /// Configurations of destination connector.
    /// </summary>
    public sealed class FlowDestinationFlowConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of destination connector profile
        /// </summary>
        [Input("connectorProfileName")]
        public Input<string>? ConnectorProfileName { get; set; }

        /// <summary>
        /// Destination connector type
        /// </summary>
        [Input("connectorType", required: true)]
        public Input<Pulumi.AwsNative.AppFlow.FlowConnectorType> ConnectorType { get; set; } = null!;

        /// <summary>
        /// Destination connector details
        /// </summary>
        [Input("destinationConnectorProperties", required: true)]
        public Input<Inputs.FlowDestinationConnectorPropertiesArgs> DestinationConnectorProperties { get; set; } = null!;

        public FlowDestinationFlowConfigArgs()
        {
        }
    }
}