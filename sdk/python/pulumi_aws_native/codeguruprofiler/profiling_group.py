# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._inputs import *

__all__ = ['ProfilingGroupArgs', 'ProfilingGroup']

@pulumi.input_type
class ProfilingGroupArgs:
    def __init__(__self__, *,
                 profiling_group_name: pulumi.Input[str],
                 agent_permissions: Optional[pulumi.Input[Union[Any, str]]] = None,
                 anomaly_detection_notification_configuration: Optional[pulumi.Input[Sequence[pulumi.Input['ProfilingGroupChannelArgs']]]] = None,
                 compute_platform: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a ProfilingGroup resource.
        :param pulumi.Input[str] profiling_group_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-profilinggroupname
        :param pulumi.Input[Union[Any, str]] agent_permissions: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-agentpermissions
        :param pulumi.Input[Sequence[pulumi.Input['ProfilingGroupChannelArgs']]] anomaly_detection_notification_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-anomalydetectionnotificationconfiguration
        :param pulumi.Input[str] compute_platform: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-computeplatform
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-tags
        """
        pulumi.set(__self__, "profiling_group_name", profiling_group_name)
        if agent_permissions is not None:
            pulumi.set(__self__, "agent_permissions", agent_permissions)
        if anomaly_detection_notification_configuration is not None:
            pulumi.set(__self__, "anomaly_detection_notification_configuration", anomaly_detection_notification_configuration)
        if compute_platform is not None:
            pulumi.set(__self__, "compute_platform", compute_platform)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="profilingGroupName")
    def profiling_group_name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-profilinggroupname
        """
        return pulumi.get(self, "profiling_group_name")

    @profiling_group_name.setter
    def profiling_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "profiling_group_name", value)

    @property
    @pulumi.getter(name="agentPermissions")
    def agent_permissions(self) -> Optional[pulumi.Input[Union[Any, str]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-agentpermissions
        """
        return pulumi.get(self, "agent_permissions")

    @agent_permissions.setter
    def agent_permissions(self, value: Optional[pulumi.Input[Union[Any, str]]]):
        pulumi.set(self, "agent_permissions", value)

    @property
    @pulumi.getter(name="anomalyDetectionNotificationConfiguration")
    def anomaly_detection_notification_configuration(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProfilingGroupChannelArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-anomalydetectionnotificationconfiguration
        """
        return pulumi.get(self, "anomaly_detection_notification_configuration")

    @anomaly_detection_notification_configuration.setter
    def anomaly_detection_notification_configuration(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProfilingGroupChannelArgs']]]]):
        pulumi.set(self, "anomaly_detection_notification_configuration", value)

    @property
    @pulumi.getter(name="computePlatform")
    def compute_platform(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-computeplatform
        """
        return pulumi.get(self, "compute_platform")

    @compute_platform.setter
    def compute_platform(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compute_platform", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


class ProfilingGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_permissions: Optional[pulumi.Input[Union[Any, str]]] = None,
                 anomaly_detection_notification_configuration: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProfilingGroupChannelArgs']]]]] = None,
                 compute_platform: Optional[pulumi.Input[str]] = None,
                 profiling_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union[Any, str]] agent_permissions: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-agentpermissions
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProfilingGroupChannelArgs']]]] anomaly_detection_notification_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-anomalydetectionnotificationconfiguration
        :param pulumi.Input[str] compute_platform: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-computeplatform
        :param pulumi.Input[str] profiling_group_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-profilinggroupname
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-tags
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProfilingGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html

        :param str resource_name: The name of the resource.
        :param ProfilingGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProfilingGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_permissions: Optional[pulumi.Input[Union[Any, str]]] = None,
                 anomaly_detection_notification_configuration: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProfilingGroupChannelArgs']]]]] = None,
                 compute_platform: Optional[pulumi.Input[str]] = None,
                 profiling_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
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
            __props__ = ProfilingGroupArgs.__new__(ProfilingGroupArgs)

            __props__.__dict__["agent_permissions"] = agent_permissions
            __props__.__dict__["anomaly_detection_notification_configuration"] = anomaly_detection_notification_configuration
            __props__.__dict__["compute_platform"] = compute_platform
            if profiling_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'profiling_group_name'")
            __props__.__dict__["profiling_group_name"] = profiling_group_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
        super(ProfilingGroup, __self__).__init__(
            'aws-native:CodeGuruProfiler:ProfilingGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ProfilingGroup':
        """
        Get an existing ProfilingGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ProfilingGroupArgs.__new__(ProfilingGroupArgs)

        __props__.__dict__["agent_permissions"] = None
        __props__.__dict__["anomaly_detection_notification_configuration"] = None
        __props__.__dict__["arn"] = None
        __props__.__dict__["compute_platform"] = None
        __props__.__dict__["profiling_group_name"] = None
        __props__.__dict__["tags"] = None
        return ProfilingGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="agentPermissions")
    def agent_permissions(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-agentpermissions
        """
        return pulumi.get(self, "agent_permissions")

    @property
    @pulumi.getter(name="anomalyDetectionNotificationConfiguration")
    def anomaly_detection_notification_configuration(self) -> pulumi.Output[Optional[Sequence['outputs.ProfilingGroupChannel']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-anomalydetectionnotificationconfiguration
        """
        return pulumi.get(self, "anomaly_detection_notification_configuration")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="computePlatform")
    def compute_platform(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-computeplatform
        """
        return pulumi.get(self, "compute_platform")

    @property
    @pulumi.getter(name="profilingGroupName")
    def profiling_group_name(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-profilinggroupname
        """
        return pulumi.get(self, "profiling_group_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-tags
        """
        return pulumi.get(self, "tags")

