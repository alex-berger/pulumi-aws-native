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

__all__ = ['JobArgs', 'Job']

@pulumi.input_type
class JobArgs:
    def __init__(__self__, *,
                 role_arn: pulumi.Input[str],
                 type: pulumi.Input['JobType'],
                 data_catalog_outputs: Optional[pulumi.Input[Sequence[pulumi.Input['JobDataCatalogOutputArgs']]]] = None,
                 database_outputs: Optional[pulumi.Input[Sequence[pulumi.Input['JobDatabaseOutputArgs']]]] = None,
                 dataset_name: Optional[pulumi.Input[str]] = None,
                 encryption_key_arn: Optional[pulumi.Input[str]] = None,
                 encryption_mode: Optional[pulumi.Input['JobEncryptionMode']] = None,
                 job_sample: Optional[pulumi.Input['JobSampleArgs']] = None,
                 log_subscription: Optional[pulumi.Input['JobLogSubscription']] = None,
                 max_capacity: Optional[pulumi.Input[int]] = None,
                 max_retries: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 output_location: Optional[pulumi.Input['JobOutputLocationArgs']] = None,
                 outputs: Optional[pulumi.Input[Sequence[pulumi.Input['JobOutputArgs']]]] = None,
                 profile_configuration: Optional[pulumi.Input['JobProfileConfigurationArgs']] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 recipe: Optional[pulumi.Input['JobRecipeArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['JobTagArgs']]]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 validation_configurations: Optional[pulumi.Input[Sequence[pulumi.Input['JobValidationConfigurationArgs']]]] = None):
        """
        The set of arguments for constructing a Job resource.
        :param pulumi.Input[str] role_arn: Role arn
        :param pulumi.Input['JobType'] type: Job type
        :param pulumi.Input[str] dataset_name: Dataset name
        :param pulumi.Input[str] encryption_key_arn: Encryption Key Arn
        :param pulumi.Input['JobEncryptionMode'] encryption_mode: Encryption mode
        :param pulumi.Input['JobSampleArgs'] job_sample: Job Sample
        :param pulumi.Input['JobLogSubscription'] log_subscription: Log subscription
        :param pulumi.Input[int] max_capacity: Max capacity
        :param pulumi.Input[int] max_retries: Max retries
        :param pulumi.Input[str] name: Job name
        :param pulumi.Input['JobOutputLocationArgs'] output_location: Output location
        :param pulumi.Input['JobProfileConfigurationArgs'] profile_configuration: Profile Job configuration
        :param pulumi.Input[str] project_name: Project name
        :param pulumi.Input[int] timeout: Timeout
        :param pulumi.Input[Sequence[pulumi.Input['JobValidationConfigurationArgs']]] validation_configurations: Data quality rules configuration
        """
        pulumi.set(__self__, "role_arn", role_arn)
        pulumi.set(__self__, "type", type)
        if data_catalog_outputs is not None:
            pulumi.set(__self__, "data_catalog_outputs", data_catalog_outputs)
        if database_outputs is not None:
            pulumi.set(__self__, "database_outputs", database_outputs)
        if dataset_name is not None:
            pulumi.set(__self__, "dataset_name", dataset_name)
        if encryption_key_arn is not None:
            pulumi.set(__self__, "encryption_key_arn", encryption_key_arn)
        if encryption_mode is not None:
            pulumi.set(__self__, "encryption_mode", encryption_mode)
        if job_sample is not None:
            pulumi.set(__self__, "job_sample", job_sample)
        if log_subscription is not None:
            pulumi.set(__self__, "log_subscription", log_subscription)
        if max_capacity is not None:
            pulumi.set(__self__, "max_capacity", max_capacity)
        if max_retries is not None:
            pulumi.set(__self__, "max_retries", max_retries)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if output_location is not None:
            pulumi.set(__self__, "output_location", output_location)
        if outputs is not None:
            pulumi.set(__self__, "outputs", outputs)
        if profile_configuration is not None:
            pulumi.set(__self__, "profile_configuration", profile_configuration)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if recipe is not None:
            pulumi.set(__self__, "recipe", recipe)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if validation_configurations is not None:
            pulumi.set(__self__, "validation_configurations", validation_configurations)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Input[str]:
        """
        Role arn
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['JobType']:
        """
        Job type
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['JobType']):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="dataCatalogOutputs")
    def data_catalog_outputs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['JobDataCatalogOutputArgs']]]]:
        return pulumi.get(self, "data_catalog_outputs")

    @data_catalog_outputs.setter
    def data_catalog_outputs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['JobDataCatalogOutputArgs']]]]):
        pulumi.set(self, "data_catalog_outputs", value)

    @property
    @pulumi.getter(name="databaseOutputs")
    def database_outputs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['JobDatabaseOutputArgs']]]]:
        return pulumi.get(self, "database_outputs")

    @database_outputs.setter
    def database_outputs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['JobDatabaseOutputArgs']]]]):
        pulumi.set(self, "database_outputs", value)

    @property
    @pulumi.getter(name="datasetName")
    def dataset_name(self) -> Optional[pulumi.Input[str]]:
        """
        Dataset name
        """
        return pulumi.get(self, "dataset_name")

    @dataset_name.setter
    def dataset_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dataset_name", value)

    @property
    @pulumi.getter(name="encryptionKeyArn")
    def encryption_key_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Encryption Key Arn
        """
        return pulumi.get(self, "encryption_key_arn")

    @encryption_key_arn.setter
    def encryption_key_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encryption_key_arn", value)

    @property
    @pulumi.getter(name="encryptionMode")
    def encryption_mode(self) -> Optional[pulumi.Input['JobEncryptionMode']]:
        """
        Encryption mode
        """
        return pulumi.get(self, "encryption_mode")

    @encryption_mode.setter
    def encryption_mode(self, value: Optional[pulumi.Input['JobEncryptionMode']]):
        pulumi.set(self, "encryption_mode", value)

    @property
    @pulumi.getter(name="jobSample")
    def job_sample(self) -> Optional[pulumi.Input['JobSampleArgs']]:
        """
        Job Sample
        """
        return pulumi.get(self, "job_sample")

    @job_sample.setter
    def job_sample(self, value: Optional[pulumi.Input['JobSampleArgs']]):
        pulumi.set(self, "job_sample", value)

    @property
    @pulumi.getter(name="logSubscription")
    def log_subscription(self) -> Optional[pulumi.Input['JobLogSubscription']]:
        """
        Log subscription
        """
        return pulumi.get(self, "log_subscription")

    @log_subscription.setter
    def log_subscription(self, value: Optional[pulumi.Input['JobLogSubscription']]):
        pulumi.set(self, "log_subscription", value)

    @property
    @pulumi.getter(name="maxCapacity")
    def max_capacity(self) -> Optional[pulumi.Input[int]]:
        """
        Max capacity
        """
        return pulumi.get(self, "max_capacity")

    @max_capacity.setter
    def max_capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_capacity", value)

    @property
    @pulumi.getter(name="maxRetries")
    def max_retries(self) -> Optional[pulumi.Input[int]]:
        """
        Max retries
        """
        return pulumi.get(self, "max_retries")

    @max_retries.setter
    def max_retries(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_retries", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Job name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="outputLocation")
    def output_location(self) -> Optional[pulumi.Input['JobOutputLocationArgs']]:
        """
        Output location
        """
        return pulumi.get(self, "output_location")

    @output_location.setter
    def output_location(self, value: Optional[pulumi.Input['JobOutputLocationArgs']]):
        pulumi.set(self, "output_location", value)

    @property
    @pulumi.getter
    def outputs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['JobOutputArgs']]]]:
        return pulumi.get(self, "outputs")

    @outputs.setter
    def outputs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['JobOutputArgs']]]]):
        pulumi.set(self, "outputs", value)

    @property
    @pulumi.getter(name="profileConfiguration")
    def profile_configuration(self) -> Optional[pulumi.Input['JobProfileConfigurationArgs']]:
        """
        Profile Job configuration
        """
        return pulumi.get(self, "profile_configuration")

    @profile_configuration.setter
    def profile_configuration(self, value: Optional[pulumi.Input['JobProfileConfigurationArgs']]):
        pulumi.set(self, "profile_configuration", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        Project name
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter
    def recipe(self) -> Optional[pulumi.Input['JobRecipeArgs']]:
        return pulumi.get(self, "recipe")

    @recipe.setter
    def recipe(self, value: Optional[pulumi.Input['JobRecipeArgs']]):
        pulumi.set(self, "recipe", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['JobTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['JobTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Timeout
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter(name="validationConfigurations")
    def validation_configurations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['JobValidationConfigurationArgs']]]]:
        """
        Data quality rules configuration
        """
        return pulumi.get(self, "validation_configurations")

    @validation_configurations.setter
    def validation_configurations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['JobValidationConfigurationArgs']]]]):
        pulumi.set(self, "validation_configurations", value)


class Job(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_catalog_outputs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['JobDataCatalogOutputArgs']]]]] = None,
                 database_outputs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['JobDatabaseOutputArgs']]]]] = None,
                 dataset_name: Optional[pulumi.Input[str]] = None,
                 encryption_key_arn: Optional[pulumi.Input[str]] = None,
                 encryption_mode: Optional[pulumi.Input['JobEncryptionMode']] = None,
                 job_sample: Optional[pulumi.Input[pulumi.InputType['JobSampleArgs']]] = None,
                 log_subscription: Optional[pulumi.Input['JobLogSubscription']] = None,
                 max_capacity: Optional[pulumi.Input[int]] = None,
                 max_retries: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 output_location: Optional[pulumi.Input[pulumi.InputType['JobOutputLocationArgs']]] = None,
                 outputs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['JobOutputArgs']]]]] = None,
                 profile_configuration: Optional[pulumi.Input[pulumi.InputType['JobProfileConfigurationArgs']]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 recipe: Optional[pulumi.Input[pulumi.InputType['JobRecipeArgs']]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['JobTagArgs']]]]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input['JobType']] = None,
                 validation_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['JobValidationConfigurationArgs']]]]] = None,
                 __props__=None):
        """
        Resource schema for AWS::DataBrew::Job.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dataset_name: Dataset name
        :param pulumi.Input[str] encryption_key_arn: Encryption Key Arn
        :param pulumi.Input['JobEncryptionMode'] encryption_mode: Encryption mode
        :param pulumi.Input[pulumi.InputType['JobSampleArgs']] job_sample: Job Sample
        :param pulumi.Input['JobLogSubscription'] log_subscription: Log subscription
        :param pulumi.Input[int] max_capacity: Max capacity
        :param pulumi.Input[int] max_retries: Max retries
        :param pulumi.Input[str] name: Job name
        :param pulumi.Input[pulumi.InputType['JobOutputLocationArgs']] output_location: Output location
        :param pulumi.Input[pulumi.InputType['JobProfileConfigurationArgs']] profile_configuration: Profile Job configuration
        :param pulumi.Input[str] project_name: Project name
        :param pulumi.Input[str] role_arn: Role arn
        :param pulumi.Input[int] timeout: Timeout
        :param pulumi.Input['JobType'] type: Job type
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['JobValidationConfigurationArgs']]]] validation_configurations: Data quality rules configuration
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: JobArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource schema for AWS::DataBrew::Job.

        :param str resource_name: The name of the resource.
        :param JobArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(JobArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_catalog_outputs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['JobDataCatalogOutputArgs']]]]] = None,
                 database_outputs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['JobDatabaseOutputArgs']]]]] = None,
                 dataset_name: Optional[pulumi.Input[str]] = None,
                 encryption_key_arn: Optional[pulumi.Input[str]] = None,
                 encryption_mode: Optional[pulumi.Input['JobEncryptionMode']] = None,
                 job_sample: Optional[pulumi.Input[pulumi.InputType['JobSampleArgs']]] = None,
                 log_subscription: Optional[pulumi.Input['JobLogSubscription']] = None,
                 max_capacity: Optional[pulumi.Input[int]] = None,
                 max_retries: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 output_location: Optional[pulumi.Input[pulumi.InputType['JobOutputLocationArgs']]] = None,
                 outputs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['JobOutputArgs']]]]] = None,
                 profile_configuration: Optional[pulumi.Input[pulumi.InputType['JobProfileConfigurationArgs']]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 recipe: Optional[pulumi.Input[pulumi.InputType['JobRecipeArgs']]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['JobTagArgs']]]]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input['JobType']] = None,
                 validation_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['JobValidationConfigurationArgs']]]]] = None,
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
            __props__ = JobArgs.__new__(JobArgs)

            __props__.__dict__["data_catalog_outputs"] = data_catalog_outputs
            __props__.__dict__["database_outputs"] = database_outputs
            __props__.__dict__["dataset_name"] = dataset_name
            __props__.__dict__["encryption_key_arn"] = encryption_key_arn
            __props__.__dict__["encryption_mode"] = encryption_mode
            __props__.__dict__["job_sample"] = job_sample
            __props__.__dict__["log_subscription"] = log_subscription
            __props__.__dict__["max_capacity"] = max_capacity
            __props__.__dict__["max_retries"] = max_retries
            __props__.__dict__["name"] = name
            __props__.__dict__["output_location"] = output_location
            __props__.__dict__["outputs"] = outputs
            __props__.__dict__["profile_configuration"] = profile_configuration
            __props__.__dict__["project_name"] = project_name
            __props__.__dict__["recipe"] = recipe
            if role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'role_arn'")
            __props__.__dict__["role_arn"] = role_arn
            __props__.__dict__["tags"] = tags
            __props__.__dict__["timeout"] = timeout
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["validation_configurations"] = validation_configurations
        super(Job, __self__).__init__(
            'aws-native:databrew:Job',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Job':
        """
        Get an existing Job resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = JobArgs.__new__(JobArgs)

        __props__.__dict__["data_catalog_outputs"] = None
        __props__.__dict__["database_outputs"] = None
        __props__.__dict__["dataset_name"] = None
        __props__.__dict__["encryption_key_arn"] = None
        __props__.__dict__["encryption_mode"] = None
        __props__.__dict__["job_sample"] = None
        __props__.__dict__["log_subscription"] = None
        __props__.__dict__["max_capacity"] = None
        __props__.__dict__["max_retries"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["output_location"] = None
        __props__.__dict__["outputs"] = None
        __props__.__dict__["profile_configuration"] = None
        __props__.__dict__["project_name"] = None
        __props__.__dict__["recipe"] = None
        __props__.__dict__["role_arn"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["timeout"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["validation_configurations"] = None
        return Job(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dataCatalogOutputs")
    def data_catalog_outputs(self) -> pulumi.Output[Optional[Sequence['outputs.JobDataCatalogOutput']]]:
        return pulumi.get(self, "data_catalog_outputs")

    @property
    @pulumi.getter(name="databaseOutputs")
    def database_outputs(self) -> pulumi.Output[Optional[Sequence['outputs.JobDatabaseOutput']]]:
        return pulumi.get(self, "database_outputs")

    @property
    @pulumi.getter(name="datasetName")
    def dataset_name(self) -> pulumi.Output[Optional[str]]:
        """
        Dataset name
        """
        return pulumi.get(self, "dataset_name")

    @property
    @pulumi.getter(name="encryptionKeyArn")
    def encryption_key_arn(self) -> pulumi.Output[Optional[str]]:
        """
        Encryption Key Arn
        """
        return pulumi.get(self, "encryption_key_arn")

    @property
    @pulumi.getter(name="encryptionMode")
    def encryption_mode(self) -> pulumi.Output[Optional['JobEncryptionMode']]:
        """
        Encryption mode
        """
        return pulumi.get(self, "encryption_mode")

    @property
    @pulumi.getter(name="jobSample")
    def job_sample(self) -> pulumi.Output[Optional['outputs.JobSample']]:
        """
        Job Sample
        """
        return pulumi.get(self, "job_sample")

    @property
    @pulumi.getter(name="logSubscription")
    def log_subscription(self) -> pulumi.Output[Optional['JobLogSubscription']]:
        """
        Log subscription
        """
        return pulumi.get(self, "log_subscription")

    @property
    @pulumi.getter(name="maxCapacity")
    def max_capacity(self) -> pulumi.Output[Optional[int]]:
        """
        Max capacity
        """
        return pulumi.get(self, "max_capacity")

    @property
    @pulumi.getter(name="maxRetries")
    def max_retries(self) -> pulumi.Output[Optional[int]]:
        """
        Max retries
        """
        return pulumi.get(self, "max_retries")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Job name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outputLocation")
    def output_location(self) -> pulumi.Output[Optional['outputs.JobOutputLocation']]:
        """
        Output location
        """
        return pulumi.get(self, "output_location")

    @property
    @pulumi.getter
    def outputs(self) -> pulumi.Output[Optional[Sequence['outputs.JobOutput']]]:
        return pulumi.get(self, "outputs")

    @property
    @pulumi.getter(name="profileConfiguration")
    def profile_configuration(self) -> pulumi.Output[Optional['outputs.JobProfileConfiguration']]:
        """
        Profile Job configuration
        """
        return pulumi.get(self, "profile_configuration")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Output[Optional[str]]:
        """
        Project name
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter
    def recipe(self) -> pulumi.Output[Optional['outputs.JobRecipe']]:
        return pulumi.get(self, "recipe")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[str]:
        """
        Role arn
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.JobTag']]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[Optional[int]]:
        """
        Timeout
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output['JobType']:
        """
        Job type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="validationConfigurations")
    def validation_configurations(self) -> pulumi.Output[Optional[Sequence['outputs.JobValidationConfiguration']]]:
        """
        Data quality rules configuration
        """
        return pulumi.get(self, "validation_configurations")

