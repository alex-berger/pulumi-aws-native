// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Redshift.Outputs
{

    /// <summary>
    /// Describes a pause cluster operation. For example, a scheduled action to run the `PauseCluster` API operation.
    /// </summary>
    [OutputType]
    public sealed class ScheduledActionPauseClusterMessage
    {
        public readonly string ClusterIdentifier;

        [OutputConstructor]
        private ScheduledActionPauseClusterMessage(string clusterIdentifier)
        {
            ClusterIdentifier = clusterIdentifier;
        }
    }
}
