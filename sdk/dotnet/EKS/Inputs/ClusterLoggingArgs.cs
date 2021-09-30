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
    /// Enable exporting the Kubernetes control plane logs for your cluster to CloudWatch Logs based on log types. By default, cluster control plane logs aren't exported to CloudWatch Logs.
    /// </summary>
    public sealed class ClusterLoggingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cluster control plane logging configuration for your cluster. 
        /// </summary>
        [Input("clusterLogging")]
        public Input<Inputs.ClusterClusterLoggingArgs>? ClusterLoggingValue { get; set; }

        public ClusterLoggingArgs()
        {
        }
    }
}
