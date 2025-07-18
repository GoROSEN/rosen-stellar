# ROSEN × Stellar

## 🌹 ROSEN × Stellar

The Next-Gen Social Earning App
Seamless global connections, powered by Stellar-Native compliant stablecoins USDC/PYUSD!

## 🏆 Why ROSEN × Stellar?

Stellar's near-zero transaction fees, sub-5-second settlement times, and compliance-focused design make it the perfect home for ROSEN's vision of borderless, inclusive social economies. Its robust architecture is ideal for high-frequency, low-value payments that power our global community.

## The Problems We Solve

1.Global Payment Friction
  - Digital natives in emerging markets face high fees and slow settlements for small cross-border payments.

2.Limited Economic Opportunities
  - Individuals struggle to monetize their skills globally due to payment barriers and lack of accessible micro-earning platforms.

3.Onboarding Complexity
  - Traditional crypto experiences deter everyday users with complex wallets and confusing gas fees.

# Our Stellar-Specific Solutions

## Cross-Border Micro-Earning

- **Instant Tipping**: Global users can instantly and affordably tip creators or helpful community members with USDC on Stellar, free from exchange rate concerns or high fees.

- **Micro-Task Rewards**: Users earn USDC on Stellar for completing micro-tasks (e.g., content creation, community support) directly within Rosen, fostering a liquid digital labor market.

## Global Decentralized Recruitment

- **Talent Matching**: Businesses can post jobs and pay incentives or salaries in USDC on Stellar, facilitating efficient and transparent global talent matching.

- **Community Incentives**: USDC on Stellar is used for in-community voting and contribution rewards, strengthening decentralized social networks and incentivizing participation.

## 🎮 Key Features

Feature | Tech Stack / Stellar Integration | Competitive Edge
------- | ---------- | -------------------
1-Click PYUSD Onboarding | Circle Programmable Wallets + Stellar Account Creation | No complex seed phrases, just social login for instant access to Stellar assets.
Gasless Microtransactions | Stellar-Native USDC enabled deposit and withdrawal + Fee Abstraction | Send $10 as easily as a text message; ROSEN covers all gas fees (gas tokes).
Cross-Border Social Economy |  AI Translation + In-Chat Stablecoins Payments  | Pay global talent and reward community members in seconds, breaking down language and payment barriers.
Non-Custodial Withdrawals |  Stellar SEP-0007 (Deep Links) for Wallet Connect | Users can easily connect and withdraw Stellar assets to their preferred non-custodial wallets (e.g., Lobstr, Albedo).

## ⚙️ Why This Matters to Stellar

Metric |	Traditional Apps |	ROSEN × Stellar
------ | ----------------- | --------------------
Min. Transfer |	$100+ |	$5+
Tx Fees |	$20+ per tx |	$0 (covered by Rosen)
Settlement Time	| 1-5 days	| < 20 seconds
Global Reach	| Single-language, several countries | AI-translated for 50+ languages, global

**Drives Stellar Adoption**: Onboards non-crypto users by offering a seamless, mobile-first experience that leverages Stellar's core strengths.

1. **Showcases Stellar's Utility**: Demonstrates the real-world power of Stellar for high-frequency, low-value cross-border payments crucial for the gig economy.
2. **Expands Ecosystem Value**: Creates new economic opportunities and use cases for USDC and eventually PYUSD on the Stellar network.

## 🛠️ Setup & Run

### Architecture

![arch_img](./images/arch.png)

### Quick Start

You can download our App from [App Store](https://apps.apple.com/us/app/rosen/id6444627514) / [Google Play](https://play.google.com/store/apps/details?id=com.rosenbridge.rosen&pli=1) for a quick experience.

And you can also download the [Beta Version](https://expo.dev/accounts/rosen-bridge/projects/rosen/builds/4efa0f19-da05-450a-a2c6-3a562d60ebbc) for Start earning and sending USDC on Stellar within your social circles!

### Build & Run Client

```
cd client
npm i
npm run dev
```

### Build & Run Server

1. requirements: 

  - golang >= 1.23.4
  - mysql >= 8.0
  - redis

2. build

```
cd server-contract
go mod download
go build
```

2. Edit ```config.yaml``` with your favorate editor. e.g. ```vim```

```
vim config.yaml
```

3. Launch for first time.

```
./rosen-apiserver server --config config.yaml --migratedb yes
```

## 📈 Future Roadmap

- **PYUSD on Stellar Integration**: Full integration of PayPal USD once it officially launches on the Stellar mainnet and secures necessary regulatory approvals.

- **Stellar-Native On/Off Ramps**: Explore direct fiat on/off-ramp solutions for Stellar stablecoins through partnerships and Stellar Ecosystem Proposals (SEPs).

- **Expanded Stellar Ecosystem Assets**: Evaluate and integrate other compliant and valuable assets on the Stellar network to diversify user options.

## 📜 License

MIT © ROSEN Team.

We're not just building a social app – we're creating the economic infrastructure for a truly global, inclusive community powered by Stellar.
