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
    'GetClientCertificateResult',
    'AwaitableGetClientCertificateResult',
    'get_client_certificate',
    'get_client_certificate_output',
]

@pulumi.output_type
class GetClientCertificateResult:
    def __init__(__self__, client_certificate_id=None, description=None, tags=None):
        if client_certificate_id and not isinstance(client_certificate_id, str):
            raise TypeError("Expected argument 'client_certificate_id' to be a str")
        pulumi.set(__self__, "client_certificate_id", client_certificate_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="clientCertificateId")
    def client_certificate_id(self) -> Optional[str]:
        """
        The Primary Identifier of the Client Certficate, generated by a Create API Call
        """
        return pulumi.get(self, "client_certificate_id")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        A description of the client certificate.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.ClientCertificateTag']]:
        """
        An array of arbitrary tags (key-value pairs) to associate with the client certificate.
        """
        return pulumi.get(self, "tags")


class AwaitableGetClientCertificateResult(GetClientCertificateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClientCertificateResult(
            client_certificate_id=self.client_certificate_id,
            description=self.description,
            tags=self.tags)


def get_client_certificate(client_certificate_id: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClientCertificateResult:
    """
    Resource Type definition for AWS::ApiGateway::ClientCertificate


    :param str client_certificate_id: The Primary Identifier of the Client Certficate, generated by a Create API Call
    """
    __args__ = dict()
    __args__['clientCertificateId'] = client_certificate_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:apigateway:getClientCertificate', __args__, opts=opts, typ=GetClientCertificateResult).value

    return AwaitableGetClientCertificateResult(
        client_certificate_id=__ret__.client_certificate_id,
        description=__ret__.description,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_client_certificate)
def get_client_certificate_output(client_certificate_id: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClientCertificateResult]:
    """
    Resource Type definition for AWS::ApiGateway::ClientCertificate


    :param str client_certificate_id: The Primary Identifier of the Client Certficate, generated by a Create API Call
    """
    ...
