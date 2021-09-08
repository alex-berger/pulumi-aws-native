// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpsWorks.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling.html
    /// </summary>
    [OutputType]
    public sealed class LayerLoadBasedAutoScaling
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling.html#cfn-opsworks-layer-loadbasedautoscaling-downscaling
        /// </summary>
        public readonly Outputs.LayerAutoScalingThresholds? DownScaling;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling.html#cfn-opsworks-layer-loadbasedautoscaling-enable
        /// </summary>
        public readonly bool? Enable;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling.html#cfn-opsworks-layer-loadbasedautoscaling-upscaling
        /// </summary>
        public readonly Outputs.LayerAutoScalingThresholds? UpScaling;

        [OutputConstructor]
        private LayerLoadBasedAutoScaling(
            Outputs.LayerAutoScalingThresholds? downScaling,

            bool? enable,

            Outputs.LayerAutoScalingThresholds? upScaling)
        {
            DownScaling = downScaling;
            Enable = enable;
            UpScaling = upScaling;
        }
    }
}
