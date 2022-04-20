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

__all__ = [
    'ClusterEndpoint',
    'ClusterLoggingProperties',
    'ClusterParameterGroupParameter',
    'ClusterParameterGroupTag',
    'ClusterSecurityGroupTag',
    'ClusterSubnetGroupTag',
    'ClusterTag',
    'EndpointAccessNetworkInterface',
    'EndpointAccessVpcSecurityGroup',
    'EventSubscriptionTag',
    'ScheduledActionPauseClusterMessage',
    'ScheduledActionResizeClusterMessage',
    'ScheduledActionResumeClusterMessage',
    'ScheduledActionType',
    'VpcEndpointProperties',
]

@pulumi.output_type
class ClusterEndpoint(dict):
    def __init__(__self__, *,
                 address: Optional[str] = None,
                 port: Optional[str] = None):
        if address is not None:
            pulumi.set(__self__, "address", address)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def address(self) -> Optional[str]:
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def port(self) -> Optional[str]:
        return pulumi.get(self, "port")


@pulumi.output_type
class ClusterLoggingProperties(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "bucketName":
            suggest = "bucket_name"
        elif key == "s3KeyPrefix":
            suggest = "s3_key_prefix"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterLoggingProperties. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterLoggingProperties.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterLoggingProperties.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 bucket_name: str,
                 s3_key_prefix: Optional[str] = None):
        pulumi.set(__self__, "bucket_name", bucket_name)
        if s3_key_prefix is not None:
            pulumi.set(__self__, "s3_key_prefix", s3_key_prefix)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> str:
        return pulumi.get(self, "bucket_name")

    @property
    @pulumi.getter(name="s3KeyPrefix")
    def s3_key_prefix(self) -> Optional[str]:
        return pulumi.get(self, "s3_key_prefix")


@pulumi.output_type
class ClusterParameterGroupParameter(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "parameterName":
            suggest = "parameter_name"
        elif key == "parameterValue":
            suggest = "parameter_value"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterParameterGroupParameter. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterParameterGroupParameter.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterParameterGroupParameter.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 parameter_name: str,
                 parameter_value: str):
        pulumi.set(__self__, "parameter_name", parameter_name)
        pulumi.set(__self__, "parameter_value", parameter_value)

    @property
    @pulumi.getter(name="parameterName")
    def parameter_name(self) -> str:
        return pulumi.get(self, "parameter_name")

    @property
    @pulumi.getter(name="parameterValue")
    def parameter_value(self) -> str:
        return pulumi.get(self, "parameter_value")


@pulumi.output_type
class ClusterParameterGroupTag(dict):
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
class ClusterSecurityGroupTag(dict):
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
class ClusterSubnetGroupTag(dict):
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
class ClusterTag(dict):
    """
    A key-value pair to associate with a resource.
    """
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        A key-value pair to associate with a resource.
        :param str key: The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        :param str value: The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class EndpointAccessNetworkInterface(dict):
    """
    Describes a network interface.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "availabilityZone":
            suggest = "availability_zone"
        elif key == "networkInterfaceId":
            suggest = "network_interface_id"
        elif key == "privateIpAddress":
            suggest = "private_ip_address"
        elif key == "subnetId":
            suggest = "subnet_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in EndpointAccessNetworkInterface. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        EndpointAccessNetworkInterface.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        EndpointAccessNetworkInterface.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 availability_zone: Optional[str] = None,
                 network_interface_id: Optional[str] = None,
                 private_ip_address: Optional[str] = None,
                 subnet_id: Optional[str] = None):
        """
        Describes a network interface.
        :param str availability_zone: The Availability Zone.
        :param str network_interface_id: The network interface identifier.
        :param str private_ip_address: The IPv4 address of the network interface within the subnet.
        :param str subnet_id: The subnet identifier.
        """
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if network_interface_id is not None:
            pulumi.set(__self__, "network_interface_id", network_interface_id)
        if private_ip_address is not None:
            pulumi.set(__self__, "private_ip_address", private_ip_address)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[str]:
        """
        The Availability Zone.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> Optional[str]:
        """
        The network interface identifier.
        """
        return pulumi.get(self, "network_interface_id")

    @property
    @pulumi.getter(name="privateIpAddress")
    def private_ip_address(self) -> Optional[str]:
        """
        The IPv4 address of the network interface within the subnet.
        """
        return pulumi.get(self, "private_ip_address")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[str]:
        """
        The subnet identifier.
        """
        return pulumi.get(self, "subnet_id")


@pulumi.output_type
class EndpointAccessVpcSecurityGroup(dict):
    """
    Describes the members of a VPC security group.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "vpcSecurityGroupId":
            suggest = "vpc_security_group_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in EndpointAccessVpcSecurityGroup. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        EndpointAccessVpcSecurityGroup.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        EndpointAccessVpcSecurityGroup.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 status: Optional[str] = None,
                 vpc_security_group_id: Optional[str] = None):
        """
        Describes the members of a VPC security group.
        :param str status: The status of the VPC security group.
        :param str vpc_security_group_id: The identifier of the VPC security group.
        """
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vpc_security_group_id is not None:
            pulumi.set(__self__, "vpc_security_group_id", vpc_security_group_id)

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the VPC security group.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vpcSecurityGroupId")
    def vpc_security_group_id(self) -> Optional[str]:
        """
        The identifier of the VPC security group.
        """
        return pulumi.get(self, "vpc_security_group_id")


@pulumi.output_type
class EventSubscriptionTag(dict):
    """
    A key-value pair to associate with a resource.
    """
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        A key-value pair to associate with a resource.
        :param str key: The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        :param str value: The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class ScheduledActionPauseClusterMessage(dict):
    """
    Describes a pause cluster operation. For example, a scheduled action to run the `PauseCluster` API operation.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "clusterIdentifier":
            suggest = "cluster_identifier"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ScheduledActionPauseClusterMessage. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ScheduledActionPauseClusterMessage.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ScheduledActionPauseClusterMessage.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cluster_identifier: str):
        """
        Describes a pause cluster operation. For example, a scheduled action to run the `PauseCluster` API operation.
        """
        pulumi.set(__self__, "cluster_identifier", cluster_identifier)

    @property
    @pulumi.getter(name="clusterIdentifier")
    def cluster_identifier(self) -> str:
        return pulumi.get(self, "cluster_identifier")


@pulumi.output_type
class ScheduledActionResizeClusterMessage(dict):
    """
    Describes a resize cluster operation. For example, a scheduled action to run the `ResizeCluster` API operation.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "clusterIdentifier":
            suggest = "cluster_identifier"
        elif key == "clusterType":
            suggest = "cluster_type"
        elif key == "nodeType":
            suggest = "node_type"
        elif key == "numberOfNodes":
            suggest = "number_of_nodes"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ScheduledActionResizeClusterMessage. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ScheduledActionResizeClusterMessage.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ScheduledActionResizeClusterMessage.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cluster_identifier: str,
                 classic: Optional[bool] = None,
                 cluster_type: Optional[str] = None,
                 node_type: Optional[str] = None,
                 number_of_nodes: Optional[int] = None):
        """
        Describes a resize cluster operation. For example, a scheduled action to run the `ResizeCluster` API operation.
        """
        pulumi.set(__self__, "cluster_identifier", cluster_identifier)
        if classic is not None:
            pulumi.set(__self__, "classic", classic)
        if cluster_type is not None:
            pulumi.set(__self__, "cluster_type", cluster_type)
        if node_type is not None:
            pulumi.set(__self__, "node_type", node_type)
        if number_of_nodes is not None:
            pulumi.set(__self__, "number_of_nodes", number_of_nodes)

    @property
    @pulumi.getter(name="clusterIdentifier")
    def cluster_identifier(self) -> str:
        return pulumi.get(self, "cluster_identifier")

    @property
    @pulumi.getter
    def classic(self) -> Optional[bool]:
        return pulumi.get(self, "classic")

    @property
    @pulumi.getter(name="clusterType")
    def cluster_type(self) -> Optional[str]:
        return pulumi.get(self, "cluster_type")

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> Optional[str]:
        return pulumi.get(self, "node_type")

    @property
    @pulumi.getter(name="numberOfNodes")
    def number_of_nodes(self) -> Optional[int]:
        return pulumi.get(self, "number_of_nodes")


@pulumi.output_type
class ScheduledActionResumeClusterMessage(dict):
    """
    Describes a resume cluster operation. For example, a scheduled action to run the `ResumeCluster` API operation.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "clusterIdentifier":
            suggest = "cluster_identifier"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ScheduledActionResumeClusterMessage. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ScheduledActionResumeClusterMessage.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ScheduledActionResumeClusterMessage.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cluster_identifier: str):
        """
        Describes a resume cluster operation. For example, a scheduled action to run the `ResumeCluster` API operation.
        """
        pulumi.set(__self__, "cluster_identifier", cluster_identifier)

    @property
    @pulumi.getter(name="clusterIdentifier")
    def cluster_identifier(self) -> str:
        return pulumi.get(self, "cluster_identifier")


@pulumi.output_type
class ScheduledActionType(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "pauseCluster":
            suggest = "pause_cluster"
        elif key == "resizeCluster":
            suggest = "resize_cluster"
        elif key == "resumeCluster":
            suggest = "resume_cluster"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ScheduledActionType. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ScheduledActionType.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ScheduledActionType.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 pause_cluster: Optional['outputs.ScheduledActionPauseClusterMessage'] = None,
                 resize_cluster: Optional['outputs.ScheduledActionResizeClusterMessage'] = None,
                 resume_cluster: Optional['outputs.ScheduledActionResumeClusterMessage'] = None):
        if pause_cluster is not None:
            pulumi.set(__self__, "pause_cluster", pause_cluster)
        if resize_cluster is not None:
            pulumi.set(__self__, "resize_cluster", resize_cluster)
        if resume_cluster is not None:
            pulumi.set(__self__, "resume_cluster", resume_cluster)

    @property
    @pulumi.getter(name="pauseCluster")
    def pause_cluster(self) -> Optional['outputs.ScheduledActionPauseClusterMessage']:
        return pulumi.get(self, "pause_cluster")

    @property
    @pulumi.getter(name="resizeCluster")
    def resize_cluster(self) -> Optional['outputs.ScheduledActionResizeClusterMessage']:
        return pulumi.get(self, "resize_cluster")

    @property
    @pulumi.getter(name="resumeCluster")
    def resume_cluster(self) -> Optional['outputs.ScheduledActionResumeClusterMessage']:
        return pulumi.get(self, "resume_cluster")


@pulumi.output_type
class VpcEndpointProperties(dict):
    """
    The connection endpoint for connecting to an Amazon Redshift cluster through the proxy.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "networkInterfaces":
            suggest = "network_interfaces"
        elif key == "vpcEndpointId":
            suggest = "vpc_endpoint_id"
        elif key == "vpcId":
            suggest = "vpc_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VpcEndpointProperties. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VpcEndpointProperties.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VpcEndpointProperties.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 network_interfaces: Optional[Sequence['outputs.EndpointAccessNetworkInterface']] = None,
                 vpc_endpoint_id: Optional[str] = None,
                 vpc_id: Optional[str] = None):
        """
        The connection endpoint for connecting to an Amazon Redshift cluster through the proxy.
        :param Sequence['EndpointAccessNetworkInterface'] network_interfaces: One or more network interfaces of the endpoint. Also known as an interface endpoint.
        :param str vpc_endpoint_id: The connection endpoint ID for connecting an Amazon Redshift cluster through the proxy.
        :param str vpc_id: The VPC identifier that the endpoint is associated.
        """
        if network_interfaces is not None:
            pulumi.set(__self__, "network_interfaces", network_interfaces)
        if vpc_endpoint_id is not None:
            pulumi.set(__self__, "vpc_endpoint_id", vpc_endpoint_id)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="networkInterfaces")
    def network_interfaces(self) -> Optional[Sequence['outputs.EndpointAccessNetworkInterface']]:
        """
        One or more network interfaces of the endpoint. Also known as an interface endpoint.
        """
        return pulumi.get(self, "network_interfaces")

    @property
    @pulumi.getter(name="vpcEndpointId")
    def vpc_endpoint_id(self) -> Optional[str]:
        """
        The connection endpoint ID for connecting an Amazon Redshift cluster through the proxy.
        """
        return pulumi.get(self, "vpc_endpoint_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        """
        The VPC identifier that the endpoint is associated.
        """
        return pulumi.get(self, "vpc_id")


