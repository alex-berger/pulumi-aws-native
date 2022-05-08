// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::RDS::GlobalCluster
 */
export function getGlobalCluster(args: GetGlobalClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetGlobalClusterResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:rds:getGlobalCluster", {
        "globalClusterIdentifier": args.globalClusterIdentifier,
    }, opts);
}

export interface GetGlobalClusterArgs {
    /**
     * The cluster identifier of the new global database cluster. This parameter is stored as a lowercase string.
     */
    globalClusterIdentifier: string;
}

export interface GetGlobalClusterResult {
    /**
     * The deletion protection setting for the new global database. The global database can't be deleted when deletion protection is enabled.
     */
    readonly deletionProtection?: boolean;
}

export function getGlobalClusterOutput(args: GetGlobalClusterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGlobalClusterResult> {
    return pulumi.output(args).apply(a => getGlobalCluster(a, opts))
}

export interface GetGlobalClusterOutputArgs {
    /**
     * The cluster identifier of the new global database cluster. This parameter is stored as a lowercase string.
     */
    globalClusterIdentifier: pulumi.Input<string>;
}
