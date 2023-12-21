// @generated by protoc-gen-connect-es v1.0.0 with parameter "target=ts,import_extension=.js"
// @generated from file mgmt/v1alpha1/user_account.proto (package mgmt.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AcceptTeamAccountInviteRequest, AcceptTeamAccountInviteResponse, ConvertPersonalToTeamAccountRequest, ConvertPersonalToTeamAccountResponse, CreateTeamAccountRequest, CreateTeamAccountResponse, GetAccountTemporalConfigRequest, GetAccountTemporalConfigResponse, GetSystemInformationRequest, GetSystemInformationResponse, GetTeamAccountInvitesRequest, GetTeamAccountInvitesResponse, GetTeamAccountMembersRequest, GetTeamAccountMembersResponse, GetUserAccountsRequest, GetUserAccountsResponse, GetUserRequest, GetUserResponse, InviteUserToTeamAccountRequest, InviteUserToTeamAccountResponse, IsUserInAccountRequest, IsUserInAccountResponse, RemoveTeamAccountInviteRequest, RemoveTeamAccountInviteResponse, RemoveTeamAccountMemberRequest, RemoveTeamAccountMemberResponse, SetAccountTemporalConfigRequest, SetAccountTemporalConfigResponse, SetPersonalAccountRequest, SetPersonalAccountResponse, SetUserRequest, SetUserResponse } from "./user_account_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service mgmt.v1alpha1.UserAccountService
 */
export const UserAccountService = {
  typeName: "mgmt.v1alpha1.UserAccountService",
  methods: {
    /**
     * @generated from rpc mgmt.v1alpha1.UserAccountService.GetUser
     */
    getUser: {
      name: "GetUser",
      I: GetUserRequest,
      O: GetUserResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.UserAccountService.SetUser
     */
    setUser: {
      name: "SetUser",
      I: SetUserRequest,
      O: SetUserResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.UserAccountService.GetUserAccounts
     */
    getUserAccounts: {
      name: "GetUserAccounts",
      I: GetUserAccountsRequest,
      O: GetUserAccountsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.UserAccountService.SetPersonalAccount
     */
    setPersonalAccount: {
      name: "SetPersonalAccount",
      I: SetPersonalAccountRequest,
      O: SetPersonalAccountResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.UserAccountService.ConvertPersonalToTeamAccount
     */
    convertPersonalToTeamAccount: {
      name: "ConvertPersonalToTeamAccount",
      I: ConvertPersonalToTeamAccountRequest,
      O: ConvertPersonalToTeamAccountResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.UserAccountService.CreateTeamAccount
     */
    createTeamAccount: {
      name: "CreateTeamAccount",
      I: CreateTeamAccountRequest,
      O: CreateTeamAccountResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.UserAccountService.IsUserInAccount
     */
    isUserInAccount: {
      name: "IsUserInAccount",
      I: IsUserInAccountRequest,
      O: IsUserInAccountResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.UserAccountService.GetAccountTemporalConfig
     */
    getAccountTemporalConfig: {
      name: "GetAccountTemporalConfig",
      I: GetAccountTemporalConfigRequest,
      O: GetAccountTemporalConfigResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.UserAccountService.SetAccountTemporalConfig
     */
    setAccountTemporalConfig: {
      name: "SetAccountTemporalConfig",
      I: SetAccountTemporalConfigRequest,
      O: SetAccountTemporalConfigResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.UserAccountService.GetTeamAccountMembers
     */
    getTeamAccountMembers: {
      name: "GetTeamAccountMembers",
      I: GetTeamAccountMembersRequest,
      O: GetTeamAccountMembersResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.UserAccountService.RemoveTeamAccountMember
     */
    removeTeamAccountMember: {
      name: "RemoveTeamAccountMember",
      I: RemoveTeamAccountMemberRequest,
      O: RemoveTeamAccountMemberResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.UserAccountService.InviteUserToTeamAccount
     */
    inviteUserToTeamAccount: {
      name: "InviteUserToTeamAccount",
      I: InviteUserToTeamAccountRequest,
      O: InviteUserToTeamAccountResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.UserAccountService.GetTeamAccountInvites
     */
    getTeamAccountInvites: {
      name: "GetTeamAccountInvites",
      I: GetTeamAccountInvitesRequest,
      O: GetTeamAccountInvitesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.UserAccountService.RemoveTeamAccountInvite
     */
    removeTeamAccountInvite: {
      name: "RemoveTeamAccountInvite",
      I: RemoveTeamAccountInviteRequest,
      O: RemoveTeamAccountInviteResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.UserAccountService.AcceptTeamAccountInvite
     */
    acceptTeamAccountInvite: {
      name: "AcceptTeamAccountInvite",
      I: AcceptTeamAccountInviteRequest,
      O: AcceptTeamAccountInviteResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.UserAccountService.GetSystemInformation
     */
    getSystemInformation: {
      name: "GetSystemInformation",
      I: GetSystemInformationRequest,
      O: GetSystemInformationResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

