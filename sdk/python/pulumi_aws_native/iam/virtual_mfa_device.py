# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['VirtualMFADeviceArgs', 'VirtualMFADevice']

@pulumi.input_type
class VirtualMFADeviceArgs:
    def __init__(__self__, *,
                 users: pulumi.Input[Sequence[pulumi.Input[str]]],
                 path: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['VirtualMFADeviceTagArgs']]]] = None,
                 virtual_mfa_device_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VirtualMFADevice resource.
        """
        pulumi.set(__self__, "users", users)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if virtual_mfa_device_name is not None:
            pulumi.set(__self__, "virtual_mfa_device_name", virtual_mfa_device_name)

    @property
    @pulumi.getter
    def users(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "users")

    @users.setter
    def users(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "users", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VirtualMFADeviceTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VirtualMFADeviceTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="virtualMfaDeviceName")
    def virtual_mfa_device_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "virtual_mfa_device_name")

    @virtual_mfa_device_name.setter
    def virtual_mfa_device_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "virtual_mfa_device_name", value)


class VirtualMFADevice(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualMFADeviceTagArgs']]]]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 virtual_mfa_device_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::IAM::VirtualMFADevice

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VirtualMFADeviceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::IAM::VirtualMFADevice

        :param str resource_name: The name of the resource.
        :param VirtualMFADeviceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VirtualMFADeviceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualMFADeviceTagArgs']]]]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 virtual_mfa_device_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VirtualMFADeviceArgs.__new__(VirtualMFADeviceArgs)

            __props__.__dict__["path"] = path
            __props__.__dict__["tags"] = tags
            if users is None and not opts.urn:
                raise TypeError("Missing required property 'users'")
            __props__.__dict__["users"] = users
            __props__.__dict__["virtual_mfa_device_name"] = virtual_mfa_device_name
            __props__.__dict__["serial_number"] = None
        super(VirtualMFADevice, __self__).__init__(
            'aws-native:iam:VirtualMFADevice',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'VirtualMFADevice':
        """
        Get an existing VirtualMFADevice resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = VirtualMFADeviceArgs.__new__(VirtualMFADeviceArgs)

        __props__.__dict__["path"] = None
        __props__.__dict__["serial_number"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["users"] = None
        __props__.__dict__["virtual_mfa_device_name"] = None
        return VirtualMFADevice(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> pulumi.Output[str]:
        return pulumi.get(self, "serial_number")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.VirtualMFADeviceTag']]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def users(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "users")

    @property
    @pulumi.getter(name="virtualMfaDeviceName")
    def virtual_mfa_device_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "virtual_mfa_device_name")

