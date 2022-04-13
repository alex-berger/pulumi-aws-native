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

__all__ = ['DatasetArgs', 'Dataset']

@pulumi.input_type
class DatasetArgs:
    def __init__(__self__, *,
                 actions: pulumi.Input[Sequence[pulumi.Input['DatasetActionArgs']]],
                 content_delivery_rules: Optional[pulumi.Input[Sequence[pulumi.Input['DatasetContentDeliveryRuleArgs']]]] = None,
                 dataset_name: Optional[pulumi.Input[str]] = None,
                 late_data_rules: Optional[pulumi.Input[Sequence[pulumi.Input['DatasetLateDataRuleArgs']]]] = None,
                 retention_period: Optional[pulumi.Input['DatasetRetentionPeriodArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['DatasetTagArgs']]]] = None,
                 triggers: Optional[pulumi.Input[Sequence[pulumi.Input['DatasetTriggerArgs']]]] = None,
                 versioning_configuration: Optional[pulumi.Input['DatasetVersioningConfigurationArgs']] = None):
        """
        The set of arguments for constructing a Dataset resource.
        """
        pulumi.set(__self__, "actions", actions)
        if content_delivery_rules is not None:
            pulumi.set(__self__, "content_delivery_rules", content_delivery_rules)
        if dataset_name is not None:
            pulumi.set(__self__, "dataset_name", dataset_name)
        if late_data_rules is not None:
            pulumi.set(__self__, "late_data_rules", late_data_rules)
        if retention_period is not None:
            pulumi.set(__self__, "retention_period", retention_period)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if triggers is not None:
            pulumi.set(__self__, "triggers", triggers)
        if versioning_configuration is not None:
            pulumi.set(__self__, "versioning_configuration", versioning_configuration)

    @property
    @pulumi.getter
    def actions(self) -> pulumi.Input[Sequence[pulumi.Input['DatasetActionArgs']]]:
        return pulumi.get(self, "actions")

    @actions.setter
    def actions(self, value: pulumi.Input[Sequence[pulumi.Input['DatasetActionArgs']]]):
        pulumi.set(self, "actions", value)

    @property
    @pulumi.getter(name="contentDeliveryRules")
    def content_delivery_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DatasetContentDeliveryRuleArgs']]]]:
        return pulumi.get(self, "content_delivery_rules")

    @content_delivery_rules.setter
    def content_delivery_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DatasetContentDeliveryRuleArgs']]]]):
        pulumi.set(self, "content_delivery_rules", value)

    @property
    @pulumi.getter(name="datasetName")
    def dataset_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dataset_name")

    @dataset_name.setter
    def dataset_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dataset_name", value)

    @property
    @pulumi.getter(name="lateDataRules")
    def late_data_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DatasetLateDataRuleArgs']]]]:
        return pulumi.get(self, "late_data_rules")

    @late_data_rules.setter
    def late_data_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DatasetLateDataRuleArgs']]]]):
        pulumi.set(self, "late_data_rules", value)

    @property
    @pulumi.getter(name="retentionPeriod")
    def retention_period(self) -> Optional[pulumi.Input['DatasetRetentionPeriodArgs']]:
        return pulumi.get(self, "retention_period")

    @retention_period.setter
    def retention_period(self, value: Optional[pulumi.Input['DatasetRetentionPeriodArgs']]):
        pulumi.set(self, "retention_period", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DatasetTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DatasetTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def triggers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DatasetTriggerArgs']]]]:
        return pulumi.get(self, "triggers")

    @triggers.setter
    def triggers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DatasetTriggerArgs']]]]):
        pulumi.set(self, "triggers", value)

    @property
    @pulumi.getter(name="versioningConfiguration")
    def versioning_configuration(self) -> Optional[pulumi.Input['DatasetVersioningConfigurationArgs']]:
        return pulumi.get(self, "versioning_configuration")

    @versioning_configuration.setter
    def versioning_configuration(self, value: Optional[pulumi.Input['DatasetVersioningConfigurationArgs']]):
        pulumi.set(self, "versioning_configuration", value)


class Dataset(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetActionArgs']]]]] = None,
                 content_delivery_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetContentDeliveryRuleArgs']]]]] = None,
                 dataset_name: Optional[pulumi.Input[str]] = None,
                 late_data_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetLateDataRuleArgs']]]]] = None,
                 retention_period: Optional[pulumi.Input[pulumi.InputType['DatasetRetentionPeriodArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetTagArgs']]]]] = None,
                 triggers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetTriggerArgs']]]]] = None,
                 versioning_configuration: Optional[pulumi.Input[pulumi.InputType['DatasetVersioningConfigurationArgs']]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::IoTAnalytics::Dataset

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatasetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::IoTAnalytics::Dataset

        :param str resource_name: The name of the resource.
        :param DatasetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatasetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetActionArgs']]]]] = None,
                 content_delivery_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetContentDeliveryRuleArgs']]]]] = None,
                 dataset_name: Optional[pulumi.Input[str]] = None,
                 late_data_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetLateDataRuleArgs']]]]] = None,
                 retention_period: Optional[pulumi.Input[pulumi.InputType['DatasetRetentionPeriodArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetTagArgs']]]]] = None,
                 triggers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetTriggerArgs']]]]] = None,
                 versioning_configuration: Optional[pulumi.Input[pulumi.InputType['DatasetVersioningConfigurationArgs']]] = None,
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
            __props__ = DatasetArgs.__new__(DatasetArgs)

            if actions is None and not opts.urn:
                raise TypeError("Missing required property 'actions'")
            __props__.__dict__["actions"] = actions
            __props__.__dict__["content_delivery_rules"] = content_delivery_rules
            __props__.__dict__["dataset_name"] = dataset_name
            __props__.__dict__["late_data_rules"] = late_data_rules
            __props__.__dict__["retention_period"] = retention_period
            __props__.__dict__["tags"] = tags
            __props__.__dict__["triggers"] = triggers
            __props__.__dict__["versioning_configuration"] = versioning_configuration
        super(Dataset, __self__).__init__(
            'aws-native:iotanalytics:Dataset',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Dataset':
        """
        Get an existing Dataset resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DatasetArgs.__new__(DatasetArgs)

        __props__.__dict__["actions"] = None
        __props__.__dict__["content_delivery_rules"] = None
        __props__.__dict__["dataset_name"] = None
        __props__.__dict__["late_data_rules"] = None
        __props__.__dict__["retention_period"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["triggers"] = None
        __props__.__dict__["versioning_configuration"] = None
        return Dataset(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def actions(self) -> pulumi.Output[Sequence['outputs.DatasetAction']]:
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter(name="contentDeliveryRules")
    def content_delivery_rules(self) -> pulumi.Output[Optional[Sequence['outputs.DatasetContentDeliveryRule']]]:
        return pulumi.get(self, "content_delivery_rules")

    @property
    @pulumi.getter(name="datasetName")
    def dataset_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dataset_name")

    @property
    @pulumi.getter(name="lateDataRules")
    def late_data_rules(self) -> pulumi.Output[Optional[Sequence['outputs.DatasetLateDataRule']]]:
        return pulumi.get(self, "late_data_rules")

    @property
    @pulumi.getter(name="retentionPeriod")
    def retention_period(self) -> pulumi.Output[Optional['outputs.DatasetRetentionPeriod']]:
        return pulumi.get(self, "retention_period")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.DatasetTag']]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def triggers(self) -> pulumi.Output[Optional[Sequence['outputs.DatasetTrigger']]]:
        return pulumi.get(self, "triggers")

    @property
    @pulumi.getter(name="versioningConfiguration")
    def versioning_configuration(self) -> pulumi.Output[Optional['outputs.DatasetVersioningConfiguration']]:
        return pulumi.get(self, "versioning_configuration")

