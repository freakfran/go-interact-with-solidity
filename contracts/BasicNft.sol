// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import {ERC721} from "lib/@openzeppelin/contracts/token/ERC721/ERC721.sol";
// errors--asm-json --ir-ast-json

/**
 * @title
 * @author ge1u
 * @notice
 */
contract BasicNft is ERC721 {
    // Type declarations

    // State variables
    uint256 private s_tokenCounter;
    mapping(uint256 => string) private s_tokenIdToUri;

    // Events

    // Modifiers

    // constructor
    constructor() ERC721("Dogie", "DOG") {
        s_tokenCounter = 0;
    }

    // receive function (if exists)
    // fallback function (if exists)

    // external

    // public
    function mintNft(string memory _tokenURI) public {
        s_tokenIdToUri[s_tokenCounter] = _tokenURI;
        //s_tokenCounter=>tokenId
        _safeMint(msg.sender, s_tokenCounter);
        s_tokenCounter++;
    }

    function tokenURI(uint256 tokenId) public view override returns (string memory) {
        return s_tokenIdToUri[tokenId];
    }

    // internal
    // private
    // internal & private view & pure functions
    // external & public view & pure functions
    function getTokenCounter() public view returns (uint256) {
        return s_tokenCounter;
    }
}
