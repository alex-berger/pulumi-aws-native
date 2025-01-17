// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Schema for AWS::EKS::Addon
 */
export function getAddon(args: GetAddonArgs, opts?: pulumi.InvokeOptions): Promise<GetAddonResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:eks:getAddon", {
        "addonName": args.addonName,
        "clusterName": args.clusterName,
    }, opts);
}

export interface GetAddonArgs {
    /**
     * Name of Addon
     */
    addonName: string;
    /**
     * Name of Cluster
     */
    clusterName: string;
}

export interface GetAddonResult {
    /**
     * Version of Addon
     */
    readonly addonVersion?: string;
    /**
     * Amazon Resource Name (ARN) of the add-on
     */
    readonly arn?: string;
    /**
     * IAM role to bind to the add-on's service account
     */
    readonly serviceAccountRoleArn?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.eks.AddonTag[];
}

export function getAddonOutput(args: GetAddonOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAddonResult> {
    return pulumi.output(args).apply(a => getAddon(a, opts))
}

export interface GetAddonOutputArgs {
    /**
     * Name of Addon
     */
    addonName: pulumi.Input<string>;
    /**
     * Name of Cluster
     */
    clusterName: pulumi.Input<string>;
}
