// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Schema of AWS::EC2::IPAMAllocation Type
 */
export function getIPAMAllocation(args: GetIPAMAllocationArgs, opts?: pulumi.InvokeOptions): Promise<GetIPAMAllocationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ec2:getIPAMAllocation", {
        "cidr": args.cidr,
        "ipamPoolAllocationId": args.ipamPoolAllocationId,
        "ipamPoolId": args.ipamPoolId,
    }, opts);
}

export interface GetIPAMAllocationArgs {
    cidr: string;
    /**
     * Id of the allocation.
     */
    ipamPoolAllocationId: string;
    /**
     * Id of the IPAM Pool.
     */
    ipamPoolId: string;
}

export interface GetIPAMAllocationResult {
    /**
     * Id of the allocation.
     */
    readonly ipamPoolAllocationId?: string;
}

export function getIPAMAllocationOutput(args: GetIPAMAllocationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIPAMAllocationResult> {
    return pulumi.output(args).apply(a => getIPAMAllocation(a, opts))
}

export interface GetIPAMAllocationOutputArgs {
    cidr: pulumi.Input<string>;
    /**
     * Id of the allocation.
     */
    ipamPoolAllocationId: pulumi.Input<string>;
    /**
     * Id of the IPAM Pool.
     */
    ipamPoolId: pulumi.Input<string>;
}
