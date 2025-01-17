// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Redshift::ClusterSecurityGroupIngress
 */
export function getClusterSecurityGroupIngress(args: GetClusterSecurityGroupIngressArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterSecurityGroupIngressResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:redshift:getClusterSecurityGroupIngress", {
        "id": args.id,
    }, opts);
}

export interface GetClusterSecurityGroupIngressArgs {
    id: string;
}

export interface GetClusterSecurityGroupIngressResult {
    readonly id?: string;
}

export function getClusterSecurityGroupIngressOutput(args: GetClusterSecurityGroupIngressOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClusterSecurityGroupIngressResult> {
    return pulumi.output(args).apply(a => getClusterSecurityGroupIngress(a, opts))
}

export interface GetClusterSecurityGroupIngressOutputArgs {
    id: pulumi.Input<string>;
}
