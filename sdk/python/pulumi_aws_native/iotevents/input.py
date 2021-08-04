# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._inputs import *

__all__ = ['InputArgs', 'Input']

@pulumi.input_type
class InputArgs:
    def __init__(__self__, *,
                 input_definition: Optional[pulumi.Input['InputInputDefinitionArgs']] = None,
                 input_description: Optional[pulumi.Input[str]] = None,
                 input_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a Input resource.
        :param pulumi.Input['InputInputDefinitionArgs'] input_definition: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputdefinition
        :param pulumi.Input[str] input_description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputdescription
        :param pulumi.Input[str] input_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputname
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-tags
        """
        if input_definition is not None:
            pulumi.set(__self__, "input_definition", input_definition)
        if input_description is not None:
            pulumi.set(__self__, "input_description", input_description)
        if input_name is not None:
            pulumi.set(__self__, "input_name", input_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="inputDefinition")
    def input_definition(self) -> Optional[pulumi.Input['InputInputDefinitionArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputdefinition
        """
        return pulumi.get(self, "input_definition")

    @input_definition.setter
    def input_definition(self, value: Optional[pulumi.Input['InputInputDefinitionArgs']]):
        pulumi.set(self, "input_definition", value)

    @property
    @pulumi.getter(name="inputDescription")
    def input_description(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputdescription
        """
        return pulumi.get(self, "input_description")

    @input_description.setter
    def input_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "input_description", value)

    @property
    @pulumi.getter(name="inputName")
    def input_name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputname
        """
        return pulumi.get(self, "input_name")

    @input_name.setter
    def input_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "input_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


class Input(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 input_definition: Optional[pulumi.Input[pulumi.InputType['InputInputDefinitionArgs']]] = None,
                 input_description: Optional[pulumi.Input[str]] = None,
                 input_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['InputInputDefinitionArgs']] input_definition: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputdefinition
        :param pulumi.Input[str] input_description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputdescription
        :param pulumi.Input[str] input_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputname
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-tags
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[InputArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html

        :param str resource_name: The name of the resource.
        :param InputArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InputArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 input_definition: Optional[pulumi.Input[pulumi.InputType['InputInputDefinitionArgs']]] = None,
                 input_description: Optional[pulumi.Input[str]] = None,
                 input_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InputArgs.__new__(InputArgs)

            __props__.__dict__["input_definition"] = input_definition
            __props__.__dict__["input_description"] = input_description
            __props__.__dict__["input_name"] = input_name
            __props__.__dict__["tags"] = tags
        super(Input, __self__).__init__(
            'aws-native:IoTEvents:Input',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Input':
        """
        Get an existing Input resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = InputArgs.__new__(InputArgs)

        __props__.__dict__["input_definition"] = None
        __props__.__dict__["input_description"] = None
        __props__.__dict__["input_name"] = None
        __props__.__dict__["tags"] = None
        return Input(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="inputDefinition")
    def input_definition(self) -> pulumi.Output[Optional['outputs.InputInputDefinition']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputdefinition
        """
        return pulumi.get(self, "input_definition")

    @property
    @pulumi.getter(name="inputDescription")
    def input_description(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputdescription
        """
        return pulumi.get(self, "input_description")

    @property
    @pulumi.getter(name="inputName")
    def input_name(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputname
        """
        return pulumi.get(self, "input_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-tags
        """
        return pulumi.get(self, "tags")
