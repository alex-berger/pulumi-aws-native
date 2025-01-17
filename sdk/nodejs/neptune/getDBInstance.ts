// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Neptune::DBInstance
 */
export function getDBInstance(args: GetDBInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetDBInstanceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:neptune:getDBInstance", {
        "id": args.id,
    }, opts);
}

export interface GetDBInstanceArgs {
    id: string;
}

export interface GetDBInstanceResult {
    readonly allowMajorVersionUpgrade?: boolean;
    readonly autoMinorVersionUpgrade?: boolean;
    readonly dBInstanceClass?: string;
    readonly dBParameterGroupName?: string;
    readonly endpoint?: string;
    readonly id?: string;
    readonly port?: string;
    readonly preferredMaintenanceWindow?: string;
    readonly tags?: outputs.neptune.DBInstanceTag[];
}

export function getDBInstanceOutput(args: GetDBInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDBInstanceResult> {
    return pulumi.output(args).apply(a => getDBInstance(a, opts))
}

export interface GetDBInstanceOutputArgs {
    id: pulumi.Input<string>;
}
