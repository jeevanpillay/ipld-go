# IPLD Go Implementation Details
This repository maintains methods that all the users to create multi-hash DAG's through the `ipfs-go-api` connector. Then, these multihashes are sent to IPFS system.

## Outputs
The out-of-the-box implementation doesn't require the user to install IPFS (or run a daemon), instead, it connects directly to Infura's Public IPFS Gateway.

```
sh = shell.NewShell("https://ipfs.infura.io:5001")
```

Change this to localhost if running locally.

## Inputs
The system stores inputs based on keyValueMap's, based on two parameters that the user specifies such as the `inputKey` and `inputValue`.
```
keyValueMap := make(map[string]interface{})
```

## Marshall
The key-value pair then gets marshalled into JSON. This will be useful later!

```
entryJSON, err := json.Marshal(keyValueMap)
```

## Conversion
The marshalled JSON then gets converted into it's CID, which, is actually a multihash construct of the key-value pair created from the input.
```
cid, err := sh.DagPut(entryJSON, "json", "cbor")

```

## Access Control
Since, we have specified to use INFURA's Public Host (as specified above), we can now view the website that stores this beautiful multi-hash value.
- [Tash Sultana](https://explore.ipld.io/#/explore/bafyreiarqq75yerg6zssp3sqtyrfeht4d4dxzzbnvtwiidsogg3oni5ali)
- [Pink Floyd](https://explore.ipld.io/#/explore/bafyreifq5edjluhqczjinj23k3matyusyojiwrmk3hfekq6sth3wbejaiy)
- [COVID-19](https://explore.ipld.io/#/explore/bafyreiebwgly5tutecxp3fohuwnz5z2c52jsww4pzjs76gcuxjbd3wsmbe)
