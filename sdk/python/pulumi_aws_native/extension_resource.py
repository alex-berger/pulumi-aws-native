# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ExtensionResourceArgs', 'ExtensionResource']

@pulumi.input_type
class ExtensionResourceArgs:
    def __init__(__self__, *,
                 properties: pulumi.Input[Mapping[str, Any]],
                 type: pulumi.Input[str]):
        """
        The set of arguments for constructing a ExtensionResource resource.
        :param pulumi.Input[Mapping[str, Any]] properties: Dictionary of the extension resource properties.
        :param pulumi.Input[str] type: CloudFormation type name.
        """
        pulumi.set(__self__, "properties", properties)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Input[Mapping[str, Any]]:
        """
        Dictionary of the extension resource properties.
        """
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: pulumi.Input[Mapping[str, Any]]):
        pulumi.set(self, "properties", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        CloudFormation type name.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


class ExtensionResource(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        A special resource that enables deploying CloudFormation Extensions (third-party resources). An extension has to be pre-registered in your AWS account in order to use this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] properties: Dictionary of the extension resource properties.
        :param pulumi.Input[str] type: CloudFormation type name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ExtensionResourceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A special resource that enables deploying CloudFormation Extensions (third-party resources). An extension has to be pre-registered in your AWS account in order to use this resource.

        :param str resource_name: The name of the resource.
        :param ExtensionResourceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ExtensionResourceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
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
            __props__ = ExtensionResourceArgs.__new__(ExtensionResourceArgs)

            if properties is None and not opts.urn:
                raise TypeError("Missing required property 'properties'")
            __props__.__dict__["properties"] = properties
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["outputs"] = None
        super(ExtensionResource, __self__).__init__(
            'aws-native:index:ExtensionResource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ExtensionResource':
        """
        Get an existing ExtensionResource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ExtensionResourceArgs.__new__(ExtensionResourceArgs)

        __props__.__dict__["outputs"] = None
        return ExtensionResource(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def outputs(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        Dictionary of the extension resource attributes.
        """
        return pulumi.get(self, "outputs")
