# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'HttpNamespaceTagArgs',
    'PrivateDnsNamespacePrivateDnsPropertiesMutableArgs',
    'PrivateDnsNamespacePropertiesArgs',
    'PrivateDnsNamespaceSOAArgs',
    'PrivateDnsNamespaceTagArgs',
    'PublicDnsNamespacePropertiesArgs',
    'PublicDnsNamespacePublicDnsPropertiesMutableArgs',
    'PublicDnsNamespaceSOAArgs',
    'PublicDnsNamespaceTagArgs',
    'ServiceDnsConfigArgs',
    'ServiceDnsRecordArgs',
    'ServiceHealthCheckConfigArgs',
    'ServiceHealthCheckCustomConfigArgs',
    'ServiceTagArgs',
]

@pulumi.input_type
class HttpNamespaceTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class PrivateDnsNamespacePrivateDnsPropertiesMutableArgs:
    def __init__(__self__, *,
                 s_oa: Optional[pulumi.Input['PrivateDnsNamespaceSOAArgs']] = None):
        if s_oa is not None:
            pulumi.set(__self__, "s_oa", s_oa)

    @property
    @pulumi.getter(name="sOA")
    def s_oa(self) -> Optional[pulumi.Input['PrivateDnsNamespaceSOAArgs']]:
        return pulumi.get(self, "s_oa")

    @s_oa.setter
    def s_oa(self, value: Optional[pulumi.Input['PrivateDnsNamespaceSOAArgs']]):
        pulumi.set(self, "s_oa", value)


@pulumi.input_type
class PrivateDnsNamespacePropertiesArgs:
    def __init__(__self__, *,
                 dns_properties: Optional[pulumi.Input['PrivateDnsNamespacePrivateDnsPropertiesMutableArgs']] = None):
        if dns_properties is not None:
            pulumi.set(__self__, "dns_properties", dns_properties)

    @property
    @pulumi.getter(name="dnsProperties")
    def dns_properties(self) -> Optional[pulumi.Input['PrivateDnsNamespacePrivateDnsPropertiesMutableArgs']]:
        return pulumi.get(self, "dns_properties")

    @dns_properties.setter
    def dns_properties(self, value: Optional[pulumi.Input['PrivateDnsNamespacePrivateDnsPropertiesMutableArgs']]):
        pulumi.set(self, "dns_properties", value)


@pulumi.input_type
class PrivateDnsNamespaceSOAArgs:
    def __init__(__self__, *,
                 t_tl: Optional[pulumi.Input[float]] = None):
        if t_tl is not None:
            pulumi.set(__self__, "t_tl", t_tl)

    @property
    @pulumi.getter(name="tTL")
    def t_tl(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "t_tl")

    @t_tl.setter
    def t_tl(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "t_tl", value)


@pulumi.input_type
class PrivateDnsNamespaceTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class PublicDnsNamespacePropertiesArgs:
    def __init__(__self__, *,
                 dns_properties: Optional[pulumi.Input['PublicDnsNamespacePublicDnsPropertiesMutableArgs']] = None):
        if dns_properties is not None:
            pulumi.set(__self__, "dns_properties", dns_properties)

    @property
    @pulumi.getter(name="dnsProperties")
    def dns_properties(self) -> Optional[pulumi.Input['PublicDnsNamespacePublicDnsPropertiesMutableArgs']]:
        return pulumi.get(self, "dns_properties")

    @dns_properties.setter
    def dns_properties(self, value: Optional[pulumi.Input['PublicDnsNamespacePublicDnsPropertiesMutableArgs']]):
        pulumi.set(self, "dns_properties", value)


@pulumi.input_type
class PublicDnsNamespacePublicDnsPropertiesMutableArgs:
    def __init__(__self__, *,
                 s_oa: Optional[pulumi.Input['PublicDnsNamespaceSOAArgs']] = None):
        if s_oa is not None:
            pulumi.set(__self__, "s_oa", s_oa)

    @property
    @pulumi.getter(name="sOA")
    def s_oa(self) -> Optional[pulumi.Input['PublicDnsNamespaceSOAArgs']]:
        return pulumi.get(self, "s_oa")

    @s_oa.setter
    def s_oa(self, value: Optional[pulumi.Input['PublicDnsNamespaceSOAArgs']]):
        pulumi.set(self, "s_oa", value)


@pulumi.input_type
class PublicDnsNamespaceSOAArgs:
    def __init__(__self__, *,
                 t_tl: Optional[pulumi.Input[float]] = None):
        if t_tl is not None:
            pulumi.set(__self__, "t_tl", t_tl)

    @property
    @pulumi.getter(name="tTL")
    def t_tl(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "t_tl")

    @t_tl.setter
    def t_tl(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "t_tl", value)


@pulumi.input_type
class PublicDnsNamespaceTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ServiceDnsConfigArgs:
    def __init__(__self__, *,
                 dns_records: pulumi.Input[Sequence[pulumi.Input['ServiceDnsRecordArgs']]],
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 routing_policy: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "dns_records", dns_records)
        if namespace_id is not None:
            pulumi.set(__self__, "namespace_id", namespace_id)
        if routing_policy is not None:
            pulumi.set(__self__, "routing_policy", routing_policy)

    @property
    @pulumi.getter(name="dnsRecords")
    def dns_records(self) -> pulumi.Input[Sequence[pulumi.Input['ServiceDnsRecordArgs']]]:
        return pulumi.get(self, "dns_records")

    @dns_records.setter
    def dns_records(self, value: pulumi.Input[Sequence[pulumi.Input['ServiceDnsRecordArgs']]]):
        pulumi.set(self, "dns_records", value)

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_id", value)

    @property
    @pulumi.getter(name="routingPolicy")
    def routing_policy(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "routing_policy")

    @routing_policy.setter
    def routing_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "routing_policy", value)


@pulumi.input_type
class ServiceDnsRecordArgs:
    def __init__(__self__, *,
                 t_tl: pulumi.Input[float],
                 type: pulumi.Input[str]):
        pulumi.set(__self__, "t_tl", t_tl)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="tTL")
    def t_tl(self) -> pulumi.Input[float]:
        return pulumi.get(self, "t_tl")

    @t_tl.setter
    def t_tl(self, value: pulumi.Input[float]):
        pulumi.set(self, "t_tl", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ServiceHealthCheckConfigArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 failure_threshold: Optional[pulumi.Input[float]] = None,
                 resource_path: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "type", type)
        if failure_threshold is not None:
            pulumi.set(__self__, "failure_threshold", failure_threshold)
        if resource_path is not None:
            pulumi.set(__self__, "resource_path", resource_path)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="failureThreshold")
    def failure_threshold(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "failure_threshold")

    @failure_threshold.setter
    def failure_threshold(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "failure_threshold", value)

    @property
    @pulumi.getter(name="resourcePath")
    def resource_path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "resource_path")

    @resource_path.setter
    def resource_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_path", value)


@pulumi.input_type
class ServiceHealthCheckCustomConfigArgs:
    def __init__(__self__, *,
                 failure_threshold: Optional[pulumi.Input[float]] = None):
        if failure_threshold is not None:
            pulumi.set(__self__, "failure_threshold", failure_threshold)

    @property
    @pulumi.getter(name="failureThreshold")
    def failure_threshold(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "failure_threshold")

    @failure_threshold.setter
    def failure_threshold(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "failure_threshold", value)


@pulumi.input_type
class ServiceTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


