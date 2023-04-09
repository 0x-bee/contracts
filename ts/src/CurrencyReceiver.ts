import BN from 'bn.js';
import BigNumber from 'bignumber.js';
import {
  PromiEvent,
  TransactionReceipt,
  EventResponse,
  EventData,
  Web3ContractContext,
} from 'ethereum-abi-types-generator';

export interface CallOptions {
  from?: string;
  gasPrice?: string;
  gas?: number;
}

export interface SendOptions {
  from: string;
  value?: number | string | BN | BigNumber;
  gasPrice?: string;
  gas?: number;
}

export interface EstimateGasOptions {
  from?: string;
  value?: number | string | BN | BigNumber;
  gas?: number;
}

export interface MethodPayableReturnContext {
  send(options: SendOptions): PromiEvent<TransactionReceipt>;
  send(
    options: SendOptions,
    callback: (error: Error, result: any) => void
  ): PromiEvent<TransactionReceipt>;
  estimateGas(options: EstimateGasOptions): Promise<number>;
  estimateGas(
    options: EstimateGasOptions,
    callback: (error: Error, result: any) => void
  ): Promise<number>;
  encodeABI(): string;
}

export interface MethodConstantReturnContext<TCallReturn> {
  call(): Promise<TCallReturn>;
  call(options: CallOptions): Promise<TCallReturn>;
  call(
    options: CallOptions,
    callback: (error: Error, result: TCallReturn) => void
  ): Promise<TCallReturn>;
  encodeABI(): string;
}

export interface MethodReturnContext extends MethodPayableReturnContext {}

export type ContractContext = Web3ContractContext<
  CurrencyReceiver,
  CurrencyReceiverMethodNames,
  CurrencyReceiverEventsContext,
  CurrencyReceiverEvents
>;
export type CurrencyReceiverEvents =
  | 'Pay'
  | 'Refund'
  | 'RoleAdminChanged'
  | 'RoleGranted'
  | 'RoleRevoked'
  | 'Withdraw';
export interface CurrencyReceiverEventsContext {
  Pay(
    parameters: {
      filter?: { currency?: string | string[]; from?: string | string[] };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  Refund(
    parameters: {
      filter?: { currency?: string | string[]; to?: string | string[] };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  RoleAdminChanged(
    parameters: {
      filter?: {
        role?: string | number[] | string | number[][];
        previousAdminRole?: string | number[] | string | number[][];
        newAdminRole?: string | number[] | string | number[][];
      };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  RoleGranted(
    parameters: {
      filter?: {
        role?: string | number[] | string | number[][];
        account?: string | string[];
        sender?: string | string[];
      };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  RoleRevoked(
    parameters: {
      filter?: {
        role?: string | number[] | string | number[][];
        account?: string | string[];
        sender?: string | string[];
      };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  Withdraw(
    parameters: {
      filter?: { currency?: string | string[]; to?: string | string[] };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
}
export type CurrencyReceiverMethodNames =
  | 'new'
  | 'DEFAULT_ADMIN_ROLE'
  | 'REFUND_ROLE'
  | 'WITHDRAW_ROLE'
  | 'balanceOf'
  | 'getRoleAdmin'
  | 'grantRole'
  | 'hasRole'
  | 'pay'
  | 'refund'
  | 'renounceRole'
  | 'revokeRole'
  | 'supportsInterface'
  | 'withdraw';
export interface PayEventEmittedResponse {
  currency: string;
  from: string;
  amount: string;
  orderId: string;
}
export interface RefundEventEmittedResponse {
  currency: string;
  amount: string;
  to: string;
  orderId: string;
}
export interface RoleAdminChangedEventEmittedResponse {
  role: string | number[];
  previousAdminRole: string | number[];
  newAdminRole: string | number[];
}
export interface RoleGrantedEventEmittedResponse {
  role: string | number[];
  account: string;
  sender: string;
}
export interface RoleRevokedEventEmittedResponse {
  role: string | number[];
  account: string;
  sender: string;
}
export interface WithdrawEventEmittedResponse {
  currency: string;
  amount: string;
  to: string;
  billId: string;
}
export interface CurrencyReceiver {
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: constructor
   */
  'new'(): MethodReturnContext;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  DEFAULT_ADMIN_ROLE(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  REFUND_ROLE(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  WITHDRAW_ROLE(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param currency Type: address, Indexed: false
   */
  balanceOf(currency: string): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param role Type: bytes32, Indexed: false
   */
  getRoleAdmin(role: string | number[]): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param role Type: bytes32, Indexed: false
   * @param account Type: address, Indexed: false
   */
  grantRole(role: string | number[], account: string): MethodReturnContext;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param role Type: bytes32, Indexed: false
   * @param account Type: address, Indexed: false
   */
  hasRole(
    role: string | number[],
    account: string
  ): MethodConstantReturnContext<boolean>;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param currency Type: address, Indexed: false
   * @param amount Type: uint256, Indexed: false
   * @param orderId Type: string, Indexed: false
   */
  pay(currency: string, amount: string, orderId: string): MethodReturnContext;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param currency Type: address[], Indexed: false
   * @param amount Type: uint256[], Indexed: false
   * @param to Type: address[], Indexed: false
   * @param orderId Type: string[], Indexed: false
   */
  refund(
    currency: string[],
    amount: string[],
    to: string[],
    orderId: string[]
  ): MethodReturnContext;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param role Type: bytes32, Indexed: false
   * @param account Type: address, Indexed: false
   */
  renounceRole(role: string | number[], account: string): MethodReturnContext;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param role Type: bytes32, Indexed: false
   * @param account Type: address, Indexed: false
   */
  revokeRole(role: string | number[], account: string): MethodReturnContext;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param interfaceId Type: bytes4, Indexed: false
   */
  supportsInterface(
    interfaceId: string | number[]
  ): MethodConstantReturnContext<boolean>;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param currency Type: address[], Indexed: false
   * @param amount Type: uint256[], Indexed: false
   * @param to Type: address[], Indexed: false
   * @param billId Type: string[], Indexed: false
   */
  withdraw(
    currency: string[],
    amount: string[],
    to: string[],
    billId: string[]
  ): MethodReturnContext;
}
