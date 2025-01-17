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
    'GetDBProxyTargetGroupResult',
    'AwaitableGetDBProxyTargetGroupResult',
    'get_db_proxy_target_group',
    'get_db_proxy_target_group_output',
]

@pulumi.output_type
class GetDBProxyTargetGroupResult:
    def __init__(__self__, connection_pool_configuration_info=None, d_b_cluster_identifiers=None, d_b_instance_identifiers=None, target_group_arn=None):
        if connection_pool_configuration_info and not isinstance(connection_pool_configuration_info, dict):
            raise TypeError("Expected argument 'connection_pool_configuration_info' to be a dict")
        pulumi.set(__self__, "connection_pool_configuration_info", connection_pool_configuration_info)
        if d_b_cluster_identifiers and not isinstance(d_b_cluster_identifiers, list):
            raise TypeError("Expected argument 'd_b_cluster_identifiers' to be a list")
        pulumi.set(__self__, "d_b_cluster_identifiers", d_b_cluster_identifiers)
        if d_b_instance_identifiers and not isinstance(d_b_instance_identifiers, list):
            raise TypeError("Expected argument 'd_b_instance_identifiers' to be a list")
        pulumi.set(__self__, "d_b_instance_identifiers", d_b_instance_identifiers)
        if target_group_arn and not isinstance(target_group_arn, str):
            raise TypeError("Expected argument 'target_group_arn' to be a str")
        pulumi.set(__self__, "target_group_arn", target_group_arn)

    @property
    @pulumi.getter(name="connectionPoolConfigurationInfo")
    def connection_pool_configuration_info(self) -> Optional['outputs.DBProxyTargetGroupConnectionPoolConfigurationInfoFormat']:
        return pulumi.get(self, "connection_pool_configuration_info")

    @property
    @pulumi.getter(name="dBClusterIdentifiers")
    def d_b_cluster_identifiers(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "d_b_cluster_identifiers")

    @property
    @pulumi.getter(name="dBInstanceIdentifiers")
    def d_b_instance_identifiers(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "d_b_instance_identifiers")

    @property
    @pulumi.getter(name="targetGroupArn")
    def target_group_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) representing the target group.
        """
        return pulumi.get(self, "target_group_arn")


class AwaitableGetDBProxyTargetGroupResult(GetDBProxyTargetGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDBProxyTargetGroupResult(
            connection_pool_configuration_info=self.connection_pool_configuration_info,
            d_b_cluster_identifiers=self.d_b_cluster_identifiers,
            d_b_instance_identifiers=self.d_b_instance_identifiers,
            target_group_arn=self.target_group_arn)


def get_db_proxy_target_group(target_group_arn: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDBProxyTargetGroupResult:
    """
    Resource schema for AWS::RDS::DBProxyTargetGroup


    :param str target_group_arn: The Amazon Resource Name (ARN) representing the target group.
    """
    __args__ = dict()
    __args__['targetGroupArn'] = target_group_arn
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:rds:getDBProxyTargetGroup', __args__, opts=opts, typ=GetDBProxyTargetGroupResult).value

    return AwaitableGetDBProxyTargetGroupResult(
        connection_pool_configuration_info=__ret__.connection_pool_configuration_info,
        d_b_cluster_identifiers=__ret__.d_b_cluster_identifiers,
        d_b_instance_identifiers=__ret__.d_b_instance_identifiers,
        target_group_arn=__ret__.target_group_arn)


@_utilities.lift_output_func(get_db_proxy_target_group)
def get_db_proxy_target_group_output(target_group_arn: Optional[pulumi.Input[str]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDBProxyTargetGroupResult]:
    """
    Resource schema for AWS::RDS::DBProxyTargetGroup


    :param str target_group_arn: The Amazon Resource Name (ARN) representing the target group.
    """
    ...
