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

__all__ = ['InstanceGroupConfigArgs', 'InstanceGroupConfig']

@pulumi.input_type
class InstanceGroupConfigArgs:
    def __init__(__self__, *,
                 instance_count: pulumi.Input[int],
                 instance_role: pulumi.Input[str],
                 instance_type: pulumi.Input[str],
                 job_flow_id: pulumi.Input[str],
                 auto_scaling_policy: Optional[pulumi.Input['InstanceGroupConfigAutoScalingPolicyArgs']] = None,
                 bid_price: Optional[pulumi.Input[str]] = None,
                 configurations: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceGroupConfigConfigurationArgs']]]] = None,
                 custom_ami_id: Optional[pulumi.Input[str]] = None,
                 ebs_configuration: Optional[pulumi.Input['InstanceGroupConfigEbsConfigurationArgs']] = None,
                 market: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InstanceGroupConfig resource.
        """
        pulumi.set(__self__, "instance_count", instance_count)
        pulumi.set(__self__, "instance_role", instance_role)
        pulumi.set(__self__, "instance_type", instance_type)
        pulumi.set(__self__, "job_flow_id", job_flow_id)
        if auto_scaling_policy is not None:
            pulumi.set(__self__, "auto_scaling_policy", auto_scaling_policy)
        if bid_price is not None:
            pulumi.set(__self__, "bid_price", bid_price)
        if configurations is not None:
            pulumi.set(__self__, "configurations", configurations)
        if custom_ami_id is not None:
            pulumi.set(__self__, "custom_ami_id", custom_ami_id)
        if ebs_configuration is not None:
            pulumi.set(__self__, "ebs_configuration", ebs_configuration)
        if market is not None:
            pulumi.set(__self__, "market", market)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="instanceCount")
    def instance_count(self) -> pulumi.Input[int]:
        return pulumi.get(self, "instance_count")

    @instance_count.setter
    def instance_count(self, value: pulumi.Input[int]):
        pulumi.set(self, "instance_count", value)

    @property
    @pulumi.getter(name="instanceRole")
    def instance_role(self) -> pulumi.Input[str]:
        return pulumi.get(self, "instance_role")

    @instance_role.setter
    def instance_role(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_role", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter(name="jobFlowId")
    def job_flow_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "job_flow_id")

    @job_flow_id.setter
    def job_flow_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "job_flow_id", value)

    @property
    @pulumi.getter(name="autoScalingPolicy")
    def auto_scaling_policy(self) -> Optional[pulumi.Input['InstanceGroupConfigAutoScalingPolicyArgs']]:
        return pulumi.get(self, "auto_scaling_policy")

    @auto_scaling_policy.setter
    def auto_scaling_policy(self, value: Optional[pulumi.Input['InstanceGroupConfigAutoScalingPolicyArgs']]):
        pulumi.set(self, "auto_scaling_policy", value)

    @property
    @pulumi.getter(name="bidPrice")
    def bid_price(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "bid_price")

    @bid_price.setter
    def bid_price(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bid_price", value)

    @property
    @pulumi.getter
    def configurations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceGroupConfigConfigurationArgs']]]]:
        return pulumi.get(self, "configurations")

    @configurations.setter
    def configurations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceGroupConfigConfigurationArgs']]]]):
        pulumi.set(self, "configurations", value)

    @property
    @pulumi.getter(name="customAmiId")
    def custom_ami_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "custom_ami_id")

    @custom_ami_id.setter
    def custom_ami_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "custom_ami_id", value)

    @property
    @pulumi.getter(name="ebsConfiguration")
    def ebs_configuration(self) -> Optional[pulumi.Input['InstanceGroupConfigEbsConfigurationArgs']]:
        return pulumi.get(self, "ebs_configuration")

    @ebs_configuration.setter
    def ebs_configuration(self, value: Optional[pulumi.Input['InstanceGroupConfigEbsConfigurationArgs']]):
        pulumi.set(self, "ebs_configuration", value)

    @property
    @pulumi.getter
    def market(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "market")

    @market.setter
    def market(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "market", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


warnings.warn("""InstanceGroupConfig is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class InstanceGroupConfig(pulumi.CustomResource):
    warnings.warn("""InstanceGroupConfig is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_scaling_policy: Optional[pulumi.Input[pulumi.InputType['InstanceGroupConfigAutoScalingPolicyArgs']]] = None,
                 bid_price: Optional[pulumi.Input[str]] = None,
                 configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceGroupConfigConfigurationArgs']]]]] = None,
                 custom_ami_id: Optional[pulumi.Input[str]] = None,
                 ebs_configuration: Optional[pulumi.Input[pulumi.InputType['InstanceGroupConfigEbsConfigurationArgs']]] = None,
                 instance_count: Optional[pulumi.Input[int]] = None,
                 instance_role: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 job_flow_id: Optional[pulumi.Input[str]] = None,
                 market: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::EMR::InstanceGroupConfig

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceGroupConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::EMR::InstanceGroupConfig

        :param str resource_name: The name of the resource.
        :param InstanceGroupConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceGroupConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_scaling_policy: Optional[pulumi.Input[pulumi.InputType['InstanceGroupConfigAutoScalingPolicyArgs']]] = None,
                 bid_price: Optional[pulumi.Input[str]] = None,
                 configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceGroupConfigConfigurationArgs']]]]] = None,
                 custom_ami_id: Optional[pulumi.Input[str]] = None,
                 ebs_configuration: Optional[pulumi.Input[pulumi.InputType['InstanceGroupConfigEbsConfigurationArgs']]] = None,
                 instance_count: Optional[pulumi.Input[int]] = None,
                 instance_role: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 job_flow_id: Optional[pulumi.Input[str]] = None,
                 market: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""InstanceGroupConfig is deprecated: InstanceGroupConfig is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceGroupConfigArgs.__new__(InstanceGroupConfigArgs)

            __props__.__dict__["auto_scaling_policy"] = auto_scaling_policy
            __props__.__dict__["bid_price"] = bid_price
            __props__.__dict__["configurations"] = configurations
            __props__.__dict__["custom_ami_id"] = custom_ami_id
            __props__.__dict__["ebs_configuration"] = ebs_configuration
            if instance_count is None and not opts.urn:
                raise TypeError("Missing required property 'instance_count'")
            __props__.__dict__["instance_count"] = instance_count
            if instance_role is None and not opts.urn:
                raise TypeError("Missing required property 'instance_role'")
            __props__.__dict__["instance_role"] = instance_role
            if instance_type is None and not opts.urn:
                raise TypeError("Missing required property 'instance_type'")
            __props__.__dict__["instance_type"] = instance_type
            if job_flow_id is None and not opts.urn:
                raise TypeError("Missing required property 'job_flow_id'")
            __props__.__dict__["job_flow_id"] = job_flow_id
            __props__.__dict__["market"] = market
            __props__.__dict__["name"] = name
        super(InstanceGroupConfig, __self__).__init__(
            'aws-native:emr:InstanceGroupConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'InstanceGroupConfig':
        """
        Get an existing InstanceGroupConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = InstanceGroupConfigArgs.__new__(InstanceGroupConfigArgs)

        __props__.__dict__["auto_scaling_policy"] = None
        __props__.__dict__["bid_price"] = None
        __props__.__dict__["configurations"] = None
        __props__.__dict__["custom_ami_id"] = None
        __props__.__dict__["ebs_configuration"] = None
        __props__.__dict__["instance_count"] = None
        __props__.__dict__["instance_role"] = None
        __props__.__dict__["instance_type"] = None
        __props__.__dict__["job_flow_id"] = None
        __props__.__dict__["market"] = None
        __props__.__dict__["name"] = None
        return InstanceGroupConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoScalingPolicy")
    def auto_scaling_policy(self) -> pulumi.Output[Optional['outputs.InstanceGroupConfigAutoScalingPolicy']]:
        return pulumi.get(self, "auto_scaling_policy")

    @property
    @pulumi.getter(name="bidPrice")
    def bid_price(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "bid_price")

    @property
    @pulumi.getter
    def configurations(self) -> pulumi.Output[Optional[Sequence['outputs.InstanceGroupConfigConfiguration']]]:
        return pulumi.get(self, "configurations")

    @property
    @pulumi.getter(name="customAmiId")
    def custom_ami_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "custom_ami_id")

    @property
    @pulumi.getter(name="ebsConfiguration")
    def ebs_configuration(self) -> pulumi.Output[Optional['outputs.InstanceGroupConfigEbsConfiguration']]:
        return pulumi.get(self, "ebs_configuration")

    @property
    @pulumi.getter(name="instanceCount")
    def instance_count(self) -> pulumi.Output[int]:
        return pulumi.get(self, "instance_count")

    @property
    @pulumi.getter(name="instanceRole")
    def instance_role(self) -> pulumi.Output[str]:
        return pulumi.get(self, "instance_role")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="jobFlowId")
    def job_flow_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "job_flow_id")

    @property
    @pulumi.getter
    def market(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "market")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name")

