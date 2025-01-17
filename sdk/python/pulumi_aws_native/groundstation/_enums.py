# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ConfigBandwidthUnits',
    'ConfigEirpUnits',
    'ConfigFrequencyUnits',
    'ConfigPolarization',
    'ConfigTrackingConfigAutotrack',
]


class ConfigBandwidthUnits(str, Enum):
    G_HZ = "GHz"
    M_HZ = "MHz"
    K_HZ = "kHz"


class ConfigEirpUnits(str, Enum):
    DBW = "dBW"


class ConfigFrequencyUnits(str, Enum):
    G_HZ = "GHz"
    M_HZ = "MHz"
    K_HZ = "kHz"


class ConfigPolarization(str, Enum):
    LEFT_HAND = "LEFT_HAND"
    RIGHT_HAND = "RIGHT_HAND"
    NONE = "NONE"


class ConfigTrackingConfigAutotrack(str, Enum):
    REQUIRED = "REQUIRED"
    PREFERRED = "PREFERRED"
    REMOVED = "REMOVED"
