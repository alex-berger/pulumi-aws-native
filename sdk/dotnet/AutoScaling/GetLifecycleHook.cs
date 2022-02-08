// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScaling
{
    public static class GetLifecycleHook
    {
        /// <summary>
        /// Resource Type definition for AWS::AutoScaling::LifecycleHook
        /// </summary>
        public static Task<GetLifecycleHookResult> InvokeAsync(GetLifecycleHookArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLifecycleHookResult>("aws-native:autoscaling:getLifecycleHook", args ?? new GetLifecycleHookArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::AutoScaling::LifecycleHook
        /// </summary>
        public static Output<GetLifecycleHookResult> Invoke(GetLifecycleHookInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetLifecycleHookResult>("aws-native:autoscaling:getLifecycleHook", args ?? new GetLifecycleHookInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLifecycleHookArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Auto Scaling group for the lifecycle hook.
        /// </summary>
        [Input("autoScalingGroupName", required: true)]
        public string AutoScalingGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the lifecycle hook.
        /// </summary>
        [Input("lifecycleHookName", required: true)]
        public string LifecycleHookName { get; set; } = null!;

        public GetLifecycleHookArgs()
        {
        }
    }

    public sealed class GetLifecycleHookInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Auto Scaling group for the lifecycle hook.
        /// </summary>
        [Input("autoScalingGroupName", required: true)]
        public Input<string> AutoScalingGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the lifecycle hook.
        /// </summary>
        [Input("lifecycleHookName", required: true)]
        public Input<string> LifecycleHookName { get; set; } = null!;

        public GetLifecycleHookInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLifecycleHookResult
    {
        /// <summary>
        /// The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs. The valid values are CONTINUE and ABANDON (default).
        /// </summary>
        public readonly string? DefaultResult;
        /// <summary>
        /// The maximum time, in seconds, that can elapse before the lifecycle hook times out. The range is from 30 to 7200 seconds. The default value is 3600 seconds (1 hour). If the lifecycle hook times out, Amazon EC2 Auto Scaling performs the action that you specified in the DefaultResult property.
        /// </summary>
        public readonly int? HeartbeatTimeout;
        /// <summary>
        /// The instance state to which you want to attach the lifecycle hook.
        /// </summary>
        public readonly string? LifecycleTransition;
        /// <summary>
        /// Additional information that is included any time Amazon EC2 Auto Scaling sends a message to the notification target.
        /// </summary>
        public readonly string? NotificationMetadata;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling uses to notify you when an instance is in the transition state for the lifecycle hook. You can specify an Amazon SQS queue or an Amazon SNS topic. The notification message includes the following information: lifecycle action token, user account ID, Auto Scaling group name, lifecycle hook name, instance ID, lifecycle transition, and notification metadata.
        /// </summary>
        public readonly string? NotificationTargetARN;
        /// <summary>
        /// The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target, for example, an Amazon SNS topic or an Amazon SQS queue.
        /// </summary>
        public readonly string? RoleARN;

        [OutputConstructor]
        private GetLifecycleHookResult(
            string? defaultResult,

            int? heartbeatTimeout,

            string? lifecycleTransition,

            string? notificationMetadata,

            string? notificationTargetARN,

            string? roleARN)
        {
            DefaultResult = defaultResult;
            HeartbeatTimeout = heartbeatTimeout;
            LifecycleTransition = lifecycleTransition;
            NotificationMetadata = notificationMetadata;
            NotificationTargetARN = notificationTargetARN;
            RoleARN = roleARN;
        }
    }
}