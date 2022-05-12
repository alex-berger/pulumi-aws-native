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

__all__ = ['FunctionConfigurationArgs', 'FunctionConfiguration']

@pulumi.input_type
class FunctionConfigurationArgs:
    def __init__(__self__, *,
                 api_id: pulumi.Input[str],
                 data_source_name: pulumi.Input[str],
                 function_version: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 max_batch_size: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 request_mapping_template: Optional[pulumi.Input[str]] = None,
                 request_mapping_template_s3_location: Optional[pulumi.Input[str]] = None,
                 response_mapping_template: Optional[pulumi.Input[str]] = None,
                 response_mapping_template_s3_location: Optional[pulumi.Input[str]] = None,
                 sync_config: Optional[pulumi.Input['FunctionConfigurationSyncConfigArgs']] = None):
        """
        The set of arguments for constructing a FunctionConfiguration resource.
        """
        pulumi.set(__self__, "api_id", api_id)
        pulumi.set(__self__, "data_source_name", data_source_name)
        pulumi.set(__self__, "function_version", function_version)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if max_batch_size is not None:
            pulumi.set(__self__, "max_batch_size", max_batch_size)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if request_mapping_template is not None:
            pulumi.set(__self__, "request_mapping_template", request_mapping_template)
        if request_mapping_template_s3_location is not None:
            pulumi.set(__self__, "request_mapping_template_s3_location", request_mapping_template_s3_location)
        if response_mapping_template is not None:
            pulumi.set(__self__, "response_mapping_template", response_mapping_template)
        if response_mapping_template_s3_location is not None:
            pulumi.set(__self__, "response_mapping_template_s3_location", response_mapping_template_s3_location)
        if sync_config is not None:
            pulumi.set(__self__, "sync_config", sync_config)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter(name="dataSourceName")
    def data_source_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "data_source_name")

    @data_source_name.setter
    def data_source_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "data_source_name", value)

    @property
    @pulumi.getter(name="functionVersion")
    def function_version(self) -> pulumi.Input[str]:
        return pulumi.get(self, "function_version")

    @function_version.setter
    def function_version(self, value: pulumi.Input[str]):
        pulumi.set(self, "function_version", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="maxBatchSize")
    def max_batch_size(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "max_batch_size")

    @max_batch_size.setter
    def max_batch_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_batch_size", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="requestMappingTemplate")
    def request_mapping_template(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_mapping_template")

    @request_mapping_template.setter
    def request_mapping_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_mapping_template", value)

    @property
    @pulumi.getter(name="requestMappingTemplateS3Location")
    def request_mapping_template_s3_location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_mapping_template_s3_location")

    @request_mapping_template_s3_location.setter
    def request_mapping_template_s3_location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_mapping_template_s3_location", value)

    @property
    @pulumi.getter(name="responseMappingTemplate")
    def response_mapping_template(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "response_mapping_template")

    @response_mapping_template.setter
    def response_mapping_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "response_mapping_template", value)

    @property
    @pulumi.getter(name="responseMappingTemplateS3Location")
    def response_mapping_template_s3_location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "response_mapping_template_s3_location")

    @response_mapping_template_s3_location.setter
    def response_mapping_template_s3_location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "response_mapping_template_s3_location", value)

    @property
    @pulumi.getter(name="syncConfig")
    def sync_config(self) -> Optional[pulumi.Input['FunctionConfigurationSyncConfigArgs']]:
        return pulumi.get(self, "sync_config")

    @sync_config.setter
    def sync_config(self, value: Optional[pulumi.Input['FunctionConfigurationSyncConfigArgs']]):
        pulumi.set(self, "sync_config", value)


warnings.warn("""FunctionConfiguration is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class FunctionConfiguration(pulumi.CustomResource):
    warnings.warn("""FunctionConfiguration is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 data_source_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 function_version: Optional[pulumi.Input[str]] = None,
                 max_batch_size: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 request_mapping_template: Optional[pulumi.Input[str]] = None,
                 request_mapping_template_s3_location: Optional[pulumi.Input[str]] = None,
                 response_mapping_template: Optional[pulumi.Input[str]] = None,
                 response_mapping_template_s3_location: Optional[pulumi.Input[str]] = None,
                 sync_config: Optional[pulumi.Input[pulumi.InputType['FunctionConfigurationSyncConfigArgs']]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::AppSync::FunctionConfiguration

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FunctionConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::AppSync::FunctionConfiguration

        :param str resource_name: The name of the resource.
        :param FunctionConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FunctionConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 data_source_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 function_version: Optional[pulumi.Input[str]] = None,
                 max_batch_size: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 request_mapping_template: Optional[pulumi.Input[str]] = None,
                 request_mapping_template_s3_location: Optional[pulumi.Input[str]] = None,
                 response_mapping_template: Optional[pulumi.Input[str]] = None,
                 response_mapping_template_s3_location: Optional[pulumi.Input[str]] = None,
                 sync_config: Optional[pulumi.Input[pulumi.InputType['FunctionConfigurationSyncConfigArgs']]] = None,
                 __props__=None):
        pulumi.log.warn("""FunctionConfiguration is deprecated: FunctionConfiguration is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FunctionConfigurationArgs.__new__(FunctionConfigurationArgs)

            if api_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_id'")
            __props__.__dict__["api_id"] = api_id
            if data_source_name is None and not opts.urn:
                raise TypeError("Missing required property 'data_source_name'")
            __props__.__dict__["data_source_name"] = data_source_name
            __props__.__dict__["description"] = description
            if function_version is None and not opts.urn:
                raise TypeError("Missing required property 'function_version'")
            __props__.__dict__["function_version"] = function_version
            __props__.__dict__["max_batch_size"] = max_batch_size
            __props__.__dict__["name"] = name
            __props__.__dict__["request_mapping_template"] = request_mapping_template
            __props__.__dict__["request_mapping_template_s3_location"] = request_mapping_template_s3_location
            __props__.__dict__["response_mapping_template"] = response_mapping_template
            __props__.__dict__["response_mapping_template_s3_location"] = response_mapping_template_s3_location
            __props__.__dict__["sync_config"] = sync_config
            __props__.__dict__["function_arn"] = None
            __props__.__dict__["function_id"] = None
        super(FunctionConfiguration, __self__).__init__(
            'aws-native:appsync:FunctionConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'FunctionConfiguration':
        """
        Get an existing FunctionConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FunctionConfigurationArgs.__new__(FunctionConfigurationArgs)

        __props__.__dict__["api_id"] = None
        __props__.__dict__["data_source_name"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["function_arn"] = None
        __props__.__dict__["function_id"] = None
        __props__.__dict__["function_version"] = None
        __props__.__dict__["max_batch_size"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["request_mapping_template"] = None
        __props__.__dict__["request_mapping_template_s3_location"] = None
        __props__.__dict__["response_mapping_template"] = None
        __props__.__dict__["response_mapping_template_s3_location"] = None
        __props__.__dict__["sync_config"] = None
        return FunctionConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter(name="dataSourceName")
    def data_source_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "data_source_name")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="functionArn")
    def function_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "function_arn")

    @property
    @pulumi.getter(name="functionId")
    def function_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "function_id")

    @property
    @pulumi.getter(name="functionVersion")
    def function_version(self) -> pulumi.Output[str]:
        return pulumi.get(self, "function_version")

    @property
    @pulumi.getter(name="maxBatchSize")
    def max_batch_size(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "max_batch_size")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="requestMappingTemplate")
    def request_mapping_template(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "request_mapping_template")

    @property
    @pulumi.getter(name="requestMappingTemplateS3Location")
    def request_mapping_template_s3_location(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "request_mapping_template_s3_location")

    @property
    @pulumi.getter(name="responseMappingTemplate")
    def response_mapping_template(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "response_mapping_template")

    @property
    @pulumi.getter(name="responseMappingTemplateS3Location")
    def response_mapping_template_s3_location(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "response_mapping_template_s3_location")

    @property
    @pulumi.getter(name="syncConfig")
    def sync_config(self) -> pulumi.Output[Optional['outputs.FunctionConfigurationSyncConfig']]:
        return pulumi.get(self, "sync_config")

