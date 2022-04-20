// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Outputs
{

    [OutputType]
    public sealed class LaunchTemplatePlacement
    {
        public readonly string? Affinity;
        public readonly string? AvailabilityZone;
        public readonly string? GroupName;
        public readonly string? HostId;
        public readonly string? HostResourceGroupArn;
        public readonly int? PartitionNumber;
        public readonly string? SpreadDomain;
        public readonly string? Tenancy;

        [OutputConstructor]
        private LaunchTemplatePlacement(
            string? affinity,

            string? availabilityZone,

            string? groupName,

            string? hostId,

            string? hostResourceGroupArn,

            int? partitionNumber,

            string? spreadDomain,

            string? tenancy)
        {
            Affinity = affinity;
            AvailabilityZone = availabilityZone;
            GroupName = groupName;
            HostId = hostId;
            HostResourceGroupArn = hostResourceGroupArn;
            PartitionNumber = partitionNumber;
            SpreadDomain = spreadDomain;
            Tenancy = tenancy;
        }
    }
}
