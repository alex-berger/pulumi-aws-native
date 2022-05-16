# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'AcceleratorTagArgs',
    'EndpointGroupEndpointConfigurationArgs',
    'EndpointGroupPortOverrideArgs',
    'ListenerPortRangeArgs',
]

@pulumi.input_type
class AcceleratorTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        Tag is a key-value pair associated with accelerator.
        :param pulumi.Input[str] key: Key of the tag. Value can be 1 to 127 characters.
        :param pulumi.Input[str] value: Value for the tag. Value can be 1 to 255 characters.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        Key of the tag. Value can be 1 to 127 characters.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        Value for the tag. Value can be 1 to 255 characters.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class EndpointGroupEndpointConfigurationArgs:
    def __init__(__self__, *,
                 endpoint_id: pulumi.Input[str],
                 client_ip_preservation_enabled: Optional[pulumi.Input[bool]] = None,
                 weight: Optional[pulumi.Input[int]] = None):
        """
        The configuration for a given endpoint
        :param pulumi.Input[str] endpoint_id: Id of the endpoint. For Network/Application Load Balancer this value is the ARN.  For EIP, this value is the allocation ID.  For EC2 instances, this is the EC2 instance ID
        :param pulumi.Input[bool] client_ip_preservation_enabled: true if client ip should be preserved
        :param pulumi.Input[int] weight: The weight for the endpoint.
        """
        pulumi.set(__self__, "endpoint_id", endpoint_id)
        if client_ip_preservation_enabled is not None:
            pulumi.set(__self__, "client_ip_preservation_enabled", client_ip_preservation_enabled)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> pulumi.Input[str]:
        """
        Id of the endpoint. For Network/Application Load Balancer this value is the ARN.  For EIP, this value is the allocation ID.  For EC2 instances, this is the EC2 instance ID
        """
        return pulumi.get(self, "endpoint_id")

    @endpoint_id.setter
    def endpoint_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint_id", value)

    @property
    @pulumi.getter(name="clientIPPreservationEnabled")
    def client_ip_preservation_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        true if client ip should be preserved
        """
        return pulumi.get(self, "client_ip_preservation_enabled")

    @client_ip_preservation_enabled.setter
    def client_ip_preservation_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "client_ip_preservation_enabled", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        """
        The weight for the endpoint.
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class EndpointGroupPortOverrideArgs:
    def __init__(__self__, *,
                 endpoint_port: pulumi.Input[int],
                 listener_port: pulumi.Input[int]):
        """
        listener to endpoint port mapping.
        """
        pulumi.set(__self__, "endpoint_port", endpoint_port)
        pulumi.set(__self__, "listener_port", listener_port)

    @property
    @pulumi.getter(name="endpointPort")
    def endpoint_port(self) -> pulumi.Input[int]:
        return pulumi.get(self, "endpoint_port")

    @endpoint_port.setter
    def endpoint_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "endpoint_port", value)

    @property
    @pulumi.getter(name="listenerPort")
    def listener_port(self) -> pulumi.Input[int]:
        return pulumi.get(self, "listener_port")

    @listener_port.setter
    def listener_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "listener_port", value)


@pulumi.input_type
class ListenerPortRangeArgs:
    def __init__(__self__, *,
                 from_port: pulumi.Input[int],
                 to_port: pulumi.Input[int]):
        """
        A port range to support for connections from  clients to your accelerator.
        """
        pulumi.set(__self__, "from_port", from_port)
        pulumi.set(__self__, "to_port", to_port)

    @property
    @pulumi.getter(name="fromPort")
    def from_port(self) -> pulumi.Input[int]:
        return pulumi.get(self, "from_port")

    @from_port.setter
    def from_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "from_port", value)

    @property
    @pulumi.getter(name="toPort")
    def to_port(self) -> pulumi.Input[int]:
        return pulumi.get(self, "to_port")

    @to_port.setter
    def to_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "to_port", value)

