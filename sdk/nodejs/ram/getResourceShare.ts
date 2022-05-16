// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::RAM::ResourceShare
 */
export function getResourceShare(args: GetResourceShareArgs, opts?: pulumi.InvokeOptions): Promise<GetResourceShareResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ram:getResourceShare", {
        "id": args.id,
    }, opts);
}

export interface GetResourceShareArgs {
    id: string;
}

export interface GetResourceShareResult {
    readonly allowExternalPrincipals?: boolean;
    readonly arn?: string;
    readonly id?: string;
    readonly name?: string;
    readonly permissionArns?: string[];
    readonly principals?: string[];
    readonly resourceArns?: string[];
    readonly tags?: outputs.ram.ResourceShareTag[];
}

export function getResourceShareOutput(args: GetResourceShareOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetResourceShareResult> {
    return pulumi.output(args).apply(a => getResourceShare(a, opts))
}

export interface GetResourceShareOutputArgs {
    id: pulumi.Input<string>;
}
