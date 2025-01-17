# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetApplicationEntitlementAssociationResult',
    'AwaitableGetApplicationEntitlementAssociationResult',
    'get_application_entitlement_association',
    'get_application_entitlement_association_output',
]

@pulumi.output_type
class GetApplicationEntitlementAssociationResult:
    def __init__(__self__):


class AwaitableGetApplicationEntitlementAssociationResult(GetApplicationEntitlementAssociationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationEntitlementAssociationResult(
)


def get_application_entitlement_association(application_identifier: Optional[str] = None,
                                            entitlement_name: Optional[str] = None,
                                            stack_name: Optional[str] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationEntitlementAssociationResult:
    """
    Resource Type definition for AWS::AppStream::ApplicationEntitlementAssociation
    """
    __args__ = dict()
    __args__['applicationIdentifier'] = application_identifier
    __args__['entitlementName'] = entitlement_name
    __args__['stackName'] = stack_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:appstream:getApplicationEntitlementAssociation', __args__, opts=opts, typ=GetApplicationEntitlementAssociationResult).value

    return AwaitableGetApplicationEntitlementAssociationResult(


@_utilities.lift_output_func(get_application_entitlement_association)
def get_application_entitlement_association_output(application_identifier: Optional[pulumi.Input[str]] = None,
                                                   entitlement_name: Optional[pulumi.Input[str]] = None,
                                                   stack_name: Optional[pulumi.Input[str]] = None,
                                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetApplicationEntitlementAssociationResult]:
    """
    Resource Type definition for AWS::AppStream::ApplicationEntitlementAssociation
    """
    ...
