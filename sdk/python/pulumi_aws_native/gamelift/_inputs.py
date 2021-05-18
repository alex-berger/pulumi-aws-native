# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from .. import _inputs as _root_inputs

__all__ = [
    'AliasRoutingStrategyArgs',
    'GameServerGroupAutoScalingPolicyArgs',
    'GameServerGroupInstanceDefinitionArgs',
    'GameServerGroupInstanceDefinitionsArgs',
    'GameServerGroupLaunchTemplateArgs',
    'GameServerGroupTagsArgs',
    'GameServerGroupTargetTrackingConfigurationArgs',
    'GameServerGroupVpcSubnetsArgs',
]

@pulumi.input_type
class AliasRoutingStrategyArgs:
    def __init__(__self__, *,
                 fleet_id: Optional[pulumi.Input[str]] = None,
                 message: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-alias-routingstrategy.html
        :param pulumi.Input[str] fleet_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-alias-routingstrategy.html#cfn-gamelift-alias-routingstrategy-fleetid
        :param pulumi.Input[str] message: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-alias-routingstrategy.html#cfn-gamelift-alias-routingstrategy-message
        :param pulumi.Input[str] type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-alias-routingstrategy.html#cfn-gamelift-alias-routingstrategy-type
        """
        if fleet_id is not None:
            pulumi.set(__self__, "fleet_id", fleet_id)
        if message is not None:
            pulumi.set(__self__, "message", message)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="fleetId")
    def fleet_id(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-alias-routingstrategy.html#cfn-gamelift-alias-routingstrategy-fleetid
        """
        return pulumi.get(self, "fleet_id")

    @fleet_id.setter
    def fleet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fleet_id", value)

    @property
    @pulumi.getter
    def message(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-alias-routingstrategy.html#cfn-gamelift-alias-routingstrategy-message
        """
        return pulumi.get(self, "message")

    @message.setter
    def message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "message", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-alias-routingstrategy.html#cfn-gamelift-alias-routingstrategy-type
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class GameServerGroupAutoScalingPolicyArgs:
    def __init__(__self__, *,
                 target_tracking_configuration: pulumi.Input['GameServerGroupTargetTrackingConfigurationArgs'],
                 estimated_instance_warmup: Optional[pulumi.Input[float]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-autoscalingpolicy.html
        :param pulumi.Input['GameServerGroupTargetTrackingConfigurationArgs'] target_tracking_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-autoscalingpolicy.html#cfn-gamelift-gameservergroup-autoscalingpolicy-targettrackingconfiguration
        :param pulumi.Input[float] estimated_instance_warmup: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-autoscalingpolicy.html#cfn-gamelift-gameservergroup-autoscalingpolicy-estimatedinstancewarmup
        """
        pulumi.set(__self__, "target_tracking_configuration", target_tracking_configuration)
        if estimated_instance_warmup is not None:
            pulumi.set(__self__, "estimated_instance_warmup", estimated_instance_warmup)

    @property
    @pulumi.getter(name="targetTrackingConfiguration")
    def target_tracking_configuration(self) -> pulumi.Input['GameServerGroupTargetTrackingConfigurationArgs']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-autoscalingpolicy.html#cfn-gamelift-gameservergroup-autoscalingpolicy-targettrackingconfiguration
        """
        return pulumi.get(self, "target_tracking_configuration")

    @target_tracking_configuration.setter
    def target_tracking_configuration(self, value: pulumi.Input['GameServerGroupTargetTrackingConfigurationArgs']):
        pulumi.set(self, "target_tracking_configuration", value)

    @property
    @pulumi.getter(name="estimatedInstanceWarmup")
    def estimated_instance_warmup(self) -> Optional[pulumi.Input[float]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-autoscalingpolicy.html#cfn-gamelift-gameservergroup-autoscalingpolicy-estimatedinstancewarmup
        """
        return pulumi.get(self, "estimated_instance_warmup")

    @estimated_instance_warmup.setter
    def estimated_instance_warmup(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "estimated_instance_warmup", value)


@pulumi.input_type
class GameServerGroupInstanceDefinitionArgs:
    def __init__(__self__, *,
                 instance_type: pulumi.Input[str],
                 weighted_capacity: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-instancedefinition.html
        :param pulumi.Input[str] instance_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-instancedefinition.html#cfn-gamelift-gameservergroup-instancedefinition-instancetype
        :param pulumi.Input[str] weighted_capacity: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-instancedefinition.html#cfn-gamelift-gameservergroup-instancedefinition-weightedcapacity
        """
        pulumi.set(__self__, "instance_type", instance_type)
        if weighted_capacity is not None:
            pulumi.set(__self__, "weighted_capacity", weighted_capacity)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-instancedefinition.html#cfn-gamelift-gameservergroup-instancedefinition-instancetype
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter(name="weightedCapacity")
    def weighted_capacity(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-instancedefinition.html#cfn-gamelift-gameservergroup-instancedefinition-weightedcapacity
        """
        return pulumi.get(self, "weighted_capacity")

    @weighted_capacity.setter
    def weighted_capacity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "weighted_capacity", value)


@pulumi.input_type
class GameServerGroupInstanceDefinitionsArgs:
    def __init__(__self__, *,
                 instance_definitions: Optional[pulumi.Input[Sequence[pulumi.Input['GameServerGroupInstanceDefinitionArgs']]]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-instancedefinitions.html
        :param pulumi.Input[Sequence[pulumi.Input['GameServerGroupInstanceDefinitionArgs']]] instance_definitions: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-instancedefinitions.html#cfn-gamelift-gameservergroup-instancedefinitions-instancedefinitions
        """
        if instance_definitions is not None:
            pulumi.set(__self__, "instance_definitions", instance_definitions)

    @property
    @pulumi.getter(name="instanceDefinitions")
    def instance_definitions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GameServerGroupInstanceDefinitionArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-instancedefinitions.html#cfn-gamelift-gameservergroup-instancedefinitions-instancedefinitions
        """
        return pulumi.get(self, "instance_definitions")

    @instance_definitions.setter
    def instance_definitions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GameServerGroupInstanceDefinitionArgs']]]]):
        pulumi.set(self, "instance_definitions", value)


@pulumi.input_type
class GameServerGroupLaunchTemplateArgs:
    def __init__(__self__, *,
                 launch_template_id: Optional[pulumi.Input[str]] = None,
                 launch_template_name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-launchtemplate.html
        :param pulumi.Input[str] launch_template_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-launchtemplate.html#cfn-gamelift-gameservergroup-launchtemplate-launchtemplateid
        :param pulumi.Input[str] launch_template_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-launchtemplate.html#cfn-gamelift-gameservergroup-launchtemplate-launchtemplatename
        :param pulumi.Input[str] version: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-launchtemplate.html#cfn-gamelift-gameservergroup-launchtemplate-version
        """
        if launch_template_id is not None:
            pulumi.set(__self__, "launch_template_id", launch_template_id)
        if launch_template_name is not None:
            pulumi.set(__self__, "launch_template_name", launch_template_name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="launchTemplateId")
    def launch_template_id(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-launchtemplate.html#cfn-gamelift-gameservergroup-launchtemplate-launchtemplateid
        """
        return pulumi.get(self, "launch_template_id")

    @launch_template_id.setter
    def launch_template_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "launch_template_id", value)

    @property
    @pulumi.getter(name="launchTemplateName")
    def launch_template_name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-launchtemplate.html#cfn-gamelift-gameservergroup-launchtemplate-launchtemplatename
        """
        return pulumi.get(self, "launch_template_name")

    @launch_template_name.setter
    def launch_template_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "launch_template_name", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-launchtemplate.html#cfn-gamelift-gameservergroup-launchtemplate-version
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class GameServerGroupTagsArgs:
    def __init__(__self__, *,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-tags.html
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-tags.html#cfn-gamelift-gameservergroup-tags-tags
        """
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-tags.html#cfn-gamelift-gameservergroup-tags-tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class GameServerGroupTargetTrackingConfigurationArgs:
    def __init__(__self__, *,
                 target_value: pulumi.Input[float]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-targettrackingconfiguration.html
        :param pulumi.Input[float] target_value: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-targettrackingconfiguration.html#cfn-gamelift-gameservergroup-targettrackingconfiguration-targetvalue
        """
        pulumi.set(__self__, "target_value", target_value)

    @property
    @pulumi.getter(name="targetValue")
    def target_value(self) -> pulumi.Input[float]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-targettrackingconfiguration.html#cfn-gamelift-gameservergroup-targettrackingconfiguration-targetvalue
        """
        return pulumi.get(self, "target_value")

    @target_value.setter
    def target_value(self, value: pulumi.Input[float]):
        pulumi.set(self, "target_value", value)


@pulumi.input_type
class GameServerGroupVpcSubnetsArgs:
    def __init__(__self__, *,
                 vpc_subnets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-vpcsubnets.html
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpc_subnets: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-vpcsubnets.html#cfn-gamelift-gameservergroup-vpcsubnets-vpcsubnets
        """
        if vpc_subnets is not None:
            pulumi.set(__self__, "vpc_subnets", vpc_subnets)

    @property
    @pulumi.getter(name="vpcSubnets")
    def vpc_subnets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-vpcsubnets.html#cfn-gamelift-gameservergroup-vpcsubnets-vpcsubnets
        """
        return pulumi.get(self, "vpc_subnets")

    @vpc_subnets.setter
    def vpc_subnets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "vpc_subnets", value)


