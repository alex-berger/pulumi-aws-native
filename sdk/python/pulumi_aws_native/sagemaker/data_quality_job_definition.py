# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['DataQualityJobDefinitionArgs', 'DataQualityJobDefinition']

@pulumi.input_type
class DataQualityJobDefinitionArgs:
    def __init__(__self__, *,
                 data_quality_app_specification: pulumi.Input['DataQualityJobDefinitionDataQualityAppSpecificationArgs'],
                 data_quality_job_input: pulumi.Input['DataQualityJobDefinitionDataQualityJobInputArgs'],
                 data_quality_job_output_config: pulumi.Input['DataQualityJobDefinitionMonitoringOutputConfigArgs'],
                 job_resources: pulumi.Input['DataQualityJobDefinitionMonitoringResourcesArgs'],
                 role_arn: pulumi.Input[str],
                 data_quality_baseline_config: Optional[pulumi.Input['DataQualityJobDefinitionDataQualityBaselineConfigArgs']] = None,
                 job_definition_name: Optional[pulumi.Input[str]] = None,
                 network_config: Optional[pulumi.Input['DataQualityJobDefinitionNetworkConfigArgs']] = None,
                 stopping_condition: Optional[pulumi.Input['DataQualityJobDefinitionStoppingConditionArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['DataQualityJobDefinitionTagArgs']]]] = None):
        """
        The set of arguments for constructing a DataQualityJobDefinition resource.
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
        :param pulumi.Input[Sequence[pulumi.Input['DataQualityJobDefinitionTagArgs']]] tags: An array of key-value pairs to apply to this resource.
        """
        pulumi.set(__self__, "data_quality_app_specification", data_quality_app_specification)
        pulumi.set(__self__, "data_quality_job_input", data_quality_job_input)
        pulumi.set(__self__, "data_quality_job_output_config", data_quality_job_output_config)
        pulumi.set(__self__, "job_resources", job_resources)
        pulumi.set(__self__, "role_arn", role_arn)
        if data_quality_baseline_config is not None:
            pulumi.set(__self__, "data_quality_baseline_config", data_quality_baseline_config)
        if job_definition_name is not None:
            pulumi.set(__self__, "job_definition_name", job_definition_name)
        if network_config is not None:
            pulumi.set(__self__, "network_config", network_config)
        if stopping_condition is not None:
            pulumi.set(__self__, "stopping_condition", stopping_condition)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="dataQualityAppSpecification")
    def data_quality_app_specification(self) -> pulumi.Input['DataQualityJobDefinitionDataQualityAppSpecificationArgs']:
        return pulumi.get(self, "data_quality_app_specification")

    @data_quality_app_specification.setter
    def data_quality_app_specification(self, value: pulumi.Input['DataQualityJobDefinitionDataQualityAppSpecificationArgs']):
        pulumi.set(self, "data_quality_app_specification", value)

    @property
    @pulumi.getter(name="dataQualityJobInput")
    def data_quality_job_input(self) -> pulumi.Input['DataQualityJobDefinitionDataQualityJobInputArgs']:
        return pulumi.get(self, "data_quality_job_input")

    @data_quality_job_input.setter
    def data_quality_job_input(self, value: pulumi.Input['DataQualityJobDefinitionDataQualityJobInputArgs']):
        pulumi.set(self, "data_quality_job_input", value)

    @property
    @pulumi.getter(name="dataQualityJobOutputConfig")
    def data_quality_job_output_config(self) -> pulumi.Input['DataQualityJobDefinitionMonitoringOutputConfigArgs']:
        return pulumi.get(self, "data_quality_job_output_config")

    @data_quality_job_output_config.setter
    def data_quality_job_output_config(self, value: pulumi.Input['DataQualityJobDefinitionMonitoringOutputConfigArgs']):
        pulumi.set(self, "data_quality_job_output_config", value)

    @property
    @pulumi.getter(name="jobResources")
    def job_resources(self) -> pulumi.Input['DataQualityJobDefinitionMonitoringResourcesArgs']:
        return pulumi.get(self, "job_resources")

    @job_resources.setter
    def job_resources(self, value: pulumi.Input['DataQualityJobDefinitionMonitoringResourcesArgs']):
        pulumi.set(self, "job_resources", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter(name="dataQualityBaselineConfig")
    def data_quality_baseline_config(self) -> Optional[pulumi.Input['DataQualityJobDefinitionDataQualityBaselineConfigArgs']]:
        return pulumi.get(self, "data_quality_baseline_config")

    @data_quality_baseline_config.setter
    def data_quality_baseline_config(self, value: Optional[pulumi.Input['DataQualityJobDefinitionDataQualityBaselineConfigArgs']]):
        pulumi.set(self, "data_quality_baseline_config", value)

    @property
    @pulumi.getter(name="jobDefinitionName")
    def job_definition_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "job_definition_name")

    @job_definition_name.setter
    def job_definition_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "job_definition_name", value)

    @property
    @pulumi.getter(name="networkConfig")
    def network_config(self) -> Optional[pulumi.Input['DataQualityJobDefinitionNetworkConfigArgs']]:
        return pulumi.get(self, "network_config")

    @network_config.setter
    def network_config(self, value: Optional[pulumi.Input['DataQualityJobDefinitionNetworkConfigArgs']]):
        pulumi.set(self, "network_config", value)

    @property
    @pulumi.getter(name="stoppingCondition")
    def stopping_condition(self) -> Optional[pulumi.Input['DataQualityJobDefinitionStoppingConditionArgs']]:
        return pulumi.get(self, "stopping_condition")

    @stopping_condition.setter
    def stopping_condition(self, value: Optional[pulumi.Input['DataQualityJobDefinitionStoppingConditionArgs']]):
        pulumi.set(self, "stopping_condition", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataQualityJobDefinitionTagArgs']]]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataQualityJobDefinitionTagArgs']]]]):
        pulumi.set(self, "tags", value)


class DataQualityJobDefinition(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_quality_app_specification: Optional[pulumi.Input[pulumi.InputType['DataQualityJobDefinitionDataQualityAppSpecificationArgs']]] = None,
                 data_quality_baseline_config: Optional[pulumi.Input[pulumi.InputType['DataQualityJobDefinitionDataQualityBaselineConfigArgs']]] = None,
                 data_quality_job_input: Optional[pulumi.Input[pulumi.InputType['DataQualityJobDefinitionDataQualityJobInputArgs']]] = None,
                 data_quality_job_output_config: Optional[pulumi.Input[pulumi.InputType['DataQualityJobDefinitionMonitoringOutputConfigArgs']]] = None,
                 job_definition_name: Optional[pulumi.Input[str]] = None,
                 job_resources: Optional[pulumi.Input[pulumi.InputType['DataQualityJobDefinitionMonitoringResourcesArgs']]] = None,
                 network_config: Optional[pulumi.Input[pulumi.InputType['DataQualityJobDefinitionNetworkConfigArgs']]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 stopping_condition: Optional[pulumi.Input[pulumi.InputType['DataQualityJobDefinitionStoppingConditionArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataQualityJobDefinitionTagArgs']]]]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::SageMaker::DataQualityJobDefinition

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataQualityJobDefinitionTagArgs']]]] tags: An array of key-value pairs to apply to this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DataQualityJobDefinitionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::SageMaker::DataQualityJobDefinition

        :param str resource_name: The name of the resource.
        :param DataQualityJobDefinitionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataQualityJobDefinitionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_quality_app_specification: Optional[pulumi.Input[pulumi.InputType['DataQualityJobDefinitionDataQualityAppSpecificationArgs']]] = None,
                 data_quality_baseline_config: Optional[pulumi.Input[pulumi.InputType['DataQualityJobDefinitionDataQualityBaselineConfigArgs']]] = None,
                 data_quality_job_input: Optional[pulumi.Input[pulumi.InputType['DataQualityJobDefinitionDataQualityJobInputArgs']]] = None,
                 data_quality_job_output_config: Optional[pulumi.Input[pulumi.InputType['DataQualityJobDefinitionMonitoringOutputConfigArgs']]] = None,
                 job_definition_name: Optional[pulumi.Input[str]] = None,
                 job_resources: Optional[pulumi.Input[pulumi.InputType['DataQualityJobDefinitionMonitoringResourcesArgs']]] = None,
                 network_config: Optional[pulumi.Input[pulumi.InputType['DataQualityJobDefinitionNetworkConfigArgs']]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 stopping_condition: Optional[pulumi.Input[pulumi.InputType['DataQualityJobDefinitionStoppingConditionArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataQualityJobDefinitionTagArgs']]]]] = None,
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
            __props__ = DataQualityJobDefinitionArgs.__new__(DataQualityJobDefinitionArgs)

            if data_quality_app_specification is None and not opts.urn:
                raise TypeError("Missing required property 'data_quality_app_specification'")
            __props__.__dict__["data_quality_app_specification"] = data_quality_app_specification
            __props__.__dict__["data_quality_baseline_config"] = data_quality_baseline_config
            if data_quality_job_input is None and not opts.urn:
                raise TypeError("Missing required property 'data_quality_job_input'")
            __props__.__dict__["data_quality_job_input"] = data_quality_job_input
            if data_quality_job_output_config is None and not opts.urn:
                raise TypeError("Missing required property 'data_quality_job_output_config'")
            __props__.__dict__["data_quality_job_output_config"] = data_quality_job_output_config
            __props__.__dict__["job_definition_name"] = job_definition_name
            if job_resources is None and not opts.urn:
                raise TypeError("Missing required property 'job_resources'")
            __props__.__dict__["job_resources"] = job_resources
            __props__.__dict__["network_config"] = network_config
            if role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'role_arn'")
            __props__.__dict__["role_arn"] = role_arn
            __props__.__dict__["stopping_condition"] = stopping_condition
            __props__.__dict__["tags"] = tags
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["job_definition_arn"] = None
        super(DataQualityJobDefinition, __self__).__init__(
            'aws-native:sagemaker:DataQualityJobDefinition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DataQualityJobDefinition':
        """
        Get an existing DataQualityJobDefinition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DataQualityJobDefinitionArgs.__new__(DataQualityJobDefinitionArgs)

        __props__.__dict__["creation_time"] = None
        __props__.__dict__["data_quality_app_specification"] = None
        __props__.__dict__["data_quality_baseline_config"] = None
        __props__.__dict__["data_quality_job_input"] = None
        __props__.__dict__["data_quality_job_output_config"] = None
        __props__.__dict__["job_definition_arn"] = None
        __props__.__dict__["job_definition_name"] = None
        __props__.__dict__["job_resources"] = None
        __props__.__dict__["network_config"] = None
        __props__.__dict__["role_arn"] = None
        __props__.__dict__["stopping_condition"] = None
        __props__.__dict__["tags"] = None
        return DataQualityJobDefinition(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        The time at which the job definition was created.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="dataQualityAppSpecification")
    def data_quality_app_specification(self) -> pulumi.Output['outputs.DataQualityJobDefinitionDataQualityAppSpecification']:
        return pulumi.get(self, "data_quality_app_specification")

    @property
    @pulumi.getter(name="dataQualityBaselineConfig")
    def data_quality_baseline_config(self) -> pulumi.Output[Optional['outputs.DataQualityJobDefinitionDataQualityBaselineConfig']]:
        return pulumi.get(self, "data_quality_baseline_config")

    @property
    @pulumi.getter(name="dataQualityJobInput")
    def data_quality_job_input(self) -> pulumi.Output['outputs.DataQualityJobDefinitionDataQualityJobInput']:
        return pulumi.get(self, "data_quality_job_input")

    @property
    @pulumi.getter(name="dataQualityJobOutputConfig")
    def data_quality_job_output_config(self) -> pulumi.Output['outputs.DataQualityJobDefinitionMonitoringOutputConfig']:
        return pulumi.get(self, "data_quality_job_output_config")

    @property
    @pulumi.getter(name="jobDefinitionArn")
    def job_definition_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of job definition.
        """
        return pulumi.get(self, "job_definition_arn")

    @property
    @pulumi.getter(name="jobDefinitionName")
    def job_definition_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "job_definition_name")

    @property
    @pulumi.getter(name="jobResources")
    def job_resources(self) -> pulumi.Output['outputs.DataQualityJobDefinitionMonitoringResources']:
        return pulumi.get(self, "job_resources")

    @property
    @pulumi.getter(name="networkConfig")
    def network_config(self) -> pulumi.Output[Optional['outputs.DataQualityJobDefinitionNetworkConfig']]:
        return pulumi.get(self, "network_config")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="stoppingCondition")
    def stopping_condition(self) -> pulumi.Output[Optional['outputs.DataQualityJobDefinitionStoppingCondition']]:
        return pulumi.get(self, "stopping_condition")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.DataQualityJobDefinitionTag']]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

