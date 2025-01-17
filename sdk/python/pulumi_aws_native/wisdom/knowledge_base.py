# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['KnowledgeBaseArgs', 'KnowledgeBase']

@pulumi.input_type
class KnowledgeBaseArgs:
    def __init__(__self__, *,
                 knowledge_base_type: pulumi.Input['KnowledgeBaseType'],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rendering_configuration: Optional[pulumi.Input['KnowledgeBaseRenderingConfigurationArgs']] = None,
                 server_side_encryption_configuration: Optional[pulumi.Input['KnowledgeBaseServerSideEncryptionConfigurationArgs']] = None,
                 source_configuration: Optional[pulumi.Input['KnowledgeBaseSourceConfigurationArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['KnowledgeBaseTagArgs']]]] = None):
        """
        The set of arguments for constructing a KnowledgeBase resource.
        """
        pulumi.set(__self__, "knowledge_base_type", knowledge_base_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rendering_configuration is not None:
            pulumi.set(__self__, "rendering_configuration", rendering_configuration)
        if server_side_encryption_configuration is not None:
            pulumi.set(__self__, "server_side_encryption_configuration", server_side_encryption_configuration)
        if source_configuration is not None:
            pulumi.set(__self__, "source_configuration", source_configuration)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="knowledgeBaseType")
    def knowledge_base_type(self) -> pulumi.Input['KnowledgeBaseType']:
        return pulumi.get(self, "knowledge_base_type")

    @knowledge_base_type.setter
    def knowledge_base_type(self, value: pulumi.Input['KnowledgeBaseType']):
        pulumi.set(self, "knowledge_base_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="renderingConfiguration")
    def rendering_configuration(self) -> Optional[pulumi.Input['KnowledgeBaseRenderingConfigurationArgs']]:
        return pulumi.get(self, "rendering_configuration")

    @rendering_configuration.setter
    def rendering_configuration(self, value: Optional[pulumi.Input['KnowledgeBaseRenderingConfigurationArgs']]):
        pulumi.set(self, "rendering_configuration", value)

    @property
    @pulumi.getter(name="serverSideEncryptionConfiguration")
    def server_side_encryption_configuration(self) -> Optional[pulumi.Input['KnowledgeBaseServerSideEncryptionConfigurationArgs']]:
        return pulumi.get(self, "server_side_encryption_configuration")

    @server_side_encryption_configuration.setter
    def server_side_encryption_configuration(self, value: Optional[pulumi.Input['KnowledgeBaseServerSideEncryptionConfigurationArgs']]):
        pulumi.set(self, "server_side_encryption_configuration", value)

    @property
    @pulumi.getter(name="sourceConfiguration")
    def source_configuration(self) -> Optional[pulumi.Input['KnowledgeBaseSourceConfigurationArgs']]:
        return pulumi.get(self, "source_configuration")

    @source_configuration.setter
    def source_configuration(self, value: Optional[pulumi.Input['KnowledgeBaseSourceConfigurationArgs']]):
        pulumi.set(self, "source_configuration", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['KnowledgeBaseTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['KnowledgeBaseTagArgs']]]]):
        pulumi.set(self, "tags", value)


class KnowledgeBase(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 knowledge_base_type: Optional[pulumi.Input['KnowledgeBaseType']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rendering_configuration: Optional[pulumi.Input[pulumi.InputType['KnowledgeBaseRenderingConfigurationArgs']]] = None,
                 server_side_encryption_configuration: Optional[pulumi.Input[pulumi.InputType['KnowledgeBaseServerSideEncryptionConfigurationArgs']]] = None,
                 source_configuration: Optional[pulumi.Input[pulumi.InputType['KnowledgeBaseSourceConfigurationArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KnowledgeBaseTagArgs']]]]] = None,
                 __props__=None):
        """
        Definition of AWS::Wisdom::KnowledgeBase Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KnowledgeBaseArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::Wisdom::KnowledgeBase Resource Type

        :param str resource_name: The name of the resource.
        :param KnowledgeBaseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KnowledgeBaseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 knowledge_base_type: Optional[pulumi.Input['KnowledgeBaseType']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rendering_configuration: Optional[pulumi.Input[pulumi.InputType['KnowledgeBaseRenderingConfigurationArgs']]] = None,
                 server_side_encryption_configuration: Optional[pulumi.Input[pulumi.InputType['KnowledgeBaseServerSideEncryptionConfigurationArgs']]] = None,
                 source_configuration: Optional[pulumi.Input[pulumi.InputType['KnowledgeBaseSourceConfigurationArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KnowledgeBaseTagArgs']]]]] = None,
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
            __props__ = KnowledgeBaseArgs.__new__(KnowledgeBaseArgs)

            __props__.__dict__["description"] = description
            if knowledge_base_type is None and not opts.urn:
                raise TypeError("Missing required property 'knowledge_base_type'")
            __props__.__dict__["knowledge_base_type"] = knowledge_base_type
            __props__.__dict__["name"] = name
            __props__.__dict__["rendering_configuration"] = rendering_configuration
            __props__.__dict__["server_side_encryption_configuration"] = server_side_encryption_configuration
            __props__.__dict__["source_configuration"] = source_configuration
            __props__.__dict__["tags"] = tags
            __props__.__dict__["knowledge_base_arn"] = None
            __props__.__dict__["knowledge_base_id"] = None
        super(KnowledgeBase, __self__).__init__(
            'aws-native:wisdom:KnowledgeBase',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'KnowledgeBase':
        """
        Get an existing KnowledgeBase resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = KnowledgeBaseArgs.__new__(KnowledgeBaseArgs)

        __props__.__dict__["description"] = None
        __props__.__dict__["knowledge_base_arn"] = None
        __props__.__dict__["knowledge_base_id"] = None
        __props__.__dict__["knowledge_base_type"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["rendering_configuration"] = None
        __props__.__dict__["server_side_encryption_configuration"] = None
        __props__.__dict__["source_configuration"] = None
        __props__.__dict__["tags"] = None
        return KnowledgeBase(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="knowledgeBaseArn")
    def knowledge_base_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "knowledge_base_arn")

    @property
    @pulumi.getter(name="knowledgeBaseId")
    def knowledge_base_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "knowledge_base_id")

    @property
    @pulumi.getter(name="knowledgeBaseType")
    def knowledge_base_type(self) -> pulumi.Output['KnowledgeBaseType']:
        return pulumi.get(self, "knowledge_base_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="renderingConfiguration")
    def rendering_configuration(self) -> pulumi.Output[Optional['outputs.KnowledgeBaseRenderingConfiguration']]:
        return pulumi.get(self, "rendering_configuration")

    @property
    @pulumi.getter(name="serverSideEncryptionConfiguration")
    def server_side_encryption_configuration(self) -> pulumi.Output[Optional['outputs.KnowledgeBaseServerSideEncryptionConfiguration']]:
        return pulumi.get(self, "server_side_encryption_configuration")

    @property
    @pulumi.getter(name="sourceConfiguration")
    def source_configuration(self) -> pulumi.Output[Optional['outputs.KnowledgeBaseSourceConfiguration']]:
        return pulumi.get(self, "source_configuration")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.KnowledgeBaseTag']]]:
        return pulumi.get(self, "tags")

