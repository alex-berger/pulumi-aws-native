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
    public sealed class LaunchTemplateData
    {
        public readonly ImmutableArray<Outputs.LaunchTemplateBlockDeviceMapping> BlockDeviceMappings;
        public readonly Outputs.LaunchTemplateCapacityReservationSpecification? CapacityReservationSpecification;
        public readonly Outputs.LaunchTemplateCpuOptions? CpuOptions;
        public readonly Outputs.LaunchTemplateCreditSpecification? CreditSpecification;
        public readonly bool? DisableApiTermination;
        public readonly bool? EbsOptimized;
        public readonly ImmutableArray<Outputs.LaunchTemplateElasticGpuSpecification> ElasticGpuSpecifications;
        public readonly ImmutableArray<Outputs.LaunchTemplateElasticInferenceAccelerator> ElasticInferenceAccelerators;
        public readonly Outputs.LaunchTemplateEnclaveOptions? EnclaveOptions;
        public readonly Outputs.LaunchTemplateHibernationOptions? HibernationOptions;
        public readonly Outputs.LaunchTemplateIamInstanceProfile? IamInstanceProfile;
        public readonly string? ImageId;
        public readonly string? InstanceInitiatedShutdownBehavior;
        public readonly Outputs.LaunchTemplateInstanceMarketOptions? InstanceMarketOptions;
        public readonly Outputs.LaunchTemplateInstanceRequirements? InstanceRequirements;
        public readonly string? InstanceType;
        public readonly string? KernelId;
        public readonly string? KeyName;
        public readonly ImmutableArray<Outputs.LaunchTemplateLicenseSpecification> LicenseSpecifications;
        public readonly Outputs.LaunchTemplateMaintenanceOptions? MaintenanceOptions;
        public readonly Outputs.LaunchTemplateMetadataOptions? MetadataOptions;
        public readonly Outputs.LaunchTemplateMonitoring? Monitoring;
        public readonly ImmutableArray<Outputs.LaunchTemplateNetworkInterface> NetworkInterfaces;
        public readonly Outputs.LaunchTemplatePlacement? Placement;
        public readonly Outputs.LaunchTemplatePrivateDnsNameOptions? PrivateDnsNameOptions;
        public readonly string? RamDiskId;
        public readonly ImmutableArray<string> SecurityGroupIds;
        public readonly ImmutableArray<string> SecurityGroups;
        public readonly ImmutableArray<Outputs.LaunchTemplateTagSpecification> TagSpecifications;
        public readonly string? UserData;

        [OutputConstructor]
        private LaunchTemplateData(
            ImmutableArray<Outputs.LaunchTemplateBlockDeviceMapping> blockDeviceMappings,

            Outputs.LaunchTemplateCapacityReservationSpecification? capacityReservationSpecification,

            Outputs.LaunchTemplateCpuOptions? cpuOptions,

            Outputs.LaunchTemplateCreditSpecification? creditSpecification,

            bool? disableApiTermination,

            bool? ebsOptimized,

            ImmutableArray<Outputs.LaunchTemplateElasticGpuSpecification> elasticGpuSpecifications,

            ImmutableArray<Outputs.LaunchTemplateElasticInferenceAccelerator> elasticInferenceAccelerators,

            Outputs.LaunchTemplateEnclaveOptions? enclaveOptions,

            Outputs.LaunchTemplateHibernationOptions? hibernationOptions,

            Outputs.LaunchTemplateIamInstanceProfile? iamInstanceProfile,

            string? imageId,

            string? instanceInitiatedShutdownBehavior,

            Outputs.LaunchTemplateInstanceMarketOptions? instanceMarketOptions,

            Outputs.LaunchTemplateInstanceRequirements? instanceRequirements,

            string? instanceType,

            string? kernelId,

            string? keyName,

            ImmutableArray<Outputs.LaunchTemplateLicenseSpecification> licenseSpecifications,

            Outputs.LaunchTemplateMaintenanceOptions? maintenanceOptions,

            Outputs.LaunchTemplateMetadataOptions? metadataOptions,

            Outputs.LaunchTemplateMonitoring? monitoring,

            ImmutableArray<Outputs.LaunchTemplateNetworkInterface> networkInterfaces,

            Outputs.LaunchTemplatePlacement? placement,

            Outputs.LaunchTemplatePrivateDnsNameOptions? privateDnsNameOptions,

            string? ramDiskId,

            ImmutableArray<string> securityGroupIds,

            ImmutableArray<string> securityGroups,

            ImmutableArray<Outputs.LaunchTemplateTagSpecification> tagSpecifications,

            string? userData)
        {
            BlockDeviceMappings = blockDeviceMappings;
            CapacityReservationSpecification = capacityReservationSpecification;
            CpuOptions = cpuOptions;
            CreditSpecification = creditSpecification;
            DisableApiTermination = disableApiTermination;
            EbsOptimized = ebsOptimized;
            ElasticGpuSpecifications = elasticGpuSpecifications;
            ElasticInferenceAccelerators = elasticInferenceAccelerators;
            EnclaveOptions = enclaveOptions;
            HibernationOptions = hibernationOptions;
            IamInstanceProfile = iamInstanceProfile;
            ImageId = imageId;
            InstanceInitiatedShutdownBehavior = instanceInitiatedShutdownBehavior;
            InstanceMarketOptions = instanceMarketOptions;
            InstanceRequirements = instanceRequirements;
            InstanceType = instanceType;
            KernelId = kernelId;
            KeyName = keyName;
            LicenseSpecifications = licenseSpecifications;
            MaintenanceOptions = maintenanceOptions;
            MetadataOptions = metadataOptions;
            Monitoring = monitoring;
            NetworkInterfaces = networkInterfaces;
            Placement = placement;
            PrivateDnsNameOptions = privateDnsNameOptions;
            RamDiskId = ramDiskId;
            SecurityGroupIds = securityGroupIds;
            SecurityGroups = securityGroups;
            TagSpecifications = tagSpecifications;
            UserData = userData;
        }
    }
}