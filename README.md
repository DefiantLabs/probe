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

## :satellite: Supported Message Type Protos


This section provides an overview of the protocol buffer messages included in the package for customized Cosmos modules. Protocol buffers are a platform-neutral, extendable mechanism for serializing structured data.

### :arrow_forward: Tendermint Liquidity v1beta1 Transactions

#### :rocket: tendermint.liquidity.v1beta1

This protobuf file outlines the `tendermint.liquidity.v1beta1` package and the messages associated with liquidity pool operations on the Cosmos blockchain. It includes four RPC services for the liquidity pool:

- **CreatePool:** Submit a liquidity pool creation transaction.
- **DepositWithinBatch:** Submit a deposit to the liquidity pool batch.
- **WithdrawWithinBatch:** Submit a withdrawal from the liquidity pool batch.
- **Swap:** Submit a swap to the liquidity pool batch.

Each operation has a corresponding response message. This definition allows Probe to comprehend these liquidity pool operations and handle them accurately when interacting with the Cosmos blockchain.

---
## :rocket: Probe vs Hubl: A Comparative Overview :milky_way:

Both **Probe** and **Hubl** are mighty tools for querying Cosmos SDK based blockchains, each catering to different needs and offering unique capabilities.

---

### :flying_saucer: Probe: A Deep Dive into Cosmos SDK Types

Probe excels in its ability to handle Cosmos SDK types and specialized modules. It provides a detailed process for unpacking on-chain message types into their underlying Go types, shining particularly bright when working with chains that lack reflection support.

Key Highlights:

- :dart: **Generalized Querying**: Finely tuned for transaction and message type querying, Probe focuses on what matters most.
- :package: **Module Support**: Comprehensive solution for unpacking base Cosmos SDK types and handling customized modules.
- :dove_of_peace: **Lightweight Design**: A minimalistic design that places querying at the forefront, making it a go-to tool for developers.
- :medal: **Continuity and Evolution**: Probe was originally inspired by the now-discontinued Strangelove Ventures Lens, showing a strong lineage and progressive improvement in the Cosmos SDK querying tools.

---

### :stars: Hubl: Harnessing the Power of AutoCLI and Chain-Registry

Hubl makes full use of the new AutoCLI feature of the Cosmos SDK, enabling querying of any Cosmos SDK based blockchain that supports reflection. It's the tool of choice for those seeking compatibility and ease of use with the chain-registry.

Key Highlights:

- :control_knobs: **AutoCLI Advantage**: Hubl allows you to query any Cosmos SDK based blockchain that supports reflection, making it a versatile tool for diverse chains.
- :books: **Chain-Registry Compatibility**: Hubl integrates seamlessly with the chain-registry, simplifying the process of configuring and querying multiple chains.
- :three_button_mouse: **User-Friendly Commands**: With straightforward commands, Hubl enables easy interaction with a wide range of chains.

---

The choice between Probe and Hubl largely depends on your specific needs, the features of the chains you're working with, and your preferred method of interaction. Both tools offer unique advantages, and there might be additional nuanced differences that can be uncovered with a deeper dive into the functionalities and use cases of each tool.

## :memo: Attribution

Probe was originally inspired by [Strangelove Ventures Lens](https://github.com/strangelove-ventures/lens) (discontinued). Files containing modified source code from this repository are distinctly marked with the following header:

```markdown
// APACHE NOTICE
// Sourced with modifications from https://github.com/strangelove-ventures/lens
```
