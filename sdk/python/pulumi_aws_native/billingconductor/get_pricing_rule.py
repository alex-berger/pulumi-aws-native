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
    'GetPricingRuleResult',
    'AwaitableGetPricingRuleResult',
    'get_pricing_rule',
    'get_pricing_rule_output',
]

@pulumi.output_type
class GetPricingRuleResult:
    def __init__(__self__, arn=None, associated_pricing_plan_count=None, creation_time=None, description=None, last_modified_time=None, modifier_percentage=None, name=None, tags=None, type=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if associated_pricing_plan_count and not isinstance(associated_pricing_plan_count, int):
            raise TypeError("Expected argument 'associated_pricing_plan_count' to be a int")
        pulumi.set(__self__, "associated_pricing_plan_count", associated_pricing_plan_count)
        if creation_time and not isinstance(creation_time, int):
            raise TypeError("Expected argument 'creation_time' to be a int")
        pulumi.set(__self__, "creation_time", creation_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if last_modified_time and not isinstance(last_modified_time, int):
            raise TypeError("Expected argument 'last_modified_time' to be a int")
        pulumi.set(__self__, "last_modified_time", last_modified_time)
        if modifier_percentage and not isinstance(modifier_percentage, float):
            raise TypeError("Expected argument 'modifier_percentage' to be a float")
        pulumi.set(__self__, "modifier_percentage", modifier_percentage)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        Pricing rule ARN
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="associatedPricingPlanCount")
    def associated_pricing_plan_count(self) -> Optional[int]:
        """
        The number of pricing plans associated with pricing rule
        """
        return pulumi.get(self, "associated_pricing_plan_count")

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
        """
        Pricing rule description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> Optional[int]:
        """
        Latest modified timestamp in UNIX epoch time format
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter(name="modifierPercentage")
    def modifier_percentage(self) -> Optional[float]:
        """
        Pricing rule modifier percentage
        """
        return pulumi.get(self, "modifier_percentage")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Pricing rule name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.PricingRuleTag']]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> Optional['PricingRuleType']:
        """
        One of MARKUP or DISCOUNT that describes the direction of the rate that is applied to a pricing plan.
        """
        return pulumi.get(self, "type")


class AwaitableGetPricingRuleResult(GetPricingRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPricingRuleResult(
            arn=self.arn,
            associated_pricing_plan_count=self.associated_pricing_plan_count,
            creation_time=self.creation_time,
            description=self.description,
            last_modified_time=self.last_modified_time,
            modifier_percentage=self.modifier_percentage,
            name=self.name,
            tags=self.tags,
            type=self.type)


def get_pricing_rule(arn: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPricingRuleResult:
    """
    A markup/discount that is defined for a specific set of services that can later be associated with a pricing plan.


    :param str arn: Pricing rule ARN
    """
    __args__ = dict()
    __args__['arn'] = arn
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:billingconductor:getPricingRule', __args__, opts=opts, typ=GetPricingRuleResult).value

    return AwaitableGetPricingRuleResult(
        arn=__ret__.arn,
        associated_pricing_plan_count=__ret__.associated_pricing_plan_count,
        creation_time=__ret__.creation_time,
        description=__ret__.description,
        last_modified_time=__ret__.last_modified_time,
        modifier_percentage=__ret__.modifier_percentage,
        name=__ret__.name,
        tags=__ret__.tags,
        type=__ret__.type)


@_utilities.lift_output_func(get_pricing_rule)
def get_pricing_rule_output(arn: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPricingRuleResult]:
    """
    A markup/discount that is defined for a specific set of services that can later be associated with a pricing plan.


    :param str arn: Pricing rule ARN
    """
    ...
