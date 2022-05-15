// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::MediaLive::InputSecurityGroup
 */
export function getInputSecurityGroup(args: GetInputSecurityGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetInputSecurityGroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:medialive:getInputSecurityGroup", {
        "id": args.id,
    }, opts);
}

export interface GetInputSecurityGroupArgs {
    id: string;
}

export interface GetInputSecurityGroupResult {
    readonly arn?: string;
    readonly id?: string;
    readonly tags?: any;
    readonly whitelistRules?: outputs.medialive.InputSecurityGroupInputWhitelistRuleCidr[];
}

export function getInputSecurityGroupOutput(args: GetInputSecurityGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInputSecurityGroupResult> {
    return pulumi.output(args).apply(a => getInputSecurityGroup(a, opts))
}

export interface GetInputSecurityGroupOutputArgs {
    id: pulumi.Input<string>;
}
