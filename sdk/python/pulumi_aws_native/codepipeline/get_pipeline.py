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
    'GetPipelineResult',
    'AwaitableGetPipelineResult',
    'get_pipeline',
    'get_pipeline_output',
]

@pulumi.output_type
class GetPipelineResult:
    def __init__(__self__, artifact_store=None, artifact_stores=None, disable_inbound_stage_transitions=None, id=None, restart_execution_on_update=None, role_arn=None, stages=None, tags=None, version=None):
        if artifact_store and not isinstance(artifact_store, dict):
            raise TypeError("Expected argument 'artifact_store' to be a dict")
        pulumi.set(__self__, "artifact_store", artifact_store)
        if artifact_stores and not isinstance(artifact_stores, list):
            raise TypeError("Expected argument 'artifact_stores' to be a list")
        pulumi.set(__self__, "artifact_stores", artifact_stores)
        if disable_inbound_stage_transitions and not isinstance(disable_inbound_stage_transitions, list):
            raise TypeError("Expected argument 'disable_inbound_stage_transitions' to be a list")
        pulumi.set(__self__, "disable_inbound_stage_transitions", disable_inbound_stage_transitions)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if restart_execution_on_update and not isinstance(restart_execution_on_update, bool):
            raise TypeError("Expected argument 'restart_execution_on_update' to be a bool")
        pulumi.set(__self__, "restart_execution_on_update", restart_execution_on_update)
        if role_arn and not isinstance(role_arn, str):
            raise TypeError("Expected argument 'role_arn' to be a str")
        pulumi.set(__self__, "role_arn", role_arn)
        if stages and not isinstance(stages, list):
            raise TypeError("Expected argument 'stages' to be a list")
        pulumi.set(__self__, "stages", stages)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="artifactStore")
    def artifact_store(self) -> Optional['outputs.PipelineArtifactStore']:
        return pulumi.get(self, "artifact_store")

    @property
    @pulumi.getter(name="artifactStores")
    def artifact_stores(self) -> Optional[Sequence['outputs.PipelineArtifactStoreMap']]:
        return pulumi.get(self, "artifact_stores")

    @property
    @pulumi.getter(name="disableInboundStageTransitions")
    def disable_inbound_stage_transitions(self) -> Optional[Sequence['outputs.PipelineStageTransition']]:
        return pulumi.get(self, "disable_inbound_stage_transitions")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="restartExecutionOnUpdate")
    def restart_execution_on_update(self) -> Optional[bool]:
        return pulumi.get(self, "restart_execution_on_update")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[str]:
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter
    def stages(self) -> Optional[Sequence['outputs.PipelineStageDeclaration']]:
        return pulumi.get(self, "stages")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.PipelineTag']]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        return pulumi.get(self, "version")


class AwaitableGetPipelineResult(GetPipelineResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPipelineResult(
            artifact_store=self.artifact_store,
            artifact_stores=self.artifact_stores,
            disable_inbound_stage_transitions=self.disable_inbound_stage_transitions,
            id=self.id,
            restart_execution_on_update=self.restart_execution_on_update,
            role_arn=self.role_arn,
            stages=self.stages,
            tags=self.tags,
            version=self.version)


def get_pipeline(id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPipelineResult:
    """
    Resource Type definition for AWS::CodePipeline::Pipeline
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:codepipeline:getPipeline', __args__, opts=opts, typ=GetPipelineResult).value

    return AwaitableGetPipelineResult(
        artifact_store=__ret__.artifact_store,
        artifact_stores=__ret__.artifact_stores,
        disable_inbound_stage_transitions=__ret__.disable_inbound_stage_transitions,
        id=__ret__.id,
        restart_execution_on_update=__ret__.restart_execution_on_update,
        role_arn=__ret__.role_arn,
        stages=__ret__.stages,
        tags=__ret__.tags,
        version=__ret__.version)


@_utilities.lift_output_func(get_pipeline)
def get_pipeline_output(id: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPipelineResult]:
    """
    Resource Type definition for AWS::CodePipeline::Pipeline
    """
    ...
