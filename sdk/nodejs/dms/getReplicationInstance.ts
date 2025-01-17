// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::DMS::ReplicationInstance
 */
export function getReplicationInstance(args: GetReplicationInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetReplicationInstanceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:dms:getReplicationInstance", {
        "id": args.id,
    }, opts);
}

export interface GetReplicationInstanceArgs {
    id: string;
}

export interface GetReplicationInstanceResult {
    readonly allocatedStorage?: number;
    readonly allowMajorVersionUpgrade?: boolean;
    readonly autoMinorVersionUpgrade?: boolean;
    readonly availabilityZone?: string;
    readonly engineVersion?: string;
    readonly id?: string;
    readonly multiAZ?: boolean;
    readonly preferredMaintenanceWindow?: string;
    readonly replicationInstanceClass?: string;
    readonly replicationInstanceIdentifier?: string;
    readonly replicationInstancePrivateIpAddresses?: string;
    readonly replicationInstancePublicIpAddresses?: string;
    readonly vpcSecurityGroupIds?: string[];
}

export function getReplicationInstanceOutput(args: GetReplicationInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReplicationInstanceResult> {
    return pulumi.output(args).apply(a => getReplicationInstance(a, opts))
}

export interface GetReplicationInstanceOutputArgs {
    id: pulumi.Input<string>;
}
