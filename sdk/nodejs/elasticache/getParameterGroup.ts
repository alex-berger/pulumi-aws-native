// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ElastiCache::ParameterGroup
 */
export function getParameterGroup(args: GetParameterGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetParameterGroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:elasticache:getParameterGroup", {
        "id": args.id,
    }, opts);
}

export interface GetParameterGroupArgs {
    id: string;
}

export interface GetParameterGroupResult {
    readonly description?: string;
    readonly id?: string;
    readonly properties?: any;
    readonly tags?: outputs.elasticache.ParameterGroupTag[];
}

export function getParameterGroupOutput(args: GetParameterGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetParameterGroupResult> {
    return pulumi.output(args).apply(a => getParameterGroup(a, opts))
}

export interface GetParameterGroupOutputArgs {
    id: pulumi.Input<string>;
}
