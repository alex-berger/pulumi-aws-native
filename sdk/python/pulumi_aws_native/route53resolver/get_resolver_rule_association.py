# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetResolverRuleAssociationResult',
    'AwaitableGetResolverRuleAssociationResult',
    'get_resolver_rule_association',
    'get_resolver_rule_association_output',
]

@pulumi.output_type
class GetResolverRuleAssociationResult:
    def __init__(__self__, resolver_rule_association_id=None):
        if resolver_rule_association_id and not isinstance(resolver_rule_association_id, str):
            raise TypeError("Expected argument 'resolver_rule_association_id' to be a str")
        pulumi.set(__self__, "resolver_rule_association_id", resolver_rule_association_id)

    @property
    @pulumi.getter(name="resolverRuleAssociationId")
    def resolver_rule_association_id(self) -> Optional[str]:
        """
        Primary Identifier for Resolver Rule Association
        """
        return pulumi.get(self, "resolver_rule_association_id")


class AwaitableGetResolverRuleAssociationResult(GetResolverRuleAssociationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResolverRuleAssociationResult(
            resolver_rule_association_id=self.resolver_rule_association_id)


def get_resolver_rule_association(resolver_rule_association_id: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResolverRuleAssociationResult:
    """
    Resource Type definition for AWS::Route53Resolver::ResolverRuleAssociation


    :param str resolver_rule_association_id: Primary Identifier for Resolver Rule Association
    """
    __args__ = dict()
    __args__['resolverRuleAssociationId'] = resolver_rule_association_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:route53resolver:getResolverRuleAssociation', __args__, opts=opts, typ=GetResolverRuleAssociationResult).value

    return AwaitableGetResolverRuleAssociationResult(
        resolver_rule_association_id=__ret__.resolver_rule_association_id)


@_utilities.lift_output_func(get_resolver_rule_association)
def get_resolver_rule_association_output(resolver_rule_association_id: Optional[pulumi.Input[str]] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetResolverRuleAssociationResult]:
    """
    Resource Type definition for AWS::Route53Resolver::ResolverRuleAssociation


    :param str resolver_rule_association_id: Primary Identifier for Resolver Rule Association
    """
    ...
