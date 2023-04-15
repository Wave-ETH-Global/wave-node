import { ethers } from "hardhat";

async function main() {
  const WaveUserHandle = await ethers.getContractFactory("WaveUserHandle");
  const wuh = await WaveUserHandle.deploy();

  await wuh.deployed();

  console.log(
    `The contract was sucessfully deployed at address ${wuh.address}!`
  );
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
