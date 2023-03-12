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
  BeeWharf,
  BeeWharfMethodNames,
  BeeWharfEventsContext,
  BeeWharfEvents
>;
export type BeeWharfEvents =
  | 'OwnershipTransferred'
  | 'PayERC20'
  | 'PayEth'
  | 'RefundERC20'
  | 'RefundEth'
  | 'WithdrawERC20'
  | 'WithdrawEth';
export interface BeeWharfEventsContext {
  OwnershipTransferred(
    parameters: {
      filter?: {
        previousOwner?: string | string[];
        newOwner?: string | string[];
      };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  PayERC20(
    parameters: {
      filter?: { from?: string | string[]; tokenAddress?: string | string[] };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  PayEth(
    parameters: {
      filter?: { from?: string | string[] };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  RefundERC20(
    parameters: {
      filter?: { to?: string | string[]; tokenAddress?: string | string[] };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  RefundEth(
    parameters: {
      filter?: { to?: string | string[] };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  WithdrawERC20(
    parameters: {
      filter?: { to?: string | string[]; tokenAddress?: string | string[] };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  WithdrawEth(
    parameters: {
      filter?: { to?: string | string[] };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
}
export type BeeWharfMethodNames =
  | 'new'
  | 'addNewSupportToken'
  | 'balanceOfERC20'
  | 'balanceOfETH'
  | 'owner'
  | 'payERC20'
  | 'payERC20From'
  | 'payEth'
  | 'refundERC20'
  | 'refundERC20From'
  | 'refundETH'
  | 'renounceOwnership'
  | 'transferOwnership'
  | 'withdrawERC20'
  | 'withdrawEth';
export interface OwnershipTransferredEventEmittedResponse {
  previousOwner: string;
  newOwner: string;
}
export interface PayERC20EventEmittedResponse {
  from: string;
  amount: string;
  orderId: string;
  tokenAddress: string;
}
export interface PayEthEventEmittedResponse {
  from: string;
  amount: string;
  orderId: string;
}
export interface RefundERC20EventEmittedResponse {
  to: string;
  amount: string;
  billId: string;
  tokenAddress: string;
}
export interface RefundEthEventEmittedResponse {
  to: string;
  amount: string;
  billId: string;
}
export interface WithdrawERC20EventEmittedResponse {
  to: string;
  amount: string;
  billId: string;
  tokenAddress: string;
}
export interface WithdrawEthEventEmittedResponse {
  to: string;
  amount: string;
  billId: string;
}
export interface BeeWharf {
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: constructor
   */
  'new'(): MethodReturnContext;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param tokenAddress Type: address, Indexed: false
   */
  addNewSupportToken(tokenAddress: string): MethodReturnContext;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param tokenAddress Type: address, Indexed: false
   */
  balanceOfERC20(tokenAddress: string): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  balanceOfETH(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  owner(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param amount Type: uint256, Indexed: false
   * @param orderId Type: string, Indexed: false
   * @param tokenAddress Type: address, Indexed: false
   */
  payERC20(
    amount: string,
    orderId: string,
    tokenAddress: string
  ): MethodReturnContext;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param from Type: address, Indexed: false
   * @param amount Type: uint256, Indexed: false
   * @param orderId Type: string, Indexed: false
   * @param tokenAddress Type: address, Indexed: false
   */
  payERC20From(
    from: string,
    amount: string,
    orderId: string,
    tokenAddress: string
  ): MethodReturnContext;
  /**
   * Payable: true
   * Constant: false
   * StateMutability: payable
   * Type: function
   * @param orderId Type: string, Indexed: false
   */
  payEth(orderId: string): MethodPayableReturnContext;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param to Type: address, Indexed: false
   * @param amount Type: uint256, Indexed: false
   * @param billId Type: string, Indexed: false
   * @param tokenAddress Type: address, Indexed: false
   */
  refundERC20(
    to: string,
    amount: string,
    billId: string,
    tokenAddress: string
  ): MethodReturnContext;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param from Type: address, Indexed: false
   * @param to Type: address, Indexed: false
   * @param amount Type: uint256, Indexed: false
   * @param billId Type: string, Indexed: false
   * @param tokenAddress Type: address, Indexed: false
   */
  refundERC20From(
    from: string,
    to: string,
    amount: string,
    billId: string,
    tokenAddress: string
  ): MethodReturnContext;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param to Type: address, Indexed: false
   * @param amount Type: uint256, Indexed: false
   * @param billId Type: string, Indexed: false
   */
  refundETH(to: string, amount: string, billId: string): MethodReturnContext;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   */
  renounceOwnership(): MethodReturnContext;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param newOwner Type: address, Indexed: false
   */
  transferOwnership(newOwner: string): MethodReturnContext;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param amount Type: uint256, Indexed: false
   * @param billId Type: string, Indexed: false
   * @param tokenAddress Type: address, Indexed: false
   */
  withdrawERC20(
    amount: string,
    billId: string,
    tokenAddress: string
  ): MethodReturnContext;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param amount Type: uint256, Indexed: false
   * @param billId Type: string, Indexed: false
   */
  withdrawEth(amount: string, billId: string): MethodReturnContext;
}
