// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::VPCEndpointServicePermissions
 */
export function getVPCEndpointServicePermissions(args: GetVPCEndpointServicePermissionsArgs, opts?: pulumi.InvokeOptions): Promise<GetVPCEndpointServicePermissionsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ec2:getVPCEndpointServicePermissions", {
        "id": args.id,
    }, opts);
}

export interface GetVPCEndpointServicePermissionsArgs {
    id: string;
}

export interface GetVPCEndpointServicePermissionsResult {
    readonly allowedPrincipals?: string[];
    readonly id?: string;
}

export function getVPCEndpointServicePermissionsOutput(args: GetVPCEndpointServicePermissionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVPCEndpointServicePermissionsResult> {
    return pulumi.output(args).apply(a => getVPCEndpointServicePermissions(a, opts))
}

export interface GetVPCEndpointServicePermissionsOutputArgs {
    id: pulumi.Input<string>;
}
