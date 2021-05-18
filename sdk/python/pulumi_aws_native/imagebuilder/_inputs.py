# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'DistributionConfigurationDistributionArgs',
    'ImageImageTestsConfigurationArgs',
    'ImagePipelineImageTestsConfigurationArgs',
    'ImagePipelineScheduleArgs',
    'ImageRecipeComponentConfigurationArgs',
    'ImageRecipeEbsInstanceBlockDeviceSpecificationArgs',
    'ImageRecipeInstanceBlockDeviceMappingArgs',
]

@pulumi.input_type
class DistributionConfigurationDistributionArgs:
    def __init__(__self__, *,
                 region: pulumi.Input[str],
                 ami_distribution_configuration: Optional[pulumi.Input[Union[Any, str]]] = None,
                 license_configuration_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html
        :param pulumi.Input[str] region: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html#cfn-imagebuilder-distributionconfiguration-distribution-region
        :param pulumi.Input[Union[Any, str]] ami_distribution_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html#cfn-imagebuilder-distributionconfiguration-distribution-amidistributionconfiguration
        :param pulumi.Input[Sequence[pulumi.Input[str]]] license_configuration_arns: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html#cfn-imagebuilder-distributionconfiguration-distribution-licenseconfigurationarns
        """
        pulumi.set(__self__, "region", region)
        if ami_distribution_configuration is not None:
            pulumi.set(__self__, "ami_distribution_configuration", ami_distribution_configuration)
        if license_configuration_arns is not None:
            pulumi.set(__self__, "license_configuration_arns", license_configuration_arns)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html#cfn-imagebuilder-distributionconfiguration-distribution-region
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="amiDistributionConfiguration")
    def ami_distribution_configuration(self) -> Optional[pulumi.Input[Union[Any, str]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html#cfn-imagebuilder-distributionconfiguration-distribution-amidistributionconfiguration
        """
        return pulumi.get(self, "ami_distribution_configuration")

    @ami_distribution_configuration.setter
    def ami_distribution_configuration(self, value: Optional[pulumi.Input[Union[Any, str]]]):
        pulumi.set(self, "ami_distribution_configuration", value)

    @property
    @pulumi.getter(name="licenseConfigurationArns")
    def license_configuration_arns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html#cfn-imagebuilder-distributionconfiguration-distribution-licenseconfigurationarns
        """
        return pulumi.get(self, "license_configuration_arns")

    @license_configuration_arns.setter
    def license_configuration_arns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "license_configuration_arns", value)


@pulumi.input_type
class ImageImageTestsConfigurationArgs:
    def __init__(__self__, *,
                 image_tests_enabled: Optional[pulumi.Input[bool]] = None,
                 timeout_minutes: Optional[pulumi.Input[int]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-imagetestsconfiguration.html
        :param pulumi.Input[bool] image_tests_enabled: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-imagetestsconfiguration.html#cfn-imagebuilder-image-imagetestsconfiguration-imagetestsenabled
        :param pulumi.Input[int] timeout_minutes: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-imagetestsconfiguration.html#cfn-imagebuilder-image-imagetestsconfiguration-timeoutminutes
        """
        if image_tests_enabled is not None:
            pulumi.set(__self__, "image_tests_enabled", image_tests_enabled)
        if timeout_minutes is not None:
            pulumi.set(__self__, "timeout_minutes", timeout_minutes)

    @property
    @pulumi.getter(name="imageTestsEnabled")
    def image_tests_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-imagetestsconfiguration.html#cfn-imagebuilder-image-imagetestsconfiguration-imagetestsenabled
        """
        return pulumi.get(self, "image_tests_enabled")

    @image_tests_enabled.setter
    def image_tests_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "image_tests_enabled", value)

    @property
    @pulumi.getter(name="timeoutMinutes")
    def timeout_minutes(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-imagetestsconfiguration.html#cfn-imagebuilder-image-imagetestsconfiguration-timeoutminutes
        """
        return pulumi.get(self, "timeout_minutes")

    @timeout_minutes.setter
    def timeout_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_minutes", value)


@pulumi.input_type
class ImagePipelineImageTestsConfigurationArgs:
    def __init__(__self__, *,
                 image_tests_enabled: Optional[pulumi.Input[bool]] = None,
                 timeout_minutes: Optional[pulumi.Input[int]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-imagetestsconfiguration.html
        :param pulumi.Input[bool] image_tests_enabled: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-imagetestsconfiguration.html#cfn-imagebuilder-imagepipeline-imagetestsconfiguration-imagetestsenabled
        :param pulumi.Input[int] timeout_minutes: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-imagetestsconfiguration.html#cfn-imagebuilder-imagepipeline-imagetestsconfiguration-timeoutminutes
        """
        if image_tests_enabled is not None:
            pulumi.set(__self__, "image_tests_enabled", image_tests_enabled)
        if timeout_minutes is not None:
            pulumi.set(__self__, "timeout_minutes", timeout_minutes)

    @property
    @pulumi.getter(name="imageTestsEnabled")
    def image_tests_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-imagetestsconfiguration.html#cfn-imagebuilder-imagepipeline-imagetestsconfiguration-imagetestsenabled
        """
        return pulumi.get(self, "image_tests_enabled")

    @image_tests_enabled.setter
    def image_tests_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "image_tests_enabled", value)

    @property
    @pulumi.getter(name="timeoutMinutes")
    def timeout_minutes(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-imagetestsconfiguration.html#cfn-imagebuilder-imagepipeline-imagetestsconfiguration-timeoutminutes
        """
        return pulumi.get(self, "timeout_minutes")

    @timeout_minutes.setter
    def timeout_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_minutes", value)


@pulumi.input_type
class ImagePipelineScheduleArgs:
    def __init__(__self__, *,
                 pipeline_execution_start_condition: Optional[pulumi.Input[str]] = None,
                 schedule_expression: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-schedule.html
        :param pulumi.Input[str] pipeline_execution_start_condition: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-schedule.html#cfn-imagebuilder-imagepipeline-schedule-pipelineexecutionstartcondition
        :param pulumi.Input[str] schedule_expression: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-schedule.html#cfn-imagebuilder-imagepipeline-schedule-scheduleexpression
        """
        if pipeline_execution_start_condition is not None:
            pulumi.set(__self__, "pipeline_execution_start_condition", pipeline_execution_start_condition)
        if schedule_expression is not None:
            pulumi.set(__self__, "schedule_expression", schedule_expression)

    @property
    @pulumi.getter(name="pipelineExecutionStartCondition")
    def pipeline_execution_start_condition(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-schedule.html#cfn-imagebuilder-imagepipeline-schedule-pipelineexecutionstartcondition
        """
        return pulumi.get(self, "pipeline_execution_start_condition")

    @pipeline_execution_start_condition.setter
    def pipeline_execution_start_condition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pipeline_execution_start_condition", value)

    @property
    @pulumi.getter(name="scheduleExpression")
    def schedule_expression(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-schedule.html#cfn-imagebuilder-imagepipeline-schedule-scheduleexpression
        """
        return pulumi.get(self, "schedule_expression")

    @schedule_expression.setter
    def schedule_expression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule_expression", value)


@pulumi.input_type
class ImageRecipeComponentConfigurationArgs:
    def __init__(__self__, *,
                 component_arn: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-componentconfiguration.html
        :param pulumi.Input[str] component_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-componentconfiguration.html#cfn-imagebuilder-imagerecipe-componentconfiguration-componentarn
        """
        if component_arn is not None:
            pulumi.set(__self__, "component_arn", component_arn)

    @property
    @pulumi.getter(name="componentArn")
    def component_arn(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-componentconfiguration.html#cfn-imagebuilder-imagerecipe-componentconfiguration-componentarn
        """
        return pulumi.get(self, "component_arn")

    @component_arn.setter
    def component_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "component_arn", value)


@pulumi.input_type
class ImageRecipeEbsInstanceBlockDeviceSpecificationArgs:
    def __init__(__self__, *,
                 delete_on_termination: Optional[pulumi.Input[bool]] = None,
                 encrypted: Optional[pulumi.Input[bool]] = None,
                 iops: Optional[pulumi.Input[int]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 volume_size: Optional[pulumi.Input[int]] = None,
                 volume_type: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html
        :param pulumi.Input[bool] delete_on_termination: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-deleteontermination
        :param pulumi.Input[bool] encrypted: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-encrypted
        :param pulumi.Input[int] iops: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-iops
        :param pulumi.Input[str] kms_key_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-kmskeyid
        :param pulumi.Input[str] snapshot_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-snapshotid
        :param pulumi.Input[int] volume_size: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-volumesize
        :param pulumi.Input[str] volume_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-volumetype
        """
        if delete_on_termination is not None:
            pulumi.set(__self__, "delete_on_termination", delete_on_termination)
        if encrypted is not None:
            pulumi.set(__self__, "encrypted", encrypted)
        if iops is not None:
            pulumi.set(__self__, "iops", iops)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if snapshot_id is not None:
            pulumi.set(__self__, "snapshot_id", snapshot_id)
        if volume_size is not None:
            pulumi.set(__self__, "volume_size", volume_size)
        if volume_type is not None:
            pulumi.set(__self__, "volume_type", volume_type)

    @property
    @pulumi.getter(name="deleteOnTermination")
    def delete_on_termination(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-deleteontermination
        """
        return pulumi.get(self, "delete_on_termination")

    @delete_on_termination.setter
    def delete_on_termination(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_on_termination", value)

    @property
    @pulumi.getter
    def encrypted(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-encrypted
        """
        return pulumi.get(self, "encrypted")

    @encrypted.setter
    def encrypted(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "encrypted", value)

    @property
    @pulumi.getter
    def iops(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-iops
        """
        return pulumi.get(self, "iops")

    @iops.setter
    def iops(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "iops", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-kmskeyid
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-snapshotid
        """
        return pulumi.get(self, "snapshot_id")

    @snapshot_id.setter
    def snapshot_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_id", value)

    @property
    @pulumi.getter(name="volumeSize")
    def volume_size(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-volumesize
        """
        return pulumi.get(self, "volume_size")

    @volume_size.setter
    def volume_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "volume_size", value)

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-volumetype
        """
        return pulumi.get(self, "volume_type")

    @volume_type.setter
    def volume_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "volume_type", value)


@pulumi.input_type
class ImageRecipeInstanceBlockDeviceMappingArgs:
    def __init__(__self__, *,
                 device_name: Optional[pulumi.Input[str]] = None,
                 ebs: Optional[pulumi.Input['ImageRecipeEbsInstanceBlockDeviceSpecificationArgs']] = None,
                 no_device: Optional[pulumi.Input[str]] = None,
                 virtual_name: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html
        :param pulumi.Input[str] device_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html#cfn-imagebuilder-imagerecipe-instanceblockdevicemapping-devicename
        :param pulumi.Input['ImageRecipeEbsInstanceBlockDeviceSpecificationArgs'] ebs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html#cfn-imagebuilder-imagerecipe-instanceblockdevicemapping-ebs
        :param pulumi.Input[str] no_device: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html#cfn-imagebuilder-imagerecipe-instanceblockdevicemapping-nodevice
        :param pulumi.Input[str] virtual_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html#cfn-imagebuilder-imagerecipe-instanceblockdevicemapping-virtualname
        """
        if device_name is not None:
            pulumi.set(__self__, "device_name", device_name)
        if ebs is not None:
            pulumi.set(__self__, "ebs", ebs)
        if no_device is not None:
            pulumi.set(__self__, "no_device", no_device)
        if virtual_name is not None:
            pulumi.set(__self__, "virtual_name", virtual_name)

    @property
    @pulumi.getter(name="deviceName")
    def device_name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html#cfn-imagebuilder-imagerecipe-instanceblockdevicemapping-devicename
        """
        return pulumi.get(self, "device_name")

    @device_name.setter
    def device_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_name", value)

    @property
    @pulumi.getter
    def ebs(self) -> Optional[pulumi.Input['ImageRecipeEbsInstanceBlockDeviceSpecificationArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html#cfn-imagebuilder-imagerecipe-instanceblockdevicemapping-ebs
        """
        return pulumi.get(self, "ebs")

    @ebs.setter
    def ebs(self, value: Optional[pulumi.Input['ImageRecipeEbsInstanceBlockDeviceSpecificationArgs']]):
        pulumi.set(self, "ebs", value)

    @property
    @pulumi.getter(name="noDevice")
    def no_device(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html#cfn-imagebuilder-imagerecipe-instanceblockdevicemapping-nodevice
        """
        return pulumi.get(self, "no_device")

    @no_device.setter
    def no_device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "no_device", value)

    @property
    @pulumi.getter(name="virtualName")
    def virtual_name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html#cfn-imagebuilder-imagerecipe-instanceblockdevicemapping-virtualname
        """
        return pulumi.get(self, "virtual_name")

    @virtual_name.setter
    def virtual_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "virtual_name", value)


