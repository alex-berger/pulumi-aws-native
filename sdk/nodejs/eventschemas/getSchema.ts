// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EventSchemas::Schema
 */
export function getSchema(args: GetSchemaArgs, opts?: pulumi.InvokeOptions): Promise<GetSchemaResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:eventschemas:getSchema", {
        "id": args.id,
    }, opts);
}

export interface GetSchemaArgs {
    id: string;
}

export interface GetSchemaResult {
    readonly content?: string;
    readonly description?: string;
    readonly id?: string;
    readonly schemaArn?: string;
    readonly schemaVersion?: string;
    readonly tags?: outputs.eventschemas.SchemaTagsEntry[];
    readonly type?: string;
}

export function getSchemaOutput(args: GetSchemaOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSchemaResult> {
    return pulumi.output(args).apply(a => getSchema(a, opts))
}

export interface GetSchemaOutputArgs {
    id: pulumi.Input<string>;
}
