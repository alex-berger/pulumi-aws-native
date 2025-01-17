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
    'GetMLTransformResult',
    'AwaitableGetMLTransformResult',
    'get_ml_transform',
    'get_ml_transform_output',
]

@pulumi.output_type
class GetMLTransformResult:
    def __init__(__self__, description=None, glue_version=None, id=None, max_capacity=None, max_retries=None, name=None, number_of_workers=None, role=None, tags=None, timeout=None, transform_encryption=None, transform_parameters=None, worker_type=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if glue_version and not isinstance(glue_version, str):
            raise TypeError("Expected argument 'glue_version' to be a str")
        pulumi.set(__self__, "glue_version", glue_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if max_capacity and not isinstance(max_capacity, float):
            raise TypeError("Expected argument 'max_capacity' to be a float")
        pulumi.set(__self__, "max_capacity", max_capacity)
        if max_retries and not isinstance(max_retries, int):
            raise TypeError("Expected argument 'max_retries' to be a int")
        pulumi.set(__self__, "max_retries", max_retries)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if number_of_workers and not isinstance(number_of_workers, int):
            raise TypeError("Expected argument 'number_of_workers' to be a int")
        pulumi.set(__self__, "number_of_workers", number_of_workers)
        if role and not isinstance(role, str):
            raise TypeError("Expected argument 'role' to be a str")
        pulumi.set(__self__, "role", role)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if timeout and not isinstance(timeout, int):
            raise TypeError("Expected argument 'timeout' to be a int")
        pulumi.set(__self__, "timeout", timeout)
        if transform_encryption and not isinstance(transform_encryption, dict):
            raise TypeError("Expected argument 'transform_encryption' to be a dict")
        pulumi.set(__self__, "transform_encryption", transform_encryption)
        if transform_parameters and not isinstance(transform_parameters, dict):
            raise TypeError("Expected argument 'transform_parameters' to be a dict")
        pulumi.set(__self__, "transform_parameters", transform_parameters)
        if worker_type and not isinstance(worker_type, str):
            raise TypeError("Expected argument 'worker_type' to be a str")
        pulumi.set(__self__, "worker_type", worker_type)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="glueVersion")
    def glue_version(self) -> Optional[str]:
        return pulumi.get(self, "glue_version")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="maxCapacity")
    def max_capacity(self) -> Optional[float]:
        return pulumi.get(self, "max_capacity")

    @property
    @pulumi.getter(name="maxRetries")
    def max_retries(self) -> Optional[int]:
        return pulumi.get(self, "max_retries")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="numberOfWorkers")
    def number_of_workers(self) -> Optional[int]:
        return pulumi.get(self, "number_of_workers")

    @property
    @pulumi.getter
    def role(self) -> Optional[str]:
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Any]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def timeout(self) -> Optional[int]:
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter(name="transformEncryption")
    def transform_encryption(self) -> Optional['outputs.MLTransformTransformEncryption']:
        return pulumi.get(self, "transform_encryption")

    @property
    @pulumi.getter(name="transformParameters")
    def transform_parameters(self) -> Optional['outputs.MLTransformTransformParameters']:
        return pulumi.get(self, "transform_parameters")

    @property
    @pulumi.getter(name="workerType")
    def worker_type(self) -> Optional[str]:
        return pulumi.get(self, "worker_type")


class AwaitableGetMLTransformResult(GetMLTransformResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMLTransformResult(
            description=self.description,
            glue_version=self.glue_version,
            id=self.id,
            max_capacity=self.max_capacity,
            max_retries=self.max_retries,
            name=self.name,
            number_of_workers=self.number_of_workers,
            role=self.role,
            tags=self.tags,
            timeout=self.timeout,
            transform_encryption=self.transform_encryption,
            transform_parameters=self.transform_parameters,
            worker_type=self.worker_type)


def get_ml_transform(id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMLTransformResult:
    """
    Resource Type definition for AWS::Glue::MLTransform
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:glue:getMLTransform', __args__, opts=opts, typ=GetMLTransformResult).value

    return AwaitableGetMLTransformResult(
        description=__ret__.description,
        glue_version=__ret__.glue_version,
        id=__ret__.id,
        max_capacity=__ret__.max_capacity,
        max_retries=__ret__.max_retries,
        name=__ret__.name,
        number_of_workers=__ret__.number_of_workers,
        role=__ret__.role,
        tags=__ret__.tags,
        timeout=__ret__.timeout,
        transform_encryption=__ret__.transform_encryption,
        transform_parameters=__ret__.transform_parameters,
        worker_type=__ret__.worker_type)


@_utilities.lift_output_func(get_ml_transform)
def get_ml_transform_output(id: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMLTransformResult]:
    """
    Resource Type definition for AWS::Glue::MLTransform
    """
    ...
