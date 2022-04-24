// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::DHCPOptions
 */
export function getDHCPOptions(args: GetDHCPOptionsArgs, opts?: pulumi.InvokeOptions): Promise<GetDHCPOptionsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ec2:getDHCPOptions", {
        "dhcpOptionsId": args.dhcpOptionsId,
    }, opts);
}

export interface GetDHCPOptionsArgs {
    dhcpOptionsId: string;
}

export interface GetDHCPOptionsResult {
    readonly dhcpOptionsId?: string;
    /**
     * Any tags assigned to the DHCP options set.
     */
    readonly tags?: outputs.ec2.DHCPOptionsTag[];
}

export function getDHCPOptionsOutput(args: GetDHCPOptionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDHCPOptionsResult> {
    return pulumi.output(args).apply(a => getDHCPOptions(a, opts))
}

export interface GetDHCPOptionsOutputArgs {
    dhcpOptionsId: pulumi.Input<string>;
}
