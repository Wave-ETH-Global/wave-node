pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

/**
 * @title Creature
 * Creature - a contract for my non-fungible creatures.
 */
contract WaveUserHandle is ERC721URIStorage {
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIds;    

    constructor()
        ERC721("WaveUserHandle", "WUH")
    {}

    string[] handles;

    mapping (string => string) name;
    mapping (string => bool) _exists;
    mapping (uint256 => string) tokens;

    function claimUserHandle(string memory userHandle, string memory uuid, string memory tokenURI) public returns (uint256) {
        require(!_exists[userHandle], "User handle exists!");

        name[userHandle] = uuid;
        _exists[userHandle] = true;
        handles.push(userHandle);
        uint256 newTokenId = _tokenIds.current();
        _mint(msg.sender, newTokenId);
        _setTokenURI(newTokenId, tokenURI);
        tokens[newTokenId] = userHandle;

        _tokenIds.increment();

        return newTokenId;
    }

    function getUUIDByUserHandle(string memory userHandle) public view returns (string memory) {
        return name[userHandle];
    }

    function getUserHandleByTokenId(uint256 tokenId) public view returns (string memory) {
        return tokens[tokenId];
    }
}