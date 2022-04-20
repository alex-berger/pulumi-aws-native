// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * An example resource schema demonstrating some basic constructs and validation rules.
 */
export function getCertificate(args: GetCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetCertificateResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:lightsail:getCertificate", {
        "certificateName": args.certificateName,
    }, opts);
}

export interface GetCertificateArgs {
    /**
     * The name for the certificate.
     */
    certificateName: string;
}

export interface GetCertificateResult {
    readonly certificateArn?: string;
    /**
     * The validation status of the certificate.
     */
    readonly status?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.lightsail.CertificateTag[];
}

export function getCertificateOutput(args: GetCertificateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCertificateResult> {
    return pulumi.output(args).apply(a => getCertificate(a, opts))
}

export interface GetCertificateOutputArgs {
    /**
     * The name for the certificate.
     */
    certificateName: pulumi.Input<string>;
}
