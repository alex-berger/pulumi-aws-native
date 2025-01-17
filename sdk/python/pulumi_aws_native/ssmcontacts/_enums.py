# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ContactChannelChannelType',
    'ContactType',
]


class ContactChannelChannelType(str, Enum):
    """
    Device type, which specify notification channel. Currently supported values: “SMS”, “VOICE”, “EMAIL”, “CHATBOT.
    """
    SMS = "SMS"
    VOICE = "VOICE"
    EMAIL = "EMAIL"


class ContactType(str, Enum):
    """
    Contact type, which specify type of contact. Currently supported values: “PERSONAL”, “SHARED”, “OTHER“.
    """
    PERSONAL = "PERSONAL"
    CUSTOM = "CUSTOM"
    SERVICE = "SERVICE"
    ESCALATION = "ESCALATION"
