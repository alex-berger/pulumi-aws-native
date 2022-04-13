// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppStream::ImageBuilder
 */
export function getImageBuilder(args: GetImageBuilderArgs, opts?: pulumi.InvokeOptions): Promise<GetImageBuilderResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:appstream:getImageBuilder", {
        "id": args.id,
    }, opts);
}

export interface GetImageBuilderArgs {
    id: string;
}

export interface GetImageBuilderResult {
    readonly accessEndpoints?: outputs.appstream.ImageBuilderAccessEndpoint[];
    readonly appstreamAgentVersion?: string;
    readonly description?: string;
    readonly displayName?: string;
    readonly domainJoinInfo?: outputs.appstream.ImageBuilderDomainJoinInfo;
    readonly enableDefaultInternetAccess?: boolean;
    readonly iamRoleArn?: string;
    readonly id?: string;
    readonly imageArn?: string;
    readonly imageName?: string;
    readonly instanceType?: string;
    readonly name?: string;
    readonly streamingUrl?: string;
    readonly tags?: outputs.appstream.ImageBuilderTag[];
    readonly vpcConfig?: outputs.appstream.ImageBuilderVpcConfig;
}

export function getImageBuilderOutput(args: GetImageBuilderOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetImageBuilderResult> {
    return pulumi.output(args).apply(a => getImageBuilder(a, opts))
}

export interface GetImageBuilderOutputArgs {
    id: pulumi.Input<string>;
}
