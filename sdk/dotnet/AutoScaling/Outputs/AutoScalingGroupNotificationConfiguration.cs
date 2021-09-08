// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScaling.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-notificationconfigurations.html
    /// </summary>
    [OutputType]
    public sealed class AutoScalingGroupNotificationConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-notificationconfigurations.html#cfn-as-group-notificationconfigurations-notificationtypes
        /// </summary>
        public readonly ImmutableArray<string> NotificationTypes;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-notificationconfigurations.html#cfn-autoscaling-autoscalinggroup-notificationconfigurations-topicarn
        /// </summary>
        public readonly string TopicARN;

        [OutputConstructor]
        private AutoScalingGroupNotificationConfiguration(
            ImmutableArray<string> notificationTypes,

            string topicARN)
        {
            NotificationTypes = notificationTypes;
            TopicARN = topicARN;
        }
    }
}
