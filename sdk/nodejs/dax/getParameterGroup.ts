// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::DAX::ParameterGroup
 */
export function getParameterGroup(args: GetParameterGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetParameterGroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:dax:getParameterGroup", {
        "id": args.id,
    }, opts);
}

export interface GetParameterGroupArgs {
    id: string;
}

export interface GetParameterGroupResult {
    readonly description?: string;
    readonly id?: string;
    readonly parameterNameValues?: any;
}

export function getParameterGroupOutput(args: GetParameterGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetParameterGroupResult> {
    return pulumi.output(args).apply(a => getParameterGroup(a, opts))
}

export interface GetParameterGroupOutputArgs {
    id: pulumi.Input<string>;
}
