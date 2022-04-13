// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppStream::ApplicationEntitlementAssociation
 */
export function getApplicationEntitlementAssociation(args: GetApplicationEntitlementAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationEntitlementAssociationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:appstream:getApplicationEntitlementAssociation", {
        "applicationIdentifier": args.applicationIdentifier,
        "entitlementName": args.entitlementName,
        "stackName": args.stackName,
    }, opts);
}

export interface GetApplicationEntitlementAssociationArgs {
    applicationIdentifier: string;
    entitlementName: string;
    stackName: string;
}

export interface GetApplicationEntitlementAssociationResult {
}

export function getApplicationEntitlementAssociationOutput(args: GetApplicationEntitlementAssociationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApplicationEntitlementAssociationResult> {
    return pulumi.output(args).apply(a => getApplicationEntitlementAssociation(a, opts))
}

export interface GetApplicationEntitlementAssociationOutputArgs {
    applicationIdentifier: pulumi.Input<string>;
    entitlementName: pulumi.Input<string>;
    stackName: pulumi.Input<string>;
}
