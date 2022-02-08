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
    'GetJobTemplateResult',
    'AwaitableGetJobTemplateResult',
    'get_job_template',
    'get_job_template_output',
]

@pulumi.output_type
class GetJobTemplateResult:
    def __init__(__self__, acceleration_settings=None, arn=None, category=None, description=None, hop_destinations=None, id=None, priority=None, queue=None, settings_json=None, status_update_interval=None, tags=None):
        if acceleration_settings and not isinstance(acceleration_settings, dict):
            raise TypeError("Expected argument 'acceleration_settings' to be a dict")
        pulumi.set(__self__, "acceleration_settings", acceleration_settings)
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if category and not isinstance(category, str):
            raise TypeError("Expected argument 'category' to be a str")
        pulumi.set(__self__, "category", category)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if hop_destinations and not isinstance(hop_destinations, list):
            raise TypeError("Expected argument 'hop_destinations' to be a list")
        pulumi.set(__self__, "hop_destinations", hop_destinations)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if priority and not isinstance(priority, int):
            raise TypeError("Expected argument 'priority' to be a int")
        pulumi.set(__self__, "priority", priority)
        if queue and not isinstance(queue, str):
            raise TypeError("Expected argument 'queue' to be a str")
        pulumi.set(__self__, "queue", queue)
        if settings_json and not isinstance(settings_json, dict):
            raise TypeError("Expected argument 'settings_json' to be a dict")
        pulumi.set(__self__, "settings_json", settings_json)
        if status_update_interval and not isinstance(status_update_interval, str):
            raise TypeError("Expected argument 'status_update_interval' to be a str")
        pulumi.set(__self__, "status_update_interval", status_update_interval)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="accelerationSettings")
    def acceleration_settings(self) -> Optional['outputs.JobTemplateAccelerationSettings']:
        return pulumi.get(self, "acceleration_settings")

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def category(self) -> Optional[str]:
        return pulumi.get(self, "category")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="hopDestinations")
    def hop_destinations(self) -> Optional[Sequence['outputs.JobTemplateHopDestination']]:
        return pulumi.get(self, "hop_destinations")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def priority(self) -> Optional[int]:
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def queue(self) -> Optional[str]:
        return pulumi.get(self, "queue")

    @property
    @pulumi.getter(name="settingsJson")
    def settings_json(self) -> Optional[Any]:
        return pulumi.get(self, "settings_json")

    @property
    @pulumi.getter(name="statusUpdateInterval")
    def status_update_interval(self) -> Optional[str]:
        return pulumi.get(self, "status_update_interval")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Any]:
        return pulumi.get(self, "tags")


class AwaitableGetJobTemplateResult(GetJobTemplateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetJobTemplateResult(
            acceleration_settings=self.acceleration_settings,
            arn=self.arn,
            category=self.category,
            description=self.description,
            hop_destinations=self.hop_destinations,
            id=self.id,
            priority=self.priority,
            queue=self.queue,
            settings_json=self.settings_json,
            status_update_interval=self.status_update_interval,
            tags=self.tags)


def get_job_template(id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetJobTemplateResult:
    """
    Resource Type definition for AWS::MediaConvert::JobTemplate
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:mediaconvert:getJobTemplate', __args__, opts=opts, typ=GetJobTemplateResult).value

    return AwaitableGetJobTemplateResult(
        acceleration_settings=__ret__.acceleration_settings,
        arn=__ret__.arn,
        category=__ret__.category,
        description=__ret__.description,
        hop_destinations=__ret__.hop_destinations,
        id=__ret__.id,
        priority=__ret__.priority,
        queue=__ret__.queue,
        settings_json=__ret__.settings_json,
        status_update_interval=__ret__.status_update_interval,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_job_template)
def get_job_template_output(id: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetJobTemplateResult]:
    """
    Resource Type definition for AWS::MediaConvert::JobTemplate
    """
    ...