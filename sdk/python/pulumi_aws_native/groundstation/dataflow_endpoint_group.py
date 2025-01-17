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

__all__ = ['DataflowEndpointGroupArgs', 'DataflowEndpointGroup']

@pulumi.input_type
class DataflowEndpointGroupArgs:
    def __init__(__self__, *,
                 endpoint_details: pulumi.Input[Sequence[pulumi.Input['DataflowEndpointGroupEndpointDetailsArgs']]],
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['DataflowEndpointGroupTagArgs']]]] = None):
        """
        The set of arguments for constructing a DataflowEndpointGroup resource.
        """
        pulumi.set(__self__, "endpoint_details", endpoint_details)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="endpointDetails")
    def endpoint_details(self) -> pulumi.Input[Sequence[pulumi.Input['DataflowEndpointGroupEndpointDetailsArgs']]]:
        return pulumi.get(self, "endpoint_details")

    @endpoint_details.setter
    def endpoint_details(self, value: pulumi.Input[Sequence[pulumi.Input['DataflowEndpointGroupEndpointDetailsArgs']]]):
        pulumi.set(self, "endpoint_details", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataflowEndpointGroupTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataflowEndpointGroupTagArgs']]]]):
        pulumi.set(self, "tags", value)


class DataflowEndpointGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint_details: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataflowEndpointGroupEndpointDetailsArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataflowEndpointGroupTagArgs']]]]] = None,
                 __props__=None):
        """
        AWS Ground Station DataflowEndpointGroup schema for CloudFormation

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DataflowEndpointGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        AWS Ground Station DataflowEndpointGroup schema for CloudFormation

        :param str resource_name: The name of the resource.
        :param DataflowEndpointGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataflowEndpointGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint_details: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataflowEndpointGroupEndpointDetailsArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataflowEndpointGroupTagArgs']]]]] = None,
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
            __props__ = DataflowEndpointGroupArgs.__new__(DataflowEndpointGroupArgs)

            if endpoint_details is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint_details'")
            __props__.__dict__["endpoint_details"] = endpoint_details
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
        super(DataflowEndpointGroup, __self__).__init__(
            'aws-native:groundstation:DataflowEndpointGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DataflowEndpointGroup':
        """
        Get an existing DataflowEndpointGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DataflowEndpointGroupArgs.__new__(DataflowEndpointGroupArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["endpoint_details"] = None
        __props__.__dict__["tags"] = None
        return DataflowEndpointGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="endpointDetails")
    def endpoint_details(self) -> pulumi.Output[Sequence['outputs.DataflowEndpointGroupEndpointDetails']]:
        return pulumi.get(self, "endpoint_details")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.DataflowEndpointGroupTag']]]:
        return pulumi.get(self, "tags")

