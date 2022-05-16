# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'CloudFormationProvisionedProductAcceptLanguage',
    'CloudFormationProvisionedProductProvisioningPreferencesStackSetOperationType',
    'ServiceActionAcceptLanguage',
    'ServiceActionDefinitionType',
]


class CloudFormationProvisionedProductAcceptLanguage(str, Enum):
    EN = "en"
    JP = "jp"
    ZH = "zh"


class CloudFormationProvisionedProductProvisioningPreferencesStackSetOperationType(str, Enum):
    CREATE = "CREATE"
    UPDATE = "UPDATE"
    DELETE = "DELETE"


class ServiceActionAcceptLanguage(str, Enum):
    EN = "en"
    JP = "jp"
    ZH = "zh"


class ServiceActionDefinitionType(str, Enum):
    SSM_AUTOMATION = "SSM_AUTOMATION"
