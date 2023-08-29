## epay cli

A command line tool for epay testing.

### Build

```bash
make build
```

### Usage

Based on previous build, you can find the binary file in `bin/epay-cli`.

We provide a simple command line tool for epay testing and maintenance.

#### Test for API

```bash
./bin/epay-cli test submit # test submit api
./bin/epay-cli test mapi # test mapi submit api
```

#### Migration

```bash
./bin/epay-cli migrate # generate sql schema
```

#### Help

```bash
./bin/epay-cli help
```
