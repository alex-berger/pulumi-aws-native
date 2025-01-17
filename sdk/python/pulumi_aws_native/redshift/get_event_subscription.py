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
    'GetEventSubscriptionResult',
    'AwaitableGetEventSubscriptionResult',
    'get_event_subscription',
    'get_event_subscription_output',
]

@pulumi.output_type
class GetEventSubscriptionResult:
    def __init__(__self__, cust_subscription_id=None, customer_aws_id=None, enabled=None, event_categories=None, event_categories_list=None, severity=None, sns_topic_arn=None, source_ids=None, source_ids_list=None, source_type=None, status=None, subscription_creation_time=None, tags=None):
        if cust_subscription_id and not isinstance(cust_subscription_id, str):
            raise TypeError("Expected argument 'cust_subscription_id' to be a str")
        pulumi.set(__self__, "cust_subscription_id", cust_subscription_id)
        if customer_aws_id and not isinstance(customer_aws_id, str):
            raise TypeError("Expected argument 'customer_aws_id' to be a str")
        pulumi.set(__self__, "customer_aws_id", customer_aws_id)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if event_categories and not isinstance(event_categories, list):
            raise TypeError("Expected argument 'event_categories' to be a list")
        pulumi.set(__self__, "event_categories", event_categories)
        if event_categories_list and not isinstance(event_categories_list, list):
            raise TypeError("Expected argument 'event_categories_list' to be a list")
        pulumi.set(__self__, "event_categories_list", event_categories_list)
        if severity and not isinstance(severity, str):
            raise TypeError("Expected argument 'severity' to be a str")
        pulumi.set(__self__, "severity", severity)
        if sns_topic_arn and not isinstance(sns_topic_arn, str):
            raise TypeError("Expected argument 'sns_topic_arn' to be a str")
        pulumi.set(__self__, "sns_topic_arn", sns_topic_arn)
        if source_ids and not isinstance(source_ids, list):
            raise TypeError("Expected argument 'source_ids' to be a list")
        pulumi.set(__self__, "source_ids", source_ids)
        if source_ids_list and not isinstance(source_ids_list, list):
            raise TypeError("Expected argument 'source_ids_list' to be a list")
        pulumi.set(__self__, "source_ids_list", source_ids_list)
        if source_type and not isinstance(source_type, str):
            raise TypeError("Expected argument 'source_type' to be a str")
        pulumi.set(__self__, "source_type", source_type)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if subscription_creation_time and not isinstance(subscription_creation_time, str):
            raise TypeError("Expected argument 'subscription_creation_time' to be a str")
        pulumi.set(__self__, "subscription_creation_time", subscription_creation_time)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="custSubscriptionId")
    def cust_subscription_id(self) -> Optional[str]:
        """
        The name of the Amazon Redshift event notification subscription.
        """
        return pulumi.get(self, "cust_subscription_id")

    @property
    @pulumi.getter(name="customerAwsId")
    def customer_aws_id(self) -> Optional[str]:
        """
        The AWS account associated with the Amazon Redshift event notification subscription.
        """
        return pulumi.get(self, "customer_aws_id")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        A boolean value; set to true to activate the subscription, and set to false to create the subscription but not activate it.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="eventCategories")
    def event_categories(self) -> Optional[Sequence['EventSubscriptionEventCategoriesItem']]:
        """
        Specifies the Amazon Redshift event categories to be published by the event notification subscription.
        """
        return pulumi.get(self, "event_categories")

    @property
    @pulumi.getter(name="eventCategoriesList")
    def event_categories_list(self) -> Optional[Sequence[str]]:
        """
        The list of Amazon Redshift event categories specified in the event notification subscription.
        """
        return pulumi.get(self, "event_categories_list")

    @property
    @pulumi.getter
    def severity(self) -> Optional['EventSubscriptionSeverity']:
        """
        Specifies the Amazon Redshift event severity to be published by the event notification subscription.
        """
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter(name="snsTopicArn")
    def sns_topic_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the Amazon SNS topic used to transmit the event notifications.
        """
        return pulumi.get(self, "sns_topic_arn")

    @property
    @pulumi.getter(name="sourceIds")
    def source_ids(self) -> Optional[Sequence[str]]:
        """
        A list of one or more identifiers of Amazon Redshift source objects.
        """
        return pulumi.get(self, "source_ids")

    @property
    @pulumi.getter(name="sourceIdsList")
    def source_ids_list(self) -> Optional[Sequence[str]]:
        """
        A list of the sources that publish events to the Amazon Redshift event notification subscription.
        """
        return pulumi.get(self, "source_ids_list")

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> Optional['EventSubscriptionSourceType']:
        """
        The type of source that will be generating the events.
        """
        return pulumi.get(self, "source_type")

    @property
    @pulumi.getter
    def status(self) -> Optional['EventSubscriptionStatus']:
        """
        The status of the Amazon Redshift event notification subscription.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subscriptionCreationTime")
    def subscription_creation_time(self) -> Optional[str]:
        """
        The date and time the Amazon Redshift event notification subscription was created.
        """
        return pulumi.get(self, "subscription_creation_time")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.EventSubscriptionTag']]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetEventSubscriptionResult(GetEventSubscriptionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEventSubscriptionResult(
            cust_subscription_id=self.cust_subscription_id,
            customer_aws_id=self.customer_aws_id,
            enabled=self.enabled,
            event_categories=self.event_categories,
            event_categories_list=self.event_categories_list,
            severity=self.severity,
            sns_topic_arn=self.sns_topic_arn,
            source_ids=self.source_ids,
            source_ids_list=self.source_ids_list,
            source_type=self.source_type,
            status=self.status,
            subscription_creation_time=self.subscription_creation_time,
            tags=self.tags)


def get_event_subscription(subscription_name: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEventSubscriptionResult:
    """
    The `AWS::Redshift::EventSubscription` resource creates an Amazon Redshift Event Subscription.


    :param str subscription_name: The name of the Amazon Redshift event notification subscription
    """
    __args__ = dict()
    __args__['subscriptionName'] = subscription_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:redshift:getEventSubscription', __args__, opts=opts, typ=GetEventSubscriptionResult).value

    return AwaitableGetEventSubscriptionResult(
        cust_subscription_id=__ret__.cust_subscription_id,
        customer_aws_id=__ret__.customer_aws_id,
        enabled=__ret__.enabled,
        event_categories=__ret__.event_categories,
        event_categories_list=__ret__.event_categories_list,
        severity=__ret__.severity,
        sns_topic_arn=__ret__.sns_topic_arn,
        source_ids=__ret__.source_ids,
        source_ids_list=__ret__.source_ids_list,
        source_type=__ret__.source_type,
        status=__ret__.status,
        subscription_creation_time=__ret__.subscription_creation_time,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_event_subscription)
def get_event_subscription_output(subscription_name: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEventSubscriptionResult]:
    """
    The `AWS::Redshift::EventSubscription` resource creates an Amazon Redshift Event Subscription.


    :param str subscription_name: The name of the Amazon Redshift event notification subscription
    """
    ...
