// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::RDS::DBSecurityGroup
 */
export function getDBSecurityGroup(args: GetDBSecurityGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetDBSecurityGroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:rds:getDBSecurityGroup", {
        "id": args.id,
    }, opts);
}

export interface GetDBSecurityGroupArgs {
    id: string;
}

export interface GetDBSecurityGroupResult {
    readonly dBSecurityGroupIngress?: outputs.rds.DBSecurityGroupIngress[];
    readonly id?: string;
    readonly tags?: outputs.rds.DBSecurityGroupTag[];
}

export function getDBSecurityGroupOutput(args: GetDBSecurityGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDBSecurityGroupResult> {
    return pulumi.output(args).apply(a => getDBSecurityGroup(a, opts))
}

export interface GetDBSecurityGroupOutputArgs {
    id: pulumi.Input<string>;
}
