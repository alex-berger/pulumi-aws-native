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

__all__ = ['GlobalTableArgs', 'GlobalTable']

@pulumi.input_type
class GlobalTableArgs:
    def __init__(__self__, *,
                 attribute_definitions: pulumi.Input[Sequence[pulumi.Input['GlobalTableAttributeDefinitionArgs']]],
                 key_schema: pulumi.Input[Sequence[pulumi.Input['GlobalTableKeySchemaArgs']]],
                 replicas: pulumi.Input[Sequence[pulumi.Input['GlobalTableReplicaSpecificationArgs']]],
                 billing_mode: Optional[pulumi.Input[str]] = None,
                 global_secondary_indexes: Optional[pulumi.Input[Sequence[pulumi.Input['GlobalTableGlobalSecondaryIndexArgs']]]] = None,
                 local_secondary_indexes: Optional[pulumi.Input[Sequence[pulumi.Input['GlobalTableLocalSecondaryIndexArgs']]]] = None,
                 s_se_specification: Optional[pulumi.Input['GlobalTableSSESpecificationArgs']] = None,
                 stream_specification: Optional[pulumi.Input['GlobalTableStreamSpecificationArgs']] = None,
                 table_name: Optional[pulumi.Input[str]] = None,
                 time_to_live_specification: Optional[pulumi.Input['GlobalTableTimeToLiveSpecificationArgs']] = None,
                 write_provisioned_throughput_settings: Optional[pulumi.Input['GlobalTableWriteProvisionedThroughputSettingsArgs']] = None):
        """
        The set of arguments for constructing a GlobalTable resource.
        """
        pulumi.set(__self__, "attribute_definitions", attribute_definitions)
        pulumi.set(__self__, "key_schema", key_schema)
        pulumi.set(__self__, "replicas", replicas)
        if billing_mode is not None:
            pulumi.set(__self__, "billing_mode", billing_mode)
        if global_secondary_indexes is not None:
            pulumi.set(__self__, "global_secondary_indexes", global_secondary_indexes)
        if local_secondary_indexes is not None:
            pulumi.set(__self__, "local_secondary_indexes", local_secondary_indexes)
        if s_se_specification is not None:
            pulumi.set(__self__, "s_se_specification", s_se_specification)
        if stream_specification is not None:
            pulumi.set(__self__, "stream_specification", stream_specification)
        if table_name is not None:
            pulumi.set(__self__, "table_name", table_name)
        if time_to_live_specification is not None:
            pulumi.set(__self__, "time_to_live_specification", time_to_live_specification)
        if write_provisioned_throughput_settings is not None:
            pulumi.set(__self__, "write_provisioned_throughput_settings", write_provisioned_throughput_settings)

    @property
    @pulumi.getter(name="attributeDefinitions")
    def attribute_definitions(self) -> pulumi.Input[Sequence[pulumi.Input['GlobalTableAttributeDefinitionArgs']]]:
        return pulumi.get(self, "attribute_definitions")

    @attribute_definitions.setter
    def attribute_definitions(self, value: pulumi.Input[Sequence[pulumi.Input['GlobalTableAttributeDefinitionArgs']]]):
        pulumi.set(self, "attribute_definitions", value)

    @property
    @pulumi.getter(name="keySchema")
    def key_schema(self) -> pulumi.Input[Sequence[pulumi.Input['GlobalTableKeySchemaArgs']]]:
        return pulumi.get(self, "key_schema")

    @key_schema.setter
    def key_schema(self, value: pulumi.Input[Sequence[pulumi.Input['GlobalTableKeySchemaArgs']]]):
        pulumi.set(self, "key_schema", value)

    @property
    @pulumi.getter
    def replicas(self) -> pulumi.Input[Sequence[pulumi.Input['GlobalTableReplicaSpecificationArgs']]]:
        return pulumi.get(self, "replicas")

    @replicas.setter
    def replicas(self, value: pulumi.Input[Sequence[pulumi.Input['GlobalTableReplicaSpecificationArgs']]]):
        pulumi.set(self, "replicas", value)

    @property
    @pulumi.getter(name="billingMode")
    def billing_mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "billing_mode")

    @billing_mode.setter
    def billing_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "billing_mode", value)

    @property
    @pulumi.getter(name="globalSecondaryIndexes")
    def global_secondary_indexes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GlobalTableGlobalSecondaryIndexArgs']]]]:
        return pulumi.get(self, "global_secondary_indexes")

    @global_secondary_indexes.setter
    def global_secondary_indexes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GlobalTableGlobalSecondaryIndexArgs']]]]):
        pulumi.set(self, "global_secondary_indexes", value)

    @property
    @pulumi.getter(name="localSecondaryIndexes")
    def local_secondary_indexes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GlobalTableLocalSecondaryIndexArgs']]]]:
        return pulumi.get(self, "local_secondary_indexes")

    @local_secondary_indexes.setter
    def local_secondary_indexes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GlobalTableLocalSecondaryIndexArgs']]]]):
        pulumi.set(self, "local_secondary_indexes", value)

    @property
    @pulumi.getter(name="sSESpecification")
    def s_se_specification(self) -> Optional[pulumi.Input['GlobalTableSSESpecificationArgs']]:
        return pulumi.get(self, "s_se_specification")

    @s_se_specification.setter
    def s_se_specification(self, value: Optional[pulumi.Input['GlobalTableSSESpecificationArgs']]):
        pulumi.set(self, "s_se_specification", value)

    @property
    @pulumi.getter(name="streamSpecification")
    def stream_specification(self) -> Optional[pulumi.Input['GlobalTableStreamSpecificationArgs']]:
        return pulumi.get(self, "stream_specification")

    @stream_specification.setter
    def stream_specification(self, value: Optional[pulumi.Input['GlobalTableStreamSpecificationArgs']]):
        pulumi.set(self, "stream_specification", value)

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "table_name")

    @table_name.setter
    def table_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "table_name", value)

    @property
    @pulumi.getter(name="timeToLiveSpecification")
    def time_to_live_specification(self) -> Optional[pulumi.Input['GlobalTableTimeToLiveSpecificationArgs']]:
        return pulumi.get(self, "time_to_live_specification")

    @time_to_live_specification.setter
    def time_to_live_specification(self, value: Optional[pulumi.Input['GlobalTableTimeToLiveSpecificationArgs']]):
        pulumi.set(self, "time_to_live_specification", value)

    @property
    @pulumi.getter(name="writeProvisionedThroughputSettings")
    def write_provisioned_throughput_settings(self) -> Optional[pulumi.Input['GlobalTableWriteProvisionedThroughputSettingsArgs']]:
        return pulumi.get(self, "write_provisioned_throughput_settings")

    @write_provisioned_throughput_settings.setter
    def write_provisioned_throughput_settings(self, value: Optional[pulumi.Input['GlobalTableWriteProvisionedThroughputSettingsArgs']]):
        pulumi.set(self, "write_provisioned_throughput_settings", value)


class GlobalTable(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attribute_definitions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GlobalTableAttributeDefinitionArgs']]]]] = None,
                 billing_mode: Optional[pulumi.Input[str]] = None,
                 global_secondary_indexes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GlobalTableGlobalSecondaryIndexArgs']]]]] = None,
                 key_schema: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GlobalTableKeySchemaArgs']]]]] = None,
                 local_secondary_indexes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GlobalTableLocalSecondaryIndexArgs']]]]] = None,
                 replicas: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GlobalTableReplicaSpecificationArgs']]]]] = None,
                 s_se_specification: Optional[pulumi.Input[pulumi.InputType['GlobalTableSSESpecificationArgs']]] = None,
                 stream_specification: Optional[pulumi.Input[pulumi.InputType['GlobalTableStreamSpecificationArgs']]] = None,
                 table_name: Optional[pulumi.Input[str]] = None,
                 time_to_live_specification: Optional[pulumi.Input[pulumi.InputType['GlobalTableTimeToLiveSpecificationArgs']]] = None,
                 write_provisioned_throughput_settings: Optional[pulumi.Input[pulumi.InputType['GlobalTableWriteProvisionedThroughputSettingsArgs']]] = None,
                 __props__=None):
        """
        Version: None. Resource Type definition for AWS::DynamoDB::GlobalTable

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GlobalTableArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Version: None. Resource Type definition for AWS::DynamoDB::GlobalTable

        :param str resource_name: The name of the resource.
        :param GlobalTableArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GlobalTableArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attribute_definitions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GlobalTableAttributeDefinitionArgs']]]]] = None,
                 billing_mode: Optional[pulumi.Input[str]] = None,
                 global_secondary_indexes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GlobalTableGlobalSecondaryIndexArgs']]]]] = None,
                 key_schema: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GlobalTableKeySchemaArgs']]]]] = None,
                 local_secondary_indexes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GlobalTableLocalSecondaryIndexArgs']]]]] = None,
                 replicas: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GlobalTableReplicaSpecificationArgs']]]]] = None,
                 s_se_specification: Optional[pulumi.Input[pulumi.InputType['GlobalTableSSESpecificationArgs']]] = None,
                 stream_specification: Optional[pulumi.Input[pulumi.InputType['GlobalTableStreamSpecificationArgs']]] = None,
                 table_name: Optional[pulumi.Input[str]] = None,
                 time_to_live_specification: Optional[pulumi.Input[pulumi.InputType['GlobalTableTimeToLiveSpecificationArgs']]] = None,
                 write_provisioned_throughput_settings: Optional[pulumi.Input[pulumi.InputType['GlobalTableWriteProvisionedThroughputSettingsArgs']]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GlobalTableArgs.__new__(GlobalTableArgs)

            if attribute_definitions is None and not opts.urn:
                raise TypeError("Missing required property 'attribute_definitions'")
            __props__.__dict__["attribute_definitions"] = attribute_definitions
            __props__.__dict__["billing_mode"] = billing_mode
            __props__.__dict__["global_secondary_indexes"] = global_secondary_indexes
            if key_schema is None and not opts.urn:
                raise TypeError("Missing required property 'key_schema'")
            __props__.__dict__["key_schema"] = key_schema
            __props__.__dict__["local_secondary_indexes"] = local_secondary_indexes
            if replicas is None and not opts.urn:
                raise TypeError("Missing required property 'replicas'")
            __props__.__dict__["replicas"] = replicas
            __props__.__dict__["s_se_specification"] = s_se_specification
            __props__.__dict__["stream_specification"] = stream_specification
            __props__.__dict__["table_name"] = table_name
            __props__.__dict__["time_to_live_specification"] = time_to_live_specification
            __props__.__dict__["write_provisioned_throughput_settings"] = write_provisioned_throughput_settings
            __props__.__dict__["arn"] = None
            __props__.__dict__["stream_arn"] = None
            __props__.__dict__["table_id"] = None
        super(GlobalTable, __self__).__init__(
            'aws-native:dynamodb:GlobalTable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'GlobalTable':
        """
        Get an existing GlobalTable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = GlobalTableArgs.__new__(GlobalTableArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["attribute_definitions"] = None
        __props__.__dict__["billing_mode"] = None
        __props__.__dict__["global_secondary_indexes"] = None
        __props__.__dict__["key_schema"] = None
        __props__.__dict__["local_secondary_indexes"] = None
        __props__.__dict__["replicas"] = None
        __props__.__dict__["s_se_specification"] = None
        __props__.__dict__["stream_arn"] = None
        __props__.__dict__["stream_specification"] = None
        __props__.__dict__["table_id"] = None
        __props__.__dict__["table_name"] = None
        __props__.__dict__["time_to_live_specification"] = None
        __props__.__dict__["write_provisioned_throughput_settings"] = None
        return GlobalTable(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="attributeDefinitions")
    def attribute_definitions(self) -> pulumi.Output[Sequence['outputs.GlobalTableAttributeDefinition']]:
        return pulumi.get(self, "attribute_definitions")

    @property
    @pulumi.getter(name="billingMode")
    def billing_mode(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "billing_mode")

    @property
    @pulumi.getter(name="globalSecondaryIndexes")
    def global_secondary_indexes(self) -> pulumi.Output[Optional[Sequence['outputs.GlobalTableGlobalSecondaryIndex']]]:
        return pulumi.get(self, "global_secondary_indexes")

    @property
    @pulumi.getter(name="keySchema")
    def key_schema(self) -> pulumi.Output[Sequence['outputs.GlobalTableKeySchema']]:
        return pulumi.get(self, "key_schema")

    @property
    @pulumi.getter(name="localSecondaryIndexes")
    def local_secondary_indexes(self) -> pulumi.Output[Optional[Sequence['outputs.GlobalTableLocalSecondaryIndex']]]:
        return pulumi.get(self, "local_secondary_indexes")

    @property
    @pulumi.getter
    def replicas(self) -> pulumi.Output[Sequence['outputs.GlobalTableReplicaSpecification']]:
        return pulumi.get(self, "replicas")

    @property
    @pulumi.getter(name="sSESpecification")
    def s_se_specification(self) -> pulumi.Output[Optional['outputs.GlobalTableSSESpecification']]:
        return pulumi.get(self, "s_se_specification")

    @property
    @pulumi.getter(name="streamArn")
    def stream_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "stream_arn")

    @property
    @pulumi.getter(name="streamSpecification")
    def stream_specification(self) -> pulumi.Output[Optional['outputs.GlobalTableStreamSpecification']]:
        return pulumi.get(self, "stream_specification")

    @property
    @pulumi.getter(name="tableId")
    def table_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "table_id")

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "table_name")

    @property
    @pulumi.getter(name="timeToLiveSpecification")
    def time_to_live_specification(self) -> pulumi.Output[Optional['outputs.GlobalTableTimeToLiveSpecification']]:
        return pulumi.get(self, "time_to_live_specification")

    @property
    @pulumi.getter(name="writeProvisionedThroughputSettings")
    def write_provisioned_throughput_settings(self) -> pulumi.Output[Optional['outputs.GlobalTableWriteProvisionedThroughputSettings']]:
        return pulumi.get(self, "write_provisioned_throughput_settings")

