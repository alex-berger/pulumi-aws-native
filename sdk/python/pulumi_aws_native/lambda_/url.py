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

__all__ = ['UrlArgs', 'Url']

@pulumi.input_type
class UrlArgs:
    def __init__(__self__, *,
                 auth_type: pulumi.Input['UrlAuthType'],
                 target_function_arn: pulumi.Input[str],
                 cors: Optional[pulumi.Input['UrlCorsArgs']] = None,
                 qualifier: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Url resource.
        :param pulumi.Input['UrlAuthType'] auth_type: Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.
        :param pulumi.Input[str] target_function_arn: The Amazon Resource Name (ARN) of the function associated with the Function URL.
        :param pulumi.Input[str] qualifier: The alias qualifier for the target function. If TargetFunctionArn is unqualified then Qualifier must be passed.
        """
        pulumi.set(__self__, "auth_type", auth_type)
        pulumi.set(__self__, "target_function_arn", target_function_arn)
        if cors is not None:
            pulumi.set(__self__, "cors", cors)
        if qualifier is not None:
            pulumi.set(__self__, "qualifier", qualifier)

    @property
    @pulumi.getter(name="authType")
    def auth_type(self) -> pulumi.Input['UrlAuthType']:
        """
        Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.
        """
        return pulumi.get(self, "auth_type")

    @auth_type.setter
    def auth_type(self, value: pulumi.Input['UrlAuthType']):
        pulumi.set(self, "auth_type", value)

    @property
    @pulumi.getter(name="targetFunctionArn")
    def target_function_arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN) of the function associated with the Function URL.
        """
        return pulumi.get(self, "target_function_arn")

    @target_function_arn.setter
    def target_function_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_function_arn", value)

    @property
    @pulumi.getter
    def cors(self) -> Optional[pulumi.Input['UrlCorsArgs']]:
        return pulumi.get(self, "cors")

    @cors.setter
    def cors(self, value: Optional[pulumi.Input['UrlCorsArgs']]):
        pulumi.set(self, "cors", value)

    @property
    @pulumi.getter
    def qualifier(self) -> Optional[pulumi.Input[str]]:
        """
        The alias qualifier for the target function. If TargetFunctionArn is unqualified then Qualifier must be passed.
        """
        return pulumi.get(self, "qualifier")

    @qualifier.setter
    def qualifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "qualifier", value)


class Url(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_type: Optional[pulumi.Input['UrlAuthType']] = None,
                 cors: Optional[pulumi.Input[pulumi.InputType['UrlCorsArgs']]] = None,
                 qualifier: Optional[pulumi.Input[str]] = None,
                 target_function_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::Lambda::Url

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['UrlAuthType'] auth_type: Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.
        :param pulumi.Input[str] qualifier: The alias qualifier for the target function. If TargetFunctionArn is unqualified then Qualifier must be passed.
        :param pulumi.Input[str] target_function_arn: The Amazon Resource Name (ARN) of the function associated with the Function URL.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UrlArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::Lambda::Url

        :param str resource_name: The name of the resource.
        :param UrlArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UrlArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_type: Optional[pulumi.Input['UrlAuthType']] = None,
                 cors: Optional[pulumi.Input[pulumi.InputType['UrlCorsArgs']]] = None,
                 qualifier: Optional[pulumi.Input[str]] = None,
                 target_function_arn: Optional[pulumi.Input[str]] = None,
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
            __props__ = UrlArgs.__new__(UrlArgs)

            if auth_type is None and not opts.urn:
                raise TypeError("Missing required property 'auth_type'")
            __props__.__dict__["auth_type"] = auth_type
            __props__.__dict__["cors"] = cors
            __props__.__dict__["qualifier"] = qualifier
            if target_function_arn is None and not opts.urn:
                raise TypeError("Missing required property 'target_function_arn'")
            __props__.__dict__["target_function_arn"] = target_function_arn
            __props__.__dict__["function_arn"] = None
            __props__.__dict__["function_url"] = None
        super(Url, __self__).__init__(
            'aws-native:lambda:Url',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Url':
        """
        Get an existing Url resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = UrlArgs.__new__(UrlArgs)

        __props__.__dict__["auth_type"] = None
        __props__.__dict__["cors"] = None
        __props__.__dict__["function_arn"] = None
        __props__.__dict__["function_url"] = None
        __props__.__dict__["qualifier"] = None
        __props__.__dict__["target_function_arn"] = None
        return Url(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authType")
    def auth_type(self) -> pulumi.Output['UrlAuthType']:
        """
        Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.
        """
        return pulumi.get(self, "auth_type")

    @property
    @pulumi.getter
    def cors(self) -> pulumi.Output[Optional['outputs.UrlCors']]:
        return pulumi.get(self, "cors")

    @property
    @pulumi.getter(name="functionArn")
    def function_arn(self) -> pulumi.Output[str]:
        """
        The full Amazon Resource Name (ARN) of the function associated with the Function URL.
        """
        return pulumi.get(self, "function_arn")

    @property
    @pulumi.getter(name="functionUrl")
    def function_url(self) -> pulumi.Output[str]:
        """
        The generated url for this resource.
        """
        return pulumi.get(self, "function_url")

    @property
    @pulumi.getter
    def qualifier(self) -> pulumi.Output[Optional[str]]:
        """
        The alias qualifier for the target function. If TargetFunctionArn is unqualified then Qualifier must be passed.
        """
        return pulumi.get(self, "qualifier")

    @property
    @pulumi.getter(name="targetFunctionArn")
    def target_function_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the function associated with the Function URL.
        """
        return pulumi.get(self, "target_function_arn")

