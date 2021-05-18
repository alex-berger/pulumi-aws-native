# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['TaskSetArgs', 'TaskSet']

@pulumi.input_type
class TaskSetArgs:
    def __init__(__self__, *,
                 cluster: pulumi.Input[str],
                 service: pulumi.Input[str],
                 task_definition: pulumi.Input[str],
                 external_id: Optional[pulumi.Input[str]] = None,
                 launch_type: Optional[pulumi.Input[str]] = None,
                 load_balancers: Optional[pulumi.Input[Sequence[pulumi.Input['TaskSetLoadBalancerArgs']]]] = None,
                 network_configuration: Optional[pulumi.Input['TaskSetNetworkConfigurationArgs']] = None,
                 platform_version: Optional[pulumi.Input[str]] = None,
                 scale: Optional[pulumi.Input['TaskSetScaleArgs']] = None,
                 service_registries: Optional[pulumi.Input[Sequence[pulumi.Input['TaskSetServiceRegistryArgs']]]] = None):
        """
        The set of arguments for constructing a TaskSet resource.
        :param pulumi.Input[str] cluster: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-cluster
        :param pulumi.Input[str] service: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-service
        :param pulumi.Input[str] task_definition: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-taskdefinition
        :param pulumi.Input[str] external_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-externalid
        :param pulumi.Input[str] launch_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-launchtype
        :param pulumi.Input[Sequence[pulumi.Input['TaskSetLoadBalancerArgs']]] load_balancers: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-loadbalancers
        :param pulumi.Input['TaskSetNetworkConfigurationArgs'] network_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-networkconfiguration
        :param pulumi.Input[str] platform_version: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-platformversion
        :param pulumi.Input['TaskSetScaleArgs'] scale: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-scale
        :param pulumi.Input[Sequence[pulumi.Input['TaskSetServiceRegistryArgs']]] service_registries: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-serviceregistries
        """
        pulumi.set(__self__, "cluster", cluster)
        pulumi.set(__self__, "service", service)
        pulumi.set(__self__, "task_definition", task_definition)
        if external_id is not None:
            pulumi.set(__self__, "external_id", external_id)
        if launch_type is not None:
            pulumi.set(__self__, "launch_type", launch_type)
        if load_balancers is not None:
            pulumi.set(__self__, "load_balancers", load_balancers)
        if network_configuration is not None:
            pulumi.set(__self__, "network_configuration", network_configuration)
        if platform_version is not None:
            pulumi.set(__self__, "platform_version", platform_version)
        if scale is not None:
            pulumi.set(__self__, "scale", scale)
        if service_registries is not None:
            pulumi.set(__self__, "service_registries", service_registries)

    @property
    @pulumi.getter
    def cluster(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-cluster
        """
        return pulumi.get(self, "cluster")

    @cluster.setter
    def cluster(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster", value)

    @property
    @pulumi.getter
    def service(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-service
        """
        return pulumi.get(self, "service")

    @service.setter
    def service(self, value: pulumi.Input[str]):
        pulumi.set(self, "service", value)

    @property
    @pulumi.getter(name="taskDefinition")
    def task_definition(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-taskdefinition
        """
        return pulumi.get(self, "task_definition")

    @task_definition.setter
    def task_definition(self, value: pulumi.Input[str]):
        pulumi.set(self, "task_definition", value)

    @property
    @pulumi.getter(name="externalId")
    def external_id(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-externalid
        """
        return pulumi.get(self, "external_id")

    @external_id.setter
    def external_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_id", value)

    @property
    @pulumi.getter(name="launchType")
    def launch_type(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-launchtype
        """
        return pulumi.get(self, "launch_type")

    @launch_type.setter
    def launch_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "launch_type", value)

    @property
    @pulumi.getter(name="loadBalancers")
    def load_balancers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TaskSetLoadBalancerArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-loadbalancers
        """
        return pulumi.get(self, "load_balancers")

    @load_balancers.setter
    def load_balancers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TaskSetLoadBalancerArgs']]]]):
        pulumi.set(self, "load_balancers", value)

    @property
    @pulumi.getter(name="networkConfiguration")
    def network_configuration(self) -> Optional[pulumi.Input['TaskSetNetworkConfigurationArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-networkconfiguration
        """
        return pulumi.get(self, "network_configuration")

    @network_configuration.setter
    def network_configuration(self, value: Optional[pulumi.Input['TaskSetNetworkConfigurationArgs']]):
        pulumi.set(self, "network_configuration", value)

    @property
    @pulumi.getter(name="platformVersion")
    def platform_version(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-platformversion
        """
        return pulumi.get(self, "platform_version")

    @platform_version.setter
    def platform_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "platform_version", value)

    @property
    @pulumi.getter
    def scale(self) -> Optional[pulumi.Input['TaskSetScaleArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-scale
        """
        return pulumi.get(self, "scale")

    @scale.setter
    def scale(self, value: Optional[pulumi.Input['TaskSetScaleArgs']]):
        pulumi.set(self, "scale", value)

    @property
    @pulumi.getter(name="serviceRegistries")
    def service_registries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TaskSetServiceRegistryArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-serviceregistries
        """
        return pulumi.get(self, "service_registries")

    @service_registries.setter
    def service_registries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TaskSetServiceRegistryArgs']]]]):
        pulumi.set(self, "service_registries", value)


class TaskSet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster: Optional[pulumi.Input[str]] = None,
                 external_id: Optional[pulumi.Input[str]] = None,
                 launch_type: Optional[pulumi.Input[str]] = None,
                 load_balancers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskSetLoadBalancerArgs']]]]] = None,
                 network_configuration: Optional[pulumi.Input[pulumi.InputType['TaskSetNetworkConfigurationArgs']]] = None,
                 platform_version: Optional[pulumi.Input[str]] = None,
                 scale: Optional[pulumi.Input[pulumi.InputType['TaskSetScaleArgs']]] = None,
                 service: Optional[pulumi.Input[str]] = None,
                 service_registries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskSetServiceRegistryArgs']]]]] = None,
                 task_definition: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-cluster
        :param pulumi.Input[str] external_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-externalid
        :param pulumi.Input[str] launch_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-launchtype
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskSetLoadBalancerArgs']]]] load_balancers: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-loadbalancers
        :param pulumi.Input[pulumi.InputType['TaskSetNetworkConfigurationArgs']] network_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-networkconfiguration
        :param pulumi.Input[str] platform_version: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-platformversion
        :param pulumi.Input[pulumi.InputType['TaskSetScaleArgs']] scale: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-scale
        :param pulumi.Input[str] service: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-service
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskSetServiceRegistryArgs']]]] service_registries: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-serviceregistries
        :param pulumi.Input[str] task_definition: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-taskdefinition
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TaskSetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html

        :param str resource_name: The name of the resource.
        :param TaskSetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TaskSetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster: Optional[pulumi.Input[str]] = None,
                 external_id: Optional[pulumi.Input[str]] = None,
                 launch_type: Optional[pulumi.Input[str]] = None,
                 load_balancers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskSetLoadBalancerArgs']]]]] = None,
                 network_configuration: Optional[pulumi.Input[pulumi.InputType['TaskSetNetworkConfigurationArgs']]] = None,
                 platform_version: Optional[pulumi.Input[str]] = None,
                 scale: Optional[pulumi.Input[pulumi.InputType['TaskSetScaleArgs']]] = None,
                 service: Optional[pulumi.Input[str]] = None,
                 service_registries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskSetServiceRegistryArgs']]]]] = None,
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
            __props__ = TaskSetArgs.__new__(TaskSetArgs)

            if cluster is None and not opts.urn:
                raise TypeError("Missing required property 'cluster'")
            __props__.__dict__["cluster"] = cluster
            __props__.__dict__["external_id"] = external_id
            __props__.__dict__["launch_type"] = launch_type
            __props__.__dict__["load_balancers"] = load_balancers
            __props__.__dict__["network_configuration"] = network_configuration
            __props__.__dict__["platform_version"] = platform_version
            __props__.__dict__["scale"] = scale
            if service is None and not opts.urn:
                raise TypeError("Missing required property 'service'")
            __props__.__dict__["service"] = service
            __props__.__dict__["service_registries"] = service_registries
            if task_definition is None and not opts.urn:
                raise TypeError("Missing required property 'task_definition'")
            __props__.__dict__["task_definition"] = task_definition
            __props__.__dict__["id"] = None
        super(TaskSet, __self__).__init__(
            'aws-native:ECS:TaskSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'TaskSet':
        """
        Get an existing TaskSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TaskSetArgs.__new__(TaskSetArgs)

        __props__.__dict__["cluster"] = None
        __props__.__dict__["external_id"] = None
        __props__.__dict__["id"] = None
        __props__.__dict__["launch_type"] = None
        __props__.__dict__["load_balancers"] = None
        __props__.__dict__["network_configuration"] = None
        __props__.__dict__["platform_version"] = None
        __props__.__dict__["scale"] = None
        __props__.__dict__["service"] = None
        __props__.__dict__["service_registries"] = None
        __props__.__dict__["task_definition"] = None
        return TaskSet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cluster(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-cluster
        """
        return pulumi.get(self, "cluster")

    @property
    @pulumi.getter(name="externalId")
    def external_id(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-externalid
        """
        return pulumi.get(self, "external_id")

    @property
    @pulumi.getter
    def id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="launchType")
    def launch_type(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-launchtype
        """
        return pulumi.get(self, "launch_type")

    @property
    @pulumi.getter(name="loadBalancers")
    def load_balancers(self) -> pulumi.Output[Optional[Sequence['outputs.TaskSetLoadBalancer']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-loadbalancers
        """
        return pulumi.get(self, "load_balancers")

    @property
    @pulumi.getter(name="networkConfiguration")
    def network_configuration(self) -> pulumi.Output[Optional['outputs.TaskSetNetworkConfiguration']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-networkconfiguration
        """
        return pulumi.get(self, "network_configuration")

    @property
    @pulumi.getter(name="platformVersion")
    def platform_version(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-platformversion
        """
        return pulumi.get(self, "platform_version")

    @property
    @pulumi.getter
    def scale(self) -> pulumi.Output[Optional['outputs.TaskSetScale']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-scale
        """
        return pulumi.get(self, "scale")

    @property
    @pulumi.getter
    def service(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-service
        """
        return pulumi.get(self, "service")

    @property
    @pulumi.getter(name="serviceRegistries")
    def service_registries(self) -> pulumi.Output[Optional[Sequence['outputs.TaskSetServiceRegistry']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-serviceregistries
        """
        return pulumi.get(self, "service_registries")

    @property
    @pulumi.getter(name="taskDefinition")
    def task_definition(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-taskdefinition
        """
        return pulumi.get(self, "task_definition")

