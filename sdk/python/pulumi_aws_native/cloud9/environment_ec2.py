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

__all__ = ['EnvironmentEC2Args', 'EnvironmentEC2']

@pulumi.input_type
class EnvironmentEC2Args:
    def __init__(__self__, *,
                 instance_type: pulumi.Input[str],
                 automatic_stop_time_minutes: Optional[pulumi.Input[int]] = None,
                 connection_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_arn: Optional[pulumi.Input[str]] = None,
                 repositories: Optional[pulumi.Input[Sequence[pulumi.Input['EnvironmentEC2RepositoryArgs']]]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['EnvironmentEC2TagArgs']]]] = None):
        """
        The set of arguments for constructing a EnvironmentEC2 resource.
        """
        pulumi.set(__self__, "instance_type", instance_type)
        if automatic_stop_time_minutes is not None:
            pulumi.set(__self__, "automatic_stop_time_minutes", automatic_stop_time_minutes)
        if connection_type is not None:
            pulumi.set(__self__, "connection_type", connection_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if image_id is not None:
            pulumi.set(__self__, "image_id", image_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if owner_arn is not None:
            pulumi.set(__self__, "owner_arn", owner_arn)
        if repositories is not None:
            pulumi.set(__self__, "repositories", repositories)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter(name="automaticStopTimeMinutes")
    def automatic_stop_time_minutes(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "automatic_stop_time_minutes")

    @automatic_stop_time_minutes.setter
    def automatic_stop_time_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "automatic_stop_time_minutes", value)

    @property
    @pulumi.getter(name="connectionType")
    def connection_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "connection_type")

    @connection_type.setter
    def connection_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="ownerArn")
    def owner_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "owner_arn")

    @owner_arn.setter
    def owner_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner_arn", value)

    @property
    @pulumi.getter
    def repositories(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EnvironmentEC2RepositoryArgs']]]]:
        return pulumi.get(self, "repositories")

    @repositories.setter
    def repositories(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EnvironmentEC2RepositoryArgs']]]]):
        pulumi.set(self, "repositories", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EnvironmentEC2TagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EnvironmentEC2TagArgs']]]]):
        pulumi.set(self, "tags", value)


warnings.warn("""EnvironmentEC2 is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class EnvironmentEC2(pulumi.CustomResource):
    warnings.warn("""EnvironmentEC2 is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 automatic_stop_time_minutes: Optional[pulumi.Input[int]] = None,
                 connection_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_arn: Optional[pulumi.Input[str]] = None,
                 repositories: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EnvironmentEC2RepositoryArgs']]]]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EnvironmentEC2TagArgs']]]]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::Cloud9::EnvironmentEC2

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EnvironmentEC2Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::Cloud9::EnvironmentEC2

        :param str resource_name: The name of the resource.
        :param EnvironmentEC2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EnvironmentEC2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 automatic_stop_time_minutes: Optional[pulumi.Input[int]] = None,
                 connection_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_arn: Optional[pulumi.Input[str]] = None,
                 repositories: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EnvironmentEC2RepositoryArgs']]]]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EnvironmentEC2TagArgs']]]]] = None,
                 __props__=None):
        pulumi.log.warn("""EnvironmentEC2 is deprecated: EnvironmentEC2 is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EnvironmentEC2Args.__new__(EnvironmentEC2Args)

            __props__.__dict__["automatic_stop_time_minutes"] = automatic_stop_time_minutes
            __props__.__dict__["connection_type"] = connection_type
            __props__.__dict__["description"] = description
            __props__.__dict__["image_id"] = image_id
            if instance_type is None and not opts.urn:
                raise TypeError("Missing required property 'instance_type'")
            __props__.__dict__["instance_type"] = instance_type
            __props__.__dict__["name"] = name
            __props__.__dict__["owner_arn"] = owner_arn
            __props__.__dict__["repositories"] = repositories
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
        super(EnvironmentEC2, __self__).__init__(
            'aws-native:cloud9:EnvironmentEC2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'EnvironmentEC2':
        """
        Get an existing EnvironmentEC2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = EnvironmentEC2Args.__new__(EnvironmentEC2Args)

        __props__.__dict__["arn"] = None
        __props__.__dict__["automatic_stop_time_minutes"] = None
        __props__.__dict__["connection_type"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["image_id"] = None
        __props__.__dict__["instance_type"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["owner_arn"] = None
        __props__.__dict__["repositories"] = None
        __props__.__dict__["subnet_id"] = None
        __props__.__dict__["tags"] = None
        return EnvironmentEC2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="automaticStopTimeMinutes")
    def automatic_stop_time_minutes(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "automatic_stop_time_minutes")

    @property
    @pulumi.getter(name="connectionType")
    def connection_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "connection_type")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerArn")
    def owner_arn(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "owner_arn")

    @property
    @pulumi.getter
    def repositories(self) -> pulumi.Output[Optional[Sequence['outputs.EnvironmentEC2Repository']]]:
        return pulumi.get(self, "repositories")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.EnvironmentEC2Tag']]]:
        return pulumi.get(self, "tags")

