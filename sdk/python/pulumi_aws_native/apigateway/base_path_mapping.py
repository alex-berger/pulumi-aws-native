# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['BasePathMappingArgs', 'BasePathMapping']

@pulumi.input_type
class BasePathMappingArgs:
    def __init__(__self__, *,
                 domain_name: pulumi.Input[str],
                 base_path: Optional[pulumi.Input[str]] = None,
                 rest_api_id: Optional[pulumi.Input[str]] = None,
                 stage: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BasePathMapping resource.
        """
        pulumi.set(__self__, "domain_name", domain_name)
        if base_path is not None:
            pulumi.set(__self__, "base_path", base_path)
        if rest_api_id is not None:
            pulumi.set(__self__, "rest_api_id", rest_api_id)
        if stage is not None:
            pulumi.set(__self__, "stage", stage)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter(name="basePath")
    def base_path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "base_path")

    @base_path.setter
    def base_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "base_path", value)

    @property
    @pulumi.getter(name="restApiId")
    def rest_api_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "rest_api_id")

    @rest_api_id.setter
    def rest_api_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rest_api_id", value)

    @property
    @pulumi.getter
    def stage(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "stage")

    @stage.setter
    def stage(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stage", value)


class BasePathMapping(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 base_path: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 rest_api_id: Optional[pulumi.Input[str]] = None,
                 stage: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::ApiGateway::BasePathMapping

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BasePathMappingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::ApiGateway::BasePathMapping

        :param str resource_name: The name of the resource.
        :param BasePathMappingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BasePathMappingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 base_path: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 rest_api_id: Optional[pulumi.Input[str]] = None,
                 stage: Optional[pulumi.Input[str]] = None,
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
            __props__ = BasePathMappingArgs.__new__(BasePathMappingArgs)

            __props__.__dict__["base_path"] = base_path
            if domain_name is None and not opts.urn:
                raise TypeError("Missing required property 'domain_name'")
            __props__.__dict__["domain_name"] = domain_name
            __props__.__dict__["rest_api_id"] = rest_api_id
            __props__.__dict__["stage"] = stage
        super(BasePathMapping, __self__).__init__(
            'aws-native:apigateway:BasePathMapping',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'BasePathMapping':
        """
        Get an existing BasePathMapping resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = BasePathMappingArgs.__new__(BasePathMappingArgs)

        __props__.__dict__["base_path"] = None
        __props__.__dict__["domain_name"] = None
        __props__.__dict__["rest_api_id"] = None
        __props__.__dict__["stage"] = None
        return BasePathMapping(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="basePath")
    def base_path(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "base_path")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter(name="restApiId")
    def rest_api_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "rest_api_id")

    @property
    @pulumi.getter
    def stage(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "stage")

