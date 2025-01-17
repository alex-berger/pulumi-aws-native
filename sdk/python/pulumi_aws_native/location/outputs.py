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
    'MapConfiguration',
    'PlaceIndexDataSourceConfiguration',
]

@pulumi.output_type
class MapConfiguration(dict):
    def __init__(__self__, *,
                 style: str):
        pulumi.set(__self__, "style", style)

    @property
    @pulumi.getter
    def style(self) -> str:
        return pulumi.get(self, "style")


@pulumi.output_type
class PlaceIndexDataSourceConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "intendedUse":
            suggest = "intended_use"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PlaceIndexDataSourceConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PlaceIndexDataSourceConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PlaceIndexDataSourceConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 intended_use: Optional['PlaceIndexIntendedUse'] = None):
        if intended_use is not None:
            pulumi.set(__self__, "intended_use", intended_use)

    @property
    @pulumi.getter(name="intendedUse")
    def intended_use(self) -> Optional['PlaceIndexIntendedUse']:
        return pulumi.get(self, "intended_use")


