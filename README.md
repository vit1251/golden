# Golden Point 1.x

Golden Point is Fido point package

For work with GoldenPoint after starting you will open browser on address http://127.0.0.1:8080

## Features

 - [x] Mailer
   - [x] [FTS-1026] Binkp/1.0 minimum protocol realization
     - [x] password protected sessions
     - [x] 5D addressing for Fidonet and technology compatibele networks
     - [x] exchange of netmail packets and archmail bundles in both
           directions, including poll and mail pickup, as well as transfer
           of any binary or ASCII files
     - [x] ensuring integrity of transmitted mail and files
     - [x] simultaneus bi-directional transmission
     - [x] maximizing performance over packet switched data networks
   - [x] [FTS-1027] Binkp/1.0 optional protocol extension CRAM
 - [x] Tosser
   - [x] [FTS-0009] MSGID / REPLY; A standard for unique message identifiers and reply chain linkage
   - [x] [FTS-0001] A Basic FidoNet(r) Technical Standard
   - [ ] [FTS-0039]	A Type-2 Packet Extension Proposal
   - [x] [FTS-4000] CONTROL PARAGRAPHS
   - [x] [FTS-4001] ADDRESSING CONTROL PARAGRAPHS
 - [x] Tracker
   - [ ] TIC parser
   - [ ] TIC creator
 - [x] Editor
   - [x] Create NETMAIL messages
   - [x] Create ECHOMAIL messages

## Documentation

You may reading more documentation on https://github.com/vit1251/golden/tree/master/docs

## Binary builds

You may download binary builds

 - Unstable version (night) build provided in CI/CD system on https://github.com/vit1251/golden/actions
 - Stable version (release) build provided in https://github.com/vit1251/golden/releases

## Building from source code

In most cases, there is no need to build from source. The exceptions are when these build produce
by node owners to add additional or specific attribute to their Points.

## Contributors

 * Vitold Sedyshev
 * Sergey Anohin
