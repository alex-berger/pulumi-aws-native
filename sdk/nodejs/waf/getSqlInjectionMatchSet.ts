// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::WAF::SqlInjectionMatchSet
 */
export function getSqlInjectionMatchSet(args: GetSqlInjectionMatchSetArgs, opts?: pulumi.InvokeOptions): Promise<GetSqlInjectionMatchSetResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:waf:getSqlInjectionMatchSet", {
        "id": args.id,
    }, opts);
}

export interface GetSqlInjectionMatchSetArgs {
    id: string;
}

export interface GetSqlInjectionMatchSetResult {
    readonly id?: string;
    readonly sqlInjectionMatchTuples?: outputs.waf.SqlInjectionMatchSetSqlInjectionMatchTuple[];
}

export function getSqlInjectionMatchSetOutput(args: GetSqlInjectionMatchSetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSqlInjectionMatchSetResult> {
    return pulumi.output(args).apply(a => getSqlInjectionMatchSet(a, opts))
}

export interface GetSqlInjectionMatchSetOutputArgs {
    id: pulumi.Input<string>;
}