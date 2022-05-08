// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DataBrew::Ruleset.
 */
export function getRuleset(args: GetRulesetArgs, opts?: pulumi.InvokeOptions): Promise<GetRulesetResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:databrew:getRuleset", {
        "name": args.name,
    }, opts);
}

export interface GetRulesetArgs {
    /**
     * Name of the Ruleset
     */
    name: string;
}

export interface GetRulesetResult {
    /**
     * Description of the Ruleset
     */
    readonly description?: string;
    /**
     * List of the data quality rules in the ruleset
     */
    readonly rules?: outputs.databrew.RulesetRule[];
    readonly tags?: outputs.databrew.RulesetTag[];
}

export function getRulesetOutput(args: GetRulesetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRulesetResult> {
    return pulumi.output(args).apply(a => getRuleset(a, opts))
}

export interface GetRulesetOutputArgs {
    /**
     * Name of the Ruleset
     */
    name: pulumi.Input<string>;
}
