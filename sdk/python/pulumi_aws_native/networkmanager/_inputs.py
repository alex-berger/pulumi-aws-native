# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'DeviceLocationArgs',
    'DeviceTagArgs',
    'GlobalNetworkTagArgs',
    'LinkBandwidthArgs',
    'LinkTagArgs',
    'SiteLocationArgs',
    'SiteTagArgs',
]

@pulumi.input_type
class DeviceLocationArgs:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 latitude: Optional[pulumi.Input[str]] = None,
                 longitude: Optional[pulumi.Input[str]] = None):
        """
        The site location.
        :param pulumi.Input[str] address: The physical address.
        :param pulumi.Input[str] latitude: The latitude.
        :param pulumi.Input[str] longitude: The longitude.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if latitude is not None:
            pulumi.set(__self__, "latitude", latitude)
        if longitude is not None:
            pulumi.set(__self__, "longitude", longitude)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        The physical address.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def latitude(self) -> Optional[pulumi.Input[str]]:
        """
        The latitude.
        """
        return pulumi.get(self, "latitude")

    @latitude.setter
    def latitude(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "latitude", value)

    @property
    @pulumi.getter
    def longitude(self) -> Optional[pulumi.Input[str]]:
        """
        The longitude.
        """
        return pulumi.get(self, "longitude")

    @longitude.setter
    def longitude(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "longitude", value)


@pulumi.input_type
class DeviceTagArgs:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        A key-value pair to associate with a device resource.
        """
        if key is not None:
            pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class GlobalNetworkTagArgs:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        A key-value pair to associate with a global network resource.
        """
        if key is not None:
            pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class LinkBandwidthArgs:
    def __init__(__self__, *,
                 download_speed: Optional[pulumi.Input[int]] = None,
                 upload_speed: Optional[pulumi.Input[int]] = None):
        """
        The bandwidth for the link.
        :param pulumi.Input[int] download_speed: Download speed in Mbps.
        :param pulumi.Input[int] upload_speed: Upload speed in Mbps.
        """
        if download_speed is not None:
            pulumi.set(__self__, "download_speed", download_speed)
        if upload_speed is not None:
            pulumi.set(__self__, "upload_speed", upload_speed)

    @property
    @pulumi.getter(name="downloadSpeed")
    def download_speed(self) -> Optional[pulumi.Input[int]]:
        """
        Download speed in Mbps.
        """
        return pulumi.get(self, "download_speed")

    @download_speed.setter
    def download_speed(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "download_speed", value)

    @property
    @pulumi.getter(name="uploadSpeed")
    def upload_speed(self) -> Optional[pulumi.Input[int]]:
        """
        Upload speed in Mbps.
        """
        return pulumi.get(self, "upload_speed")

    @upload_speed.setter
    def upload_speed(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "upload_speed", value)


@pulumi.input_type
class LinkTagArgs:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        A key-value pair to associate with a link resource.
        """
        if key is not None:
            pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class SiteLocationArgs:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 latitude: Optional[pulumi.Input[str]] = None,
                 longitude: Optional[pulumi.Input[str]] = None):
        """
        The location of the site
        :param pulumi.Input[str] address: The physical address.
        :param pulumi.Input[str] latitude: The latitude.
        :param pulumi.Input[str] longitude: The longitude.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if latitude is not None:
            pulumi.set(__self__, "latitude", latitude)
        if longitude is not None:
            pulumi.set(__self__, "longitude", longitude)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        The physical address.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def latitude(self) -> Optional[pulumi.Input[str]]:
        """
        The latitude.
        """
        return pulumi.get(self, "latitude")

    @latitude.setter
    def latitude(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "latitude", value)

    @property
    @pulumi.getter
    def longitude(self) -> Optional[pulumi.Input[str]]:
        """
        The longitude.
        """
        return pulumi.get(self, "longitude")

    @longitude.setter
    def longitude(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "longitude", value)


@pulumi.input_type
class SiteTagArgs:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        A key-value pair to associate with a site resource.
        """
        if key is not None:
            pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


