// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    public sealed class LaunchTemplateBlockDeviceMappingArgs : Pulumi.ResourceArgs
    {
        [Input("deviceName")]
        public Input<string>? DeviceName { get; set; }

        [Input("ebs")]
        public Input<Inputs.LaunchTemplateEbsArgs>? Ebs { get; set; }

        [Input("noDevice")]
        public Input<string>? NoDevice { get; set; }

        [Input("virtualName")]
        public Input<string>? VirtualName { get; set; }

        public LaunchTemplateBlockDeviceMappingArgs()
        {
        }
    }
}