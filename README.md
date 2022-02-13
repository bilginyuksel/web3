## Web 3

We need wallet to do any kind of action with coins.

1. Create Wallet

```bash
mkdir wallets
mkdir wallets/solana

solana-keygen new --outfile wallets/solana/first.json
```

Get the public key.

2. Check the wallet balance

Url could be changed to localhost or main-net or test-net.

```bash
solana balance <pubkey> --url https://api.devnet.solana.com
```

3. Create another wallet and make transfer

```bash
solana-keygen new --outfile wallets/solana/second.json

solana transfer --from wallets/solana/second.json <recipient-pub> 0.5 --allow-unfunded-recipient --url https://api.devnet.solana.com --fee-payer wallets/solana/second.json
```
