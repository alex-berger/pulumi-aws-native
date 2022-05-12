// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Associates a set of DHCP options with a VPC, or associates no DHCP options with the VPC.
 */
export function getVPCDHCPOptionsAssociation(args: GetVPCDHCPOptionsAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetVPCDHCPOptionsAssociationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ec2:getVPCDHCPOptionsAssociation", {
        "dhcpOptionsId": args.dhcpOptionsId,
        "vpcId": args.vpcId,
    }, opts);
}

export interface GetVPCDHCPOptionsAssociationArgs {
    /**
     * The ID of the DHCP options set, or default to associate no DHCP options with the VPC.
     */
    dhcpOptionsId: string;
    /**
     * The ID of the VPC.
     */
    vpcId: string;
}

export interface GetVPCDHCPOptionsAssociationResult {
    /**
     * The ID of the VPC DHCPOptions Association.
     */
    readonly id?: string;
}

export function getVPCDHCPOptionsAssociationOutput(args: GetVPCDHCPOptionsAssociationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVPCDHCPOptionsAssociationResult> {
    return pulumi.output(args).apply(a => getVPCDHCPOptionsAssociation(a, opts))
}

export interface GetVPCDHCPOptionsAssociationOutputArgs {
    /**
     * The ID of the DHCP options set, or default to associate no DHCP options with the VPC.
     */
    dhcpOptionsId: pulumi.Input<string>;
    /**
     * The ID of the VPC.
     */
    vpcId: pulumi.Input<string>;
}
