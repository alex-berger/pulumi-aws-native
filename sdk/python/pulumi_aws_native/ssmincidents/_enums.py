# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ResponsePlanSsmAutomationTargetAccount',
    'ResponsePlanVariableType',
]


class ResponsePlanSsmAutomationTargetAccount(str, Enum):
    """
    The account type to use when starting the SSM automation document.
    """
    IMPACTED_ACCOUNT = "IMPACTED_ACCOUNT"
    RESPONSE_PLAN_OWNER_ACCOUNT = "RESPONSE_PLAN_OWNER_ACCOUNT"


class ResponsePlanVariableType(str, Enum):
    """
    The variable types used as dynamic parameter value when starting the SSM automation document.
    """
    INCIDENT_RECORD_ARN = "INCIDENT_RECORD_ARN"
    INVOLVED_RESOURCES = "INVOLVED_RESOURCES"
