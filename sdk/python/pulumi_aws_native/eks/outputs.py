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
    'AddonTag',
    'ClusterEncryptionConfig',
    'ClusterKubernetesNetworkConfig',
    'ClusterLogging',
    'ClusterProvider',
    'ClusterResourcesVpcConfig',
    'ClusterTag',
    'FargateProfileLabel',
    'FargateProfileSelector',
    'FargateProfileTag',
    'IdentityProviderConfigOidcIdentityProviderConfig',
    'IdentityProviderConfigRequiredClaim',
    'IdentityProviderConfigTag',
    'NodegroupLaunchTemplateSpecification',
    'NodegroupRemoteAccess',
    'NodegroupScalingConfig',
    'NodegroupTaint',
    'NodegroupUpdateConfig',
]

@pulumi.output_type
class AddonTag(dict):
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
class ClusterEncryptionConfig(dict):
    """
    The encryption configuration for the cluster
    """
    def __init__(__self__, *,
                 provider: Optional['outputs.ClusterProvider'] = None,
                 resources: Optional[Sequence[str]] = None):
        """
        The encryption configuration for the cluster
        :param 'ClusterProvider' provider: The encryption provider for the cluster.
        :param Sequence[str] resources: Specifies the resources to be encrypted. The only supported value is "secrets".
        """
        if provider is not None:
            pulumi.set(__self__, "provider", provider)
        if resources is not None:
            pulumi.set(__self__, "resources", resources)

    @property
    @pulumi.getter
    def provider(self) -> Optional['outputs.ClusterProvider']:
        """
        The encryption provider for the cluster.
        """
        return pulumi.get(self, "provider")

    @property
    @pulumi.getter
    def resources(self) -> Optional[Sequence[str]]:
        """
        Specifies the resources to be encrypted. The only supported value is "secrets".
        """
        return pulumi.get(self, "resources")


@pulumi.output_type
class ClusterKubernetesNetworkConfig(dict):
    """
    The Kubernetes network configuration for the cluster.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "ipFamily":
            suggest = "ip_family"
        elif key == "serviceIpv4Cidr":
            suggest = "service_ipv4_cidr"
        elif key == "serviceIpv6Cidr":
            suggest = "service_ipv6_cidr"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterKubernetesNetworkConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterKubernetesNetworkConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterKubernetesNetworkConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 ip_family: Optional['ClusterKubernetesNetworkConfigIpFamily'] = None,
                 service_ipv4_cidr: Optional[str] = None,
                 service_ipv6_cidr: Optional[str] = None):
        """
        The Kubernetes network configuration for the cluster.
        :param 'ClusterKubernetesNetworkConfigIpFamily' ip_family: Ipv4 or Ipv6. You can only specify ipv6 for 1.21 and later clusters that use version 1.10.1 or later of the Amazon VPC CNI add-on
        :param str service_ipv4_cidr: The CIDR block to assign Kubernetes service IP addresses from. If you don't specify a block, Kubernetes assigns addresses from either the 10.100.0.0/16 or 172.20.0.0/16 CIDR blocks. We recommend that you specify a block that does not overlap with resources in other networks that are peered or connected to your VPC. 
        :param str service_ipv6_cidr: The CIDR block to assign Kubernetes service IP addresses from.
        """
        if ip_family is not None:
            pulumi.set(__self__, "ip_family", ip_family)
        if service_ipv4_cidr is not None:
            pulumi.set(__self__, "service_ipv4_cidr", service_ipv4_cidr)
        if service_ipv6_cidr is not None:
            pulumi.set(__self__, "service_ipv6_cidr", service_ipv6_cidr)

    @property
    @pulumi.getter(name="ipFamily")
    def ip_family(self) -> Optional['ClusterKubernetesNetworkConfigIpFamily']:
        """
        Ipv4 or Ipv6. You can only specify ipv6 for 1.21 and later clusters that use version 1.10.1 or later of the Amazon VPC CNI add-on
        """
        return pulumi.get(self, "ip_family")

    @property
    @pulumi.getter(name="serviceIpv4Cidr")
    def service_ipv4_cidr(self) -> Optional[str]:
        """
        The CIDR block to assign Kubernetes service IP addresses from. If you don't specify a block, Kubernetes assigns addresses from either the 10.100.0.0/16 or 172.20.0.0/16 CIDR blocks. We recommend that you specify a block that does not overlap with resources in other networks that are peered or connected to your VPC. 
        """
        return pulumi.get(self, "service_ipv4_cidr")

    @property
    @pulumi.getter(name="serviceIpv6Cidr")
    def service_ipv6_cidr(self) -> Optional[str]:
        """
        The CIDR block to assign Kubernetes service IP addresses from.
        """
        return pulumi.get(self, "service_ipv6_cidr")


@pulumi.output_type
class ClusterLogging(dict):
    """
    Enable exporting the Kubernetes control plane logs for your cluster to CloudWatch Logs based on log types. By default, cluster control plane logs aren't exported to CloudWatch Logs.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "clusterLogging":
            suggest = "cluster_logging"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterLogging. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterLogging.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterLogging.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cluster_logging: Optional['outputs.ClusterLogging'] = None):
        """
        Enable exporting the Kubernetes control plane logs for your cluster to CloudWatch Logs based on log types. By default, cluster control plane logs aren't exported to CloudWatch Logs.
        :param 'ClusterLogging' cluster_logging: The cluster control plane logging configuration for your cluster. 
        """
        if cluster_logging is not None:
            pulumi.set(__self__, "cluster_logging", cluster_logging)

    @property
    @pulumi.getter(name="clusterLogging")
    def cluster_logging(self) -> Optional['outputs.ClusterLogging']:
        """
        The cluster control plane logging configuration for your cluster. 
        """
        return pulumi.get(self, "cluster_logging")


@pulumi.output_type
class ClusterProvider(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "keyArn":
            suggest = "key_arn"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterProvider. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterProvider.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterProvider.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 key_arn: Optional[str] = None):
        """
        :param str key_arn: Amazon Resource Name (ARN) or alias of the KMS key. The KMS key must be symmetric, created in the same region as the cluster, and if the KMS key was created in a different account, the user must have access to the KMS key.
        """
        if key_arn is not None:
            pulumi.set(__self__, "key_arn", key_arn)

    @property
    @pulumi.getter(name="keyArn")
    def key_arn(self) -> Optional[str]:
        """
        Amazon Resource Name (ARN) or alias of the KMS key. The KMS key must be symmetric, created in the same region as the cluster, and if the KMS key was created in a different account, the user must have access to the KMS key.
        """
        return pulumi.get(self, "key_arn")


@pulumi.output_type
class ClusterResourcesVpcConfig(dict):
    """
    An object representing the VPC configuration to use for an Amazon EKS cluster.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "subnetIds":
            suggest = "subnet_ids"
        elif key == "endpointPrivateAccess":
            suggest = "endpoint_private_access"
        elif key == "endpointPublicAccess":
            suggest = "endpoint_public_access"
        elif key == "publicAccessCidrs":
            suggest = "public_access_cidrs"
        elif key == "securityGroupIds":
            suggest = "security_group_ids"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterResourcesVpcConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterResourcesVpcConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterResourcesVpcConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 subnet_ids: Sequence[str],
                 endpoint_private_access: Optional[bool] = None,
                 endpoint_public_access: Optional[bool] = None,
                 public_access_cidrs: Optional[Sequence[str]] = None,
                 security_group_ids: Optional[Sequence[str]] = None):
        """
        An object representing the VPC configuration to use for an Amazon EKS cluster.
        :param Sequence[str] subnet_ids: Specify subnets for your Amazon EKS nodes. Amazon EKS creates cross-account elastic network interfaces in these subnets to allow communication between your nodes and the Kubernetes control plane.
        :param bool endpoint_private_access: Set this value to true to enable private access for your cluster's Kubernetes API server endpoint. If you enable private access, Kubernetes API requests from within your cluster's VPC use the private VPC endpoint. The default value for this parameter is false, which disables private access for your Kubernetes API server. If you disable private access and you have nodes or AWS Fargate pods in the cluster, then ensure that publicAccessCidrs includes the necessary CIDR blocks for communication with the nodes or Fargate pods.
        :param bool endpoint_public_access: Set this value to false to disable public access to your cluster's Kubernetes API server endpoint. If you disable public access, your cluster's Kubernetes API server can only receive requests from within the cluster VPC. The default value for this parameter is true, which enables public access for your Kubernetes API server.
        :param Sequence[str] public_access_cidrs: The CIDR blocks that are allowed access to your cluster's public Kubernetes API server endpoint. Communication to the endpoint from addresses outside of the CIDR blocks that you specify is denied. The default value is 0.0.0.0/0. If you've disabled private endpoint access and you have nodes or AWS Fargate pods in the cluster, then ensure that you specify the necessary CIDR blocks.
        :param Sequence[str] security_group_ids: Specify one or more security groups for the cross-account elastic network interfaces that Amazon EKS creates to use to allow communication between your worker nodes and the Kubernetes control plane. If you don't specify a security group, the default security group for your VPC is used.
        """
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        if endpoint_private_access is not None:
            pulumi.set(__self__, "endpoint_private_access", endpoint_private_access)
        if endpoint_public_access is not None:
            pulumi.set(__self__, "endpoint_public_access", endpoint_public_access)
        if public_access_cidrs is not None:
            pulumi.set(__self__, "public_access_cidrs", public_access_cidrs)
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Sequence[str]:
        """
        Specify subnets for your Amazon EKS nodes. Amazon EKS creates cross-account elastic network interfaces in these subnets to allow communication between your nodes and the Kubernetes control plane.
        """
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter(name="endpointPrivateAccess")
    def endpoint_private_access(self) -> Optional[bool]:
        """
        Set this value to true to enable private access for your cluster's Kubernetes API server endpoint. If you enable private access, Kubernetes API requests from within your cluster's VPC use the private VPC endpoint. The default value for this parameter is false, which disables private access for your Kubernetes API server. If you disable private access and you have nodes or AWS Fargate pods in the cluster, then ensure that publicAccessCidrs includes the necessary CIDR blocks for communication with the nodes or Fargate pods.
        """
        return pulumi.get(self, "endpoint_private_access")

    @property
    @pulumi.getter(name="endpointPublicAccess")
    def endpoint_public_access(self) -> Optional[bool]:
        """
        Set this value to false to disable public access to your cluster's Kubernetes API server endpoint. If you disable public access, your cluster's Kubernetes API server can only receive requests from within the cluster VPC. The default value for this parameter is true, which enables public access for your Kubernetes API server.
        """
        return pulumi.get(self, "endpoint_public_access")

    @property
    @pulumi.getter(name="publicAccessCidrs")
    def public_access_cidrs(self) -> Optional[Sequence[str]]:
        """
        The CIDR blocks that are allowed access to your cluster's public Kubernetes API server endpoint. Communication to the endpoint from addresses outside of the CIDR blocks that you specify is denied. The default value is 0.0.0.0/0. If you've disabled private endpoint access and you have nodes or AWS Fargate pods in the cluster, then ensure that you specify the necessary CIDR blocks.
        """
        return pulumi.get(self, "public_access_cidrs")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[Sequence[str]]:
        """
        Specify one or more security groups for the cross-account elastic network interfaces that Amazon EKS creates to use to allow communication between your worker nodes and the Kubernetes control plane. If you don't specify a security group, the default security group for your VPC is used.
        """
        return pulumi.get(self, "security_group_ids")


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
class FargateProfileLabel(dict):
    """
    A key-value pair to associate with a pod.
    """
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        A key-value pair to associate with a pod.
        :param str key: The key name of the label.
        :param str value: The value for the label. 
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The key name of the label.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value for the label. 
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class FargateProfileSelector(dict):
    def __init__(__self__, *,
                 namespace: str,
                 labels: Optional[Sequence['outputs.FargateProfileLabel']] = None):
        pulumi.set(__self__, "namespace", namespace)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)

    @property
    @pulumi.getter
    def namespace(self) -> str:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def labels(self) -> Optional[Sequence['outputs.FargateProfileLabel']]:
        return pulumi.get(self, "labels")


@pulumi.output_type
class FargateProfileTag(dict):
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
class IdentityProviderConfigOidcIdentityProviderConfig(dict):
    """
    An object representing an OpenID Connect (OIDC) configuration.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "clientId":
            suggest = "client_id"
        elif key == "issuerUrl":
            suggest = "issuer_url"
        elif key == "groupsClaim":
            suggest = "groups_claim"
        elif key == "groupsPrefix":
            suggest = "groups_prefix"
        elif key == "requiredClaims":
            suggest = "required_claims"
        elif key == "usernameClaim":
            suggest = "username_claim"
        elif key == "usernamePrefix":
            suggest = "username_prefix"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in IdentityProviderConfigOidcIdentityProviderConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        IdentityProviderConfigOidcIdentityProviderConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        IdentityProviderConfigOidcIdentityProviderConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 client_id: str,
                 issuer_url: str,
                 groups_claim: Optional[str] = None,
                 groups_prefix: Optional[str] = None,
                 required_claims: Optional[Sequence['outputs.IdentityProviderConfigRequiredClaim']] = None,
                 username_claim: Optional[str] = None,
                 username_prefix: Optional[str] = None):
        """
        An object representing an OpenID Connect (OIDC) configuration.
        :param str client_id: This is also known as audience. The ID for the client application that makes authentication requests to the OpenID identity provider.
        :param str issuer_url: The URL of the OpenID identity provider that allows the API server to discover public signing keys for verifying tokens.
        :param str groups_claim: The JWT claim that the provider uses to return your groups.
        :param str groups_prefix: The prefix that is prepended to group claims to prevent clashes with existing names (such as system: groups).
        :param str username_claim: The JSON Web Token (JWT) claim to use as the username. The default is sub, which is expected to be a unique identifier of the end user. You can choose other claims, such as email or name, depending on the OpenID identity provider. Claims other than email are prefixed with the issuer URL to prevent naming clashes with other plug-ins.
        :param str username_prefix: The prefix that is prepended to username claims to prevent clashes with existing names. If you do not provide this field, and username is a value other than email, the prefix defaults to issuerurl#. You can use the value - to disable all prefixing.
        """
        pulumi.set(__self__, "client_id", client_id)
        pulumi.set(__self__, "issuer_url", issuer_url)
        if groups_claim is not None:
            pulumi.set(__self__, "groups_claim", groups_claim)
        if groups_prefix is not None:
            pulumi.set(__self__, "groups_prefix", groups_prefix)
        if required_claims is not None:
            pulumi.set(__self__, "required_claims", required_claims)
        if username_claim is not None:
            pulumi.set(__self__, "username_claim", username_claim)
        if username_prefix is not None:
            pulumi.set(__self__, "username_prefix", username_prefix)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        """
        This is also known as audience. The ID for the client application that makes authentication requests to the OpenID identity provider.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="issuerUrl")
    def issuer_url(self) -> str:
        """
        The URL of the OpenID identity provider that allows the API server to discover public signing keys for verifying tokens.
        """
        return pulumi.get(self, "issuer_url")

    @property
    @pulumi.getter(name="groupsClaim")
    def groups_claim(self) -> Optional[str]:
        """
        The JWT claim that the provider uses to return your groups.
        """
        return pulumi.get(self, "groups_claim")

    @property
    @pulumi.getter(name="groupsPrefix")
    def groups_prefix(self) -> Optional[str]:
        """
        The prefix that is prepended to group claims to prevent clashes with existing names (such as system: groups).
        """
        return pulumi.get(self, "groups_prefix")

    @property
    @pulumi.getter(name="requiredClaims")
    def required_claims(self) -> Optional[Sequence['outputs.IdentityProviderConfigRequiredClaim']]:
        return pulumi.get(self, "required_claims")

    @property
    @pulumi.getter(name="usernameClaim")
    def username_claim(self) -> Optional[str]:
        """
        The JSON Web Token (JWT) claim to use as the username. The default is sub, which is expected to be a unique identifier of the end user. You can choose other claims, such as email or name, depending on the OpenID identity provider. Claims other than email are prefixed with the issuer URL to prevent naming clashes with other plug-ins.
        """
        return pulumi.get(self, "username_claim")

    @property
    @pulumi.getter(name="usernamePrefix")
    def username_prefix(self) -> Optional[str]:
        """
        The prefix that is prepended to username claims to prevent clashes with existing names. If you do not provide this field, and username is a value other than email, the prefix defaults to issuerurl#. You can use the value - to disable all prefixing.
        """
        return pulumi.get(self, "username_prefix")


@pulumi.output_type
class IdentityProviderConfigRequiredClaim(dict):
    """
    The key value pairs that describe required claims in the identity token. If set, each claim is verified to be present in the token with a matching value.
    """
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        The key value pairs that describe required claims in the identity token. If set, each claim is verified to be present in the token with a matching value.
        :param str key: The key of the requiredClaims.
        :param str value: The value for the requiredClaims.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The key of the requiredClaims.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value for the requiredClaims.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class IdentityProviderConfigTag(dict):
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
class NodegroupLaunchTemplateSpecification(dict):
    """
    An object representing a launch template specification for AWS EKS Nodegroup.
    """
    def __init__(__self__, *,
                 id: Optional[str] = None,
                 name: Optional[str] = None,
                 version: Optional[str] = None):
        """
        An object representing a launch template specification for AWS EKS Nodegroup.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        return pulumi.get(self, "version")


@pulumi.output_type
class NodegroupRemoteAccess(dict):
    """
    An object representing a remote access configuration specification for AWS EKS Nodegroup.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "ec2SshKey":
            suggest = "ec2_ssh_key"
        elif key == "sourceSecurityGroups":
            suggest = "source_security_groups"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NodegroupRemoteAccess. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NodegroupRemoteAccess.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NodegroupRemoteAccess.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 ec2_ssh_key: str,
                 source_security_groups: Optional[Sequence[str]] = None):
        """
        An object representing a remote access configuration specification for AWS EKS Nodegroup.
        """
        pulumi.set(__self__, "ec2_ssh_key", ec2_ssh_key)
        if source_security_groups is not None:
            pulumi.set(__self__, "source_security_groups", source_security_groups)

    @property
    @pulumi.getter(name="ec2SshKey")
    def ec2_ssh_key(self) -> str:
        return pulumi.get(self, "ec2_ssh_key")

    @property
    @pulumi.getter(name="sourceSecurityGroups")
    def source_security_groups(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "source_security_groups")


@pulumi.output_type
class NodegroupScalingConfig(dict):
    """
    An object representing a auto scaling group specification for AWS EKS Nodegroup.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "desiredSize":
            suggest = "desired_size"
        elif key == "maxSize":
            suggest = "max_size"
        elif key == "minSize":
            suggest = "min_size"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NodegroupScalingConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NodegroupScalingConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NodegroupScalingConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 desired_size: Optional[int] = None,
                 max_size: Optional[int] = None,
                 min_size: Optional[int] = None):
        """
        An object representing a auto scaling group specification for AWS EKS Nodegroup.
        """
        if desired_size is not None:
            pulumi.set(__self__, "desired_size", desired_size)
        if max_size is not None:
            pulumi.set(__self__, "max_size", max_size)
        if min_size is not None:
            pulumi.set(__self__, "min_size", min_size)

    @property
    @pulumi.getter(name="desiredSize")
    def desired_size(self) -> Optional[int]:
        return pulumi.get(self, "desired_size")

    @property
    @pulumi.getter(name="maxSize")
    def max_size(self) -> Optional[int]:
        return pulumi.get(self, "max_size")

    @property
    @pulumi.getter(name="minSize")
    def min_size(self) -> Optional[int]:
        return pulumi.get(self, "min_size")


@pulumi.output_type
class NodegroupTaint(dict):
    """
    An object representing a Taint specification for AWS EKS Nodegroup.
    """
    def __init__(__self__, *,
                 effect: Optional[str] = None,
                 key: Optional[str] = None,
                 value: Optional[str] = None):
        """
        An object representing a Taint specification for AWS EKS Nodegroup.
        """
        if effect is not None:
            pulumi.set(__self__, "effect", effect)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def effect(self) -> Optional[str]:
        return pulumi.get(self, "effect")

    @property
    @pulumi.getter
    def key(self) -> Optional[str]:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        return pulumi.get(self, "value")


@pulumi.output_type
class NodegroupUpdateConfig(dict):
    """
    The node group update configuration.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "maxUnavailable":
            suggest = "max_unavailable"
        elif key == "maxUnavailablePercentage":
            suggest = "max_unavailable_percentage"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NodegroupUpdateConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NodegroupUpdateConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NodegroupUpdateConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 max_unavailable: Optional[float] = None,
                 max_unavailable_percentage: Optional[float] = None):
        """
        The node group update configuration.
        :param float max_unavailable: The maximum number of nodes unavailable at once during a version update. Nodes will be updated in parallel. This value or maxUnavailablePercentage is required to have a value.The maximum number is 100. 
        :param float max_unavailable_percentage: The maximum percentage of nodes unavailable during a version update. This percentage of nodes will be updated in parallel, up to 100 nodes at once. This value or maxUnavailable is required to have a value.
        """
        if max_unavailable is not None:
            pulumi.set(__self__, "max_unavailable", max_unavailable)
        if max_unavailable_percentage is not None:
            pulumi.set(__self__, "max_unavailable_percentage", max_unavailable_percentage)

    @property
    @pulumi.getter(name="maxUnavailable")
    def max_unavailable(self) -> Optional[float]:
        """
        The maximum number of nodes unavailable at once during a version update. Nodes will be updated in parallel. This value or maxUnavailablePercentage is required to have a value.The maximum number is 100. 
        """
        return pulumi.get(self, "max_unavailable")

    @property
    @pulumi.getter(name="maxUnavailablePercentage")
    def max_unavailable_percentage(self) -> Optional[float]:
        """
        The maximum percentage of nodes unavailable during a version update. This percentage of nodes will be updated in parallel, up to 100 nodes at once. This value or maxUnavailable is required to have a value.
        """
        return pulumi.get(self, "max_unavailable_percentage")


