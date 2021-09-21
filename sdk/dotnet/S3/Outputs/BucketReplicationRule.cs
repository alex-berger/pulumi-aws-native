// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Outputs
{

    [OutputType]
    public sealed class BucketReplicationRule
    {
        public readonly Outputs.BucketDeleteMarkerReplication? DeleteMarkerReplication;
        public readonly Outputs.BucketReplicationDestination Destination;
        public readonly Outputs.BucketReplicationRuleFilter? Filter;
        public readonly string? Id;
        public readonly string? Prefix;
        public readonly int? Priority;
        public readonly Outputs.BucketSourceSelectionCriteria? SourceSelectionCriteria;
        public readonly string Status;

        [OutputConstructor]
        private BucketReplicationRule(
            Outputs.BucketDeleteMarkerReplication? deleteMarkerReplication,

            Outputs.BucketReplicationDestination destination,

            Outputs.BucketReplicationRuleFilter? filter,

            string? id,

            string? prefix,

            int? priority,

            Outputs.BucketSourceSelectionCriteria? sourceSelectionCriteria,

            string status)
        {
            DeleteMarkerReplication = deleteMarkerReplication;
            Destination = destination;
            Filter = filter;
            Id = id;
            Prefix = prefix;
            Priority = priority;
            SourceSelectionCriteria = sourceSelectionCriteria;
            Status = status;
        }
    }
}