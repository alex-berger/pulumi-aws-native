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
    'GetAlarmResult',
    'AwaitableGetAlarmResult',
    'get_alarm',
    'get_alarm_output',
]

@pulumi.output_type
class GetAlarmResult:
    def __init__(__self__, actions_enabled=None, alarm_actions=None, alarm_description=None, arn=None, comparison_operator=None, datapoints_to_alarm=None, dimensions=None, evaluate_low_sample_count_percentile=None, evaluation_periods=None, extended_statistic=None, id=None, insufficient_data_actions=None, metric_name=None, metrics=None, namespace=None, o_k_actions=None, period=None, statistic=None, threshold=None, threshold_metric_id=None, treat_missing_data=None, unit=None):
        if actions_enabled and not isinstance(actions_enabled, bool):
            raise TypeError("Expected argument 'actions_enabled' to be a bool")
        pulumi.set(__self__, "actions_enabled", actions_enabled)
        if alarm_actions and not isinstance(alarm_actions, list):
            raise TypeError("Expected argument 'alarm_actions' to be a list")
        pulumi.set(__self__, "alarm_actions", alarm_actions)
        if alarm_description and not isinstance(alarm_description, str):
            raise TypeError("Expected argument 'alarm_description' to be a str")
        pulumi.set(__self__, "alarm_description", alarm_description)
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if comparison_operator and not isinstance(comparison_operator, str):
            raise TypeError("Expected argument 'comparison_operator' to be a str")
        pulumi.set(__self__, "comparison_operator", comparison_operator)
        if datapoints_to_alarm and not isinstance(datapoints_to_alarm, int):
            raise TypeError("Expected argument 'datapoints_to_alarm' to be a int")
        pulumi.set(__self__, "datapoints_to_alarm", datapoints_to_alarm)
        if dimensions and not isinstance(dimensions, list):
            raise TypeError("Expected argument 'dimensions' to be a list")
        pulumi.set(__self__, "dimensions", dimensions)
        if evaluate_low_sample_count_percentile and not isinstance(evaluate_low_sample_count_percentile, str):
            raise TypeError("Expected argument 'evaluate_low_sample_count_percentile' to be a str")
        pulumi.set(__self__, "evaluate_low_sample_count_percentile", evaluate_low_sample_count_percentile)
        if evaluation_periods and not isinstance(evaluation_periods, int):
            raise TypeError("Expected argument 'evaluation_periods' to be a int")
        pulumi.set(__self__, "evaluation_periods", evaluation_periods)
        if extended_statistic and not isinstance(extended_statistic, str):
            raise TypeError("Expected argument 'extended_statistic' to be a str")
        pulumi.set(__self__, "extended_statistic", extended_statistic)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if insufficient_data_actions and not isinstance(insufficient_data_actions, list):
            raise TypeError("Expected argument 'insufficient_data_actions' to be a list")
        pulumi.set(__self__, "insufficient_data_actions", insufficient_data_actions)
        if metric_name and not isinstance(metric_name, str):
            raise TypeError("Expected argument 'metric_name' to be a str")
        pulumi.set(__self__, "metric_name", metric_name)
        if metrics and not isinstance(metrics, list):
            raise TypeError("Expected argument 'metrics' to be a list")
        pulumi.set(__self__, "metrics", metrics)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if o_k_actions and not isinstance(o_k_actions, list):
            raise TypeError("Expected argument 'o_k_actions' to be a list")
        pulumi.set(__self__, "o_k_actions", o_k_actions)
        if period and not isinstance(period, int):
            raise TypeError("Expected argument 'period' to be a int")
        pulumi.set(__self__, "period", period)
        if statistic and not isinstance(statistic, str):
            raise TypeError("Expected argument 'statistic' to be a str")
        pulumi.set(__self__, "statistic", statistic)
        if threshold and not isinstance(threshold, float):
            raise TypeError("Expected argument 'threshold' to be a float")
        pulumi.set(__self__, "threshold", threshold)
        if threshold_metric_id and not isinstance(threshold_metric_id, str):
            raise TypeError("Expected argument 'threshold_metric_id' to be a str")
        pulumi.set(__self__, "threshold_metric_id", threshold_metric_id)
        if treat_missing_data and not isinstance(treat_missing_data, str):
            raise TypeError("Expected argument 'treat_missing_data' to be a str")
        pulumi.set(__self__, "treat_missing_data", treat_missing_data)
        if unit and not isinstance(unit, str):
            raise TypeError("Expected argument 'unit' to be a str")
        pulumi.set(__self__, "unit", unit)

    @property
    @pulumi.getter(name="actionsEnabled")
    def actions_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "actions_enabled")

    @property
    @pulumi.getter(name="alarmActions")
    def alarm_actions(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "alarm_actions")

    @property
    @pulumi.getter(name="alarmDescription")
    def alarm_description(self) -> Optional[str]:
        return pulumi.get(self, "alarm_description")

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="comparisonOperator")
    def comparison_operator(self) -> Optional[str]:
        return pulumi.get(self, "comparison_operator")

    @property
    @pulumi.getter(name="datapointsToAlarm")
    def datapoints_to_alarm(self) -> Optional[int]:
        return pulumi.get(self, "datapoints_to_alarm")

    @property
    @pulumi.getter
    def dimensions(self) -> Optional[Sequence['outputs.AlarmDimension']]:
        return pulumi.get(self, "dimensions")

    @property
    @pulumi.getter(name="evaluateLowSampleCountPercentile")
    def evaluate_low_sample_count_percentile(self) -> Optional[str]:
        return pulumi.get(self, "evaluate_low_sample_count_percentile")

    @property
    @pulumi.getter(name="evaluationPeriods")
    def evaluation_periods(self) -> Optional[int]:
        return pulumi.get(self, "evaluation_periods")

    @property
    @pulumi.getter(name="extendedStatistic")
    def extended_statistic(self) -> Optional[str]:
        return pulumi.get(self, "extended_statistic")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="insufficientDataActions")
    def insufficient_data_actions(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "insufficient_data_actions")

    @property
    @pulumi.getter(name="metricName")
    def metric_name(self) -> Optional[str]:
        return pulumi.get(self, "metric_name")

    @property
    @pulumi.getter
    def metrics(self) -> Optional[Sequence['outputs.AlarmMetricDataQuery']]:
        return pulumi.get(self, "metrics")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="oKActions")
    def o_k_actions(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "o_k_actions")

    @property
    @pulumi.getter
    def period(self) -> Optional[int]:
        return pulumi.get(self, "period")

    @property
    @pulumi.getter
    def statistic(self) -> Optional[str]:
        return pulumi.get(self, "statistic")

    @property
    @pulumi.getter
    def threshold(self) -> Optional[float]:
        return pulumi.get(self, "threshold")

    @property
    @pulumi.getter(name="thresholdMetricId")
    def threshold_metric_id(self) -> Optional[str]:
        return pulumi.get(self, "threshold_metric_id")

    @property
    @pulumi.getter(name="treatMissingData")
    def treat_missing_data(self) -> Optional[str]:
        return pulumi.get(self, "treat_missing_data")

    @property
    @pulumi.getter
    def unit(self) -> Optional[str]:
        return pulumi.get(self, "unit")


class AwaitableGetAlarmResult(GetAlarmResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAlarmResult(
            actions_enabled=self.actions_enabled,
            alarm_actions=self.alarm_actions,
            alarm_description=self.alarm_description,
            arn=self.arn,
            comparison_operator=self.comparison_operator,
            datapoints_to_alarm=self.datapoints_to_alarm,
            dimensions=self.dimensions,
            evaluate_low_sample_count_percentile=self.evaluate_low_sample_count_percentile,
            evaluation_periods=self.evaluation_periods,
            extended_statistic=self.extended_statistic,
            id=self.id,
            insufficient_data_actions=self.insufficient_data_actions,
            metric_name=self.metric_name,
            metrics=self.metrics,
            namespace=self.namespace,
            o_k_actions=self.o_k_actions,
            period=self.period,
            statistic=self.statistic,
            threshold=self.threshold,
            threshold_metric_id=self.threshold_metric_id,
            treat_missing_data=self.treat_missing_data,
            unit=self.unit)


def get_alarm(id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAlarmResult:
    """
    Resource Type definition for AWS::CloudWatch::Alarm
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:cloudwatch:getAlarm', __args__, opts=opts, typ=GetAlarmResult).value

    return AwaitableGetAlarmResult(
        actions_enabled=__ret__.actions_enabled,
        alarm_actions=__ret__.alarm_actions,
        alarm_description=__ret__.alarm_description,
        arn=__ret__.arn,
        comparison_operator=__ret__.comparison_operator,
        datapoints_to_alarm=__ret__.datapoints_to_alarm,
        dimensions=__ret__.dimensions,
        evaluate_low_sample_count_percentile=__ret__.evaluate_low_sample_count_percentile,
        evaluation_periods=__ret__.evaluation_periods,
        extended_statistic=__ret__.extended_statistic,
        id=__ret__.id,
        insufficient_data_actions=__ret__.insufficient_data_actions,
        metric_name=__ret__.metric_name,
        metrics=__ret__.metrics,
        namespace=__ret__.namespace,
        o_k_actions=__ret__.o_k_actions,
        period=__ret__.period,
        statistic=__ret__.statistic,
        threshold=__ret__.threshold,
        threshold_metric_id=__ret__.threshold_metric_id,
        treat_missing_data=__ret__.treat_missing_data,
        unit=__ret__.unit)


@_utilities.lift_output_func(get_alarm)
def get_alarm_output(id: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAlarmResult]:
    """
    Resource Type definition for AWS::CloudWatch::Alarm
    """
    ...