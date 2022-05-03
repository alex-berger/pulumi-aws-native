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

__all__ = ['ServiceArgs', 'Service']

@pulumi.input_type
class ServiceArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 dns_config: Optional[pulumi.Input['ServiceDnsConfigArgs']] = None,
                 health_check_config: Optional[pulumi.Input['ServiceHealthCheckConfigArgs']] = None,
                 health_check_custom_config: Optional[pulumi.Input['ServiceHealthCheckCustomConfigArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceTagArgs']]]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Service resource.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dns_config is not None:
            pulumi.set(__self__, "dns_config", dns_config)
        if health_check_config is not None:
            pulumi.set(__self__, "health_check_config", health_check_config)
        if health_check_custom_config is not None:
            pulumi.set(__self__, "health_check_custom_config", health_check_custom_config)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace_id is not None:
            pulumi.set(__self__, "namespace_id", namespace_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dnsConfig")
    def dns_config(self) -> Optional[pulumi.Input['ServiceDnsConfigArgs']]:
        return pulumi.get(self, "dns_config")

    @dns_config.setter
    def dns_config(self, value: Optional[pulumi.Input['ServiceDnsConfigArgs']]):
        pulumi.set(self, "dns_config", value)

    @property
    @pulumi.getter(name="healthCheckConfig")
    def health_check_config(self) -> Optional[pulumi.Input['ServiceHealthCheckConfigArgs']]:
        return pulumi.get(self, "health_check_config")

    @health_check_config.setter
    def health_check_config(self, value: Optional[pulumi.Input['ServiceHealthCheckConfigArgs']]):
        pulumi.set(self, "health_check_config", value)

    @property
    @pulumi.getter(name="healthCheckCustomConfig")
    def health_check_custom_config(self) -> Optional[pulumi.Input['ServiceHealthCheckCustomConfigArgs']]:
        return pulumi.get(self, "health_check_custom_config")

    @health_check_custom_config.setter
    def health_check_custom_config(self, value: Optional[pulumi.Input['ServiceHealthCheckCustomConfigArgs']]):
        pulumi.set(self, "health_check_custom_config", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServiceTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


warnings.warn("""Service is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class Service(pulumi.CustomResource):
    warnings.warn("""Service is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dns_config: Optional[pulumi.Input[pulumi.InputType['ServiceDnsConfigArgs']]] = None,
                 health_check_config: Optional[pulumi.Input[pulumi.InputType['ServiceHealthCheckConfigArgs']]] = None,
                 health_check_custom_config: Optional[pulumi.Input[pulumi.InputType['ServiceHealthCheckCustomConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceTagArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::ServiceDiscovery::Service

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ServiceArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::ServiceDiscovery::Service

        :param str resource_name: The name of the resource.
        :param ServiceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dns_config: Optional[pulumi.Input[pulumi.InputType['ServiceDnsConfigArgs']]] = None,
                 health_check_config: Optional[pulumi.Input[pulumi.InputType['ServiceHealthCheckConfigArgs']]] = None,
                 health_check_custom_config: Optional[pulumi.Input[pulumi.InputType['ServiceHealthCheckCustomConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceTagArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""Service is deprecated: Service is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceArgs.__new__(ServiceArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["dns_config"] = dns_config
            __props__.__dict__["health_check_config"] = health_check_config
            __props__.__dict__["health_check_custom_config"] = health_check_custom_config
            __props__.__dict__["name"] = name
            __props__.__dict__["namespace_id"] = namespace_id
            __props__.__dict__["tags"] = tags
            __props__.__dict__["type"] = type
            __props__.__dict__["arn"] = None
        super(Service, __self__).__init__(
            'aws-native:servicediscovery:Service',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Service':
        """
        Get an existing Service resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ServiceArgs.__new__(ServiceArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["dns_config"] = None
        __props__.__dict__["health_check_config"] = None
        __props__.__dict__["health_check_custom_config"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["namespace_id"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        return Service(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dnsConfig")
    def dns_config(self) -> pulumi.Output[Optional['outputs.ServiceDnsConfig']]:
        return pulumi.get(self, "dns_config")

    @property
    @pulumi.getter(name="healthCheckConfig")
    def health_check_config(self) -> pulumi.Output[Optional['outputs.ServiceHealthCheckConfig']]:
        return pulumi.get(self, "health_check_config")

    @property
    @pulumi.getter(name="healthCheckCustomConfig")
    def health_check_custom_config(self) -> pulumi.Output[Optional['outputs.ServiceHealthCheckCustomConfig']]:
        return pulumi.get(self, "health_check_custom_config")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "namespace_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.ServiceTag']]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "type")

