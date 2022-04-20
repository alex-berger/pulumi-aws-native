// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppSync::DomainNameApiAssociation
 */
export function getDomainNameApiAssociation(args: GetDomainNameApiAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainNameApiAssociationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:appsync:getDomainNameApiAssociation", {
        "apiAssociationIdentifier": args.apiAssociationIdentifier,
    }, opts);
}

export interface GetDomainNameApiAssociationArgs {
    apiAssociationIdentifier: string;
}

export interface GetDomainNameApiAssociationResult {
    readonly apiAssociationIdentifier?: string;
    readonly apiId?: string;
}

export function getDomainNameApiAssociationOutput(args: GetDomainNameApiAssociationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDomainNameApiAssociationResult> {
    return pulumi.output(args).apply(a => getDomainNameApiAssociation(a, opts))
}

export interface GetDomainNameApiAssociationOutputArgs {
    apiAssociationIdentifier: pulumi.Input<string>;
}
