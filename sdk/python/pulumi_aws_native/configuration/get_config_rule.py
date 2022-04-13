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
    'GetConfigRuleResult',
    'AwaitableGetConfigRuleResult',
    'get_config_rule',
    'get_config_rule_output',
]

@pulumi.output_type
class GetConfigRuleResult:
    def __init__(__self__, arn=None, compliance_type=None, config_rule_id=None, description=None, input_parameters=None, maximum_execution_frequency=None, scope=None, source=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if compliance_type and not isinstance(compliance_type, str):
            raise TypeError("Expected argument 'compliance_type' to be a str")
        pulumi.set(__self__, "compliance_type", compliance_type)
        if config_rule_id and not isinstance(config_rule_id, str):
            raise TypeError("Expected argument 'config_rule_id' to be a str")
        pulumi.set(__self__, "config_rule_id", config_rule_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if input_parameters and not isinstance(input_parameters, dict):
            raise TypeError("Expected argument 'input_parameters' to be a dict")
        pulumi.set(__self__, "input_parameters", input_parameters)
        if maximum_execution_frequency and not isinstance(maximum_execution_frequency, str):
            raise TypeError("Expected argument 'maximum_execution_frequency' to be a str")
        pulumi.set(__self__, "maximum_execution_frequency", maximum_execution_frequency)
        if scope and not isinstance(scope, dict):
            raise TypeError("Expected argument 'scope' to be a dict")
        pulumi.set(__self__, "scope", scope)
        if source and not isinstance(source, dict):
            raise TypeError("Expected argument 'source' to be a dict")
        pulumi.set(__self__, "source", source)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="complianceType")
    def compliance_type(self) -> Optional[str]:
        return pulumi.get(self, "compliance_type")

    @property
    @pulumi.getter(name="configRuleId")
    def config_rule_id(self) -> Optional[str]:
        return pulumi.get(self, "config_rule_id")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="inputParameters")
    def input_parameters(self) -> Optional[Any]:
        return pulumi.get(self, "input_parameters")

    @property
    @pulumi.getter(name="maximumExecutionFrequency")
    def maximum_execution_frequency(self) -> Optional[str]:
        return pulumi.get(self, "maximum_execution_frequency")

    @property
    @pulumi.getter
    def scope(self) -> Optional['outputs.ConfigRuleScope']:
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def source(self) -> Optional['outputs.ConfigRuleSource']:
        return pulumi.get(self, "source")


class AwaitableGetConfigRuleResult(GetConfigRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConfigRuleResult(
            arn=self.arn,
            compliance_type=self.compliance_type,
            config_rule_id=self.config_rule_id,
            description=self.description,
            input_parameters=self.input_parameters,
            maximum_execution_frequency=self.maximum_execution_frequency,
            scope=self.scope,
            source=self.source)


def get_config_rule(config_rule_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConfigRuleResult:
    """
    Resource Type definition for AWS::Config::ConfigRule
    """
    __args__ = dict()
    __args__['configRuleId'] = config_rule_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:configuration:getConfigRule', __args__, opts=opts, typ=GetConfigRuleResult).value

    return AwaitableGetConfigRuleResult(
        arn=__ret__.arn,
        compliance_type=__ret__.compliance_type,
        config_rule_id=__ret__.config_rule_id,
        description=__ret__.description,
        input_parameters=__ret__.input_parameters,
        maximum_execution_frequency=__ret__.maximum_execution_frequency,
        scope=__ret__.scope,
        source=__ret__.source)


@_utilities.lift_output_func(get_config_rule)
def get_config_rule_output(config_rule_id: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetConfigRuleResult]:
    """
    Resource Type definition for AWS::Config::ConfigRule
    """
    ...
