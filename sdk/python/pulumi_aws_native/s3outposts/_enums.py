# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'BucketRuleStatus',
    'EndpointAccessType',
    'EndpointStatus',
]


class BucketRuleStatus(str, Enum):
    ENABLED = "Enabled"
    DISABLED = "Disabled"


class EndpointAccessType(str, Enum):
    """
    The type of access for the on-premise network connectivity for the Outpost endpoint. To access endpoint from an on-premises network, you must specify the access type and provide the customer owned Ipv4 pool.
    """
    CUSTOMER_OWNED_IP = "CustomerOwnedIp"
    PRIVATE = "Private"


class EndpointStatus(str, Enum):
    AVAILABLE = "Available"
    PENDING = "Pending"
    DELETING = "Deleting"