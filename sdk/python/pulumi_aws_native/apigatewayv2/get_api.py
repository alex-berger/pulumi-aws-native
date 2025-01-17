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
    'GetApiResult',
    'AwaitableGetApiResult',
    'get_api',
    'get_api_output',
]

@pulumi.output_type
class GetApiResult:
    def __init__(__self__, api_endpoint=None, api_key_selection_expression=None, base_path=None, body=None, body_s3_location=None, cors_configuration=None, credentials_arn=None, description=None, disable_execute_api_endpoint=None, disable_schema_validation=None, fail_on_warnings=None, id=None, name=None, route_key=None, route_selection_expression=None, tags=None, target=None, version=None):
        if api_endpoint and not isinstance(api_endpoint, str):
            raise TypeError("Expected argument 'api_endpoint' to be a str")
        pulumi.set(__self__, "api_endpoint", api_endpoint)
        if api_key_selection_expression and not isinstance(api_key_selection_expression, str):
            raise TypeError("Expected argument 'api_key_selection_expression' to be a str")
        pulumi.set(__self__, "api_key_selection_expression", api_key_selection_expression)
        if base_path and not isinstance(base_path, str):
            raise TypeError("Expected argument 'base_path' to be a str")
        pulumi.set(__self__, "base_path", base_path)
        if body and not isinstance(body, dict):
            raise TypeError("Expected argument 'body' to be a dict")
        pulumi.set(__self__, "body", body)
        if body_s3_location and not isinstance(body_s3_location, dict):
            raise TypeError("Expected argument 'body_s3_location' to be a dict")
        pulumi.set(__self__, "body_s3_location", body_s3_location)
        if cors_configuration and not isinstance(cors_configuration, dict):
            raise TypeError("Expected argument 'cors_configuration' to be a dict")
        pulumi.set(__self__, "cors_configuration", cors_configuration)
        if credentials_arn and not isinstance(credentials_arn, str):
            raise TypeError("Expected argument 'credentials_arn' to be a str")
        pulumi.set(__self__, "credentials_arn", credentials_arn)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disable_execute_api_endpoint and not isinstance(disable_execute_api_endpoint, bool):
            raise TypeError("Expected argument 'disable_execute_api_endpoint' to be a bool")
        pulumi.set(__self__, "disable_execute_api_endpoint", disable_execute_api_endpoint)
        if disable_schema_validation and not isinstance(disable_schema_validation, bool):
            raise TypeError("Expected argument 'disable_schema_validation' to be a bool")
        pulumi.set(__self__, "disable_schema_validation", disable_schema_validation)
        if fail_on_warnings and not isinstance(fail_on_warnings, bool):
            raise TypeError("Expected argument 'fail_on_warnings' to be a bool")
        pulumi.set(__self__, "fail_on_warnings", fail_on_warnings)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if route_key and not isinstance(route_key, str):
            raise TypeError("Expected argument 'route_key' to be a str")
        pulumi.set(__self__, "route_key", route_key)
        if route_selection_expression and not isinstance(route_selection_expression, str):
            raise TypeError("Expected argument 'route_selection_expression' to be a str")
        pulumi.set(__self__, "route_selection_expression", route_selection_expression)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if target and not isinstance(target, str):
            raise TypeError("Expected argument 'target' to be a str")
        pulumi.set(__self__, "target", target)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="apiEndpoint")
    def api_endpoint(self) -> Optional[str]:
        return pulumi.get(self, "api_endpoint")

    @property
    @pulumi.getter(name="apiKeySelectionExpression")
    def api_key_selection_expression(self) -> Optional[str]:
        return pulumi.get(self, "api_key_selection_expression")

    @property
    @pulumi.getter(name="basePath")
    def base_path(self) -> Optional[str]:
        return pulumi.get(self, "base_path")

    @property
    @pulumi.getter
    def body(self) -> Optional[Any]:
        return pulumi.get(self, "body")

    @property
    @pulumi.getter(name="bodyS3Location")
    def body_s3_location(self) -> Optional['outputs.ApiBodyS3Location']:
        return pulumi.get(self, "body_s3_location")

    @property
    @pulumi.getter(name="corsConfiguration")
    def cors_configuration(self) -> Optional['outputs.ApiCors']:
        return pulumi.get(self, "cors_configuration")

    @property
    @pulumi.getter(name="credentialsArn")
    def credentials_arn(self) -> Optional[str]:
        return pulumi.get(self, "credentials_arn")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="disableExecuteApiEndpoint")
    def disable_execute_api_endpoint(self) -> Optional[bool]:
        return pulumi.get(self, "disable_execute_api_endpoint")

    @property
    @pulumi.getter(name="disableSchemaValidation")
    def disable_schema_validation(self) -> Optional[bool]:
        return pulumi.get(self, "disable_schema_validation")

    @property
    @pulumi.getter(name="failOnWarnings")
    def fail_on_warnings(self) -> Optional[bool]:
        return pulumi.get(self, "fail_on_warnings")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="routeKey")
    def route_key(self) -> Optional[str]:
        return pulumi.get(self, "route_key")

    @property
    @pulumi.getter(name="routeSelectionExpression")
    def route_selection_expression(self) -> Optional[str]:
        return pulumi.get(self, "route_selection_expression")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Any]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def target(self) -> Optional[str]:
        return pulumi.get(self, "target")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        return pulumi.get(self, "version")


class AwaitableGetApiResult(GetApiResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApiResult(
            api_endpoint=self.api_endpoint,
            api_key_selection_expression=self.api_key_selection_expression,
            base_path=self.base_path,
            body=self.body,
            body_s3_location=self.body_s3_location,
            cors_configuration=self.cors_configuration,
            credentials_arn=self.credentials_arn,
            description=self.description,
            disable_execute_api_endpoint=self.disable_execute_api_endpoint,
            disable_schema_validation=self.disable_schema_validation,
            fail_on_warnings=self.fail_on_warnings,
            id=self.id,
            name=self.name,
            route_key=self.route_key,
            route_selection_expression=self.route_selection_expression,
            tags=self.tags,
            target=self.target,
            version=self.version)


def get_api(id: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApiResult:
    """
    Resource Type definition for AWS::ApiGatewayV2::Api
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:apigatewayv2:getApi', __args__, opts=opts, typ=GetApiResult).value

    return AwaitableGetApiResult(
        api_endpoint=__ret__.api_endpoint,
        api_key_selection_expression=__ret__.api_key_selection_expression,
        base_path=__ret__.base_path,
        body=__ret__.body,
        body_s3_location=__ret__.body_s3_location,
        cors_configuration=__ret__.cors_configuration,
        credentials_arn=__ret__.credentials_arn,
        description=__ret__.description,
        disable_execute_api_endpoint=__ret__.disable_execute_api_endpoint,
        disable_schema_validation=__ret__.disable_schema_validation,
        fail_on_warnings=__ret__.fail_on_warnings,
        id=__ret__.id,
        name=__ret__.name,
        route_key=__ret__.route_key,
        route_selection_expression=__ret__.route_selection_expression,
        tags=__ret__.tags,
        target=__ret__.target,
        version=__ret__.version)


@_utilities.lift_output_func(get_api)
def get_api_output(id: Optional[pulumi.Input[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetApiResult]:
    """
    Resource Type definition for AWS::ApiGatewayV2::Api
    """
    ...
