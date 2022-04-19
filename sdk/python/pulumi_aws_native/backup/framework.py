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

__all__ = ['FrameworkArgs', 'Framework']

@pulumi.input_type
class FrameworkArgs:
    def __init__(__self__, *,
                 framework_controls: pulumi.Input[Sequence[pulumi.Input['FrameworkControlArgs']]],
                 framework_description: Optional[pulumi.Input[str]] = None,
                 framework_name: Optional[pulumi.Input[str]] = None,
                 framework_tags: Optional[pulumi.Input[Sequence[pulumi.Input['FrameworkTagArgs']]]] = None):
        """
        The set of arguments for constructing a Framework resource.
        :param pulumi.Input[Sequence[pulumi.Input['FrameworkControlArgs']]] framework_controls: Contains detailed information about all of the controls of a framework. Each framework must contain at least one control.
        :param pulumi.Input[str] framework_description: An optional description of the framework with a maximum 1,024 characters.
        :param pulumi.Input[str] framework_name: The unique name of a framework. This name is between 1 and 256 characters, starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_).
        :param pulumi.Input[Sequence[pulumi.Input['FrameworkTagArgs']]] framework_tags: Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.
        """
        pulumi.set(__self__, "framework_controls", framework_controls)
        if framework_description is not None:
            pulumi.set(__self__, "framework_description", framework_description)
        if framework_name is not None:
            pulumi.set(__self__, "framework_name", framework_name)
        if framework_tags is not None:
            pulumi.set(__self__, "framework_tags", framework_tags)

    @property
    @pulumi.getter(name="frameworkControls")
    def framework_controls(self) -> pulumi.Input[Sequence[pulumi.Input['FrameworkControlArgs']]]:
        """
        Contains detailed information about all of the controls of a framework. Each framework must contain at least one control.
        """
        return pulumi.get(self, "framework_controls")

    @framework_controls.setter
    def framework_controls(self, value: pulumi.Input[Sequence[pulumi.Input['FrameworkControlArgs']]]):
        pulumi.set(self, "framework_controls", value)

    @property
    @pulumi.getter(name="frameworkDescription")
    def framework_description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of the framework with a maximum 1,024 characters.
        """
        return pulumi.get(self, "framework_description")

    @framework_description.setter
    def framework_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "framework_description", value)

    @property
    @pulumi.getter(name="frameworkName")
    def framework_name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of a framework. This name is between 1 and 256 characters, starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_).
        """
        return pulumi.get(self, "framework_name")

    @framework_name.setter
    def framework_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "framework_name", value)

    @property
    @pulumi.getter(name="frameworkTags")
    def framework_tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FrameworkTagArgs']]]]:
        """
        Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.
        """
        return pulumi.get(self, "framework_tags")

    @framework_tags.setter
    def framework_tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FrameworkTagArgs']]]]):
        pulumi.set(self, "framework_tags", value)


class Framework(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 framework_controls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FrameworkControlArgs']]]]] = None,
                 framework_description: Optional[pulumi.Input[str]] = None,
                 framework_name: Optional[pulumi.Input[str]] = None,
                 framework_tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FrameworkTagArgs']]]]] = None,
                 __props__=None):
        """
        Contains detailed information about a framework. Frameworks contain controls, which evaluate and report on your backup events and resources. Frameworks generate daily compliance results.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FrameworkControlArgs']]]] framework_controls: Contains detailed information about all of the controls of a framework. Each framework must contain at least one control.
        :param pulumi.Input[str] framework_description: An optional description of the framework with a maximum 1,024 characters.
        :param pulumi.Input[str] framework_name: The unique name of a framework. This name is between 1 and 256 characters, starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FrameworkTagArgs']]]] framework_tags: Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FrameworkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Contains detailed information about a framework. Frameworks contain controls, which evaluate and report on your backup events and resources. Frameworks generate daily compliance results.

        :param str resource_name: The name of the resource.
        :param FrameworkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FrameworkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 framework_controls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FrameworkControlArgs']]]]] = None,
                 framework_description: Optional[pulumi.Input[str]] = None,
                 framework_name: Optional[pulumi.Input[str]] = None,
                 framework_tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FrameworkTagArgs']]]]] = None,
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
            __props__ = FrameworkArgs.__new__(FrameworkArgs)

            if framework_controls is None and not opts.urn:
                raise TypeError("Missing required property 'framework_controls'")
            __props__.__dict__["framework_controls"] = framework_controls
            __props__.__dict__["framework_description"] = framework_description
            __props__.__dict__["framework_name"] = framework_name
            __props__.__dict__["framework_tags"] = framework_tags
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["deployment_status"] = None
            __props__.__dict__["framework_arn"] = None
            __props__.__dict__["framework_status"] = None
        super(Framework, __self__).__init__(
            'aws-native:backup:Framework',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Framework':
        """
        Get an existing Framework resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FrameworkArgs.__new__(FrameworkArgs)

        __props__.__dict__["creation_time"] = None
        __props__.__dict__["deployment_status"] = None
        __props__.__dict__["framework_arn"] = None
        __props__.__dict__["framework_controls"] = None
        __props__.__dict__["framework_description"] = None
        __props__.__dict__["framework_name"] = None
        __props__.__dict__["framework_status"] = None
        __props__.__dict__["framework_tags"] = None
        return Framework(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[float]:
        """
        The date and time that a framework is created, in Unix format and Coordinated Universal Time (UTC). The value of `CreationTime` is accurate to milliseconds. For example, the value 1516925490.087 represents Friday, January 26, 2018 12:11:30.087 AM.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="deploymentStatus")
    def deployment_status(self) -> pulumi.Output[str]:
        """
        The deployment status of a framework. The statuses are: `CREATE_IN_PROGRESS | UPDATE_IN_PROGRESS | DELETE_IN_PROGRESS | COMPLETED | FAILED`
        """
        return pulumi.get(self, "deployment_status")

    @property
    @pulumi.getter(name="frameworkArn")
    def framework_arn(self) -> pulumi.Output[str]:
        """
        An Amazon Resource Name (ARN) that uniquely identifies Framework as a resource
        """
        return pulumi.get(self, "framework_arn")

    @property
    @pulumi.getter(name="frameworkControls")
    def framework_controls(self) -> pulumi.Output[Sequence['outputs.FrameworkControl']]:
        """
        Contains detailed information about all of the controls of a framework. Each framework must contain at least one control.
        """
        return pulumi.get(self, "framework_controls")

    @property
    @pulumi.getter(name="frameworkDescription")
    def framework_description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional description of the framework with a maximum 1,024 characters.
        """
        return pulumi.get(self, "framework_description")

    @property
    @pulumi.getter(name="frameworkName")
    def framework_name(self) -> pulumi.Output[Optional[str]]:
        """
        The unique name of a framework. This name is between 1 and 256 characters, starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_).
        """
        return pulumi.get(self, "framework_name")

    @property
    @pulumi.getter(name="frameworkStatus")
    def framework_status(self) -> pulumi.Output[str]:
        """
        A framework consists of one or more controls. Each control governs a resource, such as backup plans, backup selections, backup vaults, or recovery points. You can also turn AWS Config recording on or off for each resource. The statuses are:

        `ACTIVE` when recording is turned on for all resources governed by the framework.

        `PARTIALLY_ACTIVE` when recording is turned off for at least one resource governed by the framework.

        `INACTIVE` when recording is turned off for all resources governed by the framework.

        `UNAVAILABLE` when AWS Backup is unable to validate recording status at this time.
        """
        return pulumi.get(self, "framework_status")

    @property
    @pulumi.getter(name="frameworkTags")
    def framework_tags(self) -> pulumi.Output[Optional[Sequence['outputs.FrameworkTag']]]:
        """
        Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.
        """
        return pulumi.get(self, "framework_tags")

