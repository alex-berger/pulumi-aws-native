## HEAD (Unreleased)
(None)

## 0.16.1 (Apr 29, 2022)
- Fix getAzs [#472](https://github.com/pulumi/pulumi-aws-native/pull/472)

## 0.16.0 (Apr 25, 2022)
- Implement support for getUrlSuffix.
  [#456](https://github.com/pulumi/pulumi-aws-native/pull/456)
- Ensure that all create-only properties are captured as input properties.
  [#466](https://github.com/pulumi/pulumi-aws-native/pull/466)
- Upgrade to Pulumi/Pulumi v3.30.0
  [#458](https://github.com/pulumi/pulumi-aws-native/pull/458)

---

## 0.15.0 (Apr 06, 2022)

- Update pulumi codegen dependency to fix secret property handling in Go SDK
- Update to include the latest resource definitions

## 0.14.0 (Mar 25, 2022)

- Update pulumi codegen dependency to fix secret property handling in Go SDK
- Update to include the latest resource definitions

## 0.13.0 (Feb 23, 2022)

- Update to include the latest resource definitions

## 0.12.0 (Feb 8, 2022)

- Implement get functions for all resources
  [#346](https://github.com/pulumi/pulumi-aws-native/pull/346)
- Update Pulumi dependencies to v3.23.0
- Use sequence numbers to generate deterministic autonames.
  [#341](https://github.com/pulumi/pulumi-aws-native/pull/341)
- Update to include the latest resource definitions

## 0.11.0 (Jan 24, 2022)

- Update to include the latest resource definitions

## 0.10.0 (Jan 10, 2022)

- Update to include the latest resource definitions
- Serve decompressed schema from GetSchema API
  [#287](https://github.com/pulumi/pulumi-aws-native/pull/287)

## 0.9.0 (December 15, 2021)

- Update to include the latest resource definitions
- Explain how to add missing region config
  [#220](https://github.com/pulumi/pulumi-aws-native/issues/220)
- Allow cancellation during ongoing operations
  [#43](https://github.com/pulumi/pulumi-aws-native/issues/43)

## 0.8.0 (November 30, 2021)

Update to include the latest resource definitions

## 0.7.1 (November 23 2021)

- Fix for "Custom providers can leak credentials to state file from environment variables"
  [#236](https://github.com/pulumi/pulumi-aws-native/issues/236)

  **PLEASE READ**

  If you set credentials through environment variables (e.g. `AWS_SECRET_ACCESS_KEY`) AND
  use the SDK to create a provider where these values are not explicitly set, (e.g. `new awsnative.Provider("...");`)
  prior versions of the `aws-native` provider may have included the credentials in the state in clear text.

  All users are recommended to upgrade their provider version to this or newer version and run a `pulumi up`.
  Please also rotate the affected credentials after all relevant stacks have been updated.

  You can check if your state file contains credentials by running `pulumi stack export | grep -A 3 "accessKey\|secretKey\|token"`
  and checking if any unencrypted values are produced. After the update these values will either not be present
  or be stored as encrypted secrets using your stack's preferred encryption provider.

  Note that the Pulumi state backend also encrypts the state as a whole and other state backends
  support a similar mechanism which should significantly limit exposure of the creds.
  Nonetheless, We sincerely regret the inconvenience this causes.

## 0.7.0 (November 22, 2021)

Update to include the latest resource definitions

## 0.6.0 (November 18, 2021)

- Add support for specifying max retries
  [#112](https://github.com/pulumi/pulumi-aws-native/issues/112)
- Fix modifying S3 bucket tags
  [#230](https://github.com/pulumi/pulumi-aws-native/issues/230)

## 0.5.0 (November 15, 2021)

- Improve retry logic to avoid excessive requests to APIs
  [#192](https://github.com/pulumi/pulumi-aws-native/issues/192)

## 0.4.0 (November 8, 2021)

- Update to pulumi v3.16.0 deps
  [#202](https://github.com/pulumi/pulumi-aws-native/pull/202)
- Add autonaming support
  [#156](https://github.com/pulumi/pulumi-aws-native/issues/156)
- Turn off replacement detection to prevent false replacements
  [#186](https://github.com/pulumi/pulumi-aws-native/issues/186)
- Support partial initialization errors
  [#208](https://github.com/pulumi/pulumi-aws-native/pull/208)

## 0.3.0 (November 2, 2021)

- Latest resource definitions

## 0.2.0 (October 8, 2021)

- Deduplicate type names [#160](https://github.com/pulumi/pulumi-aws-native/issues/160)
- [cf2pulumi] Fix supported check for config module [#165](https://github.com/pulumi/pulumi-aws-native/issues/165)
- Improve inline schema type representation [#167](https://github.com/pulumi/pulumi-aws-native/pull/167)

## 0.1.0 (September 30, 2021)

First public release
