# Probe 
![image](https://github.com/DefiantLabs/probe/assets/807940/1ac446be-4ba7-4ee2-b317-856930b899fb)


*Your lightweight Cosmos SDK Querier*

**Delve into the depths of the Cosmos with Probe**

---

## :question: What is Probe?

Probe is a streamlined Cosmos SDK querying package. It deftly manages the unpacking process for base Cosmos SDK types and provides a detailed process for handling specialized modules.

:sparkles: Key features:

- Direct import of Cosmos SDK and setting up the Codec
- Comprehensive support for customized modules
- Lightweight design focusing on query operations

---

## :point_right: Why should you use Probe?

Probe is a powerful tool for any Cosmos SDK developer seeking efficient, generalized querying on Cosmos blockchains. The package is finely tuned for transaction and message type querying.

:white_check_mark: **Benefits of using Probe:**

1. **Generalized Querying:** Simplified process for unpacking on-chain message types into their underlying Go types.

2. **Module Support:** Unpacking process for base Cosmos SDK types, and a comprehensive solution for customized modules.

3. **Lightweight Design:** Minimalistic design that places querying at the forefront.

---

## :satellite: Supported Protos

### :arrow_forward: Tendermint Liquidity v1beta1 Transactions

This section provides an overview of the protocol buffer messages for Tendermint Liquidity v1beta1 transaction messages. Protocol buffers are a platform-neutral, extendable mechanism for serializing structured data.

---

#### :bookmark_tabs: Transaction Types

The file encompasses several types of transactions that can be sent to the Tendermint Liquidity protocol:

- **Deposit:** Deposit funds into a liquidity pool.
- **Withdraw:** Withdraw funds from a liquidity pool.
- **Swap:** Swap tokens between two liquidity pools.
- **ProvideLiquidity:** Provide liquidity to a liquidity pool.
- **RemoveLiquidity:** Remove liquidity from a liquidity pool.
- **SwapExactTokensForTokens:** Swap tokens between two liquidity pools, ensuring the user receives the same total amount of tokens.

---

#### :rocket: tendermint.liquidity.v1beta1

This protobuf file outlines the `tendermint.liquidity.v1beta1` package and the messages associated with liquidity pool operations on the Cosmos blockchain. It includes four RPC services for the liquidity pool:

- **CreatePool:** Submit a liquidity pool creation transaction.
- **DepositWithinBatch:** Submit a deposit to the liquidity pool batch.
- **WithdrawWithinBatch:** Submit a withdrawal from the liquidity pool batch.
- **Swap:** Submit a swap to the liquidity pool batch.

Each operation has a corresponding response message. This definition allows Probe to comprehend these liquidity pool operations and handle them accurately when interacting with the Cosmos blockchain.

---

## :memo: Attribution

Probe was originally inspired by [Strangelove Ventures Lens](https://github.com/strangelove-ventures/lens) (discontinued). Files containing modified source code from this repository are distinctly marked with the following header:

```markdown
// APACHE NOTICE
// Sourced with modifications from https://github.com/strangelove-ventures/lens
```
