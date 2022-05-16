# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ComponentBindingPropertiesArgs',
    'ComponentChildArgs',
    'ComponentCollectionPropertiesArgs',
    'ComponentEventsArgs',
    'ComponentOverridesArgs',
    'ComponentPropertiesArgs',
    'ComponentTagsArgs',
    'ComponentVariantValuesArgs',
    'ComponentVariantArgs',
    'ThemeTagsArgs',
    'ThemeValuesArgs',
    'ThemeValueArgs',
]

@pulumi.input_type
class ComponentBindingPropertiesArgs:
    def __init__(__self__):
        pass


@pulumi.input_type
class ComponentChildArgs:
    def __init__(__self__, *,
                 component_type: pulumi.Input[str],
                 name: pulumi.Input[str],
                 properties: pulumi.Input['ComponentPropertiesArgs'],
                 children: Optional[pulumi.Input[Sequence[pulumi.Input['ComponentChildArgs']]]] = None,
                 events: Optional[pulumi.Input['ComponentEventsArgs']] = None):
        pulumi.set(__self__, "component_type", component_type)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "properties", properties)
        if children is not None:
            pulumi.set(__self__, "children", children)
        if events is not None:
            pulumi.set(__self__, "events", events)

    @property
    @pulumi.getter(name="componentType")
    def component_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "component_type")

    @component_type.setter
    def component_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "component_type", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Input['ComponentPropertiesArgs']:
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: pulumi.Input['ComponentPropertiesArgs']):
        pulumi.set(self, "properties", value)

    @property
    @pulumi.getter
    def children(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ComponentChildArgs']]]]:
        return pulumi.get(self, "children")

    @children.setter
    def children(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ComponentChildArgs']]]]):
        pulumi.set(self, "children", value)

    @property
    @pulumi.getter
    def events(self) -> Optional[pulumi.Input['ComponentEventsArgs']]:
        return pulumi.get(self, "events")

    @events.setter
    def events(self, value: Optional[pulumi.Input['ComponentEventsArgs']]):
        pulumi.set(self, "events", value)


@pulumi.input_type
class ComponentCollectionPropertiesArgs:
    def __init__(__self__):
        pass


@pulumi.input_type
class ComponentEventsArgs:
    def __init__(__self__):
        pass


@pulumi.input_type
class ComponentOverridesArgs:
    def __init__(__self__):
        pass


@pulumi.input_type
class ComponentPropertiesArgs:
    def __init__(__self__):
        pass


@pulumi.input_type
class ComponentTagsArgs:
    def __init__(__self__):
        pass


@pulumi.input_type
class ComponentVariantValuesArgs:
    def __init__(__self__):
        pass


@pulumi.input_type
class ComponentVariantArgs:
    def __init__(__self__, *,
                 overrides: Optional[pulumi.Input['ComponentOverridesArgs']] = None,
                 variant_values: Optional[pulumi.Input['ComponentVariantValuesArgs']] = None):
        if overrides is not None:
            pulumi.set(__self__, "overrides", overrides)
        if variant_values is not None:
            pulumi.set(__self__, "variant_values", variant_values)

    @property
    @pulumi.getter
    def overrides(self) -> Optional[pulumi.Input['ComponentOverridesArgs']]:
        return pulumi.get(self, "overrides")

    @overrides.setter
    def overrides(self, value: Optional[pulumi.Input['ComponentOverridesArgs']]):
        pulumi.set(self, "overrides", value)

    @property
    @pulumi.getter(name="variantValues")
    def variant_values(self) -> Optional[pulumi.Input['ComponentVariantValuesArgs']]:
        return pulumi.get(self, "variant_values")

    @variant_values.setter
    def variant_values(self, value: Optional[pulumi.Input['ComponentVariantValuesArgs']]):
        pulumi.set(self, "variant_values", value)


@pulumi.input_type
class ThemeTagsArgs:
    def __init__(__self__):
        pass


@pulumi.input_type
class ThemeValuesArgs:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input['ThemeValueArgs']] = None):
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
    def value(self) -> Optional[pulumi.Input['ThemeValueArgs']]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input['ThemeValueArgs']]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ThemeValueArgs:
    def __init__(__self__, *,
                 children: Optional[pulumi.Input[Sequence[pulumi.Input['ThemeValuesArgs']]]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        if children is not None:
            pulumi.set(__self__, "children", children)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def children(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ThemeValuesArgs']]]]:
        return pulumi.get(self, "children")

    @children.setter
    def children(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ThemeValuesArgs']]]]):
        pulumi.set(self, "children", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)

