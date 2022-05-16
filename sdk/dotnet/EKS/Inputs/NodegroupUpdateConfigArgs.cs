// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EKS.Inputs
{

    /// <summary>
    /// The node group update configuration.
    /// </summary>
    public sealed class NodegroupUpdateConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum number of nodes unavailable at once during a version update. Nodes will be updated in parallel. This value or maxUnavailablePercentage is required to have a value.The maximum number is 100. 
        /// </summary>
        [Input("maxUnavailable")]
        public Input<double>? MaxUnavailable { get; set; }

        /// <summary>
        /// The maximum percentage of nodes unavailable during a version update. This percentage of nodes will be updated in parallel, up to 100 nodes at once. This value or maxUnavailable is required to have a value.
        /// </summary>
        [Input("maxUnavailablePercentage")]
        public Input<double>? MaxUnavailablePercentage { get; set; }

        public NodegroupUpdateConfigArgs()
        {
        }
    }
}