# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'FindingsFilterCriterion',
    'FindingsFilterFindingCriteria',
    'FindingsFilterListItem',
]

@pulumi.output_type
class FindingsFilterCriterion(dict):
    """
    Map of filter criteria.
    """
    def __init__(__self__):
        """
        Map of filter criteria.
        """
        pass


@pulumi.output_type
class FindingsFilterFindingCriteria(dict):
    def __init__(__self__, *,
                 criterion: Optional['outputs.FindingsFilterCriterion'] = None):
        if criterion is not None:
            pulumi.set(__self__, "criterion", criterion)

    @property
    @pulumi.getter
    def criterion(self) -> Optional['outputs.FindingsFilterCriterion']:
        return pulumi.get(self, "criterion")


@pulumi.output_type
class FindingsFilterListItem(dict):
    """
    Returned by ListHandler representing filter name and ID.
    """
    def __init__(__self__, *,
                 id: Optional[str] = None,
                 name: Optional[str] = None):
        """
        Returned by ListHandler representing filter name and ID.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")


