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

__all__ = ['CustomMetricArgs', 'CustomMetric']

@pulumi.input_type
class CustomMetricArgs:
    def __init__(__self__, *,
                 metric_type: pulumi.Input['CustomMetricMetricType'],
                 display_name: Optional[pulumi.Input[str]] = None,
                 metric_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['CustomMetricTagArgs']]]] = None):
        """
        The set of arguments for constructing a CustomMetric resource.
        :param pulumi.Input['CustomMetricMetricType'] metric_type: The type of the custom metric. Types include string-list, ip-address-list, number-list, and number.
        :param pulumi.Input[str] display_name: Field represents a friendly name in the console for the custom metric; it doesn't have to be unique. Don't use this name as the metric identifier in the device metric report. Can be updated once defined.
        :param pulumi.Input[str] metric_name: The name of the custom metric. This will be used in the metric report submitted from the device/thing. Shouldn't begin with aws: . Cannot be updated once defined.
        :param pulumi.Input[Sequence[pulumi.Input['CustomMetricTagArgs']]] tags: An array of key-value pairs to apply to this resource.
        """
        pulumi.set(__self__, "metric_type", metric_type)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if metric_name is not None:
            pulumi.set(__self__, "metric_name", metric_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="metricType")
    def metric_type(self) -> pulumi.Input['CustomMetricMetricType']:
        """
        The type of the custom metric. Types include string-list, ip-address-list, number-list, and number.
        """
        return pulumi.get(self, "metric_type")

    @metric_type.setter
    def metric_type(self, value: pulumi.Input['CustomMetricMetricType']):
        pulumi.set(self, "metric_type", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Field represents a friendly name in the console for the custom metric; it doesn't have to be unique. Don't use this name as the metric identifier in the device metric report. Can be updated once defined.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="metricName")
    def metric_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the custom metric. This will be used in the metric report submitted from the device/thing. Shouldn't begin with aws: . Cannot be updated once defined.
        """
        return pulumi.get(self, "metric_name")

    @metric_name.setter
    def metric_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metric_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CustomMetricTagArgs']]]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CustomMetricTagArgs']]]]):
        pulumi.set(self, "tags", value)


class CustomMetric(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 metric_name: Optional[pulumi.Input[str]] = None,
                 metric_type: Optional[pulumi.Input['CustomMetricMetricType']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomMetricTagArgs']]]]] = None,
                 __props__=None):
        """
        A custom metric published by your devices to Device Defender.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Field represents a friendly name in the console for the custom metric; it doesn't have to be unique. Don't use this name as the metric identifier in the device metric report. Can be updated once defined.
        :param pulumi.Input[str] metric_name: The name of the custom metric. This will be used in the metric report submitted from the device/thing. Shouldn't begin with aws: . Cannot be updated once defined.
        :param pulumi.Input['CustomMetricMetricType'] metric_type: The type of the custom metric. Types include string-list, ip-address-list, number-list, and number.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomMetricTagArgs']]]] tags: An array of key-value pairs to apply to this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CustomMetricArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A custom metric published by your devices to Device Defender.

        :param str resource_name: The name of the resource.
        :param CustomMetricArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CustomMetricArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 metric_name: Optional[pulumi.Input[str]] = None,
                 metric_type: Optional[pulumi.Input['CustomMetricMetricType']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomMetricTagArgs']]]]] = None,
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
            __props__ = CustomMetricArgs.__new__(CustomMetricArgs)

            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["metric_name"] = metric_name
            if metric_type is None and not opts.urn:
                raise TypeError("Missing required property 'metric_type'")
            __props__.__dict__["metric_type"] = metric_type
            __props__.__dict__["tags"] = tags
            __props__.__dict__["metric_arn"] = None
        super(CustomMetric, __self__).__init__(
            'aws-native:iot:CustomMetric',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'CustomMetric':
        """
        Get an existing CustomMetric resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = CustomMetricArgs.__new__(CustomMetricArgs)

        __props__.__dict__["display_name"] = None
        __props__.__dict__["metric_arn"] = None
        __props__.__dict__["metric_name"] = None
        __props__.__dict__["metric_type"] = None
        __props__.__dict__["tags"] = None
        return CustomMetric(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        Field represents a friendly name in the console for the custom metric; it doesn't have to be unique. Don't use this name as the metric identifier in the device metric report. Can be updated once defined.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="metricArn")
    def metric_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Number (ARN) of the custom metric.
        """
        return pulumi.get(self, "metric_arn")

    @property
    @pulumi.getter(name="metricName")
    def metric_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the custom metric. This will be used in the metric report submitted from the device/thing. Shouldn't begin with aws: . Cannot be updated once defined.
        """
        return pulumi.get(self, "metric_name")

    @property
    @pulumi.getter(name="metricType")
    def metric_type(self) -> pulumi.Output['CustomMetricMetricType']:
        """
        The type of the custom metric. Types include string-list, ip-address-list, number-list, and number.
        """
        return pulumi.get(self, "metric_type")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.CustomMetricTag']]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

