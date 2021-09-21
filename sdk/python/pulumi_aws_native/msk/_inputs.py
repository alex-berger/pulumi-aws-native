# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ClusterBrokerLogsArgs',
    'ClusterBrokerNodeGroupInfoArgs',
    'ClusterClientAuthenticationArgs',
    'ClusterCloudWatchLogsArgs',
    'ClusterConfigurationInfoArgs',
    'ClusterEBSStorageInfoArgs',
    'ClusterEncryptionAtRestArgs',
    'ClusterEncryptionInTransitArgs',
    'ClusterEncryptionInfoArgs',
    'ClusterFirehoseArgs',
    'ClusterIamArgs',
    'ClusterJmxExporterArgs',
    'ClusterLoggingInfoArgs',
    'ClusterNodeExporterArgs',
    'ClusterOpenMonitoringArgs',
    'ClusterPrometheusArgs',
    'ClusterS3Args',
    'ClusterSaslArgs',
    'ClusterScramArgs',
    'ClusterStorageInfoArgs',
    'ClusterTlsArgs',
]

@pulumi.input_type
class ClusterBrokerLogsArgs:
    def __init__(__self__, *,
                 cloud_watch_logs: Optional[pulumi.Input['ClusterCloudWatchLogsArgs']] = None,
                 firehose: Optional[pulumi.Input['ClusterFirehoseArgs']] = None,
                 s3: Optional[pulumi.Input['ClusterS3Args']] = None):
        if cloud_watch_logs is not None:
            pulumi.set(__self__, "cloud_watch_logs", cloud_watch_logs)
        if firehose is not None:
            pulumi.set(__self__, "firehose", firehose)
        if s3 is not None:
            pulumi.set(__self__, "s3", s3)

    @property
    @pulumi.getter(name="cloudWatchLogs")
    def cloud_watch_logs(self) -> Optional[pulumi.Input['ClusterCloudWatchLogsArgs']]:
        return pulumi.get(self, "cloud_watch_logs")

    @cloud_watch_logs.setter
    def cloud_watch_logs(self, value: Optional[pulumi.Input['ClusterCloudWatchLogsArgs']]):
        pulumi.set(self, "cloud_watch_logs", value)

    @property
    @pulumi.getter
    def firehose(self) -> Optional[pulumi.Input['ClusterFirehoseArgs']]:
        return pulumi.get(self, "firehose")

    @firehose.setter
    def firehose(self, value: Optional[pulumi.Input['ClusterFirehoseArgs']]):
        pulumi.set(self, "firehose", value)

    @property
    @pulumi.getter
    def s3(self) -> Optional[pulumi.Input['ClusterS3Args']]:
        return pulumi.get(self, "s3")

    @s3.setter
    def s3(self, value: Optional[pulumi.Input['ClusterS3Args']]):
        pulumi.set(self, "s3", value)


@pulumi.input_type
class ClusterBrokerNodeGroupInfoArgs:
    def __init__(__self__, *,
                 client_subnets: pulumi.Input[Sequence[pulumi.Input[str]]],
                 instance_type: pulumi.Input[str],
                 broker_az_distribution: Optional[pulumi.Input[str]] = None,
                 security_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 storage_info: Optional[pulumi.Input['ClusterStorageInfoArgs']] = None):
        pulumi.set(__self__, "client_subnets", client_subnets)
        pulumi.set(__self__, "instance_type", instance_type)
        if broker_az_distribution is not None:
            pulumi.set(__self__, "broker_az_distribution", broker_az_distribution)
        if security_groups is not None:
            pulumi.set(__self__, "security_groups", security_groups)
        if storage_info is not None:
            pulumi.set(__self__, "storage_info", storage_info)

    @property
    @pulumi.getter(name="clientSubnets")
    def client_subnets(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "client_subnets")

    @client_subnets.setter
    def client_subnets(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "client_subnets", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter(name="brokerAZDistribution")
    def broker_az_distribution(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "broker_az_distribution")

    @broker_az_distribution.setter
    def broker_az_distribution(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "broker_az_distribution", value)

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "security_groups")

    @security_groups.setter
    def security_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_groups", value)

    @property
    @pulumi.getter(name="storageInfo")
    def storage_info(self) -> Optional[pulumi.Input['ClusterStorageInfoArgs']]:
        return pulumi.get(self, "storage_info")

    @storage_info.setter
    def storage_info(self, value: Optional[pulumi.Input['ClusterStorageInfoArgs']]):
        pulumi.set(self, "storage_info", value)


@pulumi.input_type
class ClusterClientAuthenticationArgs:
    def __init__(__self__, *,
                 sasl: Optional[pulumi.Input['ClusterSaslArgs']] = None,
                 tls: Optional[pulumi.Input['ClusterTlsArgs']] = None):
        if sasl is not None:
            pulumi.set(__self__, "sasl", sasl)
        if tls is not None:
            pulumi.set(__self__, "tls", tls)

    @property
    @pulumi.getter
    def sasl(self) -> Optional[pulumi.Input['ClusterSaslArgs']]:
        return pulumi.get(self, "sasl")

    @sasl.setter
    def sasl(self, value: Optional[pulumi.Input['ClusterSaslArgs']]):
        pulumi.set(self, "sasl", value)

    @property
    @pulumi.getter
    def tls(self) -> Optional[pulumi.Input['ClusterTlsArgs']]:
        return pulumi.get(self, "tls")

    @tls.setter
    def tls(self, value: Optional[pulumi.Input['ClusterTlsArgs']]):
        pulumi.set(self, "tls", value)


@pulumi.input_type
class ClusterCloudWatchLogsArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 log_group: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "enabled", enabled)
        if log_group is not None:
            pulumi.set(__self__, "log_group", log_group)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="logGroup")
    def log_group(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "log_group")

    @log_group.setter
    def log_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_group", value)


@pulumi.input_type
class ClusterConfigurationInfoArgs:
    def __init__(__self__, *,
                 arn: pulumi.Input[str],
                 revision: pulumi.Input[int]):
        pulumi.set(__self__, "arn", arn)
        pulumi.set(__self__, "revision", revision)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Input[str]:
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def revision(self) -> pulumi.Input[int]:
        return pulumi.get(self, "revision")

    @revision.setter
    def revision(self, value: pulumi.Input[int]):
        pulumi.set(self, "revision", value)


@pulumi.input_type
class ClusterEBSStorageInfoArgs:
    def __init__(__self__, *,
                 volume_size: Optional[pulumi.Input[int]] = None):
        if volume_size is not None:
            pulumi.set(__self__, "volume_size", volume_size)

    @property
    @pulumi.getter(name="volumeSize")
    def volume_size(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "volume_size")

    @volume_size.setter
    def volume_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "volume_size", value)


@pulumi.input_type
class ClusterEncryptionAtRestArgs:
    def __init__(__self__, *,
                 data_volume_kms_key_id: pulumi.Input[str]):
        pulumi.set(__self__, "data_volume_kms_key_id", data_volume_kms_key_id)

    @property
    @pulumi.getter(name="dataVolumeKMSKeyId")
    def data_volume_kms_key_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "data_volume_kms_key_id")

    @data_volume_kms_key_id.setter
    def data_volume_kms_key_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "data_volume_kms_key_id", value)


@pulumi.input_type
class ClusterEncryptionInTransitArgs:
    def __init__(__self__, *,
                 client_broker: Optional[pulumi.Input[str]] = None,
                 in_cluster: Optional[pulumi.Input[bool]] = None):
        if client_broker is not None:
            pulumi.set(__self__, "client_broker", client_broker)
        if in_cluster is not None:
            pulumi.set(__self__, "in_cluster", in_cluster)

    @property
    @pulumi.getter(name="clientBroker")
    def client_broker(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "client_broker")

    @client_broker.setter
    def client_broker(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_broker", value)

    @property
    @pulumi.getter(name="inCluster")
    def in_cluster(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "in_cluster")

    @in_cluster.setter
    def in_cluster(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "in_cluster", value)


@pulumi.input_type
class ClusterEncryptionInfoArgs:
    def __init__(__self__, *,
                 encryption_at_rest: Optional[pulumi.Input['ClusterEncryptionAtRestArgs']] = None,
                 encryption_in_transit: Optional[pulumi.Input['ClusterEncryptionInTransitArgs']] = None):
        if encryption_at_rest is not None:
            pulumi.set(__self__, "encryption_at_rest", encryption_at_rest)
        if encryption_in_transit is not None:
            pulumi.set(__self__, "encryption_in_transit", encryption_in_transit)

    @property
    @pulumi.getter(name="encryptionAtRest")
    def encryption_at_rest(self) -> Optional[pulumi.Input['ClusterEncryptionAtRestArgs']]:
        return pulumi.get(self, "encryption_at_rest")

    @encryption_at_rest.setter
    def encryption_at_rest(self, value: Optional[pulumi.Input['ClusterEncryptionAtRestArgs']]):
        pulumi.set(self, "encryption_at_rest", value)

    @property
    @pulumi.getter(name="encryptionInTransit")
    def encryption_in_transit(self) -> Optional[pulumi.Input['ClusterEncryptionInTransitArgs']]:
        return pulumi.get(self, "encryption_in_transit")

    @encryption_in_transit.setter
    def encryption_in_transit(self, value: Optional[pulumi.Input['ClusterEncryptionInTransitArgs']]):
        pulumi.set(self, "encryption_in_transit", value)


@pulumi.input_type
class ClusterFirehoseArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 delivery_stream: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "enabled", enabled)
        if delivery_stream is not None:
            pulumi.set(__self__, "delivery_stream", delivery_stream)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="deliveryStream")
    def delivery_stream(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "delivery_stream")

    @delivery_stream.setter
    def delivery_stream(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delivery_stream", value)


@pulumi.input_type
class ClusterIamArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool]):
        pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class ClusterJmxExporterArgs:
    def __init__(__self__, *,
                 enabled_in_broker: pulumi.Input[bool]):
        pulumi.set(__self__, "enabled_in_broker", enabled_in_broker)

    @property
    @pulumi.getter(name="enabledInBroker")
    def enabled_in_broker(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "enabled_in_broker")

    @enabled_in_broker.setter
    def enabled_in_broker(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled_in_broker", value)


@pulumi.input_type
class ClusterLoggingInfoArgs:
    def __init__(__self__, *,
                 broker_logs: pulumi.Input['ClusterBrokerLogsArgs']):
        pulumi.set(__self__, "broker_logs", broker_logs)

    @property
    @pulumi.getter(name="brokerLogs")
    def broker_logs(self) -> pulumi.Input['ClusterBrokerLogsArgs']:
        return pulumi.get(self, "broker_logs")

    @broker_logs.setter
    def broker_logs(self, value: pulumi.Input['ClusterBrokerLogsArgs']):
        pulumi.set(self, "broker_logs", value)


@pulumi.input_type
class ClusterNodeExporterArgs:
    def __init__(__self__, *,
                 enabled_in_broker: pulumi.Input[bool]):
        pulumi.set(__self__, "enabled_in_broker", enabled_in_broker)

    @property
    @pulumi.getter(name="enabledInBroker")
    def enabled_in_broker(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "enabled_in_broker")

    @enabled_in_broker.setter
    def enabled_in_broker(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled_in_broker", value)


@pulumi.input_type
class ClusterOpenMonitoringArgs:
    def __init__(__self__, *,
                 prometheus: pulumi.Input['ClusterPrometheusArgs']):
        pulumi.set(__self__, "prometheus", prometheus)

    @property
    @pulumi.getter
    def prometheus(self) -> pulumi.Input['ClusterPrometheusArgs']:
        return pulumi.get(self, "prometheus")

    @prometheus.setter
    def prometheus(self, value: pulumi.Input['ClusterPrometheusArgs']):
        pulumi.set(self, "prometheus", value)


@pulumi.input_type
class ClusterPrometheusArgs:
    def __init__(__self__, *,
                 jmx_exporter: Optional[pulumi.Input['ClusterJmxExporterArgs']] = None,
                 node_exporter: Optional[pulumi.Input['ClusterNodeExporterArgs']] = None):
        if jmx_exporter is not None:
            pulumi.set(__self__, "jmx_exporter", jmx_exporter)
        if node_exporter is not None:
            pulumi.set(__self__, "node_exporter", node_exporter)

    @property
    @pulumi.getter(name="jmxExporter")
    def jmx_exporter(self) -> Optional[pulumi.Input['ClusterJmxExporterArgs']]:
        return pulumi.get(self, "jmx_exporter")

    @jmx_exporter.setter
    def jmx_exporter(self, value: Optional[pulumi.Input['ClusterJmxExporterArgs']]):
        pulumi.set(self, "jmx_exporter", value)

    @property
    @pulumi.getter(name="nodeExporter")
    def node_exporter(self) -> Optional[pulumi.Input['ClusterNodeExporterArgs']]:
        return pulumi.get(self, "node_exporter")

    @node_exporter.setter
    def node_exporter(self, value: Optional[pulumi.Input['ClusterNodeExporterArgs']]):
        pulumi.set(self, "node_exporter", value)


@pulumi.input_type
class ClusterS3Args:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 bucket: Optional[pulumi.Input[str]] = None,
                 prefix: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "enabled", enabled)
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)


@pulumi.input_type
class ClusterSaslArgs:
    def __init__(__self__, *,
                 iam: Optional[pulumi.Input['ClusterIamArgs']] = None,
                 scram: Optional[pulumi.Input['ClusterScramArgs']] = None):
        if iam is not None:
            pulumi.set(__self__, "iam", iam)
        if scram is not None:
            pulumi.set(__self__, "scram", scram)

    @property
    @pulumi.getter
    def iam(self) -> Optional[pulumi.Input['ClusterIamArgs']]:
        return pulumi.get(self, "iam")

    @iam.setter
    def iam(self, value: Optional[pulumi.Input['ClusterIamArgs']]):
        pulumi.set(self, "iam", value)

    @property
    @pulumi.getter
    def scram(self) -> Optional[pulumi.Input['ClusterScramArgs']]:
        return pulumi.get(self, "scram")

    @scram.setter
    def scram(self, value: Optional[pulumi.Input['ClusterScramArgs']]):
        pulumi.set(self, "scram", value)


@pulumi.input_type
class ClusterScramArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool]):
        pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class ClusterStorageInfoArgs:
    def __init__(__self__, *,
                 e_bs_storage_info: Optional[pulumi.Input['ClusterEBSStorageInfoArgs']] = None):
        if e_bs_storage_info is not None:
            pulumi.set(__self__, "e_bs_storage_info", e_bs_storage_info)

    @property
    @pulumi.getter(name="eBSStorageInfo")
    def e_bs_storage_info(self) -> Optional[pulumi.Input['ClusterEBSStorageInfoArgs']]:
        return pulumi.get(self, "e_bs_storage_info")

    @e_bs_storage_info.setter
    def e_bs_storage_info(self, value: Optional[pulumi.Input['ClusterEBSStorageInfoArgs']]):
        pulumi.set(self, "e_bs_storage_info", value)


@pulumi.input_type
class ClusterTlsArgs:
    def __init__(__self__, *,
                 certificate_authority_arn_list: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        if certificate_authority_arn_list is not None:
            pulumi.set(__self__, "certificate_authority_arn_list", certificate_authority_arn_list)

    @property
    @pulumi.getter(name="certificateAuthorityArnList")
    def certificate_authority_arn_list(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "certificate_authority_arn_list")

    @certificate_authority_arn_list.setter
    def certificate_authority_arn_list(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "certificate_authority_arn_list", value)

