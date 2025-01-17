// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::Host
 */
export function getHost(args: GetHostArgs, opts?: pulumi.InvokeOptions): Promise<GetHostResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ec2:getHost", {
        "hostId": args.hostId,
    }, opts);
}

export interface GetHostArgs {
    /**
     * Id of the host created.
     */
    hostId: string;
}

export interface GetHostResult {
    /**
     * Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.
     */
    readonly autoPlacement?: string;
    /**
     * Id of the host created.
     */
    readonly hostId?: string;
    /**
     * Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default.
     */
    readonly hostRecovery?: string;
}

export function getHostOutput(args: GetHostOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetHostResult> {
    return pulumi.output(args).apply(a => getHost(a, opts))
}

export interface GetHostOutputArgs {
    /**
     * Id of the host created.
     */
    hostId: pulumi.Input<string>;
}
