// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    public static class GetInstance
    {
        /// <summary>
        /// Resource Type definition for AWS::EC2::Instance
        /// </summary>
        public static Task<GetInstanceResult> InvokeAsync(GetInstanceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceResult>("aws-native:ec2:getInstance", args ?? new GetInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::EC2::Instance
        /// </summary>
        public static Output<GetInstanceResult> Invoke(GetInstanceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstanceResult>("aws-native:ec2:getInstance", args ?? new GetInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetInstanceArgs()
        {
        }
    }

    public sealed class GetInstanceInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetInstanceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceResult
    {
        public readonly string? AdditionalInfo;
        public readonly string? Affinity;
        public readonly ImmutableArray<Outputs.InstanceBlockDeviceMapping> BlockDeviceMappings;
        public readonly Outputs.InstanceCreditSpecification? CreditSpecification;
        public readonly bool? DisableApiTermination;
        public readonly bool? EbsOptimized;
        public readonly string? HostId;
        public readonly string? IamInstanceProfile;
        public readonly string? Id;
        public readonly string? InstanceInitiatedShutdownBehavior;
        public readonly string? InstanceType;
        public readonly string? KernelId;
        public readonly bool? Monitoring;
        public readonly string? PrivateDnsName;
        public readonly Outputs.InstancePrivateDnsNameOptions? PrivateDnsNameOptions;
        public readonly string? PrivateIp;
        public readonly bool? PropagateTagsToVolumeOnCreation;
        public readonly string? PublicDnsName;
        public readonly string? PublicIp;
        public readonly string? RamdiskId;
        public readonly ImmutableArray<string> SecurityGroupIds;
        public readonly bool? SourceDestCheck;
        public readonly ImmutableArray<Outputs.InstanceSsmAssociation> SsmAssociations;
        public readonly ImmutableArray<Outputs.InstanceTag> Tags;
        public readonly string? Tenancy;
        public readonly string? UserData;
        public readonly ImmutableArray<Outputs.InstanceVolume> Volumes;

        [OutputConstructor]
        private GetInstanceResult(
            string? additionalInfo,

            string? affinity,

            ImmutableArray<Outputs.InstanceBlockDeviceMapping> blockDeviceMappings,

            Outputs.InstanceCreditSpecification? creditSpecification,

            bool? disableApiTermination,

            bool? ebsOptimized,

            string? hostId,

            string? iamInstanceProfile,

            string? id,

            string? instanceInitiatedShutdownBehavior,

            string? instanceType,

            string? kernelId,

            bool? monitoring,

            string? privateDnsName,

            Outputs.InstancePrivateDnsNameOptions? privateDnsNameOptions,

            string? privateIp,

            bool? propagateTagsToVolumeOnCreation,

            string? publicDnsName,

            string? publicIp,

            string? ramdiskId,

            ImmutableArray<string> securityGroupIds,

            bool? sourceDestCheck,

            ImmutableArray<Outputs.InstanceSsmAssociation> ssmAssociations,

            ImmutableArray<Outputs.InstanceTag> tags,

            string? tenancy,

            string? userData,

            ImmutableArray<Outputs.InstanceVolume> volumes)
        {
            AdditionalInfo = additionalInfo;
            Affinity = affinity;
            BlockDeviceMappings = blockDeviceMappings;
            CreditSpecification = creditSpecification;
            DisableApiTermination = disableApiTermination;
            EbsOptimized = ebsOptimized;
            HostId = hostId;
            IamInstanceProfile = iamInstanceProfile;
            Id = id;
            InstanceInitiatedShutdownBehavior = instanceInitiatedShutdownBehavior;
            InstanceType = instanceType;
            KernelId = kernelId;
            Monitoring = monitoring;
            PrivateDnsName = privateDnsName;
            PrivateDnsNameOptions = privateDnsNameOptions;
            PrivateIp = privateIp;
            PropagateTagsToVolumeOnCreation = propagateTagsToVolumeOnCreation;
            PublicDnsName = publicDnsName;
            PublicIp = publicIp;
            RamdiskId = ramdiskId;
            SecurityGroupIds = securityGroupIds;
            SourceDestCheck = sourceDestCheck;
            SsmAssociations = ssmAssociations;
            Tags = tags;
            Tenancy = tenancy;
            UserData = userData;
            Volumes = volumes;
        }
    }
}
