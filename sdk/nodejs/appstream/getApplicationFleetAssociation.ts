// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppStream::ApplicationFleetAssociation
 */
export function getApplicationFleetAssociation(args: GetApplicationFleetAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationFleetAssociationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:appstream:getApplicationFleetAssociation", {
        "applicationArn": args.applicationArn,
        "fleetName": args.fleetName,
    }, opts);
}

export interface GetApplicationFleetAssociationArgs {
    applicationArn: string;
    fleetName: string;
}

export interface GetApplicationFleetAssociationResult {
}

export function getApplicationFleetAssociationOutput(args: GetApplicationFleetAssociationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApplicationFleetAssociationResult> {
    return pulumi.output(args).apply(a => getApplicationFleetAssociation(a, opts))
}

export interface GetApplicationFleetAssociationOutputArgs {
    applicationArn: pulumi.Input<string>;
    fleetName: pulumi.Input<string>;
}