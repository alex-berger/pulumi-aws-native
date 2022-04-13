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

__all__ = ['ConnectorArgs', 'Connector']

@pulumi.input_type
class ConnectorArgs:
    def __init__(__self__, *,
                 capacity: pulumi.Input['ConnectorCapacityArgs'],
                 connector_configuration: Any,
                 kafka_cluster: pulumi.Input['ConnectorKafkaClusterArgs'],
                 kafka_cluster_client_authentication: pulumi.Input['ConnectorKafkaClusterClientAuthenticationArgs'],
                 kafka_cluster_encryption_in_transit: pulumi.Input['ConnectorKafkaClusterEncryptionInTransitArgs'],
                 kafka_connect_version: pulumi.Input[str],
                 plugins: pulumi.Input[Sequence[pulumi.Input['ConnectorPluginArgs']]],
                 service_execution_role_arn: pulumi.Input[str],
                 connector_description: Optional[pulumi.Input[str]] = None,
                 connector_name: Optional[pulumi.Input[str]] = None,
                 log_delivery: Optional[pulumi.Input['ConnectorLogDeliveryArgs']] = None,
                 worker_configuration: Optional[pulumi.Input['ConnectorWorkerConfigurationArgs']] = None):
        """
        The set of arguments for constructing a Connector resource.
        :param Any connector_configuration: The configuration for the connector.
        :param pulumi.Input[str] kafka_connect_version: The version of Kafka Connect. It has to be compatible with both the Kafka cluster's version and the plugins.
        :param pulumi.Input[Sequence[pulumi.Input['ConnectorPluginArgs']]] plugins: List of plugins to use with the connector.
        :param pulumi.Input[str] service_execution_role_arn: The Amazon Resource Name (ARN) of the IAM role used by the connector to access Amazon S3 objects and other external resources.
        :param pulumi.Input[str] connector_description: A summary description of the connector.
        :param pulumi.Input[str] connector_name: The name of the connector.
        """
        pulumi.set(__self__, "capacity", capacity)
        pulumi.set(__self__, "connector_configuration", connector_configuration)
        pulumi.set(__self__, "kafka_cluster", kafka_cluster)
        pulumi.set(__self__, "kafka_cluster_client_authentication", kafka_cluster_client_authentication)
        pulumi.set(__self__, "kafka_cluster_encryption_in_transit", kafka_cluster_encryption_in_transit)
        pulumi.set(__self__, "kafka_connect_version", kafka_connect_version)
        pulumi.set(__self__, "plugins", plugins)
        pulumi.set(__self__, "service_execution_role_arn", service_execution_role_arn)
        if connector_description is not None:
            pulumi.set(__self__, "connector_description", connector_description)
        if connector_name is not None:
            pulumi.set(__self__, "connector_name", connector_name)
        if log_delivery is not None:
            pulumi.set(__self__, "log_delivery", log_delivery)
        if worker_configuration is not None:
            pulumi.set(__self__, "worker_configuration", worker_configuration)

    @property
    @pulumi.getter
    def capacity(self) -> pulumi.Input['ConnectorCapacityArgs']:
        return pulumi.get(self, "capacity")

    @capacity.setter
    def capacity(self, value: pulumi.Input['ConnectorCapacityArgs']):
        pulumi.set(self, "capacity", value)

    @property
    @pulumi.getter(name="connectorConfiguration")
    def connector_configuration(self) -> Any:
        """
        The configuration for the connector.
        """
        return pulumi.get(self, "connector_configuration")

    @connector_configuration.setter
    def connector_configuration(self, value: Any):
        pulumi.set(self, "connector_configuration", value)

    @property
    @pulumi.getter(name="kafkaCluster")
    def kafka_cluster(self) -> pulumi.Input['ConnectorKafkaClusterArgs']:
        return pulumi.get(self, "kafka_cluster")

    @kafka_cluster.setter
    def kafka_cluster(self, value: pulumi.Input['ConnectorKafkaClusterArgs']):
        pulumi.set(self, "kafka_cluster", value)

    @property
    @pulumi.getter(name="kafkaClusterClientAuthentication")
    def kafka_cluster_client_authentication(self) -> pulumi.Input['ConnectorKafkaClusterClientAuthenticationArgs']:
        return pulumi.get(self, "kafka_cluster_client_authentication")

    @kafka_cluster_client_authentication.setter
    def kafka_cluster_client_authentication(self, value: pulumi.Input['ConnectorKafkaClusterClientAuthenticationArgs']):
        pulumi.set(self, "kafka_cluster_client_authentication", value)

    @property
    @pulumi.getter(name="kafkaClusterEncryptionInTransit")
    def kafka_cluster_encryption_in_transit(self) -> pulumi.Input['ConnectorKafkaClusterEncryptionInTransitArgs']:
        return pulumi.get(self, "kafka_cluster_encryption_in_transit")

    @kafka_cluster_encryption_in_transit.setter
    def kafka_cluster_encryption_in_transit(self, value: pulumi.Input['ConnectorKafkaClusterEncryptionInTransitArgs']):
        pulumi.set(self, "kafka_cluster_encryption_in_transit", value)

    @property
    @pulumi.getter(name="kafkaConnectVersion")
    def kafka_connect_version(self) -> pulumi.Input[str]:
        """
        The version of Kafka Connect. It has to be compatible with both the Kafka cluster's version and the plugins.
        """
        return pulumi.get(self, "kafka_connect_version")

    @kafka_connect_version.setter
    def kafka_connect_version(self, value: pulumi.Input[str]):
        pulumi.set(self, "kafka_connect_version", value)

    @property
    @pulumi.getter
    def plugins(self) -> pulumi.Input[Sequence[pulumi.Input['ConnectorPluginArgs']]]:
        """
        List of plugins to use with the connector.
        """
        return pulumi.get(self, "plugins")

    @plugins.setter
    def plugins(self, value: pulumi.Input[Sequence[pulumi.Input['ConnectorPluginArgs']]]):
        pulumi.set(self, "plugins", value)

    @property
    @pulumi.getter(name="serviceExecutionRoleArn")
    def service_execution_role_arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN) of the IAM role used by the connector to access Amazon S3 objects and other external resources.
        """
        return pulumi.get(self, "service_execution_role_arn")

    @service_execution_role_arn.setter
    def service_execution_role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_execution_role_arn", value)

    @property
    @pulumi.getter(name="connectorDescription")
    def connector_description(self) -> Optional[pulumi.Input[str]]:
        """
        A summary description of the connector.
        """
        return pulumi.get(self, "connector_description")

    @connector_description.setter
    def connector_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connector_description", value)

    @property
    @pulumi.getter(name="connectorName")
    def connector_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the connector.
        """
        return pulumi.get(self, "connector_name")

    @connector_name.setter
    def connector_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connector_name", value)

    @property
    @pulumi.getter(name="logDelivery")
    def log_delivery(self) -> Optional[pulumi.Input['ConnectorLogDeliveryArgs']]:
        return pulumi.get(self, "log_delivery")

    @log_delivery.setter
    def log_delivery(self, value: Optional[pulumi.Input['ConnectorLogDeliveryArgs']]):
        pulumi.set(self, "log_delivery", value)

    @property
    @pulumi.getter(name="workerConfiguration")
    def worker_configuration(self) -> Optional[pulumi.Input['ConnectorWorkerConfigurationArgs']]:
        return pulumi.get(self, "worker_configuration")

    @worker_configuration.setter
    def worker_configuration(self, value: Optional[pulumi.Input['ConnectorWorkerConfigurationArgs']]):
        pulumi.set(self, "worker_configuration", value)


class Connector(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capacity: Optional[pulumi.Input[pulumi.InputType['ConnectorCapacityArgs']]] = None,
                 connector_configuration: Optional[Any] = None,
                 connector_description: Optional[pulumi.Input[str]] = None,
                 connector_name: Optional[pulumi.Input[str]] = None,
                 kafka_cluster: Optional[pulumi.Input[pulumi.InputType['ConnectorKafkaClusterArgs']]] = None,
                 kafka_cluster_client_authentication: Optional[pulumi.Input[pulumi.InputType['ConnectorKafkaClusterClientAuthenticationArgs']]] = None,
                 kafka_cluster_encryption_in_transit: Optional[pulumi.Input[pulumi.InputType['ConnectorKafkaClusterEncryptionInTransitArgs']]] = None,
                 kafka_connect_version: Optional[pulumi.Input[str]] = None,
                 log_delivery: Optional[pulumi.Input[pulumi.InputType['ConnectorLogDeliveryArgs']]] = None,
                 plugins: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ConnectorPluginArgs']]]]] = None,
                 service_execution_role_arn: Optional[pulumi.Input[str]] = None,
                 worker_configuration: Optional[pulumi.Input[pulumi.InputType['ConnectorWorkerConfigurationArgs']]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::KafkaConnect::Connector

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param Any connector_configuration: The configuration for the connector.
        :param pulumi.Input[str] connector_description: A summary description of the connector.
        :param pulumi.Input[str] connector_name: The name of the connector.
        :param pulumi.Input[str] kafka_connect_version: The version of Kafka Connect. It has to be compatible with both the Kafka cluster's version and the plugins.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ConnectorPluginArgs']]]] plugins: List of plugins to use with the connector.
        :param pulumi.Input[str] service_execution_role_arn: The Amazon Resource Name (ARN) of the IAM role used by the connector to access Amazon S3 objects and other external resources.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConnectorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::KafkaConnect::Connector

        :param str resource_name: The name of the resource.
        :param ConnectorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConnectorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capacity: Optional[pulumi.Input[pulumi.InputType['ConnectorCapacityArgs']]] = None,
                 connector_configuration: Optional[Any] = None,
                 connector_description: Optional[pulumi.Input[str]] = None,
                 connector_name: Optional[pulumi.Input[str]] = None,
                 kafka_cluster: Optional[pulumi.Input[pulumi.InputType['ConnectorKafkaClusterArgs']]] = None,
                 kafka_cluster_client_authentication: Optional[pulumi.Input[pulumi.InputType['ConnectorKafkaClusterClientAuthenticationArgs']]] = None,
                 kafka_cluster_encryption_in_transit: Optional[pulumi.Input[pulumi.InputType['ConnectorKafkaClusterEncryptionInTransitArgs']]] = None,
                 kafka_connect_version: Optional[pulumi.Input[str]] = None,
                 log_delivery: Optional[pulumi.Input[pulumi.InputType['ConnectorLogDeliveryArgs']]] = None,
                 plugins: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ConnectorPluginArgs']]]]] = None,
                 service_execution_role_arn: Optional[pulumi.Input[str]] = None,
                 worker_configuration: Optional[pulumi.Input[pulumi.InputType['ConnectorWorkerConfigurationArgs']]] = None,
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
            __props__ = ConnectorArgs.__new__(ConnectorArgs)

            if capacity is None and not opts.urn:
                raise TypeError("Missing required property 'capacity'")
            __props__.__dict__["capacity"] = capacity
            if connector_configuration is None and not opts.urn:
                raise TypeError("Missing required property 'connector_configuration'")
            __props__.__dict__["connector_configuration"] = connector_configuration
            __props__.__dict__["connector_description"] = connector_description
            __props__.__dict__["connector_name"] = connector_name
            if kafka_cluster is None and not opts.urn:
                raise TypeError("Missing required property 'kafka_cluster'")
            __props__.__dict__["kafka_cluster"] = kafka_cluster
            if kafka_cluster_client_authentication is None and not opts.urn:
                raise TypeError("Missing required property 'kafka_cluster_client_authentication'")
            __props__.__dict__["kafka_cluster_client_authentication"] = kafka_cluster_client_authentication
            if kafka_cluster_encryption_in_transit is None and not opts.urn:
                raise TypeError("Missing required property 'kafka_cluster_encryption_in_transit'")
            __props__.__dict__["kafka_cluster_encryption_in_transit"] = kafka_cluster_encryption_in_transit
            if kafka_connect_version is None and not opts.urn:
                raise TypeError("Missing required property 'kafka_connect_version'")
            __props__.__dict__["kafka_connect_version"] = kafka_connect_version
            __props__.__dict__["log_delivery"] = log_delivery
            if plugins is None and not opts.urn:
                raise TypeError("Missing required property 'plugins'")
            __props__.__dict__["plugins"] = plugins
            if service_execution_role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'service_execution_role_arn'")
            __props__.__dict__["service_execution_role_arn"] = service_execution_role_arn
            __props__.__dict__["worker_configuration"] = worker_configuration
            __props__.__dict__["connector_arn"] = None
        super(Connector, __self__).__init__(
            'aws-native:kafkaconnect:Connector',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Connector':
        """
        Get an existing Connector resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ConnectorArgs.__new__(ConnectorArgs)

        __props__.__dict__["capacity"] = None
        __props__.__dict__["connector_arn"] = None
        __props__.__dict__["connector_configuration"] = None
        __props__.__dict__["connector_description"] = None
        __props__.__dict__["connector_name"] = None
        __props__.__dict__["kafka_cluster"] = None
        __props__.__dict__["kafka_cluster_client_authentication"] = None
        __props__.__dict__["kafka_cluster_encryption_in_transit"] = None
        __props__.__dict__["kafka_connect_version"] = None
        __props__.__dict__["log_delivery"] = None
        __props__.__dict__["plugins"] = None
        __props__.__dict__["service_execution_role_arn"] = None
        __props__.__dict__["worker_configuration"] = None
        return Connector(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def capacity(self) -> pulumi.Output['outputs.ConnectorCapacity']:
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter(name="connectorArn")
    def connector_arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name for the created Connector.
        """
        return pulumi.get(self, "connector_arn")

    @property
    @pulumi.getter(name="connectorConfiguration")
    def connector_configuration(self) -> pulumi.Output[Any]:
        """
        The configuration for the connector.
        """
        return pulumi.get(self, "connector_configuration")

    @property
    @pulumi.getter(name="connectorDescription")
    def connector_description(self) -> pulumi.Output[Optional[str]]:
        """
        A summary description of the connector.
        """
        return pulumi.get(self, "connector_description")

    @property
    @pulumi.getter(name="connectorName")
    def connector_name(self) -> pulumi.Output[str]:
        """
        The name of the connector.
        """
        return pulumi.get(self, "connector_name")

    @property
    @pulumi.getter(name="kafkaCluster")
    def kafka_cluster(self) -> pulumi.Output['outputs.ConnectorKafkaCluster']:
        return pulumi.get(self, "kafka_cluster")

    @property
    @pulumi.getter(name="kafkaClusterClientAuthentication")
    def kafka_cluster_client_authentication(self) -> pulumi.Output['outputs.ConnectorKafkaClusterClientAuthentication']:
        return pulumi.get(self, "kafka_cluster_client_authentication")

    @property
    @pulumi.getter(name="kafkaClusterEncryptionInTransit")
    def kafka_cluster_encryption_in_transit(self) -> pulumi.Output['outputs.ConnectorKafkaClusterEncryptionInTransit']:
        return pulumi.get(self, "kafka_cluster_encryption_in_transit")

    @property
    @pulumi.getter(name="kafkaConnectVersion")
    def kafka_connect_version(self) -> pulumi.Output[str]:
        """
        The version of Kafka Connect. It has to be compatible with both the Kafka cluster's version and the plugins.
        """
        return pulumi.get(self, "kafka_connect_version")

    @property
    @pulumi.getter(name="logDelivery")
    def log_delivery(self) -> pulumi.Output[Optional['outputs.ConnectorLogDelivery']]:
        return pulumi.get(self, "log_delivery")

    @property
    @pulumi.getter
    def plugins(self) -> pulumi.Output[Sequence['outputs.ConnectorPlugin']]:
        """
        List of plugins to use with the connector.
        """
        return pulumi.get(self, "plugins")

    @property
    @pulumi.getter(name="serviceExecutionRoleArn")
    def service_execution_role_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the IAM role used by the connector to access Amazon S3 objects and other external resources.
        """
        return pulumi.get(self, "service_execution_role_arn")

    @property
    @pulumi.getter(name="workerConfiguration")
    def worker_configuration(self) -> pulumi.Output[Optional['outputs.ConnectorWorkerConfiguration']]:
        return pulumi.get(self, "worker_configuration")

