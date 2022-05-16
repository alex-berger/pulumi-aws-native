# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetCapacityReservationFleetResult',
    'AwaitableGetCapacityReservationFleetResult',
    'get_capacity_reservation_fleet',
    'get_capacity_reservation_fleet_output',
]

@pulumi.output_type
class GetCapacityReservationFleetResult:
    def __init__(__self__, capacity_reservation_fleet_id=None, no_remove_end_date=None, remove_end_date=None, total_target_capacity=None):
        if capacity_reservation_fleet_id and not isinstance(capacity_reservation_fleet_id, str):
            raise TypeError("Expected argument 'capacity_reservation_fleet_id' to be a str")
        pulumi.set(__self__, "capacity_reservation_fleet_id", capacity_reservation_fleet_id)
        if no_remove_end_date and not isinstance(no_remove_end_date, bool):
            raise TypeError("Expected argument 'no_remove_end_date' to be a bool")
        pulumi.set(__self__, "no_remove_end_date", no_remove_end_date)
        if remove_end_date and not isinstance(remove_end_date, bool):
            raise TypeError("Expected argument 'remove_end_date' to be a bool")
        pulumi.set(__self__, "remove_end_date", remove_end_date)
        if total_target_capacity and not isinstance(total_target_capacity, int):
            raise TypeError("Expected argument 'total_target_capacity' to be a int")
        pulumi.set(__self__, "total_target_capacity", total_target_capacity)

    @property
    @pulumi.getter(name="capacityReservationFleetId")
    def capacity_reservation_fleet_id(self) -> Optional[str]:
        return pulumi.get(self, "capacity_reservation_fleet_id")

    @property
    @pulumi.getter(name="noRemoveEndDate")
    def no_remove_end_date(self) -> Optional[bool]:
        return pulumi.get(self, "no_remove_end_date")

    @property
    @pulumi.getter(name="removeEndDate")
    def remove_end_date(self) -> Optional[bool]:
        return pulumi.get(self, "remove_end_date")

    @property
    @pulumi.getter(name="totalTargetCapacity")
    def total_target_capacity(self) -> Optional[int]:
        return pulumi.get(self, "total_target_capacity")


class AwaitableGetCapacityReservationFleetResult(GetCapacityReservationFleetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCapacityReservationFleetResult(
            capacity_reservation_fleet_id=self.capacity_reservation_fleet_id,
            no_remove_end_date=self.no_remove_end_date,
            remove_end_date=self.remove_end_date,
            total_target_capacity=self.total_target_capacity)


def get_capacity_reservation_fleet(capacity_reservation_fleet_id: Optional[str] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCapacityReservationFleetResult:
    """
    Resource Type definition for AWS::EC2::CapacityReservationFleet
    """
    __args__ = dict()
    __args__['capacityReservationFleetId'] = capacity_reservation_fleet_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:ec2:getCapacityReservationFleet', __args__, opts=opts, typ=GetCapacityReservationFleetResult).value

    return AwaitableGetCapacityReservationFleetResult(
        capacity_reservation_fleet_id=__ret__.capacity_reservation_fleet_id,
        no_remove_end_date=__ret__.no_remove_end_date,
        remove_end_date=__ret__.remove_end_date,
        total_target_capacity=__ret__.total_target_capacity)


@_utilities.lift_output_func(get_capacity_reservation_fleet)
def get_capacity_reservation_fleet_output(capacity_reservation_fleet_id: Optional[pulumi.Input[str]] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCapacityReservationFleetResult]:
    """
    Resource Type definition for AWS::EC2::CapacityReservationFleet
    """
    ...