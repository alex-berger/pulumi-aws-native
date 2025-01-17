# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetDBSecurityGroupIngressResult',
    'AwaitableGetDBSecurityGroupIngressResult',
    'get_db_security_group_ingress',
    'get_db_security_group_ingress_output',
]

@pulumi.output_type
class GetDBSecurityGroupIngressResult:
    def __init__(__self__, c_idrip=None, d_b_security_group_name=None, e_c2_security_group_id=None, e_c2_security_group_name=None, e_c2_security_group_owner_id=None, id=None):
        if c_idrip and not isinstance(c_idrip, str):
            raise TypeError("Expected argument 'c_idrip' to be a str")
        pulumi.set(__self__, "c_idrip", c_idrip)
        if d_b_security_group_name and not isinstance(d_b_security_group_name, str):
            raise TypeError("Expected argument 'd_b_security_group_name' to be a str")
        pulumi.set(__self__, "d_b_security_group_name", d_b_security_group_name)
        if e_c2_security_group_id and not isinstance(e_c2_security_group_id, str):
            raise TypeError("Expected argument 'e_c2_security_group_id' to be a str")
        pulumi.set(__self__, "e_c2_security_group_id", e_c2_security_group_id)
        if e_c2_security_group_name and not isinstance(e_c2_security_group_name, str):
            raise TypeError("Expected argument 'e_c2_security_group_name' to be a str")
        pulumi.set(__self__, "e_c2_security_group_name", e_c2_security_group_name)
        if e_c2_security_group_owner_id and not isinstance(e_c2_security_group_owner_id, str):
            raise TypeError("Expected argument 'e_c2_security_group_owner_id' to be a str")
        pulumi.set(__self__, "e_c2_security_group_owner_id", e_c2_security_group_owner_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="cIDRIP")
    def c_idrip(self) -> Optional[str]:
        return pulumi.get(self, "c_idrip")

    @property
    @pulumi.getter(name="dBSecurityGroupName")
    def d_b_security_group_name(self) -> Optional[str]:
        return pulumi.get(self, "d_b_security_group_name")

    @property
    @pulumi.getter(name="eC2SecurityGroupId")
    def e_c2_security_group_id(self) -> Optional[str]:
        return pulumi.get(self, "e_c2_security_group_id")

    @property
    @pulumi.getter(name="eC2SecurityGroupName")
    def e_c2_security_group_name(self) -> Optional[str]:
        return pulumi.get(self, "e_c2_security_group_name")

    @property
    @pulumi.getter(name="eC2SecurityGroupOwnerId")
    def e_c2_security_group_owner_id(self) -> Optional[str]:
        return pulumi.get(self, "e_c2_security_group_owner_id")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")


class AwaitableGetDBSecurityGroupIngressResult(GetDBSecurityGroupIngressResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDBSecurityGroupIngressResult(
            c_idrip=self.c_idrip,
            d_b_security_group_name=self.d_b_security_group_name,
            e_c2_security_group_id=self.e_c2_security_group_id,
            e_c2_security_group_name=self.e_c2_security_group_name,
            e_c2_security_group_owner_id=self.e_c2_security_group_owner_id,
            id=self.id)


def get_db_security_group_ingress(id: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDBSecurityGroupIngressResult:
    """
    Resource Type definition for AWS::RDS::DBSecurityGroupIngress
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:rds:getDBSecurityGroupIngress', __args__, opts=opts, typ=GetDBSecurityGroupIngressResult).value

    return AwaitableGetDBSecurityGroupIngressResult(
        c_idrip=__ret__.c_idrip,
        d_b_security_group_name=__ret__.d_b_security_group_name,
        e_c2_security_group_id=__ret__.e_c2_security_group_id,
        e_c2_security_group_name=__ret__.e_c2_security_group_name,
        e_c2_security_group_owner_id=__ret__.e_c2_security_group_owner_id,
        id=__ret__.id)


@_utilities.lift_output_func(get_db_security_group_ingress)
def get_db_security_group_ingress_output(id: Optional[pulumi.Input[str]] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDBSecurityGroupIngressResult]:
    """
    Resource Type definition for AWS::RDS::DBSecurityGroupIngress
    """
    ...
