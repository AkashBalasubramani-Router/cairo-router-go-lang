// contract ArrayDecoder {
//     function decodeArray(bytes[] memory input)
//         public
//         pure
//         returns (
//             string memory,
//             uint256,
//             string memory,
//             uint256,
//             bytes32
//         )
//     {
//         require(input.length == 7, "Input array length must be 7");

//         // Decode the elements based on their positions and types
//         string memory string1 = decodeHexString(input[0]);
//         uint256 uint1 = combineLowHighBytes32(input[1], input[2]);
//         string memory string2 = decodeHexString(input[3]);
//         uint256 uint2 = combineLowHighBytes32(input[4], input[5]);
//         bytes32 bytesValue = bytesToBytes32(input[6]);

//         return (string1, uint1, string2, uint2, bytesValue);
//     }

//     function decodeHexString(bytes memory hexString)
//         internal
//         pure
//         returns (string memory)
//     {
//         return string(hexString);
//     }

//     function combineLowHighBytes32(bytes memory low, bytes memory high)
//         public
//         pure
//         returns (uint256)
//     {
//         bytes32 lowBytesa32 = convertToBytes32(low);
//         bytes32 highBytesa32 = convertToBytes32(high);
//         uint256 lowBytes32 = bytes32ToUint256(lowBytesa32);
//         uint256 highBytes32 = bytes32ToUint256(highBytesa32);

//         uint256 combinedValue = (highBytes32 << 128) | lowBytes32;
//         return combinedValue;
//     }

  
//     function convertToBytes32(bytes memory data) public pure returns (bytes32) {
//         bytes memory padded = new bytes(32);

//         // Left pad data
//         for (uint256 i = 0; i < data.length; i++) {
//             padded[i + 32 - data.length] = data[i];
//         }

//         bytes32 result;

//         assembly {
//             result := mload(add(padded, 32))
//         }

//         return result;
//     }

//     function bytesToUint256(bytes memory data) public pure returns (uint256) {
//         require(data.length >= 32, "Data length must be at least 32 bytes");

//         uint256 result;
//         for (uint256 i = 0; i < 32; i++) {
//             result = (result << 8) | uint256(uint8(data[i]));
//         }

//         return result;
//     }

//     function bytes32ToUint256(bytes32 data) public pure returns (uint256) {
//         require(data.length >= 32, "Data length must be at least 32 bytes");

//         uint256 result;
//         for (uint256 i = 0; i < 32; i++) {
//             result = (result << 8) | uint256(uint8(data[i]));
//         }

//         return result;
//     }

//     function bytesToBytes32(bytes memory data) internal pure returns (bytes32) {
//         require(data.length <= 32, "Data length must be at most 32 bytes");

//         bytes32 result;
//         assembly {
//             result := mload(add(data, 32))
//         }
//         return result;
//     }
// }
