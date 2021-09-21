# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'HttpNamespaceTag',
    'PrivateDnsNamespacePrivateDnsPropertiesMutable',
    'PrivateDnsNamespaceProperties',
    'PrivateDnsNamespaceSOA',
    'PrivateDnsNamespaceTag',
    'PublicDnsNamespaceProperties',
    'PublicDnsNamespacePublicDnsPropertiesMutable',
    'PublicDnsNamespaceSOA',
    'PublicDnsNamespaceTag',
    'ServiceDnsConfig',
    'ServiceDnsRecord',
    'ServiceHealthCheckConfig',
    'ServiceHealthCheckCustomConfig',
    'ServiceTag',
]

@pulumi.output_type
class HttpNamespaceTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


@pulumi.output_type
class PrivateDnsNamespacePrivateDnsPropertiesMutable(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "sOA":
            suggest = "s_oa"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PrivateDnsNamespacePrivateDnsPropertiesMutable. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PrivateDnsNamespacePrivateDnsPropertiesMutable.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PrivateDnsNamespacePrivateDnsPropertiesMutable.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 s_oa: Optional['outputs.PrivateDnsNamespaceSOA'] = None):
        if s_oa is not None:
            pulumi.set(__self__, "s_oa", s_oa)

    @property
    @pulumi.getter(name="sOA")
    def s_oa(self) -> Optional['outputs.PrivateDnsNamespaceSOA']:
        return pulumi.get(self, "s_oa")


@pulumi.output_type
class PrivateDnsNamespaceProperties(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "dnsProperties":
            suggest = "dns_properties"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PrivateDnsNamespaceProperties. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PrivateDnsNamespaceProperties.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PrivateDnsNamespaceProperties.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 dns_properties: Optional['outputs.PrivateDnsNamespacePrivateDnsPropertiesMutable'] = None):
        if dns_properties is not None:
            pulumi.set(__self__, "dns_properties", dns_properties)

    @property
    @pulumi.getter(name="dnsProperties")
    def dns_properties(self) -> Optional['outputs.PrivateDnsNamespacePrivateDnsPropertiesMutable']:
        return pulumi.get(self, "dns_properties")


@pulumi.output_type
class PrivateDnsNamespaceSOA(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "tTL":
            suggest = "t_tl"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PrivateDnsNamespaceSOA. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PrivateDnsNamespaceSOA.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PrivateDnsNamespaceSOA.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 t_tl: Optional[float] = None):
        if t_tl is not None:
            pulumi.set(__self__, "t_tl", t_tl)

    @property
    @pulumi.getter(name="tTL")
    def t_tl(self) -> Optional[float]:
        return pulumi.get(self, "t_tl")


@pulumi.output_type
class PrivateDnsNamespaceTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


@pulumi.output_type
class PublicDnsNamespaceProperties(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "dnsProperties":
            suggest = "dns_properties"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PublicDnsNamespaceProperties. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PublicDnsNamespaceProperties.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PublicDnsNamespaceProperties.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 dns_properties: Optional['outputs.PublicDnsNamespacePublicDnsPropertiesMutable'] = None):
        if dns_properties is not None:
            pulumi.set(__self__, "dns_properties", dns_properties)

    @property
    @pulumi.getter(name="dnsProperties")
    def dns_properties(self) -> Optional['outputs.PublicDnsNamespacePublicDnsPropertiesMutable']:
        return pulumi.get(self, "dns_properties")


@pulumi.output_type
class PublicDnsNamespacePublicDnsPropertiesMutable(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "sOA":
            suggest = "s_oa"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PublicDnsNamespacePublicDnsPropertiesMutable. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PublicDnsNamespacePublicDnsPropertiesMutable.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PublicDnsNamespacePublicDnsPropertiesMutable.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 s_oa: Optional['outputs.PublicDnsNamespaceSOA'] = None):
        if s_oa is not None:
            pulumi.set(__self__, "s_oa", s_oa)

    @property
    @pulumi.getter(name="sOA")
    def s_oa(self) -> Optional['outputs.PublicDnsNamespaceSOA']:
        return pulumi.get(self, "s_oa")


@pulumi.output_type
class PublicDnsNamespaceSOA(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "tTL":
            suggest = "t_tl"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PublicDnsNamespaceSOA. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PublicDnsNamespaceSOA.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PublicDnsNamespaceSOA.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 t_tl: Optional[float] = None):
        if t_tl is not None:
            pulumi.set(__self__, "t_tl", t_tl)

    @property
    @pulumi.getter(name="tTL")
    def t_tl(self) -> Optional[float]:
        return pulumi.get(self, "t_tl")


@pulumi.output_type
class PublicDnsNamespaceTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


@pulumi.output_type
class ServiceDnsConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "dnsRecords":
            suggest = "dns_records"
        elif key == "namespaceId":
            suggest = "namespace_id"
        elif key == "routingPolicy":
            suggest = "routing_policy"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ServiceDnsConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ServiceDnsConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ServiceDnsConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 dns_records: Sequence['outputs.ServiceDnsRecord'],
                 namespace_id: Optional[str] = None,
                 routing_policy: Optional[str] = None):
        pulumi.set(__self__, "dns_records", dns_records)
        if namespace_id is not None:
            pulumi.set(__self__, "namespace_id", namespace_id)
        if routing_policy is not None:
            pulumi.set(__self__, "routing_policy", routing_policy)

    @property
    @pulumi.getter(name="dnsRecords")
    def dns_records(self) -> Sequence['outputs.ServiceDnsRecord']:
        return pulumi.get(self, "dns_records")

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[str]:
        return pulumi.get(self, "namespace_id")

    @property
    @pulumi.getter(name="routingPolicy")
    def routing_policy(self) -> Optional[str]:
        return pulumi.get(self, "routing_policy")


@pulumi.output_type
class ServiceDnsRecord(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "tTL":
            suggest = "t_tl"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ServiceDnsRecord. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ServiceDnsRecord.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ServiceDnsRecord.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 t_tl: float,
                 type: str):
        pulumi.set(__self__, "t_tl", t_tl)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="tTL")
    def t_tl(self) -> float:
        return pulumi.get(self, "t_tl")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


@pulumi.output_type
class ServiceHealthCheckConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "failureThreshold":
            suggest = "failure_threshold"
        elif key == "resourcePath":
            suggest = "resource_path"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ServiceHealthCheckConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ServiceHealthCheckConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ServiceHealthCheckConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 type: str,
                 failure_threshold: Optional[float] = None,
                 resource_path: Optional[str] = None):
        pulumi.set(__self__, "type", type)
        if failure_threshold is not None:
            pulumi.set(__self__, "failure_threshold", failure_threshold)
        if resource_path is not None:
            pulumi.set(__self__, "resource_path", resource_path)

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="failureThreshold")
    def failure_threshold(self) -> Optional[float]:
        return pulumi.get(self, "failure_threshold")

    @property
    @pulumi.getter(name="resourcePath")
    def resource_path(self) -> Optional[str]:
        return pulumi.get(self, "resource_path")


@pulumi.output_type
class ServiceHealthCheckCustomConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "failureThreshold":
            suggest = "failure_threshold"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ServiceHealthCheckCustomConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ServiceHealthCheckCustomConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ServiceHealthCheckCustomConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 failure_threshold: Optional[float] = None):
        if failure_threshold is not None:
            pulumi.set(__self__, "failure_threshold", failure_threshold)

    @property
    @pulumi.getter(name="failureThreshold")
    def failure_threshold(self) -> Optional[float]:
        return pulumi.get(self, "failure_threshold")


@pulumi.output_type
class ServiceTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")

