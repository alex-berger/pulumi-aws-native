// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::SubnetRouteTableAssociation
 */
export function getSubnetRouteTableAssociation(args: GetSubnetRouteTableAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetSubnetRouteTableAssociationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ec2:getSubnetRouteTableAssociation", {
        "id": args.id,
    }, opts);
}

export interface GetSubnetRouteTableAssociationArgs {
    id: string;
}

export interface GetSubnetRouteTableAssociationResult {
    readonly id?: string;
}

export function getSubnetRouteTableAssociationOutput(args: GetSubnetRouteTableAssociationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSubnetRouteTableAssociationResult> {
    return pulumi.output(args).apply(a => getSubnetRouteTableAssociation(a, opts))
}

export interface GetSubnetRouteTableAssociationOutputArgs {
    id: pulumi.Input<string>;
}
