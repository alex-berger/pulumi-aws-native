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
    'GetStorageVirtualMachineResult',
    'AwaitableGetStorageVirtualMachineResult',
    'get_storage_virtual_machine',
    'get_storage_virtual_machine_output',
]

@pulumi.output_type
class GetStorageVirtualMachineResult:
    def __init__(__self__, active_directory_configuration=None, resource_arn=None, storage_virtual_machine_id=None, svm_admin_password=None, tags=None, u_uid=None):
        if active_directory_configuration and not isinstance(active_directory_configuration, dict):
            raise TypeError("Expected argument 'active_directory_configuration' to be a dict")
        pulumi.set(__self__, "active_directory_configuration", active_directory_configuration)
        if resource_arn and not isinstance(resource_arn, str):
            raise TypeError("Expected argument 'resource_arn' to be a str")
        pulumi.set(__self__, "resource_arn", resource_arn)
        if storage_virtual_machine_id and not isinstance(storage_virtual_machine_id, str):
            raise TypeError("Expected argument 'storage_virtual_machine_id' to be a str")
        pulumi.set(__self__, "storage_virtual_machine_id", storage_virtual_machine_id)
        if svm_admin_password and not isinstance(svm_admin_password, str):
            raise TypeError("Expected argument 'svm_admin_password' to be a str")
        pulumi.set(__self__, "svm_admin_password", svm_admin_password)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if u_uid and not isinstance(u_uid, str):
            raise TypeError("Expected argument 'u_uid' to be a str")
        pulumi.set(__self__, "u_uid", u_uid)

    @property
    @pulumi.getter(name="activeDirectoryConfiguration")
    def active_directory_configuration(self) -> Optional['outputs.StorageVirtualMachineActiveDirectoryConfiguration']:
        return pulumi.get(self, "active_directory_configuration")

    @property
    @pulumi.getter(name="resourceARN")
    def resource_arn(self) -> Optional[str]:
        return pulumi.get(self, "resource_arn")

    @property
    @pulumi.getter(name="storageVirtualMachineId")
    def storage_virtual_machine_id(self) -> Optional[str]:
        return pulumi.get(self, "storage_virtual_machine_id")

    @property
    @pulumi.getter(name="svmAdminPassword")
    def svm_admin_password(self) -> Optional[str]:
        return pulumi.get(self, "svm_admin_password")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.StorageVirtualMachineTag']]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="uUID")
    def u_uid(self) -> Optional[str]:
        return pulumi.get(self, "u_uid")


class AwaitableGetStorageVirtualMachineResult(GetStorageVirtualMachineResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStorageVirtualMachineResult(
            active_directory_configuration=self.active_directory_configuration,
            resource_arn=self.resource_arn,
            storage_virtual_machine_id=self.storage_virtual_machine_id,
            svm_admin_password=self.svm_admin_password,
            tags=self.tags,
            u_uid=self.u_uid)


def get_storage_virtual_machine(storage_virtual_machine_id: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStorageVirtualMachineResult:
    """
    Resource Type definition for AWS::FSx::StorageVirtualMachine
    """
    __args__ = dict()
    __args__['storageVirtualMachineId'] = storage_virtual_machine_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:fsx:getStorageVirtualMachine', __args__, opts=opts, typ=GetStorageVirtualMachineResult).value

    return AwaitableGetStorageVirtualMachineResult(
        active_directory_configuration=__ret__.active_directory_configuration,
        resource_arn=__ret__.resource_arn,
        storage_virtual_machine_id=__ret__.storage_virtual_machine_id,
        svm_admin_password=__ret__.svm_admin_password,
        tags=__ret__.tags,
        u_uid=__ret__.u_uid)


@_utilities.lift_output_func(get_storage_virtual_machine)
def get_storage_virtual_machine_output(storage_virtual_machine_id: Optional[pulumi.Input[str]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetStorageVirtualMachineResult]:
    """
    Resource Type definition for AWS::FSx::StorageVirtualMachine
    """
    ...
