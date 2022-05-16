# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'FlowTemplateDefinitionDocumentArgs',
]

@pulumi.input_type
class FlowTemplateDefinitionDocumentArgs:
    def __init__(__self__, *,
                 language: pulumi.Input[str],
                 text: pulumi.Input[str]):
        pulumi.set(__self__, "language", language)
        pulumi.set(__self__, "text", text)

    @property
    @pulumi.getter
    def language(self) -> pulumi.Input[str]:
        return pulumi.get(self, "language")

    @language.setter
    def language(self, value: pulumi.Input[str]):
        pulumi.set(self, "language", value)

    @property
    @pulumi.getter
    def text(self) -> pulumi.Input[str]:
        return pulumi.get(self, "text")

    @text.setter
    def text(self, value: pulumi.Input[str]):
        pulumi.set(self, "text", value)

