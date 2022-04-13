# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['DBInstanceArgs', 'DBInstance']

@pulumi.input_type
class DBInstanceArgs:
    def __init__(__self__, *,
                 d_b_instance_class: pulumi.Input[str],
                 allow_major_version_upgrade: Optional[pulumi.Input[bool]] = None,
                 auto_minor_version_upgrade: Optional[pulumi.Input[bool]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 d_b_cluster_identifier: Optional[pulumi.Input[str]] = None,
                 d_b_instance_identifier: Optional[pulumi.Input[str]] = None,
                 d_b_parameter_group_name: Optional[pulumi.Input[str]] = None,
                 d_b_snapshot_identifier: Optional[pulumi.Input[str]] = None,
                 d_b_subnet_group_name: Optional[pulumi.Input[str]] = None,
                 preferred_maintenance_window: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['DBInstanceTagArgs']]]] = None):
        """
        The set of arguments for constructing a DBInstance resource.
        """
        pulumi.set(__self__, "d_b_instance_class", d_b_instance_class)
        if allow_major_version_upgrade is not None:
            pulumi.set(__self__, "allow_major_version_upgrade", allow_major_version_upgrade)
        if auto_minor_version_upgrade is not None:
            pulumi.set(__self__, "auto_minor_version_upgrade", auto_minor_version_upgrade)
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if d_b_cluster_identifier is not None:
            pulumi.set(__self__, "d_b_cluster_identifier", d_b_cluster_identifier)
        if d_b_instance_identifier is not None:
            pulumi.set(__self__, "d_b_instance_identifier", d_b_instance_identifier)
        if d_b_parameter_group_name is not None:
            pulumi.set(__self__, "d_b_parameter_group_name", d_b_parameter_group_name)
        if d_b_snapshot_identifier is not None:
            pulumi.set(__self__, "d_b_snapshot_identifier", d_b_snapshot_identifier)
        if d_b_subnet_group_name is not None:
            pulumi.set(__self__, "d_b_subnet_group_name", d_b_subnet_group_name)
        if preferred_maintenance_window is not None:
            pulumi.set(__self__, "preferred_maintenance_window", preferred_maintenance_window)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="dBInstanceClass")
    def d_b_instance_class(self) -> pulumi.Input[str]:
        return pulumi.get(self, "d_b_instance_class")

    @d_b_instance_class.setter
    def d_b_instance_class(self, value: pulumi.Input[str]):
        pulumi.set(self, "d_b_instance_class", value)

    @property
    @pulumi.getter(name="allowMajorVersionUpgrade")
    def allow_major_version_upgrade(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "allow_major_version_upgrade")

    @allow_major_version_upgrade.setter
    def allow_major_version_upgrade(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_major_version_upgrade", value)

    @property
    @pulumi.getter(name="autoMinorVersionUpgrade")
    def auto_minor_version_upgrade(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "auto_minor_version_upgrade")

    @auto_minor_version_upgrade.setter
    def auto_minor_version_upgrade(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_minor_version_upgrade", value)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter(name="dBClusterIdentifier")
    def d_b_cluster_identifier(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "d_b_cluster_identifier")

    @d_b_cluster_identifier.setter
    def d_b_cluster_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "d_b_cluster_identifier", value)

    @property
    @pulumi.getter(name="dBInstanceIdentifier")
    def d_b_instance_identifier(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "d_b_instance_identifier")

    @d_b_instance_identifier.setter
    def d_b_instance_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "d_b_instance_identifier", value)

    @property
    @pulumi.getter(name="dBParameterGroupName")
    def d_b_parameter_group_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "d_b_parameter_group_name")

    @d_b_parameter_group_name.setter
    def d_b_parameter_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "d_b_parameter_group_name", value)

    @property
    @pulumi.getter(name="dBSnapshotIdentifier")
    def d_b_snapshot_identifier(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "d_b_snapshot_identifier")

    @d_b_snapshot_identifier.setter
    def d_b_snapshot_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "d_b_snapshot_identifier", value)

    @property
    @pulumi.getter(name="dBSubnetGroupName")
    def d_b_subnet_group_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "d_b_subnet_group_name")

    @d_b_subnet_group_name.setter
    def d_b_subnet_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "d_b_subnet_group_name", value)

    @property
    @pulumi.getter(name="preferredMaintenanceWindow")
    def preferred_maintenance_window(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "preferred_maintenance_window")

    @preferred_maintenance_window.setter
    def preferred_maintenance_window(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "preferred_maintenance_window", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DBInstanceTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DBInstanceTagArgs']]]]):
        pulumi.set(self, "tags", value)


warnings.warn("""DBInstance is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class DBInstance(pulumi.CustomResource):
    warnings.warn("""DBInstance is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_major_version_upgrade: Optional[pulumi.Input[bool]] = None,
                 auto_minor_version_upgrade: Optional[pulumi.Input[bool]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 d_b_cluster_identifier: Optional[pulumi.Input[str]] = None,
                 d_b_instance_class: Optional[pulumi.Input[str]] = None,
                 d_b_instance_identifier: Optional[pulumi.Input[str]] = None,
                 d_b_parameter_group_name: Optional[pulumi.Input[str]] = None,
                 d_b_snapshot_identifier: Optional[pulumi.Input[str]] = None,
                 d_b_subnet_group_name: Optional[pulumi.Input[str]] = None,
                 preferred_maintenance_window: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DBInstanceTagArgs']]]]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::Neptune::DBInstance

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DBInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::Neptune::DBInstance

        :param str resource_name: The name of the resource.
        :param DBInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DBInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_major_version_upgrade: Optional[pulumi.Input[bool]] = None,
                 auto_minor_version_upgrade: Optional[pulumi.Input[bool]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 d_b_cluster_identifier: Optional[pulumi.Input[str]] = None,
                 d_b_instance_class: Optional[pulumi.Input[str]] = None,
                 d_b_instance_identifier: Optional[pulumi.Input[str]] = None,
                 d_b_parameter_group_name: Optional[pulumi.Input[str]] = None,
                 d_b_snapshot_identifier: Optional[pulumi.Input[str]] = None,
                 d_b_subnet_group_name: Optional[pulumi.Input[str]] = None,
                 preferred_maintenance_window: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DBInstanceTagArgs']]]]] = None,
                 __props__=None):
        pulumi.log.warn("""DBInstance is deprecated: DBInstance is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DBInstanceArgs.__new__(DBInstanceArgs)

            __props__.__dict__["allow_major_version_upgrade"] = allow_major_version_upgrade
            __props__.__dict__["auto_minor_version_upgrade"] = auto_minor_version_upgrade
            __props__.__dict__["availability_zone"] = availability_zone
            __props__.__dict__["d_b_cluster_identifier"] = d_b_cluster_identifier
            if d_b_instance_class is None and not opts.urn:
                raise TypeError("Missing required property 'd_b_instance_class'")
            __props__.__dict__["d_b_instance_class"] = d_b_instance_class
            __props__.__dict__["d_b_instance_identifier"] = d_b_instance_identifier
            __props__.__dict__["d_b_parameter_group_name"] = d_b_parameter_group_name
            __props__.__dict__["d_b_snapshot_identifier"] = d_b_snapshot_identifier
            __props__.__dict__["d_b_subnet_group_name"] = d_b_subnet_group_name
            __props__.__dict__["preferred_maintenance_window"] = preferred_maintenance_window
            __props__.__dict__["tags"] = tags
            __props__.__dict__["endpoint"] = None
            __props__.__dict__["port"] = None
        super(DBInstance, __self__).__init__(
            'aws-native:neptune:DBInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DBInstance':
        """
        Get an existing DBInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DBInstanceArgs.__new__(DBInstanceArgs)

        __props__.__dict__["allow_major_version_upgrade"] = None
        __props__.__dict__["auto_minor_version_upgrade"] = None
        __props__.__dict__["availability_zone"] = None
        __props__.__dict__["d_b_cluster_identifier"] = None
        __props__.__dict__["d_b_instance_class"] = None
        __props__.__dict__["d_b_instance_identifier"] = None
        __props__.__dict__["d_b_parameter_group_name"] = None
        __props__.__dict__["d_b_snapshot_identifier"] = None
        __props__.__dict__["d_b_subnet_group_name"] = None
        __props__.__dict__["endpoint"] = None
        __props__.__dict__["port"] = None
        __props__.__dict__["preferred_maintenance_window"] = None
        __props__.__dict__["tags"] = None
        return DBInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowMajorVersionUpgrade")
    def allow_major_version_upgrade(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "allow_major_version_upgrade")

    @property
    @pulumi.getter(name="autoMinorVersionUpgrade")
    def auto_minor_version_upgrade(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "auto_minor_version_upgrade")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="dBClusterIdentifier")
    def d_b_cluster_identifier(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "d_b_cluster_identifier")

    @property
    @pulumi.getter(name="dBInstanceClass")
    def d_b_instance_class(self) -> pulumi.Output[str]:
        return pulumi.get(self, "d_b_instance_class")

    @property
    @pulumi.getter(name="dBInstanceIdentifier")
    def d_b_instance_identifier(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "d_b_instance_identifier")

    @property
    @pulumi.getter(name="dBParameterGroupName")
    def d_b_parameter_group_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "d_b_parameter_group_name")

    @property
    @pulumi.getter(name="dBSnapshotIdentifier")
    def d_b_snapshot_identifier(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "d_b_snapshot_identifier")

    @property
    @pulumi.getter(name="dBSubnetGroupName")
    def d_b_subnet_group_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "d_b_subnet_group_name")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[str]:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="preferredMaintenanceWindow")
    def preferred_maintenance_window(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "preferred_maintenance_window")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.DBInstanceTag']]]:
        return pulumi.get(self, "tags")

