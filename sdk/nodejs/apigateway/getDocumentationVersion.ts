// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A snapshot of the documentation of an API.
 */
export function getDocumentationVersion(args: GetDocumentationVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetDocumentationVersionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:apigateway:getDocumentationVersion", {
        "documentationVersion": args.documentationVersion,
        "restApiId": args.restApiId,
    }, opts);
}

export interface GetDocumentationVersionArgs {
    /**
     * The version identifier of the API documentation snapshot.
     */
    documentationVersion: string;
    /**
     * The identifier of the API.
     */
    restApiId: string;
}

export interface GetDocumentationVersionResult {
    /**
     * The description of the API documentation snapshot.
     */
    readonly description?: string;
}

export function getDocumentationVersionOutput(args: GetDocumentationVersionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDocumentationVersionResult> {
    return pulumi.output(args).apply(a => getDocumentationVersion(a, opts))
}

export interface GetDocumentationVersionOutputArgs {
    /**
     * The version identifier of the API documentation snapshot.
     */
    documentationVersion: pulumi.Input<string>;
    /**
     * The identifier of the API.
     */
    restApiId: pulumi.Input<string>;
}
