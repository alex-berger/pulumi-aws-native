# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['LifecycleHookArgs', 'LifecycleHook']

@pulumi.input_type
class LifecycleHookArgs:
    def __init__(__self__, *,
                 auto_scaling_group_name: pulumi.Input[str],
                 lifecycle_transition: pulumi.Input[str],
                 default_result: Optional[pulumi.Input[str]] = None,
                 heartbeat_timeout: Optional[pulumi.Input[int]] = None,
                 lifecycle_hook_name: Optional[pulumi.Input[str]] = None,
                 notification_metadata: Optional[pulumi.Input[str]] = None,
                 notification_target_arn: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LifecycleHook resource.
        :param pulumi.Input[str] auto_scaling_group_name: The name of the Auto Scaling group for the lifecycle hook.
        :param pulumi.Input[str] lifecycle_transition: The instance state to which you want to attach the lifecycle hook.
        :param pulumi.Input[str] default_result: The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs. The valid values are CONTINUE and ABANDON (default).
        :param pulumi.Input[int] heartbeat_timeout: The maximum time, in seconds, that can elapse before the lifecycle hook times out. The range is from 30 to 7200 seconds. The default value is 3600 seconds (1 hour). If the lifecycle hook times out, Amazon EC2 Auto Scaling performs the action that you specified in the DefaultResult property.
        :param pulumi.Input[str] lifecycle_hook_name: The name of the lifecycle hook.
        :param pulumi.Input[str] notification_metadata: Additional information that is included any time Amazon EC2 Auto Scaling sends a message to the notification target.
        :param pulumi.Input[str] notification_target_arn: The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling uses to notify you when an instance is in the transition state for the lifecycle hook. You can specify an Amazon SQS queue or an Amazon SNS topic. The notification message includes the following information: lifecycle action token, user account ID, Auto Scaling group name, lifecycle hook name, instance ID, lifecycle transition, and notification metadata.
        :param pulumi.Input[str] role_arn: The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target, for example, an Amazon SNS topic or an Amazon SQS queue.
        """
        pulumi.set(__self__, "auto_scaling_group_name", auto_scaling_group_name)
        pulumi.set(__self__, "lifecycle_transition", lifecycle_transition)
        if default_result is not None:
            pulumi.set(__self__, "default_result", default_result)
        if heartbeat_timeout is not None:
            pulumi.set(__self__, "heartbeat_timeout", heartbeat_timeout)
        if lifecycle_hook_name is not None:
            pulumi.set(__self__, "lifecycle_hook_name", lifecycle_hook_name)
        if notification_metadata is not None:
            pulumi.set(__self__, "notification_metadata", notification_metadata)
        if notification_target_arn is not None:
            pulumi.set(__self__, "notification_target_arn", notification_target_arn)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)

    @property
    @pulumi.getter(name="autoScalingGroupName")
    def auto_scaling_group_name(self) -> pulumi.Input[str]:
        """
        The name of the Auto Scaling group for the lifecycle hook.
        """
        return pulumi.get(self, "auto_scaling_group_name")

    @auto_scaling_group_name.setter
    def auto_scaling_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "auto_scaling_group_name", value)

    @property
    @pulumi.getter(name="lifecycleTransition")
    def lifecycle_transition(self) -> pulumi.Input[str]:
        """
        The instance state to which you want to attach the lifecycle hook.
        """
        return pulumi.get(self, "lifecycle_transition")

    @lifecycle_transition.setter
    def lifecycle_transition(self, value: pulumi.Input[str]):
        pulumi.set(self, "lifecycle_transition", value)

    @property
    @pulumi.getter(name="defaultResult")
    def default_result(self) -> Optional[pulumi.Input[str]]:
        """
        The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs. The valid values are CONTINUE and ABANDON (default).
        """
        return pulumi.get(self, "default_result")

    @default_result.setter
    def default_result(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_result", value)

    @property
    @pulumi.getter(name="heartbeatTimeout")
    def heartbeat_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum time, in seconds, that can elapse before the lifecycle hook times out. The range is from 30 to 7200 seconds. The default value is 3600 seconds (1 hour). If the lifecycle hook times out, Amazon EC2 Auto Scaling performs the action that you specified in the DefaultResult property.
        """
        return pulumi.get(self, "heartbeat_timeout")

    @heartbeat_timeout.setter
    def heartbeat_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "heartbeat_timeout", value)

    @property
    @pulumi.getter(name="lifecycleHookName")
    def lifecycle_hook_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the lifecycle hook.
        """
        return pulumi.get(self, "lifecycle_hook_name")

    @lifecycle_hook_name.setter
    def lifecycle_hook_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lifecycle_hook_name", value)

    @property
    @pulumi.getter(name="notificationMetadata")
    def notification_metadata(self) -> Optional[pulumi.Input[str]]:
        """
        Additional information that is included any time Amazon EC2 Auto Scaling sends a message to the notification target.
        """
        return pulumi.get(self, "notification_metadata")

    @notification_metadata.setter
    def notification_metadata(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notification_metadata", value)

    @property
    @pulumi.getter(name="notificationTargetARN")
    def notification_target_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling uses to notify you when an instance is in the transition state for the lifecycle hook. You can specify an Amazon SQS queue or an Amazon SNS topic. The notification message includes the following information: lifecycle action token, user account ID, Auto Scaling group name, lifecycle hook name, instance ID, lifecycle transition, and notification metadata.
        """
        return pulumi.get(self, "notification_target_arn")

    @notification_target_arn.setter
    def notification_target_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notification_target_arn", value)

    @property
    @pulumi.getter(name="roleARN")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target, for example, an Amazon SNS topic or an Amazon SQS queue.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)


class LifecycleHook(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_scaling_group_name: Optional[pulumi.Input[str]] = None,
                 default_result: Optional[pulumi.Input[str]] = None,
                 heartbeat_timeout: Optional[pulumi.Input[int]] = None,
                 lifecycle_hook_name: Optional[pulumi.Input[str]] = None,
                 lifecycle_transition: Optional[pulumi.Input[str]] = None,
                 notification_metadata: Optional[pulumi.Input[str]] = None,
                 notification_target_arn: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::AutoScaling::LifecycleHook

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_scaling_group_name: The name of the Auto Scaling group for the lifecycle hook.
        :param pulumi.Input[str] default_result: The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs. The valid values are CONTINUE and ABANDON (default).
        :param pulumi.Input[int] heartbeat_timeout: The maximum time, in seconds, that can elapse before the lifecycle hook times out. The range is from 30 to 7200 seconds. The default value is 3600 seconds (1 hour). If the lifecycle hook times out, Amazon EC2 Auto Scaling performs the action that you specified in the DefaultResult property.
        :param pulumi.Input[str] lifecycle_hook_name: The name of the lifecycle hook.
        :param pulumi.Input[str] lifecycle_transition: The instance state to which you want to attach the lifecycle hook.
        :param pulumi.Input[str] notification_metadata: Additional information that is included any time Amazon EC2 Auto Scaling sends a message to the notification target.
        :param pulumi.Input[str] notification_target_arn: The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling uses to notify you when an instance is in the transition state for the lifecycle hook. You can specify an Amazon SQS queue or an Amazon SNS topic. The notification message includes the following information: lifecycle action token, user account ID, Auto Scaling group name, lifecycle hook name, instance ID, lifecycle transition, and notification metadata.
        :param pulumi.Input[str] role_arn: The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target, for example, an Amazon SNS topic or an Amazon SQS queue.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LifecycleHookArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::AutoScaling::LifecycleHook

        :param str resource_name: The name of the resource.
        :param LifecycleHookArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LifecycleHookArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_scaling_group_name: Optional[pulumi.Input[str]] = None,
                 default_result: Optional[pulumi.Input[str]] = None,
                 heartbeat_timeout: Optional[pulumi.Input[int]] = None,
                 lifecycle_hook_name: Optional[pulumi.Input[str]] = None,
                 lifecycle_transition: Optional[pulumi.Input[str]] = None,
                 notification_metadata: Optional[pulumi.Input[str]] = None,
                 notification_target_arn: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LifecycleHookArgs.__new__(LifecycleHookArgs)

            if auto_scaling_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'auto_scaling_group_name'")
            __props__.__dict__["auto_scaling_group_name"] = auto_scaling_group_name
            __props__.__dict__["default_result"] = default_result
            __props__.__dict__["heartbeat_timeout"] = heartbeat_timeout
            __props__.__dict__["lifecycle_hook_name"] = lifecycle_hook_name
            if lifecycle_transition is None and not opts.urn:
                raise TypeError("Missing required property 'lifecycle_transition'")
            __props__.__dict__["lifecycle_transition"] = lifecycle_transition
            __props__.__dict__["notification_metadata"] = notification_metadata
            __props__.__dict__["notification_target_arn"] = notification_target_arn
            __props__.__dict__["role_arn"] = role_arn
        super(LifecycleHook, __self__).__init__(
            'aws-native:autoscaling:LifecycleHook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'LifecycleHook':
        """
        Get an existing LifecycleHook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = LifecycleHookArgs.__new__(LifecycleHookArgs)

        __props__.__dict__["auto_scaling_group_name"] = None
        __props__.__dict__["default_result"] = None
        __props__.__dict__["heartbeat_timeout"] = None
        __props__.__dict__["lifecycle_hook_name"] = None
        __props__.__dict__["lifecycle_transition"] = None
        __props__.__dict__["notification_metadata"] = None
        __props__.__dict__["notification_target_arn"] = None
        __props__.__dict__["role_arn"] = None
        return LifecycleHook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoScalingGroupName")
    def auto_scaling_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Auto Scaling group for the lifecycle hook.
        """
        return pulumi.get(self, "auto_scaling_group_name")

    @property
    @pulumi.getter(name="defaultResult")
    def default_result(self) -> pulumi.Output[Optional[str]]:
        """
        The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs. The valid values are CONTINUE and ABANDON (default).
        """
        return pulumi.get(self, "default_result")

    @property
    @pulumi.getter(name="heartbeatTimeout")
    def heartbeat_timeout(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum time, in seconds, that can elapse before the lifecycle hook times out. The range is from 30 to 7200 seconds. The default value is 3600 seconds (1 hour). If the lifecycle hook times out, Amazon EC2 Auto Scaling performs the action that you specified in the DefaultResult property.
        """
        return pulumi.get(self, "heartbeat_timeout")

    @property
    @pulumi.getter(name="lifecycleHookName")
    def lifecycle_hook_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the lifecycle hook.
        """
        return pulumi.get(self, "lifecycle_hook_name")

    @property
    @pulumi.getter(name="lifecycleTransition")
    def lifecycle_transition(self) -> pulumi.Output[str]:
        """
        The instance state to which you want to attach the lifecycle hook.
        """
        return pulumi.get(self, "lifecycle_transition")

    @property
    @pulumi.getter(name="notificationMetadata")
    def notification_metadata(self) -> pulumi.Output[Optional[str]]:
        """
        Additional information that is included any time Amazon EC2 Auto Scaling sends a message to the notification target.
        """
        return pulumi.get(self, "notification_metadata")

    @property
    @pulumi.getter(name="notificationTargetARN")
    def notification_target_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling uses to notify you when an instance is in the transition state for the lifecycle hook. You can specify an Amazon SQS queue or an Amazon SNS topic. The notification message includes the following information: lifecycle action token, user account ID, Auto Scaling group name, lifecycle hook name, instance ID, lifecycle transition, and notification metadata.
        """
        return pulumi.get(self, "notification_target_arn")

    @property
    @pulumi.getter(name="roleARN")
    def role_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target, for example, an Amazon SNS topic or an Amazon SQS queue.
        """
        return pulumi.get(self, "role_arn")

