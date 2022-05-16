// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Contains the Rules that identify the requests that you want to allow, block, or count. In a RuleGroup, you also specify a default action (ALLOW or BLOCK), and the action for each Rule that you add to a RuleGroup, for example, block requests from specified IP addresses or block requests from specified referrers. You also associate the RuleGroup with a CloudFront distribution to identify the requests that you want AWS WAF to filter. If you add more than one Rule to a RuleGroup, a request needs to match only one of the specifications to be allowed, blocked, or counted.
 */
export function getRuleGroup(args: GetRuleGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetRuleGroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:wafv2:getRuleGroup", {
        "id": args.id,
        "name": args.name,
        "scope": args.scope,
    }, opts);
}

export interface GetRuleGroupArgs {
    id: string;
    name: string;
    scope: enums.wafv2.RuleGroupScope;
}

export interface GetRuleGroupResult {
    readonly arn?: string;
    /**
     * Collection of Available Labels.
     */
    readonly availableLabels?: outputs.wafv2.RuleGroupLabelSummary[];
    readonly capacity?: number;
    /**
     * Collection of Consumed Labels.
     */
    readonly consumedLabels?: outputs.wafv2.RuleGroupLabelSummary[];
    readonly customResponseBodies?: outputs.wafv2.RuleGroupCustomResponseBodies;
    readonly description?: string;
    readonly id?: string;
    readonly labelNamespace?: string;
    /**
     * Collection of Rules.
     */
    readonly rules?: outputs.wafv2.RuleGroupRule[];
    readonly tags?: outputs.wafv2.RuleGroupTag[];
    readonly visibilityConfig?: outputs.wafv2.RuleGroupVisibilityConfig;
}

export function getRuleGroupOutput(args: GetRuleGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRuleGroupResult> {
    return pulumi.output(args).apply(a => getRuleGroup(a, opts))
}

export interface GetRuleGroupOutputArgs {
    id: pulumi.Input<string>;
    name: pulumi.Input<string>;
    scope: pulumi.Input<enums.wafv2.RuleGroupScope>;
}