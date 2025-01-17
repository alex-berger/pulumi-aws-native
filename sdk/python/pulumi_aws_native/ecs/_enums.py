# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'CapacityProviderAutoScalingGroupProviderManagedTerminationProtection',
    'CapacityProviderManagedScalingStatus',
    'ClusterCapacityProviderAssociationsCapacityProvider',
    'ServiceAwsVpcConfigurationAssignPublicIp',
    'ServiceDeploymentControllerType',
    'ServiceLaunchType',
    'ServicePlacementConstraintType',
    'ServicePlacementStrategyType',
    'ServicePropagateTags',
    'ServiceSchedulingStrategy',
    'TaskDefinitionAuthorizationConfigIAM',
    'TaskDefinitionEFSVolumeConfigurationTransitEncryption',
    'TaskSetAwsVpcConfigurationAssignPublicIp',
    'TaskSetLaunchType',
    'TaskSetScaleUnit',
]


class CapacityProviderAutoScalingGroupProviderManagedTerminationProtection(str, Enum):
    DISABLED = "DISABLED"
    ENABLED = "ENABLED"


class CapacityProviderManagedScalingStatus(str, Enum):
    DISABLED = "DISABLED"
    ENABLED = "ENABLED"


class ClusterCapacityProviderAssociationsCapacityProvider(str, Enum):
    FARGATE = "FARGATE"
    FARGATE_SPOT = "FARGATE_SPOT"


class ServiceAwsVpcConfigurationAssignPublicIp(str, Enum):
    DISABLED = "DISABLED"
    ENABLED = "ENABLED"


class ServiceDeploymentControllerType(str, Enum):
    CODE_DEPLOY = "CODE_DEPLOY"
    ECS = "ECS"
    EXTERNAL = "EXTERNAL"


class ServiceLaunchType(str, Enum):
    EC2 = "EC2"
    FARGATE = "FARGATE"
    EXTERNAL = "EXTERNAL"


class ServicePlacementConstraintType(str, Enum):
    DISTINCT_INSTANCE = "distinctInstance"
    MEMBER_OF = "memberOf"


class ServicePlacementStrategyType(str, Enum):
    BINPACK = "binpack"
    RANDOM = "random"
    SPREAD = "spread"


class ServicePropagateTags(str, Enum):
    SERVICE = "SERVICE"
    TASK_DEFINITION = "TASK_DEFINITION"


class ServiceSchedulingStrategy(str, Enum):
    DAEMON = "DAEMON"
    REPLICA = "REPLICA"


class TaskDefinitionAuthorizationConfigIAM(str, Enum):
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class TaskDefinitionEFSVolumeConfigurationTransitEncryption(str, Enum):
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class TaskSetAwsVpcConfigurationAssignPublicIp(str, Enum):
    """
    Whether the task's elastic network interface receives a public IP address. The default value is DISABLED.
    """
    DISABLED = "DISABLED"
    ENABLED = "ENABLED"


class TaskSetLaunchType(str, Enum):
    """
    The launch type that new tasks in the task set will use. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html in the Amazon Elastic Container Service Developer Guide. 
    """
    EC2 = "EC2"
    FARGATE = "FARGATE"


class TaskSetScaleUnit(str, Enum):
    """
    The unit of measure for the scale value.
    """
    PERCENT = "PERCENT"
