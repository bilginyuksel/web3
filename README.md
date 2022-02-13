## Web 3

We need wallet to do any kind of action with coins.

1. Create Wallet

```bash
mkdir wallets
mkdir wallets/solana

solana-keygen new --outfile wallets/solana/first.json
```

Get the public key. Verify wallet after creation

```bash
solana-keygen verify <pub> wallets/solana/first.json
```


2. Send coin from dev.net to the wallet.

```bash
solana airdrop 1 <receiver> --url https://api.devnet.solana.com
```

3. Check the wallet balance

Url could be changed to localhost or main-net or test-net.

```bash
solana balance <pubkey> --url https://api.devnet.solana.com
```

4. Create another wallet and make transfer

```bash
solana-keygen new --outfile wallets/solana/second.json

solana transfer --from wallets/solana/second.json <recipient> 0.5 \
--allow-unfunded-recipient \
--fee-payer wallets/solana/second.json \
--url https://api.devnet.solana.com
```

5. Create stake account

Create a new wallet for stake acocunt.

```bash
solana-keygen new -o wallets/solana/staking.json
```

Create stake account with the wallet already created.

```bash
solana create-stake-account --from wallets/solana/default.json \
    wallets/solana/staking.json 1.2 \
    --stake-authority <keypair> \
    --withdraw-authority <keypair> \
    --fee-payer <keypair> \
    --url https://api.devnet.solana.com
```

View the new stake account

```bash
solana stake-account <stake-account-address> --url https://api.devnet.solana.com
```
