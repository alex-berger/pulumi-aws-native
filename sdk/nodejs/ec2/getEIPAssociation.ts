// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::EIPAssociation
 */
export function getEIPAssociation(args: GetEIPAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetEIPAssociationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ec2:getEIPAssociation", {
        "id": args.id,
    }, opts);
}

export interface GetEIPAssociationArgs {
    id: string;
}

export interface GetEIPAssociationResult {
    readonly allocationId?: string;
    readonly eIP?: string;
    readonly id?: string;
    readonly instanceId?: string;
    readonly networkInterfaceId?: string;
    readonly privateIpAddress?: string;
}

export function getEIPAssociationOutput(args: GetEIPAssociationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEIPAssociationResult> {
    return pulumi.output(args).apply(a => getEIPAssociation(a, opts))
}

export interface GetEIPAssociationOutputArgs {
    id: pulumi.Input<string>;
}
