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
    'ClusterBrokerLogs',
    'ClusterBrokerNodeGroupInfo',
    'ClusterClientAuthentication',
    'ClusterCloudWatchLogs',
    'ClusterConfigurationInfo',
    'ClusterConnectivityInfo',
    'ClusterEBSStorageInfo',
    'ClusterEncryptionAtRest',
    'ClusterEncryptionInTransit',
    'ClusterEncryptionInfo',
    'ClusterFirehose',
    'ClusterIam',
    'ClusterJmxExporter',
    'ClusterLoggingInfo',
    'ClusterNodeExporter',
    'ClusterOpenMonitoring',
    'ClusterPrometheus',
    'ClusterProvisionedThroughput',
    'ClusterPublicAccess',
    'ClusterS3',
    'ClusterSasl',
    'ClusterScram',
    'ClusterStorageInfo',
    'ClusterTls',
    'ClusterUnauthenticated',
]

@pulumi.output_type
class ClusterBrokerLogs(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "cloudWatchLogs":
            suggest = "cloud_watch_logs"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterBrokerLogs. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterBrokerLogs.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterBrokerLogs.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cloud_watch_logs: Optional['outputs.ClusterCloudWatchLogs'] = None,
                 firehose: Optional['outputs.ClusterFirehose'] = None,
                 s3: Optional['outputs.ClusterS3'] = None):
        if cloud_watch_logs is not None:
            pulumi.set(__self__, "cloud_watch_logs", cloud_watch_logs)
        if firehose is not None:
            pulumi.set(__self__, "firehose", firehose)
        if s3 is not None:
            pulumi.set(__self__, "s3", s3)

    @property
    @pulumi.getter(name="cloudWatchLogs")
    def cloud_watch_logs(self) -> Optional['outputs.ClusterCloudWatchLogs']:
        return pulumi.get(self, "cloud_watch_logs")

    @property
    @pulumi.getter
    def firehose(self) -> Optional['outputs.ClusterFirehose']:
        return pulumi.get(self, "firehose")

    @property
    @pulumi.getter
    def s3(self) -> Optional['outputs.ClusterS3']:
        return pulumi.get(self, "s3")


@pulumi.output_type
class ClusterBrokerNodeGroupInfo(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "clientSubnets":
            suggest = "client_subnets"
        elif key == "instanceType":
            suggest = "instance_type"
        elif key == "brokerAZDistribution":
            suggest = "broker_az_distribution"
        elif key == "connectivityInfo":
            suggest = "connectivity_info"
        elif key == "securityGroups":
            suggest = "security_groups"
        elif key == "storageInfo":
            suggest = "storage_info"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterBrokerNodeGroupInfo. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterBrokerNodeGroupInfo.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterBrokerNodeGroupInfo.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 client_subnets: Sequence[str],
                 instance_type: str,
                 broker_az_distribution: Optional[str] = None,
                 connectivity_info: Optional['outputs.ClusterConnectivityInfo'] = None,
                 security_groups: Optional[Sequence[str]] = None,
                 storage_info: Optional['outputs.ClusterStorageInfo'] = None):
        pulumi.set(__self__, "client_subnets", client_subnets)
        pulumi.set(__self__, "instance_type", instance_type)
        if broker_az_distribution is not None:
            pulumi.set(__self__, "broker_az_distribution", broker_az_distribution)
        if connectivity_info is not None:
            pulumi.set(__self__, "connectivity_info", connectivity_info)
        if security_groups is not None:
            pulumi.set(__self__, "security_groups", security_groups)
        if storage_info is not None:
            pulumi.set(__self__, "storage_info", storage_info)

    @property
    @pulumi.getter(name="clientSubnets")
    def client_subnets(self) -> Sequence[str]:
        return pulumi.get(self, "client_subnets")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> str:
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="brokerAZDistribution")
    def broker_az_distribution(self) -> Optional[str]:
        return pulumi.get(self, "broker_az_distribution")

    @property
    @pulumi.getter(name="connectivityInfo")
    def connectivity_info(self) -> Optional['outputs.ClusterConnectivityInfo']:
        return pulumi.get(self, "connectivity_info")

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "security_groups")

    @property
    @pulumi.getter(name="storageInfo")
    def storage_info(self) -> Optional['outputs.ClusterStorageInfo']:
        return pulumi.get(self, "storage_info")


@pulumi.output_type
class ClusterClientAuthentication(dict):
    def __init__(__self__, *,
                 sasl: Optional['outputs.ClusterSasl'] = None,
                 tls: Optional['outputs.ClusterTls'] = None,
                 unauthenticated: Optional['outputs.ClusterUnauthenticated'] = None):
        if sasl is not None:
            pulumi.set(__self__, "sasl", sasl)
        if tls is not None:
            pulumi.set(__self__, "tls", tls)
        if unauthenticated is not None:
            pulumi.set(__self__, "unauthenticated", unauthenticated)

    @property
    @pulumi.getter
    def sasl(self) -> Optional['outputs.ClusterSasl']:
        return pulumi.get(self, "sasl")

    @property
    @pulumi.getter
    def tls(self) -> Optional['outputs.ClusterTls']:
        return pulumi.get(self, "tls")

    @property
    @pulumi.getter
    def unauthenticated(self) -> Optional['outputs.ClusterUnauthenticated']:
        return pulumi.get(self, "unauthenticated")


@pulumi.output_type
class ClusterCloudWatchLogs(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "logGroup":
            suggest = "log_group"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterCloudWatchLogs. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterCloudWatchLogs.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterCloudWatchLogs.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 enabled: bool,
                 log_group: Optional[str] = None):
        pulumi.set(__self__, "enabled", enabled)
        if log_group is not None:
            pulumi.set(__self__, "log_group", log_group)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="logGroup")
    def log_group(self) -> Optional[str]:
        return pulumi.get(self, "log_group")


@pulumi.output_type
class ClusterConfigurationInfo(dict):
    def __init__(__self__, *,
                 arn: str,
                 revision: int):
        pulumi.set(__self__, "arn", arn)
        pulumi.set(__self__, "revision", revision)

    @property
    @pulumi.getter
    def arn(self) -> str:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def revision(self) -> int:
        return pulumi.get(self, "revision")


@pulumi.output_type
class ClusterConnectivityInfo(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "publicAccess":
            suggest = "public_access"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterConnectivityInfo. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterConnectivityInfo.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterConnectivityInfo.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 public_access: Optional['outputs.ClusterPublicAccess'] = None):
        if public_access is not None:
            pulumi.set(__self__, "public_access", public_access)

    @property
    @pulumi.getter(name="publicAccess")
    def public_access(self) -> Optional['outputs.ClusterPublicAccess']:
        return pulumi.get(self, "public_access")


@pulumi.output_type
class ClusterEBSStorageInfo(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "provisionedThroughput":
            suggest = "provisioned_throughput"
        elif key == "volumeSize":
            suggest = "volume_size"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterEBSStorageInfo. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterEBSStorageInfo.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterEBSStorageInfo.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 provisioned_throughput: Optional['outputs.ClusterProvisionedThroughput'] = None,
                 volume_size: Optional[int] = None):
        if provisioned_throughput is not None:
            pulumi.set(__self__, "provisioned_throughput", provisioned_throughput)
        if volume_size is not None:
            pulumi.set(__self__, "volume_size", volume_size)

    @property
    @pulumi.getter(name="provisionedThroughput")
    def provisioned_throughput(self) -> Optional['outputs.ClusterProvisionedThroughput']:
        return pulumi.get(self, "provisioned_throughput")

    @property
    @pulumi.getter(name="volumeSize")
    def volume_size(self) -> Optional[int]:
        return pulumi.get(self, "volume_size")


@pulumi.output_type
class ClusterEncryptionAtRest(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "dataVolumeKMSKeyId":
            suggest = "data_volume_kms_key_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterEncryptionAtRest. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterEncryptionAtRest.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterEncryptionAtRest.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 data_volume_kms_key_id: str):
        pulumi.set(__self__, "data_volume_kms_key_id", data_volume_kms_key_id)

    @property
    @pulumi.getter(name="dataVolumeKMSKeyId")
    def data_volume_kms_key_id(self) -> str:
        return pulumi.get(self, "data_volume_kms_key_id")


@pulumi.output_type
class ClusterEncryptionInTransit(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "clientBroker":
            suggest = "client_broker"
        elif key == "inCluster":
            suggest = "in_cluster"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterEncryptionInTransit. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterEncryptionInTransit.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterEncryptionInTransit.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 client_broker: Optional['ClusterEncryptionInTransitClientBroker'] = None,
                 in_cluster: Optional[bool] = None):
        if client_broker is not None:
            pulumi.set(__self__, "client_broker", client_broker)
        if in_cluster is not None:
            pulumi.set(__self__, "in_cluster", in_cluster)

    @property
    @pulumi.getter(name="clientBroker")
    def client_broker(self) -> Optional['ClusterEncryptionInTransitClientBroker']:
        return pulumi.get(self, "client_broker")

    @property
    @pulumi.getter(name="inCluster")
    def in_cluster(self) -> Optional[bool]:
        return pulumi.get(self, "in_cluster")


@pulumi.output_type
class ClusterEncryptionInfo(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "encryptionAtRest":
            suggest = "encryption_at_rest"
        elif key == "encryptionInTransit":
            suggest = "encryption_in_transit"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterEncryptionInfo. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterEncryptionInfo.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterEncryptionInfo.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 encryption_at_rest: Optional['outputs.ClusterEncryptionAtRest'] = None,
                 encryption_in_transit: Optional['outputs.ClusterEncryptionInTransit'] = None):
        if encryption_at_rest is not None:
            pulumi.set(__self__, "encryption_at_rest", encryption_at_rest)
        if encryption_in_transit is not None:
            pulumi.set(__self__, "encryption_in_transit", encryption_in_transit)

    @property
    @pulumi.getter(name="encryptionAtRest")
    def encryption_at_rest(self) -> Optional['outputs.ClusterEncryptionAtRest']:
        return pulumi.get(self, "encryption_at_rest")

    @property
    @pulumi.getter(name="encryptionInTransit")
    def encryption_in_transit(self) -> Optional['outputs.ClusterEncryptionInTransit']:
        return pulumi.get(self, "encryption_in_transit")


@pulumi.output_type
class ClusterFirehose(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "deliveryStream":
            suggest = "delivery_stream"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterFirehose. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterFirehose.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterFirehose.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 enabled: bool,
                 delivery_stream: Optional[str] = None):
        pulumi.set(__self__, "enabled", enabled)
        if delivery_stream is not None:
            pulumi.set(__self__, "delivery_stream", delivery_stream)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="deliveryStream")
    def delivery_stream(self) -> Optional[str]:
        return pulumi.get(self, "delivery_stream")


@pulumi.output_type
class ClusterIam(dict):
    def __init__(__self__, *,
                 enabled: bool):
        pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")


@pulumi.output_type
class ClusterJmxExporter(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "enabledInBroker":
            suggest = "enabled_in_broker"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterJmxExporter. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterJmxExporter.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterJmxExporter.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 enabled_in_broker: bool):
        pulumi.set(__self__, "enabled_in_broker", enabled_in_broker)

    @property
    @pulumi.getter(name="enabledInBroker")
    def enabled_in_broker(self) -> bool:
        return pulumi.get(self, "enabled_in_broker")


@pulumi.output_type
class ClusterLoggingInfo(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "brokerLogs":
            suggest = "broker_logs"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterLoggingInfo. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterLoggingInfo.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterLoggingInfo.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 broker_logs: 'outputs.ClusterBrokerLogs'):
        pulumi.set(__self__, "broker_logs", broker_logs)

    @property
    @pulumi.getter(name="brokerLogs")
    def broker_logs(self) -> 'outputs.ClusterBrokerLogs':
        return pulumi.get(self, "broker_logs")


@pulumi.output_type
class ClusterNodeExporter(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "enabledInBroker":
            suggest = "enabled_in_broker"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterNodeExporter. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterNodeExporter.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterNodeExporter.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 enabled_in_broker: bool):
        pulumi.set(__self__, "enabled_in_broker", enabled_in_broker)

    @property
    @pulumi.getter(name="enabledInBroker")
    def enabled_in_broker(self) -> bool:
        return pulumi.get(self, "enabled_in_broker")


@pulumi.output_type
class ClusterOpenMonitoring(dict):
    def __init__(__self__, *,
                 prometheus: 'outputs.ClusterPrometheus'):
        pulumi.set(__self__, "prometheus", prometheus)

    @property
    @pulumi.getter
    def prometheus(self) -> 'outputs.ClusterPrometheus':
        return pulumi.get(self, "prometheus")


@pulumi.output_type
class ClusterPrometheus(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "jmxExporter":
            suggest = "jmx_exporter"
        elif key == "nodeExporter":
            suggest = "node_exporter"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterPrometheus. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterPrometheus.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterPrometheus.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 jmx_exporter: Optional['outputs.ClusterJmxExporter'] = None,
                 node_exporter: Optional['outputs.ClusterNodeExporter'] = None):
        if jmx_exporter is not None:
            pulumi.set(__self__, "jmx_exporter", jmx_exporter)
        if node_exporter is not None:
            pulumi.set(__self__, "node_exporter", node_exporter)

    @property
    @pulumi.getter(name="jmxExporter")
    def jmx_exporter(self) -> Optional['outputs.ClusterJmxExporter']:
        return pulumi.get(self, "jmx_exporter")

    @property
    @pulumi.getter(name="nodeExporter")
    def node_exporter(self) -> Optional['outputs.ClusterNodeExporter']:
        return pulumi.get(self, "node_exporter")


@pulumi.output_type
class ClusterProvisionedThroughput(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "volumeThroughput":
            suggest = "volume_throughput"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterProvisionedThroughput. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterProvisionedThroughput.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterProvisionedThroughput.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 enabled: Optional[bool] = None,
                 volume_throughput: Optional[int] = None):
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if volume_throughput is not None:
            pulumi.set(__self__, "volume_throughput", volume_throughput)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="volumeThroughput")
    def volume_throughput(self) -> Optional[int]:
        return pulumi.get(self, "volume_throughput")


@pulumi.output_type
class ClusterPublicAccess(dict):
    def __init__(__self__, *,
                 type: Optional[str] = None):
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")


@pulumi.output_type
class ClusterS3(dict):
    def __init__(__self__, *,
                 enabled: bool,
                 bucket: Optional[str] = None,
                 prefix: Optional[str] = None):
        pulumi.set(__self__, "enabled", enabled)
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def bucket(self) -> Optional[str]:
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def prefix(self) -> Optional[str]:
        return pulumi.get(self, "prefix")


@pulumi.output_type
class ClusterSasl(dict):
    def __init__(__self__, *,
                 iam: Optional['outputs.ClusterIam'] = None,
                 scram: Optional['outputs.ClusterScram'] = None):
        if iam is not None:
            pulumi.set(__self__, "iam", iam)
        if scram is not None:
            pulumi.set(__self__, "scram", scram)

    @property
    @pulumi.getter
    def iam(self) -> Optional['outputs.ClusterIam']:
        return pulumi.get(self, "iam")

    @property
    @pulumi.getter
    def scram(self) -> Optional['outputs.ClusterScram']:
        return pulumi.get(self, "scram")


@pulumi.output_type
class ClusterScram(dict):
    def __init__(__self__, *,
                 enabled: bool):
        pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")


@pulumi.output_type
class ClusterStorageInfo(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "eBSStorageInfo":
            suggest = "e_bs_storage_info"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterStorageInfo. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterStorageInfo.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterStorageInfo.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 e_bs_storage_info: Optional['outputs.ClusterEBSStorageInfo'] = None):
        if e_bs_storage_info is not None:
            pulumi.set(__self__, "e_bs_storage_info", e_bs_storage_info)

    @property
    @pulumi.getter(name="eBSStorageInfo")
    def e_bs_storage_info(self) -> Optional['outputs.ClusterEBSStorageInfo']:
        return pulumi.get(self, "e_bs_storage_info")


@pulumi.output_type
class ClusterTls(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "certificateAuthorityArnList":
            suggest = "certificate_authority_arn_list"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterTls. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterTls.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterTls.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 certificate_authority_arn_list: Optional[Sequence[str]] = None,
                 enabled: Optional[bool] = None):
        if certificate_authority_arn_list is not None:
            pulumi.set(__self__, "certificate_authority_arn_list", certificate_authority_arn_list)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter(name="certificateAuthorityArnList")
    def certificate_authority_arn_list(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "certificate_authority_arn_list")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        return pulumi.get(self, "enabled")


@pulumi.output_type
class ClusterUnauthenticated(dict):
    def __init__(__self__, *,
                 enabled: bool):
        pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")


