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

__all__ = ['EndpointArgs', 'Endpoint']

@pulumi.input_type
class EndpointArgs:
    def __init__(__self__, *,
                 event_buses: pulumi.Input[Sequence[pulumi.Input['EndpointEventBusArgs']]],
                 routing_config: pulumi.Input['EndpointRoutingConfigArgs'],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 replication_config: Optional[pulumi.Input['EndpointReplicationConfigArgs']] = None,
                 role_arn: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Endpoint resource.
        """
        pulumi.set(__self__, "event_buses", event_buses)
        pulumi.set(__self__, "routing_config", routing_config)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if replication_config is not None:
            pulumi.set(__self__, "replication_config", replication_config)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)

    @property
    @pulumi.getter(name="eventBuses")
    def event_buses(self) -> pulumi.Input[Sequence[pulumi.Input['EndpointEventBusArgs']]]:
        return pulumi.get(self, "event_buses")

    @event_buses.setter
    def event_buses(self, value: pulumi.Input[Sequence[pulumi.Input['EndpointEventBusArgs']]]):
        pulumi.set(self, "event_buses", value)

    @property
    @pulumi.getter(name="routingConfig")
    def routing_config(self) -> pulumi.Input['EndpointRoutingConfigArgs']:
        return pulumi.get(self, "routing_config")

    @routing_config.setter
    def routing_config(self, value: pulumi.Input['EndpointRoutingConfigArgs']):
        pulumi.set(self, "routing_config", value)

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
    @pulumi.getter(name="replicationConfig")
    def replication_config(self) -> Optional[pulumi.Input['EndpointReplicationConfigArgs']]:
        return pulumi.get(self, "replication_config")

    @replication_config.setter
    def replication_config(self, value: Optional[pulumi.Input['EndpointReplicationConfigArgs']]):
        pulumi.set(self, "replication_config", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)


class Endpoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 event_buses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointEventBusArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 replication_config: Optional[pulumi.Input[pulumi.InputType['EndpointReplicationConfigArgs']]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 routing_config: Optional[pulumi.Input[pulumi.InputType['EndpointRoutingConfigArgs']]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::Events::Endpoint.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EndpointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::Events::Endpoint.

        :param str resource_name: The name of the resource.
        :param EndpointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EndpointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 event_buses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointEventBusArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 replication_config: Optional[pulumi.Input[pulumi.InputType['EndpointReplicationConfigArgs']]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 routing_config: Optional[pulumi.Input[pulumi.InputType['EndpointRoutingConfigArgs']]] = None,
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
            __props__ = EndpointArgs.__new__(EndpointArgs)

            __props__.__dict__["description"] = description
            if event_buses is None and not opts.urn:
                raise TypeError("Missing required property 'event_buses'")
            __props__.__dict__["event_buses"] = event_buses
            __props__.__dict__["name"] = name
            __props__.__dict__["replication_config"] = replication_config
            __props__.__dict__["role_arn"] = role_arn
            if routing_config is None and not opts.urn:
                raise TypeError("Missing required property 'routing_config'")
            __props__.__dict__["routing_config"] = routing_config
            __props__.__dict__["arn"] = None
            __props__.__dict__["endpoint_id"] = None
            __props__.__dict__["endpoint_url"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["state_reason"] = None
        super(Endpoint, __self__).__init__(
            'aws-native:events:Endpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Endpoint':
        """
        Get an existing Endpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = EndpointArgs.__new__(EndpointArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["endpoint_id"] = None
        __props__.__dict__["endpoint_url"] = None
        __props__.__dict__["event_buses"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["replication_config"] = None
        __props__.__dict__["role_arn"] = None
        __props__.__dict__["routing_config"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["state_reason"] = None
        return Endpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "endpoint_id")

    @property
    @pulumi.getter(name="endpointUrl")
    def endpoint_url(self) -> pulumi.Output[str]:
        return pulumi.get(self, "endpoint_url")

    @property
    @pulumi.getter(name="eventBuses")
    def event_buses(self) -> pulumi.Output[Sequence['outputs.EndpointEventBus']]:
        return pulumi.get(self, "event_buses")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="replicationConfig")
    def replication_config(self) -> pulumi.Output[Optional['outputs.EndpointReplicationConfig']]:
        return pulumi.get(self, "replication_config")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="routingConfig")
    def routing_config(self) -> pulumi.Output['outputs.EndpointRoutingConfig']:
        return pulumi.get(self, "routing_config")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output['EndpointState']:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="stateReason")
    def state_reason(self) -> pulumi.Output[str]:
        return pulumi.get(self, "state_reason")

