# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetAutoScalingGroupResult',
    'AwaitableGetAutoScalingGroupResult',
    'get_auto_scaling_group',
    'get_auto_scaling_group_output',
]

@pulumi.output_type
class GetAutoScalingGroupResult:
    def __init__(__self__, availability_zones=None, capacity_rebalance=None, context=None, cooldown=None, desired_capacity=None, desired_capacity_type=None, health_check_grace_period=None, health_check_type=None, id=None, launch_configuration_name=None, launch_template=None, launch_template_specification=None, lifecycle_hook_specification_list=None, load_balancer_names=None, max_instance_lifetime=None, max_size=None, metrics_collection=None, min_size=None, mixed_instances_policy=None, new_instances_protected_from_scale_in=None, notification_configurations=None, placement_group=None, service_linked_role_arn=None, tags=None, target_group_arns=None, termination_policies=None, v_pc_zone_identifier=None):
        if availability_zones and not isinstance(availability_zones, list):
            raise TypeError("Expected argument 'availability_zones' to be a list")
        pulumi.set(__self__, "availability_zones", availability_zones)
        if capacity_rebalance and not isinstance(capacity_rebalance, bool):
            raise TypeError("Expected argument 'capacity_rebalance' to be a bool")
        pulumi.set(__self__, "capacity_rebalance", capacity_rebalance)
        if context and not isinstance(context, str):
            raise TypeError("Expected argument 'context' to be a str")
        pulumi.set(__self__, "context", context)
        if cooldown and not isinstance(cooldown, str):
            raise TypeError("Expected argument 'cooldown' to be a str")
        pulumi.set(__self__, "cooldown", cooldown)
        if desired_capacity and not isinstance(desired_capacity, str):
            raise TypeError("Expected argument 'desired_capacity' to be a str")
        pulumi.set(__self__, "desired_capacity", desired_capacity)
        if desired_capacity_type and not isinstance(desired_capacity_type, str):
            raise TypeError("Expected argument 'desired_capacity_type' to be a str")
        pulumi.set(__self__, "desired_capacity_type", desired_capacity_type)
        if health_check_grace_period and not isinstance(health_check_grace_period, int):
            raise TypeError("Expected argument 'health_check_grace_period' to be a int")
        pulumi.set(__self__, "health_check_grace_period", health_check_grace_period)
        if health_check_type and not isinstance(health_check_type, str):
            raise TypeError("Expected argument 'health_check_type' to be a str")
        pulumi.set(__self__, "health_check_type", health_check_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if launch_configuration_name and not isinstance(launch_configuration_name, str):
            raise TypeError("Expected argument 'launch_configuration_name' to be a str")
        pulumi.set(__self__, "launch_configuration_name", launch_configuration_name)
        if launch_template and not isinstance(launch_template, dict):
            raise TypeError("Expected argument 'launch_template' to be a dict")
        pulumi.set(__self__, "launch_template", launch_template)
        if launch_template_specification and not isinstance(launch_template_specification, str):
            raise TypeError("Expected argument 'launch_template_specification' to be a str")
        pulumi.set(__self__, "launch_template_specification", launch_template_specification)
        if lifecycle_hook_specification_list and not isinstance(lifecycle_hook_specification_list, list):
            raise TypeError("Expected argument 'lifecycle_hook_specification_list' to be a list")
        pulumi.set(__self__, "lifecycle_hook_specification_list", lifecycle_hook_specification_list)
        if load_balancer_names and not isinstance(load_balancer_names, list):
            raise TypeError("Expected argument 'load_balancer_names' to be a list")
        pulumi.set(__self__, "load_balancer_names", load_balancer_names)
        if max_instance_lifetime and not isinstance(max_instance_lifetime, int):
            raise TypeError("Expected argument 'max_instance_lifetime' to be a int")
        pulumi.set(__self__, "max_instance_lifetime", max_instance_lifetime)
        if max_size and not isinstance(max_size, str):
            raise TypeError("Expected argument 'max_size' to be a str")
        pulumi.set(__self__, "max_size", max_size)
        if metrics_collection and not isinstance(metrics_collection, list):
            raise TypeError("Expected argument 'metrics_collection' to be a list")
        pulumi.set(__self__, "metrics_collection", metrics_collection)
        if min_size and not isinstance(min_size, str):
            raise TypeError("Expected argument 'min_size' to be a str")
        pulumi.set(__self__, "min_size", min_size)
        if mixed_instances_policy and not isinstance(mixed_instances_policy, dict):
            raise TypeError("Expected argument 'mixed_instances_policy' to be a dict")
        pulumi.set(__self__, "mixed_instances_policy", mixed_instances_policy)
        if new_instances_protected_from_scale_in and not isinstance(new_instances_protected_from_scale_in, bool):
            raise TypeError("Expected argument 'new_instances_protected_from_scale_in' to be a bool")
        pulumi.set(__self__, "new_instances_protected_from_scale_in", new_instances_protected_from_scale_in)
        if notification_configurations and not isinstance(notification_configurations, list):
            raise TypeError("Expected argument 'notification_configurations' to be a list")
        pulumi.set(__self__, "notification_configurations", notification_configurations)
        if placement_group and not isinstance(placement_group, str):
            raise TypeError("Expected argument 'placement_group' to be a str")
        pulumi.set(__self__, "placement_group", placement_group)
        if service_linked_role_arn and not isinstance(service_linked_role_arn, str):
            raise TypeError("Expected argument 'service_linked_role_arn' to be a str")
        pulumi.set(__self__, "service_linked_role_arn", service_linked_role_arn)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if target_group_arns and not isinstance(target_group_arns, list):
            raise TypeError("Expected argument 'target_group_arns' to be a list")
        pulumi.set(__self__, "target_group_arns", target_group_arns)
        if termination_policies and not isinstance(termination_policies, list):
            raise TypeError("Expected argument 'termination_policies' to be a list")
        pulumi.set(__self__, "termination_policies", termination_policies)
        if v_pc_zone_identifier and not isinstance(v_pc_zone_identifier, list):
            raise TypeError("Expected argument 'v_pc_zone_identifier' to be a list")
        pulumi.set(__self__, "v_pc_zone_identifier", v_pc_zone_identifier)

    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "availability_zones")

    @property
    @pulumi.getter(name="capacityRebalance")
    def capacity_rebalance(self) -> Optional[bool]:
        return pulumi.get(self, "capacity_rebalance")

    @property
    @pulumi.getter
    def context(self) -> Optional[str]:
        return pulumi.get(self, "context")

    @property
    @pulumi.getter
    def cooldown(self) -> Optional[str]:
        return pulumi.get(self, "cooldown")

    @property
    @pulumi.getter(name="desiredCapacity")
    def desired_capacity(self) -> Optional[str]:
        return pulumi.get(self, "desired_capacity")

    @property
    @pulumi.getter(name="desiredCapacityType")
    def desired_capacity_type(self) -> Optional[str]:
        return pulumi.get(self, "desired_capacity_type")

    @property
    @pulumi.getter(name="healthCheckGracePeriod")
    def health_check_grace_period(self) -> Optional[int]:
        return pulumi.get(self, "health_check_grace_period")

    @property
    @pulumi.getter(name="healthCheckType")
    def health_check_type(self) -> Optional[str]:
        return pulumi.get(self, "health_check_type")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="launchConfigurationName")
    def launch_configuration_name(self) -> Optional[str]:
        return pulumi.get(self, "launch_configuration_name")

    @property
    @pulumi.getter(name="launchTemplate")
    def launch_template(self) -> Optional['outputs.AutoScalingGroupLaunchTemplateSpecification']:
        return pulumi.get(self, "launch_template")

    @property
    @pulumi.getter(name="launchTemplateSpecification")
    def launch_template_specification(self) -> Optional[str]:
        return pulumi.get(self, "launch_template_specification")

    @property
    @pulumi.getter(name="lifecycleHookSpecificationList")
    def lifecycle_hook_specification_list(self) -> Optional[Sequence['outputs.AutoScalingGroupLifecycleHookSpecification']]:
        return pulumi.get(self, "lifecycle_hook_specification_list")

    @property
    @pulumi.getter(name="loadBalancerNames")
    def load_balancer_names(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "load_balancer_names")

    @property
    @pulumi.getter(name="maxInstanceLifetime")
    def max_instance_lifetime(self) -> Optional[int]:
        return pulumi.get(self, "max_instance_lifetime")

    @property
    @pulumi.getter(name="maxSize")
    def max_size(self) -> Optional[str]:
        return pulumi.get(self, "max_size")

    @property
    @pulumi.getter(name="metricsCollection")
    def metrics_collection(self) -> Optional[Sequence['outputs.AutoScalingGroupMetricsCollection']]:
        return pulumi.get(self, "metrics_collection")

    @property
    @pulumi.getter(name="minSize")
    def min_size(self) -> Optional[str]:
        return pulumi.get(self, "min_size")

    @property
    @pulumi.getter(name="mixedInstancesPolicy")
    def mixed_instances_policy(self) -> Optional['outputs.AutoScalingGroupMixedInstancesPolicy']:
        return pulumi.get(self, "mixed_instances_policy")

    @property
    @pulumi.getter(name="newInstancesProtectedFromScaleIn")
    def new_instances_protected_from_scale_in(self) -> Optional[bool]:
        return pulumi.get(self, "new_instances_protected_from_scale_in")

    @property
    @pulumi.getter(name="notificationConfigurations")
    def notification_configurations(self) -> Optional[Sequence['outputs.AutoScalingGroupNotificationConfiguration']]:
        return pulumi.get(self, "notification_configurations")

    @property
    @pulumi.getter(name="placementGroup")
    def placement_group(self) -> Optional[str]:
        return pulumi.get(self, "placement_group")

    @property
    @pulumi.getter(name="serviceLinkedRoleARN")
    def service_linked_role_arn(self) -> Optional[str]:
        return pulumi.get(self, "service_linked_role_arn")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.AutoScalingGroupTagProperty']]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetGroupARNs")
    def target_group_arns(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "target_group_arns")

    @property
    @pulumi.getter(name="terminationPolicies")
    def termination_policies(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "termination_policies")

    @property
    @pulumi.getter(name="vPCZoneIdentifier")
    def v_pc_zone_identifier(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "v_pc_zone_identifier")


class AwaitableGetAutoScalingGroupResult(GetAutoScalingGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAutoScalingGroupResult(
            availability_zones=self.availability_zones,
            capacity_rebalance=self.capacity_rebalance,
            context=self.context,
            cooldown=self.cooldown,
            desired_capacity=self.desired_capacity,
            desired_capacity_type=self.desired_capacity_type,
            health_check_grace_period=self.health_check_grace_period,
            health_check_type=self.health_check_type,
            id=self.id,
            launch_configuration_name=self.launch_configuration_name,
            launch_template=self.launch_template,
            launch_template_specification=self.launch_template_specification,
            lifecycle_hook_specification_list=self.lifecycle_hook_specification_list,
            load_balancer_names=self.load_balancer_names,
            max_instance_lifetime=self.max_instance_lifetime,
            max_size=self.max_size,
            metrics_collection=self.metrics_collection,
            min_size=self.min_size,
            mixed_instances_policy=self.mixed_instances_policy,
            new_instances_protected_from_scale_in=self.new_instances_protected_from_scale_in,
            notification_configurations=self.notification_configurations,
            placement_group=self.placement_group,
            service_linked_role_arn=self.service_linked_role_arn,
            tags=self.tags,
            target_group_arns=self.target_group_arns,
            termination_policies=self.termination_policies,
            v_pc_zone_identifier=self.v_pc_zone_identifier)


def get_auto_scaling_group(id: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAutoScalingGroupResult:
    """
    Resource Type definition for AWS::AutoScaling::AutoScalingGroup
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:autoscaling:getAutoScalingGroup', __args__, opts=opts, typ=GetAutoScalingGroupResult).value

    return AwaitableGetAutoScalingGroupResult(
        availability_zones=__ret__.availability_zones,
        capacity_rebalance=__ret__.capacity_rebalance,
        context=__ret__.context,
        cooldown=__ret__.cooldown,
        desired_capacity=__ret__.desired_capacity,
        desired_capacity_type=__ret__.desired_capacity_type,
        health_check_grace_period=__ret__.health_check_grace_period,
        health_check_type=__ret__.health_check_type,
        id=__ret__.id,
        launch_configuration_name=__ret__.launch_configuration_name,
        launch_template=__ret__.launch_template,
        launch_template_specification=__ret__.launch_template_specification,
        lifecycle_hook_specification_list=__ret__.lifecycle_hook_specification_list,
        load_balancer_names=__ret__.load_balancer_names,
        max_instance_lifetime=__ret__.max_instance_lifetime,
        max_size=__ret__.max_size,
        metrics_collection=__ret__.metrics_collection,
        min_size=__ret__.min_size,
        mixed_instances_policy=__ret__.mixed_instances_policy,
        new_instances_protected_from_scale_in=__ret__.new_instances_protected_from_scale_in,
        notification_configurations=__ret__.notification_configurations,
        placement_group=__ret__.placement_group,
        service_linked_role_arn=__ret__.service_linked_role_arn,
        tags=__ret__.tags,
        target_group_arns=__ret__.target_group_arns,
        termination_policies=__ret__.termination_policies,
        v_pc_zone_identifier=__ret__.v_pc_zone_identifier)


@_utilities.lift_output_func(get_auto_scaling_group)
def get_auto_scaling_group_output(id: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAutoScalingGroupResult]:
    """
    Resource Type definition for AWS::AutoScaling::AutoScalingGroup
    """
    ...
