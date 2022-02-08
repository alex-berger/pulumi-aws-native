// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::Subnet
 */
export function getSubnet(args: GetSubnetArgs, opts?: pulumi.InvokeOptions): Promise<GetSubnetResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ec2:getSubnet", {
        "subnetId": args.subnetId,
    }, opts);
}

export interface GetSubnetArgs {
    subnetId: string;
}

export interface GetSubnetResult {
    readonly assignIpv6AddressOnCreation?: boolean;
    readonly ipv6CidrBlock?: string;
    readonly ipv6CidrBlocks?: string[];
    readonly mapPublicIpOnLaunch?: boolean;
    readonly networkAclAssociationId?: string;
    readonly subnetId?: string;
    readonly tags?: outputs.ec2.SubnetTag[];
}

export function getSubnetOutput(args: GetSubnetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSubnetResult> {
    return pulumi.output(args).apply(a => getSubnet(a, opts))
}

export interface GetSubnetOutputArgs {
    subnetId: pulumi.Input<string>;
}