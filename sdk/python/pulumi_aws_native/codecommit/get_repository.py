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
    'GetRepositoryResult',
    'AwaitableGetRepositoryResult',
    'get_repository',
    'get_repository_output',
]

@pulumi.output_type
class GetRepositoryResult:
    def __init__(__self__, arn=None, clone_url_http=None, clone_url_ssh=None, code=None, id=None, name=None, repository_description=None, repository_name=None, tags=None, triggers=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if clone_url_http and not isinstance(clone_url_http, str):
            raise TypeError("Expected argument 'clone_url_http' to be a str")
        pulumi.set(__self__, "clone_url_http", clone_url_http)
        if clone_url_ssh and not isinstance(clone_url_ssh, str):
            raise TypeError("Expected argument 'clone_url_ssh' to be a str")
        pulumi.set(__self__, "clone_url_ssh", clone_url_ssh)
        if code and not isinstance(code, dict):
            raise TypeError("Expected argument 'code' to be a dict")
        pulumi.set(__self__, "code", code)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if repository_description and not isinstance(repository_description, str):
            raise TypeError("Expected argument 'repository_description' to be a str")
        pulumi.set(__self__, "repository_description", repository_description)
        if repository_name and not isinstance(repository_name, str):
            raise TypeError("Expected argument 'repository_name' to be a str")
        pulumi.set(__self__, "repository_name", repository_name)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if triggers and not isinstance(triggers, list):
            raise TypeError("Expected argument 'triggers' to be a list")
        pulumi.set(__self__, "triggers", triggers)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="cloneUrlHttp")
    def clone_url_http(self) -> Optional[str]:
        return pulumi.get(self, "clone_url_http")

    @property
    @pulumi.getter(name="cloneUrlSsh")
    def clone_url_ssh(self) -> Optional[str]:
        return pulumi.get(self, "clone_url_ssh")

    @property
    @pulumi.getter
    def code(self) -> Optional['outputs.RepositoryCode']:
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="repositoryDescription")
    def repository_description(self) -> Optional[str]:
        return pulumi.get(self, "repository_description")

    @property
    @pulumi.getter(name="repositoryName")
    def repository_name(self) -> Optional[str]:
        return pulumi.get(self, "repository_name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.RepositoryTag']]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def triggers(self) -> Optional[Sequence['outputs.RepositoryTrigger']]:
        return pulumi.get(self, "triggers")


class AwaitableGetRepositoryResult(GetRepositoryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRepositoryResult(
            arn=self.arn,
            clone_url_http=self.clone_url_http,
            clone_url_ssh=self.clone_url_ssh,
            code=self.code,
            id=self.id,
            name=self.name,
            repository_description=self.repository_description,
            repository_name=self.repository_name,
            tags=self.tags,
            triggers=self.triggers)


def get_repository(id: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRepositoryResult:
    """
    Resource Type definition for AWS::CodeCommit::Repository
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:codecommit:getRepository', __args__, opts=opts, typ=GetRepositoryResult).value

    return AwaitableGetRepositoryResult(
        arn=__ret__.arn,
        clone_url_http=__ret__.clone_url_http,
        clone_url_ssh=__ret__.clone_url_ssh,
        code=__ret__.code,
        id=__ret__.id,
        name=__ret__.name,
        repository_description=__ret__.repository_description,
        repository_name=__ret__.repository_name,
        tags=__ret__.tags,
        triggers=__ret__.triggers)


@_utilities.lift_output_func(get_repository)
def get_repository_output(id: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRepositoryResult]:
    """
    Resource Type definition for AWS::CodeCommit::Repository
    """
    ...
