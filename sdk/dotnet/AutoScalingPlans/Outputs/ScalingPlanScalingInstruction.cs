// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScalingPlans.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html
    /// </summary>
    [OutputType]
    public sealed class ScalingPlanScalingInstruction
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-customizedloadmetricspecification
        /// </summary>
        public readonly Outputs.ScalingPlanCustomizedLoadMetricSpecification? CustomizedLoadMetricSpecification;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-disabledynamicscaling
        /// </summary>
        public readonly bool? DisableDynamicScaling;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-maxcapacity
        /// </summary>
        public readonly int MaxCapacity;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-mincapacity
        /// </summary>
        public readonly int MinCapacity;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-predefinedloadmetricspecification
        /// </summary>
        public readonly Outputs.ScalingPlanPredefinedLoadMetricSpecification? PredefinedLoadMetricSpecification;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-predictivescalingmaxcapacitybehavior
        /// </summary>
        public readonly string? PredictiveScalingMaxCapacityBehavior;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-predictivescalingmaxcapacitybuffer
        /// </summary>
        public readonly int? PredictiveScalingMaxCapacityBuffer;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-predictivescalingmode
        /// </summary>
        public readonly string? PredictiveScalingMode;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-resourceid
        /// </summary>
        public readonly string ResourceId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-scalabledimension
        /// </summary>
        public readonly string ScalableDimension;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-scalingpolicyupdatebehavior
        /// </summary>
        public readonly string? ScalingPolicyUpdateBehavior;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-scheduledactionbuffertime
        /// </summary>
        public readonly int? ScheduledActionBufferTime;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-servicenamespace
        /// </summary>
        public readonly string ServiceNamespace;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-targettrackingconfigurations
        /// </summary>
        public readonly ImmutableArray<Outputs.ScalingPlanTargetTrackingConfiguration> TargetTrackingConfigurations;

        [OutputConstructor]
        private ScalingPlanScalingInstruction(
            Outputs.ScalingPlanCustomizedLoadMetricSpecification? customizedLoadMetricSpecification,

            bool? disableDynamicScaling,

            int maxCapacity,

            int minCapacity,

            Outputs.ScalingPlanPredefinedLoadMetricSpecification? predefinedLoadMetricSpecification,

            string? predictiveScalingMaxCapacityBehavior,

            int? predictiveScalingMaxCapacityBuffer,

            string? predictiveScalingMode,

            string resourceId,

            string scalableDimension,

            string? scalingPolicyUpdateBehavior,

            int? scheduledActionBufferTime,

            string serviceNamespace,

            ImmutableArray<Outputs.ScalingPlanTargetTrackingConfiguration> targetTrackingConfigurations)
        {
            CustomizedLoadMetricSpecification = customizedLoadMetricSpecification;
            DisableDynamicScaling = disableDynamicScaling;
            MaxCapacity = maxCapacity;
            MinCapacity = minCapacity;
            PredefinedLoadMetricSpecification = predefinedLoadMetricSpecification;
            PredictiveScalingMaxCapacityBehavior = predictiveScalingMaxCapacityBehavior;
            PredictiveScalingMaxCapacityBuffer = predictiveScalingMaxCapacityBuffer;
            PredictiveScalingMode = predictiveScalingMode;
            ResourceId = resourceId;
            ScalableDimension = scalableDimension;
            ScalingPolicyUpdateBehavior = scalingPolicyUpdateBehavior;
            ScheduledActionBufferTime = scheduledActionBufferTime;
            ServiceNamespace = serviceNamespace;
            TargetTrackingConfigurations = targetTrackingConfigurations;
        }
    }
}
