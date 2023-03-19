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
  | 'OwnershipTransferred'
  | 'Pay'
  | 'Refund'
  | 'Withdraw';
export interface CurrencyReceiverEventsContext {
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
      filter?: { currency?: string[] | string[][]; to?: string[] | string[][] };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  Withdraw(
    parameters: {
      filter?: { currency?: string[] | string[][]; to?: string | string[] };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
}
export type CurrencyReceiverMethodNames =
  | 'new'
  | 'balanceOf'
  | 'owner'
  | 'pay'
  | 'refund'
  | 'renounceOwnership'
  | 'transferOwnership'
  | 'withdraw';
export interface OwnershipTransferredEventEmittedResponse {
  previousOwner: string;
  newOwner: string;
}
export interface PayEventEmittedResponse {
  currency: string;
  from: string;
  amount: string;
  orderId: string;
}
export interface RefundEventEmittedResponse {
  currency: string[];
  amount: string[];
  to: string[];
  orderIds: string[];
}
export interface WithdrawEventEmittedResponse {
  currency: string[];
  amount: string[];
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
   * @param currency Type: address, Indexed: false
   */
  balanceOf(currency: string): MethodConstantReturnContext<string>;
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
   * @param currency Type: address[], Indexed: false
   * @param amount Type: uint256[], Indexed: false
   * @param billId Type: string, Indexed: false
   */
  withdraw(
    currency: string[],
    amount: string[],
    billId: string
  ): MethodReturnContext;
}
