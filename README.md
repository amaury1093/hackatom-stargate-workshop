# "Build Your Own Module - Stargate Edition"

This repository contains support material for the online workshop "Build Your Own Module - Stargate Edition", held during the [HackAtom V](https://five.hackatom.org/) hackathon.

The accompanying slides can be found [here](TODO). If you didn't attend the workshop, you can read the slides, and follow the tutorial steps in this repository along them.

The workshop is divided into 3 steps:

- Step 1: Create a blank blockchain using the [`starport`](https://github.com/tendermint/starport) scaffolding tool.
- Step 2: Create your own Nameservice module.
- Step 3: Run the chain, and interact with your module.

### Requirements

- [`golang`](https://golang.org/doc/install) >1.15.0 installed
- The [starport](https://github.com/tendermint/starport) tool will be used to go through this tutorial. The fastest way to install it is via npm (`npm i -g @tendermint/starport`) or brew (`brew install tendermint/tap/starport`). You could also clone the repository and build it from source:

```bash
git clone https://github.com/tendermint/starport && cd starport && make
```

## Step 1: Create a blank blockchain using the [`starport`](https://github.com/tendermint/starport) scaffolding tool.

```bash
# In a clean directory, run the starport command to create a new blockchain skeleton.
starport app github.com/{your_gh_username}/nameservice --sdk-version stargate
cd nameservice
# starport v0.11.1 hasn't been updated to the latest Cosmos SDK v0.40.0-rc0, so we update it manually.
go get github.com/cosmos/cosmos-sdk@v0.40.0-rc0
# We will only use the following modules:
# - auth, bank, staking, distribution, slashing and our own nameservice
# Inside the app/ folder, delete all mentions of other modules:
# - mint, upgrade, gov, genutil, capability, ibc*

# Run the starport command to make sure the app builds
starport build
```

At the end of this step, you should have a folder structure similar to the `step1/` folder.

> Note: For the sake of this tutorial, the app in the `step1/` folder has been renamed ~`nameservice`~ -> `step1`

## Step 2

```bash
# Add a new type to your x/nameservice module.
starport type name value price

```

> Note: For clarity purpose, the app in the `step1/` folder has been renamed ~`nameservice`~ -> `step2`
