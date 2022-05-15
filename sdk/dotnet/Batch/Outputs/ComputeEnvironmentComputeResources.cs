// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Outputs
{

    [OutputType]
    public sealed class ComputeEnvironmentComputeResources
    {
        public readonly string? AllocationStrategy;
        public readonly int? BidPercentage;
        public readonly int? DesiredvCpus;
        public readonly ImmutableArray<Outputs.ComputeEnvironmentEc2ConfigurationObject> Ec2Configuration;
        public readonly string? Ec2KeyPair;
        public readonly string? ImageId;
        public readonly string? InstanceRole;
        public readonly ImmutableArray<string> InstanceTypes;
        public readonly Outputs.ComputeEnvironmentLaunchTemplateSpecification? LaunchTemplate;
        public readonly int MaxvCpus;
        public readonly int? MinvCpus;
        public readonly string? PlacementGroup;
        public readonly ImmutableArray<string> SecurityGroupIds;
        public readonly string? SpotIamFleetRole;
        public readonly ImmutableArray<string> Subnets;
        /// <summary>
        /// A key-value pair to associate with a resource.
        /// </summary>
        public readonly object? Tags;
        public readonly string Type;
        public readonly bool? UpdateToLatestImageVersion;

        [OutputConstructor]
        private ComputeEnvironmentComputeResources(
            string? allocationStrategy,

            int? bidPercentage,

            int? desiredvCpus,

            ImmutableArray<Outputs.ComputeEnvironmentEc2ConfigurationObject> ec2Configuration,

            string? ec2KeyPair,

            string? imageId,

            string? instanceRole,

            ImmutableArray<string> instanceTypes,

            Outputs.ComputeEnvironmentLaunchTemplateSpecification? launchTemplate,

            int maxvCpus,

            int? minvCpus,

            string? placementGroup,

            ImmutableArray<string> securityGroupIds,

            string? spotIamFleetRole,

            ImmutableArray<string> subnets,

            object? tags,

            string type,

            bool? updateToLatestImageVersion)
        {
            AllocationStrategy = allocationStrategy;
            BidPercentage = bidPercentage;
            DesiredvCpus = desiredvCpus;
            Ec2Configuration = ec2Configuration;
            Ec2KeyPair = ec2KeyPair;
            ImageId = imageId;
            InstanceRole = instanceRole;
            InstanceTypes = instanceTypes;
            LaunchTemplate = launchTemplate;
            MaxvCpus = maxvCpus;
            MinvCpus = minvCpus;
            PlacementGroup = placementGroup;
            SecurityGroupIds = securityGroupIds;
            SpotIamFleetRole = spotIamFleetRole;
            Subnets = subnets;
            Tags = tags;
            Type = type;
            UpdateToLatestImageVersion = updateToLatestImageVersion;
        }
    }
}
