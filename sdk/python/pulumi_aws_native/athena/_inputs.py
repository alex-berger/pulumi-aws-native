# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from .. import _inputs as _root_inputs

__all__ = [
    'DataCatalogTagsArgs',
    'WorkGroupEncryptionConfigurationArgs',
    'WorkGroupResultConfigurationArgs',
    'WorkGroupResultConfigurationUpdatesArgs',
    'WorkGroupTagsArgs',
    'WorkGroupWorkGroupConfigurationArgs',
    'WorkGroupWorkGroupConfigurationUpdatesArgs',
]

@pulumi.input_type
class DataCatalogTagsArgs:
    def __init__(__self__, *,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-datacatalog-tags.html
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-datacatalog-tags.html#cfn-athena-datacatalog-tags-tags
        """
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-datacatalog-tags.html#cfn-athena-datacatalog-tags-tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class WorkGroupEncryptionConfigurationArgs:
    def __init__(__self__, *,
                 encryption_option: pulumi.Input[str],
                 kms_key: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-encryptionconfiguration.html
        :param pulumi.Input[str] encryption_option: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-encryptionconfiguration.html#cfn-athena-workgroup-encryptionconfiguration-encryptionoption
        :param pulumi.Input[str] kms_key: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-encryptionconfiguration.html#cfn-athena-workgroup-encryptionconfiguration-kmskey
        """
        pulumi.set(__self__, "encryption_option", encryption_option)
        if kms_key is not None:
            pulumi.set(__self__, "kms_key", kms_key)

    @property
    @pulumi.getter(name="encryptionOption")
    def encryption_option(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-encryptionconfiguration.html#cfn-athena-workgroup-encryptionconfiguration-encryptionoption
        """
        return pulumi.get(self, "encryption_option")

    @encryption_option.setter
    def encryption_option(self, value: pulumi.Input[str]):
        pulumi.set(self, "encryption_option", value)

    @property
    @pulumi.getter(name="kmsKey")
    def kms_key(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-encryptionconfiguration.html#cfn-athena-workgroup-encryptionconfiguration-kmskey
        """
        return pulumi.get(self, "kms_key")

    @kms_key.setter
    def kms_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key", value)


@pulumi.input_type
class WorkGroupResultConfigurationArgs:
    def __init__(__self__, *,
                 encryption_configuration: Optional[pulumi.Input['WorkGroupEncryptionConfigurationArgs']] = None,
                 output_location: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-resultconfiguration.html
        :param pulumi.Input['WorkGroupEncryptionConfigurationArgs'] encryption_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-resultconfiguration.html#cfn-athena-workgroup-resultconfiguration-encryptionconfiguration
        :param pulumi.Input[str] output_location: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-resultconfiguration.html#cfn-athena-workgroup-resultconfiguration-outputlocation
        """
        if encryption_configuration is not None:
            pulumi.set(__self__, "encryption_configuration", encryption_configuration)
        if output_location is not None:
            pulumi.set(__self__, "output_location", output_location)

    @property
    @pulumi.getter(name="encryptionConfiguration")
    def encryption_configuration(self) -> Optional[pulumi.Input['WorkGroupEncryptionConfigurationArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-resultconfiguration.html#cfn-athena-workgroup-resultconfiguration-encryptionconfiguration
        """
        return pulumi.get(self, "encryption_configuration")

    @encryption_configuration.setter
    def encryption_configuration(self, value: Optional[pulumi.Input['WorkGroupEncryptionConfigurationArgs']]):
        pulumi.set(self, "encryption_configuration", value)

    @property
    @pulumi.getter(name="outputLocation")
    def output_location(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-resultconfiguration.html#cfn-athena-workgroup-resultconfiguration-outputlocation
        """
        return pulumi.get(self, "output_location")

    @output_location.setter
    def output_location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output_location", value)


@pulumi.input_type
class WorkGroupResultConfigurationUpdatesArgs:
    def __init__(__self__, *,
                 encryption_configuration: Optional[pulumi.Input['WorkGroupEncryptionConfigurationArgs']] = None,
                 output_location: Optional[pulumi.Input[str]] = None,
                 remove_encryption_configuration: Optional[pulumi.Input[bool]] = None,
                 remove_output_location: Optional[pulumi.Input[bool]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-resultconfigurationupdates.html
        :param pulumi.Input['WorkGroupEncryptionConfigurationArgs'] encryption_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-resultconfigurationupdates.html#cfn-athena-workgroup-resultconfigurationupdates-encryptionconfiguration
        :param pulumi.Input[str] output_location: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-resultconfigurationupdates.html#cfn-athena-workgroup-resultconfigurationupdates-outputlocation
        :param pulumi.Input[bool] remove_encryption_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-resultconfigurationupdates.html#cfn-athena-workgroup-resultconfigurationupdates-removeencryptionconfiguration
        :param pulumi.Input[bool] remove_output_location: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-resultconfigurationupdates.html#cfn-athena-workgroup-resultconfigurationupdates-removeoutputlocation
        """
        if encryption_configuration is not None:
            pulumi.set(__self__, "encryption_configuration", encryption_configuration)
        if output_location is not None:
            pulumi.set(__self__, "output_location", output_location)
        if remove_encryption_configuration is not None:
            pulumi.set(__self__, "remove_encryption_configuration", remove_encryption_configuration)
        if remove_output_location is not None:
            pulumi.set(__self__, "remove_output_location", remove_output_location)

    @property
    @pulumi.getter(name="encryptionConfiguration")
    def encryption_configuration(self) -> Optional[pulumi.Input['WorkGroupEncryptionConfigurationArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-resultconfigurationupdates.html#cfn-athena-workgroup-resultconfigurationupdates-encryptionconfiguration
        """
        return pulumi.get(self, "encryption_configuration")

    @encryption_configuration.setter
    def encryption_configuration(self, value: Optional[pulumi.Input['WorkGroupEncryptionConfigurationArgs']]):
        pulumi.set(self, "encryption_configuration", value)

    @property
    @pulumi.getter(name="outputLocation")
    def output_location(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-resultconfigurationupdates.html#cfn-athena-workgroup-resultconfigurationupdates-outputlocation
        """
        return pulumi.get(self, "output_location")

    @output_location.setter
    def output_location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output_location", value)

    @property
    @pulumi.getter(name="removeEncryptionConfiguration")
    def remove_encryption_configuration(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-resultconfigurationupdates.html#cfn-athena-workgroup-resultconfigurationupdates-removeencryptionconfiguration
        """
        return pulumi.get(self, "remove_encryption_configuration")

    @remove_encryption_configuration.setter
    def remove_encryption_configuration(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "remove_encryption_configuration", value)

    @property
    @pulumi.getter(name="removeOutputLocation")
    def remove_output_location(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-resultconfigurationupdates.html#cfn-athena-workgroup-resultconfigurationupdates-removeoutputlocation
        """
        return pulumi.get(self, "remove_output_location")

    @remove_output_location.setter
    def remove_output_location(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "remove_output_location", value)


@pulumi.input_type
class WorkGroupTagsArgs:
    def __init__(__self__, *,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-tags.html
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-tags.html#cfn-athena-workgroup-tags-tags
        """
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-tags.html#cfn-athena-workgroup-tags-tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class WorkGroupWorkGroupConfigurationArgs:
    def __init__(__self__, *,
                 bytes_scanned_cutoff_per_query: Optional[pulumi.Input[int]] = None,
                 enforce_work_group_configuration: Optional[pulumi.Input[bool]] = None,
                 publish_cloud_watch_metrics_enabled: Optional[pulumi.Input[bool]] = None,
                 requester_pays_enabled: Optional[pulumi.Input[bool]] = None,
                 result_configuration: Optional[pulumi.Input['WorkGroupResultConfigurationArgs']] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html
        :param pulumi.Input[int] bytes_scanned_cutoff_per_query: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html#cfn-athena-workgroup-workgroupconfiguration-bytesscannedcutoffperquery
        :param pulumi.Input[bool] enforce_work_group_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html#cfn-athena-workgroup-workgroupconfiguration-enforceworkgroupconfiguration
        :param pulumi.Input[bool] publish_cloud_watch_metrics_enabled: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html#cfn-athena-workgroup-workgroupconfiguration-publishcloudwatchmetricsenabled
        :param pulumi.Input[bool] requester_pays_enabled: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html#cfn-athena-workgroup-workgroupconfiguration-requesterpaysenabled
        :param pulumi.Input['WorkGroupResultConfigurationArgs'] result_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html#cfn-athena-workgroup-workgroupconfiguration-resultconfiguration
        """
        if bytes_scanned_cutoff_per_query is not None:
            pulumi.set(__self__, "bytes_scanned_cutoff_per_query", bytes_scanned_cutoff_per_query)
        if enforce_work_group_configuration is not None:
            pulumi.set(__self__, "enforce_work_group_configuration", enforce_work_group_configuration)
        if publish_cloud_watch_metrics_enabled is not None:
            pulumi.set(__self__, "publish_cloud_watch_metrics_enabled", publish_cloud_watch_metrics_enabled)
        if requester_pays_enabled is not None:
            pulumi.set(__self__, "requester_pays_enabled", requester_pays_enabled)
        if result_configuration is not None:
            pulumi.set(__self__, "result_configuration", result_configuration)

    @property
    @pulumi.getter(name="bytesScannedCutoffPerQuery")
    def bytes_scanned_cutoff_per_query(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html#cfn-athena-workgroup-workgroupconfiguration-bytesscannedcutoffperquery
        """
        return pulumi.get(self, "bytes_scanned_cutoff_per_query")

    @bytes_scanned_cutoff_per_query.setter
    def bytes_scanned_cutoff_per_query(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bytes_scanned_cutoff_per_query", value)

    @property
    @pulumi.getter(name="enforceWorkGroupConfiguration")
    def enforce_work_group_configuration(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html#cfn-athena-workgroup-workgroupconfiguration-enforceworkgroupconfiguration
        """
        return pulumi.get(self, "enforce_work_group_configuration")

    @enforce_work_group_configuration.setter
    def enforce_work_group_configuration(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enforce_work_group_configuration", value)

    @property
    @pulumi.getter(name="publishCloudWatchMetricsEnabled")
    def publish_cloud_watch_metrics_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html#cfn-athena-workgroup-workgroupconfiguration-publishcloudwatchmetricsenabled
        """
        return pulumi.get(self, "publish_cloud_watch_metrics_enabled")

    @publish_cloud_watch_metrics_enabled.setter
    def publish_cloud_watch_metrics_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "publish_cloud_watch_metrics_enabled", value)

    @property
    @pulumi.getter(name="requesterPaysEnabled")
    def requester_pays_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html#cfn-athena-workgroup-workgroupconfiguration-requesterpaysenabled
        """
        return pulumi.get(self, "requester_pays_enabled")

    @requester_pays_enabled.setter
    def requester_pays_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "requester_pays_enabled", value)

    @property
    @pulumi.getter(name="resultConfiguration")
    def result_configuration(self) -> Optional[pulumi.Input['WorkGroupResultConfigurationArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html#cfn-athena-workgroup-workgroupconfiguration-resultconfiguration
        """
        return pulumi.get(self, "result_configuration")

    @result_configuration.setter
    def result_configuration(self, value: Optional[pulumi.Input['WorkGroupResultConfigurationArgs']]):
        pulumi.set(self, "result_configuration", value)


@pulumi.input_type
class WorkGroupWorkGroupConfigurationUpdatesArgs:
    def __init__(__self__, *,
                 bytes_scanned_cutoff_per_query: Optional[pulumi.Input[int]] = None,
                 enforce_work_group_configuration: Optional[pulumi.Input[bool]] = None,
                 publish_cloud_watch_metrics_enabled: Optional[pulumi.Input[bool]] = None,
                 remove_bytes_scanned_cutoff_per_query: Optional[pulumi.Input[bool]] = None,
                 requester_pays_enabled: Optional[pulumi.Input[bool]] = None,
                 result_configuration_updates: Optional[pulumi.Input['WorkGroupResultConfigurationUpdatesArgs']] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html
        :param pulumi.Input[int] bytes_scanned_cutoff_per_query: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-bytesscannedcutoffperquery
        :param pulumi.Input[bool] enforce_work_group_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-enforceworkgroupconfiguration
        :param pulumi.Input[bool] publish_cloud_watch_metrics_enabled: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-publishcloudwatchmetricsenabled
        :param pulumi.Input[bool] remove_bytes_scanned_cutoff_per_query: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-removebytesscannedcutoffperquery
        :param pulumi.Input[bool] requester_pays_enabled: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-requesterpaysenabled
        :param pulumi.Input['WorkGroupResultConfigurationUpdatesArgs'] result_configuration_updates: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-resultconfigurationupdates
        """
        if bytes_scanned_cutoff_per_query is not None:
            pulumi.set(__self__, "bytes_scanned_cutoff_per_query", bytes_scanned_cutoff_per_query)
        if enforce_work_group_configuration is not None:
            pulumi.set(__self__, "enforce_work_group_configuration", enforce_work_group_configuration)
        if publish_cloud_watch_metrics_enabled is not None:
            pulumi.set(__self__, "publish_cloud_watch_metrics_enabled", publish_cloud_watch_metrics_enabled)
        if remove_bytes_scanned_cutoff_per_query is not None:
            pulumi.set(__self__, "remove_bytes_scanned_cutoff_per_query", remove_bytes_scanned_cutoff_per_query)
        if requester_pays_enabled is not None:
            pulumi.set(__self__, "requester_pays_enabled", requester_pays_enabled)
        if result_configuration_updates is not None:
            pulumi.set(__self__, "result_configuration_updates", result_configuration_updates)

    @property
    @pulumi.getter(name="bytesScannedCutoffPerQuery")
    def bytes_scanned_cutoff_per_query(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-bytesscannedcutoffperquery
        """
        return pulumi.get(self, "bytes_scanned_cutoff_per_query")

    @bytes_scanned_cutoff_per_query.setter
    def bytes_scanned_cutoff_per_query(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bytes_scanned_cutoff_per_query", value)

    @property
    @pulumi.getter(name="enforceWorkGroupConfiguration")
    def enforce_work_group_configuration(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-enforceworkgroupconfiguration
        """
        return pulumi.get(self, "enforce_work_group_configuration")

    @enforce_work_group_configuration.setter
    def enforce_work_group_configuration(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enforce_work_group_configuration", value)

    @property
    @pulumi.getter(name="publishCloudWatchMetricsEnabled")
    def publish_cloud_watch_metrics_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-publishcloudwatchmetricsenabled
        """
        return pulumi.get(self, "publish_cloud_watch_metrics_enabled")

    @publish_cloud_watch_metrics_enabled.setter
    def publish_cloud_watch_metrics_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "publish_cloud_watch_metrics_enabled", value)

    @property
    @pulumi.getter(name="removeBytesScannedCutoffPerQuery")
    def remove_bytes_scanned_cutoff_per_query(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-removebytesscannedcutoffperquery
        """
        return pulumi.get(self, "remove_bytes_scanned_cutoff_per_query")

    @remove_bytes_scanned_cutoff_per_query.setter
    def remove_bytes_scanned_cutoff_per_query(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "remove_bytes_scanned_cutoff_per_query", value)

    @property
    @pulumi.getter(name="requesterPaysEnabled")
    def requester_pays_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-requesterpaysenabled
        """
        return pulumi.get(self, "requester_pays_enabled")

    @requester_pays_enabled.setter
    def requester_pays_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "requester_pays_enabled", value)

    @property
    @pulumi.getter(name="resultConfigurationUpdates")
    def result_configuration_updates(self) -> Optional[pulumi.Input['WorkGroupResultConfigurationUpdatesArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-resultconfigurationupdates
        """
        return pulumi.get(self, "result_configuration_updates")

    @result_configuration_updates.setter
    def result_configuration_updates(self, value: Optional[pulumi.Input['WorkGroupResultConfigurationUpdatesArgs']]):
        pulumi.set(self, "result_configuration_updates", value)


