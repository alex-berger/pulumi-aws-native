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
    'GetAnomalyDetectorResult',
    'AwaitableGetAnomalyDetectorResult',
    'get_anomaly_detector',
    'get_anomaly_detector_output',
]

@pulumi.output_type
class GetAnomalyDetectorResult:
    def __init__(__self__, anomaly_detector_config=None, anomaly_detector_description=None, arn=None, kms_key_arn=None, metric_set_list=None):
        if anomaly_detector_config and not isinstance(anomaly_detector_config, dict):
            raise TypeError("Expected argument 'anomaly_detector_config' to be a dict")
        pulumi.set(__self__, "anomaly_detector_config", anomaly_detector_config)
        if anomaly_detector_description and not isinstance(anomaly_detector_description, str):
            raise TypeError("Expected argument 'anomaly_detector_description' to be a str")
        pulumi.set(__self__, "anomaly_detector_description", anomaly_detector_description)
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if kms_key_arn and not isinstance(kms_key_arn, str):
            raise TypeError("Expected argument 'kms_key_arn' to be a str")
        pulumi.set(__self__, "kms_key_arn", kms_key_arn)
        if metric_set_list and not isinstance(metric_set_list, list):
            raise TypeError("Expected argument 'metric_set_list' to be a list")
        pulumi.set(__self__, "metric_set_list", metric_set_list)

    @property
    @pulumi.getter(name="anomalyDetectorConfig")
    def anomaly_detector_config(self) -> Optional['outputs.AnomalyDetectorConfig']:
        """
        Configuration options for the AnomalyDetector
        """
        return pulumi.get(self, "anomaly_detector_config")

    @property
    @pulumi.getter(name="anomalyDetectorDescription")
    def anomaly_detector_description(self) -> Optional[str]:
        """
        A description for the AnomalyDetector.
        """
        return pulumi.get(self, "anomaly_detector_description")

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> Optional[str]:
        """
        KMS key used to encrypt the AnomalyDetector data
        """
        return pulumi.get(self, "kms_key_arn")

    @property
    @pulumi.getter(name="metricSetList")
    def metric_set_list(self) -> Optional[Sequence['outputs.AnomalyDetectorMetricSet']]:
        """
        List of metric sets for anomaly detection
        """
        return pulumi.get(self, "metric_set_list")


class AwaitableGetAnomalyDetectorResult(GetAnomalyDetectorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAnomalyDetectorResult(
            anomaly_detector_config=self.anomaly_detector_config,
            anomaly_detector_description=self.anomaly_detector_description,
            arn=self.arn,
            kms_key_arn=self.kms_key_arn,
            metric_set_list=self.metric_set_list)


def get_anomaly_detector(arn: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAnomalyDetectorResult:
    """
    An Amazon Lookout for Metrics Detector
    """
    __args__ = dict()
    __args__['arn'] = arn
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:lookoutmetrics:getAnomalyDetector', __args__, opts=opts, typ=GetAnomalyDetectorResult).value

    return AwaitableGetAnomalyDetectorResult(
        anomaly_detector_config=__ret__.anomaly_detector_config,
        anomaly_detector_description=__ret__.anomaly_detector_description,
        arn=__ret__.arn,
        kms_key_arn=__ret__.kms_key_arn,
        metric_set_list=__ret__.metric_set_list)


@_utilities.lift_output_func(get_anomaly_detector)
def get_anomaly_detector_output(arn: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAnomalyDetectorResult]:
    """
    An Amazon Lookout for Metrics Detector
    """
    ...
