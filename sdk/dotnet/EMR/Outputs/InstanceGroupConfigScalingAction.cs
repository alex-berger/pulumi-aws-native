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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingaction.html
    /// </summary>
    [OutputType]
    public sealed class InstanceGroupConfigScalingAction
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingaction.html#cfn-elasticmapreduce-instancegroupconfig-scalingaction-market
        /// </summary>
        public readonly string? Market;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingaction.html#cfn-elasticmapreduce-instancegroupconfig-scalingaction-simplescalingpolicyconfiguration
        /// </summary>
        public readonly Outputs.InstanceGroupConfigSimpleScalingPolicyConfiguration SimpleScalingPolicyConfiguration;

        [OutputConstructor]
        private InstanceGroupConfigScalingAction(
            string? market,

            Outputs.InstanceGroupConfigSimpleScalingPolicyConfiguration simpleScalingPolicyConfiguration)
        {
            Market = market;
            SimpleScalingPolicyConfiguration = simpleScalingPolicyConfiguration;
        }
    }
}
