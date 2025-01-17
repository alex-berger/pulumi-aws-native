// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SES::ReceiptRule
 */
export function getReceiptRule(args: GetReceiptRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetReceiptRuleResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ses:getReceiptRule", {
        "id": args.id,
    }, opts);
}

export interface GetReceiptRuleArgs {
    id: string;
}

export interface GetReceiptRuleResult {
    readonly after?: string;
    readonly id?: string;
    readonly rule?: outputs.ses.ReceiptRuleRule;
}

export function getReceiptRuleOutput(args: GetReceiptRuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReceiptRuleResult> {
    return pulumi.output(args).apply(a => getReceiptRule(a, opts))
}

export interface GetReceiptRuleOutputArgs {
    id: pulumi.Input<string>;
}
