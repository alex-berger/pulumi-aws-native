# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetResourcePolicyResult',
    'AwaitableGetResourcePolicyResult',
    'get_resource_policy',
    'get_resource_policy_output',
]

@pulumi.output_type
class GetResourcePolicyResult:
    def __init__(__self__, policy_document=None):
        if policy_document and not isinstance(policy_document, str):
            raise TypeError("Expected argument 'policy_document' to be a str")
        pulumi.set(__self__, "policy_document", policy_document)

    @property
    @pulumi.getter(name="policyDocument")
    def policy_document(self) -> Optional[str]:
        """
        The policy document
        """
        return pulumi.get(self, "policy_document")


class AwaitableGetResourcePolicyResult(GetResourcePolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResourcePolicyResult(
            policy_document=self.policy_document)


def get_resource_policy(policy_name: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResourcePolicyResult:
    """
    The resource schema for AWSLogs ResourcePolicy


    :param str policy_name: A name for resource policy
    """
    __args__ = dict()
    __args__['policyName'] = policy_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:logs:getResourcePolicy', __args__, opts=opts, typ=GetResourcePolicyResult).value

    return AwaitableGetResourcePolicyResult(
        policy_document=__ret__.policy_document)


@_utilities.lift_output_func(get_resource_policy)
def get_resource_policy_output(policy_name: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetResourcePolicyResult]:
    """
    The resource schema for AWSLogs ResourcePolicy


    :param str policy_name: A name for resource policy
    """
    ...
