#!/bin/sh

/parity/parity --chain /chain.json \
               --password /.parity-pass \
               --unlock=0x00fc9af07fb2fcfc560672630b19bb8f72732164,0x006Db1ebe1D7da13Aa4B7523F7A4Ff123Cf98E58 \
               --base-path=/eth \
               --jsonrpc-cors "*" \
               --jsonrpc-interface all \
               --jsonrpc-hosts all \
               --reseal-min-period 0 \
               --tracing on \
               --unsafe-expose \
               --stratum

#curl --data $'{
#    "method": "eth_sendTransaction",
#    "params": [{
#        "from": "0x00fc9af07fb2fcfc560672630b19bb8f72732164",
#        "to": "0x006Db1ebe1D7da13Aa4B7523F7A4Ff123Cf98E58",
#        "gas": "0x5208",
#        "gasPrice": "0x3b9aca00",
#        "value": "0x4563918244f40000",
#        "data": "0x"
#    }],
#    "id": 1,
#    "jsonrpc": "2.0"
#}' -H "Content-Type: application/json" -X POST localhost:8545
#
#curl --data $'{
#    "method": "eth_signTransaction",
#    "params": [{
#        "from": "0x00fc9af07fb2fcfc560672630b19bb8f72732164",
#        "to": "0x006Db1ebe1D7da13Aa4B7523F7A4Ff123Cf98E58",
#        "gas": "0x5208",
#        "gasPrice": "0x3b9aca00",
#        "value": "0x4563918244f40000",
#        "data": "0x"
#    }],
#    "id": 3,
#    "jsonrpc": "2.0"
#}' -H "Content-Type: application/json" -X POST localhost:8545
#
#curl --data '{"method":"eth_sendRawTransaction","params":["0xf86d01843b9aca0082520894006db1ebe1d7da13aa4b7523f7a4ff123cf98e58884563918244f400008082017da0b7072a3f203635fc3d031b49b167273ae51d73d42ac9e0705e1855ce12f70489a03900c39b286d6e90de9501d1912730d53f71a5c595006373d9068c472cf4fc97"],"id":5,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:8545
#
#{"jsonrpc":"2.0","result":{"raw":"0xf86d0a843b9aca0082520894006db1ebe1d7da13aa4b7523f7a4ff123cf98e58884563918244f400008082017ea076b12808b7263716008c1868e1fb8766268eb837a5b7ec85fd743dc0c2d1f86da033209a57f52534d90855afcb8014c34d3c7dc9826e32481302ae807acb1d22bd","tx":{"blockHash":null,"blockNumber":null,"chainId":"0xad","condition":null,"creates":null,"from":"0x00fc9af07fb2fcfc560672630b19bb8f72732164","gas":"0x5208","gasPrice":"0x3b9aca00","hash":"0x77383216ea9a1f48a769c98ff8a9aff5837aeb647a6883fea2332d8449c264df","input":"0x","nonce":"0xa","publicKey":"0xb9b104599786bf1bd0cbf924c751e65691bd38494ae334b3ce4a35279185d16ce0277a60314d29b0e696213e99c758bffd50acb694e3371d629032eaa0ef3c70","r":"0x76b12808b7263716008c1868e1fb8766268eb837a5b7ec85fd743dc0c2d1f86d","raw":"0xf86d0a843b9aca0082520894006db1ebe1d7da13aa4b7523f7a4ff123cf98e58884563918244f400008082017ea076b12808b7263716008c1868e1fb8766268eb837a5b7ec85fd743dc0c2d1f86da033209a57f52534d90855afcb8014c34d3c7dc9826e32481302ae807acb1d22bd","s":"0x33209a57f52534d90855afcb8014c34d3c7dc9826e32481302ae807acb1d22bd","standardV":"0x1","to":"0x006db1ebe1d7da13aa4b7523f7a4ff123cf98e58","transactionIndex":null,"v":"0x17e","value":"0x4563918244f40000"}},"id":2}

#               //--author cf3baf57b91a3bd96a76150585209dfb4fa0eb90 \
#               //--unlock cf3baf57b91a3bd96a76150585209dfb4fa0eb90 \
#               //--password /.parity-pass \
#Operating Options:
#    --public-node
#        Start Parity as a public web server. Account storage and transaction signing will be delegated to the UI.
#
#    --no-download
#        Normally new releases will be downloaded ready for updating. This disables it. Not recommended.
#
#    --no-consensus
#        Force the binary to run even if there are known issues regarding consensus. Not recommended.
#
#    --light
#        Experimental: run in light client mode. Light clients synchronize a bare minimum of data and fetch necessary
#        data on-demand from the network. Much lower in storage, potentially higher in bandwidth. Has no effect with
#        subcommands.
#
#    --no-hardcoded-sync
#        By default, if there is no existing database the light client will automatically jump to a block hardcoded in
#        the chain's specifications. This disables this feature.
#
#    --force-direct
#        Run the originally installed version of Parity, ignoring any updates that have since been installed.
#
#    --mode=[MODE]
#        Set the operating mode. MODE can be one of: last - Uses the last-used mode, active if none; active - Parity
#        continuously syncs the chain; passive - Parity syncs initially, then sleeps and wakes regularly to resync; dark
#        - Parity syncs only when the RPC is active; offline - Parity doesn't sync. (default: last)
#
#    --mode-timeout=[SECS]
#        Specify the number of seconds before inactivity timeout occurs when mode is dark or passive (default: 300)
#
#    --mode-alarm=[SECS]
#        Specify the number of seconds before auto sleep reawake timeout occurs when mode is passive (default: 3600)
#
#    --auto-update=[SET]
#        Set a releases set to automatically update and install. SET can be one of: all - All updates in the our release
#        track; critical - Only consensus/security updates; none - No updates will be auto-installed. (default: critical)
#
#    --auto-update-delay=[NUM]
#        Specify the maximum number of blocks used for randomly delaying updates. (default: 100)
#
#    --auto-update-check-frequency=[NUM]
#        Specify the number of blocks between each auto-update check. (default: 20)
#
#    --release-track=[TRACK]
#        Set which release track we should use for updates. TRACK can be one of: stable - Stable releases; beta - Beta
#        releases; nightly - Nightly releases (unstable); testing - Testing releases (do not use); current - Whatever
#        track this executable was released on. (default: current)
#
#    --chain=[CHAIN]
#        Specify the blockchain type. CHAIN may be either a JSON chain specification file or olympic, frontier,
#        homestead, mainnet, morden, ropsten, classic, expanse, tobalaba, musicoin, ellaism, easthub, social, testnet,
#        kovan or dev. (default: foundation)
#
#    --keys-path=[PATH]
#        Specify the path for JSON key files to be found (default: $BASE/keys)
#
#    --identity=[NAME]
#        Specify your node's name. (default: )
#
#    -d, --base-path=[PATH]
#        Specify the base data storage path.
#
#    --db-path=[PATH]
#        Specify the database directory path
#
#Convenience options:
#    --unsafe-expose
#        All servers will listen on external interfaces and will be remotely accessible. It's equivalent with setting the
#        following: --{{ws,jsonrpc,ui,ipfs,secret_store,stratum}}-interface=all --*-hosts=all    This option is UNSAFE
#        and should be used with great care!
#
#    -c, --config=[CONFIG]
#        Specify a configuration. CONFIG may be either a configuration file or a preset: dev, insecure, dev-insecure,
#        mining, or non-standard-ports. (default: $BASE/config.toml)
#
#    --ports-shift=[SHIFT]
#        Add SHIFT to all port numbers Parity is listening on. Includes network port and all servers (RPC, WebSockets,
#        UI, IPFS, SecretStore). (default: 0)
#
#Account options:
#    --no-hardware-wallets
#        Disables hardware wallet support.
#
#    --fast-unlock
#        Use drasticly faster unlocking mode. This setting causes raw secrets to be stored unprotected in memory, so use
#        with care.
#
#    --keys-iterations=[NUM]
#        Specify the number of iterations to use when deriving key from the password (bigger is more secure) (default:
#        10240)
#
#    --accounts-refresh=[TIME]
#        Specify the cache time of accounts read from disk. If you manage thousands of accounts set this to 0 to disable
#        refresh. (default: 5)
#
#    --unlock=[ACCOUNTS]
#        Unlock ACCOUNTS for the duration of the execution. ACCOUNTS is a comma-delimited list of addresses. Implies
#        --no-ui.
#
#    --password=[FILE]...
#        Provide a file containing a password for unlocking an account. Leading and trailing whitespace is trimmed.
#        (default: [])
#
#Private transactions options:
#    --private-tx-enabled
#        Enable private transactions.
#
#    --private-signer=[ACCOUNT]
#        Specify the account for signing public transaction created upon verified private transaction.
#
#    --private-validators=[ACCOUNTS]
#        Specify the accounts for validating private transactions. ACCOUNTS is a comma-delimited list of addresses.
#
#    --private-account=[ACCOUNT]
#        Specify the account for signing requests to secret store.
#
#    --private-sstore-url=[URL]
#        Specify secret store URL used for encrypting private transactions.
#
#    --private-sstore-threshold=[NUM]
#        Specify secret store threshold used for encrypting private transactions.
#
#    --private-passwords=[FILE]...
#        Provide a file containing passwords for unlocking accounts (signer, private account, validators).
#
#UI options:
#    --force-ui
#        Enable Trusted UI WebSocket endpoint, even when --unlock is in use.
#
#    --no-ui
#        Disable Trusted UI WebSocket endpoint.
#
#    --ui-no-validation
#        Disable Origin and Host headers validation for Trusted UI. WARNING: INSECURE. Used only for development.
#
#    --ui-interface=[IP]
#        Specify the hostname portion of the Trusted UI server, IP should be an interface's IP address, or local.
#        (default: local)
#
#    --ui-hosts=[HOSTS]
#        List of allowed Host header values. This option will validate the Host header sent by the browser, it is
#        additional security against some attack vectors. Special options: "all", "none",. (default: none)
#
#    --ui-path=[PATH]
#        Specify directory where Trusted UIs tokens should be stored. (default: $BASE/signer)
#
#    --ui-port=[PORT]
#        Specify the port of Trusted UI server. (default: 8180)
#
#Networking options:
#    --no-warp
#        Disable syncing from the snapshot over the network.
#
#    --no-discovery
#        Disable new peer discovery.
#
#    --reserved-only
#        Connect only to reserved nodes.
#
#    --no-ancient-blocks
#        Disable downloading old blocks after snapshot restoration or warp sync.
#
#    --no-serve-light
#        Disable serving of light peers.
#
#    --warp-barrier=[NUM]
#        When warp enabled never attempt regular sync before warping to block NUM.
#
#    --port=[PORT]
#        Override the port on which the node should listen. (default: 30303)
#
#    --min-peers=[NUM]
#        Try to maintain at least NUM peers.
#
#    --max-peers=[NUM]
#        Allow up to NUM peers.
#
#    --snapshot-peers=[NUM]
#        Allow additional NUM peers for a snapshot sync. (default: 0)
#
#    --nat=[METHOD]
#        Specify method to use for determining public address. Must be one of: any, none, upnp, extip:<IP>. (default:
#        any)
#
#    --allow-ips=[FILTER]
#        Filter outbound connections. Must be one of: private - connect to private network IP addresses only; public -
#        connect to public network IP addresses only; all - connect to any IP address. (default: all)
#
#    --max-pending-peers=[NUM]
#        Allow up to NUM pending connections. (default: 64)
#
#    --network-id=[INDEX]
#        Override the network identifier from the chain we are on.
#
#    --bootnodes=[NODES]
#        Override the bootnodes from our chain. NODES should be comma-delimited enodes.
#
#    --node-key=[KEY]
#        Specify node secret key, either as 64-character hex string or input to SHA3 operation.
#
#    --reserved-peers=[FILE]
#        Provide a file containing enodes, one per line. These nodes will always have a reserved slot on top of the
#        normal maximum peers.
#
#API and console options – RPC:
#    --no-jsonrpc
#        Disable the JSON-RPC API server.
#
#    --jsonrpc-port=[PORT]
#        Specify the port portion of the JSONRPC API server. (default: 8545)
#
#    --jsonrpc-interface=[IP]
#        Specify the hostname portion of the JSONRPC API server, IP should be an interface's IP address, or all (all
#        interfaces) or local. (default: local)
#
#    --jsonrpc-apis=[APIS]
#        Specify the APIs available through the JSONRPC interface using a comma-delimited list of API names. Possible
#        names are: all, safe, web3, net, eth, pubsub, personal, signer, parity, parity_pubsub, parity_accounts,
#        parity_set, traces, rpc, secretstore, shh, shh_pubsub. You can also disable a specific API by putting '-' in the
#        front, example: all,-personal. safe contains following apis: web3, net, eth, pubsub, parity, parity_pubsub,
#        traces, rpc, shh, shh_pubsub (default:
#        web3,eth,pubsub,net,parity,private,parity_pubsub,traces,rpc,shh,shh_pubsub)
#
#    --jsonrpc-hosts=[HOSTS]
#        List of allowed Host header values. This option will validate the Host header sent by the browser, it is
#        additional security against some attack vectors. Special options: "all", "none",. (default: none)
#
#    --jsonrpc-threads=[THREADS]
#        Turn on additional processing threads in all RPC servers. Setting this to non-zero value allows parallel cpu
#        -heavy queries execution. (default: 4)
#
#    --jsonrpc-cors=[URL]
#        Specify CORS header for JSON-RPC API responses. Special options: "all", "none". (default: none)
#
#    --jsonrpc-server-threads=[NUM]
#        Enables multiple threads handling incoming connections for HTTP JSON-RPC server.
#
#API and console options – WebSockets:
#    --no-ws
#        Disable the WebSockets server.
#
#    --ws-port=[PORT]
#        Specify the port portion of the WebSockets server. (default: 8546)
#
#    --ws-interface=[IP]
#        Specify the hostname portion of the WebSockets server, IP should be an interface's IP address, or all (all
#        interfaces) or local. (default: local)
#
#    --ws-apis=[APIS]
#        Specify the APIs available through the WebSockets interface using a comma-delimited list of API names. Possible
#        names are: all, safe, web3, net, eth, pubsub, personal, signer, parity, parity_pubsub, parity_accounts,
#        parity_set, traces, rpc, secretstore, shh, shh_pubsub. You can also disable a specific API by putting '-' in the
#        front, example: all,-personal. safe contains following apis: web3, net, eth, pubsub, parity, parity_pubsub,
#        traces, rpc, shh, shh_pubsub (default:
#        web3,eth,pubsub,net,parity,parity_pubsub,private,traces,rpc,shh,shh_pubsub)
#
#    --ws-origins=[URL]
#        Specify Origin header values allowed to connect. Special options: "all", "none". (default: parity://*,chrome
#        -extension://*,moz-extension://*)
#
#    --ws-hosts=[HOSTS]
#        List of allowed Host header values. This option will validate the Host header sent by the browser, it is
#        additional security against some attack vectors. Special options: "all", "none". (default: none)
#
#    --ws-max-connections=[CONN]
#        Maximal number of allowed concurrent WS connections. (default: 100)
#
#API and console options – IPC:
#    --no-ipc
#        Disable JSON-RPC over IPC service.
#
#    --ipc-path=[PATH]
#        Specify custom path for JSON-RPC over IPC service. (default: $BASE/jsonrpc.ipc)
#
#    --ipc-apis=[APIS]
#        Specify custom API set available via JSON-RPC over IPC using a comma-delimited list of API names. Possible names
#        are: all, safe, web3, net, eth, pubsub, personal, signer, parity, parity_pubsub, parity_accounts, parity_set,
#        traces, rpc, secretstore, shh, shh_pubsub. You can also disable a specific API by putting '-' in the front,
#        example: all,-personal. safe contains: web3, net, eth, pubsub, parity, parity_pubsub, traces, rpc, shh,
#        shh_pubsub (default: web3,eth,pubsub,net,parity,parity_pubsub,parity_accounts,private,traces,rpc,shh,shh_pubsub)
#
#API and console options – Dapps:
#    --no-dapps
#        Disable the Dapps server (e.g. status page).
#
#    --dapps-path=[PATH]
#        Specify directory where dapps should be installed. (default: $BASE/dapps)
#
#API and console options – IPFS:
#    --ipfs-api
#        Enable IPFS-compatible HTTP API.
#
#    --ipfs-api-port=[PORT]
#        Configure on which port the IPFS HTTP API should listen. (default: 5001)
#
#    --ipfs-api-interface=[IP]
#        Specify the hostname portion of the IPFS API server, IP should be an interface's IP address or local. (default:
#        local)
#
#    --ipfs-api-hosts=[HOSTS]
#        List of allowed Host header values. This option will validate the Host header sent by the browser, it is
#        additional security against some attack vectors. Special options: "all", "none". (default: none)
#
#    --ipfs-api-cors=[URL]
#        Specify CORS header for IPFS API responses. Special options: "all", "none". (default: none)
#
#Secret store options:
#    --no-secretstore
#        Disable Secret Store functionality.
#
#    --no-secretstore-http
#        Disable Secret Store HTTP API.
#
#    --no-acl-check
#        Disable ACL check (useful for test environments).
#
#    --no-secretstore-auto-migrate
#        Do not run servers set change session automatically when servers set changes. This option has no effect when
#        servers set is read from configuration file.
#
#    --secretstore-contract=[SOURCE]
#        Secret Store Service contract address source: none, registry (contract address is read from secretstore_service
#        entry in registry) or address. (default: none)
#
#    --secretstore-srv-gen-contract=[SOURCE]
#        Secret Store Service server key generation contract address source: none, registry (contract address is read
#        from secretstore_service_srv_gen entry in registry) or address. (default: none)
#
#    --secretstore-srv-retr-contract=[SOURCE]
#        Secret Store Service server key retrieval contract address source: none, registry (contract address is read from
#        secretstore_service_srv_retr entry in registry) or address. (default: none)
#
#    --secretstore-doc-store-contract=[SOURCE]
#        Secret Store Service document key store contract address source: none, registry (contract address is read from
#        secretstore_service_doc_store entry in registry) or address. (default: none)
#
#    --secretstore-doc-sretr-contract=[SOURCE]
#        Secret Store Service document key shadow retrieval contract address source: none, registry (contract address is
#        read from secretstore_service_doc_sretr entry in registry) or address. (default: none)
#
#    --secretstore-nodes=[NODES]
#        Comma-separated list of other secret store cluster nodes in form NODE_PUBLIC_KEY_IN_HEX@NODE_IP_ADDR:NODE_PORT.
#        (default: )
#
#    --secretstore-interface=[IP]
#        Specify the hostname portion for listening to Secret Store Key Server internal requests, IP should be an
#        interface's IP address, or local. (default: local)
#
#    --secretstore-port=[PORT]
#        Specify the port portion for listening to Secret Store Key Server internal requests. (default: 8083)
#
#    --secretstore-http-interface=[IP]
#        Specify the hostname portion for listening to Secret Store Key Server HTTP requests, IP should be an interface's
#        IP address, or local. (default: local)
#
#    --secretstore-http-port=[PORT]
#        Specify the port portion for listening to Secret Store Key Server HTTP requests. (default: 8082)
#
#    --secretstore-path=[PATH]
#        Specify directory where Secret Store should save its data. (default: $BASE/secretstore)
#
#    --secretstore-secret=[SECRET]
#        Hex-encoded secret key of this node.
#
#    --secretstore-admin=[PUBLIC]
#        Hex-encoded public key of secret store administrator.
#
#Sealing/Mining options:
#    --force-sealing
#        Force the node to author new blocks as if it were always sealing/mining.
#
#    --reseal-on-uncle
#        Force the node to author new blocks when a new uncle block is imported.
#
#    --remove-solved
#        Move solved blocks from the work package queue instead of cloning them. This gives a slightly faster import
#        speed, but means that extra solutions submitted for the same work package will go unused.
#
#    --tx-queue-no-unfamiliar-locals
#        Transactions recieved via local means (RPC, WS, etc) will be treated as external if the sending account is
#        unknown.
#
#    --refuse-service-transactions
#        Always refuse service transactions.
#
#    --infinite-pending-block
#        Pending block will be created with maximal possible gas limit and will execute all transactions in the queue.
#        Note that such block is invalid and should never be attempted to be mined.
#
#    --no-persistent-txqueue
#        Don't save pending local transactions to disk to be restored whenever the node restarts.
#
#    --stratum
#        Run Stratum server for miner push notification.
#
#    --reseal-on-txs=[SET]
#        Specify which transactions should force the node to reseal a block. SET is one of: none - never reseal on new
#        transactions; own - reseal only on a new local transaction; ext - reseal only on a new external transaction; all
#        - reseal on all new transactions. (default: own)
#
#    --reseal-min-period=[MS]
#        Specify the minimum time between reseals from incoming transactions. MS is time measured in milliseconds.
#        (default: 2000)
#
#    --reseal-max-period=[MS]
#        Specify the maximum time since last block to enable force-sealing. MS is time measured in milliseconds.
#        (default: 120000)
#
#    --work-queue-size=[ITEMS]
#        Specify the number of historical work packages which are kept cached lest a solution is found for them later.
#        High values take more memory but result in fewer unusable solutions. (default: 20)
#
#    --relay-set=[SET]
#        Set of transactions to relay. SET may be: cheap - Relay any transaction in the queue (this may include invalid
#        transactions); strict - Relay only executed transactions (this guarantees we don't relay invalid transactions,
#        but means we relay nothing if not mining); lenient - Same as strict when mining, and cheap when not. (default:
#        cheap)
#
#    --usd-per-tx=[USD]
#        Amount of USD to be paid for a basic transaction. The minimum gas price is set accordingly. (default: 0.0001)
#
#    --usd-per-eth=[SOURCE]
#        USD value of a single ETH. SOURCE may be either an amount in USD, a web service or 'auto' to use each web
#        service in turn and fallback on the last known good value. (default: auto)
#
#    --price-update-period=[T]
#        T will be allowed to pass between each gas price update. T may be daily, hourly, a number of seconds, or a time
#        string of the form "2 days", "30 minutes" etc.. (default: hourly)
#
#    --gas-floor-target=[GAS]
#        Amount of gas per block to target when sealing a new block. (default: 4700000)
#
#    --gas-cap=[GAS]
#        A cap on how large we will raise the gas limit per block due to transaction volume. (default: 6283184)
#
#    --tx-queue-mem-limit=[MB]
#        Maximum amount of memory that can be used by the transaction queue. Setting this parameter to 0 disables
#        limiting. (default: 4)
#
#    --tx-queue-size=[LIMIT]
#        Maximum amount of transactions in the queue (waiting to be included in next block). (default: 8192)
#
#    --tx-queue-per-sender=[LIMIT]
#        Maximum number of transactions per sender in the queue. By default it's 1% of the entire queue, but not less
#        than 16.
#
#    --tx-queue-gas=[LIMIT]
#        Maximum amount of total gas for external transactions in the queue. LIMIT can be either an amount of gas or
#        'auto' or 'off'. 'auto' sets the limit to be 20x the current block gas limit. (default: off)
#
#    --tx-queue-strategy=[S]
#        Prioritization strategy used to order transactions in the queue. S may be: gas_price - Prioritize txs with high
#        gas price (default: gas_price)
#
#    --stratum-interface=[IP]
#        Interface address for Stratum server. (default: local)
#
#    --stratum-port=[PORT]
#        Port for Stratum server to listen on. (default: 8008)
#
#    --min-gas-price=[STRING]
#        Minimum amount of Wei per GAS to be paid for a transaction to be accepted for mining. Overrides --usd-per-tx.
#
#    --gas-price-percentile=[PCT]
#        Set PCT percentile gas price value from last 100 blocks as default gas price when sending transactions.
#        (default: 50)
#
#    --author=[ADDRESS]
#        Specify the block author (aka "coinbase") address for sending block rewards from sealed blocks. NOTE: MINING
#        WILL NOT WORK WITHOUT THIS OPTION.
#
#    --engine-signer=[ADDRESS]
#        Specify the address which should be used to sign consensus messages and issue blocks. Relevant only to non-PoW
#        chains.
#
#    --tx-gas-limit=[GAS]
#        Apply a limit of GAS as the maximum amount of gas a single transaction may have for it to be mined.
#
#    --tx-time-limit=[MS]
#        Maximal time for processing single transaction. If enabled senders of transactions offending the limit will get
#        other transactions penalized.
#
#    --extra-data=[STRING]
#        Specify a custom extra-data for authored blocks, no more than 32 characters.
#
#    --notify-work=[URLS]
#        URLs to which work package notifications are pushed. URLS should be a comma-delimited list of HTTP URLs.
#
#    --stratum-secret=[STRING]
#        Secret for authorizing Stratum server for peers.
#
#Internal Options:
#    --can-restart
#        Executable will auto-restart if exiting with 69
#
#Miscellaneous options:
#    --no-color
#        Don't use terminal color codes in output.
#
#    -v, --version
#        Show information about version.
#
#    --no-config
#        Don't load a configuration file.
#
#    --ntp-servers=[HOSTS]
#        Comma separated list of NTP servers to provide current time (host:port). Used to verify node health. Parity uses
#        pool.ntp.org NTP servers; consider joining the pool: http://www.pool.ntp.org/join.html (default:
#        0.parity.pool.ntp.org:123,1.parity.pool.ntp.org:123,2.parity.pool.ntp.org:123,3.parity.pool.ntp.org:123)
#
#    -l, --logging=[LOGGING]
#        Specify the logging level. Must conform to the same format as RUST_LOG.
#
#    --log-file=[FILENAME]
#        Specify a filename into which logging should be appended.
#
#Footprint options:
#    --fast-and-loose
#        Disables DB WAL, which gives a significant speed up but means an unclean exit is unrecoverable.
#
#    --scale-verifiers
#        Automatically scale amount of verifier threads based on workload. Not guaranteed to be faster.
#
#    --tracing=[BOOL]
#        Indicates if full transaction tracing should be enabled. Works only if client had been fully synced with tracing
#        enabled. BOOL may be one of auto, on, off. auto uses last used value of this option (off if it does not exist).
#        (default: auto)
#
#    --pruning=[METHOD]
#        Configure pruning of the state/storage trie. METHOD may be one of auto, archive, fast: archive - keep all state
#        trie data. No pruning. fast - maintain journal overlay. Fast but 50MB used. auto - use the method most recently
#        synced or default to fast if none synced. (default: auto)
#
#    --pruning-history=[NUM]
#        Set a minimum number of recent states to keep in memory when pruning is active. (default: 64)
#
#    --pruning-memory=[MB]
#        The ideal amount of memory in megabytes to use to store recent states. As many states as possible will be kept
#        within this limit, and at least --pruning-history states will always be kept. (default: 32)
#
#    --cache-size-db=[MB]
#        Override database cache size. (default: 128)
#
#    --cache-size-blocks=[MB]
#        Specify the prefered size of the blockchain cache in megabytes. (default: 8)
#
#    --cache-size-queue=[MB]
#        Specify the maximum size of memory to use for block queue. (default: 40)
#
#    --cache-size-state=[MB]
#        Specify the maximum size of memory to use for the state cache. (default: 25)
#
#    --db-compaction=[TYPE]
#        Database compaction type. TYPE may be one of: ssd - suitable for SSDs and fast HDDs; hdd - suitable for slow
#        HDDs; auto - determine automatically. (default: auto)
#
#    --fat-db=[BOOL]
#        Build appropriate information to allow enumeration of all accounts and storage keys. Doubles the size of the
#        state database. BOOL may be one of on, off or auto. (default: auto)
#
#    --cache-size=[MB]
#        Set total amount of discretionary memory to use for the entire system, overrides other cache and queue options.
#
#    --num-verifiers=[INT]
#        Amount of verifier threads to use or to begin with, if verifier auto-scaling is enabled.
#
#Import/export options:
#    --no-seal-check
#        Skip block seal check.
#
#Snapshot options:
#    --no-periodic-snapshot
#        Disable automated snapshots which usually occur once every 10000 blocks.
#
#Whisper options:
#    --whisper
#        Enable the Whisper network.
#
#    --whisper-pool-size=[MB]
#        Target size of the whisper message pool in megabytes. (default: 10)
#
#Legacy options:
#    --warp
#        Does nothing; warp sync is enabled by default.
#
#    --dapps-apis-all
#        Dapps server is merged with RPC server. Use --jsonrpc-apis.
#
#    --geth
#        Run in Geth-compatibility mode. Sets the IPC path to be the same as Geth's. Overrides the --ipc-path and
#        --ipcpath options. Alters RPCs to reflect Geth bugs. Includes the personal_ RPC by default.
#
#    --testnet
#        Testnet mode. Equivalent to --chain testnet. Overrides the --keys-path option.
#
#    --import-geth-keys
#        Attempt to import keys from Geth client.
#
#    --ipcdisable
#        Equivalent to --no-ipc.
#
#    --ipc-off
#        Equivalent to --no-ipc.
#
#    --nodiscover
#        Equivalent to --no-discovery.
#
#    -j, --jsonrpc
#        Does nothing; JSON-RPC is on by default now.
#
#    --jsonrpc-off
#        Equivalent to --no-jsonrpc.
#
#    -w, --webapp
#        Does nothing; dapps server is on by default now.
#
#    --dapps-off
#        Equivalent to --no-dapps.
#
#    --rpc
#        Does nothing; JSON-RPC is on by default now.
#
#    --dapps-port=[PORT]
#        Dapps server is merged with RPC server. Use --jsonrpc-port.
#
#    --dapps-interface=[IP]
#        Dapps server is merged with RPC server. Use --jsonrpc-interface.
#
#    --dapps-hosts=[HOSTS]
#        Dapps server is merged with RPC server. Use --jsonrpc-hosts.
#
#    --dapps-cors=[URL]
#        Dapps server is merged with RPC server. Use --jsonrpc-cors.
#
#    --dapps-user=[USERNAME]
#        Dapps server authentication has been removed.
#
#    --dapps-pass=[PASSWORD]
#        Dapps server authentication has been removed.
#
#    --datadir=[PATH]
#        Equivalent to --base-path PATH.
#
#    --networkid=[INDEX]
#        Equivalent to --network-id INDEX.
#
#    --peers=[NUM]
#        Equivalent to --min-peers NUM.
#
#    --nodekey=[KEY]
#        Equivalent to --node-key KEY.
#
#    --rpcaddr=[IP]
#        Equivalent to --jsonrpc-interface IP.
#
#    --rpcport=[PORT]
#        Equivalent to --jsonrpc-port PORT.
#
#    --rpcapi=[APIS]
#        Equivalent to --jsonrpc-apis APIS.
#
#    --rpccorsdomain=[URL]
#        Equivalent to --jsonrpc-cors URL.
#
#    --ipcapi=[APIS]
#        Equivalent to --ipc-apis APIS.
#
#    --ipcpath=[PATH]
#        Equivalent to --ipc-path PATH.
#
#    --gasprice=[WEI]
#        Equivalent to --min-gas-price WEI.
#
#    --etherbase=[ADDRESS]
#        Equivalent to --author ADDRESS.
#
#    --extradata=[STRING]
#        Equivalent to --extra-data STRING.
#
#    --cache=[MB]
#        Equivalent to --cache-size MB.
#
#    --tx-queue-ban-count=[C]
#        Not supported. (default: 1)
#
#    --tx-queue-ban-time=[SEC]
#        Not supported. (default: 180)
