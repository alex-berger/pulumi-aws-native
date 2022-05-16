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
    'GetAuthorizerResult',
    'AwaitableGetAuthorizerResult',
    'get_authorizer',
    'get_authorizer_output',
]

@pulumi.output_type
class GetAuthorizerResult:
    def __init__(__self__, authorizer_credentials_arn=None, authorizer_payload_format_version=None, authorizer_result_ttl_in_seconds=None, authorizer_type=None, authorizer_uri=None, enable_simple_responses=None, id=None, identity_source=None, identity_validation_expression=None, jwt_configuration=None, name=None):
        if authorizer_credentials_arn and not isinstance(authorizer_credentials_arn, str):
            raise TypeError("Expected argument 'authorizer_credentials_arn' to be a str")
        pulumi.set(__self__, "authorizer_credentials_arn", authorizer_credentials_arn)
        if authorizer_payload_format_version and not isinstance(authorizer_payload_format_version, str):
            raise TypeError("Expected argument 'authorizer_payload_format_version' to be a str")
        pulumi.set(__self__, "authorizer_payload_format_version", authorizer_payload_format_version)
        if authorizer_result_ttl_in_seconds and not isinstance(authorizer_result_ttl_in_seconds, int):
            raise TypeError("Expected argument 'authorizer_result_ttl_in_seconds' to be a int")
        pulumi.set(__self__, "authorizer_result_ttl_in_seconds", authorizer_result_ttl_in_seconds)
        if authorizer_type and not isinstance(authorizer_type, str):
            raise TypeError("Expected argument 'authorizer_type' to be a str")
        pulumi.set(__self__, "authorizer_type", authorizer_type)
        if authorizer_uri and not isinstance(authorizer_uri, str):
            raise TypeError("Expected argument 'authorizer_uri' to be a str")
        pulumi.set(__self__, "authorizer_uri", authorizer_uri)
        if enable_simple_responses and not isinstance(enable_simple_responses, bool):
            raise TypeError("Expected argument 'enable_simple_responses' to be a bool")
        pulumi.set(__self__, "enable_simple_responses", enable_simple_responses)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity_source and not isinstance(identity_source, list):
            raise TypeError("Expected argument 'identity_source' to be a list")
        pulumi.set(__self__, "identity_source", identity_source)
        if identity_validation_expression and not isinstance(identity_validation_expression, str):
            raise TypeError("Expected argument 'identity_validation_expression' to be a str")
        pulumi.set(__self__, "identity_validation_expression", identity_validation_expression)
        if jwt_configuration and not isinstance(jwt_configuration, dict):
            raise TypeError("Expected argument 'jwt_configuration' to be a dict")
        pulumi.set(__self__, "jwt_configuration", jwt_configuration)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="authorizerCredentialsArn")
    def authorizer_credentials_arn(self) -> Optional[str]:
        return pulumi.get(self, "authorizer_credentials_arn")

    @property
    @pulumi.getter(name="authorizerPayloadFormatVersion")
    def authorizer_payload_format_version(self) -> Optional[str]:
        return pulumi.get(self, "authorizer_payload_format_version")

    @property
    @pulumi.getter(name="authorizerResultTtlInSeconds")
    def authorizer_result_ttl_in_seconds(self) -> Optional[int]:
        return pulumi.get(self, "authorizer_result_ttl_in_seconds")

    @property
    @pulumi.getter(name="authorizerType")
    def authorizer_type(self) -> Optional[str]:
        return pulumi.get(self, "authorizer_type")

    @property
    @pulumi.getter(name="authorizerUri")
    def authorizer_uri(self) -> Optional[str]:
        return pulumi.get(self, "authorizer_uri")

    @property
    @pulumi.getter(name="enableSimpleResponses")
    def enable_simple_responses(self) -> Optional[bool]:
        return pulumi.get(self, "enable_simple_responses")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="identitySource")
    def identity_source(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "identity_source")

    @property
    @pulumi.getter(name="identityValidationExpression")
    def identity_validation_expression(self) -> Optional[str]:
        return pulumi.get(self, "identity_validation_expression")

    @property
    @pulumi.getter(name="jwtConfiguration")
    def jwt_configuration(self) -> Optional['outputs.AuthorizerJWTConfiguration']:
        return pulumi.get(self, "jwt_configuration")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")


class AwaitableGetAuthorizerResult(GetAuthorizerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAuthorizerResult(
            authorizer_credentials_arn=self.authorizer_credentials_arn,
            authorizer_payload_format_version=self.authorizer_payload_format_version,
            authorizer_result_ttl_in_seconds=self.authorizer_result_ttl_in_seconds,
            authorizer_type=self.authorizer_type,
            authorizer_uri=self.authorizer_uri,
            enable_simple_responses=self.enable_simple_responses,
            id=self.id,
            identity_source=self.identity_source,
            identity_validation_expression=self.identity_validation_expression,
            jwt_configuration=self.jwt_configuration,
            name=self.name)


def get_authorizer(id: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAuthorizerResult:
    """
    Resource Type definition for AWS::ApiGatewayV2::Authorizer
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:apigatewayv2:getAuthorizer', __args__, opts=opts, typ=GetAuthorizerResult).value

    return AwaitableGetAuthorizerResult(
        authorizer_credentials_arn=__ret__.authorizer_credentials_arn,
        authorizer_payload_format_version=__ret__.authorizer_payload_format_version,
        authorizer_result_ttl_in_seconds=__ret__.authorizer_result_ttl_in_seconds,
        authorizer_type=__ret__.authorizer_type,
        authorizer_uri=__ret__.authorizer_uri,
        enable_simple_responses=__ret__.enable_simple_responses,
        id=__ret__.id,
        identity_source=__ret__.identity_source,
        identity_validation_expression=__ret__.identity_validation_expression,
        jwt_configuration=__ret__.jwt_configuration,
        name=__ret__.name)


@_utilities.lift_output_func(get_authorizer)
def get_authorizer_output(id: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAuthorizerResult]:
    """
    Resource Type definition for AWS::ApiGatewayV2::Authorizer
    """
    ...