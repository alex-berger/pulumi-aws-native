// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EMR.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration.html
    /// </summary>
    public sealed class InstanceGroupConfigEbsConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("ebsBlockDeviceConfigs")]
        private InputList<Inputs.InstanceGroupConfigEbsBlockDeviceConfigArgs>? _ebsBlockDeviceConfigs;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfigs
        /// </summary>
        public InputList<Inputs.InstanceGroupConfigEbsBlockDeviceConfigArgs> EbsBlockDeviceConfigs
        {
            get => _ebsBlockDeviceConfigs ?? (_ebsBlockDeviceConfigs = new InputList<Inputs.InstanceGroupConfigEbsBlockDeviceConfigArgs>());
            set => _ebsBlockDeviceConfigs = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration.html#cfn-emr-ebsconfiguration-ebsoptimized
        /// </summary>
        [Input("ebsOptimized")]
        public Input<bool>? EbsOptimized { get; set; }

        public InstanceGroupConfigEbsConfigurationArgs()
        {
        }
    }
}
