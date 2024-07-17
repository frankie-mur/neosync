// @generated by protoc-gen-connect-query v1.4.1 with parameter "target=ts,import_extension=.js"
// @generated from file mgmt/v1alpha1/api_key.proto (package mgmt.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { MethodKind } from "@bufbuild/protobuf";
import { CreateAccountApiKeyRequest, CreateAccountApiKeyResponse, DeleteAccountApiKeyRequest, DeleteAccountApiKeyResponse, GetAccountApiKeyRequest, GetAccountApiKeyResponse, GetAccountApiKeysRequest, GetAccountApiKeysResponse, RegenerateAccountApiKeyRequest, RegenerateAccountApiKeyResponse } from "./api_key_pb.js";

/**
 * Retrieves a list of Account API Keys
 *
 * @generated from rpc mgmt.v1alpha1.ApiKeyService.GetAccountApiKeys
 */
export const getAccountApiKeys = {
  localName: "getAccountApiKeys",
  name: "GetAccountApiKeys",
  kind: MethodKind.Unary,
  I: GetAccountApiKeysRequest,
  O: GetAccountApiKeysResponse,
  service: {
    typeName: "mgmt.v1alpha1.ApiKeyService"
  }
} as const;

/**
 * Retrieves a single API Key
 *
 * @generated from rpc mgmt.v1alpha1.ApiKeyService.GetAccountApiKey
 */
export const getAccountApiKey = {
  localName: "getAccountApiKey",
  name: "GetAccountApiKey",
  kind: MethodKind.Unary,
  I: GetAccountApiKeyRequest,
  O: GetAccountApiKeyResponse,
  service: {
    typeName: "mgmt.v1alpha1.ApiKeyService"
  }
} as const;

/**
 * Creates a single API Key
 * This method will return the decrypted contents of the API key
 *
 * @generated from rpc mgmt.v1alpha1.ApiKeyService.CreateAccountApiKey
 */
export const createAccountApiKey = {
  localName: "createAccountApiKey",
  name: "CreateAccountApiKey",
  kind: MethodKind.Unary,
  I: CreateAccountApiKeyRequest,
  O: CreateAccountApiKeyResponse,
  service: {
    typeName: "mgmt.v1alpha1.ApiKeyService"
  }
} as const;

/**
 * Regenerates a single API Key with a new expiration time
 * This method will return the decrypted contents of the API key
 *
 * @generated from rpc mgmt.v1alpha1.ApiKeyService.RegenerateAccountApiKey
 */
export const regenerateAccountApiKey = {
  localName: "regenerateAccountApiKey",
  name: "RegenerateAccountApiKey",
  kind: MethodKind.Unary,
  I: RegenerateAccountApiKeyRequest,
  O: RegenerateAccountApiKeyResponse,
  service: {
    typeName: "mgmt.v1alpha1.ApiKeyService"
  }
} as const;

/**
 * Deletes an API Key from the system.
 *
 * @generated from rpc mgmt.v1alpha1.ApiKeyService.DeleteAccountApiKey
 */
export const deleteAccountApiKey = {
  localName: "deleteAccountApiKey",
  name: "DeleteAccountApiKey",
  kind: MethodKind.Unary,
  I: DeleteAccountApiKeyRequest,
  O: DeleteAccountApiKeyResponse,
  service: {
    typeName: "mgmt.v1alpha1.ApiKeyService"
  }
} as const;
