// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ApiGateway::VpcLink
 */
export function getVpcLink(args: GetVpcLinkArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcLinkResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:apigateway:getVpcLink", {
        "id": args.id,
    }, opts);
}

export interface GetVpcLinkArgs {
    id: string;
}

export interface GetVpcLinkResult {
    readonly description?: string;
    readonly id?: string;
    readonly name?: string;
    readonly tags?: outputs.apigateway.VpcLinkTag[];
}

export function getVpcLinkOutput(args: GetVpcLinkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpcLinkResult> {
    return pulumi.output(args).apply(a => getVpcLink(a, opts))
}

export interface GetVpcLinkOutputArgs {
    id: pulumi.Input<string>;
}
