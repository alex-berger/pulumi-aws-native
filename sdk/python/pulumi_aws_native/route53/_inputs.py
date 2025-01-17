# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'HealthCheckAlarmIdentifierArgs',
    'HealthCheckConfigPropertiesArgs',
    'HealthCheckTagArgs',
    'HostedZoneConfigArgs',
    'HostedZoneQueryLoggingConfigArgs',
    'HostedZoneTagArgs',
    'HostedZoneVPCArgs',
    'RecordSetAliasTargetArgs',
    'RecordSetGeoLocationArgs',
    'RecordSetGroupAliasTargetArgs',
    'RecordSetGroupGeoLocationArgs',
    'RecordSetGroupRecordSetArgs',
]

@pulumi.input_type
class HealthCheckAlarmIdentifierArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 region: pulumi.Input[str]):
        """
        A complex type that identifies the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether the specified health check is healthy.
        :param pulumi.Input[str] name: The name of the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether this health check is healthy.
        :param pulumi.Input[str] region: For the CloudWatch alarm that you want Route 53 health checkers to use to determine whether this health check is healthy, the region that the alarm was created in.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether this health check is healthy.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        For the CloudWatch alarm that you want Route 53 health checkers to use to determine whether this health check is healthy, the region that the alarm was created in.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class HealthCheckConfigPropertiesArgs:
    def __init__(__self__, *,
                 type: pulumi.Input['HealthCheckConfigPropertiesType'],
                 alarm_identifier: Optional[pulumi.Input['HealthCheckAlarmIdentifierArgs']] = None,
                 child_health_checks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_sni: Optional[pulumi.Input[bool]] = None,
                 failure_threshold: Optional[pulumi.Input[int]] = None,
                 fully_qualified_domain_name: Optional[pulumi.Input[str]] = None,
                 health_threshold: Optional[pulumi.Input[int]] = None,
                 i_p_address: Optional[pulumi.Input[str]] = None,
                 insufficient_data_health_status: Optional[pulumi.Input['HealthCheckConfigPropertiesInsufficientDataHealthStatus']] = None,
                 inverted: Optional[pulumi.Input[bool]] = None,
                 measure_latency: Optional[pulumi.Input[bool]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 request_interval: Optional[pulumi.Input[int]] = None,
                 resource_path: Optional[pulumi.Input[str]] = None,
                 routing_control_arn: Optional[pulumi.Input[str]] = None,
                 search_string: Optional[pulumi.Input[str]] = None):
        """
        A complex type that contains information about the health check.
        """
        pulumi.set(__self__, "type", type)
        if alarm_identifier is not None:
            pulumi.set(__self__, "alarm_identifier", alarm_identifier)
        if child_health_checks is not None:
            pulumi.set(__self__, "child_health_checks", child_health_checks)
        if enable_sni is not None:
            pulumi.set(__self__, "enable_sni", enable_sni)
        if failure_threshold is not None:
            pulumi.set(__self__, "failure_threshold", failure_threshold)
        if fully_qualified_domain_name is not None:
            pulumi.set(__self__, "fully_qualified_domain_name", fully_qualified_domain_name)
        if health_threshold is not None:
            pulumi.set(__self__, "health_threshold", health_threshold)
        if i_p_address is not None:
            pulumi.set(__self__, "i_p_address", i_p_address)
        if insufficient_data_health_status is not None:
            pulumi.set(__self__, "insufficient_data_health_status", insufficient_data_health_status)
        if inverted is not None:
            pulumi.set(__self__, "inverted", inverted)
        if measure_latency is not None:
            pulumi.set(__self__, "measure_latency", measure_latency)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if regions is not None:
            pulumi.set(__self__, "regions", regions)
        if request_interval is not None:
            pulumi.set(__self__, "request_interval", request_interval)
        if resource_path is not None:
            pulumi.set(__self__, "resource_path", resource_path)
        if routing_control_arn is not None:
            pulumi.set(__self__, "routing_control_arn", routing_control_arn)
        if search_string is not None:
            pulumi.set(__self__, "search_string", search_string)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['HealthCheckConfigPropertiesType']:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['HealthCheckConfigPropertiesType']):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="alarmIdentifier")
    def alarm_identifier(self) -> Optional[pulumi.Input['HealthCheckAlarmIdentifierArgs']]:
        return pulumi.get(self, "alarm_identifier")

    @alarm_identifier.setter
    def alarm_identifier(self, value: Optional[pulumi.Input['HealthCheckAlarmIdentifierArgs']]):
        pulumi.set(self, "alarm_identifier", value)

    @property
    @pulumi.getter(name="childHealthChecks")
    def child_health_checks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "child_health_checks")

    @child_health_checks.setter
    def child_health_checks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "child_health_checks", value)

    @property
    @pulumi.getter(name="enableSNI")
    def enable_sni(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enable_sni")

    @enable_sni.setter
    def enable_sni(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_sni", value)

    @property
    @pulumi.getter(name="failureThreshold")
    def failure_threshold(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "failure_threshold")

    @failure_threshold.setter
    def failure_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "failure_threshold", value)

    @property
    @pulumi.getter(name="fullyQualifiedDomainName")
    def fully_qualified_domain_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fully_qualified_domain_name")

    @fully_qualified_domain_name.setter
    def fully_qualified_domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fully_qualified_domain_name", value)

    @property
    @pulumi.getter(name="healthThreshold")
    def health_threshold(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "health_threshold")

    @health_threshold.setter
    def health_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "health_threshold", value)

    @property
    @pulumi.getter(name="iPAddress")
    def i_p_address(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "i_p_address")

    @i_p_address.setter
    def i_p_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "i_p_address", value)

    @property
    @pulumi.getter(name="insufficientDataHealthStatus")
    def insufficient_data_health_status(self) -> Optional[pulumi.Input['HealthCheckConfigPropertiesInsufficientDataHealthStatus']]:
        return pulumi.get(self, "insufficient_data_health_status")

    @insufficient_data_health_status.setter
    def insufficient_data_health_status(self, value: Optional[pulumi.Input['HealthCheckConfigPropertiesInsufficientDataHealthStatus']]):
        pulumi.set(self, "insufficient_data_health_status", value)

    @property
    @pulumi.getter
    def inverted(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "inverted")

    @inverted.setter
    def inverted(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "inverted", value)

    @property
    @pulumi.getter(name="measureLatency")
    def measure_latency(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "measure_latency")

    @measure_latency.setter
    def measure_latency(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "measure_latency", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def regions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "regions")

    @regions.setter
    def regions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "regions", value)

    @property
    @pulumi.getter(name="requestInterval")
    def request_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "request_interval")

    @request_interval.setter
    def request_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "request_interval", value)

    @property
    @pulumi.getter(name="resourcePath")
    def resource_path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "resource_path")

    @resource_path.setter
    def resource_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_path", value)

    @property
    @pulumi.getter(name="routingControlArn")
    def routing_control_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "routing_control_arn")

    @routing_control_arn.setter
    def routing_control_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "routing_control_arn", value)

    @property
    @pulumi.getter(name="searchString")
    def search_string(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "search_string")

    @search_string.setter
    def search_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "search_string", value)


@pulumi.input_type
class HealthCheckTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A key-value pair to associate with a resource.
        :param pulumi.Input[str] key: The key name of the tag.
        :param pulumi.Input[str] value: The value for the tag.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The key name of the tag.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value for the tag.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class HostedZoneConfigArgs:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None):
        """
        A complex type that contains an optional comment.

        If you don't want to specify a comment, omit the HostedZoneConfig and Comment elements.
        :param pulumi.Input[str] comment: Any comments that you want to include about the hosted zone.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Any comments that you want to include about the hosted zone.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)


@pulumi.input_type
class HostedZoneQueryLoggingConfigArgs:
    def __init__(__self__, *,
                 cloud_watch_logs_log_group_arn: pulumi.Input[str]):
        """
        A complex type that contains information about a configuration for DNS query logging.
        :param pulumi.Input[str] cloud_watch_logs_log_group_arn: The Amazon Resource Name (ARN) of the CloudWatch Logs log group that Amazon Route 53 is publishing logs to.
        """
        pulumi.set(__self__, "cloud_watch_logs_log_group_arn", cloud_watch_logs_log_group_arn)

    @property
    @pulumi.getter(name="cloudWatchLogsLogGroupArn")
    def cloud_watch_logs_log_group_arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN) of the CloudWatch Logs log group that Amazon Route 53 is publishing logs to.
        """
        return pulumi.get(self, "cloud_watch_logs_log_group_arn")

    @cloud_watch_logs_log_group_arn.setter
    def cloud_watch_logs_log_group_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "cloud_watch_logs_log_group_arn", value)


@pulumi.input_type
class HostedZoneTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A complex type that contains information about a tag that you want to add or edit for the specified health check or hosted zone.
        :param pulumi.Input[str] key: The key name of the tag.
        :param pulumi.Input[str] value: The value for the tag.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The key name of the tag.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value for the tag.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class HostedZoneVPCArgs:
    def __init__(__self__, *,
                 v_pc_id: pulumi.Input[str],
                 v_pc_region: pulumi.Input[str]):
        """
        A complex type that contains information about an Amazon VPC. Route 53 Resolver uses the records in the private hosted zone to route traffic in that VPC.
        :param pulumi.Input[str] v_pc_id: The ID of an Amazon VPC.
        :param pulumi.Input[str] v_pc_region: The region that an Amazon VPC was created in. See https://docs.aws.amazon.com/general/latest/gr/rande.html for a list of up to date regions.
        """
        pulumi.set(__self__, "v_pc_id", v_pc_id)
        pulumi.set(__self__, "v_pc_region", v_pc_region)

    @property
    @pulumi.getter(name="vPCId")
    def v_pc_id(self) -> pulumi.Input[str]:
        """
        The ID of an Amazon VPC.
        """
        return pulumi.get(self, "v_pc_id")

    @v_pc_id.setter
    def v_pc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "v_pc_id", value)

    @property
    @pulumi.getter(name="vPCRegion")
    def v_pc_region(self) -> pulumi.Input[str]:
        """
        The region that an Amazon VPC was created in. See https://docs.aws.amazon.com/general/latest/gr/rande.html for a list of up to date regions.
        """
        return pulumi.get(self, "v_pc_region")

    @v_pc_region.setter
    def v_pc_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "v_pc_region", value)


@pulumi.input_type
class RecordSetAliasTargetArgs:
    def __init__(__self__, *,
                 d_ns_name: pulumi.Input[str],
                 hosted_zone_id: pulumi.Input[str],
                 evaluate_target_health: Optional[pulumi.Input[bool]] = None):
        pulumi.set(__self__, "d_ns_name", d_ns_name)
        pulumi.set(__self__, "hosted_zone_id", hosted_zone_id)
        if evaluate_target_health is not None:
            pulumi.set(__self__, "evaluate_target_health", evaluate_target_health)

    @property
    @pulumi.getter(name="dNSName")
    def d_ns_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "d_ns_name")

    @d_ns_name.setter
    def d_ns_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "d_ns_name", value)

    @property
    @pulumi.getter(name="hostedZoneId")
    def hosted_zone_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "hosted_zone_id")

    @hosted_zone_id.setter
    def hosted_zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "hosted_zone_id", value)

    @property
    @pulumi.getter(name="evaluateTargetHealth")
    def evaluate_target_health(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "evaluate_target_health")

    @evaluate_target_health.setter
    def evaluate_target_health(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "evaluate_target_health", value)


@pulumi.input_type
class RecordSetGeoLocationArgs:
    def __init__(__self__, *,
                 continent_code: Optional[pulumi.Input[str]] = None,
                 country_code: Optional[pulumi.Input[str]] = None,
                 subdivision_code: Optional[pulumi.Input[str]] = None):
        if continent_code is not None:
            pulumi.set(__self__, "continent_code", continent_code)
        if country_code is not None:
            pulumi.set(__self__, "country_code", country_code)
        if subdivision_code is not None:
            pulumi.set(__self__, "subdivision_code", subdivision_code)

    @property
    @pulumi.getter(name="continentCode")
    def continent_code(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "continent_code")

    @continent_code.setter
    def continent_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "continent_code", value)

    @property
    @pulumi.getter(name="countryCode")
    def country_code(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "country_code")

    @country_code.setter
    def country_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "country_code", value)

    @property
    @pulumi.getter(name="subdivisionCode")
    def subdivision_code(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "subdivision_code")

    @subdivision_code.setter
    def subdivision_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subdivision_code", value)


@pulumi.input_type
class RecordSetGroupAliasTargetArgs:
    def __init__(__self__, *,
                 d_ns_name: pulumi.Input[str],
                 hosted_zone_id: pulumi.Input[str],
                 evaluate_target_health: Optional[pulumi.Input[bool]] = None):
        pulumi.set(__self__, "d_ns_name", d_ns_name)
        pulumi.set(__self__, "hosted_zone_id", hosted_zone_id)
        if evaluate_target_health is not None:
            pulumi.set(__self__, "evaluate_target_health", evaluate_target_health)

    @property
    @pulumi.getter(name="dNSName")
    def d_ns_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "d_ns_name")

    @d_ns_name.setter
    def d_ns_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "d_ns_name", value)

    @property
    @pulumi.getter(name="hostedZoneId")
    def hosted_zone_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "hosted_zone_id")

    @hosted_zone_id.setter
    def hosted_zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "hosted_zone_id", value)

    @property
    @pulumi.getter(name="evaluateTargetHealth")
    def evaluate_target_health(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "evaluate_target_health")

    @evaluate_target_health.setter
    def evaluate_target_health(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "evaluate_target_health", value)


@pulumi.input_type
class RecordSetGroupGeoLocationArgs:
    def __init__(__self__, *,
                 continent_code: Optional[pulumi.Input[str]] = None,
                 country_code: Optional[pulumi.Input[str]] = None,
                 subdivision_code: Optional[pulumi.Input[str]] = None):
        if continent_code is not None:
            pulumi.set(__self__, "continent_code", continent_code)
        if country_code is not None:
            pulumi.set(__self__, "country_code", country_code)
        if subdivision_code is not None:
            pulumi.set(__self__, "subdivision_code", subdivision_code)

    @property
    @pulumi.getter(name="continentCode")
    def continent_code(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "continent_code")

    @continent_code.setter
    def continent_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "continent_code", value)

    @property
    @pulumi.getter(name="countryCode")
    def country_code(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "country_code")

    @country_code.setter
    def country_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "country_code", value)

    @property
    @pulumi.getter(name="subdivisionCode")
    def subdivision_code(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "subdivision_code")

    @subdivision_code.setter
    def subdivision_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subdivision_code", value)


@pulumi.input_type
class RecordSetGroupRecordSetArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 type: pulumi.Input[str],
                 alias_target: Optional[pulumi.Input['RecordSetGroupAliasTargetArgs']] = None,
                 failover: Optional[pulumi.Input[str]] = None,
                 geo_location: Optional[pulumi.Input['RecordSetGroupGeoLocationArgs']] = None,
                 health_check_id: Optional[pulumi.Input[str]] = None,
                 hosted_zone_id: Optional[pulumi.Input[str]] = None,
                 hosted_zone_name: Optional[pulumi.Input[str]] = None,
                 multi_value_answer: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 resource_records: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 set_identifier: Optional[pulumi.Input[str]] = None,
                 t_tl: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)
        if alias_target is not None:
            pulumi.set(__self__, "alias_target", alias_target)
        if failover is not None:
            pulumi.set(__self__, "failover", failover)
        if geo_location is not None:
            pulumi.set(__self__, "geo_location", geo_location)
        if health_check_id is not None:
            pulumi.set(__self__, "health_check_id", health_check_id)
        if hosted_zone_id is not None:
            pulumi.set(__self__, "hosted_zone_id", hosted_zone_id)
        if hosted_zone_name is not None:
            pulumi.set(__self__, "hosted_zone_name", hosted_zone_name)
        if multi_value_answer is not None:
            pulumi.set(__self__, "multi_value_answer", multi_value_answer)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if resource_records is not None:
            pulumi.set(__self__, "resource_records", resource_records)
        if set_identifier is not None:
            pulumi.set(__self__, "set_identifier", set_identifier)
        if t_tl is not None:
            pulumi.set(__self__, "t_tl", t_tl)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="aliasTarget")
    def alias_target(self) -> Optional[pulumi.Input['RecordSetGroupAliasTargetArgs']]:
        return pulumi.get(self, "alias_target")

    @alias_target.setter
    def alias_target(self, value: Optional[pulumi.Input['RecordSetGroupAliasTargetArgs']]):
        pulumi.set(self, "alias_target", value)

    @property
    @pulumi.getter
    def failover(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "failover")

    @failover.setter
    def failover(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "failover", value)

    @property
    @pulumi.getter(name="geoLocation")
    def geo_location(self) -> Optional[pulumi.Input['RecordSetGroupGeoLocationArgs']]:
        return pulumi.get(self, "geo_location")

    @geo_location.setter
    def geo_location(self, value: Optional[pulumi.Input['RecordSetGroupGeoLocationArgs']]):
        pulumi.set(self, "geo_location", value)

    @property
    @pulumi.getter(name="healthCheckId")
    def health_check_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "health_check_id")

    @health_check_id.setter
    def health_check_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "health_check_id", value)

    @property
    @pulumi.getter(name="hostedZoneId")
    def hosted_zone_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "hosted_zone_id")

    @hosted_zone_id.setter
    def hosted_zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hosted_zone_id", value)

    @property
    @pulumi.getter(name="hostedZoneName")
    def hosted_zone_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "hosted_zone_name")

    @hosted_zone_name.setter
    def hosted_zone_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hosted_zone_name", value)

    @property
    @pulumi.getter(name="multiValueAnswer")
    def multi_value_answer(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "multi_value_answer")

    @multi_value_answer.setter
    def multi_value_answer(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "multi_value_answer", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="resourceRecords")
    def resource_records(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "resource_records")

    @resource_records.setter
    def resource_records(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "resource_records", value)

    @property
    @pulumi.getter(name="setIdentifier")
    def set_identifier(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "set_identifier")

    @set_identifier.setter
    def set_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "set_identifier", value)

    @property
    @pulumi.getter(name="tTL")
    def t_tl(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "t_tl")

    @t_tl.setter
    def t_tl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "t_tl", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


