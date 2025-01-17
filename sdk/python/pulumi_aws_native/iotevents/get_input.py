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
    'GetInputResult',
    'AwaitableGetInputResult',
    'get_input',
    'get_input_output',
]

@pulumi.output_type
class GetInputResult:
    def __init__(__self__, input_definition=None, input_description=None, tags=None):
        if input_definition and not isinstance(input_definition, dict):
            raise TypeError("Expected argument 'input_definition' to be a dict")
        pulumi.set(__self__, "input_definition", input_definition)
        if input_description and not isinstance(input_description, str):
            raise TypeError("Expected argument 'input_description' to be a str")
        pulumi.set(__self__, "input_description", input_description)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="inputDefinition")
    def input_definition(self) -> Optional['outputs.InputDefinition']:
        return pulumi.get(self, "input_definition")

    @property
    @pulumi.getter(name="inputDescription")
    def input_description(self) -> Optional[str]:
        """
        A brief description of the input.
        """
        return pulumi.get(self, "input_description")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.InputTag']]:
        """
        An array of key-value pairs to apply to this resource.

        For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).
        """
        return pulumi.get(self, "tags")


class AwaitableGetInputResult(GetInputResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInputResult(
            input_definition=self.input_definition,
            input_description=self.input_description,
            tags=self.tags)


def get_input(input_name: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInputResult:
    """
    The AWS::IoTEvents::Input resource creates an input. To monitor your devices and processes, they must have a way to get telemetry data into AWS IoT Events. This is done by sending messages as *inputs* to AWS IoT Events. For more information, see [How to Use AWS IoT Events](https://docs.aws.amazon.com/iotevents/latest/developerguide/how-to-use-iotevents.html) in the *AWS IoT Events Developer Guide*.


    :param str input_name: The name of the input.
    """
    __args__ = dict()
    __args__['inputName'] = input_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:iotevents:getInput', __args__, opts=opts, typ=GetInputResult).value

    return AwaitableGetInputResult(
        input_definition=__ret__.input_definition,
        input_description=__ret__.input_description,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_input)
def get_input_output(input_name: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInputResult]:
    """
    The AWS::IoTEvents::Input resource creates an input. To monitor your devices and processes, they must have a way to get telemetry data into AWS IoT Events. This is done by sending messages as *inputs* to AWS IoT Events. For more information, see [How to Use AWS IoT Events](https://docs.aws.amazon.com/iotevents/latest/developerguide/how-to-use-iotevents.html) in the *AWS IoT Events Developer Guide*.


    :param str input_name: The name of the input.
    """
    ...
