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
    'GetPartnerAccountResult',
    'AwaitableGetPartnerAccountResult',
    'get_partner_account',
    'get_partner_account_output',
]

@pulumi.output_type
class GetPartnerAccountResult:
    def __init__(__self__, account_linked=None, arn=None, fingerprint=None, partner_type=None, sidewalk=None, sidewalk_response=None, sidewalk_update=None, tags=None):
        if account_linked and not isinstance(account_linked, bool):
            raise TypeError("Expected argument 'account_linked' to be a bool")
        pulumi.set(__self__, "account_linked", account_linked)
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if fingerprint and not isinstance(fingerprint, str):
            raise TypeError("Expected argument 'fingerprint' to be a str")
        pulumi.set(__self__, "fingerprint", fingerprint)
        if partner_type and not isinstance(partner_type, str):
            raise TypeError("Expected argument 'partner_type' to be a str")
        pulumi.set(__self__, "partner_type", partner_type)
        if sidewalk and not isinstance(sidewalk, dict):
            raise TypeError("Expected argument 'sidewalk' to be a dict")
        pulumi.set(__self__, "sidewalk", sidewalk)
        if sidewalk_response and not isinstance(sidewalk_response, dict):
            raise TypeError("Expected argument 'sidewalk_response' to be a dict")
        pulumi.set(__self__, "sidewalk_response", sidewalk_response)
        if sidewalk_update and not isinstance(sidewalk_update, dict):
            raise TypeError("Expected argument 'sidewalk_update' to be a dict")
        pulumi.set(__self__, "sidewalk_update", sidewalk_update)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="accountLinked")
    def account_linked(self) -> Optional[bool]:
        """
        Whether the partner account is linked to the AWS account.
        """
        return pulumi.get(self, "account_linked")

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        PartnerAccount arn. Returned after successful create.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def fingerprint(self) -> Optional[str]:
        """
        The fingerprint of the Sidewalk application server private key.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter(name="partnerType")
    def partner_type(self) -> Optional['PartnerAccountPartnerType']:
        """
        The partner type
        """
        return pulumi.get(self, "partner_type")

    @property
    @pulumi.getter
    def sidewalk(self) -> Optional['outputs.PartnerAccountSidewalkAccountInfo']:
        """
        The Sidewalk account credentials.
        """
        return pulumi.get(self, "sidewalk")

    @property
    @pulumi.getter(name="sidewalkResponse")
    def sidewalk_response(self) -> Optional['outputs.PartnerAccountSidewalkAccountInfoWithFingerprint']:
        """
        The Sidewalk account credentials.
        """
        return pulumi.get(self, "sidewalk_response")

    @property
    @pulumi.getter(name="sidewalkUpdate")
    def sidewalk_update(self) -> Optional['outputs.PartnerAccountSidewalkUpdateAccount']:
        """
        The Sidewalk account credentials.
        """
        return pulumi.get(self, "sidewalk_update")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.PartnerAccountTag']]:
        """
        A list of key-value pairs that contain metadata for the destination.
        """
        return pulumi.get(self, "tags")


class AwaitableGetPartnerAccountResult(GetPartnerAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPartnerAccountResult(
            account_linked=self.account_linked,
            arn=self.arn,
            fingerprint=self.fingerprint,
            partner_type=self.partner_type,
            sidewalk=self.sidewalk,
            sidewalk_response=self.sidewalk_response,
            sidewalk_update=self.sidewalk_update,
            tags=self.tags)


def get_partner_account(partner_account_id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPartnerAccountResult:
    """
    Create and manage partner account


    :param str partner_account_id: The partner account ID to disassociate from the AWS account
    """
    __args__ = dict()
    __args__['partnerAccountId'] = partner_account_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:iotwireless:getPartnerAccount', __args__, opts=opts, typ=GetPartnerAccountResult).value

    return AwaitableGetPartnerAccountResult(
        account_linked=__ret__.account_linked,
        arn=__ret__.arn,
        fingerprint=__ret__.fingerprint,
        partner_type=__ret__.partner_type,
        sidewalk=__ret__.sidewalk,
        sidewalk_response=__ret__.sidewalk_response,
        sidewalk_update=__ret__.sidewalk_update,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_partner_account)
def get_partner_account_output(partner_account_id: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPartnerAccountResult]:
    """
    Create and manage partner account


    :param str partner_account_id: The partner account ID to disassociate from the AWS account
    """
    ...
