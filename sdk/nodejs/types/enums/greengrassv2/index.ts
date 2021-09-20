// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ComponentVersionLambdaEventSourceType = {
    PubSub: "PUB_SUB",
    IotCore: "IOT_CORE",
} as const;

export type ComponentVersionLambdaEventSourceType = (typeof ComponentVersionLambdaEventSourceType)[keyof typeof ComponentVersionLambdaEventSourceType];

export const ComponentVersionLambdaExecutionParametersInputPayloadEncodingType = {
    Json: "json",
    Binary: "binary",
} as const;

export type ComponentVersionLambdaExecutionParametersInputPayloadEncodingType = (typeof ComponentVersionLambdaExecutionParametersInputPayloadEncodingType)[keyof typeof ComponentVersionLambdaExecutionParametersInputPayloadEncodingType];

export const ComponentVersionLambdaFilesystemPermission = {
    Ro: "ro",
    Rw: "rw",
} as const;

export type ComponentVersionLambdaFilesystemPermission = (typeof ComponentVersionLambdaFilesystemPermission)[keyof typeof ComponentVersionLambdaFilesystemPermission];

export const ComponentVersionLambdaLinuxProcessParamsIsolationMode = {
    GreengrassContainer: "GreengrassContainer",
    NoContainer: "NoContainer",
} as const;

export type ComponentVersionLambdaLinuxProcessParamsIsolationMode = (typeof ComponentVersionLambdaLinuxProcessParamsIsolationMode)[keyof typeof ComponentVersionLambdaLinuxProcessParamsIsolationMode];