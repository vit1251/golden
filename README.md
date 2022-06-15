# Golden Point 1.x

![Golden Point](/docs/images/GoldenPointMessage.png)

Golden Point is a FidoNet (FTN) point package written with in Golang to provide a mailer, tosser and other related utilities.

To work with GoldenPoint after starting you will need to open your browser on address http://127.0.0.1:8080

Project status can be found at https://github.com/vit1251/golden/projects/6

## Features

Golden Point provides:

 - [x] Mailer
   - [x] [FTS-1026] Binkp/1.0 minimum protocol realization
     - [x] password protected sessions
     - [x] 5D addressing for Fidonet and FTN technology compatible networks
     - [x] exchange of netmail packets and archmail bundles in both
           directions, including poll and mail pickup, as well as transfer
           of any binary or ASCII files
     - [x] ensuring integrity of transmitted mail and files
     - [x] simultaneous bi-directional transmission
     - [x] maximizing performance over packet switched data networks
   - [x] [FTS-1027] Binkp/1.0 optional protocol extension CRAM
 - [x] Tosser
   - [x] [FSC-0001] A Basic FidoNet(r) Technical Standard
   - [x] [FTS-0009] MSGID / REPLY; A standard for unique message identifiers and reply chain linkage
   - [x] [FSC-0039] A Type-2 Packet Extension Proposal 
   - [x] [FTS-4000] CONTROL PARAGRAPHS
   - [x] [FTS-4001] ADDRESSING CONTROL PARAGRAPHS
   - [x] [FRL-1004] Timezone information in FTN messages
 - [x] Tracker
   - [x] [FTS-5006] TIC parser
   - [x] [FTS-5006] TIC builder

## Documentation

You may read more documentation on https://github.com/vit1251/golden/tree/master/docs

## Binary builds

You may download binary builds

 - The unstable version (night) build provided in CI/CD system on https://github.com/vit1251/golden/actions
 - Stable version (release) build provided in https://github.com/vit1251/golden/releases

## Building from source code

In most cases, there is no need to build from the source. But you may read about your platform compile
scenario in the "docs" directory instructions. Common scripts are:

    # go generate
    # go build
    # ./golden --debug

That's all.
