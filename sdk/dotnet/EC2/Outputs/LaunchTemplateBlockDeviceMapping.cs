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
    public sealed class LaunchTemplateBlockDeviceMapping
    {
        public readonly string? DeviceName;
        public readonly Outputs.LaunchTemplateEbs? Ebs;
        public readonly string? NoDevice;
        public readonly string? VirtualName;

        [OutputConstructor]
        private LaunchTemplateBlockDeviceMapping(
            string? deviceName,

            Outputs.LaunchTemplateEbs? ebs,

            string? noDevice,

            string? virtualName)
        {
            DeviceName = deviceName;
            Ebs = ebs;
            NoDevice = noDevice;
            VirtualName = virtualName;
        }
    }
}
