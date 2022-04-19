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

__all__ = [
    'EnvironmentLoggingConfiguration',
    'EnvironmentModuleLoggingConfiguration',
    'EnvironmentNetworkConfiguration',
]

@pulumi.output_type
class EnvironmentLoggingConfiguration(dict):
    """
    Logging configuration for the environment.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "dagProcessingLogs":
            suggest = "dag_processing_logs"
        elif key == "schedulerLogs":
            suggest = "scheduler_logs"
        elif key == "taskLogs":
            suggest = "task_logs"
        elif key == "webserverLogs":
            suggest = "webserver_logs"
        elif key == "workerLogs":
            suggest = "worker_logs"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in EnvironmentLoggingConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        EnvironmentLoggingConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        EnvironmentLoggingConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 dag_processing_logs: Optional['outputs.EnvironmentModuleLoggingConfiguration'] = None,
                 scheduler_logs: Optional['outputs.EnvironmentModuleLoggingConfiguration'] = None,
                 task_logs: Optional['outputs.EnvironmentModuleLoggingConfiguration'] = None,
                 webserver_logs: Optional['outputs.EnvironmentModuleLoggingConfiguration'] = None,
                 worker_logs: Optional['outputs.EnvironmentModuleLoggingConfiguration'] = None):
        """
        Logging configuration for the environment.
        """
        if dag_processing_logs is not None:
            pulumi.set(__self__, "dag_processing_logs", dag_processing_logs)
        if scheduler_logs is not None:
            pulumi.set(__self__, "scheduler_logs", scheduler_logs)
        if task_logs is not None:
            pulumi.set(__self__, "task_logs", task_logs)
        if webserver_logs is not None:
            pulumi.set(__self__, "webserver_logs", webserver_logs)
        if worker_logs is not None:
            pulumi.set(__self__, "worker_logs", worker_logs)

    @property
    @pulumi.getter(name="dagProcessingLogs")
    def dag_processing_logs(self) -> Optional['outputs.EnvironmentModuleLoggingConfiguration']:
        return pulumi.get(self, "dag_processing_logs")

    @property
    @pulumi.getter(name="schedulerLogs")
    def scheduler_logs(self) -> Optional['outputs.EnvironmentModuleLoggingConfiguration']:
        return pulumi.get(self, "scheduler_logs")

    @property
    @pulumi.getter(name="taskLogs")
    def task_logs(self) -> Optional['outputs.EnvironmentModuleLoggingConfiguration']:
        return pulumi.get(self, "task_logs")

    @property
    @pulumi.getter(name="webserverLogs")
    def webserver_logs(self) -> Optional['outputs.EnvironmentModuleLoggingConfiguration']:
        return pulumi.get(self, "webserver_logs")

    @property
    @pulumi.getter(name="workerLogs")
    def worker_logs(self) -> Optional['outputs.EnvironmentModuleLoggingConfiguration']:
        return pulumi.get(self, "worker_logs")


@pulumi.output_type
class EnvironmentModuleLoggingConfiguration(dict):
    """
    Logging configuration for a specific airflow component.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "cloudWatchLogGroupArn":
            suggest = "cloud_watch_log_group_arn"
        elif key == "logLevel":
            suggest = "log_level"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in EnvironmentModuleLoggingConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        EnvironmentModuleLoggingConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        EnvironmentModuleLoggingConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cloud_watch_log_group_arn: Optional[str] = None,
                 enabled: Optional[bool] = None,
                 log_level: Optional['EnvironmentLoggingLevel'] = None):
        """
        Logging configuration for a specific airflow component.
        """
        if cloud_watch_log_group_arn is not None:
            pulumi.set(__self__, "cloud_watch_log_group_arn", cloud_watch_log_group_arn)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if log_level is not None:
            pulumi.set(__self__, "log_level", log_level)

    @property
    @pulumi.getter(name="cloudWatchLogGroupArn")
    def cloud_watch_log_group_arn(self) -> Optional[str]:
        return pulumi.get(self, "cloud_watch_log_group_arn")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="logLevel")
    def log_level(self) -> Optional['EnvironmentLoggingLevel']:
        return pulumi.get(self, "log_level")


@pulumi.output_type
class EnvironmentNetworkConfiguration(dict):
    """
    Configures the network resources of the environment.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "securityGroupIds":
            suggest = "security_group_ids"
        elif key == "subnetIds":
            suggest = "subnet_ids"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in EnvironmentNetworkConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        EnvironmentNetworkConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        EnvironmentNetworkConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 security_group_ids: Optional[Sequence[str]] = None,
                 subnet_ids: Optional[Sequence[str]] = None):
        """
        Configures the network resources of the environment.
        :param Sequence[str] security_group_ids: A list of security groups to use for the environment.
        :param Sequence[str] subnet_ids: A list of subnets to use for the environment. These must be private subnets, in the same VPC, in two different availability zones.
        """
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if subnet_ids is not None:
            pulumi.set(__self__, "subnet_ids", subnet_ids)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[Sequence[str]]:
        """
        A list of security groups to use for the environment.
        """
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Optional[Sequence[str]]:
        """
        A list of subnets to use for the environment. These must be private subnets, in the same VPC, in two different availability zones.
        """
        return pulumi.get(self, "subnet_ids")


