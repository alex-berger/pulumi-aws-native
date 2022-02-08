// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ApiGateway::ClientCertificate
 */
export function getClientCertificate(args: GetClientCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetClientCertificateResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:apigateway:getClientCertificate", {
        "clientCertificateId": args.clientCertificateId,
    }, opts);
}

export interface GetClientCertificateArgs {
    /**
     * The Primary Identifier of the Client Certficate, generated by a Create API Call
     */
    clientCertificateId: string;
}

export interface GetClientCertificateResult {
    /**
     * The Primary Identifier of the Client Certficate, generated by a Create API Call
     */
    readonly clientCertificateId?: string;
    /**
     * A description of the client certificate.
     */
    readonly description?: string;
    /**
     * An array of arbitrary tags (key-value pairs) to associate with the client certificate.
     */
    readonly tags?: outputs.apigateway.ClientCertificateTag[];
}

export function getClientCertificateOutput(args: GetClientCertificateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClientCertificateResult> {
    return pulumi.output(args).apply(a => getClientCertificate(a, opts))
}

export interface GetClientCertificateOutputArgs {
    /**
     * The Primary Identifier of the Client Certficate, generated by a Create API Call
     */
    clientCertificateId: pulumi.Input<string>;
}