// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getStream";
export * from "./getStreamConsumer";
export * from "./stream";
export * from "./streamConsumer";

// Export enums:
export * from "../types/enums/kinesis";

// Import resources to register:
import { Stream } from "./stream";
import { StreamConsumer } from "./streamConsumer";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:kinesis:Stream":
                return new Stream(name, <any>undefined, { urn })
            case "aws-native:kinesis:StreamConsumer":
                return new StreamConsumer(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "kinesis", _module)