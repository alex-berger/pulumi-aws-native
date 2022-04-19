// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Lightsail::LoadBalancerTlsCertificate
 */
export function getLoadBalancerTlsCertificate(args: GetLoadBalancerTlsCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetLoadBalancerTlsCertificateResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:lightsail:getLoadBalancerTlsCertificate", {
        "certificateName": args.certificateName,
        "loadBalancerName": args.loadBalancerName,
    }, opts);
}

export interface GetLoadBalancerTlsCertificateArgs {
    /**
     * The SSL/TLS certificate name.
     */
    certificateName: string;
    /**
     * The name of your load balancer.
     */
    loadBalancerName: string;
}

export interface GetLoadBalancerTlsCertificateResult {
    /**
     * When true, the SSL/TLS certificate is attached to the Lightsail load balancer.
     */
    readonly isAttached?: boolean;
    readonly loadBalancerTlsCertificateArn?: string;
    /**
     * The validation status of the SSL/TLS certificate.
     */
    readonly status?: string;
}

export function getLoadBalancerTlsCertificateOutput(args: GetLoadBalancerTlsCertificateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLoadBalancerTlsCertificateResult> {
    return pulumi.output(args).apply(a => getLoadBalancerTlsCertificate(a, opts))
}

export interface GetLoadBalancerTlsCertificateOutputArgs {
    /**
     * The SSL/TLS certificate name.
     */
    certificateName: pulumi.Input<string>;
    /**
     * The name of your load balancer.
     */
    loadBalancerName: pulumi.Input<string>;
}
