# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ProjectPlacementTemplate',
]

@pulumi.output_type
class ProjectPlacementTemplate(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "defaultAttributes":
            suggest = "default_attributes"
        elif key == "deviceTemplates":
            suggest = "device_templates"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProjectPlacementTemplate. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProjectPlacementTemplate.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProjectPlacementTemplate.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 default_attributes: Optional[Any] = None,
                 device_templates: Optional[Any] = None):
        if default_attributes is not None:
            pulumi.set(__self__, "default_attributes", default_attributes)
        if device_templates is not None:
            pulumi.set(__self__, "device_templates", device_templates)

    @property
    @pulumi.getter(name="defaultAttributes")
    def default_attributes(self) -> Optional[Any]:
        return pulumi.get(self, "default_attributes")

    @property
    @pulumi.getter(name="deviceTemplates")
    def device_templates(self) -> Optional[Any]:
        return pulumi.get(self, "device_templates")


