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
    'GetDeploymentGroupResult',
    'AwaitableGetDeploymentGroupResult',
    'get_deployment_group',
    'get_deployment_group_output',
]

@pulumi.output_type
class GetDeploymentGroupResult:
    def __init__(__self__, alarm_configuration=None, auto_rollback_configuration=None, auto_scaling_groups=None, blue_green_deployment_configuration=None, deployment=None, deployment_config_name=None, deployment_style=None, e_cs_services=None, ec2_tag_filters=None, ec2_tag_set=None, id=None, load_balancer_info=None, on_premises_instance_tag_filters=None, on_premises_tag_set=None, outdated_instances_strategy=None, service_role_arn=None, tags=None, trigger_configurations=None):
        if alarm_configuration and not isinstance(alarm_configuration, dict):
            raise TypeError("Expected argument 'alarm_configuration' to be a dict")
        pulumi.set(__self__, "alarm_configuration", alarm_configuration)
        if auto_rollback_configuration and not isinstance(auto_rollback_configuration, dict):
            raise TypeError("Expected argument 'auto_rollback_configuration' to be a dict")
        pulumi.set(__self__, "auto_rollback_configuration", auto_rollback_configuration)
        if auto_scaling_groups and not isinstance(auto_scaling_groups, list):
            raise TypeError("Expected argument 'auto_scaling_groups' to be a list")
        pulumi.set(__self__, "auto_scaling_groups", auto_scaling_groups)
        if blue_green_deployment_configuration and not isinstance(blue_green_deployment_configuration, dict):
            raise TypeError("Expected argument 'blue_green_deployment_configuration' to be a dict")
        pulumi.set(__self__, "blue_green_deployment_configuration", blue_green_deployment_configuration)
        if deployment and not isinstance(deployment, dict):
            raise TypeError("Expected argument 'deployment' to be a dict")
        pulumi.set(__self__, "deployment", deployment)
        if deployment_config_name and not isinstance(deployment_config_name, str):
            raise TypeError("Expected argument 'deployment_config_name' to be a str")
        pulumi.set(__self__, "deployment_config_name", deployment_config_name)
        if deployment_style and not isinstance(deployment_style, dict):
            raise TypeError("Expected argument 'deployment_style' to be a dict")
        pulumi.set(__self__, "deployment_style", deployment_style)
        if e_cs_services and not isinstance(e_cs_services, list):
            raise TypeError("Expected argument 'e_cs_services' to be a list")
        pulumi.set(__self__, "e_cs_services", e_cs_services)
        if ec2_tag_filters and not isinstance(ec2_tag_filters, list):
            raise TypeError("Expected argument 'ec2_tag_filters' to be a list")
        pulumi.set(__self__, "ec2_tag_filters", ec2_tag_filters)
        if ec2_tag_set and not isinstance(ec2_tag_set, dict):
            raise TypeError("Expected argument 'ec2_tag_set' to be a dict")
        pulumi.set(__self__, "ec2_tag_set", ec2_tag_set)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if load_balancer_info and not isinstance(load_balancer_info, dict):
            raise TypeError("Expected argument 'load_balancer_info' to be a dict")
        pulumi.set(__self__, "load_balancer_info", load_balancer_info)
        if on_premises_instance_tag_filters and not isinstance(on_premises_instance_tag_filters, list):
            raise TypeError("Expected argument 'on_premises_instance_tag_filters' to be a list")
        pulumi.set(__self__, "on_premises_instance_tag_filters", on_premises_instance_tag_filters)
        if on_premises_tag_set and not isinstance(on_premises_tag_set, dict):
            raise TypeError("Expected argument 'on_premises_tag_set' to be a dict")
        pulumi.set(__self__, "on_premises_tag_set", on_premises_tag_set)
        if outdated_instances_strategy and not isinstance(outdated_instances_strategy, str):
            raise TypeError("Expected argument 'outdated_instances_strategy' to be a str")
        pulumi.set(__self__, "outdated_instances_strategy", outdated_instances_strategy)
        if service_role_arn and not isinstance(service_role_arn, str):
            raise TypeError("Expected argument 'service_role_arn' to be a str")
        pulumi.set(__self__, "service_role_arn", service_role_arn)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if trigger_configurations and not isinstance(trigger_configurations, list):
            raise TypeError("Expected argument 'trigger_configurations' to be a list")
        pulumi.set(__self__, "trigger_configurations", trigger_configurations)

    @property
    @pulumi.getter(name="alarmConfiguration")
    def alarm_configuration(self) -> Optional['outputs.DeploymentGroupAlarmConfiguration']:
        return pulumi.get(self, "alarm_configuration")

    @property
    @pulumi.getter(name="autoRollbackConfiguration")
    def auto_rollback_configuration(self) -> Optional['outputs.DeploymentGroupAutoRollbackConfiguration']:
        return pulumi.get(self, "auto_rollback_configuration")

    @property
    @pulumi.getter(name="autoScalingGroups")
    def auto_scaling_groups(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "auto_scaling_groups")

    @property
    @pulumi.getter(name="blueGreenDeploymentConfiguration")
    def blue_green_deployment_configuration(self) -> Optional['outputs.DeploymentGroupBlueGreenDeploymentConfiguration']:
        return pulumi.get(self, "blue_green_deployment_configuration")

    @property
    @pulumi.getter
    def deployment(self) -> Optional['outputs.DeploymentGroupDeployment']:
        return pulumi.get(self, "deployment")

    @property
    @pulumi.getter(name="deploymentConfigName")
    def deployment_config_name(self) -> Optional[str]:
        return pulumi.get(self, "deployment_config_name")

    @property
    @pulumi.getter(name="deploymentStyle")
    def deployment_style(self) -> Optional['outputs.DeploymentGroupDeploymentStyle']:
        return pulumi.get(self, "deployment_style")

    @property
    @pulumi.getter(name="eCSServices")
    def e_cs_services(self) -> Optional[Sequence['outputs.DeploymentGroupECSService']]:
        return pulumi.get(self, "e_cs_services")

    @property
    @pulumi.getter(name="ec2TagFilters")
    def ec2_tag_filters(self) -> Optional[Sequence['outputs.DeploymentGroupEC2TagFilter']]:
        return pulumi.get(self, "ec2_tag_filters")

    @property
    @pulumi.getter(name="ec2TagSet")
    def ec2_tag_set(self) -> Optional['outputs.DeploymentGroupEC2TagSet']:
        return pulumi.get(self, "ec2_tag_set")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="loadBalancerInfo")
    def load_balancer_info(self) -> Optional['outputs.DeploymentGroupLoadBalancerInfo']:
        return pulumi.get(self, "load_balancer_info")

    @property
    @pulumi.getter(name="onPremisesInstanceTagFilters")
    def on_premises_instance_tag_filters(self) -> Optional[Sequence['outputs.DeploymentGroupTagFilter']]:
        return pulumi.get(self, "on_premises_instance_tag_filters")

    @property
    @pulumi.getter(name="onPremisesTagSet")
    def on_premises_tag_set(self) -> Optional['outputs.DeploymentGroupOnPremisesTagSet']:
        return pulumi.get(self, "on_premises_tag_set")

    @property
    @pulumi.getter(name="outdatedInstancesStrategy")
    def outdated_instances_strategy(self) -> Optional[str]:
        return pulumi.get(self, "outdated_instances_strategy")

    @property
    @pulumi.getter(name="serviceRoleArn")
    def service_role_arn(self) -> Optional[str]:
        return pulumi.get(self, "service_role_arn")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.DeploymentGroupTag']]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="triggerConfigurations")
    def trigger_configurations(self) -> Optional[Sequence['outputs.DeploymentGroupTriggerConfig']]:
        return pulumi.get(self, "trigger_configurations")


class AwaitableGetDeploymentGroupResult(GetDeploymentGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeploymentGroupResult(
            alarm_configuration=self.alarm_configuration,
            auto_rollback_configuration=self.auto_rollback_configuration,
            auto_scaling_groups=self.auto_scaling_groups,
            blue_green_deployment_configuration=self.blue_green_deployment_configuration,
            deployment=self.deployment,
            deployment_config_name=self.deployment_config_name,
            deployment_style=self.deployment_style,
            e_cs_services=self.e_cs_services,
            ec2_tag_filters=self.ec2_tag_filters,
            ec2_tag_set=self.ec2_tag_set,
            id=self.id,
            load_balancer_info=self.load_balancer_info,
            on_premises_instance_tag_filters=self.on_premises_instance_tag_filters,
            on_premises_tag_set=self.on_premises_tag_set,
            outdated_instances_strategy=self.outdated_instances_strategy,
            service_role_arn=self.service_role_arn,
            tags=self.tags,
            trigger_configurations=self.trigger_configurations)


def get_deployment_group(id: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeploymentGroupResult:
    """
    Resource Type definition for AWS::CodeDeploy::DeploymentGroup
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:codedeploy:getDeploymentGroup', __args__, opts=opts, typ=GetDeploymentGroupResult).value

    return AwaitableGetDeploymentGroupResult(
        alarm_configuration=__ret__.alarm_configuration,
        auto_rollback_configuration=__ret__.auto_rollback_configuration,
        auto_scaling_groups=__ret__.auto_scaling_groups,
        blue_green_deployment_configuration=__ret__.blue_green_deployment_configuration,
        deployment=__ret__.deployment,
        deployment_config_name=__ret__.deployment_config_name,
        deployment_style=__ret__.deployment_style,
        e_cs_services=__ret__.e_cs_services,
        ec2_tag_filters=__ret__.ec2_tag_filters,
        ec2_tag_set=__ret__.ec2_tag_set,
        id=__ret__.id,
        load_balancer_info=__ret__.load_balancer_info,
        on_premises_instance_tag_filters=__ret__.on_premises_instance_tag_filters,
        on_premises_tag_set=__ret__.on_premises_tag_set,
        outdated_instances_strategy=__ret__.outdated_instances_strategy,
        service_role_arn=__ret__.service_role_arn,
        tags=__ret__.tags,
        trigger_configurations=__ret__.trigger_configurations)


@_utilities.lift_output_func(get_deployment_group)
def get_deployment_group_output(id: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDeploymentGroupResult]:
    """
    Resource Type definition for AWS::CodeDeploy::DeploymentGroup
    """
    ...
