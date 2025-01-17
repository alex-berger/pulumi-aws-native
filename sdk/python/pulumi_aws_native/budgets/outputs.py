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
    'BudgetCostTypes',
    'BudgetData',
    'BudgetNotification',
    'BudgetNotificationWithSubscribers',
    'BudgetSpend',
    'BudgetSubscriber',
    'BudgetTimePeriod',
    'BudgetsActionActionThreshold',
    'BudgetsActionDefinition',
    'BudgetsActionIamActionDefinition',
    'BudgetsActionScpActionDefinition',
    'BudgetsActionSsmActionDefinition',
    'BudgetsActionSubscriber',
]

@pulumi.output_type
class BudgetCostTypes(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "includeCredit":
            suggest = "include_credit"
        elif key == "includeDiscount":
            suggest = "include_discount"
        elif key == "includeOtherSubscription":
            suggest = "include_other_subscription"
        elif key == "includeRecurring":
            suggest = "include_recurring"
        elif key == "includeRefund":
            suggest = "include_refund"
        elif key == "includeSubscription":
            suggest = "include_subscription"
        elif key == "includeSupport":
            suggest = "include_support"
        elif key == "includeTax":
            suggest = "include_tax"
        elif key == "includeUpfront":
            suggest = "include_upfront"
        elif key == "useAmortized":
            suggest = "use_amortized"
        elif key == "useBlended":
            suggest = "use_blended"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BudgetCostTypes. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BudgetCostTypes.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BudgetCostTypes.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 include_credit: Optional[bool] = None,
                 include_discount: Optional[bool] = None,
                 include_other_subscription: Optional[bool] = None,
                 include_recurring: Optional[bool] = None,
                 include_refund: Optional[bool] = None,
                 include_subscription: Optional[bool] = None,
                 include_support: Optional[bool] = None,
                 include_tax: Optional[bool] = None,
                 include_upfront: Optional[bool] = None,
                 use_amortized: Optional[bool] = None,
                 use_blended: Optional[bool] = None):
        if include_credit is not None:
            pulumi.set(__self__, "include_credit", include_credit)
        if include_discount is not None:
            pulumi.set(__self__, "include_discount", include_discount)
        if include_other_subscription is not None:
            pulumi.set(__self__, "include_other_subscription", include_other_subscription)
        if include_recurring is not None:
            pulumi.set(__self__, "include_recurring", include_recurring)
        if include_refund is not None:
            pulumi.set(__self__, "include_refund", include_refund)
        if include_subscription is not None:
            pulumi.set(__self__, "include_subscription", include_subscription)
        if include_support is not None:
            pulumi.set(__self__, "include_support", include_support)
        if include_tax is not None:
            pulumi.set(__self__, "include_tax", include_tax)
        if include_upfront is not None:
            pulumi.set(__self__, "include_upfront", include_upfront)
        if use_amortized is not None:
            pulumi.set(__self__, "use_amortized", use_amortized)
        if use_blended is not None:
            pulumi.set(__self__, "use_blended", use_blended)

    @property
    @pulumi.getter(name="includeCredit")
    def include_credit(self) -> Optional[bool]:
        return pulumi.get(self, "include_credit")

    @property
    @pulumi.getter(name="includeDiscount")
    def include_discount(self) -> Optional[bool]:
        return pulumi.get(self, "include_discount")

    @property
    @pulumi.getter(name="includeOtherSubscription")
    def include_other_subscription(self) -> Optional[bool]:
        return pulumi.get(self, "include_other_subscription")

    @property
    @pulumi.getter(name="includeRecurring")
    def include_recurring(self) -> Optional[bool]:
        return pulumi.get(self, "include_recurring")

    @property
    @pulumi.getter(name="includeRefund")
    def include_refund(self) -> Optional[bool]:
        return pulumi.get(self, "include_refund")

    @property
    @pulumi.getter(name="includeSubscription")
    def include_subscription(self) -> Optional[bool]:
        return pulumi.get(self, "include_subscription")

    @property
    @pulumi.getter(name="includeSupport")
    def include_support(self) -> Optional[bool]:
        return pulumi.get(self, "include_support")

    @property
    @pulumi.getter(name="includeTax")
    def include_tax(self) -> Optional[bool]:
        return pulumi.get(self, "include_tax")

    @property
    @pulumi.getter(name="includeUpfront")
    def include_upfront(self) -> Optional[bool]:
        return pulumi.get(self, "include_upfront")

    @property
    @pulumi.getter(name="useAmortized")
    def use_amortized(self) -> Optional[bool]:
        return pulumi.get(self, "use_amortized")

    @property
    @pulumi.getter(name="useBlended")
    def use_blended(self) -> Optional[bool]:
        return pulumi.get(self, "use_blended")


@pulumi.output_type
class BudgetData(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "budgetType":
            suggest = "budget_type"
        elif key == "timeUnit":
            suggest = "time_unit"
        elif key == "budgetLimit":
            suggest = "budget_limit"
        elif key == "budgetName":
            suggest = "budget_name"
        elif key == "costFilters":
            suggest = "cost_filters"
        elif key == "costTypes":
            suggest = "cost_types"
        elif key == "plannedBudgetLimits":
            suggest = "planned_budget_limits"
        elif key == "timePeriod":
            suggest = "time_period"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BudgetData. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BudgetData.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BudgetData.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 budget_type: str,
                 time_unit: str,
                 budget_limit: Optional['outputs.BudgetSpend'] = None,
                 budget_name: Optional[str] = None,
                 cost_filters: Optional[Any] = None,
                 cost_types: Optional['outputs.BudgetCostTypes'] = None,
                 planned_budget_limits: Optional[Any] = None,
                 time_period: Optional['outputs.BudgetTimePeriod'] = None):
        pulumi.set(__self__, "budget_type", budget_type)
        pulumi.set(__self__, "time_unit", time_unit)
        if budget_limit is not None:
            pulumi.set(__self__, "budget_limit", budget_limit)
        if budget_name is not None:
            pulumi.set(__self__, "budget_name", budget_name)
        if cost_filters is not None:
            pulumi.set(__self__, "cost_filters", cost_filters)
        if cost_types is not None:
            pulumi.set(__self__, "cost_types", cost_types)
        if planned_budget_limits is not None:
            pulumi.set(__self__, "planned_budget_limits", planned_budget_limits)
        if time_period is not None:
            pulumi.set(__self__, "time_period", time_period)

    @property
    @pulumi.getter(name="budgetType")
    def budget_type(self) -> str:
        return pulumi.get(self, "budget_type")

    @property
    @pulumi.getter(name="timeUnit")
    def time_unit(self) -> str:
        return pulumi.get(self, "time_unit")

    @property
    @pulumi.getter(name="budgetLimit")
    def budget_limit(self) -> Optional['outputs.BudgetSpend']:
        return pulumi.get(self, "budget_limit")

    @property
    @pulumi.getter(name="budgetName")
    def budget_name(self) -> Optional[str]:
        return pulumi.get(self, "budget_name")

    @property
    @pulumi.getter(name="costFilters")
    def cost_filters(self) -> Optional[Any]:
        return pulumi.get(self, "cost_filters")

    @property
    @pulumi.getter(name="costTypes")
    def cost_types(self) -> Optional['outputs.BudgetCostTypes']:
        return pulumi.get(self, "cost_types")

    @property
    @pulumi.getter(name="plannedBudgetLimits")
    def planned_budget_limits(self) -> Optional[Any]:
        return pulumi.get(self, "planned_budget_limits")

    @property
    @pulumi.getter(name="timePeriod")
    def time_period(self) -> Optional['outputs.BudgetTimePeriod']:
        return pulumi.get(self, "time_period")


@pulumi.output_type
class BudgetNotification(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "comparisonOperator":
            suggest = "comparison_operator"
        elif key == "notificationType":
            suggest = "notification_type"
        elif key == "thresholdType":
            suggest = "threshold_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BudgetNotification. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BudgetNotification.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BudgetNotification.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 comparison_operator: str,
                 notification_type: str,
                 threshold: float,
                 threshold_type: Optional[str] = None):
        pulumi.set(__self__, "comparison_operator", comparison_operator)
        pulumi.set(__self__, "notification_type", notification_type)
        pulumi.set(__self__, "threshold", threshold)
        if threshold_type is not None:
            pulumi.set(__self__, "threshold_type", threshold_type)

    @property
    @pulumi.getter(name="comparisonOperator")
    def comparison_operator(self) -> str:
        return pulumi.get(self, "comparison_operator")

    @property
    @pulumi.getter(name="notificationType")
    def notification_type(self) -> str:
        return pulumi.get(self, "notification_type")

    @property
    @pulumi.getter
    def threshold(self) -> float:
        return pulumi.get(self, "threshold")

    @property
    @pulumi.getter(name="thresholdType")
    def threshold_type(self) -> Optional[str]:
        return pulumi.get(self, "threshold_type")


@pulumi.output_type
class BudgetNotificationWithSubscribers(dict):
    def __init__(__self__, *,
                 notification: 'outputs.BudgetNotification',
                 subscribers: Sequence['outputs.BudgetSubscriber']):
        pulumi.set(__self__, "notification", notification)
        pulumi.set(__self__, "subscribers", subscribers)

    @property
    @pulumi.getter
    def notification(self) -> 'outputs.BudgetNotification':
        return pulumi.get(self, "notification")

    @property
    @pulumi.getter
    def subscribers(self) -> Sequence['outputs.BudgetSubscriber']:
        return pulumi.get(self, "subscribers")


@pulumi.output_type
class BudgetSpend(dict):
    def __init__(__self__, *,
                 amount: float,
                 unit: str):
        pulumi.set(__self__, "amount", amount)
        pulumi.set(__self__, "unit", unit)

    @property
    @pulumi.getter
    def amount(self) -> float:
        return pulumi.get(self, "amount")

    @property
    @pulumi.getter
    def unit(self) -> str:
        return pulumi.get(self, "unit")


@pulumi.output_type
class BudgetSubscriber(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "subscriptionType":
            suggest = "subscription_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BudgetSubscriber. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BudgetSubscriber.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BudgetSubscriber.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 address: str,
                 subscription_type: str):
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "subscription_type", subscription_type)

    @property
    @pulumi.getter
    def address(self) -> str:
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="subscriptionType")
    def subscription_type(self) -> str:
        return pulumi.get(self, "subscription_type")


@pulumi.output_type
class BudgetTimePeriod(dict):
    def __init__(__self__, *,
                 end: Optional[str] = None,
                 start: Optional[str] = None):
        if end is not None:
            pulumi.set(__self__, "end", end)
        if start is not None:
            pulumi.set(__self__, "start", start)

    @property
    @pulumi.getter
    def end(self) -> Optional[str]:
        return pulumi.get(self, "end")

    @property
    @pulumi.getter
    def start(self) -> Optional[str]:
        return pulumi.get(self, "start")


@pulumi.output_type
class BudgetsActionActionThreshold(dict):
    def __init__(__self__, *,
                 type: 'BudgetsActionActionThresholdType',
                 value: float):
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def type(self) -> 'BudgetsActionActionThresholdType':
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> float:
        return pulumi.get(self, "value")


@pulumi.output_type
class BudgetsActionDefinition(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "iamActionDefinition":
            suggest = "iam_action_definition"
        elif key == "scpActionDefinition":
            suggest = "scp_action_definition"
        elif key == "ssmActionDefinition":
            suggest = "ssm_action_definition"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BudgetsActionDefinition. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BudgetsActionDefinition.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BudgetsActionDefinition.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 iam_action_definition: Optional['outputs.BudgetsActionIamActionDefinition'] = None,
                 scp_action_definition: Optional['outputs.BudgetsActionScpActionDefinition'] = None,
                 ssm_action_definition: Optional['outputs.BudgetsActionSsmActionDefinition'] = None):
        if iam_action_definition is not None:
            pulumi.set(__self__, "iam_action_definition", iam_action_definition)
        if scp_action_definition is not None:
            pulumi.set(__self__, "scp_action_definition", scp_action_definition)
        if ssm_action_definition is not None:
            pulumi.set(__self__, "ssm_action_definition", ssm_action_definition)

    @property
    @pulumi.getter(name="iamActionDefinition")
    def iam_action_definition(self) -> Optional['outputs.BudgetsActionIamActionDefinition']:
        return pulumi.get(self, "iam_action_definition")

    @property
    @pulumi.getter(name="scpActionDefinition")
    def scp_action_definition(self) -> Optional['outputs.BudgetsActionScpActionDefinition']:
        return pulumi.get(self, "scp_action_definition")

    @property
    @pulumi.getter(name="ssmActionDefinition")
    def ssm_action_definition(self) -> Optional['outputs.BudgetsActionSsmActionDefinition']:
        return pulumi.get(self, "ssm_action_definition")


@pulumi.output_type
class BudgetsActionIamActionDefinition(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "policyArn":
            suggest = "policy_arn"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BudgetsActionIamActionDefinition. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BudgetsActionIamActionDefinition.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BudgetsActionIamActionDefinition.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 policy_arn: str,
                 groups: Optional[Sequence[str]] = None,
                 roles: Optional[Sequence[str]] = None,
                 users: Optional[Sequence[str]] = None):
        pulumi.set(__self__, "policy_arn", policy_arn)
        if groups is not None:
            pulumi.set(__self__, "groups", groups)
        if roles is not None:
            pulumi.set(__self__, "roles", roles)
        if users is not None:
            pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter(name="policyArn")
    def policy_arn(self) -> str:
        return pulumi.get(self, "policy_arn")

    @property
    @pulumi.getter
    def groups(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter
    def roles(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter
    def users(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "users")


@pulumi.output_type
class BudgetsActionScpActionDefinition(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "policyId":
            suggest = "policy_id"
        elif key == "targetIds":
            suggest = "target_ids"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BudgetsActionScpActionDefinition. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BudgetsActionScpActionDefinition.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BudgetsActionScpActionDefinition.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 policy_id: str,
                 target_ids: Sequence[str]):
        pulumi.set(__self__, "policy_id", policy_id)
        pulumi.set(__self__, "target_ids", target_ids)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> str:
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter(name="targetIds")
    def target_ids(self) -> Sequence[str]:
        return pulumi.get(self, "target_ids")


@pulumi.output_type
class BudgetsActionSsmActionDefinition(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "instanceIds":
            suggest = "instance_ids"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BudgetsActionSsmActionDefinition. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BudgetsActionSsmActionDefinition.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BudgetsActionSsmActionDefinition.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 instance_ids: Sequence[str],
                 region: str,
                 subtype: 'BudgetsActionSsmActionDefinitionSubtype'):
        pulumi.set(__self__, "instance_ids", instance_ids)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "subtype", subtype)

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> Sequence[str]:
        return pulumi.get(self, "instance_ids")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def subtype(self) -> 'BudgetsActionSsmActionDefinitionSubtype':
        return pulumi.get(self, "subtype")


@pulumi.output_type
class BudgetsActionSubscriber(dict):
    def __init__(__self__, *,
                 address: str,
                 type: 'BudgetsActionSubscriberType'):
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def address(self) -> str:
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def type(self) -> 'BudgetsActionSubscriberType':
        return pulumi.get(self, "type")


