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
    'GetSuiteDefinitionResult',
    'AwaitableGetSuiteDefinitionResult',
    'get_suite_definition',
    'get_suite_definition_output',
]

@pulumi.output_type
class GetSuiteDefinitionResult:
    def __init__(__self__, suite_definition_arn=None, suite_definition_configuration=None, suite_definition_id=None, suite_definition_version=None, tags=None):
        if suite_definition_arn and not isinstance(suite_definition_arn, str):
            raise TypeError("Expected argument 'suite_definition_arn' to be a str")
        pulumi.set(__self__, "suite_definition_arn", suite_definition_arn)
        if suite_definition_configuration and not isinstance(suite_definition_configuration, dict):
            raise TypeError("Expected argument 'suite_definition_configuration' to be a dict")
        pulumi.set(__self__, "suite_definition_configuration", suite_definition_configuration)
        if suite_definition_id and not isinstance(suite_definition_id, str):
            raise TypeError("Expected argument 'suite_definition_id' to be a str")
        pulumi.set(__self__, "suite_definition_id", suite_definition_id)
        if suite_definition_version and not isinstance(suite_definition_version, str):
            raise TypeError("Expected argument 'suite_definition_version' to be a str")
        pulumi.set(__self__, "suite_definition_version", suite_definition_version)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="suiteDefinitionArn")
    def suite_definition_arn(self) -> Optional[str]:
        """
        The Amazon Resource name for the suite definition.
        """
        return pulumi.get(self, "suite_definition_arn")

    @property
    @pulumi.getter(name="suiteDefinitionConfiguration")
    def suite_definition_configuration(self) -> Optional['outputs.SuiteDefinitionConfigurationProperties']:
        return pulumi.get(self, "suite_definition_configuration")

    @property
    @pulumi.getter(name="suiteDefinitionId")
    def suite_definition_id(self) -> Optional[str]:
        """
        The unique identifier for the suite definition.
        """
        return pulumi.get(self, "suite_definition_id")

    @property
    @pulumi.getter(name="suiteDefinitionVersion")
    def suite_definition_version(self) -> Optional[str]:
        """
        The suite definition version of a test suite.
        """
        return pulumi.get(self, "suite_definition_version")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.SuiteDefinitionTag']]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetSuiteDefinitionResult(GetSuiteDefinitionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSuiteDefinitionResult(
            suite_definition_arn=self.suite_definition_arn,
            suite_definition_configuration=self.suite_definition_configuration,
            suite_definition_id=self.suite_definition_id,
            suite_definition_version=self.suite_definition_version,
            tags=self.tags)


def get_suite_definition(suite_definition_id: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSuiteDefinitionResult:
    """
    An example resource schema demonstrating some basic constructs and validation rules.


    :param str suite_definition_id: The unique identifier for the suite definition.
    """
    __args__ = dict()
    __args__['suiteDefinitionId'] = suite_definition_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:iotcoredeviceadvisor:getSuiteDefinition', __args__, opts=opts, typ=GetSuiteDefinitionResult).value

    return AwaitableGetSuiteDefinitionResult(
        suite_definition_arn=__ret__.suite_definition_arn,
        suite_definition_configuration=__ret__.suite_definition_configuration,
        suite_definition_id=__ret__.suite_definition_id,
        suite_definition_version=__ret__.suite_definition_version,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_suite_definition)
def get_suite_definition_output(suite_definition_id: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSuiteDefinitionResult]:
    """
    An example resource schema demonstrating some basic constructs and validation rules.


    :param str suite_definition_id: The unique identifier for the suite definition.
    """
    ...
