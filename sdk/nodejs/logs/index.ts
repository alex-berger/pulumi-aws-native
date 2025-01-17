// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./destination";
export * from "./getDestination";
export * from "./getLogGroup";
export * from "./getLogStream";
export * from "./getMetricFilter";
export * from "./getQueryDefinition";
export * from "./getResourcePolicy";
export * from "./getSubscriptionFilter";
export * from "./logGroup";
export * from "./logStream";
export * from "./metricFilter";
export * from "./queryDefinition";
export * from "./resourcePolicy";
export * from "./subscriptionFilter";

// Import resources to register:
import { Destination } from "./destination";
import { LogGroup } from "./logGroup";
import { LogStream } from "./logStream";
import { MetricFilter } from "./metricFilter";
import { QueryDefinition } from "./queryDefinition";
import { ResourcePolicy } from "./resourcePolicy";
import { SubscriptionFilter } from "./subscriptionFilter";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:logs:Destination":
                return new Destination(name, <any>undefined, { urn })
            case "aws-native:logs:LogGroup":
                return new LogGroup(name, <any>undefined, { urn })
            case "aws-native:logs:LogStream":
                return new LogStream(name, <any>undefined, { urn })
            case "aws-native:logs:MetricFilter":
                return new MetricFilter(name, <any>undefined, { urn })
            case "aws-native:logs:QueryDefinition":
                return new QueryDefinition(name, <any>undefined, { urn })
            case "aws-native:logs:ResourcePolicy":
                return new ResourcePolicy(name, <any>undefined, { urn })
            case "aws-native:logs:SubscriptionFilter":
                return new SubscriptionFilter(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "logs", _module)
