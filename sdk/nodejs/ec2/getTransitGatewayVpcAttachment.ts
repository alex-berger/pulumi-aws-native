// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::TransitGatewayVpcAttachment
 */
export function getTransitGatewayVpcAttachment(args: GetTransitGatewayVpcAttachmentArgs, opts?: pulumi.InvokeOptions): Promise<GetTransitGatewayVpcAttachmentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ec2:getTransitGatewayVpcAttachment", {
        "id": args.id,
    }, opts);
}

export interface GetTransitGatewayVpcAttachmentArgs {
    id: string;
}

export interface GetTransitGatewayVpcAttachmentResult {
    readonly addSubnetIds?: string[];
    readonly id?: string;
    /**
     * The options for the transit gateway vpc attachment.
     */
    readonly options?: outputs.ec2.OptionsProperties;
    readonly removeSubnetIds?: string[];
    readonly tags?: outputs.ec2.TransitGatewayVpcAttachmentTag[];
}

export function getTransitGatewayVpcAttachmentOutput(args: GetTransitGatewayVpcAttachmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTransitGatewayVpcAttachmentResult> {
    return pulumi.output(args).apply(a => getTransitGatewayVpcAttachment(a, opts))
}

export interface GetTransitGatewayVpcAttachmentOutputArgs {
    id: pulumi.Input<string>;
}
