// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppStream::StackUserAssociation
 */
export function getStackUserAssociation(args: GetStackUserAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetStackUserAssociationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:appstream:getStackUserAssociation", {
        "id": args.id,
    }, opts);
}

export interface GetStackUserAssociationArgs {
    id: string;
}

export interface GetStackUserAssociationResult {
    readonly id?: string;
}

export function getStackUserAssociationOutput(args: GetStackUserAssociationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStackUserAssociationResult> {
    return pulumi.output(args).apply(a => getStackUserAssociation(a, opts))
}

export interface GetStackUserAssociationOutputArgs {
    id: pulumi.Input<string>;
}
