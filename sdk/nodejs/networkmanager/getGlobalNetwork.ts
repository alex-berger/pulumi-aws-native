// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The AWS::NetworkManager::GlobalNetwork type specifies a global network of the user's account
 */
export function getGlobalNetwork(args: GetGlobalNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetGlobalNetworkResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:networkmanager:getGlobalNetwork", {
        "id": args.id,
    }, opts);
}

export interface GetGlobalNetworkArgs {
    /**
     * The ID of the global network.
     */
    id: string;
}

export interface GetGlobalNetworkResult {
    /**
     * The Amazon Resource Name (ARN) of the global network.
     */
    readonly arn?: string;
    /**
     * The description of the global network.
     */
    readonly description?: string;
    /**
     * The ID of the global network.
     */
    readonly id?: string;
    /**
     * The tags for the global network.
     */
    readonly tags?: outputs.networkmanager.GlobalNetworkTag[];
}

export function getGlobalNetworkOutput(args: GetGlobalNetworkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGlobalNetworkResult> {
    return pulumi.output(args).apply(a => getGlobalNetwork(a, opts))
}

export interface GetGlobalNetworkOutputArgs {
    /**
     * The ID of the global network.
     */
    id: pulumi.Input<string>;
}
