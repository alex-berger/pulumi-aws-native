# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetAccountResult',
    'AwaitableGetAccountResult',
    'get_account',
    'get_account_output',
]

@pulumi.output_type
class GetAccountResult:
    def __init__(__self__, cloud_watch_role_arn=None, id=None):
        if cloud_watch_role_arn and not isinstance(cloud_watch_role_arn, str):
            raise TypeError("Expected argument 'cloud_watch_role_arn' to be a str")
        pulumi.set(__self__, "cloud_watch_role_arn", cloud_watch_role_arn)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="cloudWatchRoleArn")
    def cloud_watch_role_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of an IAM role that has write access to CloudWatch Logs in your account.
        """
        return pulumi.get(self, "cloud_watch_role_arn")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        Primary identifier which is manually generated.
        """
        return pulumi.get(self, "id")


class AwaitableGetAccountResult(GetAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccountResult(
            cloud_watch_role_arn=self.cloud_watch_role_arn,
            id=self.id)


def get_account(id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccountResult:
    """
    Resource Type definition for AWS::ApiGateway::Account


    :param str id: Primary identifier which is manually generated.
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:apigateway:getAccount', __args__, opts=opts, typ=GetAccountResult).value

    return AwaitableGetAccountResult(
        cloud_watch_role_arn=__ret__.cloud_watch_role_arn,
        id=__ret__.id)


@_utilities.lift_output_func(get_account)
def get_account_output(id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAccountResult]:
    """
    Resource Type definition for AWS::ApiGateway::Account


    :param str id: Primary identifier which is manually generated.
    """
    ...
