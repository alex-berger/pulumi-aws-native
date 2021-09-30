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

__all__ = ['ReplicationConfigurationArgs', 'ReplicationConfiguration']

@pulumi.input_type
class ReplicationConfigurationArgs:
    def __init__(__self__, *,
                 replication_configuration: pulumi.Input['ReplicationConfigurationReplicationConfigurationArgs']):
        """
        The set of arguments for constructing a ReplicationConfiguration resource.
        """
        pulumi.set(__self__, "replication_configuration", replication_configuration)

    @property
    @pulumi.getter(name="replicationConfiguration")
    def replication_configuration(self) -> pulumi.Input['ReplicationConfigurationReplicationConfigurationArgs']:
        return pulumi.get(self, "replication_configuration")

    @replication_configuration.setter
    def replication_configuration(self, value: pulumi.Input['ReplicationConfigurationReplicationConfigurationArgs']):
        pulumi.set(self, "replication_configuration", value)


class ReplicationConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 replication_configuration: Optional[pulumi.Input[pulumi.InputType['ReplicationConfigurationReplicationConfigurationArgs']]] = None,
                 __props__=None):
        """
        The AWS::ECR::ReplicationConfiguration resource configures the replication destinations for an Amazon Elastic Container Registry (Amazon Private ECR). For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/replication.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReplicationConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The AWS::ECR::ReplicationConfiguration resource configures the replication destinations for an Amazon Elastic Container Registry (Amazon Private ECR). For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/replication.html

        :param str resource_name: The name of the resource.
        :param ReplicationConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReplicationConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 replication_configuration: Optional[pulumi.Input[pulumi.InputType['ReplicationConfigurationReplicationConfigurationArgs']]] = None,
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
            __props__ = ReplicationConfigurationArgs.__new__(ReplicationConfigurationArgs)

            if replication_configuration is None and not opts.urn:
                raise TypeError("Missing required property 'replication_configuration'")
            __props__.__dict__["replication_configuration"] = replication_configuration
            __props__.__dict__["registry_id"] = None
        super(ReplicationConfiguration, __self__).__init__(
            'aws-native:ecr:ReplicationConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ReplicationConfiguration':
        """
        Get an existing ReplicationConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ReplicationConfigurationArgs.__new__(ReplicationConfigurationArgs)

        __props__.__dict__["registry_id"] = None
        __props__.__dict__["replication_configuration"] = None
        return ReplicationConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> pulumi.Output[str]:
        """
        The RegistryId associated with the aws account.
        """
        return pulumi.get(self, "registry_id")

    @property
    @pulumi.getter(name="replicationConfiguration")
    def replication_configuration(self) -> pulumi.Output['outputs.ReplicationConfigurationReplicationConfiguration']:
        return pulumi.get(self, "replication_configuration")

