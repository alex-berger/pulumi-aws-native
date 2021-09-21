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
    'LifecyclePolicyAction',
    'LifecyclePolicyCreateRule',
    'LifecyclePolicyCrossRegionCopyAction',
    'LifecyclePolicyCrossRegionCopyDeprecateRule',
    'LifecyclePolicyCrossRegionCopyRetainRule',
    'LifecyclePolicyCrossRegionCopyRule',
    'LifecyclePolicyDeprecateRule',
    'LifecyclePolicyEncryptionConfiguration',
    'LifecyclePolicyEventParameters',
    'LifecyclePolicyEventSource',
    'LifecyclePolicyFastRestoreRule',
    'LifecyclePolicyParameters',
    'LifecyclePolicyPolicyDetails',
    'LifecyclePolicyRetainRule',
    'LifecyclePolicySchedule',
    'LifecyclePolicyShareRule',
    'LifecyclePolicyTag',
]

@pulumi.output_type
class LifecyclePolicyAction(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "crossRegionCopy":
            suggest = "cross_region_copy"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LifecyclePolicyAction. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LifecyclePolicyAction.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LifecyclePolicyAction.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cross_region_copy: Sequence['outputs.LifecyclePolicyCrossRegionCopyAction'],
                 name: str):
        pulumi.set(__self__, "cross_region_copy", cross_region_copy)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="crossRegionCopy")
    def cross_region_copy(self) -> Sequence['outputs.LifecyclePolicyCrossRegionCopyAction']:
        return pulumi.get(self, "cross_region_copy")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")


@pulumi.output_type
class LifecyclePolicyCreateRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "cronExpression":
            suggest = "cron_expression"
        elif key == "intervalUnit":
            suggest = "interval_unit"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LifecyclePolicyCreateRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LifecyclePolicyCreateRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LifecyclePolicyCreateRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cron_expression: Optional[str] = None,
                 interval: Optional[int] = None,
                 interval_unit: Optional[str] = None,
                 location: Optional[str] = None,
                 times: Optional[Sequence[str]] = None):
        if cron_expression is not None:
            pulumi.set(__self__, "cron_expression", cron_expression)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if interval_unit is not None:
            pulumi.set(__self__, "interval_unit", interval_unit)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if times is not None:
            pulumi.set(__self__, "times", times)

    @property
    @pulumi.getter(name="cronExpression")
    def cron_expression(self) -> Optional[str]:
        return pulumi.get(self, "cron_expression")

    @property
    @pulumi.getter
    def interval(self) -> Optional[int]:
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter(name="intervalUnit")
    def interval_unit(self) -> Optional[str]:
        return pulumi.get(self, "interval_unit")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def times(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "times")


@pulumi.output_type
class LifecyclePolicyCrossRegionCopyAction(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "encryptionConfiguration":
            suggest = "encryption_configuration"
        elif key == "retainRule":
            suggest = "retain_rule"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LifecyclePolicyCrossRegionCopyAction. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LifecyclePolicyCrossRegionCopyAction.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LifecyclePolicyCrossRegionCopyAction.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 encryption_configuration: 'outputs.LifecyclePolicyEncryptionConfiguration',
                 target: str,
                 retain_rule: Optional['outputs.LifecyclePolicyCrossRegionCopyRetainRule'] = None):
        pulumi.set(__self__, "encryption_configuration", encryption_configuration)
        pulumi.set(__self__, "target", target)
        if retain_rule is not None:
            pulumi.set(__self__, "retain_rule", retain_rule)

    @property
    @pulumi.getter(name="encryptionConfiguration")
    def encryption_configuration(self) -> 'outputs.LifecyclePolicyEncryptionConfiguration':
        return pulumi.get(self, "encryption_configuration")

    @property
    @pulumi.getter
    def target(self) -> str:
        return pulumi.get(self, "target")

    @property
    @pulumi.getter(name="retainRule")
    def retain_rule(self) -> Optional['outputs.LifecyclePolicyCrossRegionCopyRetainRule']:
        return pulumi.get(self, "retain_rule")


@pulumi.output_type
class LifecyclePolicyCrossRegionCopyDeprecateRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "intervalUnit":
            suggest = "interval_unit"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LifecyclePolicyCrossRegionCopyDeprecateRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LifecyclePolicyCrossRegionCopyDeprecateRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LifecyclePolicyCrossRegionCopyDeprecateRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 interval: int,
                 interval_unit: str):
        pulumi.set(__self__, "interval", interval)
        pulumi.set(__self__, "interval_unit", interval_unit)

    @property
    @pulumi.getter
    def interval(self) -> int:
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter(name="intervalUnit")
    def interval_unit(self) -> str:
        return pulumi.get(self, "interval_unit")


@pulumi.output_type
class LifecyclePolicyCrossRegionCopyRetainRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "intervalUnit":
            suggest = "interval_unit"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LifecyclePolicyCrossRegionCopyRetainRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LifecyclePolicyCrossRegionCopyRetainRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LifecyclePolicyCrossRegionCopyRetainRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 interval: int,
                 interval_unit: str):
        pulumi.set(__self__, "interval", interval)
        pulumi.set(__self__, "interval_unit", interval_unit)

    @property
    @pulumi.getter
    def interval(self) -> int:
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter(name="intervalUnit")
    def interval_unit(self) -> str:
        return pulumi.get(self, "interval_unit")


@pulumi.output_type
class LifecyclePolicyCrossRegionCopyRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "cmkArn":
            suggest = "cmk_arn"
        elif key == "copyTags":
            suggest = "copy_tags"
        elif key == "deprecateRule":
            suggest = "deprecate_rule"
        elif key == "retainRule":
            suggest = "retain_rule"
        elif key == "targetRegion":
            suggest = "target_region"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LifecyclePolicyCrossRegionCopyRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LifecyclePolicyCrossRegionCopyRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LifecyclePolicyCrossRegionCopyRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 encrypted: bool,
                 cmk_arn: Optional[str] = None,
                 copy_tags: Optional[bool] = None,
                 deprecate_rule: Optional['outputs.LifecyclePolicyCrossRegionCopyDeprecateRule'] = None,
                 retain_rule: Optional['outputs.LifecyclePolicyCrossRegionCopyRetainRule'] = None,
                 target: Optional[str] = None,
                 target_region: Optional[str] = None):
        pulumi.set(__self__, "encrypted", encrypted)
        if cmk_arn is not None:
            pulumi.set(__self__, "cmk_arn", cmk_arn)
        if copy_tags is not None:
            pulumi.set(__self__, "copy_tags", copy_tags)
        if deprecate_rule is not None:
            pulumi.set(__self__, "deprecate_rule", deprecate_rule)
        if retain_rule is not None:
            pulumi.set(__self__, "retain_rule", retain_rule)
        if target is not None:
            pulumi.set(__self__, "target", target)
        if target_region is not None:
            pulumi.set(__self__, "target_region", target_region)

    @property
    @pulumi.getter
    def encrypted(self) -> bool:
        return pulumi.get(self, "encrypted")

    @property
    @pulumi.getter(name="cmkArn")
    def cmk_arn(self) -> Optional[str]:
        return pulumi.get(self, "cmk_arn")

    @property
    @pulumi.getter(name="copyTags")
    def copy_tags(self) -> Optional[bool]:
        return pulumi.get(self, "copy_tags")

    @property
    @pulumi.getter(name="deprecateRule")
    def deprecate_rule(self) -> Optional['outputs.LifecyclePolicyCrossRegionCopyDeprecateRule']:
        return pulumi.get(self, "deprecate_rule")

    @property
    @pulumi.getter(name="retainRule")
    def retain_rule(self) -> Optional['outputs.LifecyclePolicyCrossRegionCopyRetainRule']:
        return pulumi.get(self, "retain_rule")

    @property
    @pulumi.getter
    def target(self) -> Optional[str]:
        return pulumi.get(self, "target")

    @property
    @pulumi.getter(name="targetRegion")
    def target_region(self) -> Optional[str]:
        return pulumi.get(self, "target_region")


@pulumi.output_type
class LifecyclePolicyDeprecateRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "intervalUnit":
            suggest = "interval_unit"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LifecyclePolicyDeprecateRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LifecyclePolicyDeprecateRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LifecyclePolicyDeprecateRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 count: Optional[int] = None,
                 interval: Optional[int] = None,
                 interval_unit: Optional[str] = None):
        if count is not None:
            pulumi.set(__self__, "count", count)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if interval_unit is not None:
            pulumi.set(__self__, "interval_unit", interval_unit)

    @property
    @pulumi.getter
    def count(self) -> Optional[int]:
        return pulumi.get(self, "count")

    @property
    @pulumi.getter
    def interval(self) -> Optional[int]:
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter(name="intervalUnit")
    def interval_unit(self) -> Optional[str]:
        return pulumi.get(self, "interval_unit")


@pulumi.output_type
class LifecyclePolicyEncryptionConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "cmkArn":
            suggest = "cmk_arn"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LifecyclePolicyEncryptionConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LifecyclePolicyEncryptionConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LifecyclePolicyEncryptionConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 encrypted: bool,
                 cmk_arn: Optional[str] = None):
        pulumi.set(__self__, "encrypted", encrypted)
        if cmk_arn is not None:
            pulumi.set(__self__, "cmk_arn", cmk_arn)

    @property
    @pulumi.getter
    def encrypted(self) -> bool:
        return pulumi.get(self, "encrypted")

    @property
    @pulumi.getter(name="cmkArn")
    def cmk_arn(self) -> Optional[str]:
        return pulumi.get(self, "cmk_arn")


@pulumi.output_type
class LifecyclePolicyEventParameters(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "eventType":
            suggest = "event_type"
        elif key == "snapshotOwner":
            suggest = "snapshot_owner"
        elif key == "descriptionRegex":
            suggest = "description_regex"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LifecyclePolicyEventParameters. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LifecyclePolicyEventParameters.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LifecyclePolicyEventParameters.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 event_type: str,
                 snapshot_owner: Sequence[str],
                 description_regex: Optional[str] = None):
        pulumi.set(__self__, "event_type", event_type)
        pulumi.set(__self__, "snapshot_owner", snapshot_owner)
        if description_regex is not None:
            pulumi.set(__self__, "description_regex", description_regex)

    @property
    @pulumi.getter(name="eventType")
    def event_type(self) -> str:
        return pulumi.get(self, "event_type")

    @property
    @pulumi.getter(name="snapshotOwner")
    def snapshot_owner(self) -> Sequence[str]:
        return pulumi.get(self, "snapshot_owner")

    @property
    @pulumi.getter(name="descriptionRegex")
    def description_regex(self) -> Optional[str]:
        return pulumi.get(self, "description_regex")


@pulumi.output_type
class LifecyclePolicyEventSource(dict):
    def __init__(__self__, *,
                 type: str,
                 parameters: Optional['outputs.LifecyclePolicyEventParameters'] = None):
        pulumi.set(__self__, "type", type)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def parameters(self) -> Optional['outputs.LifecyclePolicyEventParameters']:
        return pulumi.get(self, "parameters")


@pulumi.output_type
class LifecyclePolicyFastRestoreRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "availabilityZones":
            suggest = "availability_zones"
        elif key == "intervalUnit":
            suggest = "interval_unit"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LifecyclePolicyFastRestoreRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LifecyclePolicyFastRestoreRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LifecyclePolicyFastRestoreRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 availability_zones: Optional[Sequence[str]] = None,
                 count: Optional[int] = None,
                 interval: Optional[int] = None,
                 interval_unit: Optional[str] = None):
        if availability_zones is not None:
            pulumi.set(__self__, "availability_zones", availability_zones)
        if count is not None:
            pulumi.set(__self__, "count", count)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if interval_unit is not None:
            pulumi.set(__self__, "interval_unit", interval_unit)

    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "availability_zones")

    @property
    @pulumi.getter
    def count(self) -> Optional[int]:
        return pulumi.get(self, "count")

    @property
    @pulumi.getter
    def interval(self) -> Optional[int]:
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter(name="intervalUnit")
    def interval_unit(self) -> Optional[str]:
        return pulumi.get(self, "interval_unit")


@pulumi.output_type
class LifecyclePolicyParameters(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "excludeBootVolume":
            suggest = "exclude_boot_volume"
        elif key == "noReboot":
            suggest = "no_reboot"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LifecyclePolicyParameters. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LifecyclePolicyParameters.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LifecyclePolicyParameters.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 exclude_boot_volume: Optional[bool] = None,
                 no_reboot: Optional[bool] = None):
        if exclude_boot_volume is not None:
            pulumi.set(__self__, "exclude_boot_volume", exclude_boot_volume)
        if no_reboot is not None:
            pulumi.set(__self__, "no_reboot", no_reboot)

    @property
    @pulumi.getter(name="excludeBootVolume")
    def exclude_boot_volume(self) -> Optional[bool]:
        return pulumi.get(self, "exclude_boot_volume")

    @property
    @pulumi.getter(name="noReboot")
    def no_reboot(self) -> Optional[bool]:
        return pulumi.get(self, "no_reboot")


@pulumi.output_type
class LifecyclePolicyPolicyDetails(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "eventSource":
            suggest = "event_source"
        elif key == "policyType":
            suggest = "policy_type"
        elif key == "resourceLocations":
            suggest = "resource_locations"
        elif key == "resourceTypes":
            suggest = "resource_types"
        elif key == "targetTags":
            suggest = "target_tags"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LifecyclePolicyPolicyDetails. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LifecyclePolicyPolicyDetails.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LifecyclePolicyPolicyDetails.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 actions: Optional[Sequence['outputs.LifecyclePolicyAction']] = None,
                 event_source: Optional['outputs.LifecyclePolicyEventSource'] = None,
                 parameters: Optional['outputs.LifecyclePolicyParameters'] = None,
                 policy_type: Optional[str] = None,
                 resource_locations: Optional[Sequence[str]] = None,
                 resource_types: Optional[Sequence[str]] = None,
                 schedules: Optional[Sequence['outputs.LifecyclePolicySchedule']] = None,
                 target_tags: Optional[Sequence['outputs.LifecyclePolicyTag']] = None):
        if actions is not None:
            pulumi.set(__self__, "actions", actions)
        if event_source is not None:
            pulumi.set(__self__, "event_source", event_source)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if policy_type is not None:
            pulumi.set(__self__, "policy_type", policy_type)
        if resource_locations is not None:
            pulumi.set(__self__, "resource_locations", resource_locations)
        if resource_types is not None:
            pulumi.set(__self__, "resource_types", resource_types)
        if schedules is not None:
            pulumi.set(__self__, "schedules", schedules)
        if target_tags is not None:
            pulumi.set(__self__, "target_tags", target_tags)

    @property
    @pulumi.getter
    def actions(self) -> Optional[Sequence['outputs.LifecyclePolicyAction']]:
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter(name="eventSource")
    def event_source(self) -> Optional['outputs.LifecyclePolicyEventSource']:
        return pulumi.get(self, "event_source")

    @property
    @pulumi.getter
    def parameters(self) -> Optional['outputs.LifecyclePolicyParameters']:
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> Optional[str]:
        return pulumi.get(self, "policy_type")

    @property
    @pulumi.getter(name="resourceLocations")
    def resource_locations(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "resource_locations")

    @property
    @pulumi.getter(name="resourceTypes")
    def resource_types(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "resource_types")

    @property
    @pulumi.getter
    def schedules(self) -> Optional[Sequence['outputs.LifecyclePolicySchedule']]:
        return pulumi.get(self, "schedules")

    @property
    @pulumi.getter(name="targetTags")
    def target_tags(self) -> Optional[Sequence['outputs.LifecyclePolicyTag']]:
        return pulumi.get(self, "target_tags")


@pulumi.output_type
class LifecyclePolicyRetainRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "intervalUnit":
            suggest = "interval_unit"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LifecyclePolicyRetainRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LifecyclePolicyRetainRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LifecyclePolicyRetainRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 count: Optional[int] = None,
                 interval: Optional[int] = None,
                 interval_unit: Optional[str] = None):
        if count is not None:
            pulumi.set(__self__, "count", count)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if interval_unit is not None:
            pulumi.set(__self__, "interval_unit", interval_unit)

    @property
    @pulumi.getter
    def count(self) -> Optional[int]:
        return pulumi.get(self, "count")

    @property
    @pulumi.getter
    def interval(self) -> Optional[int]:
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter(name="intervalUnit")
    def interval_unit(self) -> Optional[str]:
        return pulumi.get(self, "interval_unit")


@pulumi.output_type
class LifecyclePolicySchedule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "copyTags":
            suggest = "copy_tags"
        elif key == "createRule":
            suggest = "create_rule"
        elif key == "crossRegionCopyRules":
            suggest = "cross_region_copy_rules"
        elif key == "deprecateRule":
            suggest = "deprecate_rule"
        elif key == "fastRestoreRule":
            suggest = "fast_restore_rule"
        elif key == "retainRule":
            suggest = "retain_rule"
        elif key == "shareRules":
            suggest = "share_rules"
        elif key == "tagsToAdd":
            suggest = "tags_to_add"
        elif key == "variableTags":
            suggest = "variable_tags"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LifecyclePolicySchedule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LifecyclePolicySchedule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LifecyclePolicySchedule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 copy_tags: Optional[bool] = None,
                 create_rule: Optional['outputs.LifecyclePolicyCreateRule'] = None,
                 cross_region_copy_rules: Optional[Sequence['outputs.LifecyclePolicyCrossRegionCopyRule']] = None,
                 deprecate_rule: Optional['outputs.LifecyclePolicyDeprecateRule'] = None,
                 fast_restore_rule: Optional['outputs.LifecyclePolicyFastRestoreRule'] = None,
                 name: Optional[str] = None,
                 retain_rule: Optional['outputs.LifecyclePolicyRetainRule'] = None,
                 share_rules: Optional[Sequence['outputs.LifecyclePolicyShareRule']] = None,
                 tags_to_add: Optional[Sequence['outputs.LifecyclePolicyTag']] = None,
                 variable_tags: Optional[Sequence['outputs.LifecyclePolicyTag']] = None):
        if copy_tags is not None:
            pulumi.set(__self__, "copy_tags", copy_tags)
        if create_rule is not None:
            pulumi.set(__self__, "create_rule", create_rule)
        if cross_region_copy_rules is not None:
            pulumi.set(__self__, "cross_region_copy_rules", cross_region_copy_rules)
        if deprecate_rule is not None:
            pulumi.set(__self__, "deprecate_rule", deprecate_rule)
        if fast_restore_rule is not None:
            pulumi.set(__self__, "fast_restore_rule", fast_restore_rule)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if retain_rule is not None:
            pulumi.set(__self__, "retain_rule", retain_rule)
        if share_rules is not None:
            pulumi.set(__self__, "share_rules", share_rules)
        if tags_to_add is not None:
            pulumi.set(__self__, "tags_to_add", tags_to_add)
        if variable_tags is not None:
            pulumi.set(__self__, "variable_tags", variable_tags)

    @property
    @pulumi.getter(name="copyTags")
    def copy_tags(self) -> Optional[bool]:
        return pulumi.get(self, "copy_tags")

    @property
    @pulumi.getter(name="createRule")
    def create_rule(self) -> Optional['outputs.LifecyclePolicyCreateRule']:
        return pulumi.get(self, "create_rule")

    @property
    @pulumi.getter(name="crossRegionCopyRules")
    def cross_region_copy_rules(self) -> Optional[Sequence['outputs.LifecyclePolicyCrossRegionCopyRule']]:
        return pulumi.get(self, "cross_region_copy_rules")

    @property
    @pulumi.getter(name="deprecateRule")
    def deprecate_rule(self) -> Optional['outputs.LifecyclePolicyDeprecateRule']:
        return pulumi.get(self, "deprecate_rule")

    @property
    @pulumi.getter(name="fastRestoreRule")
    def fast_restore_rule(self) -> Optional['outputs.LifecyclePolicyFastRestoreRule']:
        return pulumi.get(self, "fast_restore_rule")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="retainRule")
    def retain_rule(self) -> Optional['outputs.LifecyclePolicyRetainRule']:
        return pulumi.get(self, "retain_rule")

    @property
    @pulumi.getter(name="shareRules")
    def share_rules(self) -> Optional[Sequence['outputs.LifecyclePolicyShareRule']]:
        return pulumi.get(self, "share_rules")

    @property
    @pulumi.getter(name="tagsToAdd")
    def tags_to_add(self) -> Optional[Sequence['outputs.LifecyclePolicyTag']]:
        return pulumi.get(self, "tags_to_add")

    @property
    @pulumi.getter(name="variableTags")
    def variable_tags(self) -> Optional[Sequence['outputs.LifecyclePolicyTag']]:
        return pulumi.get(self, "variable_tags")


@pulumi.output_type
class LifecyclePolicyShareRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "targetAccounts":
            suggest = "target_accounts"
        elif key == "unshareInterval":
            suggest = "unshare_interval"
        elif key == "unshareIntervalUnit":
            suggest = "unshare_interval_unit"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LifecyclePolicyShareRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LifecyclePolicyShareRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LifecyclePolicyShareRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 target_accounts: Optional[Sequence[str]] = None,
                 unshare_interval: Optional[int] = None,
                 unshare_interval_unit: Optional[str] = None):
        if target_accounts is not None:
            pulumi.set(__self__, "target_accounts", target_accounts)
        if unshare_interval is not None:
            pulumi.set(__self__, "unshare_interval", unshare_interval)
        if unshare_interval_unit is not None:
            pulumi.set(__self__, "unshare_interval_unit", unshare_interval_unit)

    @property
    @pulumi.getter(name="targetAccounts")
    def target_accounts(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "target_accounts")

    @property
    @pulumi.getter(name="unshareInterval")
    def unshare_interval(self) -> Optional[int]:
        return pulumi.get(self, "unshare_interval")

    @property
    @pulumi.getter(name="unshareIntervalUnit")
    def unshare_interval_unit(self) -> Optional[str]:
        return pulumi.get(self, "unshare_interval_unit")


@pulumi.output_type
class LifecyclePolicyTag(dict):
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

