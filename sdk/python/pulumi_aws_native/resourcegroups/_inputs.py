# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'GroupConfigurationItemArgs',
    'GroupConfigurationParameterArgs',
    'GroupQueryArgs',
    'GroupResourceQueryArgs',
    'GroupTagFilterArgs',
    'GroupTagArgs',
]

@pulumi.input_type
class GroupConfigurationItemArgs:
    def __init__(__self__, *,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input['GroupConfigurationParameterArgs']]]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GroupConfigurationParameterArgs']]]]:
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GroupConfigurationParameterArgs']]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class GroupConfigurationParameterArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        if name is not None:
            pulumi.set(__self__, "name", name)
        if values is not None:
            pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "values", value)


@pulumi.input_type
class GroupQueryArgs:
    def __init__(__self__, *,
                 resource_type_filters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 stack_identifier: Optional[pulumi.Input[str]] = None,
                 tag_filters: Optional[pulumi.Input[Sequence[pulumi.Input['GroupTagFilterArgs']]]] = None):
        if resource_type_filters is not None:
            pulumi.set(__self__, "resource_type_filters", resource_type_filters)
        if stack_identifier is not None:
            pulumi.set(__self__, "stack_identifier", stack_identifier)
        if tag_filters is not None:
            pulumi.set(__self__, "tag_filters", tag_filters)

    @property
    @pulumi.getter(name="resourceTypeFilters")
    def resource_type_filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "resource_type_filters")

    @resource_type_filters.setter
    def resource_type_filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "resource_type_filters", value)

    @property
    @pulumi.getter(name="stackIdentifier")
    def stack_identifier(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "stack_identifier")

    @stack_identifier.setter
    def stack_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stack_identifier", value)

    @property
    @pulumi.getter(name="tagFilters")
    def tag_filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GroupTagFilterArgs']]]]:
        return pulumi.get(self, "tag_filters")

    @tag_filters.setter
    def tag_filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GroupTagFilterArgs']]]]):
        pulumi.set(self, "tag_filters", value)


@pulumi.input_type
class GroupResourceQueryArgs:
    def __init__(__self__, *,
                 query: Optional[pulumi.Input['GroupQueryArgs']] = None,
                 type: Optional[pulumi.Input['GroupResourceQueryType']] = None):
        if query is not None:
            pulumi.set(__self__, "query", query)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def query(self) -> Optional[pulumi.Input['GroupQueryArgs']]:
        return pulumi.get(self, "query")

    @query.setter
    def query(self, value: Optional[pulumi.Input['GroupQueryArgs']]):
        pulumi.set(self, "query", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input['GroupResourceQueryType']]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input['GroupResourceQueryType']]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class GroupTagFilterArgs:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        if key is not None:
            pulumi.set(__self__, "key", key)
        if values is not None:
            pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "values", value)


@pulumi.input_type
class GroupTagArgs:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        if key is not None:
            pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


