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
    'GetFuotaTaskResult',
    'AwaitableGetFuotaTaskResult',
    'get_fuota_task',
    'get_fuota_task_output',
]

@pulumi.output_type
class GetFuotaTaskResult:
    def __init__(__self__, arn=None, associate_multicast_group=None, associate_wireless_device=None, description=None, disassociate_multicast_group=None, disassociate_wireless_device=None, firmware_update_image=None, firmware_update_role=None, fuota_task_status=None, id=None, lo_ra_wan=None, name=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if associate_multicast_group and not isinstance(associate_multicast_group, str):
            raise TypeError("Expected argument 'associate_multicast_group' to be a str")
        pulumi.set(__self__, "associate_multicast_group", associate_multicast_group)
        if associate_wireless_device and not isinstance(associate_wireless_device, str):
            raise TypeError("Expected argument 'associate_wireless_device' to be a str")
        pulumi.set(__self__, "associate_wireless_device", associate_wireless_device)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disassociate_multicast_group and not isinstance(disassociate_multicast_group, str):
            raise TypeError("Expected argument 'disassociate_multicast_group' to be a str")
        pulumi.set(__self__, "disassociate_multicast_group", disassociate_multicast_group)
        if disassociate_wireless_device and not isinstance(disassociate_wireless_device, str):
            raise TypeError("Expected argument 'disassociate_wireless_device' to be a str")
        pulumi.set(__self__, "disassociate_wireless_device", disassociate_wireless_device)
        if firmware_update_image and not isinstance(firmware_update_image, str):
            raise TypeError("Expected argument 'firmware_update_image' to be a str")
        pulumi.set(__self__, "firmware_update_image", firmware_update_image)
        if firmware_update_role and not isinstance(firmware_update_role, str):
            raise TypeError("Expected argument 'firmware_update_role' to be a str")
        pulumi.set(__self__, "firmware_update_role", firmware_update_role)
        if fuota_task_status and not isinstance(fuota_task_status, str):
            raise TypeError("Expected argument 'fuota_task_status' to be a str")
        pulumi.set(__self__, "fuota_task_status", fuota_task_status)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lo_ra_wan and not isinstance(lo_ra_wan, dict):
            raise TypeError("Expected argument 'lo_ra_wan' to be a dict")
        pulumi.set(__self__, "lo_ra_wan", lo_ra_wan)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        FUOTA task arn. Returned after successful create.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="associateMulticastGroup")
    def associate_multicast_group(self) -> Optional[str]:
        """
        Multicast group to associate. Only for update request.
        """
        return pulumi.get(self, "associate_multicast_group")

    @property
    @pulumi.getter(name="associateWirelessDevice")
    def associate_wireless_device(self) -> Optional[str]:
        """
        Wireless device to associate. Only for update request.
        """
        return pulumi.get(self, "associate_wireless_device")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        FUOTA task description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="disassociateMulticastGroup")
    def disassociate_multicast_group(self) -> Optional[str]:
        """
        Multicast group to disassociate. Only for update request.
        """
        return pulumi.get(self, "disassociate_multicast_group")

    @property
    @pulumi.getter(name="disassociateWirelessDevice")
    def disassociate_wireless_device(self) -> Optional[str]:
        """
        Wireless device to disassociate. Only for update request.
        """
        return pulumi.get(self, "disassociate_wireless_device")

    @property
    @pulumi.getter(name="firmwareUpdateImage")
    def firmware_update_image(self) -> Optional[str]:
        """
        FUOTA task firmware update image binary S3 link
        """
        return pulumi.get(self, "firmware_update_image")

    @property
    @pulumi.getter(name="firmwareUpdateRole")
    def firmware_update_role(self) -> Optional[str]:
        """
        FUOTA task firmware IAM role for reading S3
        """
        return pulumi.get(self, "firmware_update_role")

    @property
    @pulumi.getter(name="fuotaTaskStatus")
    def fuota_task_status(self) -> Optional[str]:
        """
        FUOTA task status. Returned after successful read.
        """
        return pulumi.get(self, "fuota_task_status")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        FUOTA task id. Returned after successful create.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="loRaWAN")
    def lo_ra_wan(self) -> Optional['outputs.FuotaTaskLoRaWAN']:
        """
        FUOTA task LoRaWAN
        """
        return pulumi.get(self, "lo_ra_wan")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of FUOTA task
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.FuotaTaskTag']]:
        """
        A list of key-value pairs that contain metadata for the FUOTA task.
        """
        return pulumi.get(self, "tags")


class AwaitableGetFuotaTaskResult(GetFuotaTaskResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFuotaTaskResult(
            arn=self.arn,
            associate_multicast_group=self.associate_multicast_group,
            associate_wireless_device=self.associate_wireless_device,
            description=self.description,
            disassociate_multicast_group=self.disassociate_multicast_group,
            disassociate_wireless_device=self.disassociate_wireless_device,
            firmware_update_image=self.firmware_update_image,
            firmware_update_role=self.firmware_update_role,
            fuota_task_status=self.fuota_task_status,
            id=self.id,
            lo_ra_wan=self.lo_ra_wan,
            name=self.name,
            tags=self.tags)


def get_fuota_task(id: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFuotaTaskResult:
    """
    Create and manage FUOTA tasks.


    :param str id: FUOTA task id. Returned after successful create.
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:iotwireless:getFuotaTask', __args__, opts=opts, typ=GetFuotaTaskResult).value

    return AwaitableGetFuotaTaskResult(
        arn=__ret__.arn,
        associate_multicast_group=__ret__.associate_multicast_group,
        associate_wireless_device=__ret__.associate_wireless_device,
        description=__ret__.description,
        disassociate_multicast_group=__ret__.disassociate_multicast_group,
        disassociate_wireless_device=__ret__.disassociate_wireless_device,
        firmware_update_image=__ret__.firmware_update_image,
        firmware_update_role=__ret__.firmware_update_role,
        fuota_task_status=__ret__.fuota_task_status,
        id=__ret__.id,
        lo_ra_wan=__ret__.lo_ra_wan,
        name=__ret__.name,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_fuota_task)
def get_fuota_task_output(id: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFuotaTaskResult]:
    """
    Create and manage FUOTA tasks.


    :param str id: FUOTA task id. Returned after successful create.
    """
    ...
