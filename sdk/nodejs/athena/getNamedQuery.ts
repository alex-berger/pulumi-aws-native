// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::Athena::NamedQuery
 */
export function getNamedQuery(args: GetNamedQueryArgs, opts?: pulumi.InvokeOptions): Promise<GetNamedQueryResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:athena:getNamedQuery", {
        "namedQueryId": args.namedQueryId,
    }, opts);
}

export interface GetNamedQueryArgs {
    /**
     * The unique ID of the query.
     */
    namedQueryId: string;
}

export interface GetNamedQueryResult {
    /**
     * The unique ID of the query.
     */
    readonly namedQueryId?: string;
}

export function getNamedQueryOutput(args: GetNamedQueryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNamedQueryResult> {
    return pulumi.output(args).apply(a => getNamedQuery(a, opts))
}

export interface GetNamedQueryOutputArgs {
    /**
     * The unique ID of the query.
     */
    namedQueryId: pulumi.Input<string>;
}
