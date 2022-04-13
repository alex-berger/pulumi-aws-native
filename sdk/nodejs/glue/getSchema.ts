// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * This resource represents a schema of Glue Schema Registry.
 */
export function getSchema(args: GetSchemaArgs, opts?: pulumi.InvokeOptions): Promise<GetSchemaResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:glue:getSchema", {
        "arn": args.arn,
    }, opts);
}

export interface GetSchemaArgs {
    /**
     * Amazon Resource Name for the Schema.
     */
    arn: string;
}

export interface GetSchemaResult {
    /**
     * Amazon Resource Name for the Schema.
     */
    readonly arn?: string;
    readonly checkpointVersion?: outputs.glue.SchemaVersion;
    /**
     * Compatibility setting for the schema.
     */
    readonly compatibility?: enums.glue.SchemaCompatibility;
    /**
     * A description of the schema. If description is not provided, there will not be any default value for this.
     */
    readonly description?: string;
    /**
     * Represents the version ID associated with the initial schema version.
     */
    readonly initialSchemaVersionId?: string;
    /**
     * List of tags to tag the schema
     */
    readonly tags?: outputs.glue.SchemaTag[];
}

export function getSchemaOutput(args: GetSchemaOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSchemaResult> {
    return pulumi.output(args).apply(a => getSchema(a, opts))
}

export interface GetSchemaOutputArgs {
    /**
     * Amazon Resource Name for the Schema.
     */
    arn: pulumi.Input<string>;
}
