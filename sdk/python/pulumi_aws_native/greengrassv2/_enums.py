# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ComponentVersionLambdaEventSourceType',
    'ComponentVersionLambdaExecutionParametersInputPayloadEncodingType',
    'ComponentVersionLambdaFilesystemPermission',
    'ComponentVersionLambdaLinuxProcessParamsIsolationMode',
]


class ComponentVersionLambdaEventSourceType(str, Enum):
    PUB_SUB = "PUB_SUB"
    IOT_CORE = "IOT_CORE"


class ComponentVersionLambdaExecutionParametersInputPayloadEncodingType(str, Enum):
    JSON = "json"
    BINARY = "binary"


class ComponentVersionLambdaFilesystemPermission(str, Enum):
    RO = "ro"
    RW = "rw"


class ComponentVersionLambdaLinuxProcessParamsIsolationMode(str, Enum):
    GREENGRASS_CONTAINER = "GreengrassContainer"
    NO_CONTAINER = "NoContainer"
