# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetDeploymentResult',
    'AwaitableGetDeploymentResult',
    'get_deployment',
    'get_deployment_output',
]

@pulumi.output_type
class GetDeploymentResult:
    def __init__(__self__, deployment_id=None, description=None):
        if deployment_id and not isinstance(deployment_id, str):
            raise TypeError("Expected argument 'deployment_id' to be a str")
        pulumi.set(__self__, "deployment_id", deployment_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="deploymentId")
    def deployment_id(self) -> Optional[str]:
        """
        Primary Id for this resource
        """
        return pulumi.get(self, "deployment_id")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        A description of the purpose of the API Gateway deployment.
        """
        return pulumi.get(self, "description")


class AwaitableGetDeploymentResult(GetDeploymentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeploymentResult(
            deployment_id=self.deployment_id,
            description=self.description)


def get_deployment(deployment_id: Optional[str] = None,
                   rest_api_id: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeploymentResult:
    """
    Resource Type definition for AWS::ApiGateway::Deployment


    :param str deployment_id: Primary Id for this resource
    :param str rest_api_id: The ID of the RestApi resource to deploy. 
    """
    __args__ = dict()
    __args__['deploymentId'] = deployment_id
    __args__['restApiId'] = rest_api_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:apigateway:getDeployment', __args__, opts=opts, typ=GetDeploymentResult).value

    return AwaitableGetDeploymentResult(
        deployment_id=__ret__.deployment_id,
        description=__ret__.description)


@_utilities.lift_output_func(get_deployment)
def get_deployment_output(deployment_id: Optional[pulumi.Input[str]] = None,
                          rest_api_id: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDeploymentResult]:
    """
    Resource Type definition for AWS::ApiGateway::Deployment


    :param str deployment_id: Primary Id for this resource
    :param str rest_api_id: The ID of the RestApi resource to deploy. 
    """
    ...
