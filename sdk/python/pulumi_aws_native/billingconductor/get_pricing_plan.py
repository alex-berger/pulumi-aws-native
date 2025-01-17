# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetPricingPlanResult',
    'AwaitableGetPricingPlanResult',
    'get_pricing_plan',
    'get_pricing_plan_output',
]

@pulumi.output_type
class GetPricingPlanResult:
    def __init__(__self__, arn=None, creation_time=None, description=None, last_modified_time=None, name=None, pricing_rule_arns=None, size=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if creation_time and not isinstance(creation_time, int):
            raise TypeError("Expected argument 'creation_time' to be a int")
        pulumi.set(__self__, "creation_time", creation_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if last_modified_time and not isinstance(last_modified_time, int):
            raise TypeError("Expected argument 'last_modified_time' to be a int")
        pulumi.set(__self__, "last_modified_time", last_modified_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if pricing_rule_arns and not isinstance(pricing_rule_arns, list):
            raise TypeError("Expected argument 'pricing_rule_arns' to be a list")
        pulumi.set(__self__, "pricing_rule_arns", pricing_rule_arns)
        if size and not isinstance(size, int):
            raise TypeError("Expected argument 'size' to be a int")
        pulumi.set(__self__, "size", size)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        Pricing Plan ARN
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[int]:
        """
        Creation timestamp in UNIX epoch time format
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> Optional[int]:
        """
        Latest modified timestamp in UNIX epoch time format
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pricingRuleArns")
    def pricing_rule_arns(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "pricing_rule_arns")

    @property
    @pulumi.getter
    def size(self) -> Optional[int]:
        """
        Number of associated pricing rules
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.PricingPlanTag']]:
        return pulumi.get(self, "tags")


class AwaitableGetPricingPlanResult(GetPricingPlanResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPricingPlanResult(
            arn=self.arn,
            creation_time=self.creation_time,
            description=self.description,
            last_modified_time=self.last_modified_time,
            name=self.name,
            pricing_rule_arns=self.pricing_rule_arns,
            size=self.size,
            tags=self.tags)


def get_pricing_plan(arn: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPricingPlanResult:
    """
    Pricing Plan enables you to customize your billing details consistent with the usage that accrues in each of your billing groups.


    :param str arn: Pricing Plan ARN
    """
    __args__ = dict()
    __args__['arn'] = arn
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:billingconductor:getPricingPlan', __args__, opts=opts, typ=GetPricingPlanResult).value

    return AwaitableGetPricingPlanResult(
        arn=__ret__.arn,
        creation_time=__ret__.creation_time,
        description=__ret__.description,
        last_modified_time=__ret__.last_modified_time,
        name=__ret__.name,
        pricing_rule_arns=__ret__.pricing_rule_arns,
        size=__ret__.size,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_pricing_plan)
def get_pricing_plan_output(arn: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPricingPlanResult]:
    """
    Pricing Plan enables you to customize your billing details consistent with the usage that accrues in each of your billing groups.


    :param str arn: Pricing Plan ARN
    """
    ...
