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
    'MapConfigurationArgs',
    'PlaceIndexDataSourceConfigurationArgs',
]

@pulumi.input_type
class MapConfigurationArgs:
    def __init__(__self__, *,
                 style: pulumi.Input[str]):
        pulumi.set(__self__, "style", style)

    @property
    @pulumi.getter
    def style(self) -> pulumi.Input[str]:
        return pulumi.get(self, "style")

    @style.setter
    def style(self, value: pulumi.Input[str]):
        pulumi.set(self, "style", value)


@pulumi.input_type
class PlaceIndexDataSourceConfigurationArgs:
    def __init__(__self__, *,
                 intended_use: Optional[pulumi.Input['PlaceIndexIntendedUse']] = None):
        if intended_use is not None:
            pulumi.set(__self__, "intended_use", intended_use)

    @property
    @pulumi.getter(name="intendedUse")
    def intended_use(self) -> Optional[pulumi.Input['PlaceIndexIntendedUse']]:
        return pulumi.get(self, "intended_use")

    @intended_use.setter
    def intended_use(self, value: Optional[pulumi.Input['PlaceIndexIntendedUse']]):
        pulumi.set(self, "intended_use", value)

