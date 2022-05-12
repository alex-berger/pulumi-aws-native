# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'GetIPAMPoolResult',
    'AwaitableGetIPAMPoolResult',
    'get_ipam_pool',
    'get_ipam_pool_output',
]

@pulumi.output_type
class GetIPAMPoolResult:
    def __init__(__self__, allocation_default_netmask_length=None, allocation_max_netmask_length=None, allocation_min_netmask_length=None, allocation_resource_tags=None, arn=None, auto_import=None, description=None, ipam_arn=None, ipam_pool_id=None, ipam_scope_arn=None, ipam_scope_type=None, pool_depth=None, provisioned_cidrs=None, state=None, state_message=None, tags=None):
        if allocation_default_netmask_length and not isinstance(allocation_default_netmask_length, int):
            raise TypeError("Expected argument 'allocation_default_netmask_length' to be a int")
        pulumi.set(__self__, "allocation_default_netmask_length", allocation_default_netmask_length)
        if allocation_max_netmask_length and not isinstance(allocation_max_netmask_length, int):
            raise TypeError("Expected argument 'allocation_max_netmask_length' to be a int")
        pulumi.set(__self__, "allocation_max_netmask_length", allocation_max_netmask_length)
        if allocation_min_netmask_length and not isinstance(allocation_min_netmask_length, int):
            raise TypeError("Expected argument 'allocation_min_netmask_length' to be a int")
        pulumi.set(__self__, "allocation_min_netmask_length", allocation_min_netmask_length)
        if allocation_resource_tags and not isinstance(allocation_resource_tags, list):
            raise TypeError("Expected argument 'allocation_resource_tags' to be a list")
        pulumi.set(__self__, "allocation_resource_tags", allocation_resource_tags)
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if auto_import and not isinstance(auto_import, bool):
            raise TypeError("Expected argument 'auto_import' to be a bool")
        pulumi.set(__self__, "auto_import", auto_import)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if ipam_arn and not isinstance(ipam_arn, str):
            raise TypeError("Expected argument 'ipam_arn' to be a str")
        pulumi.set(__self__, "ipam_arn", ipam_arn)
        if ipam_pool_id and not isinstance(ipam_pool_id, str):
            raise TypeError("Expected argument 'ipam_pool_id' to be a str")
        pulumi.set(__self__, "ipam_pool_id", ipam_pool_id)
        if ipam_scope_arn and not isinstance(ipam_scope_arn, str):
            raise TypeError("Expected argument 'ipam_scope_arn' to be a str")
        pulumi.set(__self__, "ipam_scope_arn", ipam_scope_arn)
        if ipam_scope_type and not isinstance(ipam_scope_type, str):
            raise TypeError("Expected argument 'ipam_scope_type' to be a str")
        pulumi.set(__self__, "ipam_scope_type", ipam_scope_type)
        if pool_depth and not isinstance(pool_depth, int):
            raise TypeError("Expected argument 'pool_depth' to be a int")
        pulumi.set(__self__, "pool_depth", pool_depth)
        if provisioned_cidrs and not isinstance(provisioned_cidrs, list):
            raise TypeError("Expected argument 'provisioned_cidrs' to be a list")
        pulumi.set(__self__, "provisioned_cidrs", provisioned_cidrs)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if state_message and not isinstance(state_message, str):
            raise TypeError("Expected argument 'state_message' to be a str")
        pulumi.set(__self__, "state_message", state_message)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="allocationDefaultNetmaskLength")
    def allocation_default_netmask_length(self) -> Optional[int]:
        """
        The default netmask length for allocations made from this pool. This value is used when the netmask length of an allocation isn't specified.
        """
        return pulumi.get(self, "allocation_default_netmask_length")

    @property
    @pulumi.getter(name="allocationMaxNetmaskLength")
    def allocation_max_netmask_length(self) -> Optional[int]:
        """
        The maximum allowed netmask length for allocations made from this pool.
        """
        return pulumi.get(self, "allocation_max_netmask_length")

    @property
    @pulumi.getter(name="allocationMinNetmaskLength")
    def allocation_min_netmask_length(self) -> Optional[int]:
        """
        The minimum allowed netmask length for allocations made from this pool.
        """
        return pulumi.get(self, "allocation_min_netmask_length")

    @property
    @pulumi.getter(name="allocationResourceTags")
    def allocation_resource_tags(self) -> Optional[Sequence['outputs.IPAMPoolTag']]:
        """
        When specified, an allocation will not be allowed unless a resource has a matching set of tags.
        """
        return pulumi.get(self, "allocation_resource_tags")

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the IPAM Pool.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="autoImport")
    def auto_import(self) -> Optional[bool]:
        """
        Determines what to do if IPAM discovers resources that haven't been assigned an allocation. If set to true, an allocation will be made automatically.
        """
        return pulumi.get(self, "auto_import")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ipamArn")
    def ipam_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the IPAM this pool is a part of.
        """
        return pulumi.get(self, "ipam_arn")

    @property
    @pulumi.getter(name="ipamPoolId")
    def ipam_pool_id(self) -> Optional[str]:
        """
        Id of the IPAM Pool.
        """
        return pulumi.get(self, "ipam_pool_id")

    @property
    @pulumi.getter(name="ipamScopeArn")
    def ipam_scope_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the scope this pool is a part of.
        """
        return pulumi.get(self, "ipam_scope_arn")

    @property
    @pulumi.getter(name="ipamScopeType")
    def ipam_scope_type(self) -> Optional['IPAMPoolIpamScopeType']:
        """
        Determines whether this scope contains publicly routable space or space for a private network
        """
        return pulumi.get(self, "ipam_scope_type")

    @property
    @pulumi.getter(name="poolDepth")
    def pool_depth(self) -> Optional[int]:
        """
        The depth of this pool in the source pool hierarchy.
        """
        return pulumi.get(self, "pool_depth")

    @property
    @pulumi.getter(name="provisionedCidrs")
    def provisioned_cidrs(self) -> Optional[Sequence['outputs.IPAMPoolProvisionedCidr']]:
        """
        A list of cidrs representing the address space available for allocation in this pool.
        """
        return pulumi.get(self, "provisioned_cidrs")

    @property
    @pulumi.getter
    def state(self) -> Optional['IPAMPoolState']:
        """
        The state of this pool. This can be one of the following values: "create-in-progress", "create-complete", "modify-in-progress", "modify-complete", "delete-in-progress", or "delete-complete"
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="stateMessage")
    def state_message(self) -> Optional[str]:
        """
        An explanation of how the pool arrived at it current state.
        """
        return pulumi.get(self, "state_message")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.IPAMPoolTag']]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetIPAMPoolResult(GetIPAMPoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIPAMPoolResult(
            allocation_default_netmask_length=self.allocation_default_netmask_length,
            allocation_max_netmask_length=self.allocation_max_netmask_length,
            allocation_min_netmask_length=self.allocation_min_netmask_length,
            allocation_resource_tags=self.allocation_resource_tags,
            arn=self.arn,
            auto_import=self.auto_import,
            description=self.description,
            ipam_arn=self.ipam_arn,
            ipam_pool_id=self.ipam_pool_id,
            ipam_scope_arn=self.ipam_scope_arn,
            ipam_scope_type=self.ipam_scope_type,
            pool_depth=self.pool_depth,
            provisioned_cidrs=self.provisioned_cidrs,
            state=self.state,
            state_message=self.state_message,
            tags=self.tags)


def get_ipam_pool(ipam_pool_id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIPAMPoolResult:
    """
    Resource Schema of AWS::EC2::IPAMPool Type


    :param str ipam_pool_id: Id of the IPAM Pool.
    """
    __args__ = dict()
    __args__['ipamPoolId'] = ipam_pool_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:ec2:getIPAMPool', __args__, opts=opts, typ=GetIPAMPoolResult).value

    return AwaitableGetIPAMPoolResult(
        allocation_default_netmask_length=__ret__.allocation_default_netmask_length,
        allocation_max_netmask_length=__ret__.allocation_max_netmask_length,
        allocation_min_netmask_length=__ret__.allocation_min_netmask_length,
        allocation_resource_tags=__ret__.allocation_resource_tags,
        arn=__ret__.arn,
        auto_import=__ret__.auto_import,
        description=__ret__.description,
        ipam_arn=__ret__.ipam_arn,
        ipam_pool_id=__ret__.ipam_pool_id,
        ipam_scope_arn=__ret__.ipam_scope_arn,
        ipam_scope_type=__ret__.ipam_scope_type,
        pool_depth=__ret__.pool_depth,
        provisioned_cidrs=__ret__.provisioned_cidrs,
        state=__ret__.state,
        state_message=__ret__.state_message,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_ipam_pool)
def get_ipam_pool_output(ipam_pool_id: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIPAMPoolResult]:
    """
    Resource Schema of AWS::EC2::IPAMPool Type


    :param str ipam_pool_id: Id of the IPAM Pool.
    """
    ...
