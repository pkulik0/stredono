# Stredono

## Requirements

- Node.js
- pnpm
- Go
- Terraform CLI
- GCP SDK
- Python 3

## Hooks

### Setup

```bash
./scripts/setup_hooks.sh
```

### Available Hooks

- `pre-commit` - enforces `terraform fmt` and `terraform validate` before commit
- `post-checkout` switches to the correct terraform workspace based on the branch name and regenerates the output files


## Donation Flow

### Flowchart

```mermaid
flowchart TB
    V((Viewer)) --Visits--> DP([Donation Page])
    DP --Send--> SDF[Send Donate Function]
    
    SDF --Start payment--> PP(Payment Provider)
    PP --Return payment link--> SDF
    SDF --Save payment-->DB
    
    SDF --Return payment link--> DP
    DP --Redirects to payment link--> V


    PP(Payment Provider) --Confirms payment--> CPF[Confirm Payment Function]
    CPF --Update payment--> DB[(Database)]
    CPF --Publishes confirmation--> PS[[Pub/Sub]]
    
    PS --Subscribes--> WF[Websocket Feed]

    WF --> AP([Alert Page])
    AP --Shows alert--> S((Streamer))


    SD([Streamer Dashboard]) --- SS
    SS[Streamer Service] --- DB
    WF --> SD
    SD --- S
    
```

### Components

```mermaid
flowchart TB
    A((User))
    subgraph Need to be implemented
        B([Front-facing App])
        C[Cloud Function]
    end
    D(External Service)
    E[[Google Cloud Service]]
    F[(Database)]
```