// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Events::Rule
 */
export function getRule(args: GetRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetRuleResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:events:getRule", {
        "id": args.id,
    }, opts);
}

export interface GetRuleArgs {
    id: string;
}

export interface GetRuleResult {
    readonly arn?: string;
    readonly description?: string;
    readonly eventPattern?: any;
    readonly id?: string;
    readonly roleArn?: string;
    readonly scheduleExpression?: string;
    readonly state?: string;
    readonly targets?: outputs.events.RuleTarget[];
}

export function getRuleOutput(args: GetRuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRuleResult> {
    return pulumi.output(args).apply(a => getRule(a, opts))
}

export interface GetRuleOutputArgs {
    id: pulumi.Input<string>;
}