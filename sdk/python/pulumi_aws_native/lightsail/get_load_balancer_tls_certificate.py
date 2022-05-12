# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetLoadBalancerTlsCertificateResult',
    'AwaitableGetLoadBalancerTlsCertificateResult',
    'get_load_balancer_tls_certificate',
    'get_load_balancer_tls_certificate_output',
]

@pulumi.output_type
class GetLoadBalancerTlsCertificateResult:
    def __init__(__self__, is_attached=None, load_balancer_tls_certificate_arn=None, status=None):
        if is_attached and not isinstance(is_attached, bool):
            raise TypeError("Expected argument 'is_attached' to be a bool")
        pulumi.set(__self__, "is_attached", is_attached)
        if load_balancer_tls_certificate_arn and not isinstance(load_balancer_tls_certificate_arn, str):
            raise TypeError("Expected argument 'load_balancer_tls_certificate_arn' to be a str")
        pulumi.set(__self__, "load_balancer_tls_certificate_arn", load_balancer_tls_certificate_arn)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="isAttached")
    def is_attached(self) -> Optional[bool]:
        """
        When true, the SSL/TLS certificate is attached to the Lightsail load balancer.
        """
        return pulumi.get(self, "is_attached")

    @property
    @pulumi.getter(name="loadBalancerTlsCertificateArn")
    def load_balancer_tls_certificate_arn(self) -> Optional[str]:
        return pulumi.get(self, "load_balancer_tls_certificate_arn")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The validation status of the SSL/TLS certificate.
        """
        return pulumi.get(self, "status")


class AwaitableGetLoadBalancerTlsCertificateResult(GetLoadBalancerTlsCertificateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLoadBalancerTlsCertificateResult(
            is_attached=self.is_attached,
            load_balancer_tls_certificate_arn=self.load_balancer_tls_certificate_arn,
            status=self.status)


def get_load_balancer_tls_certificate(certificate_name: Optional[str] = None,
                                      load_balancer_name: Optional[str] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLoadBalancerTlsCertificateResult:
    """
    Resource Type definition for AWS::Lightsail::LoadBalancerTlsCertificate


    :param str certificate_name: The SSL/TLS certificate name.
    :param str load_balancer_name: The name of your load balancer.
    """
    __args__ = dict()
    __args__['certificateName'] = certificate_name
    __args__['loadBalancerName'] = load_balancer_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:lightsail:getLoadBalancerTlsCertificate', __args__, opts=opts, typ=GetLoadBalancerTlsCertificateResult).value

    return AwaitableGetLoadBalancerTlsCertificateResult(
        is_attached=__ret__.is_attached,
        load_balancer_tls_certificate_arn=__ret__.load_balancer_tls_certificate_arn,
        status=__ret__.status)


@_utilities.lift_output_func(get_load_balancer_tls_certificate)
def get_load_balancer_tls_certificate_output(certificate_name: Optional[pulumi.Input[str]] = None,
                                             load_balancer_name: Optional[pulumi.Input[str]] = None,
                                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLoadBalancerTlsCertificateResult]:
    """
    Resource Type definition for AWS::Lightsail::LoadBalancerTlsCertificate


    :param str certificate_name: The SSL/TLS certificate name.
    :param str load_balancer_name: The name of your load balancer.
    """
    ...
