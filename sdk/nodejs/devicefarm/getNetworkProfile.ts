// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * AWS::DeviceFarm::NetworkProfile creates a new DF Network Profile
 */
export function getNetworkProfile(args: GetNetworkProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkProfileResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:devicefarm:getNetworkProfile", {
        "arn": args.arn,
    }, opts);
}

export interface GetNetworkProfileArgs {
    arn: string;
}

export interface GetNetworkProfileResult {
    readonly arn?: string;
    readonly description?: string;
    readonly downlinkBandwidthBits?: number;
    readonly downlinkDelayMs?: number;
    readonly downlinkJitterMs?: number;
    readonly downlinkLossPercent?: number;
    readonly name?: string;
    readonly tags?: outputs.devicefarm.NetworkProfileTag[];
    readonly uplinkBandwidthBits?: number;
    readonly uplinkDelayMs?: number;
    readonly uplinkJitterMs?: number;
    readonly uplinkLossPercent?: number;
}

export function getNetworkProfileOutput(args: GetNetworkProfileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetworkProfileResult> {
    return pulumi.output(args).apply(a => getNetworkProfile(a, opts))
}

export interface GetNetworkProfileOutputArgs {
    arn: pulumi.Input<string>;
}
