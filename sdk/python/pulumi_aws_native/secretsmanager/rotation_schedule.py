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

__all__ = ['RotationScheduleArgs', 'RotationSchedule']

@pulumi.input_type
class RotationScheduleArgs:
    def __init__(__self__, *,
                 secret_id: pulumi.Input[str],
                 hosted_rotation_lambda: Optional[pulumi.Input['RotationScheduleHostedRotationLambdaArgs']] = None,
                 rotate_immediately_on_update: Optional[pulumi.Input[bool]] = None,
                 rotation_lambda_arn: Optional[pulumi.Input[str]] = None,
                 rotation_rules: Optional[pulumi.Input['RotationScheduleRotationRulesArgs']] = None):
        """
        The set of arguments for constructing a RotationSchedule resource.
        """
        pulumi.set(__self__, "secret_id", secret_id)
        if hosted_rotation_lambda is not None:
            pulumi.set(__self__, "hosted_rotation_lambda", hosted_rotation_lambda)
        if rotate_immediately_on_update is not None:
            pulumi.set(__self__, "rotate_immediately_on_update", rotate_immediately_on_update)
        if rotation_lambda_arn is not None:
            pulumi.set(__self__, "rotation_lambda_arn", rotation_lambda_arn)
        if rotation_rules is not None:
            pulumi.set(__self__, "rotation_rules", rotation_rules)

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "secret_id")

    @secret_id.setter
    def secret_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret_id", value)

    @property
    @pulumi.getter(name="hostedRotationLambda")
    def hosted_rotation_lambda(self) -> Optional[pulumi.Input['RotationScheduleHostedRotationLambdaArgs']]:
        return pulumi.get(self, "hosted_rotation_lambda")

    @hosted_rotation_lambda.setter
    def hosted_rotation_lambda(self, value: Optional[pulumi.Input['RotationScheduleHostedRotationLambdaArgs']]):
        pulumi.set(self, "hosted_rotation_lambda", value)

    @property
    @pulumi.getter(name="rotateImmediatelyOnUpdate")
    def rotate_immediately_on_update(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "rotate_immediately_on_update")

    @rotate_immediately_on_update.setter
    def rotate_immediately_on_update(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "rotate_immediately_on_update", value)

    @property
    @pulumi.getter(name="rotationLambdaARN")
    def rotation_lambda_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "rotation_lambda_arn")

    @rotation_lambda_arn.setter
    def rotation_lambda_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rotation_lambda_arn", value)

    @property
    @pulumi.getter(name="rotationRules")
    def rotation_rules(self) -> Optional[pulumi.Input['RotationScheduleRotationRulesArgs']]:
        return pulumi.get(self, "rotation_rules")

    @rotation_rules.setter
    def rotation_rules(self, value: Optional[pulumi.Input['RotationScheduleRotationRulesArgs']]):
        pulumi.set(self, "rotation_rules", value)


warnings.warn("""RotationSchedule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class RotationSchedule(pulumi.CustomResource):
    warnings.warn("""RotationSchedule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 hosted_rotation_lambda: Optional[pulumi.Input[pulumi.InputType['RotationScheduleHostedRotationLambdaArgs']]] = None,
                 rotate_immediately_on_update: Optional[pulumi.Input[bool]] = None,
                 rotation_lambda_arn: Optional[pulumi.Input[str]] = None,
                 rotation_rules: Optional[pulumi.Input[pulumi.InputType['RotationScheduleRotationRulesArgs']]] = None,
                 secret_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::SecretsManager::RotationSchedule

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RotationScheduleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::SecretsManager::RotationSchedule

        :param str resource_name: The name of the resource.
        :param RotationScheduleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RotationScheduleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 hosted_rotation_lambda: Optional[pulumi.Input[pulumi.InputType['RotationScheduleHostedRotationLambdaArgs']]] = None,
                 rotate_immediately_on_update: Optional[pulumi.Input[bool]] = None,
                 rotation_lambda_arn: Optional[pulumi.Input[str]] = None,
                 rotation_rules: Optional[pulumi.Input[pulumi.InputType['RotationScheduleRotationRulesArgs']]] = None,
                 secret_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""RotationSchedule is deprecated: RotationSchedule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RotationScheduleArgs.__new__(RotationScheduleArgs)

            __props__.__dict__["hosted_rotation_lambda"] = hosted_rotation_lambda
            __props__.__dict__["rotate_immediately_on_update"] = rotate_immediately_on_update
            __props__.__dict__["rotation_lambda_arn"] = rotation_lambda_arn
            __props__.__dict__["rotation_rules"] = rotation_rules
            if secret_id is None and not opts.urn:
                raise TypeError("Missing required property 'secret_id'")
            __props__.__dict__["secret_id"] = secret_id
        super(RotationSchedule, __self__).__init__(
            'aws-native:secretsmanager:RotationSchedule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RotationSchedule':
        """
        Get an existing RotationSchedule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RotationScheduleArgs.__new__(RotationScheduleArgs)

        __props__.__dict__["hosted_rotation_lambda"] = None
        __props__.__dict__["rotate_immediately_on_update"] = None
        __props__.__dict__["rotation_lambda_arn"] = None
        __props__.__dict__["rotation_rules"] = None
        __props__.__dict__["secret_id"] = None
        return RotationSchedule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="hostedRotationLambda")
    def hosted_rotation_lambda(self) -> pulumi.Output[Optional['outputs.RotationScheduleHostedRotationLambda']]:
        return pulumi.get(self, "hosted_rotation_lambda")

    @property
    @pulumi.getter(name="rotateImmediatelyOnUpdate")
    def rotate_immediately_on_update(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "rotate_immediately_on_update")

    @property
    @pulumi.getter(name="rotationLambdaARN")
    def rotation_lambda_arn(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "rotation_lambda_arn")

    @property
    @pulumi.getter(name="rotationRules")
    def rotation_rules(self) -> pulumi.Output[Optional['outputs.RotationScheduleRotationRules']]:
        return pulumi.get(self, "rotation_rules")

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "secret_id")

