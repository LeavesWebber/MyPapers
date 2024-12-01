import Web3 from "web3";
const ERC721contractABI = [
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "to",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "tokenId",
                "type": "uint256"
            }
        ],
        "name": "approve",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "initialOwner",
                "type": "address"
            }
        ],
        "stateMutability": "nonpayable",
        "type": "constructor"
    },
    {
        "inputs": [],
        "name": "ERC721EnumerableForbiddenBatchMint",
        "type": "error"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "sender",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "tokenId",
                "type": "uint256"
            },
            {
                "internalType": "address",
                "name": "owner",
                "type": "address"
            }
        ],
        "name": "ERC721IncorrectOwner",
        "type": "error"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "operator",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "tokenId",
                "type": "uint256"
            }
        ],
        "name": "ERC721InsufficientApproval",
        "type": "error"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "approver",
                "type": "address"
            }
        ],
        "name": "ERC721InvalidApprover",
        "type": "error"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "operator",
                "type": "address"
            }
        ],
        "name": "ERC721InvalidOperator",
        "type": "error"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "owner",
                "type": "address"
            }
        ],
        "name": "ERC721InvalidOwner",
        "type": "error"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "receiver",
                "type": "address"
            }
        ],
        "name": "ERC721InvalidReceiver",
        "type": "error"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "sender",
                "type": "address"
            }
        ],
        "name": "ERC721InvalidSender",
        "type": "error"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "tokenId",
                "type": "uint256"
            }
        ],
        "name": "ERC721NonexistentToken",
        "type": "error"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "owner",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "index",
                "type": "uint256"
            }
        ],
        "name": "ERC721OutOfBoundsIndex",
        "type": "error"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "owner",
                "type": "address"
            }
        ],
        "name": "OwnableInvalidOwner",
        "type": "error"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "account",
                "type": "address"
            }
        ],
        "name": "OwnableUnauthorizedAccount",
        "type": "error"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "address",
                "name": "owner",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "approved",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "tokenId",
                "type": "uint256"
            }
        ],
        "name": "Approval",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "address",
                "name": "owner",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "operator",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "bool",
                "name": "approved",
                "type": "bool"
            }
        ],
        "name": "ApprovalForAll",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "_fromTokenId",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "_toTokenId",
                "type": "uint256"
            }
        ],
        "name": "BatchMetadataUpdate",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "_tokenId",
                "type": "uint256"
            }
        ],
        "name": "MetadataUpdate",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "address",
                "name": "previousOwner",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "newOwner",
                "type": "address"
            }
        ],
        "name": "OwnershipTransferred",
        "type": "event"
    },
    {
        "inputs": [],
        "name": "renounceOwnership",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "to",
                "type": "address"
            },
            {
                "internalType": "string",
                "name": "uri",
                "type": "string"
            }
        ],
        "name": "safeMint",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "from",
                "type": "address"
            },
            {
                "internalType": "address",
                "name": "to",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "tokenId",
                "type": "uint256"
            }
        ],
        "name": "safeTransferFrom",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "from",
                "type": "address"
            },
            {
                "internalType": "address",
                "name": "to",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "tokenId",
                "type": "uint256"
            },
            {
                "internalType": "bytes",
                "name": "data",
                "type": "bytes"
            }
        ],
        "name": "safeTransferFrom",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "operator",
                "type": "address"
            },
            {
                "internalType": "bool",
                "name": "approved",
                "type": "bool"
            }
        ],
        "name": "setApprovalForAll",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "address",
                "name": "from",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "to",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "tokenId",
                "type": "uint256"
            }
        ],
        "name": "Transfer",
        "type": "event"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "from",
                "type": "address"
            },
            {
                "internalType": "address",
                "name": "to",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "tokenId",
                "type": "uint256"
            }
        ],
        "name": "transferFrom",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "newOwner",
                "type": "address"
            }
        ],
        "name": "transferOwnership",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "owner",
                "type": "address"
            }
        ],
        "name": "balanceOf",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "tokenId",
                "type": "uint256"
            }
        ],
        "name": "getApproved",
        "outputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "owner",
                "type": "address"
            },
            {
                "internalType": "address",
                "name": "operator",
                "type": "address"
            }
        ],
        "name": "isApprovedForAll",
        "outputs": [
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "name",
        "outputs": [
            {
                "internalType": "string",
                "name": "",
                "type": "string"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "owner",
        "outputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "tokenId",
                "type": "uint256"
            }
        ],
        "name": "ownerOf",
        "outputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "bytes4",
                "name": "interfaceId",
                "type": "bytes4"
            }
        ],
        "name": "supportsInterface",
        "outputs": [
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "symbol",
        "outputs": [
            {
                "internalType": "string",
                "name": "",
                "type": "string"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "index",
                "type": "uint256"
            }
        ],
        "name": "tokenByIndex",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "owner",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "index",
                "type": "uint256"
            }
        ],
        "name": "tokenOfOwnerByIndex",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "tokenId",
                "type": "uint256"
            }
        ],
        "name": "tokenURI",
        "outputs": [
            {
                "internalType": "string",
                "name": "",
                "type": "string"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "totalSupply",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    }
];
export const ERC721contractAddress = "0x813Dfc691e771bDbAD90B68ABb584457FA69D8a9";
const web3 = new Web3(window.ethereum);
export const ERC721contractInstance = new web3.eth.Contract(
    ERC721contractABI,
    ERC721contractAddress
);

// 没加审稿内容的
// const ERC20contractABI = [
//     {
//         inputs: [
//             {
//                 internalType: "uint256",
//                 name: "initialSupply",
//                 type: "uint256",
//             },
//         ],
//         stateMutability: "nonpayable",
//         type: "constructor",
//     },
//     {
//         inputs: [
//             {
//                 internalType: "address",
//                 name: "spender",
//                 type: "address",
//             },
//             {
//                 internalType: "uint256",
//                 name: "allowance",
//                 type: "uint256",
//             },
//             {
//                 internalType: "uint256",
//                 name: "needed",
//                 type: "uint256",
//             },
//         ],
//         name: "ERC20InsufficientAllowance",
//         type: "error",
//     },
//     {
//         inputs: [
//             {
//                 internalType: "address",
//                 name: "sender",
//                 type: "address",
//             },
//             {
//                 internalType: "uint256",
//                 name: "balance",
//                 type: "uint256",
//             },
//             {
//                 internalType: "uint256",
//                 name: "needed",
//                 type: "uint256",
//             },
//         ],
//         name: "ERC20InsufficientBalance",
//         type: "error",
//     },
//     {
//         inputs: [
//             {
//                 internalType: "address",
//                 name: "approver",
//                 type: "address",
//             },
//         ],
//         name: "ERC20InvalidApprover",
//         type: "error",
//     },
//     {
//         inputs: [
//             {
//                 internalType: "address",
//                 name: "receiver",
//                 type: "address",
//             },
//         ],
//         name: "ERC20InvalidReceiver",
//         type: "error",
//     },
//     {
//         inputs: [
//             {
//                 internalType: "address",
//                 name: "sender",
//                 type: "address",
//             },
//         ],
//         name: "ERC20InvalidSender",
//         type: "error",
//     },
//     {
//         inputs: [
//             {
//                 internalType: "address",
//                 name: "spender",
//                 type: "address",
//             },
//         ],
//         name: "ERC20InvalidSpender",
//         type: "error",
//     },
//     {
//         anonymous: false,
//         inputs: [
//             {
//                 indexed: true,
//                 internalType: "address",
//                 name: "owner",
//                 type: "address",
//             },
//             {
//                 indexed: true,
//                 internalType: "address",
//                 name: "spender",
//                 type: "address",
//             },
//             {
//                 indexed: false,
//                 internalType: "uint256",
//                 name: "value",
//                 type: "uint256",
//             },
//         ],
//         name: "Approval",
//         type: "event",
//     },
//     {
//         anonymous: false,
//         inputs: [
//             {
//                 indexed: true,
//                 internalType: "address",
//                 name: "from",
//                 type: "address",
//             },
//             {
//                 indexed: false,
//                 internalType: "uint256",
//                 name: "value",
//                 type: "uint256",
//             },
//         ],
//         name: "Burn",
//         type: "event",
//     },
//     {
//         anonymous: false,
//         inputs: [
//             {
//                 indexed: false,
//                 internalType: "string",
//                 name: "hash",
//                 type: "string",
//             },
//             {
//                 indexed: true,
//                 internalType: "address",
//                 name: "sender",
//                 type: "address",
//             },
//         ],
//         name: "HashStored",
//         type: "event",
//     },
//     {
//         anonymous: false,
//         inputs: [
//             {
//                 indexed: true,
//                 internalType: "address",
//                 name: "to",
//                 type: "address",
//             },
//             {
//                 indexed: false,
//                 internalType: "uint256",
//                 name: "value",
//                 type: "uint256",
//             },
//         ],
//         name: "Mint",
//         type: "event",
//     },
//     {
//         anonymous: false,
//         inputs: [
//             {
//                 indexed: true,
//                 internalType: "address",
//                 name: "from",
//                 type: "address",
//             },
//             {
//                 indexed: true,
//                 internalType: "address",
//                 name: "to",
//                 type: "address",
//             },
//             {
//                 indexed: false,
//                 internalType: "uint256",
//                 name: "value",
//                 type: "uint256",
//             },
//         ],
//         name: "Transfer",
//         type: "event",
//     },
//     {
//         anonymous: false,
//         inputs: [
//             {
//                 indexed: true,
//                 internalType: "address",
//                 name: "from",
//                 type: "address",
//             },
//             {
//                 indexed: true,
//                 internalType: "address",
//                 name: "to",
//                 type: "address",
//             },
//             {
//                 indexed: false,
//                 internalType: "uint256",
//                 name: "value",
//                 type: "uint256",
//             },
//         ],
//         name: "TransferWithHash",
//         type: "event",
//     },
//     {
//         inputs: [
//             {
//                 internalType: "address",
//                 name: "owner",
//                 type: "address",
//             },
//             {
//                 internalType: "address",
//                 name: "spender",
//                 type: "address",
//             },
//         ],
//         name: "allowance",
//         outputs: [
//             {
//                 internalType: "uint256",
//                 name: "",
//                 type: "uint256",
//             },
//         ],
//         stateMutability: "view",
//         type: "function",
//     },
//     {
//         inputs: [
//             {
//                 internalType: "address",
//                 name: "spender",
//                 type: "address",
//             },
//             {
//                 internalType: "uint256",
//                 name: "value",
//                 type: "uint256",
//             },
//         ],
//         name: "approve",
//         outputs: [
//             {
//                 internalType: "bool",
//                 name: "",
//                 type: "bool",
//             },
//         ],
//         stateMutability: "nonpayable",
//         type: "function",
//     },
//     {
//         inputs: [
//             {
//                 internalType: "address",
//                 name: "account",
//                 type: "address",
//             },
//         ],
//         name: "balanceOf",
//         outputs: [
//             {
//                 internalType: "uint256",
//                 name: "",
//                 type: "uint256",
//             },
//         ],
//         stateMutability: "view",
//         type: "function",
//     },
//     {
//         inputs: [
//             {
//                 internalType: "address",
//                 name: "account",
//                 type: "address",
//             },
//             {
//                 internalType: "uint256",
//                 name: "amount",
//                 type: "uint256",
//             },
//         ],
//         name: "burnFrom",
//         outputs: [],
//         stateMutability: "nonpayable",
//         type: "function",
//     },
//     {
//         inputs: [],
//         name: "decimals",
//         outputs: [
//             {
//                 internalType: "uint8",
//                 name: "",
//                 type: "uint8",
//             },
//         ],
//         stateMutability: "view",
//         type: "function",
//     },
//     {
//         inputs: [
//             {
//                 internalType: "string",
//                 name: "hash",
//                 type: "string",
//             },
//         ],
//         name: "getRecipientByHash",
//         outputs: [
//             {
//                 internalType: "address",
//                 name: "",
//                 type: "address",
//             },
//         ],
//         stateMutability: "view",
//         type: "function",
//     },
//     {
//         inputs: [
//             {
//                 internalType: "address[]",
//                 name: "toAddresses",
//                 type: "address[]",
//             },
//             {
//                 internalType: "uint256",
//                 name: "amount",
//                 type: "uint256",
//             },
//         ],
//         name: "mint",
//         outputs: [
//             {
//                 internalType: "bool",
//                 name: "",
//                 type: "bool",
//             },
//         ],
//         stateMutability: "nonpayable",
//         type: "function",
//     },
//     {
//         inputs: [],
//         name: "name",
//         outputs: [
//             {
//                 internalType: "string",
//                 name: "",
//                 type: "string",
//             },
//         ],
//         stateMutability: "view",
//         type: "function",
//     },
//     {
//         inputs: [
//             {
//                 internalType: "string",
//                 name: "hash",
//                 type: "string",
//             },
//         ],
//         name: "storeHash",
//         outputs: [],
//         stateMutability: "nonpayable",
//         type: "function",
//     },
//     {
//         inputs: [],
//         name: "symbol",
//         outputs: [
//             {
//                 internalType: "string",
//                 name: "",
//                 type: "string",
//             },
//         ],
//         stateMutability: "view",
//         type: "function",
//     },
//     {
//         inputs: [],
//         name: "totalSupply",
//         outputs: [
//             {
//                 internalType: "uint256",
//                 name: "",
//                 type: "uint256",
//             },
//         ],
//         stateMutability: "view",
//         type: "function",
//     },
//     {
//         inputs: [
//             {
//                 internalType: "address",
//                 name: "recipient",
//                 type: "address",
//             },
//             {
//                 internalType: "uint256",
//                 name: "amount",
//                 type: "uint256",
//             },
//         ],
//         name: "transfer",
//         outputs: [
//             {
//                 internalType: "bool",
//                 name: "",
//                 type: "bool",
//             },
//         ],
//         stateMutability: "nonpayable",
//         type: "function",
//     },
//     {
//         inputs: [
//             {
//                 internalType: "address",
//                 name: "from",
//                 type: "address",
//             },
//             {
//                 internalType: "address",
//                 name: "to",
//                 type: "address",
//             },
//             {
//                 internalType: "uint256",
//                 name: "value",
//                 type: "uint256",
//             },
//         ],
//         name: "transferFrom",
//         outputs: [
//             {
//                 internalType: "bool",
//                 name: "",
//                 type: "bool",
//             },
//         ],
//         stateMutability: "nonpayable",
//         type: "function",
//     },
// ];
// const ERC20contractAddress = "0x9064cC938Dd23e43d8DB22a4Fa5D8CbB356F5d5F";
// export const ERC20contractInstance = new web3.eth.Contract(
//     ERC20contractABI,
//     ERC20contractAddress
// );

const ERC20contractABI = [
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "initialSupply",
                "type": "uint256"
            }
        ],
        "stateMutability": "nonpayable",
        "type": "constructor"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "spender",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "allowance",
                "type": "uint256"
            },
            {
                "internalType": "uint256",
                "name": "needed",
                "type": "uint256"
            }
        ],
        "name": "ERC20InsufficientAllowance",
        "type": "error"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "sender",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "balance",
                "type": "uint256"
            },
            {
                "internalType": "uint256",
                "name": "needed",
                "type": "uint256"
            }
        ],
        "name": "ERC20InsufficientBalance",
        "type": "error"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "approver",
                "type": "address"
            }
        ],
        "name": "ERC20InvalidApprover",
        "type": "error"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "receiver",
                "type": "address"
            }
        ],
        "name": "ERC20InvalidReceiver",
        "type": "error"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "sender",
                "type": "address"
            }
        ],
        "name": "ERC20InvalidSender",
        "type": "error"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "spender",
                "type": "address"
            }
        ],
        "name": "ERC20InvalidSpender",
        "type": "error"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "address",
                "name": "owner",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "spender",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "value",
                "type": "uint256"
            }
        ],
        "name": "Approval",
        "type": "event"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "spender",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "value",
                "type": "uint256"
            }
        ],
        "name": "approve",
        "outputs": [
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "address",
                "name": "from",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "value",
                "type": "uint256"
            }
        ],
        "name": "Burn",
        "type": "event"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "account",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "amount",
                "type": "uint256"
            }
        ],
        "name": "burnFrom",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "string",
                "name": "hash",
                "type": "string"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            }
        ],
        "name": "HashStored",
        "type": "event"
    },
    {
        "inputs": [
            {
                "internalType": "address[]",
                "name": "toAddresses",
                "type": "address[]"
            },
            {
                "internalType": "uint256",
                "name": "amount",
                "type": "uint256"
            }
        ],
        "name": "mint",
        "outputs": [
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "address",
                "name": "to",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "value",
                "type": "uint256"
            }
        ],
        "name": "Mint",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "string",
                "name": "content",
                "type": "string"
            }
        ],
        "name": "ReviewStored",
        "type": "event"
    },
    {
        "inputs": [
            {
                "internalType": "string",
                "name": "hash",
                "type": "string"
            }
        ],
        "name": "storeHash",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "string",
                "name": "content",
                "type": "string"
            }
        ],
        "name": "storeReview",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "recipient",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "amount",
                "type": "uint256"
            }
        ],
        "name": "transfer",
        "outputs": [
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "address",
                "name": "from",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "to",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "value",
                "type": "uint256"
            }
        ],
        "name": "Transfer",
        "type": "event"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "from",
                "type": "address"
            },
            {
                "internalType": "address",
                "name": "to",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "value",
                "type": "uint256"
            }
        ],
        "name": "transferFrom",
        "outputs": [
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "address",
                "name": "from",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "to",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "value",
                "type": "uint256"
            }
        ],
        "name": "TransferWithHash",
        "type": "event"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "owner",
                "type": "address"
            },
            {
                "internalType": "address",
                "name": "spender",
                "type": "address"
            }
        ],
        "name": "allowance",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "account",
                "type": "address"
            }
        ],
        "name": "balanceOf",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "decimals",
        "outputs": [
            {
                "internalType": "uint8",
                "name": "",
                "type": "uint8"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "string",
                "name": "hash",
                "type": "string"
            }
        ],
        "name": "getRecipientByHash",
        "outputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "string",
                "name": "content",
                "type": "string"
            }
        ],
        "name": "getReviewByHash",
        "outputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "name",
        "outputs": [
            {
                "internalType": "string",
                "name": "",
                "type": "string"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "symbol",
        "outputs": [
            {
                "internalType": "string",
                "name": "",
                "type": "string"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "totalSupply",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    }
]
const ERC20contractAddress = "0x548746c5DE169B603f804aB8ABBCE133981f6B36";
export const ERC20contractInstance = new web3.eth.Contract(
    ERC20contractABI,
    ERC20contractAddress
);
const MarketplacecontractABI = [
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "_tokenId",
                "type": "uint256"
            }
        ],
        "name": "buy",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "_tokenId",
                "type": "uint256"
            }
        ],
        "name": "cancelOrder",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "_tokenId",
                "type": "uint256"
            },
            {
                "internalType": "uint256",
                "name": "_price",
                "type": "uint256"
            }
        ],
        "name": "changePrice",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "operator",
                "type": "address"
            },
            {
                "internalType": "address",
                "name": "from",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "tokenId",
                "type": "uint256"
            },
            {
                "internalType": "bytes",
                "name": "data",
                "type": "bytes"
            }
        ],
        "name": "onERC721Received",
        "outputs": [
            {
                "internalType": "bytes4",
                "name": "",
                "type": "bytes4"
            }
        ],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "_erc20",
                "type": "address"
            },
            {
                "internalType": "address",
                "name": "_erc721",
                "type": "address"
            }
        ],
        "stateMutability": "nonpayable",
        "type": "constructor"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "address",
                "name": "seller",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "address",
                "name": "buyer",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "tokenId",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "price",
                "type": "uint256"
            }
        ],
        "name": "Deal",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "address",
                "name": "seller",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "tokenId",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "price",
                "type": "uint256"
            }
        ],
        "name": "NewOrder",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "address",
                "name": "seller",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "tokenId",
                "type": "uint256"
            }
        ],
        "name": "OrderChancelled",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "address",
                "name": "seller",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "tokenId",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "previousPrice",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "price",
                "type": "uint256"
            }
        ],
        "name": "PriceChanged",
        "type": "event"
    },
    {
        "inputs": [],
        "name": "erc20",
        "outputs": [
            {
                "internalType": "contract IERC20",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "erc721",
        "outputs": [
            {
                "internalType": "contract IERC721",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "getAllNFTs",
        "outputs": [
            {
                "components": [
                    {
                        "internalType": "address",
                        "name": "seller",
                        "type": "address"
                    },
                    {
                        "internalType": "uint256",
                        "name": "tokenId",
                        "type": "uint256"
                    },
                    {
                        "internalType": "uint256",
                        "name": "price",
                        "type": "uint256"
                    }
                ],
                "internalType": "struct Market.Order[]",
                "name": "",
                "type": "tuple[]"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "getMyNFTs",
        "outputs": [
            {
                "components": [
                    {
                        "internalType": "address",
                        "name": "seller",
                        "type": "address"
                    },
                    {
                        "internalType": "uint256",
                        "name": "tokenId",
                        "type": "uint256"
                    },
                    {
                        "internalType": "uint256",
                        "name": "price",
                        "type": "uint256"
                    }
                ],
                "internalType": "struct Market.Order[]",
                "name": "",
                "type": "tuple[]"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "getOrderLength",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "name": "idToOrderIndex",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "_tokenId",
                "type": "uint256"
            }
        ],
        "name": "isListed",
        "outputs": [
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "name": "orderOfId",
        "outputs": [
            {
                "internalType": "address",
                "name": "seller",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "tokenId",
                "type": "uint256"
            },
            {
                "internalType": "uint256",
                "name": "price",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "name": "orders",
        "outputs": [
            {
                "internalType": "address",
                "name": "seller",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "tokenId",
                "type": "uint256"
            },
            {
                "internalType": "uint256",
                "name": "price",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "bytes",
                "name": "_bytes",
                "type": "bytes"
            },
            {
                "internalType": "uint256",
                "name": "_start",
                "type": "uint256"
            }
        ],
        "name": "toUint256",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "pure",
        "type": "function"
    }
]
export const MarketplacecontractAddress = "0xB44C3C0E60bB72C39730A6999aD378954E18DEeB";
export const MarketplacecontractInstance = new web3.eth.Contract(
    MarketplacecontractABI,
    MarketplacecontractAddress
);

export const ERC20contract =
    `
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

// ERC20代币合约
contract MPER is ERC20 {
    using Counters for Counters.Counter;
    Counters.Counter private _hashCounter;

    // 哈希值到账户地址的映射
    mapping(string => address) private _hashToAddress;
    // 审稿内容到账户地址的映射
    mapping(string => address) private _reviewToAddress;

    // 事件：转账
    event TransferWithHash(address indexed from, address indexed to, uint256 value);

    // 事件：铸币
    event Mint(address indexed to, uint256 value);

    // 事件：存储哈希值
    event HashStored(string hash, address indexed sender);

    // 事件：扣减币
    event Burn(address indexed from, uint256 value);

    // 事件：存储审稿内容
    event ReviewStored(string content);

    // 初始化函数，分配总量给初始用户
    constructor(uint256 initialSupply) ERC20("MPER", "MPER") {
        _mint(msg.sender, initialSupply);
    }

    // 获取哈希值对应的账户地址
    function getRecipientByHash(string memory hash) public view returns (address) {
        return _hashToAddress[hash];
    }

    // 转账函数
    function transfer(address recipient, uint256 amount) public override returns (bool) {
        require(recipient != address(0), "Invalid address");
        super.transfer(recipient, amount);
        emit TransferWithHash(msg.sender, recipient, amount);
        return true;
    }
    
    // 铸币函数
    // function mint(address to, uint256 amount) public returns (bool) {
    //     _mint(to, amount);
    //     emit Mint(to, amount);
    //     return true;
    // }
    function mint(address[] memory toAddresses, uint256 amount) public returns (bool) {
        for (uint256 i = 0; i < toAddresses.length; i++) {
            _mint(toAddresses[i], amount);
            emit Mint(toAddresses[i], amount);
        }
        return true;
    }


    // 存储哈希值到区块链
    function storeHash(string memory hash) public {
        // require(_hashToAddress[hash] == address(0), "Hash already stored");
        require(_hashToAddress[hash] == address(0) || _hashToAddress[hash] == msg.sender, "Hash already stored");
        _hashToAddress[hash] = msg.sender;
        emit HashStored(hash, msg.sender);
    }

    // 扣减币函数
    function burnFrom(address account, uint256 amount) public {
        _burn(account, amount);
        emit Burn(account, amount);
    }

    // 存储审稿内容
    function storeReview(string memory content) public {
        _reviewToAddress[content] = msg.sender;
        emit ReviewStored(content);
    }

    // 获取审稿内容的账户地址
    function getReviewByHash(string memory content) public view returns (address) {
        return _reviewToAddress[content];
    }
}
`

export const ERC721contract =
    `
// SPDX-License-Identifier: MIT
// Compatible with OpenZeppelin Contracts ^5.0.0
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MyNFT is ERC721, ERC721Enumerable, ERC721URIStorage, Ownable {
    uint256 private _nextTokenId;

    constructor(address initialOwner)
        ERC721("MyNFT", "MNT")
        Ownable(initialOwner)
    {}

    function safeMint(address to, string memory uri) public {
        uint256 tokenId = _nextTokenId++;
        _safeMint(to, tokenId);
        _setTokenURI(tokenId, uri);
    }

    // The following functions are overrides required by Solidity.

    function _update(address to, uint256 tokenId, address auth)
        internal
        override(ERC721, ERC721Enumerable)
        returns (address)
    {
        return super._update(to, tokenId, auth);
    }

    function _increaseBalance(address account, uint128 value)
        internal
        override(ERC721, ERC721Enumerable)
    {
        super._increaseBalance(account, value);
    }

    function tokenURI(uint256 tokenId)
        public
        view
        override(ERC721, ERC721URIStorage)
        returns (string memory)
    {
        return super.tokenURI(tokenId);
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721, ERC721Enumerable, ERC721URIStorage)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}
`

export const Marketcontract =
    `
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/interfaces/IERC20.sol";
import "@openzeppelin/contracts/interfaces/IERC721.sol";

contract Market {
    IERC20 public erc20;
    IERC721 public erc721;

    bytes4 internal constant MAGIC_ON_ERC721_RECEIVED = 0x150b7a02;
    struct Order {
        address seller;
        uint256 tokenId;
        uint256 price;
    }

    mapping(uint256 => Order) public orderOfId; // token id to order
    Order[] public orders;
    mapping(uint256 => uint256) public idToOrderIndex; // token id to index orders

    event Deal(address seller, address buyer, uint256 tokenId, uint256 price); // 成交
    event NewOrder(address seller, uint256 tokenId, uint256 price);
    event PriceChanged(
        address seller,
        uint256 tokenId,
        uint256 previousPrice,
        uint256 price
    );
    event OrderChancelled(address seller, uint256 tokenId);

    constructor(address _erc20, address _erc721) {
        require(_erc20 != address(0), "zero address");
        require(_erc721 != address(0), "zero address");
        erc20 = IERC20(_erc20);
        erc721 = IERC721(_erc721);
    }

    function buy(uint256 _tokenId) external {
        address seller = orderOfId[_tokenId].seller;
        address buyer = msg.sender;
        uint256 price = orderOfId[_tokenId].price;

        require(
            erc20.transferFrom(buyer, seller, price),
            "transfer not successful"
        );
        erc721.safeTransferFrom(address(this), buyer, _tokenId);

        removeOrder(_tokenId);
        emit Deal(seller, buyer, _tokenId, price);
    }

    function cancelOrder(uint256 _tokenId) external {
        address seller = orderOfId[_tokenId].seller;
        require(msg.sender == seller, "only seller can cancel order");
        erc721.safeTransferFrom(address(this), seller, _tokenId);
        
        removeOrder(_tokenId);

        emit OrderChancelled(seller, _tokenId);
    }

    function changePrice(uint256 _tokenId, uint256 _price) external {
        address seller = orderOfId[_tokenId].seller;
        require(msg.sender == seller, "only seller can change price");
        
        uint256 previousPrice = orderOfId[_tokenId].price;
        orderOfId[_tokenId].price = _price;

        Order storage order = orders[idToOrderIndex[_tokenId]];
        order.price = _price;

        emit PriceChanged(seller, _tokenId, previousPrice, _price);
    }

    // function createOrder(uint256 _tokenId, uint256 _price) external {
    //     address seller = msg.sender;
    //     erc721.safeTransferFrom(seller, address(this), _tokenId);
    //     Order memory order = Order(seller, _tokenId, _price);
    //     orders.push(order);
    //     orderOfId[_tokenId] = order;
    //     idToOrderIndex[_tokenId] = orders.length - 1;
    //     emit NewOrder(seller, _tokenId, _price);
    // }

    function isListed(uint256 _tokenId) public view returns (bool) {
        return orderOfId[_tokenId].seller != address(0);
    }

    function onERC721Received(
        address operator,
        address from,
        uint256 tokenId,
        bytes calldata data
    ) external returns (bytes4) {
        uint256 price = toUint256(data, 0);
        require(price > 0, "price must be greater than 0");

        orders.push(Order(from, tokenId, price));
        orderOfId[tokenId] = Order(from, tokenId, price);
        idToOrderIndex[tokenId] = orders.length - 1;
        emit NewOrder(from, tokenId, price);
        return MAGIC_ON_ERC721_RECEIVED;
    }

    function removeOrder(uint256 _tokenId) internal {
        uint256 index = idToOrderIndex[_tokenId];
        uint256 lastIndex = orders.length - 1;
        if (index != lastIndex) {
            Order storage lastOrder = orders[lastIndex];
            orders[index] = lastOrder;
            idToOrderIndex[lastOrder.tokenId] = index;
        }
        orders.pop();
        delete orderOfId[_tokenId];
        delete idToOrderIndex[_tokenId];
    }
    // how-to-use-bytestouint-function-in-solidity-the-one-with-assembly
    function toUint256(
        bytes memory _bytes,
        uint256 _start
    ) public pure returns (uint256) {
        require(_start + 32 >= _start, "Market: toUint256_overflow");
        require(_bytes.length >= _start + 32, "Market: toUint256_outOfBounds");
        uint256 tempUint;

        assembly {
            tempUint := mload(add(add(_bytes, 0x20), _start))
        }
        return tempUint;
    }

    function getOrderLength() external view returns (uint256) {
        return orders.length;
    }

    function getAllNFTs() external view returns (Order[] memory) {
        return orders;
    }

    function getMyNFTs() external view returns (Order[] memory) {
        Order[] memory myOrders = new Order[](orders.length);
        uint256 count = 0;
        for (uint256 i = 0; i < orders.length; i++) {
            if (orders[i].seller == msg.sender) {
                myOrders[count] = orders[i];
                count++;
            }
        }
        return myOrders;
    }
}
`