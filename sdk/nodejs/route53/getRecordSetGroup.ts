// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Route53::RecordSetGroup
 */
export function getRecordSetGroup(args: GetRecordSetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetRecordSetGroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:route53:getRecordSetGroup", {
        "id": args.id,
    }, opts);
}

export interface GetRecordSetGroupArgs {
    id: string;
}

export interface GetRecordSetGroupResult {
    readonly comment?: string;
    readonly id?: string;
    readonly recordSets?: outputs.route53.RecordSetGroupRecordSet[];
}

export function getRecordSetGroupOutput(args: GetRecordSetGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRecordSetGroupResult> {
    return pulumi.output(args).apply(a => getRecordSetGroup(a, opts))
}

export interface GetRecordSetGroupOutputArgs {
    id: pulumi.Input<string>;
}
