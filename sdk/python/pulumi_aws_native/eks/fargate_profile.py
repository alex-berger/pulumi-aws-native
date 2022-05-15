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

__all__ = ['FargateProfileArgs', 'FargateProfile']

@pulumi.input_type
class FargateProfileArgs:
    def __init__(__self__, *,
                 cluster_name: pulumi.Input[str],
                 pod_execution_role_arn: pulumi.Input[str],
                 selectors: pulumi.Input[Sequence[pulumi.Input['FargateProfileSelectorArgs']]],
                 fargate_profile_name: Optional[pulumi.Input[str]] = None,
                 subnets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['FargateProfileTagArgs']]]] = None):
        """
        The set of arguments for constructing a FargateProfile resource.
        :param pulumi.Input[str] cluster_name: Name of the Cluster
        :param pulumi.Input[str] pod_execution_role_arn: The IAM policy arn for pods
        :param pulumi.Input[str] fargate_profile_name: Name of FargateProfile
        :param pulumi.Input[Sequence[pulumi.Input['FargateProfileTagArgs']]] tags: An array of key-value pairs to apply to this resource.
        """
        pulumi.set(__self__, "cluster_name", cluster_name)
        pulumi.set(__self__, "pod_execution_role_arn", pod_execution_role_arn)
        pulumi.set(__self__, "selectors", selectors)
        if fargate_profile_name is not None:
            pulumi.set(__self__, "fargate_profile_name", fargate_profile_name)
        if subnets is not None:
            pulumi.set(__self__, "subnets", subnets)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Input[str]:
        """
        Name of the Cluster
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter(name="podExecutionRoleArn")
    def pod_execution_role_arn(self) -> pulumi.Input[str]:
        """
        The IAM policy arn for pods
        """
        return pulumi.get(self, "pod_execution_role_arn")

    @pod_execution_role_arn.setter
    def pod_execution_role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "pod_execution_role_arn", value)

    @property
    @pulumi.getter
    def selectors(self) -> pulumi.Input[Sequence[pulumi.Input['FargateProfileSelectorArgs']]]:
        return pulumi.get(self, "selectors")

    @selectors.setter
    def selectors(self, value: pulumi.Input[Sequence[pulumi.Input['FargateProfileSelectorArgs']]]):
        pulumi.set(self, "selectors", value)

    @property
    @pulumi.getter(name="fargateProfileName")
    def fargate_profile_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of FargateProfile
        """
        return pulumi.get(self, "fargate_profile_name")

    @fargate_profile_name.setter
    def fargate_profile_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fargate_profile_name", value)

    @property
    @pulumi.getter
    def subnets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "subnets")

    @subnets.setter
    def subnets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "subnets", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FargateProfileTagArgs']]]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FargateProfileTagArgs']]]]):
        pulumi.set(self, "tags", value)


class FargateProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 fargate_profile_name: Optional[pulumi.Input[str]] = None,
                 pod_execution_role_arn: Optional[pulumi.Input[str]] = None,
                 selectors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FargateProfileSelectorArgs']]]]] = None,
                 subnets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FargateProfileTagArgs']]]]] = None,
                 __props__=None):
        """
        Resource Schema for AWS::EKS::FargateProfile

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_name: Name of the Cluster
        :param pulumi.Input[str] fargate_profile_name: Name of FargateProfile
        :param pulumi.Input[str] pod_execution_role_arn: The IAM policy arn for pods
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FargateProfileTagArgs']]]] tags: An array of key-value pairs to apply to this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FargateProfileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Schema for AWS::EKS::FargateProfile

        :param str resource_name: The name of the resource.
        :param FargateProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FargateProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 fargate_profile_name: Optional[pulumi.Input[str]] = None,
                 pod_execution_role_arn: Optional[pulumi.Input[str]] = None,
                 selectors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FargateProfileSelectorArgs']]]]] = None,
                 subnets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FargateProfileTagArgs']]]]] = None,
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
            __props__ = FargateProfileArgs.__new__(FargateProfileArgs)

            if cluster_name is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_name'")
            __props__.__dict__["cluster_name"] = cluster_name
            __props__.__dict__["fargate_profile_name"] = fargate_profile_name
            if pod_execution_role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'pod_execution_role_arn'")
            __props__.__dict__["pod_execution_role_arn"] = pod_execution_role_arn
            if selectors is None and not opts.urn:
                raise TypeError("Missing required property 'selectors'")
            __props__.__dict__["selectors"] = selectors
            __props__.__dict__["subnets"] = subnets
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
        super(FargateProfile, __self__).__init__(
            'aws-native:eks:FargateProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'FargateProfile':
        """
        Get an existing FargateProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FargateProfileArgs.__new__(FargateProfileArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["cluster_name"] = None
        __props__.__dict__["fargate_profile_name"] = None
        __props__.__dict__["pod_execution_role_arn"] = None
        __props__.__dict__["selectors"] = None
        __props__.__dict__["subnets"] = None
        __props__.__dict__["tags"] = None
        return FargateProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Output[str]:
        """
        Name of the Cluster
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="fargateProfileName")
    def fargate_profile_name(self) -> pulumi.Output[Optional[str]]:
        """
        Name of FargateProfile
        """
        return pulumi.get(self, "fargate_profile_name")

    @property
    @pulumi.getter(name="podExecutionRoleArn")
    def pod_execution_role_arn(self) -> pulumi.Output[str]:
        """
        The IAM policy arn for pods
        """
        return pulumi.get(self, "pod_execution_role_arn")

    @property
    @pulumi.getter
    def selectors(self) -> pulumi.Output[Sequence['outputs.FargateProfileSelector']]:
        return pulumi.get(self, "selectors")

    @property
    @pulumi.getter
    def subnets(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "subnets")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.FargateProfileTag']]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

