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

solana transfer --from wallets/solana/second.json <recipient> 0.5 --allow-unfunded-recipient --url https://api.devnet.solana.com --fee-payer wallets/solana/second.json
```
