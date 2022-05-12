// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The AWS::Timestream::Database resource creates a Timestream database.
 */
export function getDatabase(args: GetDatabaseArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabaseResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:timestream:getDatabase", {
        "databaseName": args.databaseName,
    }, opts);
}

export interface GetDatabaseArgs {
    /**
     * The name for the database. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the database name.
     */
    databaseName: string;
}

export interface GetDatabaseResult {
    readonly arn?: string;
    /**
     * The KMS key for the database. If the KMS key is not specified, the database will be encrypted with a Timestream managed KMS key located in your account.
     */
    readonly kmsKeyId?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.timestream.DatabaseTag[];
}

export function getDatabaseOutput(args: GetDatabaseOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatabaseResult> {
    return pulumi.output(args).apply(a => getDatabase(a, opts))
}

export interface GetDatabaseOutputArgs {
    /**
     * The name for the database. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the database name.
     */
    databaseName: pulumi.Input<string>;
}
