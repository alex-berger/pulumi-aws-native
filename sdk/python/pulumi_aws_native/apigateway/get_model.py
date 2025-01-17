# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetModelResult',
    'AwaitableGetModelResult',
    'get_model',
    'get_model_output',
]

@pulumi.output_type
class GetModelResult:
    def __init__(__self__, description=None, schema=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if schema and not isinstance(schema, dict):
            raise TypeError("Expected argument 'schema' to be a dict")
        pulumi.set(__self__, "schema", schema)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        A description that identifies this model.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def schema(self) -> Optional[Any]:
        """
        The schema to use to transform data to one or more output formats. Specify null ({}) if you don't want to specify a schema.
        """
        return pulumi.get(self, "schema")


class AwaitableGetModelResult(GetModelResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetModelResult(
            description=self.description,
            schema=self.schema)


def get_model(name: Optional[str] = None,
              rest_api_id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetModelResult:
    """
    Resource Type definition for AWS::ApiGateway::Model


    :param str name: A name for the model. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the model name.
    :param str rest_api_id: The ID of a REST API with which to associate this model.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['restApiId'] = rest_api_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:apigateway:getModel', __args__, opts=opts, typ=GetModelResult).value

    return AwaitableGetModelResult(
        description=__ret__.description,
        schema=__ret__.schema)


@_utilities.lift_output_func(get_model)
def get_model_output(name: Optional[pulumi.Input[str]] = None,
                     rest_api_id: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetModelResult]:
    """
    Resource Type definition for AWS::ApiGateway::Model


    :param str name: A name for the model. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the model name.
    :param str rest_api_id: The ID of a REST API with which to associate this model.
    """
    ...
