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

__all__ = ['ServiceArgs', 'Service']

@pulumi.input_type
class ServiceArgs:
    def __init__(__self__, *,
                 capacity_provider_strategy: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceCapacityProviderStrategyItemArgs']]]] = None,
                 cluster: Optional[pulumi.Input[str]] = None,
                 deployment_configuration: Optional[pulumi.Input['ServiceDeploymentConfigurationArgs']] = None,
                 deployment_controller: Optional[pulumi.Input['ServiceDeploymentControllerArgs']] = None,
                 desired_count: Optional[pulumi.Input[int]] = None,
                 enable_ecs_managed_tags: Optional[pulumi.Input[bool]] = None,
                 enable_execute_command: Optional[pulumi.Input[bool]] = None,
                 health_check_grace_period_seconds: Optional[pulumi.Input[int]] = None,
                 launch_type: Optional[pulumi.Input['ServiceLaunchType']] = None,
                 load_balancers: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceLoadBalancerArgs']]]] = None,
                 network_configuration: Optional[pulumi.Input['ServiceNetworkConfigurationArgs']] = None,
                 placement_constraints: Optional[pulumi.Input[Sequence[pulumi.Input['ServicePlacementConstraintArgs']]]] = None,
                 placement_strategies: Optional[pulumi.Input[Sequence[pulumi.Input['ServicePlacementStrategyArgs']]]] = None,
                 platform_version: Optional[pulumi.Input[str]] = None,
                 propagate_tags: Optional[pulumi.Input['ServicePropagateTags']] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 scheduling_strategy: Optional[pulumi.Input['ServiceSchedulingStrategy']] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 service_registries: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceRegistryArgs']]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceTagArgs']]]] = None,
                 task_definition: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Service resource.
        """
        if capacity_provider_strategy is not None:
            pulumi.set(__self__, "capacity_provider_strategy", capacity_provider_strategy)
        if cluster is not None:
            pulumi.set(__self__, "cluster", cluster)
        if deployment_configuration is not None:
            pulumi.set(__self__, "deployment_configuration", deployment_configuration)
        if deployment_controller is not None:
            pulumi.set(__self__, "deployment_controller", deployment_controller)
        if desired_count is not None:
            pulumi.set(__self__, "desired_count", desired_count)
        if enable_ecs_managed_tags is not None:
            pulumi.set(__self__, "enable_ecs_managed_tags", enable_ecs_managed_tags)
        if enable_execute_command is not None:
            pulumi.set(__self__, "enable_execute_command", enable_execute_command)
        if health_check_grace_period_seconds is not None:
            pulumi.set(__self__, "health_check_grace_period_seconds", health_check_grace_period_seconds)
        if launch_type is not None:
            pulumi.set(__self__, "launch_type", launch_type)
        if load_balancers is not None:
            pulumi.set(__self__, "load_balancers", load_balancers)
        if network_configuration is not None:
            pulumi.set(__self__, "network_configuration", network_configuration)
        if placement_constraints is not None:
            pulumi.set(__self__, "placement_constraints", placement_constraints)
        if placement_strategies is not None:
            pulumi.set(__self__, "placement_strategies", placement_strategies)
        if platform_version is not None:
            pulumi.set(__self__, "platform_version", platform_version)
        if propagate_tags is not None:
            pulumi.set(__self__, "propagate_tags", propagate_tags)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if scheduling_strategy is not None:
            pulumi.set(__self__, "scheduling_strategy", scheduling_strategy)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if service_registries is not None:
            pulumi.set(__self__, "service_registries", service_registries)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if task_definition is not None:
            pulumi.set(__self__, "task_definition", task_definition)

    @property
    @pulumi.getter(name="capacityProviderStrategy")
    def capacity_provider_strategy(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServiceCapacityProviderStrategyItemArgs']]]]:
        return pulumi.get(self, "capacity_provider_strategy")

    @capacity_provider_strategy.setter
    def capacity_provider_strategy(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceCapacityProviderStrategyItemArgs']]]]):
        pulumi.set(self, "capacity_provider_strategy", value)

    @property
    @pulumi.getter
    def cluster(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cluster")

    @cluster.setter
    def cluster(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster", value)

    @property
    @pulumi.getter(name="deploymentConfiguration")
    def deployment_configuration(self) -> Optional[pulumi.Input['ServiceDeploymentConfigurationArgs']]:
        return pulumi.get(self, "deployment_configuration")

    @deployment_configuration.setter
    def deployment_configuration(self, value: Optional[pulumi.Input['ServiceDeploymentConfigurationArgs']]):
        pulumi.set(self, "deployment_configuration", value)

    @property
    @pulumi.getter(name="deploymentController")
    def deployment_controller(self) -> Optional[pulumi.Input['ServiceDeploymentControllerArgs']]:
        return pulumi.get(self, "deployment_controller")

    @deployment_controller.setter
    def deployment_controller(self, value: Optional[pulumi.Input['ServiceDeploymentControllerArgs']]):
        pulumi.set(self, "deployment_controller", value)

    @property
    @pulumi.getter(name="desiredCount")
    def desired_count(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "desired_count")

    @desired_count.setter
    def desired_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "desired_count", value)

    @property
    @pulumi.getter(name="enableECSManagedTags")
    def enable_ecs_managed_tags(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enable_ecs_managed_tags")

    @enable_ecs_managed_tags.setter
    def enable_ecs_managed_tags(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_ecs_managed_tags", value)

    @property
    @pulumi.getter(name="enableExecuteCommand")
    def enable_execute_command(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enable_execute_command")

    @enable_execute_command.setter
    def enable_execute_command(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_execute_command", value)

    @property
    @pulumi.getter(name="healthCheckGracePeriodSeconds")
    def health_check_grace_period_seconds(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "health_check_grace_period_seconds")

    @health_check_grace_period_seconds.setter
    def health_check_grace_period_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "health_check_grace_period_seconds", value)

    @property
    @pulumi.getter(name="launchType")
    def launch_type(self) -> Optional[pulumi.Input['ServiceLaunchType']]:
        return pulumi.get(self, "launch_type")

    @launch_type.setter
    def launch_type(self, value: Optional[pulumi.Input['ServiceLaunchType']]):
        pulumi.set(self, "launch_type", value)

    @property
    @pulumi.getter(name="loadBalancers")
    def load_balancers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServiceLoadBalancerArgs']]]]:
        return pulumi.get(self, "load_balancers")

    @load_balancers.setter
    def load_balancers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceLoadBalancerArgs']]]]):
        pulumi.set(self, "load_balancers", value)

    @property
    @pulumi.getter(name="networkConfiguration")
    def network_configuration(self) -> Optional[pulumi.Input['ServiceNetworkConfigurationArgs']]:
        return pulumi.get(self, "network_configuration")

    @network_configuration.setter
    def network_configuration(self, value: Optional[pulumi.Input['ServiceNetworkConfigurationArgs']]):
        pulumi.set(self, "network_configuration", value)

    @property
    @pulumi.getter(name="placementConstraints")
    def placement_constraints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServicePlacementConstraintArgs']]]]:
        return pulumi.get(self, "placement_constraints")

    @placement_constraints.setter
    def placement_constraints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServicePlacementConstraintArgs']]]]):
        pulumi.set(self, "placement_constraints", value)

    @property
    @pulumi.getter(name="placementStrategies")
    def placement_strategies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServicePlacementStrategyArgs']]]]:
        return pulumi.get(self, "placement_strategies")

    @placement_strategies.setter
    def placement_strategies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServicePlacementStrategyArgs']]]]):
        pulumi.set(self, "placement_strategies", value)

    @property
    @pulumi.getter(name="platformVersion")
    def platform_version(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "platform_version")

    @platform_version.setter
    def platform_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "platform_version", value)

    @property
    @pulumi.getter(name="propagateTags")
    def propagate_tags(self) -> Optional[pulumi.Input['ServicePropagateTags']]:
        return pulumi.get(self, "propagate_tags")

    @propagate_tags.setter
    def propagate_tags(self, value: Optional[pulumi.Input['ServicePropagateTags']]):
        pulumi.set(self, "propagate_tags", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="schedulingStrategy")
    def scheduling_strategy(self) -> Optional[pulumi.Input['ServiceSchedulingStrategy']]:
        return pulumi.get(self, "scheduling_strategy")

    @scheduling_strategy.setter
    def scheduling_strategy(self, value: Optional[pulumi.Input['ServiceSchedulingStrategy']]):
        pulumi.set(self, "scheduling_strategy", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="serviceRegistries")
    def service_registries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServiceRegistryArgs']]]]:
        return pulumi.get(self, "service_registries")

    @service_registries.setter
    def service_registries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceRegistryArgs']]]]):
        pulumi.set(self, "service_registries", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServiceTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="taskDefinition")
    def task_definition(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "task_definition")

    @task_definition.setter
    def task_definition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "task_definition", value)


class Service(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capacity_provider_strategy: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceCapacityProviderStrategyItemArgs']]]]] = None,
                 cluster: Optional[pulumi.Input[str]] = None,
                 deployment_configuration: Optional[pulumi.Input[pulumi.InputType['ServiceDeploymentConfigurationArgs']]] = None,
                 deployment_controller: Optional[pulumi.Input[pulumi.InputType['ServiceDeploymentControllerArgs']]] = None,
                 desired_count: Optional[pulumi.Input[int]] = None,
                 enable_ecs_managed_tags: Optional[pulumi.Input[bool]] = None,
                 enable_execute_command: Optional[pulumi.Input[bool]] = None,
                 health_check_grace_period_seconds: Optional[pulumi.Input[int]] = None,
                 launch_type: Optional[pulumi.Input['ServiceLaunchType']] = None,
                 load_balancers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceLoadBalancerArgs']]]]] = None,
                 network_configuration: Optional[pulumi.Input[pulumi.InputType['ServiceNetworkConfigurationArgs']]] = None,
                 placement_constraints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServicePlacementConstraintArgs']]]]] = None,
                 placement_strategies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServicePlacementStrategyArgs']]]]] = None,
                 platform_version: Optional[pulumi.Input[str]] = None,
                 propagate_tags: Optional[pulumi.Input['ServicePropagateTags']] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 scheduling_strategy: Optional[pulumi.Input['ServiceSchedulingStrategy']] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 service_registries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceRegistryArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceTagArgs']]]]] = None,
                 task_definition: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::ECS::Service

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ServiceArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::ECS::Service

        :param str resource_name: The name of the resource.
        :param ServiceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capacity_provider_strategy: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceCapacityProviderStrategyItemArgs']]]]] = None,
                 cluster: Optional[pulumi.Input[str]] = None,
                 deployment_configuration: Optional[pulumi.Input[pulumi.InputType['ServiceDeploymentConfigurationArgs']]] = None,
                 deployment_controller: Optional[pulumi.Input[pulumi.InputType['ServiceDeploymentControllerArgs']]] = None,
                 desired_count: Optional[pulumi.Input[int]] = None,
                 enable_ecs_managed_tags: Optional[pulumi.Input[bool]] = None,
                 enable_execute_command: Optional[pulumi.Input[bool]] = None,
                 health_check_grace_period_seconds: Optional[pulumi.Input[int]] = None,
                 launch_type: Optional[pulumi.Input['ServiceLaunchType']] = None,
                 load_balancers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceLoadBalancerArgs']]]]] = None,
                 network_configuration: Optional[pulumi.Input[pulumi.InputType['ServiceNetworkConfigurationArgs']]] = None,
                 placement_constraints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServicePlacementConstraintArgs']]]]] = None,
                 placement_strategies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServicePlacementStrategyArgs']]]]] = None,
                 platform_version: Optional[pulumi.Input[str]] = None,
                 propagate_tags: Optional[pulumi.Input['ServicePropagateTags']] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 scheduling_strategy: Optional[pulumi.Input['ServiceSchedulingStrategy']] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 service_registries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceRegistryArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceTagArgs']]]]] = None,
                 task_definition: Optional[pulumi.Input[str]] = None,
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
            __props__ = ServiceArgs.__new__(ServiceArgs)

            __props__.__dict__["capacity_provider_strategy"] = capacity_provider_strategy
            __props__.__dict__["cluster"] = cluster
            __props__.__dict__["deployment_configuration"] = deployment_configuration
            __props__.__dict__["deployment_controller"] = deployment_controller
            __props__.__dict__["desired_count"] = desired_count
            __props__.__dict__["enable_ecs_managed_tags"] = enable_ecs_managed_tags
            __props__.__dict__["enable_execute_command"] = enable_execute_command
            __props__.__dict__["health_check_grace_period_seconds"] = health_check_grace_period_seconds
            __props__.__dict__["launch_type"] = launch_type
            __props__.__dict__["load_balancers"] = load_balancers
            __props__.__dict__["network_configuration"] = network_configuration
            __props__.__dict__["placement_constraints"] = placement_constraints
            __props__.__dict__["placement_strategies"] = placement_strategies
            __props__.__dict__["platform_version"] = platform_version
            __props__.__dict__["propagate_tags"] = propagate_tags
            __props__.__dict__["role"] = role
            __props__.__dict__["scheduling_strategy"] = scheduling_strategy
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["service_registries"] = service_registries
            __props__.__dict__["tags"] = tags
            __props__.__dict__["task_definition"] = task_definition
            __props__.__dict__["name"] = None
            __props__.__dict__["service_arn"] = None
        super(Service, __self__).__init__(
            'aws-native:ecs:Service',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Service':
        """
        Get an existing Service resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ServiceArgs.__new__(ServiceArgs)

        __props__.__dict__["capacity_provider_strategy"] = None
        __props__.__dict__["cluster"] = None
        __props__.__dict__["deployment_configuration"] = None
        __props__.__dict__["deployment_controller"] = None
        __props__.__dict__["desired_count"] = None
        __props__.__dict__["enable_ecs_managed_tags"] = None
        __props__.__dict__["enable_execute_command"] = None
        __props__.__dict__["health_check_grace_period_seconds"] = None
        __props__.__dict__["launch_type"] = None
        __props__.__dict__["load_balancers"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["network_configuration"] = None
        __props__.__dict__["placement_constraints"] = None
        __props__.__dict__["placement_strategies"] = None
        __props__.__dict__["platform_version"] = None
        __props__.__dict__["propagate_tags"] = None
        __props__.__dict__["role"] = None
        __props__.__dict__["scheduling_strategy"] = None
        __props__.__dict__["service_arn"] = None
        __props__.__dict__["service_name"] = None
        __props__.__dict__["service_registries"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["task_definition"] = None
        return Service(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="capacityProviderStrategy")
    def capacity_provider_strategy(self) -> pulumi.Output[Optional[Sequence['outputs.ServiceCapacityProviderStrategyItem']]]:
        return pulumi.get(self, "capacity_provider_strategy")

    @property
    @pulumi.getter
    def cluster(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "cluster")

    @property
    @pulumi.getter(name="deploymentConfiguration")
    def deployment_configuration(self) -> pulumi.Output[Optional['outputs.ServiceDeploymentConfiguration']]:
        return pulumi.get(self, "deployment_configuration")

    @property
    @pulumi.getter(name="deploymentController")
    def deployment_controller(self) -> pulumi.Output[Optional['outputs.ServiceDeploymentController']]:
        return pulumi.get(self, "deployment_controller")

    @property
    @pulumi.getter(name="desiredCount")
    def desired_count(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "desired_count")

    @property
    @pulumi.getter(name="enableECSManagedTags")
    def enable_ecs_managed_tags(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enable_ecs_managed_tags")

    @property
    @pulumi.getter(name="enableExecuteCommand")
    def enable_execute_command(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enable_execute_command")

    @property
    @pulumi.getter(name="healthCheckGracePeriodSeconds")
    def health_check_grace_period_seconds(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "health_check_grace_period_seconds")

    @property
    @pulumi.getter(name="launchType")
    def launch_type(self) -> pulumi.Output[Optional['ServiceLaunchType']]:
        return pulumi.get(self, "launch_type")

    @property
    @pulumi.getter(name="loadBalancers")
    def load_balancers(self) -> pulumi.Output[Optional[Sequence['outputs.ServiceLoadBalancer']]]:
        return pulumi.get(self, "load_balancers")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkConfiguration")
    def network_configuration(self) -> pulumi.Output[Optional['outputs.ServiceNetworkConfiguration']]:
        return pulumi.get(self, "network_configuration")

    @property
    @pulumi.getter(name="placementConstraints")
    def placement_constraints(self) -> pulumi.Output[Optional[Sequence['outputs.ServicePlacementConstraint']]]:
        return pulumi.get(self, "placement_constraints")

    @property
    @pulumi.getter(name="placementStrategies")
    def placement_strategies(self) -> pulumi.Output[Optional[Sequence['outputs.ServicePlacementStrategy']]]:
        return pulumi.get(self, "placement_strategies")

    @property
    @pulumi.getter(name="platformVersion")
    def platform_version(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "platform_version")

    @property
    @pulumi.getter(name="propagateTags")
    def propagate_tags(self) -> pulumi.Output[Optional['ServicePropagateTags']]:
        return pulumi.get(self, "propagate_tags")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="schedulingStrategy")
    def scheduling_strategy(self) -> pulumi.Output[Optional['ServiceSchedulingStrategy']]:
        return pulumi.get(self, "scheduling_strategy")

    @property
    @pulumi.getter(name="serviceArn")
    def service_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "service_arn")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="serviceRegistries")
    def service_registries(self) -> pulumi.Output[Optional[Sequence['outputs.ServiceRegistry']]]:
        return pulumi.get(self, "service_registries")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.ServiceTag']]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="taskDefinition")
    def task_definition(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "task_definition")

