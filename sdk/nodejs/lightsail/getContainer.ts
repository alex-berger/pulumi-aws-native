// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Lightsail::Container
 */
export function getContainer(args: GetContainerArgs, opts?: pulumi.InvokeOptions): Promise<GetContainerResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:lightsail:getContainer", {
        "serviceName": args.serviceName,
    }, opts);
}

export interface GetContainerArgs {
    /**
     * The name for the container service.
     */
    serviceName: string;
}

export interface GetContainerResult {
    readonly containerArn?: string;
    /**
     * Describes a container deployment configuration of an Amazon Lightsail container service.
     */
    readonly containerServiceDeployment?: outputs.lightsail.ContainerServiceDeployment;
    /**
     * A Boolean value to indicate whether the container service is disabled.
     */
    readonly isDisabled?: boolean;
    /**
     * The power specification for the container service.
     */
    readonly power?: string;
    /**
     * The public domain names to use with the container service, such as example.com and www.example.com.
     */
    readonly publicDomainNames?: outputs.lightsail.ContainerPublicDomainName[];
    /**
     * The scale specification for the container service.
     */
    readonly scale?: number;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.lightsail.ContainerTag[];
    /**
     * The publicly accessible URL of the container service.
     */
    readonly url?: string;
}

export function getContainerOutput(args: GetContainerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetContainerResult> {
    return pulumi.output(args).apply(a => getContainer(a, opts))
}

export interface GetContainerOutputArgs {
    /**
     * The name for the container service.
     */
    serviceName: pulumi.Input<string>;
}