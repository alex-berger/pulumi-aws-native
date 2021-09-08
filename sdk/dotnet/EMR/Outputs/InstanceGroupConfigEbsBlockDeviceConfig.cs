// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EMR.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig.html
    /// </summary>
    [OutputType]
    public sealed class InstanceGroupConfigEbsBlockDeviceConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification
        /// </summary>
        public readonly Outputs.InstanceGroupConfigVolumeSpecification VolumeSpecification;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfig-volumesperinstance
        /// </summary>
        public readonly int? VolumesPerInstance;

        [OutputConstructor]
        private InstanceGroupConfigEbsBlockDeviceConfig(
            Outputs.InstanceGroupConfigVolumeSpecification volumeSpecification,

            int? volumesPerInstance)
        {
            VolumeSpecification = volumeSpecification;
            VolumesPerInstance = volumesPerInstance;
        }
    }
}
