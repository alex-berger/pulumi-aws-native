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
    'GetRestApiResult',
    'AwaitableGetRestApiResult',
    'get_rest_api',
    'get_rest_api_output',
]

@pulumi.output_type
class GetRestApiResult:
    def __init__(__self__, api_key_source_type=None, binary_media_types=None, body=None, body_s3_location=None, clone_from=None, description=None, disable_execute_api_endpoint=None, endpoint_configuration=None, fail_on_warnings=None, id=None, minimum_compression_size=None, mode=None, name=None, parameters=None, policy=None, root_resource_id=None, tags=None):
        if api_key_source_type and not isinstance(api_key_source_type, str):
            raise TypeError("Expected argument 'api_key_source_type' to be a str")
        pulumi.set(__self__, "api_key_source_type", api_key_source_type)
        if binary_media_types and not isinstance(binary_media_types, list):
            raise TypeError("Expected argument 'binary_media_types' to be a list")
        pulumi.set(__self__, "binary_media_types", binary_media_types)
        if body and not isinstance(body, dict):
            raise TypeError("Expected argument 'body' to be a dict")
        pulumi.set(__self__, "body", body)
        if body_s3_location and not isinstance(body_s3_location, dict):
            raise TypeError("Expected argument 'body_s3_location' to be a dict")
        pulumi.set(__self__, "body_s3_location", body_s3_location)
        if clone_from and not isinstance(clone_from, str):
            raise TypeError("Expected argument 'clone_from' to be a str")
        pulumi.set(__self__, "clone_from", clone_from)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disable_execute_api_endpoint and not isinstance(disable_execute_api_endpoint, bool):
            raise TypeError("Expected argument 'disable_execute_api_endpoint' to be a bool")
        pulumi.set(__self__, "disable_execute_api_endpoint", disable_execute_api_endpoint)
        if endpoint_configuration and not isinstance(endpoint_configuration, dict):
            raise TypeError("Expected argument 'endpoint_configuration' to be a dict")
        pulumi.set(__self__, "endpoint_configuration", endpoint_configuration)
        if fail_on_warnings and not isinstance(fail_on_warnings, bool):
            raise TypeError("Expected argument 'fail_on_warnings' to be a bool")
        pulumi.set(__self__, "fail_on_warnings", fail_on_warnings)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if minimum_compression_size and not isinstance(minimum_compression_size, int):
            raise TypeError("Expected argument 'minimum_compression_size' to be a int")
        pulumi.set(__self__, "minimum_compression_size", minimum_compression_size)
        if mode and not isinstance(mode, str):
            raise TypeError("Expected argument 'mode' to be a str")
        pulumi.set(__self__, "mode", mode)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if parameters and not isinstance(parameters, dict):
            raise TypeError("Expected argument 'parameters' to be a dict")
        pulumi.set(__self__, "parameters", parameters)
        if policy and not isinstance(policy, dict):
            raise TypeError("Expected argument 'policy' to be a dict")
        pulumi.set(__self__, "policy", policy)
        if root_resource_id and not isinstance(root_resource_id, str):
            raise TypeError("Expected argument 'root_resource_id' to be a str")
        pulumi.set(__self__, "root_resource_id", root_resource_id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="apiKeySourceType")
    def api_key_source_type(self) -> Optional[str]:
        return pulumi.get(self, "api_key_source_type")

    @property
    @pulumi.getter(name="binaryMediaTypes")
    def binary_media_types(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "binary_media_types")

    @property
    @pulumi.getter
    def body(self) -> Optional[Any]:
        return pulumi.get(self, "body")

    @property
    @pulumi.getter(name="bodyS3Location")
    def body_s3_location(self) -> Optional['outputs.RestApiS3Location']:
        return pulumi.get(self, "body_s3_location")

    @property
    @pulumi.getter(name="cloneFrom")
    def clone_from(self) -> Optional[str]:
        return pulumi.get(self, "clone_from")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="disableExecuteApiEndpoint")
    def disable_execute_api_endpoint(self) -> Optional[bool]:
        return pulumi.get(self, "disable_execute_api_endpoint")

    @property
    @pulumi.getter(name="endpointConfiguration")
    def endpoint_configuration(self) -> Optional['outputs.RestApiEndpointConfiguration']:
        return pulumi.get(self, "endpoint_configuration")

    @property
    @pulumi.getter(name="failOnWarnings")
    def fail_on_warnings(self) -> Optional[bool]:
        return pulumi.get(self, "fail_on_warnings")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="minimumCompressionSize")
    def minimum_compression_size(self) -> Optional[int]:
        return pulumi.get(self, "minimum_compression_size")

    @property
    @pulumi.getter
    def mode(self) -> Optional[str]:
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Any]:
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter
    def policy(self) -> Optional[Any]:
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter(name="rootResourceId")
    def root_resource_id(self) -> Optional[str]:
        return pulumi.get(self, "root_resource_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.RestApiTag']]:
        return pulumi.get(self, "tags")


class AwaitableGetRestApiResult(GetRestApiResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRestApiResult(
            api_key_source_type=self.api_key_source_type,
            binary_media_types=self.binary_media_types,
            body=self.body,
            body_s3_location=self.body_s3_location,
            clone_from=self.clone_from,
            description=self.description,
            disable_execute_api_endpoint=self.disable_execute_api_endpoint,
            endpoint_configuration=self.endpoint_configuration,
            fail_on_warnings=self.fail_on_warnings,
            id=self.id,
            minimum_compression_size=self.minimum_compression_size,
            mode=self.mode,
            name=self.name,
            parameters=self.parameters,
            policy=self.policy,
            root_resource_id=self.root_resource_id,
            tags=self.tags)


def get_rest_api(id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRestApiResult:
    """
    Resource Type definition for AWS::ApiGateway::RestApi
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:apigateway:getRestApi', __args__, opts=opts, typ=GetRestApiResult).value

    return AwaitableGetRestApiResult(
        api_key_source_type=__ret__.api_key_source_type,
        binary_media_types=__ret__.binary_media_types,
        body=__ret__.body,
        body_s3_location=__ret__.body_s3_location,
        clone_from=__ret__.clone_from,
        description=__ret__.description,
        disable_execute_api_endpoint=__ret__.disable_execute_api_endpoint,
        endpoint_configuration=__ret__.endpoint_configuration,
        fail_on_warnings=__ret__.fail_on_warnings,
        id=__ret__.id,
        minimum_compression_size=__ret__.minimum_compression_size,
        mode=__ret__.mode,
        name=__ret__.name,
        parameters=__ret__.parameters,
        policy=__ret__.policy,
        root_resource_id=__ret__.root_resource_id,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_rest_api)
def get_rest_api_output(id: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRestApiResult]:
    """
    Resource Type definition for AWS::ApiGateway::RestApi
    """
    ...
