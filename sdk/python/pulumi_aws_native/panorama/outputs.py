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
    'ApplicationInstanceManifestOverridesPayload',
    'ApplicationInstanceManifestPayload',
    'ApplicationInstanceTag',
    'PackageStorageLocation',
    'PackageTag',
]

@pulumi.output_type
class ApplicationInstanceManifestOverridesPayload(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "payloadData":
            suggest = "payload_data"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApplicationInstanceManifestOverridesPayload. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApplicationInstanceManifestOverridesPayload.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApplicationInstanceManifestOverridesPayload.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 payload_data: Optional[str] = None):
        if payload_data is not None:
            pulumi.set(__self__, "payload_data", payload_data)

    @property
    @pulumi.getter(name="payloadData")
    def payload_data(self) -> Optional[str]:
        return pulumi.get(self, "payload_data")


@pulumi.output_type
class ApplicationInstanceManifestPayload(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "payloadData":
            suggest = "payload_data"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApplicationInstanceManifestPayload. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApplicationInstanceManifestPayload.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApplicationInstanceManifestPayload.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 payload_data: Optional[str] = None):
        if payload_data is not None:
            pulumi.set(__self__, "payload_data", payload_data)

    @property
    @pulumi.getter(name="payloadData")
    def payload_data(self) -> Optional[str]:
        return pulumi.get(self, "payload_data")


@pulumi.output_type
class ApplicationInstanceTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        :param str key: A string used to identify this tag
        :param str value: A string containing the value for the tag
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        A string used to identify this tag
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        A string containing the value for the tag
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class PackageStorageLocation(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "binaryPrefixLocation":
            suggest = "binary_prefix_location"
        elif key == "generatedPrefixLocation":
            suggest = "generated_prefix_location"
        elif key == "manifestPrefixLocation":
            suggest = "manifest_prefix_location"
        elif key == "repoPrefixLocation":
            suggest = "repo_prefix_location"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PackageStorageLocation. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PackageStorageLocation.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PackageStorageLocation.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 binary_prefix_location: Optional[str] = None,
                 bucket: Optional[str] = None,
                 generated_prefix_location: Optional[str] = None,
                 manifest_prefix_location: Optional[str] = None,
                 repo_prefix_location: Optional[str] = None):
        if binary_prefix_location is not None:
            pulumi.set(__self__, "binary_prefix_location", binary_prefix_location)
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if generated_prefix_location is not None:
            pulumi.set(__self__, "generated_prefix_location", generated_prefix_location)
        if manifest_prefix_location is not None:
            pulumi.set(__self__, "manifest_prefix_location", manifest_prefix_location)
        if repo_prefix_location is not None:
            pulumi.set(__self__, "repo_prefix_location", repo_prefix_location)

    @property
    @pulumi.getter(name="binaryPrefixLocation")
    def binary_prefix_location(self) -> Optional[str]:
        return pulumi.get(self, "binary_prefix_location")

    @property
    @pulumi.getter
    def bucket(self) -> Optional[str]:
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter(name="generatedPrefixLocation")
    def generated_prefix_location(self) -> Optional[str]:
        return pulumi.get(self, "generated_prefix_location")

    @property
    @pulumi.getter(name="manifestPrefixLocation")
    def manifest_prefix_location(self) -> Optional[str]:
        return pulumi.get(self, "manifest_prefix_location")

    @property
    @pulumi.getter(name="repoPrefixLocation")
    def repo_prefix_location(self) -> Optional[str]:
        return pulumi.get(self, "repo_prefix_location")


@pulumi.output_type
class PackageTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


