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
    'GetServerResult',
    'AwaitableGetServerResult',
    'get_server',
    'get_server_output',
]

@pulumi.output_type
class GetServerResult:
    def __init__(__self__, arn=None, backup_retention_count=None, disable_automated_backup=None, endpoint=None, id=None, preferred_backup_window=None, preferred_maintenance_window=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if backup_retention_count and not isinstance(backup_retention_count, int):
            raise TypeError("Expected argument 'backup_retention_count' to be a int")
        pulumi.set(__self__, "backup_retention_count", backup_retention_count)
        if disable_automated_backup and not isinstance(disable_automated_backup, bool):
            raise TypeError("Expected argument 'disable_automated_backup' to be a bool")
        pulumi.set(__self__, "disable_automated_backup", disable_automated_backup)
        if endpoint and not isinstance(endpoint, str):
            raise TypeError("Expected argument 'endpoint' to be a str")
        pulumi.set(__self__, "endpoint", endpoint)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if preferred_backup_window and not isinstance(preferred_backup_window, str):
            raise TypeError("Expected argument 'preferred_backup_window' to be a str")
        pulumi.set(__self__, "preferred_backup_window", preferred_backup_window)
        if preferred_maintenance_window and not isinstance(preferred_maintenance_window, str):
            raise TypeError("Expected argument 'preferred_maintenance_window' to be a str")
        pulumi.set(__self__, "preferred_maintenance_window", preferred_maintenance_window)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="backupRetentionCount")
    def backup_retention_count(self) -> Optional[int]:
        return pulumi.get(self, "backup_retention_count")

    @property
    @pulumi.getter(name="disableAutomatedBackup")
    def disable_automated_backup(self) -> Optional[bool]:
        return pulumi.get(self, "disable_automated_backup")

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[str]:
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="preferredBackupWindow")
    def preferred_backup_window(self) -> Optional[str]:
        return pulumi.get(self, "preferred_backup_window")

    @property
    @pulumi.getter(name="preferredMaintenanceWindow")
    def preferred_maintenance_window(self) -> Optional[str]:
        return pulumi.get(self, "preferred_maintenance_window")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.ServerTag']]:
        return pulumi.get(self, "tags")


class AwaitableGetServerResult(GetServerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServerResult(
            arn=self.arn,
            backup_retention_count=self.backup_retention_count,
            disable_automated_backup=self.disable_automated_backup,
            endpoint=self.endpoint,
            id=self.id,
            preferred_backup_window=self.preferred_backup_window,
            preferred_maintenance_window=self.preferred_maintenance_window,
            tags=self.tags)


def get_server(server_name: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServerResult:
    """
    Resource Type definition for AWS::OpsWorksCM::Server
    """
    __args__ = dict()
    __args__['serverName'] = server_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:opsworkscm:getServer', __args__, opts=opts, typ=GetServerResult).value

    return AwaitableGetServerResult(
        arn=__ret__.arn,
        backup_retention_count=__ret__.backup_retention_count,
        disable_automated_backup=__ret__.disable_automated_backup,
        endpoint=__ret__.endpoint,
        id=__ret__.id,
        preferred_backup_window=__ret__.preferred_backup_window,
        preferred_maintenance_window=__ret__.preferred_maintenance_window,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_server)
def get_server_output(server_name: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServerResult]:
    """
    Resource Type definition for AWS::OpsWorksCM::Server
    """
    ...
