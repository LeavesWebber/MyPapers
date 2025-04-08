import Web3 from 'web3';
import { ERC721contractABI } from '../constant.js';

// 创建Web3实例
const web3 = new Web3(window.ethereum);

// 合约地址
const contractAddress = '0xCf3150B891e5176545c9EC8BfD2321bf13168848'; // 这里需要填入实际的合约地址

// 创建合约实例
const contractInstance = new web3.eth.Contract(ERC721contractABI, contractAddress);

export { contractInstance }; 