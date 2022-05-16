// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Schema for the AWS Route53 Recovery Readiness ResourceSet Resource and API.
 */
export function getResourceSet(args: GetResourceSetArgs, opts?: pulumi.InvokeOptions): Promise<GetResourceSetResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:route53recoveryreadiness:getResourceSet", {
        "resourceSetName": args.resourceSetName,
    }, opts);
}

export interface GetResourceSetArgs {
    /**
     * The name of the resource set to create.
     */
    resourceSetName: string;
}

export interface GetResourceSetResult {
    /**
     * The Amazon Resource Name (ARN) of the resource set.
     */
    readonly resourceSetArn?: string;
    /**
     * A list of resource objects in the resource set.
     */
    readonly resources?: outputs.route53recoveryreadiness.ResourceSetResource[];
    /**
     * A tag to associate with the parameters for a resource set.
     */
    readonly tags?: outputs.route53recoveryreadiness.ResourceSetTag[];
}

export function getResourceSetOutput(args: GetResourceSetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetResourceSetResult> {
    return pulumi.output(args).apply(a => getResourceSet(a, opts))
}

export interface GetResourceSetOutputArgs {
    /**
     * The name of the resource set to create.
     */
    resourceSetName: pulumi.Input<string>;
}