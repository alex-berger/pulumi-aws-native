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
    'GetWebACLResult',
    'AwaitableGetWebACLResult',
    'get_web_acl',
    'get_web_acl_output',
]

@pulumi.output_type
class GetWebACLResult:
    def __init__(__self__, default_action=None, id=None, rules=None):
        if default_action and not isinstance(default_action, dict):
            raise TypeError("Expected argument 'default_action' to be a dict")
        pulumi.set(__self__, "default_action", default_action)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> Optional['outputs.WebACLAction']:
        return pulumi.get(self, "default_action")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def rules(self) -> Optional[Sequence['outputs.WebACLRule']]:
        return pulumi.get(self, "rules")


class AwaitableGetWebACLResult(GetWebACLResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWebACLResult(
            default_action=self.default_action,
            id=self.id,
            rules=self.rules)


def get_web_acl(id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWebACLResult:
    """
    Resource Type definition for AWS::WAFRegional::WebACL
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:wafregional:getWebACL', __args__, opts=opts, typ=GetWebACLResult).value

    return AwaitableGetWebACLResult(
        default_action=__ret__.default_action,
        id=__ret__.id,
        rules=__ret__.rules)


@_utilities.lift_output_func(get_web_acl)
def get_web_acl_output(id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWebACLResult]:
    """
    Resource Type definition for AWS::WAFRegional::WebACL
    """
    ...
