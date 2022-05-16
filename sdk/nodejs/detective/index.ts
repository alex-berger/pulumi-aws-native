// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getGraph";
export * from "./getMemberInvitation";
export * from "./graph";
export * from "./memberInvitation";

// Import resources to register:
import { Graph } from "./graph";
import { MemberInvitation } from "./memberInvitation";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:detective:Graph":
                return new Graph(name, <any>undefined, { urn })
            case "aws-native:detective:MemberInvitation":
                return new MemberInvitation(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "detective", _module)