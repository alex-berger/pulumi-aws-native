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

__all__ = ['ClusterArgs', 'Cluster']

@pulumi.input_type
class ClusterArgs:
    def __init__(__self__, *,
                 broker_node_group_info: pulumi.Input['ClusterBrokerNodeGroupInfoArgs'],
                 cluster_name: pulumi.Input[str],
                 kafka_version: pulumi.Input[str],
                 number_of_broker_nodes: pulumi.Input[int],
                 client_authentication: Optional[pulumi.Input['ClusterClientAuthenticationArgs']] = None,
                 configuration_info: Optional[pulumi.Input['ClusterConfigurationInfoArgs']] = None,
                 encryption_info: Optional[pulumi.Input['ClusterEncryptionInfoArgs']] = None,
                 enhanced_monitoring: Optional[pulumi.Input[str]] = None,
                 logging_info: Optional[pulumi.Input['ClusterLoggingInfoArgs']] = None,
                 open_monitoring: Optional[pulumi.Input['ClusterOpenMonitoringArgs']] = None,
                 tags: Optional[Any] = None):
        """
        The set of arguments for constructing a Cluster resource.
        """
        pulumi.set(__self__, "broker_node_group_info", broker_node_group_info)
        pulumi.set(__self__, "cluster_name", cluster_name)
        pulumi.set(__self__, "kafka_version", kafka_version)
        pulumi.set(__self__, "number_of_broker_nodes", number_of_broker_nodes)
        if client_authentication is not None:
            pulumi.set(__self__, "client_authentication", client_authentication)
        if configuration_info is not None:
            pulumi.set(__self__, "configuration_info", configuration_info)
        if encryption_info is not None:
            pulumi.set(__self__, "encryption_info", encryption_info)
        if enhanced_monitoring is not None:
            pulumi.set(__self__, "enhanced_monitoring", enhanced_monitoring)
        if logging_info is not None:
            pulumi.set(__self__, "logging_info", logging_info)
        if open_monitoring is not None:
            pulumi.set(__self__, "open_monitoring", open_monitoring)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="brokerNodeGroupInfo")
    def broker_node_group_info(self) -> pulumi.Input['ClusterBrokerNodeGroupInfoArgs']:
        return pulumi.get(self, "broker_node_group_info")

    @broker_node_group_info.setter
    def broker_node_group_info(self, value: pulumi.Input['ClusterBrokerNodeGroupInfoArgs']):
        pulumi.set(self, "broker_node_group_info", value)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter(name="kafkaVersion")
    def kafka_version(self) -> pulumi.Input[str]:
        return pulumi.get(self, "kafka_version")

    @kafka_version.setter
    def kafka_version(self, value: pulumi.Input[str]):
        pulumi.set(self, "kafka_version", value)

    @property
    @pulumi.getter(name="numberOfBrokerNodes")
    def number_of_broker_nodes(self) -> pulumi.Input[int]:
        return pulumi.get(self, "number_of_broker_nodes")

    @number_of_broker_nodes.setter
    def number_of_broker_nodes(self, value: pulumi.Input[int]):
        pulumi.set(self, "number_of_broker_nodes", value)

    @property
    @pulumi.getter(name="clientAuthentication")
    def client_authentication(self) -> Optional[pulumi.Input['ClusterClientAuthenticationArgs']]:
        return pulumi.get(self, "client_authentication")

    @client_authentication.setter
    def client_authentication(self, value: Optional[pulumi.Input['ClusterClientAuthenticationArgs']]):
        pulumi.set(self, "client_authentication", value)

    @property
    @pulumi.getter(name="configurationInfo")
    def configuration_info(self) -> Optional[pulumi.Input['ClusterConfigurationInfoArgs']]:
        return pulumi.get(self, "configuration_info")

    @configuration_info.setter
    def configuration_info(self, value: Optional[pulumi.Input['ClusterConfigurationInfoArgs']]):
        pulumi.set(self, "configuration_info", value)

    @property
    @pulumi.getter(name="encryptionInfo")
    def encryption_info(self) -> Optional[pulumi.Input['ClusterEncryptionInfoArgs']]:
        return pulumi.get(self, "encryption_info")

    @encryption_info.setter
    def encryption_info(self, value: Optional[pulumi.Input['ClusterEncryptionInfoArgs']]):
        pulumi.set(self, "encryption_info", value)

    @property
    @pulumi.getter(name="enhancedMonitoring")
    def enhanced_monitoring(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "enhanced_monitoring")

    @enhanced_monitoring.setter
    def enhanced_monitoring(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enhanced_monitoring", value)

    @property
    @pulumi.getter(name="loggingInfo")
    def logging_info(self) -> Optional[pulumi.Input['ClusterLoggingInfoArgs']]:
        return pulumi.get(self, "logging_info")

    @logging_info.setter
    def logging_info(self, value: Optional[pulumi.Input['ClusterLoggingInfoArgs']]):
        pulumi.set(self, "logging_info", value)

    @property
    @pulumi.getter(name="openMonitoring")
    def open_monitoring(self) -> Optional[pulumi.Input['ClusterOpenMonitoringArgs']]:
        return pulumi.get(self, "open_monitoring")

    @open_monitoring.setter
    def open_monitoring(self, value: Optional[pulumi.Input['ClusterOpenMonitoringArgs']]):
        pulumi.set(self, "open_monitoring", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[Any]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[Any]):
        pulumi.set(self, "tags", value)


warnings.warn("""Cluster is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class Cluster(pulumi.CustomResource):
    warnings.warn("""Cluster is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 broker_node_group_info: Optional[pulumi.Input[pulumi.InputType['ClusterBrokerNodeGroupInfoArgs']]] = None,
                 client_authentication: Optional[pulumi.Input[pulumi.InputType['ClusterClientAuthenticationArgs']]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 configuration_info: Optional[pulumi.Input[pulumi.InputType['ClusterConfigurationInfoArgs']]] = None,
                 encryption_info: Optional[pulumi.Input[pulumi.InputType['ClusterEncryptionInfoArgs']]] = None,
                 enhanced_monitoring: Optional[pulumi.Input[str]] = None,
                 kafka_version: Optional[pulumi.Input[str]] = None,
                 logging_info: Optional[pulumi.Input[pulumi.InputType['ClusterLoggingInfoArgs']]] = None,
                 number_of_broker_nodes: Optional[pulumi.Input[int]] = None,
                 open_monitoring: Optional[pulumi.Input[pulumi.InputType['ClusterOpenMonitoringArgs']]] = None,
                 tags: Optional[Any] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::MSK::Cluster

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::MSK::Cluster

        :param str resource_name: The name of the resource.
        :param ClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 broker_node_group_info: Optional[pulumi.Input[pulumi.InputType['ClusterBrokerNodeGroupInfoArgs']]] = None,
                 client_authentication: Optional[pulumi.Input[pulumi.InputType['ClusterClientAuthenticationArgs']]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 configuration_info: Optional[pulumi.Input[pulumi.InputType['ClusterConfigurationInfoArgs']]] = None,
                 encryption_info: Optional[pulumi.Input[pulumi.InputType['ClusterEncryptionInfoArgs']]] = None,
                 enhanced_monitoring: Optional[pulumi.Input[str]] = None,
                 kafka_version: Optional[pulumi.Input[str]] = None,
                 logging_info: Optional[pulumi.Input[pulumi.InputType['ClusterLoggingInfoArgs']]] = None,
                 number_of_broker_nodes: Optional[pulumi.Input[int]] = None,
                 open_monitoring: Optional[pulumi.Input[pulumi.InputType['ClusterOpenMonitoringArgs']]] = None,
                 tags: Optional[Any] = None,
                 __props__=None):
        pulumi.log.warn("""Cluster is deprecated: Cluster is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClusterArgs.__new__(ClusterArgs)

            if broker_node_group_info is None and not opts.urn:
                raise TypeError("Missing required property 'broker_node_group_info'")
            __props__.__dict__["broker_node_group_info"] = broker_node_group_info
            __props__.__dict__["client_authentication"] = client_authentication
            if cluster_name is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_name'")
            __props__.__dict__["cluster_name"] = cluster_name
            __props__.__dict__["configuration_info"] = configuration_info
            __props__.__dict__["encryption_info"] = encryption_info
            __props__.__dict__["enhanced_monitoring"] = enhanced_monitoring
            if kafka_version is None and not opts.urn:
                raise TypeError("Missing required property 'kafka_version'")
            __props__.__dict__["kafka_version"] = kafka_version
            __props__.__dict__["logging_info"] = logging_info
            if number_of_broker_nodes is None and not opts.urn:
                raise TypeError("Missing required property 'number_of_broker_nodes'")
            __props__.__dict__["number_of_broker_nodes"] = number_of_broker_nodes
            __props__.__dict__["open_monitoring"] = open_monitoring
            __props__.__dict__["tags"] = tags
        super(Cluster, __self__).__init__(
            'aws-native:msk:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Cluster':
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ClusterArgs.__new__(ClusterArgs)

        __props__.__dict__["broker_node_group_info"] = None
        __props__.__dict__["client_authentication"] = None
        __props__.__dict__["cluster_name"] = None
        __props__.__dict__["configuration_info"] = None
        __props__.__dict__["encryption_info"] = None
        __props__.__dict__["enhanced_monitoring"] = None
        __props__.__dict__["kafka_version"] = None
        __props__.__dict__["logging_info"] = None
        __props__.__dict__["number_of_broker_nodes"] = None
        __props__.__dict__["open_monitoring"] = None
        __props__.__dict__["tags"] = None
        return Cluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="brokerNodeGroupInfo")
    def broker_node_group_info(self) -> pulumi.Output['outputs.ClusterBrokerNodeGroupInfo']:
        return pulumi.get(self, "broker_node_group_info")

    @property
    @pulumi.getter(name="clientAuthentication")
    def client_authentication(self) -> pulumi.Output[Optional['outputs.ClusterClientAuthentication']]:
        return pulumi.get(self, "client_authentication")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="configurationInfo")
    def configuration_info(self) -> pulumi.Output[Optional['outputs.ClusterConfigurationInfo']]:
        return pulumi.get(self, "configuration_info")

    @property
    @pulumi.getter(name="encryptionInfo")
    def encryption_info(self) -> pulumi.Output[Optional['outputs.ClusterEncryptionInfo']]:
        return pulumi.get(self, "encryption_info")

    @property
    @pulumi.getter(name="enhancedMonitoring")
    def enhanced_monitoring(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "enhanced_monitoring")

    @property
    @pulumi.getter(name="kafkaVersion")
    def kafka_version(self) -> pulumi.Output[str]:
        return pulumi.get(self, "kafka_version")

    @property
    @pulumi.getter(name="loggingInfo")
    def logging_info(self) -> pulumi.Output[Optional['outputs.ClusterLoggingInfo']]:
        return pulumi.get(self, "logging_info")

    @property
    @pulumi.getter(name="numberOfBrokerNodes")
    def number_of_broker_nodes(self) -> pulumi.Output[int]:
        return pulumi.get(self, "number_of_broker_nodes")

    @property
    @pulumi.getter(name="openMonitoring")
    def open_monitoring(self) -> pulumi.Output[Optional['outputs.ClusterOpenMonitoring']]:
        return pulumi.get(self, "open_monitoring")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Any]]:
        return pulumi.get(self, "tags")
