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
    'GetFlowTemplateResult',
    'AwaitableGetFlowTemplateResult',
    'get_flow_template',
    'get_flow_template_output',
]

@pulumi.output_type
class GetFlowTemplateResult:
    def __init__(__self__, compatible_namespace_version=None, definition=None, id=None):
        if compatible_namespace_version and not isinstance(compatible_namespace_version, float):
            raise TypeError("Expected argument 'compatible_namespace_version' to be a float")
        pulumi.set(__self__, "compatible_namespace_version", compatible_namespace_version)
        if definition and not isinstance(definition, dict):
            raise TypeError("Expected argument 'definition' to be a dict")
        pulumi.set(__self__, "definition", definition)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="compatibleNamespaceVersion")
    def compatible_namespace_version(self) -> Optional[float]:
        return pulumi.get(self, "compatible_namespace_version")

    @property
    @pulumi.getter
    def definition(self) -> Optional['outputs.FlowTemplateDefinitionDocument']:
        return pulumi.get(self, "definition")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")


class AwaitableGetFlowTemplateResult(GetFlowTemplateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFlowTemplateResult(
            compatible_namespace_version=self.compatible_namespace_version,
            definition=self.definition,
            id=self.id)


def get_flow_template(id: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFlowTemplateResult:
    """
    Resource Type definition for AWS::IoTThingsGraph::FlowTemplate
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:iotthingsgraph:getFlowTemplate', __args__, opts=opts, typ=GetFlowTemplateResult).value

    return AwaitableGetFlowTemplateResult(
        compatible_namespace_version=__ret__.compatible_namespace_version,
        definition=__ret__.definition,
        id=__ret__.id)


@_utilities.lift_output_func(get_flow_template)
def get_flow_template_output(id: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFlowTemplateResult]:
    """
    Resource Type definition for AWS::IoTThingsGraph::FlowTemplate
    """
    ...