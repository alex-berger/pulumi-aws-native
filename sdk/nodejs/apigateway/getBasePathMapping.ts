// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ApiGateway::BasePathMapping
 */
export function getBasePathMapping(args: GetBasePathMappingArgs, opts?: pulumi.InvokeOptions): Promise<GetBasePathMappingResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:apigateway:getBasePathMapping", {
        "basePath": args.basePath,
        "domainName": args.domainName,
    }, opts);
}

export interface GetBasePathMappingArgs {
    /**
     * The base path name that callers of the API must provide in the URL after the domain name.
     */
    basePath: string;
    /**
     * The DomainName of an AWS::ApiGateway::DomainName resource.
     */
    domainName: string;
}

export interface GetBasePathMappingResult {
    readonly id?: string;
    /**
     * The ID of the API.
     */
    readonly restApiId?: string;
    /**
     * The name of the API's stage.
     */
    readonly stage?: string;
}

export function getBasePathMappingOutput(args: GetBasePathMappingOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBasePathMappingResult> {
    return pulumi.output(args).apply(a => getBasePathMapping(a, opts))
}

export interface GetBasePathMappingOutputArgs {
    /**
     * The base path name that callers of the API must provide in the URL after the domain name.
     */
    basePath: pulumi.Input<string>;
    /**
     * The DomainName of an AWS::ApiGateway::DomainName resource.
     */
    domainName: pulumi.Input<string>;
}
