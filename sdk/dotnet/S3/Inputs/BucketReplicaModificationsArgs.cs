// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    public sealed class BucketReplicaModificationsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether Amazon S3 replicates modifications on replicas.
        /// </summary>
        [Input("status", required: true)]
        public Input<Pulumi.AwsNative.S3.BucketReplicaModificationsStatus> Status { get; set; } = null!;

        public BucketReplicaModificationsArgs()
        {
        }
    }
}