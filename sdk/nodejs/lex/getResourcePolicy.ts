// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * A resource policy with specified policy statements that attaches to a Lex bot or bot alias.
 */
export function getResourcePolicy(args: GetResourcePolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetResourcePolicyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:lex:getResourcePolicy", {
        "id": args.id,
    }, opts);
}

export interface GetResourcePolicyArgs {
    id: string;
}

export interface GetResourcePolicyResult {
    readonly id?: string;
    readonly policy?: outputs.lex.ResourcePolicyPolicy;
    readonly resourceArn?: string;
    readonly revisionId?: string;
}

export function getResourcePolicyOutput(args: GetResourcePolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetResourcePolicyResult> {
    return pulumi.output(args).apply(a => getResourcePolicy(a, opts))
}

export interface GetResourcePolicyOutputArgs {
    id: pulumi.Input<string>;
}
