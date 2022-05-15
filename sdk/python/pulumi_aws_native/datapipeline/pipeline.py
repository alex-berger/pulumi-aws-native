# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['PipelineArgs', 'Pipeline']

@pulumi.input_type
class PipelineArgs:
    def __init__(__self__, *,
                 parameter_objects: pulumi.Input[Sequence[pulumi.Input['PipelineParameterObjectArgs']]],
                 activate: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameter_values: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineParameterValueArgs']]]] = None,
                 pipeline_objects: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineObjectArgs']]]] = None,
                 pipeline_tags: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineTagArgs']]]] = None):
        """
        The set of arguments for constructing a Pipeline resource.
        """
        pulumi.set(__self__, "parameter_objects", parameter_objects)
        if activate is not None:
            pulumi.set(__self__, "activate", activate)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parameter_values is not None:
            pulumi.set(__self__, "parameter_values", parameter_values)
        if pipeline_objects is not None:
            pulumi.set(__self__, "pipeline_objects", pipeline_objects)
        if pipeline_tags is not None:
            pulumi.set(__self__, "pipeline_tags", pipeline_tags)

    @property
    @pulumi.getter(name="parameterObjects")
    def parameter_objects(self) -> pulumi.Input[Sequence[pulumi.Input['PipelineParameterObjectArgs']]]:
        return pulumi.get(self, "parameter_objects")

    @parameter_objects.setter
    def parameter_objects(self, value: pulumi.Input[Sequence[pulumi.Input['PipelineParameterObjectArgs']]]):
        pulumi.set(self, "parameter_objects", value)

    @property
    @pulumi.getter
    def activate(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "activate")

    @activate.setter
    def activate(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "activate", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="parameterValues")
    def parameter_values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PipelineParameterValueArgs']]]]:
        return pulumi.get(self, "parameter_values")

    @parameter_values.setter
    def parameter_values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineParameterValueArgs']]]]):
        pulumi.set(self, "parameter_values", value)

    @property
    @pulumi.getter(name="pipelineObjects")
    def pipeline_objects(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PipelineObjectArgs']]]]:
        return pulumi.get(self, "pipeline_objects")

    @pipeline_objects.setter
    def pipeline_objects(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineObjectArgs']]]]):
        pulumi.set(self, "pipeline_objects", value)

    @property
    @pulumi.getter(name="pipelineTags")
    def pipeline_tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PipelineTagArgs']]]]:
        return pulumi.get(self, "pipeline_tags")

    @pipeline_tags.setter
    def pipeline_tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineTagArgs']]]]):
        pulumi.set(self, "pipeline_tags", value)


warnings.warn("""Pipeline is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class Pipeline(pulumi.CustomResource):
    warnings.warn("""Pipeline is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 activate: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameter_objects: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineParameterObjectArgs']]]]] = None,
                 parameter_values: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineParameterValueArgs']]]]] = None,
                 pipeline_objects: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineObjectArgs']]]]] = None,
                 pipeline_tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineTagArgs']]]]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::DataPipeline::Pipeline

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PipelineArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::DataPipeline::Pipeline

        :param str resource_name: The name of the resource.
        :param PipelineArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PipelineArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 activate: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameter_objects: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineParameterObjectArgs']]]]] = None,
                 parameter_values: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineParameterValueArgs']]]]] = None,
                 pipeline_objects: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineObjectArgs']]]]] = None,
                 pipeline_tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineTagArgs']]]]] = None,
                 __props__=None):
        pulumi.log.warn("""Pipeline is deprecated: Pipeline is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PipelineArgs.__new__(PipelineArgs)

            __props__.__dict__["activate"] = activate
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if parameter_objects is None and not opts.urn:
                raise TypeError("Missing required property 'parameter_objects'")
            __props__.__dict__["parameter_objects"] = parameter_objects
            __props__.__dict__["parameter_values"] = parameter_values
            __props__.__dict__["pipeline_objects"] = pipeline_objects
            __props__.__dict__["pipeline_tags"] = pipeline_tags
        super(Pipeline, __self__).__init__(
            'aws-native:datapipeline:Pipeline',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Pipeline':
        """
        Get an existing Pipeline resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PipelineArgs.__new__(PipelineArgs)

        __props__.__dict__["activate"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["parameter_objects"] = None
        __props__.__dict__["parameter_values"] = None
        __props__.__dict__["pipeline_objects"] = None
        __props__.__dict__["pipeline_tags"] = None
        return Pipeline(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def activate(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "activate")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parameterObjects")
    def parameter_objects(self) -> pulumi.Output[Sequence['outputs.PipelineParameterObject']]:
        return pulumi.get(self, "parameter_objects")

    @property
    @pulumi.getter(name="parameterValues")
    def parameter_values(self) -> pulumi.Output[Optional[Sequence['outputs.PipelineParameterValue']]]:
        return pulumi.get(self, "parameter_values")

    @property
    @pulumi.getter(name="pipelineObjects")
    def pipeline_objects(self) -> pulumi.Output[Optional[Sequence['outputs.PipelineObject']]]:
        return pulumi.get(self, "pipeline_objects")

    @property
    @pulumi.getter(name="pipelineTags")
    def pipeline_tags(self) -> pulumi.Output[Optional[Sequence['outputs.PipelineTag']]]:
        return pulumi.get(self, "pipeline_tags")

