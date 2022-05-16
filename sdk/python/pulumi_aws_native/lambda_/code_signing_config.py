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

__all__ = ['CodeSigningConfigArgs', 'CodeSigningConfig']

@pulumi.input_type
class CodeSigningConfigArgs:
    def __init__(__self__, *,
                 allowed_publishers: pulumi.Input['CodeSigningConfigAllowedPublishersArgs'],
                 code_signing_policies: Optional[pulumi.Input['CodeSigningConfigCodeSigningPoliciesArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CodeSigningConfig resource.
        :param pulumi.Input['CodeSigningConfigAllowedPublishersArgs'] allowed_publishers: When the CodeSigningConfig is later on attached to a function, the function code will be expected to be signed by profiles from this list
        :param pulumi.Input['CodeSigningConfigCodeSigningPoliciesArgs'] code_signing_policies: Policies to control how to act if a signature is invalid
        :param pulumi.Input[str] description: A description of the CodeSigningConfig
        """
        pulumi.set(__self__, "allowed_publishers", allowed_publishers)
        if code_signing_policies is not None:
            pulumi.set(__self__, "code_signing_policies", code_signing_policies)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="allowedPublishers")
    def allowed_publishers(self) -> pulumi.Input['CodeSigningConfigAllowedPublishersArgs']:
        """
        When the CodeSigningConfig is later on attached to a function, the function code will be expected to be signed by profiles from this list
        """
        return pulumi.get(self, "allowed_publishers")

    @allowed_publishers.setter
    def allowed_publishers(self, value: pulumi.Input['CodeSigningConfigAllowedPublishersArgs']):
        pulumi.set(self, "allowed_publishers", value)

    @property
    @pulumi.getter(name="codeSigningPolicies")
    def code_signing_policies(self) -> Optional[pulumi.Input['CodeSigningConfigCodeSigningPoliciesArgs']]:
        """
        Policies to control how to act if a signature is invalid
        """
        return pulumi.get(self, "code_signing_policies")

    @code_signing_policies.setter
    def code_signing_policies(self, value: Optional[pulumi.Input['CodeSigningConfigCodeSigningPoliciesArgs']]):
        pulumi.set(self, "code_signing_policies", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the CodeSigningConfig
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


class CodeSigningConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_publishers: Optional[pulumi.Input[pulumi.InputType['CodeSigningConfigAllowedPublishersArgs']]] = None,
                 code_signing_policies: Optional[pulumi.Input[pulumi.InputType['CodeSigningConfigCodeSigningPoliciesArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::Lambda::CodeSigningConfig.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['CodeSigningConfigAllowedPublishersArgs']] allowed_publishers: When the CodeSigningConfig is later on attached to a function, the function code will be expected to be signed by profiles from this list
        :param pulumi.Input[pulumi.InputType['CodeSigningConfigCodeSigningPoliciesArgs']] code_signing_policies: Policies to control how to act if a signature is invalid
        :param pulumi.Input[str] description: A description of the CodeSigningConfig
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CodeSigningConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::Lambda::CodeSigningConfig.

        :param str resource_name: The name of the resource.
        :param CodeSigningConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CodeSigningConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_publishers: Optional[pulumi.Input[pulumi.InputType['CodeSigningConfigAllowedPublishersArgs']]] = None,
                 code_signing_policies: Optional[pulumi.Input[pulumi.InputType['CodeSigningConfigCodeSigningPoliciesArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
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
            __props__ = CodeSigningConfigArgs.__new__(CodeSigningConfigArgs)

            if allowed_publishers is None and not opts.urn:
                raise TypeError("Missing required property 'allowed_publishers'")
            __props__.__dict__["allowed_publishers"] = allowed_publishers
            __props__.__dict__["code_signing_policies"] = code_signing_policies
            __props__.__dict__["description"] = description
            __props__.__dict__["code_signing_config_arn"] = None
            __props__.__dict__["code_signing_config_id"] = None
        super(CodeSigningConfig, __self__).__init__(
            'aws-native:lambda:CodeSigningConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'CodeSigningConfig':
        """
        Get an existing CodeSigningConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = CodeSigningConfigArgs.__new__(CodeSigningConfigArgs)

        __props__.__dict__["allowed_publishers"] = None
        __props__.__dict__["code_signing_config_arn"] = None
        __props__.__dict__["code_signing_config_id"] = None
        __props__.__dict__["code_signing_policies"] = None
        __props__.__dict__["description"] = None
        return CodeSigningConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowedPublishers")
    def allowed_publishers(self) -> pulumi.Output['outputs.CodeSigningConfigAllowedPublishers']:
        """
        When the CodeSigningConfig is later on attached to a function, the function code will be expected to be signed by profiles from this list
        """
        return pulumi.get(self, "allowed_publishers")

    @property
    @pulumi.getter(name="codeSigningConfigArn")
    def code_signing_config_arn(self) -> pulumi.Output[str]:
        """
        A unique Arn for CodeSigningConfig resource
        """
        return pulumi.get(self, "code_signing_config_arn")

    @property
    @pulumi.getter(name="codeSigningConfigId")
    def code_signing_config_id(self) -> pulumi.Output[str]:
        """
        A unique identifier for CodeSigningConfig resource
        """
        return pulumi.get(self, "code_signing_config_id")

    @property
    @pulumi.getter(name="codeSigningPolicies")
    def code_signing_policies(self) -> pulumi.Output[Optional['outputs.CodeSigningConfigCodeSigningPolicies']]:
        """
        Policies to control how to act if a signature is invalid
        """
        return pulumi.get(self, "code_signing_policies")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the CodeSigningConfig
        """
        return pulumi.get(self, "description")
