# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['LaunchArgs', 'Launch']

@pulumi.input_type
class LaunchArgs:
    def __init__(__self__, *,
                 groups: pulumi.Input[Sequence[pulumi.Input['LaunchGroupObjectArgs']]],
                 project: pulumi.Input[str],
                 scheduled_splits_config: pulumi.Input[Sequence[pulumi.Input['LaunchStepConfigArgs']]],
                 description: Optional[pulumi.Input[str]] = None,
                 execution_status: Optional[pulumi.Input['LaunchExecutionStatusObjectArgs']] = None,
                 metric_monitors: Optional[pulumi.Input[Sequence[pulumi.Input['LaunchMetricDefinitionObjectArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 randomization_salt: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['LaunchTagArgs']]]] = None):
        """
        The set of arguments for constructing a Launch resource.
        :param pulumi.Input['LaunchExecutionStatusObjectArgs'] execution_status: Start or Stop Launch Launch. Default is not started.
        :param pulumi.Input[Sequence[pulumi.Input['LaunchTagArgs']]] tags: An array of key-value pairs to apply to this resource.
        """
        pulumi.set(__self__, "groups", groups)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "scheduled_splits_config", scheduled_splits_config)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if execution_status is not None:
            pulumi.set(__self__, "execution_status", execution_status)
        if metric_monitors is not None:
            pulumi.set(__self__, "metric_monitors", metric_monitors)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if randomization_salt is not None:
            pulumi.set(__self__, "randomization_salt", randomization_salt)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def groups(self) -> pulumi.Input[Sequence[pulumi.Input['LaunchGroupObjectArgs']]]:
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: pulumi.Input[Sequence[pulumi.Input['LaunchGroupObjectArgs']]]):
        pulumi.set(self, "groups", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="scheduledSplitsConfig")
    def scheduled_splits_config(self) -> pulumi.Input[Sequence[pulumi.Input['LaunchStepConfigArgs']]]:
        return pulumi.get(self, "scheduled_splits_config")

    @scheduled_splits_config.setter
    def scheduled_splits_config(self, value: pulumi.Input[Sequence[pulumi.Input['LaunchStepConfigArgs']]]):
        pulumi.set(self, "scheduled_splits_config", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="executionStatus")
    def execution_status(self) -> Optional[pulumi.Input['LaunchExecutionStatusObjectArgs']]:
        """
        Start or Stop Launch Launch. Default is not started.
        """
        return pulumi.get(self, "execution_status")

    @execution_status.setter
    def execution_status(self, value: Optional[pulumi.Input['LaunchExecutionStatusObjectArgs']]):
        pulumi.set(self, "execution_status", value)

    @property
    @pulumi.getter(name="metricMonitors")
    def metric_monitors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LaunchMetricDefinitionObjectArgs']]]]:
        return pulumi.get(self, "metric_monitors")

    @metric_monitors.setter
    def metric_monitors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LaunchMetricDefinitionObjectArgs']]]]):
        pulumi.set(self, "metric_monitors", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="randomizationSalt")
    def randomization_salt(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "randomization_salt")

    @randomization_salt.setter
    def randomization_salt(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "randomization_salt", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LaunchTagArgs']]]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LaunchTagArgs']]]]):
        pulumi.set(self, "tags", value)


class Launch(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 execution_status: Optional[pulumi.Input[pulumi.InputType['LaunchExecutionStatusObjectArgs']]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LaunchGroupObjectArgs']]]]] = None,
                 metric_monitors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LaunchMetricDefinitionObjectArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 randomization_salt: Optional[pulumi.Input[str]] = None,
                 scheduled_splits_config: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LaunchStepConfigArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LaunchTagArgs']]]]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::Evidently::Launch.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['LaunchExecutionStatusObjectArgs']] execution_status: Start or Stop Launch Launch. Default is not started.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LaunchTagArgs']]]] tags: An array of key-value pairs to apply to this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LaunchArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::Evidently::Launch.

        :param str resource_name: The name of the resource.
        :param LaunchArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LaunchArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 execution_status: Optional[pulumi.Input[pulumi.InputType['LaunchExecutionStatusObjectArgs']]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LaunchGroupObjectArgs']]]]] = None,
                 metric_monitors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LaunchMetricDefinitionObjectArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 randomization_salt: Optional[pulumi.Input[str]] = None,
                 scheduled_splits_config: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LaunchStepConfigArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LaunchTagArgs']]]]] = None,
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
            __props__ = LaunchArgs.__new__(LaunchArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["execution_status"] = execution_status
            if groups is None and not opts.urn:
                raise TypeError("Missing required property 'groups'")
            __props__.__dict__["groups"] = groups
            __props__.__dict__["metric_monitors"] = metric_monitors
            __props__.__dict__["name"] = name
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["randomization_salt"] = randomization_salt
            if scheduled_splits_config is None and not opts.urn:
                raise TypeError("Missing required property 'scheduled_splits_config'")
            __props__.__dict__["scheduled_splits_config"] = scheduled_splits_config
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
        super(Launch, __self__).__init__(
            'aws-native:evidently:Launch',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Launch':
        """
        Get an existing Launch resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = LaunchArgs.__new__(LaunchArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["execution_status"] = None
        __props__.__dict__["groups"] = None
        __props__.__dict__["metric_monitors"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["randomization_salt"] = None
        __props__.__dict__["scheduled_splits_config"] = None
        __props__.__dict__["tags"] = None
        return Launch(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="executionStatus")
    def execution_status(self) -> pulumi.Output[Optional['outputs.LaunchExecutionStatusObject']]:
        """
        Start or Stop Launch Launch. Default is not started.
        """
        return pulumi.get(self, "execution_status")

    @property
    @pulumi.getter
    def groups(self) -> pulumi.Output[Sequence['outputs.LaunchGroupObject']]:
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter(name="metricMonitors")
    def metric_monitors(self) -> pulumi.Output[Optional[Sequence['outputs.LaunchMetricDefinitionObject']]]:
        return pulumi.get(self, "metric_monitors")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="randomizationSalt")
    def randomization_salt(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "randomization_salt")

    @property
    @pulumi.getter(name="scheduledSplitsConfig")
    def scheduled_splits_config(self) -> pulumi.Output[Sequence['outputs.LaunchStepConfig']]:
        return pulumi.get(self, "scheduled_splits_config")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.LaunchTag']]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

